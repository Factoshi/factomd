package msgorder

import (
	"github.com/FactomProject/factomd/common"
	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/modules/event"
	"github.com/FactomProject/factomd/pubsub"
	"github.com/FactomProject/factomd/state"
	"github.com/FactomProject/factomd/worker"
)

type Handler struct {
	Pub
	Sub
	*Events
	exit chan interface{}
	//ticker  chan interface{}
	//logfile string
}

func New(nodeName string) *Handler {
	v := new(Handler)
	v.Events = &Events{
		DBHT:   nil,
		Ack:    nil,
		Config: &event.LeaderConfig{NodeName: nodeName}, // FIXME should use pubsub.Config
	}
	v.exit = make(chan interface{})
}

type Pub struct {
	UnAck pubsub.IPublisher
}

// isolate deps on state package - eventually functions will be relocated
var GetFedServerIndexHash = state.GetFedServerIndexHash

// create and start all publishers
func (p *Pub) Init(nodeName string) {
	p.UnAck = pubsub.PubFactory.Threaded(100).Publish(
		pubsub.GetPath(nodeName, event.Path.UnAckMsgs),
	)
	go p.UnAck.Start()
}

type Sub struct {
	MsgInput      *pubsub.SubChannel
	MovedToHeight *pubsub.SubChannel
}

// Create all subscribers
func (s *Sub) Init() {
	s.MovedToHeight = pubsub.SubFactory.Channel(1000)
	s.MsgInput = pubsub.SubFactory.Channel(1000)
}

// start subscriptions
func (s *Sub) Start(nodeName string) {
	s.MovedToHeight.Subscribe(pubsub.GetPath(nodeName, event.Path.DBHT))
	s.MsgInput.Subscribe(pubsub.GetPath(nodeName, event.Path.BVM))
}

type Events struct {
	*event.DBHT                     // from move-to-ht
	*event.Ack                      // record of last sent ack by leader
	Config      *event.LeaderConfig // FIXME: use pubsub.Config obj
}

func (h *Handler) Start(w *worker.Thread) {
	w.Spawn("MsgOrderThread", func(w *worker.Thread) {
		w.OnReady(func() {
			go h.waitForEOM()
		})
		w.OnRun(h.Run)
		w.OnExit(func() {
			close(h.exit)
			h.Pub.UnAck.Close()
		})
		h.Pub.Init(h.Config.NodeName)
		h.Sub.Init()
	})
}

func (h *Handler) Run() {
runLoop:
	for {
		select {
		case v := <-h.MsgInput.Updates:
			m := v.(interfaces.IMsg)
			if constants.NeedsAck(m.Type()) {
				// FIXME: match ACK/Reveals
			}
		case v := <-h.MovedToHeight.Updates:
			evt := v.(*event.DBHT)

			if evt.Minute == 10 {
				continue // skip min 10
			}

			if h.DBHT.Minute == evt.Minute && h.DBHT.DBHeight == evt.DBHeight {
				continue // skip duplicates
			}

			h.DBHT = evt

			// TODO: send UnAcked messages to leader
			continue runLoop
		case <-h.exit:
			return
		}
	}
}

type heldMessage struct {
	dependentHash [32]byte
	offset        int
}

type HoldingList struct {
	common.Name
	holding    map[[32]byte][]interfaces.IMsg
	dependents map[[32]byte]heldMessage // used to avoid duplicate entries & track position in holding
}