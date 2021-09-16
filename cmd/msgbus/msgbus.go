package msgbus

import (
	"fmt"
	"os"
	"time"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
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
	opSearch   operation = "opSearch"
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
	SearchEvent       EventType = "SearchEvent"
	SearchIndexEvent  EventType = "SearchIndexEvent"
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
	ConnNewState          ConnectionState = "NewState"
	ConnSrcSubscribeState ConnectionState = "SrcSubscribeState"
	ConnAuthState         ConnectionState = "AuthState"
	ConnVerifyState       ConnectionState = "VerifyState"
	ConnRoutingState      ConnectionState = "RoutingState"
	ConnConnectingState   ConnectionState = "ConnectingState"
	ConnConnectedState    ConnectionState = "ConnectedState"
	ConnConnectErrState   ConnectionState = "ConnectErrState"
	ConnMsgErrState       ConnectionState = "MsgErrState"
	ConnRouteChangeState  ConnectionState = "RouteChangeState"
	ConnDstCloseState     ConnectionState = "DstCloseState"
	ConnSrcCloseState     ConnectionState = "SrcCloseState"
	ConnShutdownState     ConnectionState = "ShutdownState"
	ConnErrorState        ConnectionState = "ErrorState"
	ConnClosedState       ConnectionState = "ClosedState"
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

const DEFAULT_DEST_ID DestID = "DefaultDestID"

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

type DestNetProto string
type DestNetHost string
type DestNetPort string
type Dest struct {
	ID       DestID
	NetHost  DestNetHost
	NetPort  DestNetPort
	NetProto DestNetProto
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

//
// Created & Updated by Connection Manager
//
type Miner struct {
	ID                      MinerID
	Name                    string
	IP                      string
	MAC                     string
	State                   MinerState
	Seller                  SellerID
	Dest                    DestID // Updated by Connection Scheduler
	InitialMeasuredHashRate int
	CurrentHashRate         int
}

//
// Created & Updated by Connection Manager
//
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
	MsgBusErrNoErr         MsgBusError = "NoErr"
	MsgBusErrNoMsg         MsgBusError = "NoMsg"
	MsgBusErrNoID          MsgBusError = "NoID"
	MsgBusErrNoSub         MsgBusError = "NoSub"
	MsgBusErrNoData        MsgBusError = "NoData"
	MsgBusErrNoSearchTerm  MsgBusError = "NoSearchTerm"
	MsgBusErrNoEventChan   MsgBusError = "NoEventChan"
	MsgBusErrBadMsg        MsgBusError = "BadMsg"
	MsgBusErrBadID         MsgBusError = "BadID"
	MsgBusErrBadData       MsgBusError = "BadData"
	MsgBusErrBadSearchTerm MsgBusError = "BadSearchTerm"
	MsgBusErrDupID         MsgBusError = "DupID"
	MsgBusErrDupData       MsgBusError = "DupData"
)

type cmd struct {
	op       operation
	sync     bool
	msg      MsgType
	ID       IDString
	Name     string
	IP       string
	MAC      string
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
func (ps *PubSub) GetWait(msg MsgType, id IDString) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	c := cmd{
		op:      opGet,
		sync:    true,
		msg:     msg,
		ID:      id,
		data:    nil,
		eventch: nil,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) SearchIP(msg MsgType, ip string, ech EventChan) (err error) {

	if msg == NoMsg {
		return getCommandError(MsgBusErrNoMsg)
	}

	if ech == nil {
		return getCommandError(MsgBusErrNoEventChan)
	}

	if ip == "" {
		return getCommandError(MsgBusErrNoSearchTerm)
	}

	c := cmd{
		op:      opSearch,
		sync:    false,
		msg:     msg,
		IP:      ip,
		data:    nil,
		eventch: ech,
	}

	_, err = ps.dispatch(&c)

	return err

}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) SearchIPWait(msg MsgType, ip string) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	if ip == "" {
		return e, getCommandError(MsgBusErrNoSearchTerm)
	}

	c := cmd{
		op:      opSearch,
		sync:    true,
		msg:     msg,
		IP:      ip,
		data:    nil,
		eventch: nil,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) SearchMAC(msg MsgType, mac string, ech EventChan) (err error) {

	if msg == NoMsg {
		return getCommandError(MsgBusErrNoMsg)
	}

	if ech == nil {
		return getCommandError(MsgBusErrNoEventChan)
	}

	if mac == "" {
		return getCommandError(MsgBusErrNoSearchTerm)
	}

	c := cmd{
		op:      opSearch,
		sync:    false,
		msg:     msg,
		MAC:     mac,
		data:    nil,
		eventch: ech,
	}

	_, err = ps.dispatch(&c)

	return err

}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) SearchMACWait(msg MsgType, mac string) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	if mac == "" {
		return e, getCommandError(MsgBusErrNoSearchTerm)
	}

	c := cmd{
		op:      opSearch,
		sync:    true,
		msg:     msg,
		MAC:     mac,
		data:    nil,
		eventch: nil,
	}

	return ps.dispatch(&c)

}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) SearchName(msg MsgType, name string, ech EventChan) (err error) {

	if msg == NoMsg {
		return getCommandError(MsgBusErrNoMsg)
	}

	if ech == nil {
		return getCommandError(MsgBusErrNoEventChan)
	}

	if name == "" {
		return getCommandError(MsgBusErrNoSearchTerm)
	}

	c := cmd{
		op:      opSearch,
		sync:    false,
		msg:     msg,
		Name:    name,
		data:    nil,
		eventch: ech,
	}

	_, err = ps.dispatch(&c)

	return err

}

//--------------------------------------------------------------------------------
//
//--------------------------------------------------------------------------------
func (ps *PubSub) SearchNameWait(msg MsgType, name string) (e Event, err error) {

	if msg == NoMsg {
		return e, getCommandError(MsgBusErrNoMsg)
	}

	if name == "" {
		return e, getCommandError(MsgBusErrNoSearchTerm)
	}

	c := cmd{
		op:      opSearch,
		sync:    true,
		msg:     msg,
		Name:    name,
		data:    nil,
		eventch: nil,
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
		data:   make(map[MsgType]map[IDString]registryData),
		notify: make(map[MsgType]map[chan Event]interface{}),
	}

	reg.data[ConfigMsg] = make(map[IDString]registryData)
	reg.data[DestMsg] = make(map[IDString]registryData)
	reg.data[SellerMsg] = make(map[IDString]registryData)
	reg.data[ContractMsg] = make(map[IDString]registryData)
	reg.data[MinerMsg] = make(map[IDString]registryData)
	reg.data[ConnectionMsg] = make(map[IDString]registryData)

	reg.notify[ConfigMsg] = make(map[chan Event]interface{})
	reg.notify[DestMsg] = make(map[chan Event]interface{})
	reg.notify[SellerMsg] = make(map[chan Event]interface{})
	reg.notify[ContractMsg] = make(map[chan Event]interface{})
	reg.notify[MinerMsg] = make(map[chan Event]interface{})
	reg.notify[ConnectionMsg] = make(map[chan Event]interface{})

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

		case opSearch:
			reg.search(cmdptr)

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
//
//-----------------------------------------
func (event *Event) send(e EventChan) {

	go func(e EventChan) {
		e <- *event
	}(e)

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
		// sendEvent(c.returnch, event)
		event.send(c.returnch)
	}

	if c.eventch != nil {
		// sendEvent(c.eventch, event)
		event.send(c.eventch)
	}

	// If no error, copy the event to everyone interested
	if event.Err != nil {
		for ech, _ := range reg.notify[c.msg] {
			//sendEvent(ech, event)
			event.send(ech)
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
		event.send(c.returnch)
	}

	if c.eventch != nil {
		event.send(c.eventch)
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
		fmt.Printf(lumerinlib.FileLine()+"Error:%s\n", event.Err)
	} else if _, ok := reg.data[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrBadMsg)
		fmt.Printf(lumerinlib.FileLine()+"Error:%s\n", event.Err)
	} else if _, ok := reg.notify[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrBadMsg)
		fmt.Printf(lumerinlib.FileLine()+"Error:%s\n", event.Err)
	} else if _, ok := reg.data[c.msg][c.ID]; !ok {
		event.Err = getCommandError(MsgBusErrBadID)
		fmt.Printf(lumerinlib.FileLine()+"Error:%s\n", event.Err)
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
		event.send(c.returnch)
	}

	if c.eventch != nil {
		event.send(c.eventch)
	}

	// Notify anyone listening
	for ech, _ := range reg.data[c.msg][c.ID].sub.eventchan {
		if _, ok := reg.data[c.msg][c.ID].sub.eventchan[ech]; ok {
			event.send(ech)
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
		event.send(c.returnch)
	}
	if c.eventch != nil {
		event.send(c.eventch)
	}

}

//-----------------------------------------
//
//
//-----------------------------------------
func (reg *registry) search(c *cmd) {

	event := Event{
		EventType: SearchEvent,
		Msg:       c.msg,
		Data:      nil,
		Err:       nil,
	}

	// Only works for Miner messages at the moment.
	if c.msg != MinerMsg {
		event.Err = getCommandError(MsgBusErrBadMsg)
		return
	}

	if c.Name == "" && c.IP == "" && c.MAC != "" {
		event.Err = getCommandError(MsgBusErrBadSearchTerm)
		return
	}

	if _, ok := reg.data[c.msg]; !ok {
		event.Err = getCommandError(MsgBusErrNoData)
		return
	}

	var index IDIndex
	switch {
	case c.Name != "":
		for i, _ := range reg.data[c.msg] {
			if reg.data[c.msg][i].data.(Miner).Name == c.Name {
				index = append(index, i)
			}
		}
	case c.IP != "":
		for i, _ := range reg.data[c.msg] {
			if reg.data[c.msg][i].data.(Miner).IP == c.IP {
				index = append(index, i)
			}
		}
	case c.MAC != "":
		for i, _ := range reg.data[c.msg] {
			if reg.data[c.msg][i].data.(Miner).MAC == c.MAC {
				index = append(index, i)
			}
		}
	default:
		panic("")
	}
	event.EventType = SearchIndexEvent
	event.Data = index

	if c.sync {
		event.send(c.returnch)
	}
	if c.eventch != nil {
		event.send(c.eventch)
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
		event.send(c.returnch)
	}
	if c.eventch != nil {
		event.send(c.eventch)
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
		event.send(c.returnch)
	}

	if c.eventch != nil {
		event.send(c.eventch)
	}

	for ech, _ := range reg.data[c.msg][c.ID].sub.eventchan {
		event.send(ech)
	}

	for ech, _ := range reg.notify[c.msg] {
		event.send(ech)
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
		event.send(c.returnch)
	}

}

//-----------------------------------------
//
//-----------------------------------------
func GetRandomIDString() (i IDString) {

	f, err := os.Open("/dev/urandom")
	if err != nil {
		fmt.Printf("Error reading /dev/urandom: %s\n", err)
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
