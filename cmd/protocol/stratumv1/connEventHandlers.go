package stratumv1

import (
	"fmt"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/connectionmanager"
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
	dest := scoe.Dest()

	// Need a new Dst Conn Connection Record.

	e = svs.protocol.GetDstStruct().NewProtocolDstStruct(scoe)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" NewProtocolDstStruct() error:%s", e)
		return e
	}
	svs.dstState[uid] = DstStateNew
	svs.dstDest[uid] = dest

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
	//if svs.protocol.GetDefaultRouteUID() < 0 {
	//	svs.protocol.SetDefaultRouteUID(uid)
	// }

	return e
}

//
// handleConnReadEvent()
// Look up the connection index
// Parse the message
// Run the message through a handle routine
//
func (svs *StratumV1Struct) handleConnReadEvent(scre *simple.SimpleConnReadEvent) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called ")

	uid := scre.UniqueID()

	// Validate the index is good HERE

	e = scre.Err()
	if nil != e {
		switch e {
		case connectionmanager.ErrConnMgrClosed:
			contextlib.Logf(svs.Ctx(), contextlib.LevelDebug, lumerinlib.FileLineFunc()+" Connection Manager Closed, closing down the stratum connection here")
			svs.Cancel()
			return nil
		case connectionmanager.ErrConnDstClosed:
			contextlib.Logf(svs.Ctx(), contextlib.LevelDebug, lumerinlib.FileLineFunc()+" Dest:%d  Closed Reconnecting", uid)
			svs.SetDstStateUid(uid, DstStateRedialing)
			e = svs.protocol.AsyncReDial(uid)
			return e
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Dest:%d  Error not handled:%s", uid, e)
		}

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
			svs.handleSrcReqCapabilities(request)
		case string(CLIENT_MINING_EXTRANONCE):
			svs.handleSrcReqExtranonce(request)
		case string(CLIENT_MINING_AUTHORIZE):
			e = svs.handleSrcReqAuthorize(request)
		case string(CLIENT_MINING_SUBSCRIBE):
			e = svs.handleSrcReqSubscribe(request)
		case string(CLIENT_MINING_SUBMIT):
			svs.handleSrcReqSubmit(request)
		case string(CLIENT_MINING_SUGGEST_DIFFICULTY):
			svs.handleSrcReqDifficulty(request)
		case string(CLIENT_MINING_SUGGEST_TARGET):
			svs.handleSrcReqSuggestTarget(request)
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Bad Destination Message Type Recieved:%s", request.Method)
		}

		// Received Dst message
	} else {
		switch request.Method {
		case string(SERVER_MINING_NOTIFY):
			svs.handleDstReqNotify(uid, request)
		case string(SERVER_GET_VERSION):
			svs.handleDstReqGetVersion(uid, request)
		case string(SERVER_SHOW_MESSAGE):
			svs.handleDstReqShowMessage(uid, request)
		case string(SERVER_MINING_PING):
			svs.handleDstReqPing(uid, request)
		case string(SERVER_MINING_SET_EXTRANONCE):
			svs.handleDstReqSetExtranonce(uid, request)
		case string(SERVER_MINING_SET_GOAL):
			svs.handleDstReqSetGoal(uid, request)
		case string(SERVER_MINING_SET_DIFFICULTY):
			svs.handleDstReqSetDifficulty(uid, request)
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Bad Destination Message Type Recieved:%s", request.Method)
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
		srcstate := svs.GetSrcState()
		switch srcstate {
		case SrcStateSubscribed:
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" not handled yet")
		}

	} else {
		dststate := svs.GetDstStateUid(uid)
		switch dststate {
		//
		// Got Response to Subscribe, send authorize now
		// drop the response message
		// push the auth message back
		//
		case DstStateSubscribing:
			svs.SetDstStateUid(uid, DstStateAuthorizing)
			request := svs.srcAuthRequest
			username := svs.dstDest[uid].Username()
			password := svs.dstDest[uid].Password()
			msg, e := request.createAuthorizeRequestMsg(username, password)
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

		//
		// Got Response to Authorize, Set state to running now
		// drop the response message
		//
		case DstStateAuthorizing:
			svs.SetDstStateUid(uid, DstStateStandBy)

			// Called in case this is the only connection opened.
			svs.switchDest()

		//
		// Pass response messages when in Running State
		//
		case DstStateRunning:

			msg, e := response.createResponseMsg()
			if e != nil {
				contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createResponseMsg() error:%s", e)
			}

			// Write to the current destination

			LogJson(svs.Ctx(), "Response DST -> [SRC]:", msg)

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

		case DstStateStandBy:

			msg, e := response.createResponseMsg()
			if e != nil {
				contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createResponseMsg() error:%s", e)
			}
			LogJson(svs.Ctx(), "Response in StandBy - DROPPED:", msg)

			contextlib.Logf(svs.Ctx(), contextlib.LevelDebug, lumerinlib.FileLineFunc()+" state not handled yet:%s", dststate)

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
			svs.handleDstNoticeNotify(uid, notice)
		case string(SERVER_MINING_SET_DIFFICULTY):
			svs.handleDstNoticeSetDifficulty(uid, notice)
		case string(SERVER_RECONNECT):
			svs.handleDstNoticeReconnect(uid, notice)
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
// handleSrcReqSubscribe()
// State Transition: SrcStateNew -> SrcStateSubscribed
//
func (svs *StratumV1Struct) handleSrcReqSubscribe(request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	state := svs.GetSrcState()

	// Validate the current sstate of the SRC connection
	switch state {
	case SrcStateNew:
		// This is what we expect, so skip
	case SrcStateSubscribed:
		// Returing subscribe, something went wrong.
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Got Subscribe, but already subscribed")
		return ErrBadSrcState
	case SrcStateAuthorized:
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Got Subscribe, but already authorized")
		return ErrBadSrcState
	default:
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

	dst := contextlib.GetDest(svs.Ctx())
	if dst == nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetDest() returned nil")
	}

	response := &stratumResponse{}
	msg, e := response.createSrcSubscribeResponseMsg(id)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createResponseMsg error:%s", e)
		return e
	}

	LogJson(svs.Ctx(), "SRC -> [SRC]:", msg)

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
// handleSrcReqAuthorize()
// State Transition: SrcStateSubscribed -> SrcStateAuthorized
// Open Comm to Default Dst Here
//
func (svs *StratumV1Struct) handleSrcReqAuthorize(request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	state := svs.GetSrcState()
	// Validate the current sstate of the SRC connection
	switch state {
	case SrcStateNew:
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Got Authorize, expect subscribe")
		return ErrBadSrcState
	case SrcStateSubscribed:
		// This is what we expect, so continue
	case SrcStateAuthorized:
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Got Authorize, but already authorized")
		return ErrBadSrcState
	default:
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

	dstID := contextlib.GetDest(svs.Ctx())
	if dstID == nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetDest() returned nil")
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

	LogJson(svs.Ctx(), "SRC -> [SRC]:", msg)

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

	//	// Fire up the default destination connction here
	//	// If it does not alread exist
	//	dr := svs.protocol.GetDefaultRouteUID()
	//	if dr < 0 {
	//		err := svs.protocol.AsyncDial(dst)
	//		if err != nil {
	//			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" AsyncDial returned error:%s", e)
	//		}
	//	}

	return nil
}

//
// handleSrcReqSubmit()
// The system should be authorized at least, possibly other states too
//
func (svs *StratumV1Struct) handleSrcReqSubmit(request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	state := svs.GetSrcState()
	// Validate the current sstate of the SRC connection
	switch state {
	case SrcStateNew:
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Got Submit, expecting Subscribe")
		return ErrBadSrcState
	case SrcStateSubscribed:
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Got Submit, expecting Authorize")
		return ErrBadSrcState
	case SrcStateAuthorized:
	case SrcStateRunning:
		// This is what we expect, so skip
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Src state:%s", state)
	}

	// Track ID HERE?

	//id, e := request.getID()
	//if e != nil {
	//	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" getID() returned error:%s ", e)
	//	return e
	//}

	//
	// Get the current default route UID
	//
	uid := svs.protocol.GetDefaultRouteUID()

	//
	// Get the username of the default route
	//
	username := svs.dstDest[uid].Username()

	// msg, e := request.createRequestMsg()
	msg, e := request.createSubmitRequestMsg(username)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createRequestMsg() error:%s", e)
	}

	// Write to the current destination

	LogJson(svs.Ctx(), "SRC -> [DST]:", msg)

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

	// Call switchDest to change destinations oif needed (happens on a submit)
	svs.switchDest()

	// Change Destination HERE
	// Log Submission HERE

	return nil
}

//
//
//
func (svs *StratumV1Struct) handleSrcReqCapabilities(request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" this is not handled yet")
	return ErrSrcReqNotSupported
}

//
//
//
func (svs *StratumV1Struct) handleSrcReqDifficulty(request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" this is not handled yet")
	return ErrSrcReqNotSupported

}

//
//
//
func (svs *StratumV1Struct) handleSrcReqExtranonce(request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" this is not handled yet")
	return ErrSrcReqNotSupported

}

//
//
//
func (svs *StratumV1Struct) handleSrcReqSuggestTarget(request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" this is not handled yet")
	return ErrSrcReqNotSupported

}

// -------------------------------------------------------------------------
// Incoming Comm Dst REQUESTs
// -------------------------------------------------------------------------

//
// handleDstReqNotify()
// if the state of the connection is accetable, and the connection is the default pass the notification
//
//
func (svs *StratumV1Struct) handleDstReqNotify(uid simple.ConnUniqueID, request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	// is uid the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src
	dststate := svs.GetDstStateUid(uid)
	switch dststate {
	case DstStateSubscribing:
		contextlib.Logf(svs.Ctx(), contextlib.LevelDebug, lumerinlib.FileLineFunc()+" Dst Not Subscribed yet, dropping Notifying Request")
		return nil
	case DstStateAuthorizing:
		contextlib.Logf(svs.Ctx(), contextlib.LevelDebug, lumerinlib.FileLineFunc()+" Dst Not Authorized yet, dropping Notifying Request")
		return nil
	case DstStateRunning:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" passing notify for state:%s", dststate)
	case DstStateStandBy:
		contextlib.Logf(svs.Ctx(), contextlib.LevelDebug, lumerinlib.FileLineFunc()+" state not handled yet:%s", dststate)
		msg, e := request.createRequestMsg()
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createResponseMsg() error:%s", e)
		}
		LogJson(svs.Ctx(), "Request in StandBy - DROPPED:", msg)
	// case DstStateError:
	// case DstStateNew:
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Dst ReqNotify state:%s not handled", dststate)
		return ErrDstReqNotSupported
	}

	defRouteUid := svs.protocol.GetDefaultRouteUID()

	// This is the default route
	if defRouteUid == uid {
		msg, e := request.createReqMiningNotify()
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createNoticeMiningNotify() returned error:%s", e)
			return e
		}

		LogJson(svs.Ctx(), "DST -> [SRC]:", msg)

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
func (svs *StratumV1Struct) handleDstReqGetVersion(UID simple.ConnUniqueID, request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" this is not handled yet")
	return ErrDstReqNotSupported

}

//
//
//
func (svs *StratumV1Struct) handleDstReqPing(UID simple.ConnUniqueID, request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" this is not handled yet")
	return ErrDstReqNotSupported

}

//
//
//
func (svs *StratumV1Struct) handleDstReqShowMessage(UID simple.ConnUniqueID, request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" this is not handled yet")
	return ErrDstReqNotSupported

}

//
//
//
func (svs *StratumV1Struct) handleDstReqSetExtranonce(UID simple.ConnUniqueID, request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" this is not handled yet")
	return ErrDstReqNotSupported

}

//
//
//
func (svs *StratumV1Struct) handleDstReqSetGoal(UID simple.ConnUniqueID, request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" this is not handled yet")
	return ErrDstReqNotSupported

}

//
// handleDstReqSetDifficulty()
// handles incomin set difficulty message from a pool connection
//
func (svs *StratumV1Struct) handleDstReqSetDifficulty(uid simple.ConnUniqueID, request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	dststate := svs.GetDstStateUid(uid)
	switch dststate {
	case DstStateSubscribing:
		contextlib.Logf(svs.Ctx(), contextlib.LevelDebug, lumerinlib.FileLineFunc()+" Dst Not Subscribed yet, dropping Notifying Request")
		return nil
	case DstStateAuthorizing:
		contextlib.Logf(svs.Ctx(), contextlib.LevelDebug, lumerinlib.FileLineFunc()+" Dst Not Authorized yet, dropping Notifying Request")
		return nil
	case DstStateRunning:
		contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" passing notify for state:%s", dststate)
	case DstStateStandBy:
		contextlib.Logf(svs.Ctx(), contextlib.LevelDebug, lumerinlib.FileLineFunc()+" state not handled yet:%s", dststate)
		msg, e := request.createRequestMsg()
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createResponseMsg() error:%s", e)
		}
		LogJson(svs.Ctx(), "Request in StandBy - DROPPED:", msg)

	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Dst ReqNotify state:%s not handled", dststate)
		return ErrDstReqNotSupported
	}

	defRouteUid := svs.protocol.GetDefaultRouteUID()
	// This is the default route
	if defRouteUid == uid {
		msg, e := request.createRequestMsg()
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createNoticeSetDifficultyMsg() returned error:%s", e)
			return e
		}

		LogJson(svs.Ctx(), "DST -> [SRC]:", msg)

		svs.protocol.WriteSrc(msg)
	} else {
		// Store or drop the message?
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" uid:%d is not the default dst:%d, should we store or drop the message", uid, defRouteUid)
	}

	return e
}

// -------------------------------------------------------------------------
// Incoming Comm NOTICEs
// -------------------------------------------------------------------------

//
// handleDstNotify()
// if the state of the connection is accetable, and the connection is the default pass the notification
//
//
func (svs *StratumV1Struct) handleDstNoticeNotify(uid simple.ConnUniqueID, notice *stratumNotice) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	// is uid the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src
	dststate := svs.GetDstStateUid(uid)
	switch dststate {
	case DstStateAuthorizing:
		fallthrough
	case DstStateSubscribing:
		fallthrough
	case DstStateRunning:
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" passing notify for state:%s", dststate)
	case DstStateStandBy:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" state not handled yet:%s", dststate)
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

		LogJson(svs.Ctx(), "DST -> [SRC]:", msg)

		svs.protocol.WriteSrc(msg)
	} else {
		// Store or drop the message?
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" uid:%d is not the default dst:%d, should we store or drop the message", uid, defRouteUid)
	}

	return e
}

//
// handleDstSetNoticeDifficulty()
// handles incomin set difficulty message from a pool connection
//
func (svs *StratumV1Struct) handleDstNoticeSetDifficulty(uid simple.ConnUniqueID, notice *stratumNotice) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	// is index the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src

	dststate := svs.GetDstStateUid(uid)
	switch dststate {
	case DstStateSubscribing:
		fallthrough
	case DstStateAuthorizing:
		fallthrough
	case DstStateRunning:
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" passing set diff for state:%s", dststate)
	case DstStateStandBy:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" state not handled yet:%s", dststate)
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

		LogJson(svs.Ctx(), "DST -> [SRC]:", msg)

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
func (svs *StratumV1Struct) handleDstNoticeReconnect(uid simple.ConnUniqueID, notice *stratumNotice) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter")

	// is index the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src

	dststate := svs.GetDstStateUid(uid)
	switch dststate {
	case DstStateSubscribing:
		fallthrough
	case DstStateAuthorizing:
		fallthrough
	case DstStateRunning:
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" passing set diff for state:%s", dststate)
	case DstStateStandBy:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" state not handled yet:%s", dststate)
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

		LogJson(svs.Ctx(), "DST -> [SRC]:", msg)

		svs.protocol.WriteSrc(msg)
	} else {
		// Store or drop the message?
		contextlib.Logf(svs.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" uid:%d is not the default dst:%d, should we store or drop the message", uid, defRouteUid)
	}

	return e

}
