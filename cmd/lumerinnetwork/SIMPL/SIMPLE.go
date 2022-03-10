package simple

import (
	"context"
	"errors"
	"fmt"
	_"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	_"gitlab.com/TitanInd/lumerin/cmd/log"
	"net"
	_"time"
	//the below packages need to have their gitlab branches sorted out prior to being
	//imported via go mod tidy
	_"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/lumerinconnection"
	//_ "gitlab.com/TitanInd/lumerin/cmd/config"
	_"gitlab.com/TitanInd/lumerin/lumerinlib"
)

/*
The purpose of the simple layer is to accept any messages from the
protocol layer and pass messages to;
lower down the stack
to the message bus
It is also designed to return messages from the msg bus to the protocol
layer.
Refer to proxy router document
https://titanind.atlassian.net/wiki/spaces/PR/pages/5570561/Lumerin+Node
*/

//type SimpleStructProtocolFunc func(*SimpleStruct) chan *SimpleEvent

type ConnUniqueID uint
type URL string
type MsgType string
type ID string
type Data string
type EventHandler string
type SearchString string

type SimpleContextValue string

const SimpleContext SimpleContextValue = "SimpleContextKey"

/*
The simple listen struct is used to establish a Listen port
(TCP, UDP, or TRUNK) and accept connections. The accepted
connections create a SimpleStruct{}, and are passed up to the
protocol layer where the connection is initialized with a new
context, which contains a protocol structure that allows for event handling.
*/
type SimpleListenStruct struct {
	ctx    context.Context
	cancel func()
	accept chan *SimpleStruct //channel to accept simple structs and process their message
}

/*
The simple struct is used to point to a specific instance
of a connection manager and MsgBus. The structure ties these
to a protocol struct where events are directed to be handled.
*/
type SimpleStruct struct {
	ctx    context.Context
	cancel func() //it might make sense to use the WithCancel function instead
	//the event handler portion can be removed since the
	//EventHandler method in implemented on the SimpleStruct
	eventHandler interface{}      //this is a SimpleEvent struct
	eventChan    chan SimpleEvent //channel to listen for simple events
	protocolChan chan SimpleEvent //channel for protocol to receive simple events
	commChan     chan []byte      //channel to listen for simple events
	maxMessageSize uint //this value is not initially set so defaults to 0
}


type EventType string
const NoEvent EventType = "noevent"
const MsgUpdateEvent EventType = "msgupdate"
const MsgDeleteEvent EventType = "msgdelete"
const MsgGetEvent EventType = "msgget"
const MsgGetIndexEvent EventType = "msgindex"
const MsgSearchEvent EventType = "msgsearch"
const MsgSearchIndexEvent EventType = "msgsearchindex"
const MsgPublishEvent EventType = "msgpublish"
const MsgUnpublishEvent EventType = "msgunpublish"
const MsgSubscribedEvent EventType = "msgsubscribe"
const MsgUnsubscribedEvent EventType = "msgunsubscribe"
const MsgRemovedEvent EventType = "msgremoved"
const ConnReadEvent EventType = "connread"
const ConnEOFEvent EventType = "conneof"
const ConnErrorEvent EventType = "connerror"
const ErrorEvent EventType = "error"
const MsgToProtocol EventType = "msgUp"

// this is a temporary function used to initialize a SimpleListenStruct
func dummyFunc() {}

// this is a dummy interface
type dummyInterface interface {
	dummy()
}

type dummyStruct struct {
}

func (d *dummyStruct) dummy() {
}

/*
create and return a struct with channels to listen to
call goroutine embedded in the struct
*/
func New(ctx context.Context, listen net.Addr, newproto interface{}) (SimpleListenStruct, error) {
	myStruct := SimpleListenStruct{
		ctx:    ctx,
		cancel: dummyFunc, //need to replace dummy func with an actual cancel function
		accept: make(chan *SimpleStruct),
	}
	// determine if a more robust error message is needed
	return myStruct, nil
}

//consider calling this as a gorouting from protocol layer, assuming
//protocll layer will have a layer to communicate with a chan over
func (s *SimpleListenStruct) Run() error {
	go func() {
		// continuously listen for messages coming in on the accept channel
		for {
			//consider moving event handler login into here
			x := <-s.accept //receive a value from the accept
			fmt.Printf("%+v", x)
		}
	}()
	return nil
}

func (s *SimpleListenStruct) Accept() <-chan *SimpleStruct {
	return s.accept
}

// replacing the channel with a return statement containing the new simple struct
func (s *SimpleListenStruct) NewSimpleStruct(ctx context.Context) {
	go func() {
		myStruct := &SimpleStruct{ //generate a new SimpleStruct
			ctx:          ctx,
			cancel:       dummyFunc,
			eventHandler: dummyStruct{},
			eventChan:    make(chan SimpleEvent),
		}
		s.accept <- myStruct //push a SimpleStruct onto the SimpleListenStruct's accept channel
	}()
}

// Calls the listen context cancel function, which closes out the listener routine
func (s *SimpleListenStruct) Close() {
	_, cancel := context.WithCancel(s.ctx)
	cancel() //cancel is a function which terminates the associated goroutine
}

/*
Start a new go routine to handle the new connection context
after initialization by the protocol layer. There will be a
variable in the context that points to the protocol structure
containing all of the pertinent data for the state of the protocol
and event handler routines
All of the SimpleStruct functions that follow can be called
before and after Run() is called
It is assumed that Run() can only be called once
*/
/*
TODO pass context to SimpleListenStruct's designated connection layer
*/
func (s *SimpleStruct) Run(c context.Context) error {
	// loop to continuously listen for messages coming in
	// on the channels assigned to the connection layer
	// and the msgbus

	if s.maxMessageSize == 0 {
		s.maxMessageSize = 10 //setting the default max message size to 10 bytes
	}

	var res error

	go func() {
		for {
			select {
			case x := <-s.commChan:
				//create SimpleEvent and pass to event handler
				newMessage := SimpleEvent{
					EventType: MsgToProtocol,
					Data:      x,
				}
				s.EventHandler(newMessage)
			default:
				res = errors.New("error in receiving commchan value")
				return
			}
		}
	}()
	return res
}

/*
Calls the connection context cancel function which closes out the
currently established SRC connection and all of the associated DST connections
*/
func (s *SimpleStruct) Close() {
	_, cancel := context.WithCancel(s.ctx)
	cancel()
}

// Set IO buffer parameters
// this IO buffer parameters apply to the deque used to stage/proess messages
// for stage 1 this can be assumed to be unconfigurable and use defaults only
func (s *SimpleStruct) SetBuffer() {}

// Set message buffering to a certain delimiter, for example a newline character: ‘\n’
// for stage 1 this will assumed to be unconfigurable and only use '\n' as the 
// new line
func (s *SimpleStruct) SetMessageDelimiterDefault() {}

// Set message buffering to be of a certain size
func (s *SimpleStruct) SetMessageSizeDefault(mSize uint) {
	s.maxMessageSize = mSize
}

// TODO not part of stage 1
// Set encryption parameters
func (s *SimpleStruct) SetEncryptionDefault() {}

// TODO not part of stage 1
// Set Compression parameters
func (s *SimpleStruct) SetCompressionDefault() {}

// Dial the a destination address (DST)
func (s *SimpleStruct) Dial(dst net.Addr) (ConnUniqueID, error) { return 0, nil } //return of 1 to appease compiler

// Reconnect dropped connection
func (s *SimpleStruct) Redial(u ConnUniqueID) {} 

// Used later to direct the default route
func (s *SimpleStruct) SetRoute(u ConnUniqueID) {} 

// Used later to direct the default route
func (s *SimpleStruct) GetRoute() {} //return of 1 to appease compiler

// Used later to direct the default route
func (s *SimpleStruct) GetLocalAddr(ConnUniqueID) {} //return of 1 to appease compiler

/*

network connection functions

*/

// Used later to direct the default route
func (s *SimpleStruct) GetRemoteAddr(ConnUniqueID) {} //return of 1 to appease compiler

func (s *SimpleStruct) SetDefaultReadHandler() {}

// Supply a handler function for incoming data reads for the connection ID
func (s *SimpleStruct) SetReadHandler() {}

// Writes buffer to the specified connection
func (s *SimpleStruct) Write(ConnUniqueID, []byte) {}

// Automatic duplication of writes to a MsgBus data channel
func (s *SimpleStruct) DupWrite() {}

// Flushes all IO Buffers
func (s *SimpleStruct) Flush() {}

// Reads low level connection status information
func (s *SimpleStruct) Status() {}

/*

msg bus functions

*/

func (s *SimpleStruct) Pub(MsgType, ID, Data) error            { return errors.New("") }
func (s *SimpleStruct) Unpub(MsgType, ID) error                { return errors.New("") }
func (s *SimpleStruct) Sub(MsgType, ID, EventHandler) error    { return errors.New("") }
func (s *SimpleStruct) Unsub(MsgType, ID, EventHandler) error  { return errors.New("") }
func (s *SimpleStruct) Get(MsgType, ID, EventHandler) error    { return errors.New("") }
func (s *SimpleStruct) Set(MsgType, ID, Data) error            { return errors.New("") }
func (s *SimpleStruct) SearchIP(MsgType, SearchString) error   { return errors.New("") }
func (s *SimpleStruct) SearchMac(MsgType, SearchString) error  { return errors.New("") }
func (s *SimpleStruct) SearchName(MsgType, SearchString) error { return errors.New("") }

func (ss *SimpleStruct) Ctx() context.Context {
	return ss.ctx
}

func (ss *SimpleStruct) Cancel() {
	ss.cancel()
}

/*


event handler related functionality


*/

// type EventType string

var eventOne EventType = "eventOne"

type SimpleEvent struct {
	EventType EventType
	Data      interface{}
}

//event handler function for the SimpleStruct which is viewable from the protocol layer
func (s *SimpleStruct) EventHandler(e SimpleEvent) {
	for {
		switch e.EventType {
		case NoEvent:
			fallthrough
		case MsgUpdateEvent:
			fallthrough
		case MsgDeleteEvent:
			fallthrough
		case MsgGetEvent:
			fallthrough
		case MsgGetIndexEvent:
			fallthrough
		case MsgSearchEvent:
			fallthrough
		case MsgSearchIndexEvent:
			fallthrough
		case MsgPublishEvent:
			fallthrough
		case MsgUnpublishEvent:
			fallthrough
		case MsgSubscribedEvent:
			fallthrough
		case MsgUnsubscribedEvent:
			fallthrough
		case MsgRemovedEvent:
			fallthrough
		case ConnReadEvent:
			fallthrough
		case ConnEOFEvent:
			fallthrough
		case ConnErrorEvent:
			fallthrough
		case ErrorEvent:
			fallthrough
		case MsgToProtocol:
			fallthrough
		default:
			return
		}
	}
}
