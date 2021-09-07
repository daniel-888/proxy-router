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
type socketconnType string

const (
	SRC socketconnType = "SRC"
	DST socketconnType = "DST"
)

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
	name        socketconnType
	netConn     net.Conn
	bufReader   *bufio.Reader
	bufWriter   *bufio.Writer
	bufScanner  *bufio.Scanner
	ch          chan msgBuffer
	msgRequest  *request
	msgResponce *responce
	msgNotice   *notice
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
		panic(fmt.Sprintf(lumerinlib.FileLine()+"%s: netConn is nil\n", s.name))
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
	s.done = make(chan bool)
	s.msgRequest = nil
	s.msgResponce = nil

	return nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) close() {

	if !s.isClosed() {
		close(s.done)
	}

	if s.netConn != nil {
		s.netConn.Close()
	}
	s.bufReader = nil
	s.bufWriter = nil
	s.bufScanner = nil
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
	return s.netConn == nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isMsgReady() bool {
	return s.msgRequest != nil || s.msgResponce != nil || s.msgNotice != nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isRequestMsgReady() bool {
	return s.msgRequest != nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isResponceMsgReady() bool {
	return s.msgResponce != nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) isNoticeMsgReady() bool {
	return s.msgNotice != nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) clearMsg() {
	s.msgRequest = nil
	s.msgResponce = nil
	s.msgNotice = nil
}

//---------------------------------------------
//
//---------------------------------------------
func (s *socketconn) runSocketReader() {

	fmt.Printf("Running %s SocketReader\n", s.name)

	if s.netConn == nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"%s: netConn == nil\n", s.name))
	}

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
				fmt.Printf(lumerinlib.FileLine()+"Read %s: %s\n", s.name, buf)
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
func (s *socketconn) dial(proto string, host string, port string) error {

	c, err := net.Dial(proto, host+":"+port)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"net.Dial() failed:%s\n", err)
		return err
	}

	s.netConn = c

	err = s.setupSocket()
	if err == nil {
		s.runSocketReader()
	}

	return err
}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) send(b []byte) error {

	// Add a newline
	fmt.Printf("Socket send(): %s\n", b)

	// Stratum Protocol uses a "\n" as delimiter, it will not process until it sees this
	// The JSON package does not add one, so it is added here.
	b = append(b, "\n"...)

	msgLen := len(b)

	len, err := s.bufWriter.Write(b)
	if err != nil {
		return err
	} else if msgLen != len {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"msgLen:%d not eq len:%d\n", msgLen, len))
	}

	return s.bufWriter.Flush()

}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) sendRequest(request *request) error {

	r, err := createRequestMsg(request)
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"json.Marshal errored:%s\n", err))
	}

	return s.send(r)

}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) sendResponce(responce *responce) error {

	r, err := createResponceMsg(responce)
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"json.Marshal errored:%s\n", err))
	}

	return s.send(r)
}

//------------------------------------------
//
//------------------------------------------
func (s *socketconn) sendNotice(notice *notice) error {

	n, err := createNoticeMsg(notice)
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"json.Marshal errored:%s\n", err))
	}

	return s.send(n)
}

//--------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------

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
	c.eventMinerChan = nil
	c.minerID = ""
	c.connectionID = ""
	c.minerEvent = ps.NewEvent()
	c.stratumState = StratumNew

	c.srcConn = socketconn{
		name:        SRC,
		ch:          make(chan msgBuffer),
		done:        make(chan bool),
		msgRequest:  nil,
		msgResponce: nil,
	}

	c.dstConn = socketconn{
		name:        DST,
		ch:          make(chan msgBuffer),
		done:        make(chan bool),
		msgRequest:  nil,
		msgResponce: nil,
	}

	if conn == nil {
		fmt.Printf(lumerinlib.FileLine() + "conn is nil\n")
		return nil, fmt.Errorf("conn (net.Conn) is nil")
	}

	c.srcConn.netConn = conn

	err = c.srcConn.setupSocket()
	if err != nil {
		fmt.Printf(lumerinlib.FileLine() + "setupSocket() returned error\n")
		// This thing is basically stillborn
		c.connectionState = stateShutdown
	} else {
		err = c.dstConn.dial(defaultMinerProto, defaultMinerHost, defaultMinerPort)
		if err != nil {
			fmt.Printf(lumerinlib.FileLine() + "setupSocket() returned error\n")
			// This thing is basically stillborn
			c.connectionState = stateShutdown
		}
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
func (c *connection) isSrcMsgReady() bool {
	return c.srcConn.isMsgReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstMsgReady() bool {
	return c.dstConn.isMsgReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcRequestReady() bool {
	return c.srcConn.isRequestMsgReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstRequestReady() bool {
	return c.dstConn.isRequestMsgReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcResponceReady() bool {
	return c.srcConn.isResponceMsgReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstResponceReady() bool {
	return c.dstConn.isResponceMsgReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcNoticeReady() bool {
	return c.srcConn.isNoticeMsgReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstNoticeReady() bool {
	return c.dstConn.isNoticeMsgReady()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) clearSrcMsg() {
	c.srcConn.clearMsg()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) clearDstMsg() {
	c.dstConn.clearMsg()
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
	return c.isSrcMsgReady() || c.isDstMsgReady() || c.isMinerEventFull()
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

		ok := false
		select {
		case srcMsgBuf = <-c.srcConn.ch:
			if len(srcMsgBuf) > 0 {
				msg, err := getStratumMsg(srcMsgBuf)
				if err != nil {
					panic("Bad Src Json message")
				}

				fmt.Printf(lumerinlib.FileLine()+"MSG: type:%T\n--\tdata:%v", msg, msg)
				c.srcConn.msgNotice = nil
				c.srcConn.msgResponce = nil
				c.srcConn.msgRequest = nil
				switch msg.(type) {
				case *notice:
					c.srcConn.msgNotice = msg.(*notice)
				case *responce:
					c.srcConn.msgResponce = msg.(*responce)
				case *request:
					c.srcConn.msgRequest = msg.(*request)
				default:
					panic("")
				}
			} else {
				fmt.Printf(lumerinlib.FileLine() + "Zero lenth SRC Message\n")
			}
			ok = true

		case dstMsgBuf = <-c.dstConn.ch:
			if len(dstMsgBuf) > 0 {
				msg, err := getStratumMsg(dstMsgBuf)
				if err != nil {
					panic("Bad Dst Json message")
				}
				c.dstConn.msgNotice = nil
				c.dstConn.msgResponce = nil
				c.dstConn.msgRequest = nil
				switch msg.(type) {
				case *notice:
					c.dstConn.msgNotice = msg.(*notice)
				case *responce:
					c.dstConn.msgResponce = msg.(*responce)
				case *request:
					c.dstConn.msgRequest = msg.(*request)
				default:
					panic("")
				}
			} else {
				fmt.Printf(lumerinlib.FileLine() + "Zero lenth DST Message\n")
			}
			ok = true

		case c.minerEvent = <-c.eventMinerChan:
			ok = true
		}

		if !ok {
			panic("should not be here")
		}

	}
	return nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) updateConnectionState(s connectionStates) {

	if c.connectionState == s {
		fmt.Printf("update_connectionState() Current state is the desired state %v\n", s)
		return
	}

	// If connectionID is not defined yet, there is no msgbus entry to update
	if c.connectionID == "" {
		c.connectionState = s
		return
	}

	event, err := c.ps.GetWait(msgbus.ConnectionMsg, msgbus.IDString(c.connectionID))
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+"Error returned for pubsub Get %s\n", err))
	}

	// Need a function that will do this but time out
	// getEvent := <-ech

	// fmt.Printf("Event Data Type: %T", event.Data)

	switch t := event.Data.(type) {
	case msgbus.IDIndex:
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Data Type %s\n", t))

	case msgbus.Connection:
		connection := event.Data.(msgbus.Connection)

		if connection.State != connectionStateMap[c.connectionState] {
			fmt.Printf("connction.State: %v\n", connection.State)
			fmt.Printf("c.connectionState: %v\n", c.connectionState)
		}

		connection.State = connectionStateMap[s]
		_, err = c.ps.SetWait(msgbus.ConnectionMsg, msgbus.IDString(c.connectionID), connection)
		if err != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+"Error returned for pubsub Set %s\n", err))
		}

		fmt.Printf(lumerinlib.FileLine()+" STATE changed to:%v for ID: %s\n", c.connectionState, c.connectionID)
		c.connectionState = s

	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Unknown Type %T\n", event.Data))
	}
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleSrcRequest() {

	method := c.srcConn.msgRequest.Method
	switch method {
	case string(MINING_SUBSCRIBE):
	case string(MINING_AUTHORIZE):
	case string(MINING_CONFIGURE):
	case string(MINING_SET_TARGET):
	case string(MINING_SUBMIT):
	case string(MINING_NOTIFY):
	default:
		panic(fmt.Sprintf("Method not handled: %s", method))
	}

	c.dstConn.sendRequest(c.srcConn.msgRequest)
	c.clearSrcMsg()

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleSrcResponce() {
	c.dstConn.sendResponce(c.srcConn.msgResponce)
	c.clearSrcMsg()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleSrcNotice() {
	c.dstConn.sendNotice(c.srcConn.msgNotice)
	c.clearSrcMsg()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleDstRequest() {

	method := c.dstConn.msgRequest.Method
	switch method {
	case string(MINING_AUTHORIZE):
	case string(MINING_CONFIGURE):
	case string(MINING_NOTIFY):
	case string(MINING_SET_DIFFICULTY):
	case string(MINING_SET_TARGET):
	case string(MINING_SUBMIT):
	case string(MINING_SUBSCRIBE):
	default:
		panic(fmt.Sprintf("Method not handled: %s", method))
	}

	c.srcConn.sendRequest(c.dstConn.msgRequest)
	c.clearDstMsg()

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleDstResponce() {
	c.srcConn.sendResponce(c.dstConn.msgResponce)
	c.clearDstMsg()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleDstNotice() {
	c.srcConn.sendNotice(c.dstConn.msgNotice)
	c.clearDstMsg()
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

	c.updateConnectionState(stateOpen)

}

//------------------------------------------
//
// Is the SrcSocketConn closed?
// Is there a Connection Subscribed Event
//
// {"id": 1, "method": "mining.subscribe", "params": ["cpuminer/2.5.1"]}
// {"result":[[["mining.notify","61320eac"]],"ac0e3261",8],"id":1,"error":null}
//
//------------------------------------------
func (c *connection) handleStateOpen() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	for c.connectionState == stateOpen {

		if c.isMsgEventReady() {
			switch {

			case c.isSrcSocketClosed():
				c.updateConnectionState(stateSrcClosed)
				return

			case c.isDstSocketClosed():
				c.updateConnectionState(stateError)
				return

			case c.isSrcRequestReady():
				c.handleSrcRequest()

			case c.isSrcResponceReady():
				c.handleSrcResponce()

			case c.isSrcNoticeReady():
				c.handleSrcNotice()

			case c.isDstRequestReady():
				c.handleDstRequest()

			case c.isDstResponceReady():
				c.handleDstResponce()

			case c.isDstNoticeReady():
				c.handleDstNotice()

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "Miner Event found in StateOpen\n")
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
//
//
// {"id": 2, "method": "mining.authorize", "params": ["testrig", ""]}
// {"params":[32],"id":null,"method":"mining.set_difficulty"}
//
//------------------------------------------
func (c *connection) handleStateAuth() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	for c.connectionState == stateAuth {
		if c.isMsgEventReady() {
			switch {
			case c.isSrcSocketClosed():
				c.updateConnectionState(stateSrcClosed)
				return

			case c.isDstSocketClosed():
				c.updateConnectionState(stateError)
				return

			case c.isSrcRequestReady():
				c.handleSrcRequest()

			case c.isSrcResponceReady():
				c.handleSrcResponce()

			case c.isSrcNoticeReady():
				c.handleSrcNotice()

			case c.isDstRequestReady():
				c.handleDstRequest()

			case c.isDstResponceReady():
				c.handleDstResponce()

			case c.isDstNoticeReady():
				c.handleDstNotice()

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "Miner Event found in StateAuth\n")
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

	// Skip Verify for Now, go right to Connected
	c.updateConnectionState(stateConnected)

	for c.connectionState == stateAuth {

		if c.isMsgEventReady() {
			switch {

			case c.isSrcSocketClosed():
				c.updateConnectionState(stateSrcClosed)
				return

			case c.isDstSocketClosed():
				c.updateConnectionState(stateError)
				return

			case c.isSrcRequestReady():
				c.handleSrcRequest()

			case c.isSrcResponceReady():
				c.handleSrcResponce()

			case c.isSrcNoticeReady():
				c.handleSrcNotice()

			case c.isDstRequestReady():
				c.handleDstRequest()

			case c.isDstResponceReady():
				c.handleDstResponce()

			case c.isDstNoticeReady():
				c.handleDstNotice()

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "MinerEvent in handleStateHashVerify()")
				c.connectionState = stateError
				return

			default:
				fmt.Printf("Default reached in handleStateAuth")
			}
		}

		c.updateConnectionState(stateRouting)
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
	c.updateConnectionState(stateShutdown)

	for c.connectionState == stateRouting {
		if c.isMsgEventReady() {
			switch {

			case c.isSrcSocketClosed():
				c.updateConnectionState(stateSrcClosed)
				return

			case c.isDstSocketClosed():
				c.updateConnectionState(stateError)
				return

			case c.isSrcRequestReady():
				c.handleSrcRequest()

			case c.isSrcResponceReady():
				c.handleSrcResponce()

			case c.isSrcNoticeReady():
				c.handleSrcNotice()

			case c.isDstRequestReady():
				c.handleDstRequest()

			case c.isDstResponceReady():
				c.handleDstResponce()

			case c.isDstNoticeReady():
				c.handleDstNotice()

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "isminerEventFull in handleStateRouting()")
				c.connectionState = stateError
				return

			default:
				fmt.Printf("default reached in handleStateRouting()")
			}
		}

		c.updateConnectionState(stateConnecting)
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
				c.updateConnectionState(stateSrcClosed)
				return

			case c.isDstSocketClosed():
				c.updateConnectionState(stateError)
				return

			case c.isSrcRequestReady():
				c.handleSrcRequest()

			case c.isSrcResponceReady():
				c.handleSrcResponce()

			case c.isSrcNoticeReady():
				c.handleSrcNotice()

			case c.isDstRequestReady():
				c.handleDstRequest()

			case c.isDstResponceReady():
				c.handleDstResponce()

			case c.isDstNoticeReady():
				c.handleDstNotice()

			case c.isMinerEventFull():
				fmt.Printf(lumerinlib.FileLine() + "isminerEventFull in handleStateConnecting()")
				c.connectionState = stateError
				return

			default:
				fmt.Printf("default reached in handleStateConnecting()")
			}
		}

		c.updateConnectionState(stateConnected)

	}

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateConnected() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	for c.connectionState == stateConnected {
		if c.isMsgEventReady() {
			switch {
			case c.isSrcSocketClosed():
				c.updateConnectionState(stateSrcClosed)
				return

			case c.isDstSocketClosed():
				c.updateConnectionState(stateDstClosed)
				return

			case c.isSrcRequestReady():
				c.handleSrcRequest()

			case c.isSrcResponceReady():
				c.handleSrcResponce()

			case c.isSrcNoticeReady():
				c.handleSrcNotice()

			case c.isDstRequestReady():
				c.handleDstRequest()

			case c.isDstResponceReady():
				c.handleDstResponce()

			case c.isDstNoticeReady():
				c.handleDstNotice()

			case c.isMinerEventFull():
			default:
				panic("Default reached in handleStateAuth")
			}
		}
	}

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateRouteChange() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	switch {
	case c.isSrcSocketClosed():
		c.updateConnectionState(stateSrcClosed)
		return

	case c.isDstSocketClosed():
		c.updateConnectionState(stateDstClosed)
		return

	case c.isSrcRequestReady():
		c.handleSrcRequest()

	case c.isSrcResponceReady():
		c.handleSrcResponce()

	case c.isSrcNoticeReady():
		c.handleSrcNotice()

	case c.isDstRequestReady():
		c.handleDstRequest()

	case c.isDstResponceReady():
		c.handleDstResponce()

	case c.isDstNoticeReady():
		c.handleDstNotice()

	case c.isMinerEventFull():

	default:
		panic("Default reached in handleStateAuth")
	}

	panic("not handled yet")
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateMsgError() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	switch {
	case c.isSrcSocketClosed():
	case c.isDstSocketClosed():

	case c.isMinerEventFull():
	case c.isSrcRequestReady():
		c.handleSrcRequest()

	case c.isSrcResponceReady():
		c.handleSrcResponce()

	case c.isSrcNoticeReady():
		c.handleSrcNotice()

	case c.isDstRequestReady():
		c.handleDstRequest()

	case c.isDstResponceReady():
		c.handleDstResponce()

	case c.isDstNoticeReady():
		c.handleDstNotice()

	default:
		panic("Default reached in handleStateAuth")
	}

	c.connectionState = stateShutdown
	panic("not handled yet")
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateConnectError() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	switch {
	case c.isSrcSocketClosed():
	case c.isDstSocketClosed():
		c.connectionState = stateDstClosed
	case c.isMinerEventFull():
	case c.isSrcRequestReady():
		c.handleSrcRequest()

	case c.isSrcResponceReady():
		c.handleSrcResponce()

	case c.isSrcNoticeReady():
		c.handleSrcNotice()

	case c.isDstRequestReady():
		panic("")
	case c.isDstResponceReady():
		panic("")
	case c.isDstNoticeReady():
		panic("")

	default:
		panic("Default reached in handleStateAuth")
	}

	c.connectionState = stateShutdown
	panic("not handled yet")
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

	c.updateConnectionState(stateShutdown)
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateDstClosed() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	//if c.isSrcSocketClosed() {
	//	return
	//}

	// Close out any resources here
	// Update PubSub

	c.updateConnectionState(stateRouting)
	panic("not handled yet")
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateError() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	fmt.Printf("Error State Reached, shutting down...\n")

	c.updateConnectionState(stateShutdown)
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

	c.updateConnectionState(stateClosed)

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
			fmt.Printf(lumerinlib.FileLine()+" newConnection() failed with %s\n", err)
		}
		go c.dispatchLoop()
	}

}
