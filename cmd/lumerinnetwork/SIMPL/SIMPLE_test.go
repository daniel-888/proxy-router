package simple

import (
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

func protocolMessage() ProtocolMessage {
	return ProtocolMessage{
		WorkerName:      "Josh's magic money laundering machine",
		MessageContents: []byte("I wish I was as cool as Josh"), //create a byte stringb
		MessageActions:  []uint{1, 2},
	}
}

func msgbusMessage() MSGBusMessage {
	return MSGBusMessage{
		WorkerName:      "Josh's magic money laundering machine",
		MessageContents: []byte("I wish I was as cool as Josh"), //create a byte stringb
		MessageActions:  []uint{1, 2},
	}
}

func connectionMessage() ConnectionMessage {
	return ConnectionMessage{
		WorkerName:      "Josh's magic money laundering machine",
		MessageContents: []byte("I wish I was as cool as Josh"), //create a byte stringb
		MessageActions:  []uint{1, 2},
	}
}

/*
below are the actual tests
*/

func TestSendMessageFromProtocol(t *testing.T) {
	simple := New()      //creation of simple layer which provides entry/exit points
	simple.ActivateSIMPLELayer()       //creates a for loop that checks the deque for new messages
	simple.ListenForIncomingMessages() //creates a for loop that listens to all channels
	ProtocolChan := simple.ProtocolChan
	pMessage := protocolMessage()
	ProtocolChan <- pMessage
	simple.Close()
}

func TestSendMessageFromConnectionLayer(t *testing.T) {
	simple := New()      //creation of simple layer which provides entry/exit points
	simple.ActivateSIMPLELayer()       //creates a for loop that checks the deque for new messages
	simple.ListenForIncomingMessages() //creates a for loop that listens to all channels
	ConnectionChan := simple.ConnectionChan
	lMessage := connectionMessage()
	ConnectionChan <- lMessage
	simple.Close()
}

func TestReceiveMessageFromMSGBus(t *testing.T) {
	simple := New()      //creation of simple layer which provides entry/exit points
	simple.ActivateSIMPLELayer()       //creates a for loop that checks the deque for new messages
	simple.ListenForIncomingMessages() //creates a for loop that listens to all channels
	MsgBusChan := simple.MSGChan
	mMessage := msgbusMessage()
	MsgBusChan <- mMessage
	simple.Close()
}

func TestPushMessageToProtocol(t *testing.T) {
}

func TestPushMessageToConnectionLayer(t *testing.T) {
}

func TestPushMessageToMSGBus(t *testing.T) {
}

func TestHashrateCountMessage(t *testing.T) {
}

func TestValidationRequestMessage(t *testing.T) {
}

func TestMessageFrom3Sources(t *testing.T) {
}

func TestMultipleMessagesFromProtocol(t *testing.T) {
}

func TestMultipleMessagesFromConnectionLayer(t *testing.T) {
}

func TestCorruptMessage(t *testing.T) {
}

func TestMessageWithInvalidActions(t *testing.T) {
}
