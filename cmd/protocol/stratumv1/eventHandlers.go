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

}

//
// handleConnEOFEvent()
// Close out the connection index
// Could reopen the connection, or perform some error handling
//
func (svs *StratumV1Struct) handleConnEOFEvent(event *simple.SimpleEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleConnErrorEvent()
// Perform some error handling
//
func (svs *StratumV1Struct) handleConnErrorEvent(event *simple.SimpleEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}

//
// handleErrorEvent()
// Perform Error handling
//
func (svs *StratumV1Struct) handleErrorEvent(event *simple.SimpleEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" Called")

}
