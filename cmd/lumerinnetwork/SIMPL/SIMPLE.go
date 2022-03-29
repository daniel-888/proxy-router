package simple

import (
	"context"
	"errors"
	"fmt"
	"net"

	"gitlab.com/TitanInd/lumerin/cmd/log"
	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/connectionmanager"
	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/lumerinconnection"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

type ConnUniqueID int
type URL string
type ID string
type Data string
type EventHandler string
type SearchString string

type IDString msgbus.IDString
type MsgType msgbus.MsgType

const (
	NoMsg                    MsgType = MsgType(msgbus.NoMsg)
	ConfigMsg                MsgType = MsgType(msgbus.ConfigMsg)
	ContractManagerConfigMsg MsgType = MsgType(msgbus.ContractManagerConfigMsg)
	DestMsg                  MsgType = MsgType(msgbus.DestMsg)
	NodeOperatorMsg          MsgType = MsgType(msgbus.NodeOperatorMsg)
	ContractMsg              MsgType = MsgType(msgbus.ContractMsg)
	MinerMsg                 MsgType = MsgType(msgbus.MinerMsg)
	ConnectionMsg            MsgType = MsgType(msgbus.ConnectionMsg)
	LogMsg                   MsgType = MsgType(msgbus.LogMsg)
)

type SimpleListenStruct struct {
	ctx              context.Context
	cancel           func()
	accept           chan *SimpleStruct //channel to accept simple structs and process their message
	connectionListen *connectionmanager.ConnectionListenStruct
	msgbus           *msgbus.PubSub
	logger           *log.Logger
}

/*
The simple struct is used to point to a specific instance
of a connection manager and MsgBus. The structure ties these
to a protocol struct where events are directed to be handled.
*/
type SimpleStruct struct {
	ctx               context.Context
	cancel            func()             //it might make sense to use the WithCancel function instead
	eventChan         chan *SimpleEvent  // Channle to get
	msgbusChan        chan *msgbus.Event //
	openChan          chan *SimpleConnOpenEvent
	connectionMapping map[ConnUniqueID]*lumerinconnection.LumerinSocketStruct //mapping of uint to connections
	ConnectionStruct  *connectionmanager.ConnectionStruct
	msgbus            *msgbus.PubSub
	logger            *log.Logger
}

/*
a struct that contains the data and the event type being passed into the SimpleStruct
*/
type SimpleEvent struct {
	EventType     EventType
	ConnReadEvent *SimpleConnReadEvent
	ConnOpenEvent *SimpleConnOpenEvent
	MsgBusEvent   *SimpleMsgBusEvent
}

type SimpleConnReadEvent struct {
	uID   ConnUniqueID
	data  []byte
	count int
	err   error
}

type SimpleConnOpenEvent struct {
	uID  ConnUniqueID
	dest *msgbus.Dest
	err  error
}

func (s *SimpleConnReadEvent) UniqueID() ConnUniqueID { return s.uID }
func (s *SimpleConnReadEvent) Data() []byte           { return s.data }
func (s *SimpleConnReadEvent) Count() int             { return s.count }
func (s *SimpleConnReadEvent) Err() error             { return s.err }

func (s *SimpleConnOpenEvent) UniqueID() ConnUniqueID { return s.uID }
func (s *SimpleConnOpenEvent) Dest() *msgbus.Dest     { return s.dest }
func (s *SimpleConnOpenEvent) Err() error             { return s.err }

type SimpleMsgBusEvent struct {
	EventType EventType
	Msg       MsgType
	ID        IDString
	RequestID int
	Data      interface{}
	Err       error
}

/*
struct that tells the SimpleStruct which connection to provide
the encoded data to
*/

type EventType string

const NoEvent EventType = "noevent"
const MsgBusEvent EventType = "msgbus"

const MsgUpdateEvent EventType = EventType(msgbus.UpdateEvent)
const MsgDeleteEvent EventType = EventType(msgbus.DeleteEvent)
const MsgGetEvent EventType = EventType(msgbus.GetEvent)
const MsgGetIndexEvent EventType = EventType(msgbus.GetIndexEvent)
const MsgSearchEvent EventType = EventType(msgbus.SearchEvent)
const MsgSearchIndexEvent EventType = EventType(msgbus.SearchIndexEvent)
const MsgPublishEvent EventType = EventType(msgbus.PublishEvent)
const MsgUnpublishEvent EventType = EventType(msgbus.UnpublishEvent)
const MsgSubscribedEvent EventType = EventType(msgbus.SubscribedEvent)
const MsgUnsubscribedEvent EventType = EventType(msgbus.UnsubscribedEvent)
const MsgRemovedEvent EventType = EventType(msgbus.RemovedEvent)
const ConnOpenEvent EventType = "connopen"
const ConnReadEvent EventType = "connread"
const ConnEOFEvent EventType = "conneof"
const ConnErrorEvent EventType = "connerror"
const ErrorEvent EventType = "error"
const MsgToProtocol EventType = "msgUp"

// ----------------------------------------------------------------------
//  SimpleListenStruct Functions
// ----------------------------------------------------------------------

func NewListen(ctx context.Context, listen net.Addr) (SimpleListenStruct, error) {

	ctx, cancel := context.WithCancel(ctx)

	c := ctx.Value(contextlib.ContextKey)
	if c == nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" called")
	}

	cs, ok := c.(*contextlib.ContextStruct)
	if !ok {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLine()+" Context Structre not correct")
	}

	if cs.GetSrc() == nil {
		cs.Logf(contextlib.LevelPanic, "Context Src Addr not defined")
	}

	if cs.GetMsgBus() == nil {
		cs.Logf(contextlib.LevelPanic, "Context MsgBus not defined")
	}

	if cs.GetLog() == nil {
		cs.Logf(contextlib.LevelPanic, "Context Logger not defined")
	}

	cls, e := connectionmanager.NewListen(ctx)
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Lumerin Listen() returne error:%s", e)
	}

	myStruct := SimpleListenStruct{
		ctx:              ctx,
		cancel:           cancel,
		accept:           make(chan *SimpleStruct),
		connectionListen: cls,
		msgbus:           cs.GetMsgBus(),
		logger:           cs.GetLog(),
	}
	// determine if a more robust error message is needed
	return myStruct, nil
}

//consider calling this as a gorouting from protocol layer, assuming
//protocll layer will have a layer to communicate with a chan over
func (s *SimpleListenStruct) Run() {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	s.connectionListen.Run()
	go s.goListenAccept()

}

//
//
//
func (s *SimpleListenStruct) goListenAccept() {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	cs := contextlib.GetContextStruct(s.ctx)
	if cs == nil {
		contextlib.Logf(s.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Structre not correct")
	}

	// if cs.GetProtocol() == nil {
	// 	contextlib.Logf(s.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context New Protocol Function not defined")
	// }

	// This needs error checking....
	// proto := cs.GetProtocol()
	// if proto == nil {
	// 	contextlib.Logf(s.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetProtocol() returned nil")
	// }

	connectionStructChan := s.connectionListen.Accept()

FORLOOP:
	for {
		select {
		case <-s.ctx.Done():
			contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" context canceled")
			break FORLOOP
		case connectionStruct := <-connectionStructChan:

			if connectionStruct == nil {
				contextlib.Logf(s.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Connection Listen Accept returned nil")
				s.Cancel()
				break
			}

			//create a cancel function from the context in the SimpleListenStruct
			newctx, cancel := context.WithCancel(s.ctx)
			eventchan := make(chan *SimpleEvent)

			//creating a new simple struct to pass to the protocol layer
			newSimpleStruct := &SimpleStruct{
				ctx:               newctx,
				cancel:            cancel,
				eventChan:         eventchan,
				msgbusChan:        make(chan *msgbus.Event),
				openChan:          make(chan *SimpleConnOpenEvent),
				connectionMapping: map[ConnUniqueID]*lumerinconnection.LumerinSocketStruct{},
				ConnectionStruct:  connectionStruct,
				msgbus:            s.msgbus,
				logger:            s.logger,
			}

			s.accept <- newSimpleStruct
		}
	}

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" exit")
}

//
//
//
func (s *SimpleListenStruct) GetAccept() <-chan *SimpleStruct {
	return s.accept
}

//
//
//
func (s *SimpleListenStruct) Done() bool {
	select {
	case <-s.ctx.Done():
		return true
	default:
		return false
	}
}

//
// Calls the listen context cancel function, which closes out the listener routine
//
func (s *SimpleListenStruct) Close() {
	if s.Done() {
		return
	}

	// Close any open structurs here?

	s.Cancel()
}

//
//
//
func (s *SimpleListenStruct) Cancel() {

	if s.cancel == nil {
		contextlib.Logf(s.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" cancel function is nul, struct:%v", s)
		return
	}

	if s.Done() {
		return
	}

	close(s.accept)
	s.cancel()
}

// ----------------------------------------------------------------------
//  SimpleStruct Functions
// ----------------------------------------------------------------------

//
//
//
func (s *SimpleStruct) GetEventChan() <-chan *SimpleEvent {

	if s.eventChan == nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" EventChan is nil SimpleStruct:%v", s)
	}

	return s.eventChan
}

//
//
//
func (s *SimpleStruct) Run() {
	contextlib.Logf(s.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" enter")

	if s == nil {
		panic(lumerinlib.FileLineFunc() + " SimpleStruct is nil")
	}
	if s.ConnectionStruct == nil {
		panic(lumerinlib.FileLineFunc() + " SimpleStruct.ConnectionStruct is nil")
	}

	// Just checking for good measure
	cs := contextlib.GetContextStruct(s.ctx)
	if cs == nil {
		contextlib.Logf(s.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Structre not correct")
	}

	go s.goEvent()
}

func (s *SimpleStruct) goEvent() {
	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter ")

	//
	// Using connection managers index as the UniqueID (for now?)
	//

	readchan := s.ConnectionStruct.GetReadChan()
	openchan := s.openChan
	msgbuschan := s.msgbusChan

FORLOOP:
	for {
		select {
		case <-s.Ctx().Done():
			contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Closing down")
			break FORLOOP
		case comm := <-readchan:
			contextlib.Logf(s.ctx, contextlib.LevelInfo, lumerinlib.FileLineFunc()+" ReadChan Event ")

			if comm == nil {
				contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" readchan returned nil")
				break FORLOOP
			}
			//
			// Grossly inefficient... will fix later...
			scre := &SimpleConnReadEvent{}
			scre.uID = ConnUniqueID(comm.Index())
			scre.data = comm.Data()
			scre.count = comm.Count()
			scre.err = comm.Err()

			ev := &SimpleEvent{
				EventType:     ConnReadEvent,
				ConnReadEvent: scre,
				MsgBusEvent:   nil,
				ConnOpenEvent: nil,
			}
			s.eventChan <- ev

		case open := <-openchan:
			contextlib.Logf(s.ctx, contextlib.LevelInfo, lumerinlib.FileLineFunc()+" OpenChan Event ")

			ev := &SimpleEvent{
				EventType:     ConnOpenEvent,
				ConnOpenEvent: open,
				ConnReadEvent: nil,
				MsgBusEvent:   nil,
			}

			s.eventChan <- ev
		case msg := <-msgbuschan:
			contextlib.Logf(s.ctx, contextlib.LevelInfo, lumerinlib.FileLineFunc()+" MsgBusChan Event ")

			if msg == nil {
				contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" msgbuschan returned nil")
				break FORLOOP
			}
			smbe := &SimpleMsgBusEvent{}
			smbe.Data = msg.Data
			smbe.Msg = MsgType(msg.Msg)
			smbe.EventType = EventType(msg.EventType)
			smbe.ID = IDString(msg.ID)
			smbe.RequestID = msg.RequestID
			smbe.Err = msg.Err

			ev := &SimpleEvent{
				EventType:     MsgBusEvent,
				MsgBusEvent:   smbe,
				ConnReadEvent: nil,
				ConnOpenEvent: nil,
			}
			s.eventChan <- ev
		}
	}

	contextlib.Logf(s.ctx, contextlib.LevelWarn, lumerinlib.FileLineFunc()+" exit")
}

//
//
//
func (s *SimpleStruct) Done() bool {
	select {
	case <-s.ctx.Done():
		return true
	default:
		return false
	}
}

//
//
//
func (s *SimpleStruct) Close() {
	if s.Done() {
		return
	}

	s.Cancel()
}

//
//
//
func (s *SimpleStruct) Cancel() {

	if s.cancel == nil {
		contextlib.Logf(s.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" cancel function is nul, struct:%v", s)
		return
	}

	if s.Done() {
		return
	}

	close(s.eventChan)
	close(s.msgbusChan)
	close(s.openChan)
	s.cancel()
}

// Set IO buffer parameters
// this IO buffer parameters apply to the deque used to stage/proess messages
// for stage 1 this can be assumed to be unconfigurable and use defaults only
// func (s *SimpleStruct) SetBuffer() {}

// Set message buffering to a certain delimiter, for example a newline character: ‘\n’
// for stage 1 this will assumed to be unconfigurable and only use '\n' as the
// new line
//func (s *SimpleStruct) SetMessageDelimiterDefault() {}

// Set message buffering to be of a certain size
//func (s *SimpleStruct) SetMessageSizeDefault(mSize uint) {
//	s.maxMessageSize = mSize
//}

// TODO not part of stage 1
// Set encryption parameters
//func (s *SimpleStruct) SetEncryptionDefault() {}

// TODO not part of stage 1
// Set Compression parameters
// func (s *SimpleStruct) SetCompressionDefault() {}

/*
Dial the a destination address (DST)
takes in a net.Addr object and feeds into the net.Dial function
the resulting Conn is then added to the SimpleStructs mapping and and associated
ConnUniqueID is returned from this function

id is the calling ID info, the uid is returned from the connectionmanager layer, and is used to index the connection
*/
func (s *SimpleStruct) AsyncDial(dest *msgbus.Dest) (e error) {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	if s == nil {
		return fmt.Errorf(lumerinlib.FileLineFunc() + " SimpleStruct == nil ")
	}
	if s.ConnectionStruct == nil {
		return fmt.Errorf(lumerinlib.FileLineFunc() + " SimpleStruct.ConnectionStruct == nil ")
	}

	addr, e := dest.NetAddr()
	if e != nil {
		return e
	}

	go func() {

		uid, e := s.ConnectionStruct.Dial(addr)

		open := &SimpleConnOpenEvent{
			uID:  ConnUniqueID(uid),
			dest: dest,
			err:  e,
		}

		if !s.Done() {
			s.openChan <- open
		}
	}()

	return nil
}

//
// AsyncReDial
// Reconnect dropped connection
//
func (s *SimpleStruct) AsyncReDial(uid ConnUniqueID) error {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	if s == nil {
		return errors.New(lumerinlib.FileLineFunc() + " SimpleStruct == nil ")
	}
	if s.ConnectionStruct == nil {
		return errors.New(lumerinlib.FileLineFunc() + " SimpleStruct.ConnectionStruct == nil ")
	}

	go func() {
		e := s.ConnectionStruct.ReDialIdx(int(uid))

		open := &SimpleConnOpenEvent{
			uID: ConnUniqueID(uid),
			err: e,
		}

		if !s.Done() {
			s.openChan <- open
		}
	}()

	return nil
}

/*
function to retrieve the connection mapped to a unique id
*/
func (s *SimpleStruct) GetConnBasedOnConnUniqueID(x ConnUniqueID) *lumerinconnection.LumerinSocketStruct {
	return s.connectionMapping[x]
}

// Used later to direct the default route
func (s *SimpleStruct) SetRoute(u ConnUniqueID) error {
	if u < 0 {
		contextlib.Logf(s.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index out of range:%d", u)
	}
	return s.ConnectionStruct.SetRoute(int(u))
}

// Used later to direct the default route
func (s *SimpleStruct) GetRoute() {} //return of 1 to appease compiler

// Used later to direct the default route
func (s *SimpleStruct) GetLocalAddr(ConnUniqueID) {} //return of 1 to appease compiler

/*

network connection functions

*/

//
//
//
func (s *SimpleStruct) GetRemoteAddr(uid ConnUniqueID) {

}

// Writes buffer to the specified connection
func (s *SimpleStruct) Write(uid ConnUniqueID, msg []byte) (count int, e error) {
	if uid < 0 {
		count, e = s.ConnectionStruct.SrcWrite(msg)
		// Need to so some error checking here, if src is closed, then the whole thing needs to be shutdown.
		if e != nil {
			contextlib.Logf(s.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" ConnectionStruct.SrcRead() error:%s", e)
		}

	} else {
		count, e = s.ConnectionStruct.IdxWrite(int(uid), msg)
	}

	return count, e
}

// Automatic duplication of writes to a MsgBus data channel
func (s *SimpleStruct) DupWrite() {}

// Flushes all IO Buffers
func (s *SimpleStruct) Flush() {}

// Reads low level connection status information
func (s *SimpleStruct) Status() {}

//
// Message Bus Functions
// Pass through to the msgbus package
//
func (s *SimpleStruct) Pub(msgtype MsgType, id IDString, data interface{}) (rid int, e error) {

	//
	// MsgType validation here
	//

	rid, e = s.msgbus.Pub(msgbus.MsgType(msgtype), msgbus.IDString(id), data, s.msgbusChan)
	return rid, e
}

//
//
//
func (s *SimpleStruct) Unpub(msgtype MsgType, id IDString) (rid int, e error) {

	rid, e = s.msgbus.Unpub(msgbus.MsgType(msgtype), msgbus.IDString(id))
	return rid, e
}

//
//
//
func (s *SimpleStruct) Sub(msgtype MsgType, id IDString) (rid int, e error) {

	rid, e = s.msgbus.Sub(msgbus.MsgType(msgtype), msgbus.IDString(id), s.msgbusChan)
	return rid, e
}

//
//
//
func (s *SimpleStruct) Unsub(msgtype MsgType, id IDString) (rid int, e error) {

	rid, e = s.msgbus.Unsub(msgbus.MsgType(msgtype), msgbus.IDString(id), s.msgbusChan)
	return rid, e
}

//
//
//
func (s *SimpleStruct) Get(msgtype MsgType, id IDString) (rid int, e error) {

	rid, e = s.msgbus.Get(msgbus.MsgType(msgtype), msgbus.IDString(id), s.msgbusChan)
	return rid, e
}

//
//
//
func (s *SimpleStruct) Set(msgtype MsgType, id IDString, data interface{}) (rid int, e error) {

	rid, e = s.msgbus.Set(msgbus.MsgType(msgtype), msgbus.IDString(id), data)
	return rid, e
}

//
//
//
func (s *SimpleStruct) SearchIP(msgtype MsgType, search SearchString) (rid int, e error) {

	rid, e = s.msgbus.SearchIP(msgbus.MsgType(msgtype), string(search), s.msgbusChan)
	return rid, e
}

//
//
//
func (s *SimpleStruct) SearchMac(msgtype MsgType, search SearchString) (rid int, e error) {

	rid, e = s.msgbus.SearchMAC(msgbus.MsgType(msgtype), string(search), s.msgbusChan)
	return rid, e
}

//
//
//
func (s *SimpleStruct) SearchName(msgtype MsgType, name SearchString) (rid int, e error) {

	rid, e = s.msgbus.SearchName(msgbus.MsgType(msgtype), string(name), s.msgbusChan)
	return rid, e
}

//
//
//
func (s *SimpleStruct) Ctx() context.Context {
	return s.ctx
}
