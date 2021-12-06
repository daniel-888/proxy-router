package connectionmanager

import (
	"fmt"
	"net"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/config"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

//const (
//	defaultMinerHost  msgbus.DestNetHost  = "localhost"
//	defaultMinerPort  msgbus.DestNetPort  = "3334"
//	defaultMinerProto msgbus.DestNetProto = "tcp"
//)

type connectionStates string
type socketconnType string

const (
	SRC socketconnType = "SRC"
	DST socketconnType = "DST"
)

const (
	stateNew            connectionStates = connectionStates(msgbus.ConnNewState)
	stateRouting        connectionStates = connectionStates(msgbus.ConnRoutingState)
	stateConnecting     connectionStates = connectionStates(msgbus.ConnConnectingState)
	stateSubscribeStep0 connectionStates = "stateSubscribeStep0"
	stateSubscribeStep1 connectionStates = "stateSubscribeStep1"
	stateAuthStep0      connectionStates = "stateAuthStep0"
	stateAuthStep1      connectionStates = "stateAuthStep1"
	stateHashVerify     connectionStates = connectionStates(msgbus.ConnVerifyState)
	stateConnected      connectionStates = connectionStates(msgbus.ConnConnectedState)
	stateRouteChange    connectionStates = connectionStates(msgbus.ConnRouteChangeState)
	stateMgsError       connectionStates = connectionStates(msgbus.ConnMsgErrState)
	stateConnectError   connectionStates = connectionStates(msgbus.ConnConnectErrState)
	stateSrcClosed      connectionStates = connectionStates(msgbus.ConnSrcCloseState)
	stateDstClosed      connectionStates = connectionStates(msgbus.ConnDstCloseState)
	stateError          connectionStates = connectionStates(msgbus.ConnErrorState)
	stateShutdown       connectionStates = connectionStates(msgbus.ConnShutdownState)
	stateClosed         connectionStates = connectionStates(msgbus.ConnClosedState)
)

var connectionStateMap = map[connectionStates]msgbus.ConnectionState{
	stateNew:            msgbus.ConnNewState,
	stateRouting:        msgbus.ConnRoutingState,
	stateConnecting:     msgbus.ConnConnectingState,
	stateSubscribeStep0: msgbus.ConnSrcSubscribeState,
	stateSubscribeStep1: msgbus.ConnSrcSubscribeState,
	stateAuthStep0:      msgbus.ConnAuthState,
	stateAuthStep1:      msgbus.ConnAuthState,
	stateHashVerify:     msgbus.ConnVerifyState,
	stateConnected:      msgbus.ConnConnectedState,
	stateRouteChange:    msgbus.ConnRouteChangeState,
	stateMgsError:       msgbus.ConnMsgErrState,
	stateConnectError:   msgbus.ConnConnectErrState,
	stateSrcClosed:      msgbus.ConnSrcCloseState,
	stateDstClosed:      msgbus.ConnDstCloseState,
	stateShutdown:       msgbus.ConnShutdownState,
	stateError:          msgbus.ConnErrorState,
	stateClosed:         msgbus.ConnClosedState,
}

type ConnectionManager struct {
	ps *msgbus.PubSub
}

type msgBuffer []byte

//
// Add these to consolodate network functions
//

type connection struct {
	ps              *msgbus.PubSub
	connectionState connectionStates
	srcConn         socketconn
	dstConn         socketconn
	connectionID    msgbus.ConnectionID
	minerID         msgbus.MinerID
	destID          msgbus.DestID
	eventMinerChan  msgbus.EventChan
	minerEvent      msgbus.Event
	stratumState    stratumStates
	srcMsgBuf       msgBuffer
	dstMsgBuf       msgBuffer
	requestid       int
	done            chan bool
}

//--------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------

//------------------------------------------
//
//------------------------------------------
// func New(ps *msgbus.PubSub, config map[string]interface{}) (cm *ConnectionManager, err error) {
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
	c.destID = ""
	c.connectionID = ""
	c.minerEvent = ps.NewEvent()
	c.stratumState = StratumNew
	c.srcMsgBuf = nil
	c.dstMsgBuf = nil
	c.requestid = 0

	c.srcConn = newSocketConn(SRC)
	c.dstConn = newSocketConn(DST)

	c.srcConn.netConn = conn

	return c, err
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isSrcSocketClosed() bool {
	return c.srcConn.isSocketClosed()
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) isDstSocketClosed() bool {
	return c.dstConn.isSocketClosed()
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
func (c *connection) isSrcOrDstMsgReady() bool {
	return c.srcConn.isMsgReady() || c.dstConn.isMsgReady()
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
func (c *connection) getSrcRequestMsg() *request {
	return c.srcConn.getRequestMsg()
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
func (c *connection) getDstRequestMsg() *request {
	return c.dstConn.getRequestMsg()
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
func (c *connection) getSrcResponceMsg() *responce {
	return c.srcConn.getResponceMsg()
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
func (c *connection) getDstResponceMsg() *responce {
	return c.dstConn.getResponceMsg()
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
func (c *connection) getSrcNoticeMsg() *notice {
	return c.srcConn.getNoticeMsg()
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
func (c *connection) getDstNoticeMsg() *notice {
	return c.dstConn.getNoticeMsg()
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

//---------------------------------------------
//
//---------------------------------------------
func (c *connection) getDefaultDest() msgbus.Dest {

	dest, err := c.ps.DestGetWait(msgbus.DEFAULT_DEST_ID)

	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Default Destination not in message bus %s\n", err))
	}

	return *dest
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) readSrcMsg() error {

	if len(c.srcMsgBuf) > 0 {
		msg, err := getStratumMsg(c.srcMsgBuf)
		if err != nil {
			panic("Bad Src Json message")
		}

		fmt.Printf(lumerinlib.FileLine()+"SRC MSG: type:%T\n--\tdata:%v\n", msg, msg)
		// c.srcConn.msgNotice = nil
		// c.srcConn.msgResponce = nil
		// c.srcConn.msgRequest = nil
		switch msg.(type) {
		case *notice:
			c.srcConn.addNoticeMsg(msg.(*notice))
		case *responce:
			c.srcConn.addResponceMsg(msg.(*responce))
		case *request:
			c.srcConn.addRequestMsg(msg.(*request))
		default:
			panic(fmt.Sprintf("what the hell is this?  :%T\n", msg))
		}
	} else {
		fmt.Printf(lumerinlib.FileLine() + "Zero lenth SRC Message\n")
	}
	return nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) readDstMsg() error {
	if len(c.dstMsgBuf) > 0 {
		msg, err := getStratumMsg(c.dstMsgBuf)
		if err != nil {
			panic("Bad Dst Json message")
		}

		fmt.Printf(lumerinlib.FileLine()+"DST MSG: type:%T\n--\tdata:%v\n", msg, msg)
		// c.dstConn.msgNotice = nil
		// c.dstConn.msgResponce = nil
		// c.dstConn.msgRequest = nil
		switch msg.(type) {
		case *notice:
			c.dstConn.addNoticeMsg(msg.(*notice))
		case *responce:
			c.dstConn.addResponceMsg(msg.(*responce))
		case *request:
			c.dstConn.addRequestMsg(msg.(*request))
		default:
			panic(fmt.Sprintf("what the hell is this?  :%T\n", msg))
		}
	} else {
		fmt.Printf(lumerinlib.FileLine() + "Zero lenth DST Message\n")
	}
	return nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) waitSrcMsg() error {
	fmt.Printf("waitSrcMsg()...\n")

	if c.connectionState == stateClosed || c.connectionState == stateSrcClosed || c.connectionState == stateShutdown {
		panic("should not be here")
	}

	select {
	case c.srcMsgBuf = <-c.srcConn.ch:
		c.readSrcMsg()
	}
	return nil
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) waitDstMsg() error {
	fmt.Printf("waitDstMsg()...\n")

	if c.connectionState == stateClosed || c.connectionState == stateSrcClosed || c.connectionState == stateShutdown {
		panic("should not be here")
	}

	select {
	case c.dstMsgBuf = <-c.dstConn.ch:
		c.readDstMsg()
	}
	return nil
}

//------------------------------------------
// Need to add timeout in here
//------------------------------------------
func (c *connection) waitMsgEvent() error {

	if !c.isMsgEventReady() {

		fmt.Printf(lumerinlib.Funcname() + " Msg Ready\n")

		if c.connectionState == stateClosed || c.connectionState == stateSrcClosed || c.connectionState == stateShutdown {
			panic(lumerinlib.FileLine() + " should not be here")
		}

		//
		// Wait on an event, then Read the event from
		// src/dst Socket Chan, Miner Event
		//

		ok := false
		select {
		case c.srcMsgBuf = <-c.srcConn.ch:
			c.readSrcMsg()
			ok = true

		case c.dstMsgBuf = <-c.dstConn.ch:
			c.readDstMsg()
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

	fmt.Printf(lumerinlib.Funcname()+" minerID: %s, State:%s -> %s\n", c.minerID, c.connectionState, s)

	if c.connectionState == s {
		fmt.Printf("update_connectionState() Current state is the desired state %v\n", s)
		return
	}

	// If connectionID is not defined yet, there is no msgbus entry to update
	if c.connectionID != "" {
		c.ps.ConnUpdateStateWait(c.connectionID, connectionStateMap[s])
	}

	c.connectionState = s

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) getMinerDestID() (ret msgbus.DestID, err error) {

	if c.minerID == "" {
		return "", fmt.Errorf("MinerID not set")
	}

	miner, err := c.ps.MinerGetWait(c.minerID)

	if err != nil {
		panic(fmt.Sprintf(lumerinlib.Funcname()+" Unable to get Record %s\n", err))
	}

	ret = miner.Dest

	return ret, err
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) getOrCreateMinerByName(name string) (ret msgbus.MinerID, err error) {

	var minerStruct msgbus.Miner = msgbus.Miner{
		ID:   msgbus.MinerID(""),
		Name: name,
	}
	//
	// Create entry into connection table
	//
	searchEvent, err := c.ps.SearchNameWait(msgbus.MinerMsg, name)
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" SearchNameWait failed %s\n", err))
	}
	if _, ok := searchEvent.Data.(msgbus.IDIndex); !ok {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" SearchNameWait returned wrong type %T\n", searchEvent.Data))
	}

	if len(searchEvent.Data.(msgbus.IDIndex)) == 0 {

		// No miner record found, so create one

		minerStruct.ID = msgbus.MinerID(msgbus.GetRandomIDString())
		minerStruct.Name = name

		pubEvent, err := c.ps.PubWait(
			msgbus.MinerMsg,
			msgbus.IDString(minerStruct.ID),
			minerStruct)

		if err != nil {
			return ret, err
		}
		if pubEvent.Err != nil {
			return ret, pubEvent.Err
		}

		ech := c.ps.NewEventChan()
		e, err := c.ps.SubWait(
			msgbus.MinerMsg,
			msgbus.IDString(minerStruct.ID),
			ech,
		)
		if err != nil {
			return ret, err
		}
		if e.Err != nil {
			return ret, e.Err
		}

		go c.goMonitorMinerRecord(ech)

		ret = minerStruct.ID
		return ret, nil

	} else {

		// Miner found, so return Miner ID

		return msgbus.MinerID((searchEvent.Data.(msgbus.IDIndex))[0]), nil

	}

}

func (c *connection) goMonitorMinerRecord(ech msgbus.EventChan) {

	for minerevent := range ech {

		// Need to chech if things have changed or not

		switch minerevent.EventType {
		case msgbus.SubscribedEvent:
			// Ignore
		case msgbus.UpdateEvent:
			c.updateConnectionState(stateRouteChange)

		default:
			panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", minerevent))

		}
	}

}

//------------------------------------------
//
// This should all be moved to the protocol handler,
// and this becomes an abstration
//------------------------------------------
func (c *connection) handleSrcRequest() {

	r := c.getSrcRequestMsg()
	if r == nil {
		fmt.Printf(lumerinlib.FileLine() + "Nothing returned\n")
		return
	}

	method := r.Method
	id := r.ID

	if c.requestid != id-1 {
		fmt.Printf(lumerinlib.FileLine()+"Request ID out of sequence newID: %d, currID: %d\n", id, c.requestid)
	}

	switch method {
	case string(CLIENT_MINING_AUTHORIZE):

		if c.connectionState != stateAuthStep0 {
			panic(fmt.Sprintf("Authorization in state %s\n", c.connectionState))
		}

		name, err := r.getAuthName()
		if err != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+" gatAuthName failed: %s\n", err))
		}

		// Pull username from DestID
		if c.destID != "" {
			dest, err := c.ps.DestGetWait(c.destID)
			if err != nil {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" DestGetWait failed: %s\n", err))
			}

			n := dest.Username()
			if n != "" {
				name = n
			}
		}

		minerID, err := c.getOrCreateMinerByName(name)
		if err != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+" getOrCreateMinerByName failed: %s\n", err))
		}
		c.minerID = minerID
		// c.connectionState = stateAuthStep1
		c.updateConnectionState(stateAuthStep1)

		// destID, err := c.getMinerDestID()
		// if err != nil {
		// 	panic(fmt.Sprintf(lumerinlib.FileLine()+" getOrCreateMinerByName failed: %s\n", err))
		// }
		// c.destID = destID

	case string(MINING_CONFIGURE):
	// case string(SERVER_MINING_NOTIFY):
	// case string(SERVER_MINING_SET_DIFFICULTY):
	case string(MINING_SET_TARGET):
	case string(CLIENT_MINING_SUBMIT):
		// {
		//   "method": "mining.submit",
		//   "id":4,
		//   "params": ["testrig2", "613b47ea000002b9", "0000000000000000", "613ba9ff", "a4360200"]
		// }
		//   0) Worker Name 1) Job ID 2) Extra Nonce HEX 3) nTime 4) nOnce
		//

		// Skim submit info here

	case string(CLIENT_MINING_SUBSCRIBE):
		// Grab subscription info from the miner

		if c.connectionState != stateSubscribeStep0 {
			panic(fmt.Sprintf("Subscribing in state %s\n", c.connectionState))
		}
		// c.connectionState = stateSubscribeStep1
		c.updateConnectionState(stateSubscribeStep1)

	default:
		panic(fmt.Sprintf("Method not handled: %s\n", method))
	}

	c.dstConn.sendRequest(r)

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleSrcResponce() {

	r := c.getSrcResponceMsg()
	if r == nil {
		fmt.Printf(lumerinlib.FileLine() + "Nothing returned\n")
		return
	}

	// Add in handling here
	// switch c.connectionState {
	// case stateSubscribeStep1:
	// c.connectionState = stateAuthStep1
	// }
	fmt.Printf(lumerinlib.Funcname()+" connectionState:%s\n", c.connectionState)

	c.dstConn.sendResponce(r)
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleSrcNotice() {

	n := c.getSrcNoticeMsg()
	if n == nil {
		fmt.Printf(lumerinlib.FileLine() + "Nothing returned\n")
		return
	}
	// Add in handling here

	// switch c.connectionState {
	// case stateSubscribeStep1:
	//c.connectionState = stateAuthStep0
	//	c.updateConnectionState(stateAuthStep0)
	// }
	// c.connectionState = stateAuthStep1
	fmt.Printf(lumerinlib.Funcname()+" connectionState:%s\n", c.connectionState)

	c.dstConn.sendNotice(n)
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleDstRequest() {

	r := c.getDstRequestMsg()
	if r == nil {
		fmt.Printf(lumerinlib.FileLine() + "Nothing returned\n")
		return
	}
	method := r.Method

	switch method {
	// case string(CLIENT_MINING_AUTHORIZE):
	case string(MINING_CONFIGURE):
		fallthrough
	case string(SERVER_MINING_NOTIFY):
		fallthrough
	case string(SERVER_MINING_PING):
		fallthrough
	case string(SERVER_MINING_SET_DIFFICULTY):
		fallthrough
	case string(MINING_SET_TARGET):
		fmt.Printf(lumerinlib.Funcname()+" connectionState:%s\n", c.connectionState)
	// case string(CLIENT_MINING_SUBMIT):
	// case string(CLIENT_MINING_SUBSCRIBE):
	//	if c.connectionState != stateOpenStep1 {
	//		panic(fmt.Sprintf("Subscribing in state %s\n", c.connectionState))
	//	}

	default:
		panic(fmt.Sprintf("Method not handled: %s\n", method))
	}

	c.srcConn.sendRequest(r)

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleDstResponce() {

	r := c.getDstResponceMsg()
	if r == nil {
		fmt.Printf(lumerinlib.FileLine() + "Nothing returned\n")
		return
	}
	// Looking for a result with a True
	switch c.connectionState {

	case stateAuthStep1:
		result, err := r.getAuthResult()
		if err != nil {
			panic("")
		}

		if result {
			c.updateConnectionState(stateHashVerify)
		} else {
			panic("")
		}

	case stateSubscribeStep1:
		c.updateConnectionState(stateAuthStep0)

	default:
		fmt.Printf(lumerinlib.Funcname()+" connectionState:%s\n", c.connectionState)
	}

	c.srcConn.sendResponce(r)
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleDstNotice() {

	n := c.getDstNoticeMsg()
	if n == nil {
		fmt.Printf(lumerinlib.FileLine() + "Nothing returned\n")
		return
	}
	// Add in handling here
	switch c.connectionState {
	case stateSubscribeStep1:
		//c.connectionState = stateAuthStep0
		c.updateConnectionState(stateAuthStep0)
	case stateAuthStep1:
		fmt.Printf(" stateAuthStep1 Notice\n")
	default:
		fmt.Printf(lumerinlib.Funcname()+" connectionState:%s\n", c.connectionState)
	}

	c.srcConn.sendNotice(n)
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleMinerEvent() {

	event, err := c.ps.GetWait(msgbus.MinerMsg, msgbus.IDString(c.minerID))
	if err != nil {
		panic("")
	}
	if event.Err != nil {
		panic("")
	}

	if event.Data.(msgbus.Miner).ID != c.minerID {

		c.updateConnectionState(stateRouteChange)

	} else {
		fmt.Printf("Recieved Miner Event, but the ID has not changed, ignoring\n")
	}
}

//------------------------------------------
// Starting point for new src connections
//
// Create src socket reader channel
// Create entry in connection table (Pub)
// Subscribe to updates for the connection table entry(Sub)
// Set State to Routing
//
// Does not "read" any channels, passes that to the next state handler
//------------------------------------------
func (c *connection) handleStateNew() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	err := c.srcConn.setupSocket()
	if err != nil {
		fmt.Printf(lumerinlib.FileLine() + "setupSocket() returned error\n")
		// This thing is basically stillborn
		// c.updateConnectionState(stateShutdown)
		c.connectionState = stateShutdown
	}

	if c.connectionID == "" {

		fmt.Printf("Create new SRC connection\n")

		//
		// Assign the connection a random ID
		//
		c.connectionID = msgbus.ConnectionID(msgbus.GetRandomIDString())
		c.connectionState = stateNew

		//
		// Start the src socket reader
		//
		c.srcConn.runSocketReader()

		dest := c.getDefaultDest()

		var connStruct msgbus.Connection = msgbus.Connection{
			ID:        c.connectionID,
			Miner:     msgbus.MinerID(""),
			Dest:      dest.ID,
			State:     msgbus.ConnNewState,
			TotalHash: 0,
			StartDate: time.Now(),
		}

		//
		// Create entry into connection table
		//
		event, err := c.ps.PubWait(
			msgbus.ConnectionMsg,
			msgbus.IDString(c.connectionID),
			connStruct)
		if err != nil {
			panic(lumerinlib.FileLine() + fmt.Sprintf("Error PubWait() error: %s\n", err))
		}
		if event.Err != nil {
			panic(lumerinlib.FileLine() + fmt.Sprintf("Error PubWait() Event error: %s\n", err))
		}

	}

	c.updateConnectionState(stateRouting)

}

//------------------------------------------
//
// Lookup Routing Data
// Should get a default pool address to start
//------------------------------------------
func (c *connection) handleStateRouting() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	if c.isSrcSocketClosed() {
		c.updateConnectionState(stateSrcClosed)
		return
	}

	//
	// Close any existing connection
	//
	c.dstConn.close()

	// Is there a MinerID
	// Get current miner DestID
	// -> if empty -> default DestID
	// -> if not empty -> us it

	var destid msgbus.DestID = ""

	if c.minerID != "" {
		var err error

		destid, err = c.getMinerDestID()
		if err != nil {
			fmt.Printf(lumerinlib.FileLine()+"getMinerDestID() returned error:%s\n", err)
			//c.connectionState = stateShutdown
			c.updateConnectionState(stateShutdown)
			return
		}
	}

	var dest msgbus.Dest

	if destid == "" {
		dest = c.getDefaultDest()
		c.destID = dest.ID
	} else {
		if c.ps.DestExistsWait(destid) == false {
			fmt.Printf(lumerinlib.FileLine() + " DestExistsWait() is false\n")
			c.updateConnectionState(stateShutdown)
			return
		}
		c.destID = destid
	}

	c.updateConnectionState(stateConnecting)

}

//------------------------------------------
//
// Connect to a the destination
//------------------------------------------
func (c *connection) handleStateConnecting() {

	fmt.Printf(lumerinlib.Funcname() + "\n")

	if c.isSrcSocketClosed() {
		c.updateConnectionState(stateSrcClosed)
		return
	}

	if c.destID == "" {
		fmt.Printf(lumerinlib.FileLine() + " c.destid not set\n")
	}

	dest, err := c.ps.DestGetWait(c.destID)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine() + "GetWait() returned error\n")
		c.updateConnectionState(stateShutdown)
		return
	}

	proto := "tcp" // Hardcoding for now...
	host := dest.Host()
	port := dest.Port()

	// Open connection
	fmt.Printf(lumerinlib.Funcname()+" Open Connection to: %s:%s\n", host, port)
	err = c.dstConn.dial(string(proto), string(host), string(port))
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+" setupSocket() returned error: %s\n", err)
		// c.connectionState = stateConnectError
		c.updateConnectionState(stateConnectError)
		return
	}

	fmt.Printf(lumerinlib.Funcname()+" Connected to: %s:%s\n", host, port)
	c.updateConnectionState(stateSubscribeStep0)

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
func (c *connection) handleStateSubscribeStep0() {

	fmt.Printf(lumerinlib.Funcname() + "\n")

	if c.connectionState != stateSubscribeStep0 {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" bad State: %s\n", c.connectionState))
	}

	for c.connectionState == stateSubscribeStep0 {

		if c.isSrcMsgReady() {
			switch {
			case c.isSrcSocketClosed():
				c.updateConnectionState(stateSrcClosed)

			case c.isDstSocketClosed():
				c.updateConnectionState(stateError)

			case c.isSrcRequestReady():
				c.handleSrcRequest()

			default:
				panic(fmt.Sprintf(lumerinlib.FileLine() + "Default reached in StateOpen\n"))
			}
		}

		if c.connectionState == stateSubscribeStep0 {
			err := c.waitSrcMsg()
			if err != nil {
				fmt.Printf(lumerinlib.FileLine() + "waitMsgEvent() returned error\n")
				c.connectionState = stateError
				return
			}
		}
	}

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
func (c *connection) handleStateSubscribeStep1() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	if c.connectionState != stateSubscribeStep1 {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" bad State: %s\n", c.connectionState))
	}

	for c.connectionState == stateSubscribeStep1 {

		if c.isDstMsgReady() {
			switch {
			case c.isSrcSocketClosed():
				c.updateConnectionState(stateSrcClosed)

			case c.isDstSocketClosed():
				c.updateConnectionState(stateError)

			case c.isDstResponceReady():
				c.handleDstResponce()

			case c.isDstNoticeReady():
				c.handleDstNotice()

			default:
				panic("Default reached in " + lumerinlib.Funcname() + "\n")
			}
		}

		if c.connectionState == stateSubscribeStep1 {
			err := c.waitDstMsg()
			if err != nil {
				fmt.Printf(lumerinlib.FileLine() + "waitMsgEvent() returned error\n")
				c.connectionState = stateError
				return
			}
		}
	}

}

//------------------------------------------
//
// {"id": 2, "method": "mining.authorize", "params": ["testrig", ""]}
// {"params":[32],"id":null,"method":"mining.set_difficulty"}
//
//------------------------------------------
func (c *connection) handleStateAuthStep0() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	if c.connectionState != stateAuthStep0 {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" bad State: %s\n", c.connectionState))
	}

	for c.connectionState == stateAuthStep0 {
		if c.isSrcMsgReady() {
			switch {
			case c.isSrcSocketClosed():
				c.updateConnectionState(stateSrcClosed)

			case c.isDstSocketClosed():
				c.updateConnectionState(stateError)

			case c.isSrcRequestReady():
				c.handleSrcRequest()

			default:
				panic("Default reached in " + lumerinlib.Funcname() + "\n")
			}
		}

		if c.connectionState == stateAuthStep0 {
			err := c.waitSrcMsg()
			if err != nil {
				fmt.Printf(lumerinlib.FileLine() + "waitSrcMsg() returned error\n")
				c.connectionState = stateError
				return
			}
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
func (c *connection) handleStateAuthStep1() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	for c.connectionState == stateAuthStep1 {
		if c.isDstMsgReady() {
			switch {
			case c.isSrcSocketClosed():
				c.updateConnectionState(stateSrcClosed)

			case c.isDstSocketClosed():
				c.updateConnectionState(stateError)

			case c.isDstResponceReady():
				c.handleDstResponce()

			case c.isDstNoticeReady():
				c.handleDstNotice()

			default:
				panic("Default reached in " + lumerinlib.Funcname() + "\n")
			}
		}

		if c.connectionState == stateAuthStep1 {
			err := c.waitDstMsg()
			if err != nil {
				fmt.Printf(lumerinlib.FileLine() + "waitDstMsg() returned error\n")
				c.connectionState = stateError
				return
			}
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

	for c.connectionState == stateHashVerify {

		if c.isSrcOrDstMsgReady() {
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

			default:
				panic("Default reached in " + lumerinlib.Funcname() + "\n")
			}
		}

		c.updateConnectionState(stateRouting)
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

			case c.isDstSocketClosed():
				c.updateConnectionState(stateDstClosed)

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
				c.handleMinerEvent()

			default:
				panic("Default reached in " + lumerinlib.Funcname() + "\n")
			}
		}

		// Whats this for again?
		//if c.connectionState == stateConnected {
		//	err := c.waitDstMsg()
		//	if err != nil {
		//		fmt.Printf(lumerinlib.FileLine() + "waitDstMsg() returned error\n")
		//		c.connectionState = stateError
		//		return
		//	}
		//}
	}

}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateRouteChange() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	c.dstConn.close()

	c.updateConnectionState(stateRouting)
}

//------------------------------------------
//
//------------------------------------------
func (c *connection) handleStateMsgError() {

	fmt.Printf("Enter " + lumerinlib.Funcname() + "\n")

	switch {
	case c.isSrcSocketClosed():
		c.updateConnectionState(stateSrcClosed)
		return
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
		c.updateConnectionState(stateSrcClosed)
		return
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
		panic(lumerinlib.FileLine() + " isDstRequestReady\n")
	case c.isDstResponceReady():
		panic(lumerinlib.FileLine() + " isDstResponseReady\n")
	case c.isDstNoticeReady():
		panic(lumerinlib.FileLine() + " isDstNoticeReady\n")

	default:
		panic(lumerinlib.FileLine() + " default reached\n")
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

	fmt.Printf(lumerinlib.Funcname() + "\n")

	if c.isSrcSocketClosed() {
		return
	}

	// Close out any resources here
	// Update PubSub

	c.updateConnectionState(stateRouting)

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

	c.handleStateNew()
	c.handleStateRouting()
	c.handleStateConnecting()

	for {

		if c.isSrcSocketClosed() {
			c.updateConnectionState(stateSrcClosed)
		}

		if c.isDstSocketClosed() {
			c.updateConnectionState(stateDstClosed)
		}

		switch c.connectionState {
		case stateNew:
			c.handleStateNew()

		//
		//
		case stateRouting:
			c.handleStateRouting()

			//
			//
		case stateConnecting:
			c.handleStateConnecting()

			// Open Default Dest
			//
		case stateSubscribeStep0:
			c.handleStateSubscribeStep0()

			// Request miner.subscribe  ID #n
		case stateSubscribeStep1:
			c.handleStateSubscribeStep1()

			// Result  ID #n
		case stateAuthStep0:
			c.handleStateAuthStep0()

			//
			// Request miner.authorize  ID #n
		case stateAuthStep1:
			c.handleStateAuthStep1()

			//
			// Result ID #n
		case stateHashVerify:
			c.updateConnectionState(stateConnected)
		//	c.handleStateHashVerify()

		//
		//
		case stateConnected:
			c.handleStateConnected()

			//
			//
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
		if c.connectionState == stateClosed ||
			c.connectionState == stateSrcClosed ||
			c.connectionState == stateShutdown ||
			c.connectionState == stateRouting ||
			c.connectionState == stateConnecting {
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
func (cm *ConnectionManager) Start() (err error) {

	ip, err := config.ConfigGetVal(config.ConfigConnectionListenIP)
	if err != nil {
		panic(err)
	}
	port, err := config.ConfigGetVal(config.ConfigConnectionListenPort)
	if err != nil {
		panic(err)
	}

	listener := ip + ":" + port

	l, err := net.Listen("tcp", listener)
	if err != nil {
		fmt.Printf("Listener Error for %s, %s\n", listener, err)
		return err
	}

	go cm.listenForIncomingConnections(l)

	fmt.Printf("Connection Manager Started\n")

	return err
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
