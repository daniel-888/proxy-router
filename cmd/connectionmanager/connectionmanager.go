package connectionmanager

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

type connectionStates int

// const RECVBUFSIZE int =	2048

const (
	stateNone connectionStates = iota
	stateOpen
	stateAuth
	stateHashVerify
	stateRouting
	stateConnecting
	stateConnected
	stateRouteChange
	stateMgsError
	stateConnectError
	stateSrcClosed
	stateDstClosed
	stateShutdown
	stateClosed
)

type ConnectionManager struct {
	ps *msgbus.PubSub
}

type msgBuffer []byte

//type msgBuffer [RECVBUFSIZE]byte

type connection struct {
	ps                  *msgbus.PubSub
	state               connectionStates
	srcSocketConn       net.Conn
	dstSocketConn       net.Conn
	srcSocketChan       chan msgBuffer
	dstSocketChan       chan msgBuffer
	connectionID        msgbus.ConnectionID
	eventConnectionChan msgbus.EventChan
	minerID             msgbus.MinerID
	eventMinerChan      msgbus.EventChan
	srcMsgBuf           msgBuffer
	dstMsgBuf           msgBuffer
	connectionEvent     msgbus.Event
	minerEvent          msgbus.Event
}

//------------------------------------------
//
//------------------------------------------
func New(ps *msgbus.PubSub) (cm *ConnectionManager, err error) {
	cm = &ConnectionManager{
		ps: ps,
	}
	return cm, err
}

//------------------------------------------
//
//------------------------------------------
func newConnection(conn net.Conn, ps *msgbus.PubSub) *connection {
	c := connection{}
	c.ps = ps
	c.state = stateNone
	c.srcSocketConn = conn
	c.dstSocketConn = nil
	c.srcSocketChan = nil
	c.dstSocketChan = nil
	c.eventMinerChan = nil
	c.eventConnectionChan = nil
	c.minerID = ""
	c.connectionID = ""
	c.srcMsgBuf = nil
	c.dstMsgBuf = nil
	c.connectionEvent = ps.NewEvent()
	c.minerEvent = ps.NewEvent()

	return &c
}

//---------------------------------------------
//
//---------------------------------------------
func (c *connection) runSrcSocketReader() {

	//go func(socket net.Conn, ch chan<- msgBuffer) {
	go func(c *connection) {

		defer func() {
			close(c.srcSocketChan)
			c.srcSocketConn.Close()
		}()

		scanner := bufio.NewScanner(c.srcSocketConn)

	loop:
		for {
			if !scanner.Scan() {
				err := scanner.Err()
				if err == nil {
					fmt.Printf("SRC Socket Closed TCP connection")

				} else {

					fmt.Printf("Error recieved on src TCP connection: %s", err)
				}

				c.srcSocketConn = nil
				break loop
			}

			c.srcSocketChan <- scanner.Bytes()
		}
	}(c)

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcSocketClosed() bool {
	return c.srcSocketConn == nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstSocketClosed() bool {
	return c.dstSocketConn == nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcMsgBufFull() bool {
	return len(c.srcMsgBuf) > 0
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstMsgBufFull() bool {
	return len(c.srcMsgBuf) > 0
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isConnectionEvent() bool {
	return c.connectionEvent.EventType != msgbus.NoEvent
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isMinerEvent() bool {
	return c.minerEvent.EventType != msgbus.NoEvent
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcSocketChanReady() bool {
	return len(c.srcSocketChan) > 0
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstSocketChanReady() bool {
	return len(c.dstSocketChan) > 0
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isEventConnectionChanReady() bool {
	return len(c.eventConnectionChan) > 0
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isEventMinerChanReady() bool {
	return len(c.eventMinerChan) > 0
}

//------------------------------------------
// Starting point for new src connections
//
// Create entry in connection table
// Create src connection handler
// 	Create reader channel
// 	Create for loop to feed the reader channel
//------------------------------------------
func (c *connection) handleStateNone() {

	fmt.Printf("Create new SRC connection")

	c.srcSocketChan = make(chan msgBuffer)
	c.connectionID = msgbus.ConnectionID(msgbus.GetRandomIDString())
	c.state = stateOpen
	c.eventConnectionChan = c.ps.NewEventChan()

	// This is called when there is a new src connection
	// Start the src socket reader
	c.runSrcSocketReader()

	var connStruct msgbus.Connection = msgbus.Connection{
		ID:        c.connectionID,
		Miner:     msgbus.MinerID(""),
		Dest:      msgbus.DestID(""),
		State:     msgbus.SrcOpenState,
		TotalHash: 0,
		StartDate: time.Now(),
	}

	// Create entry into connection table
	err1 := c.ps.Pub(
		msgbus.ConnectionMsg,
		msgbus.IDString(c.connectionID),
		connStruct)
	if err1 != nil {
		fmt.Printf("Error pubsub.Pub return error: %s", err1)
		panic(err1)
	}

	// Subscribe to events for the connection table
	err2 := c.ps.Sub(
		msgbus.ConnectionMsg,
		msgbus.IDString(c.connectionID),
		c.eventConnectionChan)
	if err2 != nil {
		fmt.Printf("Error pubsub.Sub return error: %s", err2)
		panic(err2)
	}

}

//------------------------------------------
//
// Is the SrcSocketConn closed?
// Is there a Connection Subscribed Event
//
//------------------------------------------
func (c *connection) handleStateOpen() {

	switch {

	case c.isSrcSocketClosed():
		c.state = stateSrcClosed
		return

	case c.isConnectionEvent():
		if c.connectionEvent.EventType == msgbus.SubscribedEvent {
			c.connectionEvent.EventType = msgbus.NoEvent
			c.state = stateAuth
		} else {
			panic("non-Subscribed Connection Event found in StateOpen")
		}
		return

	case c.isMinerEvent():
		panic("Miner Event found in StateOpen")

	case c.isSrcMsgBufFull():
		panic("src Msg Buffer found in StateOpen")

	case c.isDstMsgBufFull():
		panic("Dst Msg Buffer found in StateOpen")

	default:
		panic("Default reached in handleStateOpen")
	}

}

//------------------------------------------
//
// Figure out who we have
// Set the miner ID
//------------------------------------------
func (c *connection) handleStateAuth() {

	switch {
	case c.isSrcSocketClosed():
		c.state = stateSrcClosed
		return

	case c.isSrcMsgBufFull():

		c.state = stateHashVerify
		return

	case c.isConnectionEvent():
		panic("Connection Event found in StateAuth")

	case c.isMinerEvent():
		panic("Miner Event found in StateAuth")

	case c.isDstMsgBufFull():
		panic("Dst Msg Buffer found in StateAuth")

	default:
		panic("Default reached in handleStateAuth")
	}

}

//------------------------------------------
//
// Connect to spoof pool and monitor the messages going back and forth
//------------------------------------------
func (c *connection) handleStateHashVerify() {

	if c.isSrcSocketClosed() {
		c.state = stateSrcClosed
		return
	}

	c.state = stateRouting
}

//------------------------------------------
//
// Lookup Routing Data
// Should get a default pool address to start
// The connection scedular will redirect with a route change
//------------------------------------------
func (c *connection) handleStateRouting() {

	if c.isSrcSocketClosed() {
		c.state = stateSrcClosed
		return
	}

	c.state = stateConnecting
}

//------------------------------------------
//
// Connect to a the destination
//------------------------------------------
func (c *connection) handleStateConnecting() {
	if c.isSrcSocketClosed() {
		c.state = stateSrcClosed
		return
	}

	c.state = stateConnected
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateConnected() {

	if c.isSrcSocketClosed() {
		c.state = stateSrcClosed
		return
	}

	if c.isDstSocketClosed() {
		c.state = stateDstClosed
		return
	}

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateRouteChange() {
	if c.isSrcSocketClosed() {
		c.state = stateSrcClosed
		return
	}

	panic("")
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateMsgError() {
	if c.isSrcSocketClosed() {
		c.state = stateSrcClosed
		return
	}

	c.state = stateShutdown
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateConnectError() {
	if c.isSrcSocketClosed() {
		c.state = stateSrcClosed
		return
	}

	c.state = stateShutdown
}

//------------------------------------------
//
// Close Dst connection
// Update PubSub Values and wait for the return
//------------------------------------------
func (c *connection) handleStateSrcClosed() {

	c.state = stateShutdown
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateDstClosed() {
	if c.isSrcSocketClosed() {
		c.state = stateSrcClosed
		return
	}

	c.state = stateRouting
}

//------------------------------------------
//
// Close Dst
// Close Src
// Close ConnectionChan
// Close MinerChan
// Other Clean up?
//------------------------------------------
func (c *connection) handleStateShutdown() {

	c.state = stateClosed
}

//------------------------------------------
//
// Read from an array of channels
//	Src socket (c.srcChan)
//	Dst socket (c.dstChan)
//	PS Connection Event ()
// Read Current State and handle it
//------------------------------------------
func (c *connection) dispatchLoop() {

	for {
		switch c.state {
		case stateNone:
			c.handleStateNone()

		case stateOpen:
			c.handleStateOpen()

		case stateAuth:
			c.handleStateAuth()

		case stateHashVerify:
			c.handleStateHashVerify()

		case stateRouting:
			c.handleStateRouting()

		case stateConnecting:
			c.handleStateConnecting()

		case stateConnected:
			c.handleStateConnected()

		case stateRouteChange:
			c.handleStateRouteChange()

		case stateMgsError:
			c.handleStateMsgError()

		case stateConnectError:
			c.handleStateConnectError()

		case stateSrcClosed:
			c.handleStateSrcClosed()

		case stateDstClosed:
			c.handleStateDstClosed()

		case stateShutdown:
			c.handleStateShutdown()

		case stateClosed:
			return

		default:
			panic("Default Reached, dazed and confused")

		}

		// Skip the select process if any of these states are in effect
		if c.state == stateClosed || c.state == stateSrcClosed || c.state == stateShutdown {
			continue
		}

		// Wait on an event, then Read the event from
		// src/dst Socket Chan, Connection Event, Miner Event
		select {
		case c.srcMsgBuf = <-c.srcSocketChan:
		case c.dstMsgBuf = <-c.dstSocketChan:
		case c.connectionEvent = <-c.eventConnectionChan:
		case c.minerEvent = <-c.eventMinerChan:
		}

	}

}

//------------------------------------------
// Start the listener here
// Need to be able to take into account configuration variables
//
// Need context pointer to close out when the system is shutting down
//------------------------------------------
func (cm *ConnectionManager) start() error {
	l, err := net.Listen("tcp", "localhost"+":"+"3333")
	if err != nil {
		fmt.Printf("Listener Error %s", err)
		return err
	}

	go cm.listenForIncomingConnections(l)

	fmt.Printf("Connection Manager Started")

	return nil
}

//------------------------------------------
//
//------------------------------------------
func (cm *ConnectionManager) listenForIncomingConnections(l net.Listener) {
	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Printf("Error Accepting connection: %s", err)
			break
		}
		if conn == nil {
			fmt.Printf("Error no connection returned")
			break
		}

		c := newConnection(conn, cm.ps)
		go c.dispatchLoop()
	}

}
