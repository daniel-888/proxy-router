package simple

import (
	"context" //this can probably be removed once gitlab packages can be imported
	"errors"
	"fmt"
	"net"
	_ "time"
	//the below packages need to have their gitlab branches sorted out prior to being
	//imported via go mod tidy
	//_ "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/lumerinconnection"
	//double check that these imports were formatted correctly
	//_ "gitlab.com/TitanInd/lumerin/cmd/config"
	//_ "gitlab.com/TitanInd/lumerin/cmd/msgbus"
	//_ "gitlab.com/TitanInd/lumerin/lumerinlib"
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

type ConnUniqueID uint
type URL string
type MsgType string
type ID string
type Data string
type EventHandler string
type SearchString string

type ContextValue string
type EventType string

const SimpleMsgBusValue ContextValue = "MSGBUS"
const SimpleSrcAddrValue ContextValue = "SRCADDR"
const SimpleDstAddrValue ContextValue = "DSTADDR"

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

// takes the byte array destined for the protocol layer and unmarshals it into a ProtocolMessage struct
// then it pushes the ProtocolMessage onto the ProtocolChan
func (s *SimpleStruct) msgToProtocol(b []byte) {
	//create an in-memory temporary struct to pass to the ProtocolChan
	//pass the struct to the protocol chan
}

// takes the byte array destined for the protocol layer and unmarshals it into a MSGBusMessage struct
// then it pushes the MSGBusMessage onto the MSGChan
func (s *SimpleStruct) msgToMSGBus(b []byte) {
	//create an in-memory temporary struct to pass to the MSGChan
	//pass the struct to the protocol chan
}

// takes the byte array destined for the protocol layer and unmarshals it into a ConnectionMessage struct
// then it pushes the ConnectionMessage onto the ConnectionChan
func (s *SimpleStruct) msgToConnection(b []byte) {
	//create an in-memory temporary struct to pass to the ConnectionChan
	//pass the struct to the protocol chan
}

/*
this function is where the majority of the work for the SIMPLE layer will be done
*/
func (s *SimpleStruct) processIncomingMessage(m uint) {
	switch m {
	case 0: //route message to protocol channel
		s.msgToProtocol([]byte{})
	case 1: //route message to msgbus channel
		s.msgToMSGBus([]byte{})
	case 2: //route message to connection channel
		s.msgToConnection([]byte{})
	default:
		fmt.Println("lord bogdanoff demands elon tank the price of dogecoin")
	}
}

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
func New(ctx context.Context, listen net.Addr) (SimpleListenStruct, error) {
	myStruct := SimpleListenStruct{
		ctx:    ctx,
		cancel: dummyFunc,
		accept: make(chan *SimpleStruct),
	}
	// determine if a more robust error message is needed
	// return myStruct, errors.New("unable to create a SimpleListenStruct")
	return myStruct, nil
}

func NewSimpleStruct(ctx context.Context) (SimpleStruct, error) {
	myStruct := SimpleStruct{
		ctx:          ctx,
		cancel:       dummyFunc,
		eventHandler: dummyStruct{},
		eventChan:    make(chan SimpleEvent),
	}
	// determine if a more robust error message is needed
	return myStruct, errors.New("unable to create a SimpleListenStruct")
}

func (s *SimpleListenStruct) Run() error {
	go func() {
		// continuously listen for messages coming in on the accept channel
		for {
			//consider moving event handler login into here
			x := <-s.accept //receive a value from the accept
			fmt.Printf("%+v", x)
		}
	}()
	return errors.New("meow")
}

func (s *SimpleListenStruct) Accept() <-chan *SimpleStruct {
	return s.accept
}

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
func (s *SimpleStruct) Run(c context.Context) {
	msgDeque := []byte{}
	go func() {
		for {
			fmt.Printf("%+v", msgDeque)
		}
	}()
}

/*
Calls the connection context cancel function which closes out the
currently established SRC connection and all of the associated DST connections
*/
func (s *SimpleStruct) Close() {}

// Set IO buffer parameters
func (s *SimpleStruct) SetBuffer() {}

// Set message buffering to a certain delimiter, for example a newline character: ‘\n’
func (s *SimpleStruct) SetMessageDelimiterDefault() {}

// Set message buffering to be of a certain size
func (s *SimpleStruct) SetMessageSizeDefault() {}

// TODO not part of stage 1
// Set encryption parameters
func (s *SimpleStruct) SetEncryptionDefault() {}

// TODO not part of stage 1
// Set Compression parameters
func (s *SimpleStruct) SetCompressionDefault() {}

// Dial the a destination address (DST)
func (s *SimpleStruct) Dial(u URL) ConnUniqueID { return 1 } //return of 1 to appease compiler

// Reconnect dropped connection
func (s *SimpleStruct) Redial(u ConnUniqueID) {} //return of 1 to appease compiler

// Used later to direct the default route
func (s *SimpleStruct) SetRoute(u ConnUniqueID) {} //return of 1 to appease compiler

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

/*
The simple listen struct is used to establish a Listen port
(TCP, UDP, or TRUNK) and accept connections. The accepted
connections create a SimpleStruct{}, and are passed up to the
protocol layer where the connection is initialized with a new
context, which contains a protocol structure that allows for event handling.
*/
/*
accept chan should have a SimpleStruct pushed into it when creating a
new SimpleStruct for an individual connection
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
	ctx          context.Context
	cancel       func()           //it might make sense to use the WithCancel function instead
	eventHandler interface{}      //this is a SimpleEvent struct
	eventChan    chan SimpleEvent //channel to listen for simple events
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

type SimpleProtocolInterface interface {
	EventHandler(*SimpleEvent)
}

//event handler function for the SimpleStruct which is viewable from the protocol layer
func (s *SimpleStruct) EventHandler() {
	for {
		x := <-s.eventChan
		switch x.EventType {
		case eventOne:
			fallthrough
		default:
			return
		}
	}
}
