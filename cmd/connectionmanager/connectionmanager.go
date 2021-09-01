package connectionmanager

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

const (
	defaultMinerHost  string = "localhost"
	defaultMinerPort  string = "3334"
	defaultMinerProto string = "tcp"
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

//
// Add these to consolodate network functions
//
type socketconn struct {
	name        string
	netConn     net.Conn
	bufReader   *bufio.Reader
	bufWriter   *bufio.Writer
	bufScanner  *bufio.Scanner
	ch          chan msgBuffer
	msgRequest  *request
	msgResponce *responce
	done        chan bool
}

type connection struct {
	ps              *msgbus.PubSub
	connectionState connectionStates
	srcConn         socketconn
	dstConn         socketconn
	connectionID    msgbus.ConnectionID
	minerID         msgbus.MinerID
	eventMinerChan  msgbus.EventChan
	minerEvent      msgbus.Event
	stratumState    stratumStates
	done            chan bool
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) setupSocket() (err error) {

	if s.netConn == nil {
		fmt.Printf(lumerinlib.FileLine() + "SocketConn is nil\n")
		panic("socketconn is nil\n")
	}

	if s.bufScanner == nil {
		s.bufScanner = bufio.NewScanner(s.netConn)
	}

	if s.bufReader == nil {
		s.bufReader = bufio.NewReader(s.netConn)
	}

	if s.bufWriter == nil {
		s.bufWriter = bufio.NewWriter(s.netConn)
	}

	s.ch = make(chan msgBuffer)

	s.msgRequest = nil
	s.msgResponce = nil

	return nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) close() {

	close(s.done)

	if s.netConn != nil {
		s.netConn.Close()
	}
	s.bufReader = nil
	s.bufWriter = nil
	s.bufScanner = nil
	//	s.ch          chan msgBuffer
	s.msgRequest = nil
	s.msgResponce = nil

}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isClosed() bool {
	select {
	case <-s.done:
		return true
	default:
	}
	return false
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isSocketClosed() bool {
	if s.netConn == nil {
		return true
	}
	return false
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isRequestReady() bool {
	return s.msgRequest != nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isResponceReady() bool {
	return s.msgResponce != nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) clearRequest() {
	s.msgRequest = nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) clearResponce() {
	s.msgResponce = nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) runSocketReader() {

	if s.bufReader == nil {
		s.setupSocket()
	}

	go func(s *socketconn) {

		defer func() {
			fmt.Printf("Closing %s Socket\n", s.name)
			s.netConn.Close()
		}()

	loop:
		for {
			fmt.Printf("%s Socket Scan()...\n", s.name)

			if !s.bufScanner.Scan() {
				err := s.bufScanner.Err()

				if err == nil {
					fmt.Printf(lumerinlib.FileLine()+"%s Socket Closed TCP connection\n", s.name)
				} else {
					fmt.Printf(lumerinlib.FileLine()+"Error recieved on %s TCP connection: %s\n", s.name, err)
				}

				break loop
			}

			buf := s.bufScanner.Bytes()

			if len(buf) > 0 {
				fmt.Printf("Read %s: %s\n", s.name, buf)
				s.ch <- buf
			} else {
				fmt.Printf(lumerinlib.FileLine()+"Warning: Read %s Zero Len, skipping\n", s.name)
			}
		}
	}(s)

}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) send() error {

	return nil

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
func newConnection(conn net.Conn, ps *msgbus.PubSub) (c *connection, err error) {
	c = &connection{}

	c.ps = ps
	c.connectionState = stateNew
	//	c.srcSocketConn = conn
	//	c.dstSocketConn = nil
	//	c.srcSocketChan = nil
	//	c.dstSocketChan = nil
	c.eventMinerChan = nil
	c.minerID = ""
	c.connectionID = ""
	c.minerEvent = ps.NewEvent()
	c.stratumState = StratumNew

	c.srcConn = socketconn{name: "SRC"}
	c.dstConn = socketconn{name: "DST"}

	if conn == nil {
		fmt.Printf(lumerinlib.FileLine() + "conn is nil\n")
		return nil, fmt.Errorf("conn (net.Conn) is nil")
	}

	c.srcConn.netConn = conn

	err = c.srcConn.setupSocket()
	if err != nil {
		fmt.Printf(lumerinlib.FileLine() + "setupSocket() returned error\n")
	}

	return c, err
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcSocketClosed() bool {
	if c.srcConn.isSocketClosed() {
		return true
	}
	return false
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstSocketClosed() bool {
	if c.dstConn.isSocketClosed() {
		return true
	}
	return false
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcRequestReady() bool {
	return c.srcConn.isRequestReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstResponceReady() bool {
	return c.dstConn.isResponceReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) clearSrcRequest() {
	c.srcConn.clearRequest()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) clearDstResponce() {
	c.dstConn.clearResponce()
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
	return len(c.srcConn.ch) > 0
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstSocketChanReady() bool {
	return len(c.dstConn.ch) > 0
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
func (c *connection) clearMinerEvent() {
	c.minerEvent.EventType = msgbus.NoEvent
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isMsgEventReady() bool {
	return c.isSrcRequestReady() || c.isDstResponceReady() || c.isMinerEventFull()
}

//------------------------------------------
// Need to add timeout in here
//------------------------------------------
func (c *connection) waitMsgEvent() error {

	if !c.isMsgEventReady() {

		fmt.Printf("waitMsgEvent()...\n")

		var srcMsgBuf msgBuffer = nil
		var dstMsgBuf msgBuffer = nil

		// Wait on an event, then Read the event from
		// src/dst Socket Chan, Connection Event, Miner Event
		select {
		case srcMsgBuf = <-c.srcConn.ch:
			if len(srcMsgBuf) > 0 {
				request, err := getRequestMsg(srcMsgBuf)
				if err != nil {
					panic("Bad Src Json message")
				}
				c.srcConn.msgRequest = request
			} else {
				fmt.Printf(lumerinlib.FileLine() + "Zero lenth SRC Message\n")
			}

		case dstMsgBuf = <-c.dstConn.ch:
			if len(dstMsgBuf) > 0 {
				responce, err := getResponceMsg(dstMsgBuf)
				if err != nil {
					panic("Bad Dst Json message")
				}
				c.dstConn.msgResponce = responce
			} else {
				fmt.Printf(lumerinlib.FileLine() + "Zero lenth DST Message\n")
			}

		case c.minerEvent = <-c.eventMinerChan:
		}

	}
	return nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) updateConnectionState(s connectionStates) error {

	if c.connectionState == s {
		fmt.Printf("update_connectionState() Current state is the desired state %v\n", s)
		return nil
	}

	event, err := c.ps.GetWait(msgbus.ConnectionMsg, msgbus.IDString(c.connectionID))
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error returned for pubsub Get %s\n", err)
		return err
	}

	// Need a function that will do this but time out
	// getEvent := <-ech

	connection := event.Data.(msgbus.Connection)

	if connection.State != connectionStateMap[c.connectionState] {
		fmt.Printf("connction.State: %v\n", connection.State)
		fmt.Printf("c.connectionState: %v\n", c.connectionState)
	}

	connection.State = connectionStateMap[s]
	_, err = c.ps.SetWait(msgbus.ConnectionMsg, msgbus.IDString(c.connectionID), connection)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error returned for pubsub Set %s\n", err)
		return err
	}

	fmt.Printf(lumerinlib.FileLine()+" STATE changed to:%v for ID: %s\n", c.connectionState, c.connectionID)
	c.connectionState = s

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

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	if c.connectionID == "" {

		fmt.Printf("Create new SRC connection\n")

		c.connectionID = msgbus.ConnectionID(msgbus.GetRandomIDString())
		c.connectionState = stateNew

		// This is called when there is a new src connection
		// Start the src socket reader
		c.srcConn.runSocketReader()

		var connStruct msgbus.Connection = msgbus.Connection{
			ID:        c.connectionID,
			Miner:     msgbus.MinerID(""),
			Dest:      msgbus.DestID(""),
			State:     msgbus.ConnNewState,
			TotalHash: 0,
			StartDate: time.Now(),
		}

		// Create entry into connection table
		_, err1 := c.ps.PubWait(
			msgbus.ConnectionMsg,
			msgbus.IDString(c.connectionID),
			connStruct)
		if err1 != nil {
			fmt.Printf("Error pubsub.Pub return error: %s\n", err1)
			panic(err1)
		}

	}

	err := c.updateConnectionState(stateOpen)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine() + "updateConnectionState returned error:" + err.Error() + "\n")
		c.connectionState = stateError
		return
	}

}

//------------------------------------------
//
// Is the SrcSocketConn closed?
// Is there a Connection Subscribed Event
//
//------------------------------------------
func (c *connection) handleStateOpen() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	for c.connectionState == stateOpen {

		if c.isMsgEventReady() {
			switch {

			case c.isSrcSocketClosed():
				err := c.updateConnectionState(stateSrcClosed)
				if err != nil {
					fmt.Printf(lumerinlib.FileLine() + "updateConnectionState returned error:" + err.Error() + "\n")
					c.connectionState = stateError
				}
				return

			case c.isSrcRequestReady():

				if c.srcConn.msgRequest.Method == string(MINING_SUBSCRIBE) {
					c.stratumState = StratumSubscribed
				}

				fmt.Printf(lumerinlib.FileLine()+"Read SRC Request: %v\n", c.srcConn.msgRequest)

				c.clearSrcRequest()

				err := c.updateConnectionState(stateAuth)
				if err != nil {
					fmt.Printf(lumerinlib.FileLine() + "updateConnectionState returned error:" + err.Error() + "\n")
					c.connectionState = stateError
				}
				return

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "Miner Event found in StateOpen\n")
				c.connectionState = stateError
				return

			case c.isDstResponceReady():
				fmt.Printf(lumerinlib.FileLine() + "Dst Msg Buffer found in StateOpen\n")
				c.connectionState = stateError
				return

			default:
				panic("Default reached in StateOpen\n")
			}
		}

		err := c.waitMsgEvent()
		if err != nil {
			fmt.Printf(lumerinlib.FileLine() + "waitMsgEvent() returned error")
			c.connectionState = stateError
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

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	for c.connectionState == stateAuth {
		if c.isMsgEventReady() {
			switch {
			case c.isSrcSocketClosed():
				err := c.updateConnectionState(stateSrcClosed)
				if err != nil {
					fmt.Printf(lumerinlib.FileLine() + "updateConnectionState returned error:" + err.Error() + "\n")
					c.connectionState = stateError
				}
				return

			case c.isSrcRequestReady():

				if c.srcConn.msgRequest.Method == string(MINING_AUTHORIZE) {
					c.stratumState = StratumSubscribed
				}
				c.clearSrcRequest()

				err := c.updateConnectionState(stateHashVerify)
				if err != nil {
					fmt.Printf(lumerinlib.FileLine() + "updateConnectionState returned error:" + err.Error())
					c.connectionState = stateError
				}
				return

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "Miner Event found in StateAuth\n")
				c.connectionState = stateError
				return

			case c.isDstResponceReady():
				fmt.Printf(lumerinlib.FileLine() + "Dst Msg Buffer found in StateAuth\n")
				c.connectionState = stateError
				return

			default:
				panic("Default reached in handleStateAuth\n")
			}
		}

		err := c.waitMsgEvent()
		if err != nil {
			fmt.Printf(lumerinlib.FileLine() + "waitMsgEvent() returned error")
			c.connectionState = stateError
			return
		}
	}

}

//------------------------------------------
//
// Connect to spoof pool and monitor the messages going back and forth
//------------------------------------------
func (c *connection) handleStateHashVerify() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	// Skip Verify for Now, go right to Routing
	err := c.updateConnectionState(stateRouting)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine() + "updateConnectionState returned error:" + err.Error())
		c.connectionState = stateError
	}

	for c.connectionState == stateAuth {
		if c.isMsgEventReady() {
			switch {
			case c.isSrcSocketClosed():
				return

			case c.isSrcRequestReady():
				fmt.Printf(lumerinlib.FileLine() + "SrcRequestReady in handleStateHashVerify()")
				c.connectionState = stateError
				return

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "MinerEvent in handleStateHashVerify()")
				c.connectionState = stateError
				return

			case c.isDstResponceReady():
				fmt.Printf(lumerinlib.FileLine() + "DstResponceReady in handleStateHashVerify()")
				c.connectionState = stateError
				return

			default:
				fmt.Printf("Default reached in handleStateAuth")
			}
		}

		err := c.updateConnectionState(stateRouting)
		if err != nil {
			fmt.Printf(lumerinlib.FileLine() + " updateConnectionState returned error:" + err.Error())
			c.connectionState = stateError
			return
		}
	}
}

//------------------------------------------
//
// Lookup Routing Data
// Should get a default pool address to start
// The connection scedular will redirect with a route change
//------------------------------------------
func (c *connection) handleStateRouting() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	// Skip Routing for Now
	err := c.updateConnectionState(stateShutdown)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine() + "updateConnectionState returned error:" + err.Error())
		c.connectionState = stateError
	}

	for c.connectionState == stateRouting {
		if c.isMsgEventReady() {
			switch {
			case c.isSrcSocketClosed():
				return

			case c.isSrcRequestReady():
				fmt.Printf(lumerinlib.FileLine() + "SrcRequestReady in handleStateRouting()")
				c.connectionState = stateError
				return

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "isminerEventFull in handleStateRouting()")
				c.connectionState = stateError
				return

			case c.isDstResponceReady():
				fmt.Printf(lumerinlib.FileLine() + "dstResponceReady in handleStateRouting()")
				c.connectionState = stateError
				return

			default:
				fmt.Printf("default reached in handleStateRouting()")
			}
		}

		err := c.updateConnectionState(stateConnecting)
		if err != nil {
			fmt.Printf(lumerinlib.FileLine() + "updateConnectionState returned error:" + err.Error())
			c.connectionState = stateError
			return
		}
	}
}

//------------------------------------------
//
// Connect to a the destination
//------------------------------------------
func (c *connection) handleStateConnecting() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	for c.connectionState == stateConnecting {
		if c.isMsgEventReady() {
			switch {
			case c.isSrcSocketClosed():
				return

			case c.isSrcRequestReady():
				fmt.Printf(lumerinlib.FileLine() + "SrcRequestReady in handleStateConnecting()")
				c.connectionState = stateError
				return

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "isminerEventFull in handleStateConnecting()")
				c.connectionState = stateError
				return

			case c.isDstResponceReady():
				fmt.Printf(lumerinlib.FileLine() + "dstResponceReady in handleStateConnecting()")
				c.connectionState = stateError
				return

			default:
				fmt.Printf("default reached in handleStateConnecting()")
			}
		}

		err := c.updateConnectionState(stateConnected)
		if err != nil {
			fmt.Printf(lumerinlib.FileLine() + "updateConnectionState returned error:" + err.Error())
			c.connectionState = stateError
			return
		}

	}

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateConnected() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	switch {
	case c.isSrcSocketClosed():
		return

	case c.isDstSocketClosed():
		c.connectionState = stateDstClosed
		return

	case c.isSrcRequestReady():
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

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	switch {
	case c.isSrcSocketClosed():
		return

	case c.isSrcRequestReady():
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

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	switch {
	case c.isSrcSocketClosed():
		return

	case c.isDstSocketClosed():
		c.connectionState = stateDstClosed
		return

	case c.isSrcRequestReady():
		//	case c.isConnectionEventFull():
	case c.isMinerEventFull():
	case c.isDstResponceReady():
	default:
		panic("Default reached in handleStateAuth")
	}

	c.connectionState = stateShutdown
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateConnectError() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	switch {
	case c.isSrcSocketClosed():
		return

	case c.isDstSocketClosed():
		c.connectionState = stateDstClosed
		return

	case c.isSrcRequestReady():
		//	case c.isConnectionEventFull():
	case c.isMinerEventFull():
	case c.isDstResponceReady():
	default:
		panic("Default reached in handleStateAuth")
	}

	c.connectionState = stateShutdown
}

//------------------------------------------
//
// Close Dst connection
// Update PubSub Values and wait for the return
//------------------------------------------
func (c *connection) handleStateSrcClosed() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	if !c.isSrcSocketClosed() {
		panic(lumerinlib.FileLine() + "Src is not closed\n")
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

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

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

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	fmt.Printf("Error State Reached, shutting down...\n")

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

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	var err error
	//	connectionEventClosed := false
	minerEventClosed := false

	if c.eventMinerChan != nil {
		err = c.ps.RemoveAndCloseEventChan(c.eventMinerChan)
		if err != nil {
			fmt.Printf("Error pubsub.Unsub return error: %s\n", err)
			panic(err)
		}
	} else {
		minerEventClosed = true
	}

	if !minerEventClosed && c.isMinerEventFull() {
		if c.minerEvent.EventType == msgbus.RemovedEvent {
			minerEventClosed = true
		}
		c.clearMinerEvent()
	}

	c.srcConn.close()
	c.dstConn.close()

	if c.eventMinerChan != nil {
		close(c.eventMinerChan)
	}

	err = c.updateConnectionState(stateClosed)
	if err != nil {
		panic("updateConnectionState returned error:" + err.Error())
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

		switch c.connectionState {
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
		if c.connectionState == stateClosed || c.connectionState == stateSrcClosed || c.connectionState == stateShutdown {
			continue
		}

		c.waitMsgEvent()

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

		c, err := newConnection(conn, cm.ps)
		if err != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+" newConnection() failed with %s\n", err))
		}
		go c.dispatchLoop()
	}

}
