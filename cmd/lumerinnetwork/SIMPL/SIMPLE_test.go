package simple

import (
	"context"
	_ "fmt"
	"net"
	_ "reflect"
	"testing"
)

/*
testing system for the SIMPLE layer
The SIMPLE layer will need to be tested to
ensure that is can route packets depending on outside instructions
As additional functionality is added to the SIMPLE layer
(multithreading, compressing, rerouting, error handling, etc)
additional feature and functional tests will need to be written.

The Stratum layer, lower level layers, and MSG are not being tested in this
testing suite, however their messages may either be used or simulated for testing purposes
*/

/*
this is just a cookie cutter placehodler so new tests
can be quickly implemented
func TestTemplate(t *testing.T) {
}
*/

type testAddr struct {
	x string
}

func (t testAddr) Network() string {
	return t.x
}

func (t testAddr) String() string {
	return t.x
}

func generateTestContext() context.Context {
	returnContext := context.TODO()
	return returnContext
}

func generateTestAddr() net.Addr {
	return testAddr{x: "1"}
}

type ConnectionLayer struct {
	SimpleConnection *SimpleStruct
}

func NewConnLayer(s *SimpleStruct) ConnectionLayer {
	return ConnectionLayer{
		SimpleConnection: s,
	}
}

func (c *ConnectionLayer) ConnToSimple() {
	go func() {
		c.SimpleConnection.commChan <- []byte("test message one")
	}()
}

//function to simulate the protocol layer which will be able to listen for and
//send events to the SIMPL layer
//should run in a go-routine to simulate actual protocol layer
type ProtocolLayer struct {
	ListenStruct SimpleListenStruct
	SimpleStruct SimpleStruct
}

type ProtocolInterface interface {
	EventHandler(*SimpleEvent)
}

//generate a simplestruct for testing purposes
func generateSimpleListenStruct() SimpleListenStruct {
	myContext := generateTestContext()
	myAddr := generateTestAddr()
	myStruct, _ := New(myContext, myAddr)
	return myStruct
}

//generate a protocol layer for testing purposes
func generateProtocolLayer() ProtocolLayer {
	listenStruct := generateSimpleListenStruct()
	simpleStruct, _ := NewSimpleStruct(generateTestContext())
	return ProtocolLayer{
		ListenStruct: listenStruct,
		SimpleStruct: simpleStruct,
	}
}

// this is a basic test to create a SimpleListenEvent
// then have the connection layer request a SimpleStruct
// then have the SimpleStruct initialize a ProtocolStruct
// success will be determined by ensuring the communication layer
// simple struct, and protocol layer know about eachother
/*
test steps
1. protocol layer will initialize a SimpleListenStruct
2. protocol layer calls the New function in the SIMPLE package
	a context and an Addr are passed into New
3. protocol layer calls run method on SimpleListenStruct

*/
func TestInitializeSimpleListenStruct(t *testing.T) {
	listenStruct, _ := New(generateTestContext(), generateTestAddr())
	listenStruct.Run()

}

//send a message from the protocol layer to the simple layer
func TestSendMessageFromProtocolToConnectionLayer(t *testing.T) {
	/*
		test steps
		1. create simulated protocol layer
		2. call function on listening struct to create a new simple struct
		3. listen for simple struct on listen structs accept channel
		4. initialize the event handler on the simple struct
		5. create a dummy SimpleEvent (this could be extentiated outside of the test)
		6. pass the SimpleEvent into the eventChan channel
		7. close both structs

		TODO check to see value of SimpleEvent within connection layer
	*/
	pc := generateProtocolLayer()
	listenStruct := pc.ListenStruct
	listenStruct.NewSimpleStruct(generateTestContext())
	simpleStruct := <-listenStruct.accept
	event := SimpleEvent{ //create a simpleEvent to pass into event chan
		EventType: eventOne,
		Data:      []byte{},
	}
	simpleStruct.EventHandler(event)
	simpleStruct.Close()
	listenStruct.Close()

}

//test to initialize a simple layer and protocol layer
//connection layer will send byte information to the
//SimpleStruct and the SimpleStruct will pass that information upwards
//to the protocol layer
//message will be byte array of string "test sentence one"
func TestSendMessageFromConnectionLayer(t *testing.T) {
	pc := generateProtocolLayer()
	listenStruct := pc.ListenStruct
	listenStruct.NewSimpleStruct(generateTestContext())
	simpleStruct := <-listenStruct.accept
	go simpleStruct.Run(listenStruct.ctx)

	connLayer := NewConnLayer(simpleStruct)
	connLayer.ConnToSimple()

	// connMsg := <- simpleStruct.protocolChan
	// fmt.Printf("Event Type: %s", connMsg.EventType)
	//getting an issue where Data is being considered an Interface instead of a byetstring
	// if "test sentence one" != string(connMsg.Data) {
	//	t.Error("msg came out wrong",err)
	//}
}

func TestReceiveMessageFromMSGBus(t *testing.T) {
}

// function to route a message from the connection layer to the protocol layer
// this function creates a ConnectionMessage and specifies that it wants to push
// it to the protocol layer
// will be successful is message provided in ConnectionMessage is detected in the ProtocolChan
func TestPushMessageToProtocol(t *testing.T) {
}

// function to route a message from the protocol layer to the msgbus
// this function creates a ConnectionMessage and specifies that it wants to push
// it to the protocol layer
// will be successful is message provided in ConnectionMessage is detected in the ProtocolChan
func TestPushMessageToMSGBus(t *testing.T) {
}

// function to route a message from the protocol layer to the connection layer
// this function creates a ConnectionMessage and specifies that it wants to push
// it to the protocol layer
// will be successful is message provided in ConnectionMessage is detected in the ProtocolChan
func TestPushMessageToConnectionLayer(t *testing.T) {
}

func TestHashrateCountMessage(t *testing.T) {
}

func TestValidationRequestMessage(t *testing.T) {
}

//send the following messages
// 1. message from connection layer to protocol layer
// 2. message from msg.bus to protocol layer
// 3. message from protocol layer to msgbus
// this test will be considered successful if the messages
// are processed in order and also make it to their final destination
func TestMessageFrom3Sources(t *testing.T) {
}

//send the following messages
// 1. message from protocol to connection layer
// 2. message from protocol to msgbus
// this test will be considered successful if the messages are processed in order
// and make it to their intended destinations
func TestMultipleMessagesFromProtocol(t *testing.T) {
}

//send the following messages
// 1. message from protocol to connection layer
// 2. message from protocol to msgbus
// this test will be considered successful if the messages are processed in order
// and make it to their intended destinations
func TestMultipleMessagesFromConnectionLayer(t *testing.T) {
}

// test to send a corrupted message through the SIMPLE layer
// the corrputed message should go through as expected since
// the simple layer doesn't check for message integrity
func TestCorruptMessage(t *testing.T) {
}

//send a message from the connection channel with an option of 100
//this will not be picked up by any cases in the processIncomingMessage function
func TestMessageWithInvalidActions(t *testing.T) {
}
