package stratumv1

import (
	"context"
	"fmt"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

//
// Need a series of access functions
//
//

type StratumV1ListenStruct struct {
	protocollisten *protocol.ProtocolListenStruct
}

type StratumV1Struct struct {
	protocol *protocol.ProtocolStruct
	// Add in stratum state information here
}

//
//
//
func New(ctx context.Context, mb *msgbus.PubSub, src net.Addr, dst net.Addr) (s *StratumV1ListenStruct, e error) {

	// Validate src and dst here

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	cs := &contextlib.ContextStruct{}
	cs.SetMsgBus(mb)
	cs.SetSrc(src)
	cs.SetDst(dst)
	//
	// This is the only place that SetProtocol is called
	cs.SetProtocol(newStratumV1Func)

	ctx = context.WithValue(ctx, contextlib.ContextKey, cs)

	protocollisten, err := protocol.NewListen(ctx)
	if err != nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" NewListen returned error:%s", e)
	}

	s = &StratumV1ListenStruct{
		protocollisten: protocollisten,
	}

	return s, e
}

//
//
//
func (s *StratumV1Struct) goEvent() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	for event := range s.protocol.Event() {
		s.eventHandler(event)
	}
}

//
//
//
func (s *StratumV1ListenStruct) Run() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	s.protocollisten.Run()
}

//
//
//
func (s *StratumV1ListenStruct) Ctx() context.Context {
	return s.protocollisten.Ctx()
}

//
//
//
func (s *StratumV1ListenStruct) Cancel() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	s.protocollisten.Cancel()
}

//
// newProtoFunc() is called by the simple layer for Accept() connections.
// The system here will loop on the event channel, and handle the events one at a time
//
// SIMPL defined this function as passing in a SimpeStruct abd retuning a chan for SimpleEvents
//
func newStratumV1Func(ss *simple.SimpleStruct) chan *simple.SimpleEvent {

	contextlib.Logf(ss.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	i := ss.Ctx().Value(contextlib.ContextKey)
	cs, ok := i.(contextlib.ContextStruct)
	if !ok {
		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Context Struct not in CTX")
	}

	dst := cs.GetDst()
	if dst == nil {
		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Context Struct DST not defined")
	}

	// inialize a new ProtocolStruct to gain access to the standard protocol functions
	// The default Dst should be opened when this returns
	pls, err := protocol.NewProtocol(ss)
	if err != nil {
		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Create NewProtocol() failed: %s", err)
	}

	svs := &StratumV1Struct{
		protocol: pls,
		// Fill in other state information here
	}

	// Launch the event handler
	go svs.goEvent()

	// return the event handler channel to the caller (the simple layer accept() function )
	return svs.protocol.Event()
}

// ---------------------------------------------------------------------
//  StratumV1Struct
//

//
// returns the StratumV1Struct context pointer
//
func (s *StratumV1Struct) Ctx() context.Context {
	return s.protocol.Ctx()
}

//
// Cancels the StratumV1Struct instance
//
func (s *StratumV1Struct) Cancel() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	s.protocol.Cancel()
}

//
//
// This takes the SimpleEvent and dispatches it to the appropriate handeler, updaing the
// StratumV1Struct state along the way.
// The event hander is expected to be single threaded
//
// Event Handler
func (svs *StratumV1Struct) eventHandler(event *simple.SimpleEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	switch event.EventType {
	case simple.NoEvent:
		return

	case simple.MsgUpdateEvent:
		fallthrough
	case simple.MsgDeleteEvent:
		fallthrough
	case simple.MsgGetEvent:
		fallthrough
	case simple.MsgGetIndexEvent:
		fallthrough
	case simple.MsgSearchEvent:
		fallthrough
	case simple.MsgSearchIndexEvent:
		fallthrough
	case simple.MsgPublishEvent:
		fallthrough
	case simple.MsgUnpublishEvent:
		fallthrough
	case simple.MsgSubscribedEvent:
		fallthrough
	case simple.MsgUnsubscribedEvent:
		fallthrough
	case simple.MsgRemovedEvent:
		msg, ok := event.Data.(msgbus.Event)
		if !ok {
			contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Event Data wrong Type:%t", event.Data)
		}
		// Error checking here  ev == megbus.Event
		svs.decodeMsgBusEvent(msg)
		return

	case simple.ConnReadEvent:
		// Error checking here event == connection event
		svs.handleConnReadEvent(event)
		return

	case simple.ConnEOFEvent:
		// Error checking here event == connection event
		svs.handleConnEOFEvent(event)
		return

	case simple.ConnErrorEvent:
		// Error checking here event == connection event
		svs.handleConnErrorEvent(event)
		return

	case simple.ErrorEvent:
		// Error checking here event == Error event
		svs.handleErrorEvent(event)
		return

	default:
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" Default Reached: Event Type:%s", string(event.EventType))
	}

}

//
//
//
func (svs *StratumV1Struct) decodeMsgBusEvent(event msgbus.Event) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	switch event.EventType {
	case msgbus.NoEvent:
		fmt.Printf(lumerinlib.Funcname() + " NoEvent received, returning\n")
		return
	case msgbus.UpdateEvent:
		svs.handleMsgUpdateEvent(event)
		return
	case msgbus.DeleteEvent:
		svs.handleMsgDeleteEvent(event)
		return
	case msgbus.GetEvent:
		svs.handleMsgGetEvent(event)
		return
	case msgbus.GetIndexEvent:
		svs.handleMsgIndexEvent(event)
		return
	case msgbus.SearchEvent:
		svs.handleMsgSearchEvent(event)
		return
	case msgbus.SearchIndexEvent:
		svs.handleMsgSearchIndexEvent(event)
		return
	case msgbus.PublishEvent:
		svs.handleMsgPublishEvent(event)
		return
	case msgbus.UnpublishEvent:
		svs.handleMsgUnpublishEvent(event)
		return
	case msgbus.SubscribedEvent:
		svs.handleMsgSubscribedEvent(event)
		return
	case msgbus.UnsubscribedEvent:
		svs.handleMsgUnsubscribedEvent(event)
		return
	case msgbus.RemovedEvent:
		svs.handleMsgRemovedEvent(event)
		return

	default:
		lumerinlib.PanicHere(fmt.Sprintf(lumerinlib.FileLine()+" Default Reached: Event Type:%s", string(event.EventType)))
	}
}
