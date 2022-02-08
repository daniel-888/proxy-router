package connectionmanager

import (
	"fmt"
	_ "net"
	_ "time"

	_ "gitlab.com/TitanInd/lumerin/cmd/config"
	_ "gitlab.com/TitanInd/lumerin/cmd/msgbus"
	_ "gitlab.com/TitanInd/lumerin/lumerinlib"
)

/*
The purpose of the simple layer is to accept any messages from the
protocol layer and pass messages to;
lower down the stack
to the message bus
It is also designed to return messages from the msg bus to the protocol
layer.
Refer to proxy router document section 7.7.2.2
*/

/*
high level thinking there should be a deque to push messages onto
this will operate in a first in - first out manner where messages will
take precedence based on their seniority
The goal is to turn this into an asynchronous process so many messages can be
handled at once
*/
type messageAction string
type actionResponse string

//create a goroutine which has 3 channels associated with it
//channel 0: accessable by any protocol layer
//channel 1: accessable by the msg bus
//channel 2: accessable from the lowe layers in the stack
/*
MsgDeque is a last in first out datastructue which can accept
messages of any struct type and in constantly processed
*/
type SIMPLE struct {
	ProtocolChan chan ProtocolMessage
	MSGChan      chan MSGBusMessage
	LowerChan    chan LowerMessage
	MsgDeque     []interface{}
}

//struct to handle/accept messages from the layer 1 channel
type ProtocolMessage struct {
	WorkerName      string
	MessageContents []byte
	MessageActions  []string
}

// struct to handle/accept messages from the message bus
type MSGBusMessage struct {
	WorkerName      string
	MessageContents []byte
	MessageActions  []string
}

// struct to handle messages from further down in the stack
type LowerMessage struct {
	WorkerName      string
	MessageContents []byte
	MessageActions  []string
}

//define available actions
const (
	//constants to define requested incoming messages
	HashSubmit    messageAction = "HashSubmit"
	HashrateCount messageAction = "HashrateCount"
)

//define available return messages
const (
	HashValid     actionResponse = "HashValid"
	HashInvalid   actionResponse = "HashInvalid"
	HashrateValue actionResponse = "HashrateValue"
)

//function to constantly monitor MsgDeque and process the last item on it
func (s *SIMPLE) ActivateSIMPLELayer() {
	go func() {
		if len(s.MsgDeque) > 0 {
			//msg is the last element in the msg deque and is processed
			//newDeque is to rewrite the MsgDeque in lieu of another popping method
			msg, newDeque := s.MsgDeque[0], s.MsgDeque[1:]
			processIncomingMessage(msg)
			s.MsgDeque = newDeque

		}
	}()
}

//listens for messages coming in through the various channels
//oldest item will always be index 0
func (s *SIMPLE) ListenForIncomingMessages() {
	go func() {
		for {
			select {
			case pc := <-s.ProtocolChan:
				s.MsgDeque = append(s.MsgDeque, pc)
			case mc := <-s.ProtocolChan:
				s.MsgDeque = append(s.MsgDeque, mc)
			case lc := <-s.ProtocolChan:
				s.MsgDeque = append(s.MsgDeque, lc)
			}
		}
	}()
}

/*
this function is where the majority of the work for the SIMPLE layer will be done
Each message coming in will have a [task] field which tells the SIMPLE layer
how to process the message. The idea here is that anybody can create a [task] and
associated function and add to the processing request.
Rules to follow
1. this is a function, so for every input there's only 1 output
2. do not break the interface of the output
3. design functions in a maintainable manner
*/
func processIncomingMessage(m interface{}) {
}

/*
create and return a struct with channels to listen to
call goroutine embedded in the struct
*/
func InitializeSIMPLELayer() SIMPLE {
	var deque []interface{}
	return SIMPLE{
		ProtocolChan: make(chan ProtocolMessage),
		MSGChan:      make(chan MSGBusMessage),
		LowerChan:    make(chan LowerMessage),
		MsgDeque:     deque,
	}
}

//create a listener for the msg bus
