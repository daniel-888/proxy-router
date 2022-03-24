package stratumv1

import (
	"fmt"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

//
// States:
// Transition to a new Destination ID
//		Is the Dest Open, if not open it
//		Set State so when the source is ready it can transition
//
//

//
// handleMsgUpdateEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgUpdateEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	// Parse Event Msg

	switch event.Msg {
	// Miner updates are of interest to me
	case msgbus.MinerMsg:
		// Check that we have the correct miner ID
		currentRec, ok := event.Data.(msgbus.Miner)
		if !ok {
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" event.Data is not a msgbus.Miner struct")
		}
		if svs.minerRec.ID != currentRec.ID {
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" event.Data records do not match up")
		}

		// Compare what has changed

		if svs.minerRec.Dest != currentRec.Dest {
			contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" Miner:%s Dest changed to: %s from %s", currentRec.ID, currentRec.Dest, svs.minerRec.Dest)
			// Destination has changed...

			// Start the process of transitioning to a new Dest

		}

	// Ignore all others
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgDeleteEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgDeleteEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgGetEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else error out
//
func (svs *StratumV1Struct) handleMsgGetEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgIndexEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else error out
//
func (svs *StratumV1Struct) handleMsgIndexEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgSearchEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else error out
//
func (svs *StratumV1Struct) handleMsgSearchEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgSearchEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else error out
//
func (svs *StratumV1Struct) handleMsgSearchIndexEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgPublishEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgPublishEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgUnpublishEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgUnpublishEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgSubscribedEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgSubscribedEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgUnisubscribedEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgUnsubscribedEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleMsgRemovedEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
//
func (svs *StratumV1Struct) handleMsgRemovedEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	switch event.Msg {
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" ignoring update to: %s:%s", event.EventType, event.ID)
	}

}

//
// handleConnReadEvent()
// Look up the connection index
// Parse the message
// Run the message through a handle routine
//
func (svs *StratumV1Struct) handleConnReadEvent(scre *simple.SimpleConnReadEvent) (e error) {

	// contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called %v", scre)
	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called ")

	uid := scre.UniqueID()

	// Validate the index is good HERE

	e = scre.Err()
	if nil != e {
		// Notate it here, but set the err in the connection later
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" scre had an error:%s", e)
	}

	//
	// Should be moved into protocol.go file since this is only protocol related.
	//
	var pcs *protocol.ProtocolConnectionStruct
	if uid < 0 {
		pcs, e = svs.protocol.GetSrcConn()
	} else {
		pcs, e = svs.protocol.GetDstConn(uid)
	}

	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" GetDstConn(%d) had an error:%s", uid, e)
	}

	// Buffer data locally
	data := scre.Data()
	if len(data) == 0 {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" scre had no data")
	}

	pcs.AddBuf(data)

	// -- Move above to Protocol --

	// If there is a '\n' in the message, extract the message
	// Adjust the buffer accordingly

	if uid == 0 {
		_ = uid
	}

	// Cycle through the buffer until all of the '\n' are found
	for buf, e := pcs.GetLineTermData(); len(buf) > 0 && e == nil; buf, e = pcs.GetLineTermData() {

		if len(buf) == 0 {
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" buffer is 0")
		}

		// Process the message
		ret, e := unmarshalMsg(buf)

		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Called data:%s", data)
			return e
		}

		switch ret := ret.(type) {
		case *stratumRequest:
			e = svs.handleRequest(uid, ret)
		case *stratumResponse:
			e = svs.handleResponse(uid, ret)
		case *stratumNotice:
			e = svs.handleNotice(uid, ret)
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Called")
		}
		if e != nil {
			break
		}
	}

	return e
}

//
// handleConnOpenEvent()
// UID is the unique ID of the DST connection
// UID is what is fed to the connection functions like Write
// New connection starts off in the new state, immediatly send a Subscribe message, then Auth
//
func (svs *StratumV1Struct) handleConnOpenEvent(scoe *simple.SimpleConnOpenEvent) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called %v", scoe)

	// If there is an error do we handle it here, or put it into the connection struct?
	e = scoe.Err()
	if nil != e {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" scre had an error:%s", e)
		return e
	}

	uid := scoe.UniqueID()

	// Need a new Dst Conn Connection Record.

	e = svs.protocol.GetDstStruct().NewProtocolDstStruct(scoe)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" NewProtocolDstStruct() error:%s", e)
		return e
	}
	svs.dstState[uid] = DstStateNew

	dstconn, e := svs.protocol.GetDstConn(uid)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" GetDstConn() bad index:%d", uid)
		return e
	}
	dstconn.SetState(protocol.ConnStateReady)

	// Send initialization subscribe message here
	// Set state to DstStateSubscribing

	request := svs.srcSubscribeRequest

	msg, e := request.createRequestMsg()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetDstConn() bad UID:%d", uid)
	}

	msgsize := len(msg)
	if e == nil {
		count, e := svs.protocol.WriteDst(uid, msg)
		if e != nil {
			svs.SetDstStateUid(uid, DstStateError)
		}

		if count != msgsize {
			svs.SetDstStateUid(uid, DstStateError)
		}

		svs.SetDstStateUid(uid, DstStateSubscribing)
	}

	// Set default route if not set already
	if svs.protocol.GetDefaultRouteUID() < 0 {
		svs.protocol.SetDefaultRouteUID(uid)
	}

	return e
}

//
// handleConnEOFEvent()
// Close out the connection index
// Could reopen the connection, or perform some error handling
//
func (svs *StratumV1Struct) handleConnEOFEvent(event *simple.SimpleEvent) {

	// contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Not Handled Yet!")

}

//
// handleConnErrorEvent()
// Perform some error handling
//
func (svs *StratumV1Struct) handleConnErrorEvent(event *simple.SimpleEvent) {

	// contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Not Handled Yet!")
}

//
// handleErrorEvent()
// Perform Error handling
//
func (svs *StratumV1Struct) handleErrorEvent(event *simple.SimpleEvent) {

	// contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Not Handled Yet!")
}

//
// handleRequest()
// index: -1 = SRC, 0 = default, >0 = Dst
//
func (svs *StratumV1Struct) handleRequest(uid simple.ConnUniqueID, request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called Index: %d Method: %s", uid, request.Method)

	// Recieved Src message
	if uid < 0 {
		switch request.Method {
		case string(CLIENT_MINING_CAPABILITIES):
			svs.handleSrcCapabilities(request)
		case string(CLIENT_MINING_EXTRANONCE):
			svs.handleSrcExtranonce(request)
		case string(CLIENT_MINING_AUTHORIZE):
			e = svs.handleSrcAuthorize(request)
		case string(CLIENT_MINING_SUBSCRIBE):
			e = svs.handleSrcSubscribe(request)
		case string(CLIENT_MINING_SUBMIT):
			svs.handleSrcSubmit(request)
		case string(CLIENT_MINING_SUGGEST_DIFFICULTY):
			svs.handleSrcDifficulty(request)
		case string(CLIENT_MINING_SUGGEST_TARGET):
			svs.handleSrcSuggestTarget(request)
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Bad Destination Message Type Recieved")
		}

		// Received Dst message
	} else {
		switch request.Method {
		case string(SERVER_GET_VERSION):
			svs.handleDstGetVersion(uid, request)
		//case string(SERVER_RECONNECT):
		//	svs.handleDstReconnect(uid, request)
		case string(SERVER_SHOW_MESSAGE):
			svs.handleDstShowMessage(uid, request)
		case string(SERVER_MINING_PING):
			svs.handleDstPing(uid, request)
		case string(SERVER_MINING_SET_EXTRANONCE):
			svs.handleDstSetExtranonce(uid, request)
		case string(SERVER_MINING_SET_GOAL):
			svs.handleDstSetGoal(uid, request)
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Bad Destination Message Type Recieved")
		}
	}

	return e
}

//
// handleResponse()
// index: -1 = SRC, 0 = default, >0 = Dst
//
func (svs *StratumV1Struct) handleResponse(uid simple.ConnUniqueID, response *stratumResponse) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	if uid < 0 {
		state := svs.GetSrcState()
		switch state {
		case SrcStateSubscribed:
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" not handled yet")
		}

	} else {
		state := svs.GetDstStateUid(uid)
		switch state {
		// Got Response to Subscribe, send authorize now
		case DstStateSubscribing:
			svs.SetDstStateUid(uid, DstStateAuthorizing)
			request := svs.srcAuthRequest
			msg, e := request.createRequestMsg()
			if e != nil {
				contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" error:%s", e)
			}
			count, e := svs.protocol.WriteDst(uid, msg)
			if e != nil {
				contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" error:%s", e)
			}
			if count != len(msg) {
				contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" count:%d, len:%d", count, len(msg))
			}

		case DstStateAuthorizing:

		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" not handled yet")
		}

	}

	return nil
}

//
// handleNotice()
// index: -1 = SRC, 0 = default, >0 = Dst
//
func (svs *StratumV1Struct) handleNotice(uid simple.ConnUniqueID, notice *stratumNotice) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	if uid < 0 {
		switch notice.Method {
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Bad Destination Message Type Recieved")
		}
	} else {

		switch notice.Method {
		case string(SERVER_MINING_NOTIFY):
			svs.handleDstNotify(uid, notice)
		case string(SERVER_MINING_SET_DIFFICULTY):
			svs.handleDstSetDifficulty(uid, notice)
		case string(SERVER_RECONNECT):
			svs.handleDstReconnect(uid, notice)
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Bad Destination Message Type Recieved:%s", notice.Method)
		}
	}

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	return nil
}

// -------------------------------------------------------------------------
// Incoming Conn SRC REQUESTs
// -------------------------------------------------------------------------

//
// handleSrcSubscribe()
// State Transition: SrcStateNew -> SrcStateSubscribed
//
func (svs *StratumV1Struct) handleSrcSubscribe(request *stratumRequest) (e error) {

	state := svs.GetSrcState()

	if state != SrcStateNew {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Src state:%s", state)
	}

	if svs.srcSubscribeRequest.Method == "" {
		svs.srcSubscribeRequest = request
	}

	svs.SetSrcState(SrcStateSubscribed)

	id, e := request.getID()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" getID() error:%s", e)
		return e
	}

	dst := contextlib.GetDst(svs.Ctx())
	if dst == nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetDst() returned nil")
	}

	// Move this to JSON file

	extranonce := "1"
	extranonce2 := 1

	subscriptions := make([]string, 2)
	subscriptions[0] = string(SERVER_MINING_NOTIFY)
	subscriptions[1] = "0"

	sub2 := make([][]string, 1)
	sub2[0] = subscriptions

	result := make([]interface{}, 3)
	result[0] = sub2
	result[1] = extranonce
	result[2] = extranonce2

	response := &stratumResponse{
		ID:     id,
		Error:  nil,
		Result: result,
		Reject: nil,
	}

	msg, e := response.createResponseMsg()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createResponseMsg error:%s", e)
		return e
	}

	count, e := svs.protocol.WriteSrc(msg)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" WriteSrc error:%s, Close it down", e)
		// Error writing to Src (close it down here)
		svs.Cancel()
		return e
	}
	if count != len(msg) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
		return e
	}

	return nil
}

//
// handleSrcAuthorize()
// State Transition: SrcStateSubscribed -> SrcStateAuthorized
// Open Comm to Default Dst Here
//
func (svs *StratumV1Struct) handleSrcAuthorize(request *stratumRequest) (e error) {

	state := svs.GetSrcState()

	if state != SrcStateSubscribed {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Src state:%s", state)
	}

	if svs.srcAuthRequest.Method == "" {
		svs.srcAuthRequest = request
	}
	svs.SetSrcState(SrcStateAuthorized)

	id, e := request.getID()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" getID() returned error:%s ", e)
		return e
	}

	dst := contextlib.GetDst(svs.Ctx())
	if dst == nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetDst() returned nil")
	}

	// Move this to JSON file

	result := true

	response := &stratumResponse{
		ID:     id,
		Error:  nil,
		Result: result,
		Reject: nil,
	}

	msg, e := response.createResponseMsg()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createResponseMsg error:%s", e)
		return e
	}

	count, e := svs.protocol.WriteSrc(msg)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" WriteSrc error:%s", e)
		return e
	}
	if count != len(msg) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
		return e
	}

	// Fire up the default destination connction here
	// If it does not alread exist
	dr := svs.protocol.GetDefaultRouteUID()
	if dr < 0 {
		err := svs.protocol.AsyncDial(dst)
		if err != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" AsyncDial returned error:%s", e)
		}
	}

	return nil
}

//
// handleSrcSubmit()
// The system should be authorized at least, possibly other states too
//
func (svs *StratumV1Struct) handleSrcSubmit(request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	state := svs.GetSrcState()

	if state != SrcStateAuthorized {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Src state:%s", state)
	}

	// Track ID HERE?

	//id, e := request.getID()
	//if e != nil {
	//	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" getID() returned error:%s ", e)
	//	return e
	//}

	msg, e := request.createRequestMsg()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createRequestMsg() error:%s", e)
	}

	// Write to the current destination
	count, e := svs.protocol.Write(msg)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" WriteSrc error:%s", e)
		return e
	}
	if count != len(msg) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
		return e
	}

	// Change Destination HERE
	// Log Submission HERE

	return nil
}

//
//
//
func (svs *StratumV1Struct) handleSrcCapabilities(request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleSrcDifficulty(request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleSrcExtranonce(request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleSrcSuggestTarget(request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" this is not handled yet")

}

// -------------------------------------------------------------------------
// Incoming Comm Dst REQUESTs
// -------------------------------------------------------------------------

//
//
//
func (svs *StratumV1Struct) handleDstGetVersion(UID simple.ConnUniqueID, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstPing(UID simple.ConnUniqueID, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+"  this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstShowMessage(UID simple.ConnUniqueID, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+"  this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstSetExtranonce(UID simple.ConnUniqueID, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+"  this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstSetGoal(UID simple.ConnUniqueID, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" this is not handled yet")

}

// -------------------------------------------------------------------------
// Incoming Comm NOTICEs
// -------------------------------------------------------------------------

//
// handleDstNotify()
// if the state of the connection is accetable, and the connection is the default pass the notification
//
//
func (svs *StratumV1Struct) handleDstNotify(uid simple.ConnUniqueID, notice *stratumNotice) (e error) {

	// is uid the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src
	dststate := svs.GetDstStateUid(uid)
	switch dststate {
	case DstStateAuthorizing:
		fallthrough
	case DstStateSubscribing:
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" passing notify for state:%s", dststate)
	// case DstStateError:
	// case DstStateNew:
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+"  state:%s not handled", dststate)
	}

	defRouteUid := svs.protocol.GetDefaultRouteUID()

	// This is the default route
	if defRouteUid == uid {
		msg, e := notice.createNoticeMiningNotify()
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createNoticeMiningNotify() returned error:%s", e)
			return e
		}
		svs.protocol.WriteSrc(msg)
	} else {
		// Store or drop the message?
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" uid:%d is not the default dst:%d, should we store or drop the message", uid, defRouteUid)
	}

	return e
}

//
// handleDstSetDifficulty()
// handles incomin set difficulty message from a pool connection
//
func (svs *StratumV1Struct) handleDstSetDifficulty(uid simple.ConnUniqueID, notice *stratumNotice) (e error) {

	// is index the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src

	dststate := svs.GetDstStateUid(uid)
	switch dststate {
	case DstStateSubscribing:
		fallthrough
	case DstStateAuthorizing:
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" passing set diff for state:%s", dststate)
	// case DstStateError:
	// case DstStateNew:
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+"  state:%s not handled", dststate)
	}

	defRouteUid := svs.protocol.GetDefaultRouteUID()
	// This is the default route
	if defRouteUid == uid {
		msg, e := notice.createNoticeSetDifficultyMsg()
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createNoticeSetDifficultyMsg() returned error:%s", e)
			return e
		}
		svs.protocol.WriteSrc(msg)
	} else {
		// Store or drop the message?
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" uid:%d is not the default dst:%d, should we store or drop the message", uid, defRouteUid)
	}

	return e
}

//
//
//
func (svs *StratumV1Struct) handleDstReconnect(uid simple.ConnUniqueID, notice *stratumNotice) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	// is index the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src

	dststate := svs.GetDstStateUid(uid)
	switch dststate {
	case DstStateSubscribing:
		fallthrough
	case DstStateAuthorizing:
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" passing set diff for state:%s", dststate)
	// case DstStateError:
	// case DstStateNew:
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+"  state:%s not handled", dststate)
	}

	defRouteUid := svs.protocol.GetDefaultRouteUID()
	// This is the default route
	if defRouteUid == uid {
		// msg, e := notice.createNoticeSetDifficultyMsg()
		msg, e := notice.createNoticeMsg()
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createNoticeSetDifficultyMsg() returned error:%s", e)
			return e
		}
		svs.protocol.WriteSrc(msg)
	} else {
		// Store or drop the message?
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" uid:%d is not the default dst:%d, should we store or drop the message", uid, defRouteUid)
	}

	return e

}
