// Start fileheader template
// Code generated by go generate; DO NOT EDIT.
// This file was generated by FactomGenerate robots

// Start Generated Code

package generated

import (
	"github.com/FactomProject/factomd/common"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/telemetry"
)

// End fileheader template

// Start accountedqueue generated go code

type Queue_IMsg struct {
	common.Name
	Channel chan interfaces.IMsg
}

func (q *Queue_IMsg) Init(parent common.NamedObject, name string, size int) *Queue_IMsg {
	q.Name.Init(parent, name)
	q.Channel = make(chan interfaces.IMsg, size)
	return q
}

// construct gauge w/ proper labels
func (q *Queue_IMsg) Metric() telemetry.Gauge {
	return telemetry.ChannelSize.WithLabelValues(q.GetPath(), "current")
}

// construct counter for tracking totals
func (q *Queue_IMsg) TotalMetric() telemetry.Counter {
	return telemetry.TotalCounter.WithLabelValues(q.GetPath(), "total")
}

// Length of underlying channel
func (q Queue_IMsg) Length() int {
	return len(q.Channel)
}

// Cap of underlying channel
func (q Queue_IMsg) Cap() int {
	return cap(q.Channel)
}

// Enqueue adds item to channel and instruments based on type
func (q Queue_IMsg) Enqueue(m interfaces.IMsg) {
	q.Channel <- m
	q.TotalMetric().Inc()
	q.Metric().Inc()
}

// Enqueue adds item to channel and instruments based on
// returns true it it enqueues the data
func (q Queue_IMsg) EnqueueNonBlocking(m interfaces.IMsg) bool {
	select {
	case q.Channel <- m:
		q.TotalMetric().Inc()
		q.Metric().Inc()
		return true
	default:
		return false
	}
}

// Dequeue removes an item from channel
// Returns nil if nothing in // queue
func (q Queue_IMsg) Dequeue() interfaces.IMsg {
	select {
	case v := <-q.Channel:
		q.Metric().Dec()
		return v
	default:
		return nil
	}
}

// Dequeue removes an item from channel
func (q Queue_IMsg) BlockingDequeue() interfaces.IMsg {
	v := <-q.Channel
	q.Metric().Dec()
	return v
}

// End accountedqueue generated go code
//
// Start accountedqueue generated go code

type Queue_int struct {
	common.Name
	Channel chan int
}

func (q *Queue_int) Init(parent common.NamedObject, name string, size int) *Queue_int {
	q.Name.Init(parent, name)
	q.Channel = make(chan int, size)
	return q
}

// construct gauge w/ proper labels
func (q *Queue_int) Metric() telemetry.Gauge {
	return telemetry.ChannelSize.WithLabelValues(q.GetPath(), "current")
}

// construct counter for tracking totals
func (q *Queue_int) TotalMetric() telemetry.Counter {
	return telemetry.TotalCounter.WithLabelValues(q.GetPath(), "total")
}

// Length of underlying channel
func (q Queue_int) Length() int {
	return len(q.Channel)
}

// Cap of underlying channel
func (q Queue_int) Cap() int {
	return cap(q.Channel)
}

// Enqueue adds item to channel and instruments based on type
func (q Queue_int) Enqueue(m int) {
	q.Channel <- m
	q.TotalMetric().Inc()
	q.Metric().Inc()
}

// Enqueue adds item to channel and instruments based on
// returns true it it enqueues the data
func (q Queue_int) EnqueueNonBlocking(m int) bool {
	select {
	case q.Channel <- m:
		q.TotalMetric().Inc()
		q.Metric().Inc()
		return true
	default:
		return false
	}
}

// Dequeue removes an item from channel
// Returns nil if nothing in // queue
func (q Queue_int) Dequeue() int {
	select {
	case v := <-q.Channel:
		q.Metric().Dec()
		return v
	default:
		return nil
	}
}

// Dequeue removes an item from channel
func (q Queue_int) BlockingDequeue() int {
	v := <-q.Channel
	q.Metric().Dec()
	return v
}

// End accountedqueue generated go code
//
// Start filetail template
// Code generated by go generate; DO NOT EDIT.
// End filetail template
// End Generated Code
