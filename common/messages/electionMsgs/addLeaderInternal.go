// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package electionMsgs

import (
	"fmt"
	"reflect"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/messages/msgbase"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/state"

	llog "github.com/FactomProject/factomd/log"
	log "github.com/sirupsen/logrus"
)

//General acknowledge message
type AddLeaderInternal struct {
	msgbase.MessageBase
	NName    string
	ServerID interfaces.IHash // Hash of message acknowledged
	DBHeight uint32           // Directory Block Height that owns this ack
	Height   uint32           // Height of this ack in this process list
	//	MessageHash interfaces.IHash
}

var _ interfaces.IMsg = (*AddLeaderInternal)(nil)
var _ interfaces.IElectionMsg = (*AddLeaderInternal)(nil)

func (m *AddLeaderInternal) MarshalBinary() (data []byte, err error) {
	var buf primitives.Buffer

	if err = buf.PushByte(constants.INTERNALADDLEADER); err != nil {
		return nil, err
	}
	if e := buf.PushIHash(m.ServerID); e != nil {
		return nil, e
	}
	if e := buf.PushInt(int(m.DBHeight)); e != nil {
		return nil, e
	}
	if e := buf.PushByte(m.Minute); e != nil {
		return nil, e
	}
	if e := buf.PushByte(m.Minute); e != nil {
		return nil, e
	}
	data = buf.Bytes()
	return data, nil
}

func (m *AddLeaderInternal) GetMsgHash() (rval interfaces.IHash) {
	defer func() {
		if rval != nil && reflect.ValueOf(rval).IsNil() {
			rval = nil // convert an interface that is nil to a nil interface
			primitives.LogNilHashBug("AddLeaderInternal.GetMsgHash() saw an interface that was nil")
		}
	}()

	if m.MsgHash == nil {
		data, err := m.MarshalBinary()
		if err != nil {
			return nil
		}
		m.MsgHash = primitives.Sha(data)
	}
	return m.MsgHash
}

func (m *AddLeaderInternal) ElectionProcess(s interfaces.IState, e interfaces.IElections) {
	e.AddLeader(&state.Server{ChainID: m.ServerID, Online: true})
}

func (m *AddLeaderInternal) GetServerID() (rval interfaces.IHash) {
	defer func() {
		if rval != nil && reflect.ValueOf(rval).IsNil() {
			rval = nil // convert an interface that is nil to a nil interface
			primitives.LogNilHashBug("AddLeaderInternal.GetServerID() saw an interface that was nil")
		}
	}()

	return m.ServerID
}

func (m *AddLeaderInternal) LogFields() log.Fields {
	return log.Fields{"category": "message", "messagetype": "addleaderinternal", "dbheight": m.DBHeight, "newleader": m.ServerID.String()[4:12]}
}

func (m *AddLeaderInternal) GetRepeatHash() (rval interfaces.IHash) {
	defer func() {
		if rval != nil && reflect.ValueOf(rval).IsNil() {
			rval = nil // convert an interface that is nil to a nil interface
			primitives.LogNilHashBug("AddLeaderInternal.GetRepeatHash() saw an interface that was nil")
		}
	}()

	return m.GetMsgHash()
}

// We have to return the hash of the underlying message.
func (m *AddLeaderInternal) GetHash() (rval interfaces.IHash) {
	defer func() {
		if rval != nil && reflect.ValueOf(rval).IsNil() {
			rval = nil // convert an interface that is nil to a nil interface
			primitives.LogNilHashBug("AddLeaderInternal.GetHash() saw an interface that was nil")
		}
	}()

	return m.GetMsgHash()
}

func (m *AddLeaderInternal) GetTimestamp() interfaces.Timestamp {
	return primitives.NewTimestampNow()
}

func (m *AddLeaderInternal) Type() byte {
	return constants.INTERNALADDLEADER
}

func (m *AddLeaderInternal) ElectionValidate(ie interfaces.IElections) int {
	return 1
}

func (m *AddLeaderInternal) WellFormed() bool {
	// TODO: Flush this out
	return true
}

func (m *AddLeaderInternal) Validate(state interfaces.IState) int {
	return 1
}

// Returns true if this is a message for this server to execute as
// a leader.
func (m *AddLeaderInternal) ComputeVMIndex(state interfaces.IState) {
}

// Execute the leader functions of the given message
// Leader, follower, do the same thing.
func (m *AddLeaderInternal) LeaderExecute(state interfaces.IState) {
	m.FollowerExecute(state)
}

func (m *AddLeaderInternal) FollowerExecute(state interfaces.IState) {

}

// Acknowledgements do not go into the process list.
func (m *AddLeaderInternal) Process(dbheight uint32, state interfaces.IState) bool {
	panic("Ack object should never have its Process() method called")
}

func (m *AddLeaderInternal) JSONByte() ([]byte, error) {
	return primitives.EncodeJSON(m)
}

func (m *AddLeaderInternal) JSONString() (string, error) {
	return primitives.EncodeJSONString(m)
}

func (m *AddLeaderInternal) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling: %v", r)
			llog.LogPrintf("recovery", "Error unmarshalling: %v", r)
		}
	}()
	return
}

func (m *AddLeaderInternal) UnmarshalBinary(data []byte) error {
	_, err := m.UnmarshalBinaryData(data)
	return err
}

func (m *AddLeaderInternal) String() string {
	if m.LeaderChainID == nil {
		m.LeaderChainID = primitives.NewZeroHash()
	}
	return fmt.Sprintf(" %10s %20s %x dbheight %5d", m.NName, "Add Leader Internal", m.ServerID.Bytes(), m.DBHeight)
}

func (m *AddLeaderInternal) IsSameAs(b *AddLeaderInternal) bool {
	return true
}

func (m *AddLeaderInternal) Label() string {
	return msgbase.GetLabel(m)
}
