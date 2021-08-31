package connectionmanager

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

type connectionStates string

const (
	stateNew          connectionStates = connectionStates(msgbus.ConnNewState)
	stateOpen         connectionStates = connectionStates(msgbus.ConnSrcOpenState)
	stateAuth         connectionStates = connectionStates(msgbus.ConnAuthState)
	stateHashVerify   connectionStates = connectionStates(msgbus.ConnVerifyState)
	stateRouting      connectionStates = connectionStates(msgbus.ConnRoutingState)
	stateConnecting   connectionStates = connectionStates(msgbus.ConnConnectingState)
	stateConnected    connectionStates = connectionStates(msgbus.ConnConnectedState)
	stateRouteChange  connectionStates = connectionStates(msgbus.ConnRouteChangeState)
	stateMgsError     connectionStates = connectionStates(msgbus.ConnMsgErrState)
	stateConnectError connectionStates = connectionStates(msgbus.ConnConnectErrState)
	stateSrcClosed    connectionStates = connectionStates(msgbus.ConnSrcCloseState)
	stateDstClosed    connectionStates = connectionStates(msgbus.ConnDstCloseState)
	stateError        connectionStates = connectionStates(msgbus.ConnErrorState)
	stateShutdown     connectionStates = connectionStates(msgbus.ConnShutdownState)
	stateClosed       connectionStates = connectionStates(msgbus.ConnClosedState)
)

var connectionStateMap = map[connectionStates]msgbus.ConnectionState{
	stateNew:          msgbus.ConnNewState,
	stateOpen:         msgbus.ConnSrcOpenState,
	stateAuth:         msgbus.ConnAuthState,
	stateHashVerify:   msgbus.ConnVerifyState,
	stateRouting:      msgbus.ConnRoutingState,
	stateConnecting:   msgbus.ConnConnectingState,
	stateConnected:    msgbus.ConnConnectedState,
	stateRouteChange:  msgbus.ConnRouteChangeState,
	stateMgsError:     msgbus.ConnMsgErrState,
	stateConnectError: msgbus.ConnConnectErrState,
	stateSrcClosed:    msgbus.ConnSrcCloseState,
	stateDstClosed:    msgbus.ConnDstCloseState,
	stateShutdown:     msgbus.ConnShutdownState,
	stateError:        msgbus.ConnErrorState,
	stateClosed:       msgbus.ConnClosedState,
}

type ConnectionManager struct {
	ps *msgbus.PubSub
}

type msgBuffer []byte

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
	connectionEvent     msgbus.Event
	minerEvent          msgbus.Event
	srcRequest          *request
	dstResponce         *responce
	srcStratumState     stratumState
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
	c.state = stateNew
	c.srcSocketConn = conn
	c.dstSocketConn = nil
	c.srcSocketChan = nil
	c.dstSocketChan = nil
	c.eventMinerChan = nil
	c.eventConnectionChan = nil
	c.minerID = ""
	c.connectionID = ""
	//	c.srcMsgBuf = nil
	//	c.dstMsgBuf = nil
	c.connectionEvent = ps.NewEvent()
	c.minerEvent = ps.NewEvent()
	c.srcStratumState = StratumNew

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
			if !c.isSrcSocketClosed() {
				c.srcSocketConn.Close()
			}
		}()

		scanner := bufio.NewScanner(c.srcSocketConn)

	loop:
		for {
			if !scanner.Scan() {
				err := scanner.Err()
				if err == nil {
					fmt.Printf("SRC Socket Closed TCP connection\n")

				} else {

					fmt.Printf("Error recieved on src TCP connection: %s\n", err)
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
	if c.srcSocketConn == nil {

		err := c.updateConnectionState(stateSrcClosed)
		if err != nil {
			fmt.Printf("updateConnectionState returned error: %s\n", err)
			c.state = stateError
		}
		return true
	}
	return false
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstSocketClosed() bool {
	if c.dstSocketConn == nil {

		err := c.updateConnectionState(stateDstClosed)
		if err != nil {
			fmt.Printf("updateConnectionState returned error: %s\n", err)
			c.state = stateError
		}
		return true
	}
	return false
}

//------------------------------------------
//
//------------------------------------------
//func (c *connection) isSrcMsgBufFull() bool {
//	return len(c.srcMsgBuf) > 0
//}

//------------------------------------------
//
//------------------------------------------
//func (c *connection) isDstMsgBufFull() bool {
//	return len(c.srcMsgBuf) > 0
//}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcRequestReady() bool {
	return c.srcRequest != nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstResponceReady() bool {
	return c.dstResponce != nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) clearSrcRequest() {
	c.srcRequest = nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) clearDstResponce() {
	c.dstResponce = nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isConnectionEventFull() bool {
	return c.connectionEvent.EventType != msgbus.NoEvent
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isMinerEventFull() bool {
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
//
//------------------------------------------
func (c *connection) clearConnectionEvent() {
	c.connectionEvent.EventType = msgbus.NoEvent
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) clearMinerEvent() {
	c.minerEvent.EventType = msgbus.NoEvent
}

//------------------------------------------
// Need to add timeout in here
//------------------------------------------
func (c *connection) waitConnectionEvent() error {

	if !c.isConnectionEventFull() {

		select {
		case c.connectionEvent = <-c.eventConnectionChan:
		}
	}

	return nil
}

//------------------------------------------
// Need to add timeout in here
//------------------------------------------
func (c *connection) waitMsgEvent() error {

	if c.isSrcRequestReady() ||
		c.isDstResponceReady() ||
		c.isConnectionEventFull() ||
		c.isMinerEventFull() {

		return nil
	}

	var srcMsgBuf msgBuffer = nil
	var dstMsgBuf msgBuffer = nil

	// Wait on an event, then Read the event from
	// src/dst Socket Chan, Connection Event, Miner Event
	select {
	case srcMsgBuf = <-c.srcSocketChan:
	case dstMsgBuf = <-c.dstSocketChan:
	case c.connectionEvent = <-c.eventConnectionChan:
	case c.minerEvent = <-c.eventMinerChan:
	}

	if srcMsgBuf != nil {
		request, err := getRequestMsg(srcMsgBuf)
		if err != nil {
			panic("Bad Src Json message")
		}
		c.srcRequest = request
	}

	if dstMsgBuf != nil {
		responce, err := getResponceMsg(dstMsgBuf)
		if err != nil {
			panic("Bad Dst Json message")
		}
		c.dstResponce = responce
	}

	return nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) updateConnectionState(s connectionStates) error {

	if c.state == s {
		fmt.Printf("update_connectionState() Current state is the desired state\n")
		return nil
	}

	ech := make(msgbus.EventChan)
	err := c.ps.Get(msgbus.ConnectionMsg, msgbus.IDString(c.connectionID), ech)
	if err != nil {
		fmt.Printf("Error returned for pubsub Get %s\n", err)
		return err
	}

	// Need a function that will do this but time out
	getEvent := <-ech

	connection := getEvent.Data.(msgbus.Connection)

	if connection.State != connectionStateMap[c.state] {
		//debug.PrintStack()
		fmt.Printf("connction.State: %v\n", connection.State)
		fmt.Printf("c.state: %v\n", c.state)
		//fmt.Printf("Error: %s", lumerinlib.Errtrace())
		// panic("states are out of sync")
	}

	connection.State = connectionStateMap[s]
	err = c.ps.Set(msgbus.ConnectionMsg, msgbus.IDString(c.connectionID), connection)
	if err != nil {
		fmt.Printf("Error returned for pubsub Set %s", err)
		return err
	}

	fmt.Printf("STATE changed to:%v for ID: %s", c.state, c.connectionID)
	c.state = s

	return nil
}

//------------------------------------------
// Starting point for new src connections
//
// Create src socket reader channel
// Create entry in connection table (Pub)
// Subscribe to updates for the connection table entry(Sub)
// Set State to Open
//
// Does not "read" any channels, passes that to the next state handler
//------------------------------------------
func (c *connection) handleStateNew() {

	if c.connectionID == "" {

		fmt.Printf("Create new SRC connection\n")

		c.srcSocketChan = make(chan msgBuffer)
		c.connectionID = msgbus.ConnectionID(msgbus.GetRandomIDString())
		c.state = stateNew
		c.eventConnectionChan = c.ps.NewEventChan()

		// This is called when there is a new src connection
		// Start the src socket reader
		c.runSrcSocketReader()

		var connStruct msgbus.Connection = msgbus.Connection{
			ID:        c.connectionID,
			Miner:     msgbus.MinerID(""),
			Dest:      msgbus.DestID(""),
			State:     msgbus.ConnNewState,
			TotalHash: 0,
			StartDate: time.Now(),
		}

		// Create entry into connection table
		err1 := c.ps.Pub(
			msgbus.ConnectionMsg,
			msgbus.IDString(c.connectionID),
			connStruct)
		if err1 != nil {
			fmt.Printf("Error pubsub.Pub return error: %s\n", err1)
			panic(err1)
		}

		// Subscribe to events for the connection table
		err2 := c.ps.Sub(
			msgbus.ConnectionMsg,
			msgbus.IDString(c.connectionID),
			c.eventConnectionChan)
		if err2 != nil {
			fmt.Printf("Error pubsub.Sub return error: %s\n", err2)
			panic(err2)
		}

	}

	for c.state == stateNew {

		// Wait on connectionEvent (skipping src msg events)
		err := c.waitConnectionEvent()
		if err != nil {
			c.state = stateError
			return
		}

		if c.connectionEvent.EventType == msgbus.SubscribedEvent {
			c.clearConnectionEvent()

			err := c.updateConnectionState(stateOpen)
			if err != nil {
				fmt.Printf("updateConnectionState returned error:" + err.Error())
				c.state = stateError
				return
			}

		} else {
			fmt.Printf("non-Subscribed Connection Event found in StateNew\n")
			c.state = stateError
			return
		}

	}
}

//------------------------------------------
//
// Is the SrcSocketConn closed?
// Is there a Connection Subscribed Event
//
//------------------------------------------
func (c *connection) handleStateOpen() {

	for c.state == stateOpen {

		switch {

		case c.isSrcSocketClosed():
			return

		case c.isConnectionEventFull():
			fmt.Printf("Connection Event found in StateOpen\n")
			c.state = stateError
			return

		case c.isSrcRequestReady():
			if c.srcRequest.Method == string(MINING_SUBSCRIBE) {
				c.srcStratumState = StratumSubscribed
			}

			c.clearSrcRequest()

			err := c.updateConnectionState(stateAuth)
			if err != nil {
				fmt.Printf("updateConnectionState returned error:" + err.Error() + "\n")
				c.state = stateError
			}
			return

		case c.isMinerEventFull():
			fmt.Printf("Miner Event found in StateOpen\n")
			c.state = stateError
			return

		case c.isDstResponceReady():
			fmt.Printf("Dst Msg Buffer found in StateOpen\n")
			c.state = stateError
			return

		default:
			fmt.Printf("Default reached in StateOpen\n")
		}

		err := c.waitMsgEvent()
		if err != nil {
			fmt.Printf("waitMsgEvent() returned error")
			c.state = stateError
			return
		}
	}

}

//------------------------------------------
//
// Figure out who we have
// Set the miner ID
//------------------------------------------
func (c *connection) handleStateAuth() {

	for c.state == stateAuth {
		switch {
		case c.isSrcSocketClosed():
			return

		case c.isSrcRequestReady():

			if c.srcRequest.Method == string(MINING_AUTHORIZE) {
				c.srcStratumState = StratumSubscribed
			}
			c.clearSrcRequest()

			err := c.updateConnectionState(stateHashVerify)
			if err != nil {
				fmt.Printf("updateConnectionState returned error:" + err.Error())
				c.state = stateError
			}
			return

		case c.isConnectionEventFull():
			fmt.Printf("Connection Event found in StateAuth\n")
			c.state = stateError
			return

		case c.isMinerEventFull():
			fmt.Printf("Miner Event found in StateAuth\n")
			c.state = stateError
			return

		case c.isDstResponceReady():
			fmt.Printf("Dst Msg Buffer found in StateAuth\n")
			c.state = stateError
			return

		default:
			fmt.Printf("Default reached in handleStateAuth\n")
		}

		err := c.waitMsgEvent()
		if err != nil {
			fmt.Printf("waitMsgEvent() returned error")
			c.state = stateError
			return
		}
	}

}

//------------------------------------------
//
// Connect to spoof pool and monitor the messages going back and forth
//------------------------------------------
func (c *connection) handleStateHashVerify() {

	for c.state == stateAuth {
		switch {
		case c.isSrcSocketClosed():
			return

		case c.isSrcRequestReady():
			fmt.Printf("SrcRequestReady in handleStateHashVerify()")
			c.state = stateError
			return

		case c.isConnectionEventFull():
			fmt.Printf("ConnectionEvent in handleStateHashVerify()")
			c.state = stateError
			return

		case c.isMinerEventFull():
			fmt.Printf("MinerEvent in handleStateHashVerify()")
			c.state = stateError
			return

		case c.isDstResponceReady():
			fmt.Printf("DstResponceReady in handleStateHashVerify()")
			c.state = stateError
			return

		default:
			fmt.Printf("Default reached in handleStateAuth")
		}

		err := c.updateConnectionState(stateRouting)
		if err != nil {
			fmt.Printf("updateConnectionState returned error:" + err.Error())
			c.state = stateError
		}
		return
	}
}

//------------------------------------------
//
// Lookup Routing Data
// Should get a default pool address to start
// The connection scedular will redirect with a route change
//------------------------------------------
func (c *connection) handleStateRouting() {

	for c.state == stateRouting {
		switch {
		case c.isSrcSocketClosed():
			return

		case c.isSrcRequestReady():
			fmt.Printf("SrcRequestReady in handleStateRouting()")
			c.state = stateError
			return

		case c.isConnectionEventFull():
			fmt.Printf("isConnectionEventFull in handleStateRouting()")
			c.state = stateError
			return

		case c.isMinerEventFull():
			fmt.Printf("isminerEventFull in handleStateRouting()")
			c.state = stateError
			return

		case c.isDstResponceReady():
			fmt.Printf("dstResponceReady in handleStateRouting()")
			c.state = stateError
			return

		default:
			fmt.Printf("default reached in handleStateRouting()")
		}

		err := c.updateConnectionState(stateConnecting)
		if err != nil {
			fmt.Printf("updateConnectionState returned error:" + err.Error())
			c.state = stateError
		}
		return
	}
}

//------------------------------------------
//
// Connect to a the destination
//------------------------------------------
func (c *connection) handleStateConnecting() {

	for c.state == stateConnecting {
		switch {
		case c.isSrcSocketClosed():
			return

		case c.isSrcRequestReady():
			fmt.Printf("SrcRequestReady in handleStateConnecting()")
			c.state = stateError
			return

		case c.isConnectionEventFull():
			fmt.Printf("isConnectionEventFull in handleStateConnecting()")
			c.state = stateError
			return

		case c.isMinerEventFull():
			fmt.Printf("isminerEventFull in handleStateConnecting()")
			c.state = stateError
			return

		case c.isDstResponceReady():
			fmt.Printf("dstResponceReady in handleStateConnecting()")
			c.state = stateError
			return

		default:
			fmt.Printf("default reached in handleStateConnecting()")
		}

		err := c.updateConnectionState(stateConnected)
		if err != nil {
			fmt.Printf("updateConnectionState returned error:" + err.Error())
			c.state = stateError
		}
		return

	}

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateConnected() {

	switch {
	case c.isSrcSocketClosed():
		return

	case c.isDstSocketClosed():
		c.state = stateDstClosed
		return

	case c.isSrcRequestReady():
	case c.isConnectionEventFull():
	case c.isMinerEventFull():
	case c.isDstResponceReady():
	default:
		panic("Default reached in handleStateAuth")
	}

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateRouteChange() {

	switch {
	case c.isSrcSocketClosed():
		return

	case c.isSrcRequestReady():
	case c.isConnectionEventFull():
	case c.isMinerEventFull():
	case c.isDstResponceReady():
	default:
		panic("Default reached in handleStateAuth")
	}

	panic("")
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateMsgError() {

	switch {
	case c.isSrcSocketClosed():
		return

	case c.isDstSocketClosed():
		c.state = stateDstClosed
		return

	case c.isSrcRequestReady():
	case c.isConnectionEventFull():
	case c.isMinerEventFull():
	case c.isDstResponceReady():
	default:
		panic("Default reached in handleStateAuth")
	}

	c.state = stateShutdown
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateConnectError() {

	switch {
	case c.isSrcSocketClosed():
		return

	case c.isDstSocketClosed():
		c.state = stateDstClosed
		return

	case c.isSrcRequestReady():
	case c.isConnectionEventFull():
	case c.isMinerEventFull():
	case c.isDstResponceReady():
	default:
		panic("Default reached in handleStateAuth")
	}

	c.state = stateShutdown
}

//------------------------------------------
//
// Close Dst connection
// Update PubSub Values and wait for the return
//------------------------------------------
func (c *connection) handleStateSrcClosed() {

	if c.isSrcSocketClosed() {
		return
	}

	err := c.updateConnectionState(stateShutdown)
	if err != nil {
		fmt.Printf("updateConnectionState returned error:" + err.Error())
	}

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateDstClosed() {

	if c.isSrcSocketClosed() {
		return
	}

	// Close out any resources here
	// Update PubSub

	err := c.updateConnectionState(stateRouting)
	if err != nil {
		fmt.Printf("updateConnectionState returned error:" + err.Error())
	}
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateError() {

	fmt.Printf("Error State Reached, shutting down...")

	err := c.updateConnectionState(stateShutdown)
	if err != nil {
		panic("updateConnectionState returned error:" + err.Error())
	}
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

	var err error
	connectionEventClosed := false
	minerEventClosed := false

	err = c.updateConnectionState(stateClosed)
	if err != nil {
		panic("updateConnectionState returned error:" + err.Error())
	}

	if c.eventConnectionChan != nil {
		err = c.ps.RemoveAndCloseEventChan(c.eventConnectionChan)
		if err != nil {
			fmt.Printf("Error pubsub.remove event return error: %s\n", err)
			panic(err)
		}
	} else {
		connectionEventClosed = true
	}

	if c.eventMinerChan != nil {
		err = c.ps.RemoveAndCloseEventChan(c.eventMinerChan)
		if err != nil {
			fmt.Printf("Error pubsub.Unsub return error: %s\n", err)
			panic(err)
		}
	} else {
		minerEventClosed = true
	}

	for !(connectionEventClosed && minerEventClosed) {

		err := c.waitMsgEvent()
		if err != nil {
			panic("In Shutdown and waitMsgEvent() returned error")
		}

		if !connectionEventClosed && c.isConnectionEventFull() {
			if c.connectionEvent.EventType == msgbus.RemovedEvent {
				connectionEventClosed = true
			}
			c.clearConnectionEvent()
		}

		if !minerEventClosed && c.isMinerEventFull() {
			if c.minerEvent.EventType == msgbus.RemovedEvent {
				minerEventClosed = true
			}
			c.clearMinerEvent()
		}

	}

	if c.srcSocketConn != nil {
		c.srcSocketConn.Close()
	}

	if c.dstSocketConn != nil {
		c.dstSocketConn.Close()
	}

	if c.eventConnectionChan != nil {
		close(c.eventConnectionChan)
	}

	if c.eventMinerChan != nil {
		close(c.eventMinerChan)
	}

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
		case stateNew:
			c.handleStateNew()

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

		case stateError:
			c.handleStateError()

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

		if false && c.isSrcRequestReady() ||
			c.isDstResponceReady() ||
			c.isConnectionEventFull() ||
			c.isMinerEventFull() {

			fmt.Printf("Found Event or new Message\n\tsrcRequest %v\n\tDstResp %v\n\tConnEvent %v\n\tMinerEvent %v\n",
				c.isSrcRequestReady(),
				c.isDstResponceReady(),
				c.isConnectionEventFull(),
				c.isMinerEventFull())

			if c.isSrcRequestReady() {
				fmt.Printf("Request: %s State: %v\n", c.srcRequest.Method, c.state)
			}
			continue
		}

		if false {
			var srcMsgBuf msgBuffer = nil
			var dstMsgBuf msgBuffer = nil

			// Wait on an event, then Read the event from
			// src/dst Socket Chan, Connection Event, Miner Event
			select {
			case srcMsgBuf = <-c.srcSocketChan:
			case dstMsgBuf = <-c.dstSocketChan:
			case c.connectionEvent = <-c.eventConnectionChan:
			case c.minerEvent = <-c.eventMinerChan:
			}

			if srcMsgBuf != nil {
				request, err := getRequestMsg(srcMsgBuf)
				if err != nil {
					panic("Bad Src Json message")
				}
				c.srcRequest = request
			}

			if dstMsgBuf != nil {
				responce, err := getResponceMsg(dstMsgBuf)
				if err != nil {
					panic("Bad Dst Json message")
				}
				c.dstResponce = responce
			}
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
		fmt.Printf("Listener Error %s\n", err)
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
			fmt.Printf("Error Accepting connection: %s\n", err)
			break
		}
		if conn == nil {
			fmt.Printf("Error no connection returned\n")
			break
		}

		c := newConnection(conn, cm.ps)
		go c.dispatchLoop()
	}

}
