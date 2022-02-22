package simple

import (
	"testing"
	//"fmt"
	_ "reflect"
	"context"
	"net"
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

func generateTestContext() (context.Context) {
	returnContext := context.TODO()
	return returnContext
}

func generateTestAddr() net.Addr {
	return testAddr{x:"1"}
}

//function to simulate the protocol layer which will be able to listen for and
//send events to the SIMPL layer
//should run in a go-routine to simulate actual protocol layer
type ProtocolLayer struct {
	ListenStruct SimpleListenStruct
	SimpleStruct SimpleStruct
}

func (p *ProtocolLayer) EventHandler(s *SimpleEvent) {
	//pass the event into the simple struct in the protocol
	//layer to be processed by the SimpleStruct
	p.SimpleStruct.EventHandler(s) 
}

type ProtocolInterface interface {
	EventHandler(*SimpleEvent)
}

func TestSendMessageFromProtocol(t *testing.T) {
	simple, _ := New(generateTestContext(), generateTestAddr())                       //creation of simple layer which provides entry/exit points
	go simple.Run()       //creates a for loop that checks the deque for new messages
	simple.Close()
}

func TestSendMessageFromConnectionLayer(t *testing.T) {
	simple, _ := New(generateTestContext(),generateTestAddr())                       //creation of simple layer which provides entry/exit points
	go simple.Run()       //creates a for loop that checks the deque for new messages
	simple.Close()
}

func TestReceiveMessageFromMSGBus(t *testing.T) {
	simple, _ := New(generateTestContext(),generateTestAddr())                       //creation of simple layer which provides entry/exit points
	go simple.Run()       //creates a for loop that checks the deque for new messages
	simple.Close()
}

// function to route a message from the connection layer to the protocol layer
// this function creates a ConnectionMessage and specifies that it wants to push
// it to the protocol layer
// will be successful is message provided in ConnectionMessage is detected in the ProtocolChan
func TestPushMessageToProtocol(t *testing.T) {
	simple, _ := New(generateTestContext(),generateTestAddr())                       //creation of simple layer which provides entry/exit points
	go simple.Run()       //creates a for loop that checks the deque for new messages
	simple.Close()
}

// function to route a message from the protocol layer to the msgbus
// this function creates a ConnectionMessage and specifies that it wants to push
// it to the protocol layer
// will be successful is message provided in ConnectionMessage is detected in the ProtocolChan
func TestPushMessageToMSGBus(t *testing.T) {
	simple, _ := New(generateTestContext(),generateTestAddr())                       //creation of simple layer which provides entry/exit points
	go simple.Run()       //creates a for loop that checks the deque for new messages
	simple.Close()
}

// function to route a message from the protocol layer to the connection layer
// this function creates a ConnectionMessage and specifies that it wants to push
// it to the protocol layer
// will be successful is message provided in ConnectionMessage is detected in the ProtocolChan
func TestPushMessageToConnectionLayer(t *testing.T) {
	simple, _ := New(generateTestContext(),generateTestAddr())                       //creation of simple layer which provides entry/exit points
	go simple.Run()       //creates a for loop that checks the deque for new messages
	simple.Close()
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
