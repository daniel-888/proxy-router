package stratumv1

import (
	"fmt"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
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

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called %v", scre)

	e = scre.Err()
	if nil != e {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" scre had an error:%s", e)
		return e
	}

	// count := scre.Count()
	index := scre.Index()
	data := scre.Data()

	ret, e := unmarshalMsg(data)

	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Called")
		return e
	}

	switch ret := ret.(type) {
	case *stratumRequest:
		e = svs.handleRequest(index, ret)
	case *stratumResponse:
		e = svs.handleResponse(index, ret)
	case *stratumNotice:
		e = svs.handleNotice(index, ret)
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Called")
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
func (svs *StratumV1Struct) handleRequest(index int, request *stratumRequest) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called Index: %d Method: %s", index, request.Method)

	// Recieved Src message
	if index < 0 {
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
			svs.handleDstGetVersion(index, request)
		case string(SERVER_RECONNECT):
			svs.handleDstReconnect(index, request)
		case string(SERVER_SHOW_MESSAGE):
			svs.handleDstShowMessage(index, request)
		case string(SERVER_MINING_PING):
			svs.handleDstPing(index, request)
		case string(SERVER_MINING_SET_EXTRANONCE):
			svs.handleDstSetExtranonce(index, request)
		case string(SERVER_MINING_SET_GOAL):
			svs.handleDstSetGoal(index, request)
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
func (svs *StratumV1Struct) handleResponse(index int, request *stratumResponse) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	return nil
}

//
// handleNotice()
// index: -1 = SRC, 0 = default, >0 = Dst
//
func (svs *StratumV1Struct) handleNotice(index int, notice *stratumNotice) (e error) {

	if index < 0 {
		switch notice.Method {
		case string(SERVER_MINING_NOTIFY):
			svs.handleDstNotify(index, notice)
		case string(SERVER_MINING_SET_DIFFICULTY):
			svs.handleDstSetDifficulty(index, notice)
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Bad Destination Message Type Recieved")
		}
	} else {

		switch notice.Method {
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Bad Destination Message Type Recieved")
		}
	}

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	return nil
}

// -------------------------------------------------------------------------
// Incoming Comm REQUESTs
// -------------------------------------------------------------------------

//
//
//
func (svs *StratumV1Struct) handleSrcAuthorize(request *stratumRequest) (e error) {

	if svs.srcAuthRequest == nil {
		svs.srcAuthRequest = request
	}

	id, e := request.getID()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" getID() returned error:%s ", e)
		return e
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

	return nil
}

//
//
//
func (svs *StratumV1Struct) handleSrcCapabilities(request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleSrcDifficulty(request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleSrcExtranonce(request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

//
// handleSrcSubscribe()
// takes a parsed message from the source
//
func (svs *StratumV1Struct) handleSrcSubscribe(request *stratumRequest) (e error) {

	if svs.srcSubscribeRequest == nil {
		svs.srcSubscribeRequest = request
	}

	id, e := request.getID()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" getID() error:%s", e)
		return e
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
//
//
func (svs *StratumV1Struct) handleSrcSubmit(request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

	// Forward the request to the default destination pool

}

//
//
//
func (svs *StratumV1Struct) handleSrcSuggestTarget(request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstGetVersion(index int, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstPing(index int, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstReconnect(index int, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstShowMessage(index int, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstSetExtranonce(index int, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

//
//
//
func (svs *StratumV1Struct) handleDstSetGoal(index int, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")

}

// -------------------------------------------------------------------------
// Incoming Comm NOTICEs
// -------------------------------------------------------------------------

//
// handleDstNotify()
// passes new work from the destination pool to the source miner
//
func (svs *StratumV1Struct) handleDstNotify(index int, request *stratumNotice) {

	// is index the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src

	defroutidx := svs.protocol.GetDefaultRouteIndex()

	// This is the default route
	if defroutidx == index {
		msg, e := request.createNoticeMiningNotify()
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createNoticeMiningNotify() returned error:%s", e)
		}
		svs.protocol.WriteSrc(msg)
	} else {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")
	}

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" ... not handled yet")

}

//
// handleDstSetDifficulty()
// handles incomin set difficulty message from a pool connection
//
func (svs *StratumV1Struct) handleDstSetDifficulty(index int, request *stratumNotice) {

	// is index the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src

	defroutidx := svs.protocol.GetDefaultRouteIndex()

	// This is the default route
	if defroutidx == index {
		msg, e := request.createNoticeSetDifficultyMsg()
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createNoticeSetDifficultyMsg() returned error:%s", e)
		}
		svs.protocol.WriteSrc(msg)
	} else {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" index is not the default dst, this is not handled yet")
	}

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" ... not handled yet")

}
