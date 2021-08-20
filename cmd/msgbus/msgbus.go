package msgbus

import (
	"fmt"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
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
	UpdateEvent EventType = iota
	DeleteEvent
	GetEvent
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
	SrcOpenState ConnectionState = iota
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

type IDString string
type ConfigID IDString
type DestID IDString
type SellerID IDString
type BuyerID IDString
type ContractID IDString
type MinerID IDString
type ConnectionID IDString

type Event struct {
	eventType EventType
	msg       MsgType
	ID        IDString
	data      interface{}
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
	DefaultDest Dest
	Seller      SellerID
}

type Seller struct {
	ID                     SellerID
	DefaultDest            Dest
	TotalAvailableHashRate int
	UnusedHashRate         int
	NewContracts           []ContractID
	ReadyContracts         []ContractID
	ActiveContracts        []ContractID
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
	StartDate int
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

func BoilerPlateFunc() (string, error) {
	msg := "Logging Package"
	return lumerinlib.BoilerPlateLibFunc(msg), nil // always returns no error
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
				eventType: UpdateEvent,
				msg:       c.msg,
				ID:        c.ID,
				data:      c.data,
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
			eventType: GetEvent,
			msg:       c.msg,
			ID:        "",
			data:      ids,
		}

		c.eventch <- event
		return
	}

	if rd, ok := reg.data[c.msg][c.ID]; !ok {
		c.err = fmt.Errorf("ID-DNE")
		return
	} else {

		event := Event{
			eventType: GetEvent,
			msg:       c.msg,
			ID:        c.ID,
			data:      rd.data,
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
