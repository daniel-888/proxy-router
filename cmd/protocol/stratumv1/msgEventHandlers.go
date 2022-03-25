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
