package msgbus

import (
	"fmt"
	"os"
	"time"
)

type operation string
type EventType string
type MsgType string
type ContractState string
type MinerState string
type ConnectionState string
type MsgBusError string

const (
	opNop      operation = "opNop"
	opPub      operation = "opPub"
	opSub      operation = "opSub"
	opGet      operation = "opGet"
	opSet      operation = "opSet"
	opUnsub    operation = "opUnsub"
	opUnpub    operation = "opUnpub"
	opRemove   operation = "opRemove"
	opShutdown operation = "opShutdown"
)

const (
	NoEvent           EventType = "NoEvent"
	UpdateEvent       EventType = "UpdEvent"
	DeleteEvent       EventType = "DelEvent"
	GetEvent          EventType = "GetEvent"
	GetIndexEvent     EventType = "GetIdxEvent"
	PublishEvent      EventType = "PubEvent"
	UnpublishEvent    EventType = "UnpubEvent"
	SubscribedEvent   EventType = "SubEvent"
	UnsubscribedEvent EventType = "UnsubEvent"
	RemovedEvent      EventType = "RemovedEvent"
)

const (
	NoMsg         MsgType = "NoMsg"
	ConfigMsg     MsgType = "ConfigMsg"
	DestMsg       MsgType = "DestMsg"
	SellerMsg     MsgType = "SellerMsg"
	ContractMsg   MsgType = "ContractMsg"
	MinerMsg      MsgType = "MinerMsg"
	ConnectionMsg MsgType = "ConnectionMsg"
	LogMsg        MsgType = "LogMsg"
)

const (
	ContNewState      ContractState = "NewState"
	ContReadyState    ContractState = "ReadyState"
	ContActiveState   ContractState = "ActiveState"
	ContCompleteState ContractState = "CompleteState"
)

const (
	OnlineState  MinerState = "MinerOnlineState"
	OfflineState MinerState = "MinerOfflineState"
)

const (
	ConnNewState         ConnectionState = "NewState"
	ConnSrcOpenState     ConnectionState = "SrcOpenState"
	ConnAuthState        ConnectionState = "AuthState"
	ConnVerifyState      ConnectionState = "VerifyState"
	ConnRoutingState     ConnectionState = "RoutingState"
	ConnConnectingState  ConnectionState = "ConnectingState"
	ConnConnectedState   ConnectionState = "ConnectedState"
	ConnConnectErrState  ConnectionState = "ConnectErrState"
	ConnMsgErrState      ConnectionState = "MsgErrState"
	ConnRouteChangeState ConnectionState = "RouteChangeState"
	ConnDstCloseState    ConnectionState = "DstCloseState"
	ConnSrcCloseState    ConnectionState = "SrcCloseState"
	ConnShutdownState    ConnectionState = "ShutdownState"
	ConnErrorState       ConnectionState = "ErrorState"
	ConnClosedState      ConnectionState = "ClosedState"
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
	Err       error
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
	StartDate        time.Time
	EndDate          time.Time
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

type IDIndex []IDString

type registry struct {
	data   map[MsgType]map[IDString]registryData
	notify map[MsgType]map[chan Event]interface{}
}

// PubSub is a collection of topics.
type PubSub struct {
	cmdChan  chan *cmd
	capacity int
}

const (
	MsgBusErrNoErr       MsgBusError = "NoErr"
	MsgBusErrNoMsg       MsgBusError = "NoMsg"
	MsgBusErrNoID        MsgBusError = "NoID"
	MsgBusErrNoSub       MsgBusError = "NoSub"
	MsgBusErrNoData      MsgBusError = "NoData"
	MsgBusErrNoEventChan MsgBusError = "NoEventChan"
	MsgBusErrBadMsg      MsgBusError = "BadMsg"
	MsgBusErrBadID       MsgBusError = "BadID"
	MsgBusErrBadData     MsgBusError = "BadData"
	MsgBusErrDupID       MsgBusError = "DupID"
	MsgBusErrDupData     MsgBusError = "DupData"
)

type cmd struct {
	op       operation
	sync     bool
	msg      MsgType
	ID       IDString
	data     interface{}
	eventch  EventChan
	returnch EventChan
}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func getCommandError(e MsgBusError) error {
	return fmt.Errorf("command Error: %s", e)
}

//--------------------------------------------------------------------------------
// New creates a new PubSub and starts a goroutine for handling operations.
// The capacity of the channels created by Sub and SubOnce will be as specified.
//
// Convert this to a message bus
//
//--------------------------------------------------------------------------------
func New(capacity int) *PubSub {
	ps := &PubSub{make(chan *cmd), capacity}
	go ps.start()
	return ps
}

//--------------------------------------------------------------------------------
//--------------------------------------------------------------------------------
func (ps *PubSub) NewEventChanPtr() *EventChan {
	ech := ps.NewEventChan()
	return &ech
}

//--------------------------------------------------------------------------------
//--------------------------------------------------------------------------------
func (ps *PubSub) NewEventChan() EventChan {
	return make(EventChan)
}

//--------------------------------------------------------------------------------
//--------------------------------------------------------------------------------
func (ps *PubSub) NewEventPtr() *Event {
	e := ps.NewEvent()
	return &e
}

//--------------------------------------------------------------------------------
//--------------------------------------------------------------------------------
func (ps *PubSub) NewEvent() Event {
	e := Event{EventType: NoEvent}
	return e
}

//--------------------------------------------------------------------------------
// Create new topic structure
//
//--------------------------------------------------------------------------------
func (ps *PubSub) Pub(msg MsgType, id IDString, data interface{}) (err error) {

	if msg == NoMsg {
		return getCommandError(MsgBusErrNoMsg)
	}

	if id == "" {
		return getCommandError(MsgBusErrNoID)
	}

	if data == nil {
		return getCommandError(MsgBusErrNoData)
	}

	c := cmd{
		op:      opPub,
		sync:    false,
		msg:     msg,
		ID:      id,
		data:    data,
		eventch: nil,
	}

	_, err = ps.dispatch(&c)

	return err
}

//--------------------------------------------------------------------------------
// Create new topic structure
//
//--------------------------------------------------------------------------------
func (ps *PubSub) PubWait(msg MsgType, id IDString, data interface{}) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	if id == "" {
		return e, getCommandError(MsgBusErrNoID)
	}

	if data == nil {
		return e, getCommandError(MsgBusErrNoData)
	}

	c := cmd{
		op:      opPub,
		sync:    true,
		msg:     msg,
		ID:      id,
		data:    data,
		eventch: nil,
	}

	e, err = ps.dispatch(&c)

	return e, err
}

//--------------------------------------------------------------------------------
// Request update events for the topic
//--------------------------------------------------------------------------------
func (ps *PubSub) Sub(msg MsgType, id IDString, ech EventChan) (err error) {

	if msg == NoMsg {
		return getCommandError(MsgBusErrNoMsg)
	}

	if ech == nil {
		return getCommandError(MsgBusErrNoEventChan)
	}

	c := cmd{
		op:      opSub,
		sync:    false,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: ech,
	}

	_, err = ps.dispatch(&c)

	return err
}

//--------------------------------------------------------------------------------
// Request update events for the topic
//--------------------------------------------------------------------------------
func (ps *PubSub) SubWait(msg MsgType, id IDString, ech EventChan) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	if ech == nil {
		return e, getCommandError(MsgBusErrNoEventChan)
	}

	c := cmd{
		op:      opSub,
		sync:    true,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: ech,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) Get(msg MsgType, id IDString, ech EventChan) (err error) {

	if msg == NoMsg {
		return getCommandError(MsgBusErrNoMsg)
	}

	if ech == nil {
		return getCommandError(MsgBusErrNoEventChan)
	}

	c := cmd{
		op:      opGet,
		sync:    false,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: ech,
	}

	_, err = ps.dispatch(&c)

	return err

}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) GetWait(msg MsgType, id IDString, ech EventChan) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	if ech == nil {
		return e, getCommandError(MsgBusErrNoEventChan)
	}

	c := cmd{
		op:      opGet,
		sync:    true,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: ech,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) Set(msg MsgType, id IDString, data interface{}) (err error) {

	if msg == NoMsg {
		return getCommandError(MsgBusErrNoMsg)
	}

	if id == "" {
		return getCommandError(MsgBusErrNoID)
	}

	if data == nil {
		return getCommandError(MsgBusErrNoData)
	}

	c := cmd{
		op:      opSet,
		sync:    false,
		msg:     msg,
		ID:      id,
		data:    data,
		eventch: nil,
	}

	_, err = ps.dispatch(&c)

	return err
}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) SetWait(msg MsgType, id IDString, data interface{}) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	if id == "" {
		return e, getCommandError(MsgBusErrNoID)
	}

	if data == nil {
		return e, getCommandError(MsgBusErrNoData)
	}

	c := cmd{
		op:      opSet,
		sync:    true,
		msg:     msg,
		ID:      id,
		data:    data,
		eventch: nil,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
// Request removal of events for the topic
//--------------------------------------------------------------------------------
func (ps *PubSub) Unsub(msg MsgType, id IDString, ech EventChan) (err error) {

	if msg == NoMsg {
		return getCommandError(MsgBusErrNoMsg)
	}

	if id == "" {
		return getCommandError(MsgBusErrNoID)
	}

	if ech == nil {
		return getCommandError(MsgBusErrNoEventChan)
	}

	c := cmd{
		op:      opUnsub,
		sync:    false,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: ech,
	}

	_, err = ps.dispatch(&c)

	return err
}

//--------------------------------------------------------------------------------
// Request removal of events for the topic
//--------------------------------------------------------------------------------
func (ps *PubSub) UnsubWait(msg MsgType, id IDString, ech EventChan) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	if id == "" {
		return e, getCommandError(MsgBusErrNoID)
	}

	if ech == nil {
		return e, getCommandError(MsgBusErrNoEventChan)
	}

	c := cmd{
		op:      opUnsub,
		sync:    true,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: ech,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
// Request removal of the topic
//--------------------------------------------------------------------------------
func (ps *PubSub) Unpub(msg MsgType, id IDString) (err error) {

	if msg == NoMsg {
		return getCommandError(MsgBusErrNoMsg)
	}

	if id == "" {
		return getCommandError(MsgBusErrNoID)
	}

	c := cmd{
		op:      opUnpub,
		sync:    false,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: nil,
	}

	_, err = ps.dispatch(&c)

	return err
}

//--------------------------------------------------------------------------------
// Request removal of the topic
//--------------------------------------------------------------------------------
func (ps *PubSub) UnpubWait(msg MsgType, id IDString) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	if id == "" {
		return e, getCommandError(MsgBusErrNoID)
	}

	c := cmd{
		op:      opUnpub,
		sync:    true,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: nil,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
// Request update events for the topic
//--------------------------------------------------------------------------------
func (ps *PubSub) RemoveAndCloseEventChan(ech EventChan) (err error) {

	if ech == nil {
		return getCommandError(MsgBusErrNoEventChan)
	}

	c := cmd{
		op:      opRemove,
		sync:    false,
		msg:     NoMsg,
		ID:      "",
		data:    nil,
		eventch: ech,
	}

	_, err = ps.dispatch(&c)

	return err
}

//--------------------------------------------------------------------------------
// Request update events for the topic
//--------------------------------------------------------------------------------
func (ps *PubSub) RemoveAndCloseEventChanWait(ech EventChan) (e Event, err error) {

	if ech == nil {
		return e, getCommandError(MsgBusErrNoEventChan)
	}

	c := cmd{
		op:      opRemove,
		sync:    true,
		msg:     NoMsg,
		ID:      "",
		data:    nil,
		eventch: ech,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
//
// Shutdown closes all subscribed channels and terminates the goroutine.
//
//--------------------------------------------------------------------------------
func (ps *PubSub) Shutdown() (err error) {

	c := cmd{
		op:      opShutdown,
		sync:    false,
		msg:     NoMsg,
		ID:      "",
		data:    nil,
		eventch: nil,
	}

	_, err = ps.dispatch(&c)

	return err
}

//--------------------------------------------------------------------------------
//
// Shutdown closes all subscribed channels and terminates the goroutine.
//
//--------------------------------------------------------------------------------
func (ps *PubSub) ShutdownWait() (e Event, err error) {

	c := cmd{
		op:      opShutdown,
		sync:    true,
		msg:     NoMsg,
		ID:      "",
		data:    nil,
		eventch: nil,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
//
//
//
//--------------------------------------------------------------------------------
func (ps *PubSub) dispatch(c *cmd) (event Event, e error) {

	e = nil

	if c.op == opNop {
		return event, getCommandError(MsgBusErrNoErr)
	}

	if c.sync {
		c.returnch = make(EventChan)
	}

	ps.cmdChan <- c

	if c.sync {
		event = <-c.returnch
		e = event.Err
		close(c.returnch)
	}

	return event, e
}

//--------------------------------------------------------------------------------
//
//
//
//--------------------------------------------------------------------------------
func (ps *PubSub) start() {
	reg := registry{
		// data:   make(map[MsgType]map[IDString]interface{}),
		data:   make(map[MsgType]map[IDString]registryData),
		notify: make(map[MsgType]map[chan Event]interface{}),
	}

loop:
	for cmdptr := range ps.cmdChan {

		fmt.Printf("New Command %+v\n", *cmdptr)

		if cmdptr.op == opNop {
			// cmdptr.err = nil
			continue loop
		}

		if cmdptr.op == opShutdown {
			break loop
		}

		switch cmdptr.op {
		case opPub:
			reg.pub(cmdptr)

		case opSub:
			reg.sub(cmdptr)

		case opSet:
			reg.set(cmdptr)

		case opGet:
			reg.get(cmdptr)

		case opUnpub:
			reg.unpub(cmdptr)

		case opUnsub:
			reg.unsub(cmdptr)

		case opRemove:
			reg.removeAndClose(cmdptr)

		default:
			panic("default reached for cmd.op")
		}

	}

	fmt.Printf("Closing PubSub Command chan\n")

	// clean up here
	// Close any open channels
	// Delete registry

}

//-----------------------------------------
//-----------------------------------------
func sendEvent(e EventChan, event Event) {
	go func(e EventChan, event Event) {
		e <- event
	}(e, event)
}

//-----------------------------------------
// msg contains the message type
// ID contains the new value
// data contains the new data struct
//
//-----------------------------------------
func (reg *registry) pub(c *cmd) {

	event := Event{
		EventType: PublishEvent,
		Msg:       c.msg,
		ID:        c.ID,
		Data:      c.data,
		Err:       nil,
	}

	if _, ok := reg.data[c.msg]; !ok {
		reg.data[c.msg] = make(map[IDString]registryData)
	}

	if _, ok := reg.data[c.msg][c.ID]; ok {
		event.Err = getCommandError(MsgBusErrDupData)
	} else {
		reg.data[c.msg][c.ID] = registryData{
			sub:  Subscribers{eventchan: make(map[EventChan]int)},
			data: c.data,
		}
	}

	// If sync, return the event
	if c.sync {
		sendEvent(c.returnch, event)
	}

	if c.eventch != nil {
		sendEvent(c.eventch, event)
	}

	// If no error, copy the event to everyone interested
	if event.Err != nil {
		for ech, _ := range reg.notify[c.msg] {
			sendEvent(ech, event)
		}
	}

}

//-----------------------------------------
// msg
// ID (optional)
// ch Event Channel
//
//-----------------------------------------
func (reg *registry) sub(c *cmd) {

	event := Event{
		EventType: SubscribedEvent,
		Msg:       c.msg,
		ID:        c.ID,
		Data:      c.data,
		Err:       nil,
	}

	if c.ID == "" {
		if _, ok := reg.notify[c.msg]; !ok {
			event.Err = getCommandError(MsgBusErrBadMsg)
		} else if _, ok := reg.notify[c.msg][c.eventch]; ok {
			event.Err = getCommandError(MsgBusErrDupData)
		} else {
			reg.notify[c.msg][c.eventch] = 1
		}
	} else {
		if _, ok := reg.data[c.msg]; !ok {
			event.Err = getCommandError(MsgBusErrBadMsg)
		} else if _, ok := reg.data[c.msg][c.ID]; !ok {
			event.Err = getCommandError(MsgBusErrBadID)
		} else if _, ok := reg.data[c.msg][c.ID].sub.eventchan[c.eventch]; ok {
			event.Err = getCommandError(MsgBusErrDupData)
		} else {
			reg.data[c.msg][c.ID].sub.eventchan[c.eventch] = 1
		}
	}

	if c.sync {
		sendEvent(c.returnch, event)
	}

	if c.eventch != nil {
		sendEvent(c.eventch, event)
	}
}

//-----------------------------------------
//
//
//
//-----------------------------------------
func (reg *registry) set(c *cmd) {

	event := Event{
		EventType: UpdateEvent,
		Msg:       c.msg,
		ID:        c.ID,
		Data:      c.data,
		Err:       nil,
	}

	if c.ID == "" {
		event.Err = getCommandError(MsgBusErrNoID)
	} else if _, ok := reg.data[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrBadMsg)
	} else if _, ok := reg.notify[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrBadMsg)
	} else if _, ok := reg.data[c.msg][c.ID]; !ok {
		event.Err = getCommandError(MsgBusErrBadID)
	} else {

		//
		// Need Error Checking on the in coming data type to match msgtype
		//

		// Could do a lot here to check if the data actually changed
		// Set the data

		d := reg.data[c.msg][c.ID]
		d.data = c.data
		reg.data[c.msg][c.ID] = d

	}

	if c.sync {
		sendEvent(c.returnch, event)
	}

	if c.eventch != nil {
		sendEvent(c.eventch, event)
	}

	// Notify anyone listening
	for eventch, _ := range reg.data[c.msg][c.ID].sub.eventchan {
		if _, ok := reg.data[c.msg][c.ID].sub.eventchan[eventch]; ok {
			sendEvent(eventch, event)
		} else {
			panic("eventchannel was not ok in set()")
		}
	}
}

//-----------------------------------------
//
// Msg
// ID
// eventch - where to send the get request too
//
//-----------------------------------------
func (reg *registry) get(c *cmd) {

	event := Event{
		EventType: GetEvent,
		Msg:       c.msg,
		ID:        c.ID,
		Data:      nil,
		Err:       nil,
	}

	if _, ok := reg.data[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrBadMsg)
	} else if c.ID == "" {
		var index IDIndex
		for i, _ := range reg.data[c.msg] {
			index = append(index, i)
		}
		event.EventType = GetIndexEvent
		event.Data = index
	} else if _, ok := reg.data[c.msg][c.ID]; !ok {
		event.Err = getCommandError(MsgBusErrBadID)
	} else {
		event.Data = reg.data[c.msg][c.ID].data
	}

	if c.sync {
		sendEvent(c.returnch, event)
	}
	if c.eventch != nil {
		sendEvent(c.eventch, event)
	}

}

//-----------------------------------------
//
//
//
//-----------------------------------------
func (reg *registry) unsub(c *cmd) {

	event := Event{
		EventType: UnsubscribedEvent,
		Msg:       c.msg,
		ID:        c.ID,
		Data:      nil,
		Err:       nil,
	}

	if _, ok := reg.data[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrBadMsg)
	} else if _, ok := reg.notify[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrBadMsg)
	} else if c.ID == "" {
		if _, ok := reg.notify[c.msg][c.eventch]; !ok {
			event.Err = getCommandError(MsgBusErrNoSub)
		} else {
			delete(reg.notify[c.msg], c.eventch)
		}
	} else {
		if _, ok := reg.data[c.msg][c.ID]; !ok {
			event.Err = getCommandError(MsgBusErrBadID)
		} else if _, ok := reg.data[c.msg][c.ID].sub.eventchan[c.eventch]; !ok {
			event.Err = getCommandError(MsgBusErrNoSub)
		} else {
			delete(reg.data[c.msg][c.ID].sub.eventchan, c.eventch)
		}
	}

	if c.sync {
		sendEvent(c.returnch, event)
	}
	if c.eventch != nil {
		sendEvent(c.eventch, event)
	}

}

//-----------------------------------------
//
//-----------------------------------------
func (reg *registry) unpub(c *cmd) {

	event := Event{
		EventType: UnpublishEvent,
		Msg:       c.msg,
		ID:        c.ID,
		Data:      nil,
		Err:       nil,
	}

	if _, ok := reg.data[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrBadMsg)
	} else if _, ok := reg.notify[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrBadMsg)
	} else if c.ID == "" {
		event.Err = getCommandError(MsgBusErrNoID)
	} else if _, ok := reg.data[c.msg][c.ID]; !ok {
		event.Err = getCommandError(MsgBusErrBadID)
	}

	if c.sync {
		sendEvent(c.returnch, event)
	}

	if c.eventch != nil {
		sendEvent(c.eventch, event)
	}

	for ech, _ := range reg.data[c.msg][c.ID].sub.eventchan {
		sendEvent(ech, event)
	}

	for ech, _ := range reg.notify[c.msg] {
		sendEvent(ech, event)
	}

	delete(reg.data[c.msg], c.ID)

}

//---------------------------------------
//
//	data   map[MsgType]map[IDString]registryData
//	notify map[MsgType]map[chan Event]interface{}
//
//---------------------------------------
func (reg *registry) removeAndClose(c *cmd) {

	event := Event{
		EventType: RemovedEvent,
		Msg:       c.msg,
		ID:        c.ID,
		Data:      nil,
		Err:       nil,
	}

	if c.eventch == nil {
		event.Err = getCommandError(MsgBusErrNoEventChan)
	}

	for msg, _ := range reg.notify {
		if _, ok := reg.notify[msg][c.eventch]; ok {
			delete(reg.notify[msg], c.eventch)
		}
	}

	for msg, _ := range reg.data {
		for id, _ := range reg.data[msg] {
			if _, ok := reg.data[msg][id].sub.eventchan[c.eventch]; ok {
				delete(reg.data[c.msg][c.ID].sub.eventchan, c.eventch)
			}
		}
	}

	close(c.eventch)

	if c.sync {
		sendEvent(c.returnch, event)
	}

}

//-----------------------------------------
//
//-----------------------------------------
func GetRandomIDString() (i IDString) {

	f, err := os.Open("/dev/urandom")
	if err != nil {
		fmt.Printf("Error readong /dev/urandom: %s\n", err)
		panic(err)
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	//fmt.Printf("%08x-%08x-%08x-%08x\n", b[0:4], b[4:8], b[8:12], b[12:16])
	str := fmt.Sprintf("%08x-%08x-%08x-%08x", b[0:4], b[4:8], b[8:12], b[12:16])
	i = IDString(str)
	return i
}
