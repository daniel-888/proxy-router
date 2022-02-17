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
		MessageContents: []byte("I wish I was as cool as Josh"), //create a byte stringb
		MessageActions:  actions,
	}
}

func msgbusMessage(actions []uint) MSGBusMessage {
	return MSGBusMessage{
		WorkerName:      "Josh's magic money laundering machine",
		MessageContents: []byte("I wish I was as cool as Josh"), //create a byte stringb
		MessageActions:  actions,
	}
}

func connectionMessage(actions []uint) ConnectionMessage {
	return ConnectionMessage{
		WorkerName:      "Josh's magic money laundering machine",
		MessageContents: []byte("I wish I was as cool as Josh"), //create a byte stringb
		MessageActions:  actions,
	}
}

/*
below are the actual tests
*/

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
	var pm ProtocolMessage
	/*
		break this out into its own function for other tests to use
	*/
	go func() {

		count := 0
		for {
			temp := <-simple.ProtocolChan
			if string(temp.MessageContents) == "I Wish I was as cool as Josh" {
				pm = temp
				break
			} else if count > 10 {
				break
			}
			time.Sleep(time.Second * 1)
			count++
		}
	}()
	if string(pm.MessageContents) != "I wish I was as cool as Josh" {
		t.Errorf("messages not being sent to protocol chan.\nActual:%s\nExpected:%s", pm.MessageContents, "I Wish I was as cool as Josh")
	}
	//simple.Close()
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
