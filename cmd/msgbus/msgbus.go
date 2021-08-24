package msgbus

import (
	"fmt"
	"os"
	"time"
)

type operation int
type EventType int
type MsgType int
type ContractState int
type MinerState int
type ConnectionState int

const (
	nop operation = iota
	sub
	pub
	get
	set
	unsub
	unpub
	shutdown
)

const (
	NoEvent EventType = iota
	UpdateEvent
	DeleteEvent
	GetEvent
	GetIndexEvent
	SubscribedEvent
)

const (
	NoMsg MsgType = iota
	ConfigMsg
	DestMsg
	SellerMsg
	ContractMsg
	MinerMsg
	ConnectionMsg
	LogMsg
)

const (
	NewState ContractState = iota
	ReadyState
	ActiveState
	CompleteState
)

const (
	OnlineState MinerState = iota
	OfflineState
)

const (
	NoState ConnectionState = iota
	SrcOpenState
	AuthState
	VerifyState
	RoutingState
	ConnectingState
	ConnectedState
	ConnectErrState
	MsgErrState
	RouteChangeState
	DstCloseState
	SrcCloseState
	ShutdownState
	ClosedState
)

// Need to figure out the IDString for this, for now it is just a string
type IDString string
type ConfigID IDString
type DestID IDString
type SellerID IDString
type BuyerID IDString
type ContractID IDString
type MinerID IDString
type ConnectionID IDString

type Event struct {
	EventType EventType
	Msg       MsgType
	ID        IDString
	Data      interface{}
}

type EventChan chan Event

type Subscribers struct {
	eventchan map[EventChan]int
}

type registryData struct {
	sub  Subscribers
	data interface{}
}

type Dest struct {
	ID   DestID
	IP   string
	Port int
}

type ConfigInfo struct {
	ID          ConfigID
	DefaultDest DestID
	Seller      SellerID
}

type Seller struct {
	ID                     SellerID
	DefaultDest            DestID
	TotalAvailableHashRate int
	UnusedHashRate         int
	NewContracts           map[ContractID]bool
	ReadyContracts         map[ContractID]bool
	ActiveContracts        map[ContractID]bool
}

type Contract struct {
	ID               ContractID
	State            ContractState
	Buyer            BuyerID
	Dest             DestID
	CommitedHashRate int
	TargetHashRate   int
	CurrentHashRate  int
	Tolerance        int
	Penalty          int
	Priority         int
	StartDate        int
	EndDate          int
}

type Miner struct {
	ID                      MinerID
	State                   MinerState
	Seller                  SellerID
	Dest                    DestID
	InitialMeasuredHashRate int
	CurrentHashRate         int
}

type Connection struct {
	ID        ConnectionID
	Miner     MinerID
	Dest      DestID
	State     ConnectionState
	TotalHash int
	StartDate time.Time
}

type registry struct {
	// data   map[MsgType]map[IDString]interface{}
	data   map[MsgType]map[IDString]registryData
	notify map[MsgType]map[chan Event]interface{}
}

// PubSub is a collection of topics.
type PubSub struct {
	cmdChan  chan *cmd
	capacity int
}

type cmd struct {
	op      operation
	msg     MsgType
	ID      IDString
	data    interface{}
	eventch EventChan
	err     error
	retch   chan int
}

// New creates a new PubSub and starts a goroutine for handling operations.
// The capacity of the channels created by Sub and SubOnce will be as specified.
//
// Convert this to a message bus
//
func New(capacity int) *PubSub {
	ps := &PubSub{make(chan *cmd), capacity}
	go ps.start()
	return ps
}

func (ps *PubSub) NewEventChanPtr() *EventChan {
	ech := ps.NewEventChan()
	return &ech
}

func (ps *PubSub) NewEventChan() EventChan {
	return make(EventChan)
}

func (ps *PubSub) NewEventPtr() *Event {
	e := ps.NewEvent()
	return &e
}

func (ps *PubSub) NewEvent() Event {
	e := Event{EventType: NoEvent}
	return e
}

// Create new topic structure
func (ps *PubSub) Pub(msg MsgType, id IDString, data interface{}) (err error) {

	c := cmd{
		op:      pub,
		msg:     msg,
		ID:      id,
		data:    data,
		eventch: nil,
		err:     nil,
	}

	ps.dispatch(&c)

	return c.err
}

// Request update events for the topic
func (ps *PubSub) Sub(msg MsgType, id IDString, ech EventChan) error {

	c := cmd{
		op:      sub,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: ech,
		err:     nil,
	}

	ps.dispatch(&c)

	return c.err
}

//
func (ps *PubSub) Get(msg MsgType, id IDString, ech EventChan) error {

	c := cmd{
		op:      get,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: ech,
		err:     nil,
	}

	ps.dispatch(&c)

	return c.err

}

//
func (ps *PubSub) Set(msg MsgType, id IDString, data interface{}) error {

	c := cmd{
		op:      set,
		msg:     msg,
		ID:      id,
		data:    data,
		eventch: nil,
		err:     nil,
	}

	ps.dispatch(&c)

	return c.err
}

// Request update events for the topic
func (ps *PubSub) Unsub(msg MsgType, id IDString, ech EventChan) error {

	c := cmd{
		op:      unsub,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: ech,
		err:     nil,
	}

	ps.dispatch(&c)

	return c.err
}

// Request update events for the topic
func (ps *PubSub) Unpub(msg MsgType, id IDString) error {

	c := cmd{
		op:      unpub,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: nil,
		err:     nil,
	}

	ps.dispatch(&c)

	return c.err
}

//
// Shutdown closes all subscribed channels and terminates the goroutine.
//
func (ps *PubSub) Shutdown() {

	c := cmd{
		op:      shutdown,
		msg:     NoMsg,
		ID:      "",
		data:    nil,
		eventch: nil,
		err:     nil,
	}

	ps.dispatch(&c)
}

//
//
//
func (ps *PubSub) dispatch(c *cmd) {
	c.retch = make(chan int, 0)

	ps.cmdChan <- c

	// Wait for a response
	<-c.retch

	fmt.Printf("dispatch complete\n")
}

//
//
//
func (ps *PubSub) start() {
	reg := registry{
		// data:   make(map[MsgType]map[IDString]interface{}),
		data:   make(map[MsgType]map[IDString]registryData),
		notify: make(map[MsgType]map[chan Event]interface{}),
	}

loop:
	for cmdptr := range ps.cmdChan {

		fmt.Printf("New Command %+v\n", *cmdptr)

		if cmdptr.op == nop {
			cmdptr.err = nil
			close(cmdptr.retch)
			continue loop
		}

		if cmdptr.op == shutdown {
			close(cmdptr.retch)
			break loop
		}

		switch cmdptr.op {
		case pub:
			reg.pub(cmdptr)

		case sub:
			reg.sub(cmdptr)

		case set:
			reg.set(cmdptr)

		case get:
			reg.get(cmdptr)

		case unpub:
			reg.unpub(cmdptr)

		case unsub:
			reg.unsub(cmdptr)

		default:
			panic("default reached for cmd.op")
		}

		close(cmdptr.retch)
	}

	// clean up here
	// Close all open channels
	// Delete registry

}

func (reg *registry) noMsg(c *cmd) bool {
	if c.msg == NoMsg {
		c.err = fmt.Errorf("NOMSG")
		return true
	}
	return false
}

//
// msg contains the message type
// ID contains the new value
// data contains the new data struct
//
func (reg *registry) pub(c *cmd) {

	if reg.noMsg(c) {
		return
	}

	if c.ID == "" {
		c.err = fmt.Errorf("NOID")
		return
	}

	if c.data == nil {
		c.err = fmt.Errorf("NODATA")
		return
	}

	if _, ok := reg.data[c.msg]; !ok {
		reg.data[c.msg] = make(map[IDString]registryData)
	}

	if _, ok := reg.data[c.msg][c.ID]; ok {
		c.err = fmt.Errorf("DUPDATA")
		return
	}

	reg.data[c.msg][c.ID] = registryData{
		sub:  Subscribers{eventchan: make(map[EventChan]int)},
		data: c.data,
	}

}

//
// msg
// ID (optional)
// ch Event Channel
//
func (reg *registry) sub(c *cmd) {

	if reg.noMsg(c) {
		return
	}

	if c.eventch == nil {
		c.err = fmt.Errorf("NOEVENTCH")
		return
	}

	if c.ID == "" {

		if _, ok := reg.notify[c.msg]; !ok {
			c.err = fmt.Errorf("MSG-DNE")
			return
		}

		if _, ok := reg.notify[c.msg][c.eventch]; ok {
			c.err = fmt.Errorf("DUPDATA")
			return
		}

		reg.notify[c.msg][c.eventch] = 1
		return
	}

	if _, ok := reg.data[c.msg]; !ok {
		c.err = fmt.Errorf("MSG-DNE")
		return
	}

	if _, ok := reg.data[c.msg][c.ID]; !ok {
		c.err = fmt.Errorf("ID-DNE")
		return
	}

	if _, ok := reg.data[c.msg][c.ID].sub.eventchan[c.eventch]; ok {
		c.err = fmt.Errorf("DUPDATA")
		return
	}

	reg.data[c.msg][c.ID].sub.eventchan[c.eventch] = 1

	go func(e EventChan) {

		event := Event{
			EventType: SubscribedEvent,
			Msg:       c.msg,
			ID:        c.ID,
			Data:      nil,
		}

		e <- event

	}(c.eventch)
}

//
//
//
func (reg *registry) set(c *cmd) {

	if reg.noMsg(c) {
		return
	}

	if c.ID == "" {
		c.err = fmt.Errorf("NOID")
		return
	}

	if _, ok := reg.data[c.msg]; !ok {
		c.err = fmt.Errorf("MSG-DNE")
		return
	}

	if _, ok := reg.data[c.msg][c.ID]; !ok {
		c.err = fmt.Errorf("ID-DNE")
		return
	}

	// Could do a lot here to check if the data actually changed
	// Set the data

	d := reg.data[c.msg][c.ID]
	d.data = c.data
	reg.data[c.msg][c.ID] = d

	// Notify anyone listening
	for eventch, _ := range reg.data[c.msg][c.ID].sub.eventchan {

		go func(e EventChan) {

			event := Event{
				EventType: UpdateEvent,
				Msg:       c.msg,
				ID:        c.ID,
				Data:      c.data,
			}

			e <- event

		}(eventch)
	}

}

//
// Msg
// ID
// eventch - where to send the get request too
//
func (reg *registry) get(c *cmd) {
	if reg.noMsg(c) {
		return
	}

	if c.eventch == nil {
		c.err = fmt.Errorf("NOEVENTCH")
		return
	}

	if _, ok := reg.data[c.msg]; !ok {
		c.err = fmt.Errorf("MSG-DNE")
		return
	}

	// Get Array of ID for Msg
	if c.ID == "" {
		ids := make([]IDString, len(reg.data[c.msg]))

		for i, _ := range reg.data[c.msg] {
			ids = append(ids, i)
		}

		event := Event{
			EventType: GetIndexEvent,
			Msg:       c.msg,
			ID:        "",
			Data:      ids,
		}

		c.eventch <- event
		return
	}

	if rd, ok := reg.data[c.msg][c.ID]; !ok {
		c.err = fmt.Errorf("ID-DNE")
		return
	} else {

		event := Event{
			EventType: GetEvent,
			Msg:       c.msg,
			ID:        c.ID,
			Data:      rd.data,
		}

		c.eventch <- event
	}

}

//
//
//
func (reg *registry) unsub(c *cmd) {
	if reg.noMsg(c) {
		return
	}
}

//
//
//
func (reg *registry) unpub(c *cmd) {
	if reg.noMsg(c) {
		return
	}
}

func GetRandomIDString() (i IDString) {

	f, err := os.Open("/dev/urandom")
	if err != nil {
		fmt.Printf("Error readong /dev/urandom: %s\n", err)
		panic(err)
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	str := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	i = IDString(str)
	return i
}
