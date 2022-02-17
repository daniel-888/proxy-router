package simple

import (
	"testing"
	"time"
	//"fmt"
	_ "reflect"
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

func protocolMessage(actions []uint) ProtocolMessage {
	return ProtocolMessage{
		WorkerName:      "Josh's magic money laundering machine",
		MessageContents: []byte("Test Sentence"), //create a byte stringb
		MessageActions:  actions,
	}
}

func msgbusMessage(actions []uint) MSGBusMessage {
	return MSGBusMessage{
		WorkerName:      "Josh's magic money laundering machine",
		MessageContents: []byte("Test Sentence"), //create a byte stringb
		MessageActions:  actions,
	}
}

func connectionMessage(actions []uint) ConnectionMessage {
	return ConnectionMessage{
		WorkerName:      "Josh's magic money laundering machine",
		MessageContents: []byte("Test Sentence"), //create a byte stringb
		MessageActions:  actions,
	}
}

/*
below are the actual tests
*/

func (s *SIMPLE) listenToProtocolChan() []byte {
	var pm ProtocolMessage
	go func() {
		count := 0
		for {
			temp := <-s.ProtocolChan
			if string(temp.MessageContents) == "Test Sentence" {
				pm = temp
				break
			} else if count > 10 {
				break
			}
			time.Sleep(time.Second * 1)
			count++
		}
	}()
	return pm.MessageContents
}

func (s *SIMPLE) listenToMSGChan() []byte {
	var pm MSGBusMessage
	go func() {
		count := 0
		for {
			temp := <-s.MSGChan
			if string(temp.MessageContents) == "Test Sentence" {
				pm = temp
				break
			} else if count > 10 {
				break
			}
			time.Sleep(time.Second * 1)
			count++
		}
	}()
	return pm.MessageContents
}

func (s *SIMPLE) listenToConnectionChan() []byte {
	var pm ConnectionMessage
	go func() {
		count := 0
		for {
			temp := <-s.ConnectionChan
			if string(temp.MessageContents) == "Test Sentence" {
				pm = temp
				break
			} else if count > 10 {
				break
			}
			time.Sleep(time.Second * 1)
			count++
		}
	}()
	return pm.MessageContents
}


func TestSendMessageFromProtocol(t *testing.T) {
	simple := New()                       //creation of simple layer which provides entry/exit points
	go simple.ActivateSIMPLELayer()       //creates a for loop that checks the deque for new messages
	go simple.ListenForIncomingMessages() //creates a for loop that listens to all channels
	pMessage := protocolMessage([]uint{0, 1})
	simple.ProtocolChan <- pMessage
	time.Sleep(time.Second * 3)
	simple.Close()
}

func TestSendMessageFromConnectionLayer(t *testing.T) {
	simple := New()                       //creation of simple layer which provides entry/exit points
	go simple.ActivateSIMPLELayer()       //creates a for loop that checks the deque for new messages
	go simple.ListenForIncomingMessages() //creates a for loop that listens to all channels
	lMessage := connectionMessage([]uint{0, 1})
	simple.ConnectionChan <- lMessage
	time.Sleep(time.Second * 3)
	simple.Close()
}

func TestReceiveMessageFromMSGBus(t *testing.T) {
	simple := New()                       //creation of simple layer which provides entry/exit points
	go simple.ActivateSIMPLELayer()       //creates a for loop that checks the deque for new messages
	go simple.ListenForIncomingMessages() //creates a for loop that listens to all channels
	mMessage := msgbusMessage([]uint{0, 1})
	simple.MSGChan <- mMessage
	time.Sleep(time.Second * 3)
	simple.Close()
}

// function to route a message from the connection layer to the protocol layer
// this function creates a ConnectionMessage and specifies that it wants to push
// it to the protocol layer
// will be successful is message provided in ConnectionMessage is detected in the ProtocolChan
func TestPushMessageToProtocol(t *testing.T) {
	simple := New()                       //creation of simple layer which provides entry/exit points
	go simple.ListenForIncomingMessages() //creates a for loop that listens to all channels
	go simple.ActivateSIMPLELayer()       //creates a for loop that checks the deque for new messages
	cm := connectionMessage([]uint{0})    //creates a connection message
	simple.ConnectionChan <- cm           //pushing the connection message to the connection chan
	pm := simple.listenToProtocolChan()
	if string(pm) != "Test Sentence" {
		t.Errorf("messages not being sent to protocol chan.\nActual:%s\nExpected:%s", pm, "Test Sentence")
	}
	simple.Close()
}

// function to route a message from the protocol layer to the msgbus
// this function creates a ConnectionMessage and specifies that it wants to push
// it to the protocol layer
// will be successful is message provided in ConnectionMessage is detected in the ProtocolChan
func TestPushMessageToMSGBus(t *testing.T) {
	simple := New()                       //creation of simple layer which provides entry/exit points
	go simple.ListenForIncomingMessages() //creates a for loop that listens to all channels
	go simple.ActivateSIMPLELayer()       //creates a for loop that checks the deque for new messages
	cm := protocolMessage([]uint{1})    //creates a connection message
	simple.ProtocolChan <- cm           //pushing the connection message to the connection chan
	pm := simple.listenToMSGChan()
	if string(pm) != "Test Sentence" {
		t.Errorf("messages not being sent to protocol chan.\nActual:%s\nExpected:%s", pm, "Test Sentence")
	}
	simple.Close()
}

// function to route a message from the protocol layer to the connection layer
// this function creates a ConnectionMessage and specifies that it wants to push
// it to the protocol layer
// will be successful is message provided in ConnectionMessage is detected in the ProtocolChan
func TestPushMessageToConnectionLayer(t *testing.T) {
	simple := New()                       //creation of simple layer which provides entry/exit points
	go simple.ListenForIncomingMessages() //creates a for loop that listens to all channels
	go simple.ActivateSIMPLELayer()       //creates a for loop that checks the deque for new messages
	cm := protocolMessage([]uint{2})    //creates a connection message
	simple.ProtocolChan <- cm           //pushing the connection message to the connection chan
	pm := simple.listenToConnectionChan()
	if string(pm) != "Test Sentence" {
		t.Errorf("messages not being sent to protocol chan.\nActual:%s\nExpected:%s", pm, "Test Sentence")
	}
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
