package stratumv1

import (
	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

//
// handleMsgUpdateEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgUpdateEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

	// Parse Event Msg

}

//
// handleMsgDeleteEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgDeleteEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleMsgGetEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else error out
//
func (svs *StratumV1Struct) handleMsgGetEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleMsgIndexEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else error out
//
func (svs *StratumV1Struct) handleMsgIndexEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleMsgSearchEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else error out
//
func (svs *StratumV1Struct) handleMsgSearchEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleMsgSearchEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else error out
//
func (svs *StratumV1Struct) handleMsgSearchIndexEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleMsgPublishEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgPublishEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleMsgUnpublishEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgUnpublishEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleMsgSubscribedEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgSubscribedEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleMsgUnisubscribedEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
func (svs *StratumV1Struct) handleMsgUnsubscribedEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleMsgRemovedEvent()
// Check the request ID to see if it was requested, pull the request off and handle the result
// else
// Check to see if this is a event that we care about, and handle it
//
//
func (svs *StratumV1Struct) handleMsgRemovedEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleConnReadEvent()
// Look up the connection index
// Parse the message
// Run the message through a handle routine
//
func (svs *StratumV1Struct) handleConnReadEvent(event *simple.SimpleEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

	// index: -1 = SRC, 0 = default, >0 = Dst
	var index int = 0

	// Get ConnectionID
	// Is Src or Dst
	// Translate to index int
	// Parse Message

	ret, e := unmarshalMsg([]byte(event.Data.(string)))
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Called")
	}
	switch ret.(type) {
	case *stratumRequest:
		svs.handleRequest(index, ret.(*stratumRequest))
	case *stratumResponse:
		svs.handleResponse(index, ret.(*stratumResponse))
	case *stratumNotice:
		svs.handleNotice(index, ret.(*stratumNotice))
	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Called")
	}
}

//
// handleConnEOFEvent()
// Close out the connection index
// Could reopen the connection, or perform some error handling
//
func (svs *StratumV1Struct) handleConnEOFEvent(event *simple.SimpleEvent) {

	// contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Not Handled Yet!")

}

//
// handleConnErrorEvent()
// Perform some error handling
//
func (svs *StratumV1Struct) handleConnErrorEvent(event *simple.SimpleEvent) {

	// contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Not Handled Yet!")
}

//
// handleErrorEvent()
// Perform Error handling
//
func (svs *StratumV1Struct) handleErrorEvent(event *simple.SimpleEvent) {

	// contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

	contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Not Handled Yet!")
}

//
// handleRequest()
// index: -1 = SRC, 0 = default, >0 = Dst
//
func (svs *StratumV1Struct) handleRequest(index int, request *stratumRequest) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

	// Recieved Src message
	if index < 0 {
		switch request.Method {
		case string(CLIENT_MINING_CAPABILITIES):
			svs.handleSrcCapabilities(request)
		case string(CLIENT_MINING_EXTRANONCE):
			svs.handleSrcExtranonce(request)
		case string(CLIENT_MINING_AUTHORIZE):
			svs.handleSrcAuthorize(request)
		case string(CLIENT_MINING_SUBSCRIBE):
			svs.handleSrcSubscribe(request)
		case string(CLIENT_MINING_SUBMIT):
			svs.handleSrcSubmit(request)
		case string(CLIENT_MINING_SUGGEST_DIFFICULTY):
			svs.handleSrcDifficulty(request)
		case string(CLIENT_MINING_SUGGEST_TARGET):
			svs.handleSrcSuggestTarget(request)
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Bad Destination Message Type Recieved")
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
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Bad Destination Message Type Recieved")
		}

	}

}

//
// handleResponse()
// index: -1 = SRC, 0 = default, >0 = Dst
//
func (svs *StratumV1Struct) handleResponse(index int, request *stratumResponse) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleNotice()
// index: -1 = SRC, 0 = default, >0 = Dst
//
func (svs *StratumV1Struct) handleNotice(index int, notice *stratumNotice) {

	if index < 0 {
		switch notice.Method {
		case string(SERVER_MINING_NOTIFY):
			svs.handleDstNotify(index, notice)
		case string(SERVER_MINING_SET_DIFFICULTY):
			svs.handleDstSetDifficulty(index, notice)
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Bad Destination Message Type Recieved")
		}
	} else {

		switch notice.Method {
		default:
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Bad Destination Message Type Recieved")
		}
	}

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
//
//
func (svs *StratumV1Struct) handleSrcAuthorize(request *stratumRequest) {

}

//
//
//
func (svs *StratumV1Struct) handleSrcCapabilities(request *stratumRequest) {

}

//
//
//
func (svs *StratumV1Struct) handleSrcDifficulty(request *stratumRequest) {

}

//
//
//
func (svs *StratumV1Struct) handleSrcExtranonce(request *stratumRequest) {

}

//
//
//
func (svs *StratumV1Struct) handleSrcSubscribe(request *stratumRequest) {

	// Break down parameters into user and password
	// store these in the StratumV1Struct

	// If there is any open connection send the subscribe

}

//
//
//
func (svs *StratumV1Struct) handleSrcSubmit(request *stratumRequest) {

	// Forward the request to the default destination pool

}

//
//
//
func (svs *StratumV1Struct) handleSrcSuggestTarget(request *stratumRequest) {

}

//
//
//
func (svs *StratumV1Struct) handleDstGetVersion(index int, request *stratumRequest) {

}

//
// handleDstNotify()
// passes new work from the destination pool to the source miner
//
func (svs *StratumV1Struct) handleDstNotify(index int, request *stratumNotice) {

	// is index the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src

	defroutidx, e := svs.protocol.GetDefaultRoute()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" GetDefaultRouter returned error:%s", e)
	}

	// This is the default route
	if defroutidx == index {

	}

	// Ok, what do we do now?

}

//
//
//
func (svs *StratumV1Struct) handleDstPing(index int, request *stratumRequest) {

}

//
//
//
func (svs *StratumV1Struct) handleDstReconnect(index int, request *stratumRequest) {

}

//
//
//
func (svs *StratumV1Struct) handleDstShowMessage(index int, request *stratumRequest) {

}

//
//
//
func (svs *StratumV1Struct) handleDstSetDifficulty(index int, request *stratumNotice) {

	// is index the current default destination?
	// If not, store the notify?
	// If so, pass it to the Src

	defroutidx, e := svs.protocol.GetDefaultRoute()
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" GetDefaultRouter returned error:%s", e)
	}

	// This is the default route
	if defroutidx == index {

	}

}

//
//
//
func (svs *StratumV1Struct) handleDstSetExtranonce(index int, request *stratumRequest) {

}

//
//
//
func (svs *StratumV1Struct) handleDstSetGoal(index int, request *stratumRequest) {

}
