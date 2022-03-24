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

// type newStratumV1Func func(*simple.SimpleStruct)

// type newStratumV1Struct struct {
// 	funcptr newStratumV1Func
// }

type SrcState string
type DstState string

//
// New->Subscribed->Authorized->??
//
const SrcStateNew SrcState = "stateNew"
const SrcStateSubscribed SrcState = "stateSubscribed"
const SrcStateAuthorized SrcState = "stateAuthorized"
const SrcStateError SrcState = "stateError"

//
// New->Subscribed->Authorized->??
//
const DstStateNew DstState = "stateNew"
const DstStateOpen DstState = "stateOpen"
const DstStateSubscribing DstState = "stateSubscribing"
const DstStateAuthorizing DstState = "stateAuthorizing"
const DstStateError DstState = "stateError"

type StratumV1ListenStruct struct {
	protocollisten *protocol.ProtocolListenStruct
}

type StratumV1Struct struct {
	ctx                 context.Context
	cancel              func()
	protocol            *protocol.ProtocolStruct
	minerRec            *msgbus.Miner
	srcSubscribeRequest *stratumRequest // Copy of recieved Subscribe Request from Source
	srcAuthRequest      *stratumRequest // Copy of recieved Authorize Request from Source
	srcState            SrcState
	dstState            map[simple.ConnUniqueID]DstState

	// Add in stratum state information here
}

//
//
//
// func NewListener(ctx context.Context, src net.Addr, dst net.Addr, proto ...*newStratumV1Struct) (sls *StratumV1ListenStruct, e error) {
func NewListener(ctx context.Context, src net.Addr, dst net.Addr) (sls *StratumV1ListenStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	var cs *contextlib.ContextStruct = contextlib.GetContextStruct(ctx)

	if cs == nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" ContextStruct not defined")
	}
	if nil == cs.GetLog() {
		contextlib.Logf(ctx, contextlib.LevelWarn, lumerinlib.FileLineFunc()+" ContextStruct no Logger * defined")
	}
	if nil == cs.GetMsgBus() {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" ContextStruct no MsgBus * defined")
	}

	cs.SetSrc(src)
	cs.SetDst(dst)

	protocollisten, err := protocol.NewListen(ctx)
	if err != nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" NewListen returned error:%s", e)
	}

	sls = &StratumV1ListenStruct{
		protocollisten: protocollisten,
	}

	return sls, e
}

//
//
//
func (s *StratumV1ListenStruct) goListenAccept() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	// defer close(s.accept)

	protocolStructChan := s.protocollisten.GetAccept()
FORLOOP:
	for {
		select {
		case <-s.Ctx().Done():
			contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" context canceled")
			break FORLOOP
		case l := <-protocolStructChan:
			ss := NewStratumStruct(s.Ctx(), l)
			ss.Run()
		}
	}

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Exiting...")

}

//
//
//
func NewStratumStruct(ctx context.Context, l *protocol.ProtocolStruct) (n *StratumV1Struct) {
	ctx, cancel := context.WithCancel(ctx)
	ds := make(map[simple.ConnUniqueID]DstState)
	n = &StratumV1Struct{
		ctx:                 ctx,
		cancel:              cancel,
		protocol:            l,
		minerRec:            nil,
		srcSubscribeRequest: &stratumRequest{},
		srcAuthRequest:      &stratumRequest{},
		srcState:            SrcStateNew,
		dstState:            ds,
	}
	return n
}

//
//
//
func (sls *StratumV1ListenStruct) Run() {

	contextlib.Logf(sls.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	sls.protocollisten.Run()
	go sls.goListenAccept()
}

//
//
//
func (s *StratumV1ListenStruct) Ctx() context.Context {
	if s == nil {
		panic(lumerinlib.FileLineFunc() + " nil pointer")
	}
	contextlib.Logf(s.protocollisten.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")
	return s.protocollisten.Ctx()
}

//
//
//
func (s *StratumV1ListenStruct) Cancel() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	s.protocollisten.Cancel()
}

//
// Run() inialize the stratum running struct
//
func (s *StratumV1Struct) Run() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called ")

	cs := contextlib.GetContextStruct(s.Ctx())
	if cs == nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Struct not in CTX")
	}

	dst := cs.GetDst()
	if dst == nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Struct DST not defined")
	}

	s.protocol.Run()
	go s.goEvent()
}

// ---------------------------------------------------------------------
//  StratumV1Struct
//

//
// goEvent()
//
func (s *StratumV1Struct) goEvent() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	simplechan := s.protocol.GetSimpleEventChan()

	contextlib.Logf(s.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" Simplechan:%v", simplechan)

	for event := range simplechan {

		if event == nil {
			s.Cancel()
			contextlib.Logf(s.Ctx(), contextlib.LevelFatal, lumerinlib.FileLineFunc()+" event:%v", event)
		}

		contextlib.Logf(s.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" event:%v", event)
		e := s.eventHandler(event)

		if e != nil {
			contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" eventHandler() returned error:%s ", e)
		}
	}
}

//
// Ctx()
// returns the StratumV1Struct context pointer
//
func (s *StratumV1Struct) Ctx() context.Context {
	if s == nil {
		panic(lumerinlib.FileLineFunc() + "StratumV1Struct is nil")
	}
	if s.protocol == nil {
		panic(lumerinlib.FileLineFunc() + "StratumV1Struct.protocol is nil")
	}
	return s.protocol.Ctx()
}

//
// Cancels the StratumV1Struct instance
//
func (s *StratumV1Struct) Cancel() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	s.protocol.Cancel()
}

//
//
//
func (s *StratumV1Struct) SetDstStateUid(uid simple.ConnUniqueID, state DstState) {
	s.dstState[uid] = state
}

//
//
//
func (s *StratumV1Struct) GetDstStateUid(uid simple.ConnUniqueID) (state DstState) {
	return s.dstState[uid]
}

//
//
//
func (s *StratumV1Struct) SetSrcState(state SrcState) {
	s.srcState = state
}

//
//
//
func (s *StratumV1Struct) GetSrcState() (state SrcState) {
	return s.srcState
}

//
//
// This takes the SimpleEvent and dispatches it to the appropriate handeler, updaing the
// StratumV1Struct state along the way.
// The event hander is expected to be single threaded
//
// Event Handler
func (svs *StratumV1Struct) eventHandler(event *simple.SimpleEvent) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch event.EventType {
	case simple.NoEvent:
		return

	case simple.MsgBusEvent:
		msg := event.MsgBusEvent
		svs.decodeMsgBusEvent(msg)
		return

	case simple.ConnReadEvent:
		scre := event.ConnReadEvent
		e = svs.handleConnReadEvent(scre)
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" handleConnReadEvent() returned error:%s", e)
		}
		return

	case simple.ConnOpenEvent:
		scoe := event.ConnOpenEvent
		e = svs.handleConnOpenEvent(scoe)
		if e != nil {
			contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" handleConnOpenEvent() returned error:%s", e)
		}
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
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Default Reached: Event Type:%s", string(event.EventType))
		e = fmt.Errorf(" Default Reached: Event Type:%s", string(event.EventType))
	}

	return e
}

//
//
//
func (svs *StratumV1Struct) decodeMsgBusEvent(event *simple.SimpleMsgBusEvent) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

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
		lumerinlib.PanicHere(fmt.Sprintf(lumerinlib.FileLineFunc()+" Default Reached: Event Type:%s", string(event.EventType)))
	}
}
