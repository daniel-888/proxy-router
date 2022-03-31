package stratumv1

import (
	"context"
	"errors"
	"fmt"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

var ErrBadSrcState = errors.New("StratumV1: Bad Src State")
var ErrSrcReqNotSupported = errors.New("StratumV1: SRC Request Not Supported")
var ErrDstReqNotSupported = errors.New("StratumV1: DST Request Not Supported")

type SrcState string
type DstState string

//
// New->Subscribed->Authorized->??
//
const SrcStateNew SrcState = "stateNew"               // Freshly created Connection
const SrcStateSubscribed SrcState = "stateSubscribed" // Recieve Subscribe
const SrcStateAuthorized SrcState = "stateAuthorized" // Recieve Authorize
const SrcStateRunning SrcState = "stateRunning"       // Sent set_difficulty or work notice
const SrcStateError SrcState = "stateError"

//
// New->Subscribed->Authorized->Running
//
const DstStateNew DstState = "stateNew"
const DstStateOpen DstState = "stateOpen"
const DstStateSubscribing DstState = "stateSubscribing" // Sent Subscribe
const DstStateAuthorizing DstState = "stateAuthorizing" // Recieved Sub-response and Sent Authorize
const DstStateRunning DstState = "stateRunning"         // Recieved Auth-response (there should only be one dst connection running at any time)
const DstStateStandBy DstState = "stateStandBy"         // Inactive, not the focus of the Src connection
const DstStateRedialing DstState = "stateRedialing"     // Inactive, in the process of redialing
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
	dstDest             map[simple.ConnUniqueID]*msgbus.Dest
	switchToDestID      msgbus.DestID
	// defaultUID          simple.ConnUniqueID -- stored in protocol layer

	// Add in stratum state information here
}

var MinerCountChan chan int

//
// init()
// initializes the DstCounter
//
func init() {
	MinerCountChan = make(chan int, 5)
	go goMinerCounter(MinerCountChan)
}

//
// goDstCounter()
// Generates a UniqueID for the destination handles
//
func goMinerCounter(c chan int) {
	counter := 10000
	for {
		c <- counter
		counter += 1
	}
}

//
//
//
func NewListener(ctx context.Context, src net.Addr, dest *msgbus.Dest) (sls *StratumV1ListenStruct, e error) {

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
	cs.SetDest(dest)

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
			ss := NewStratumV1Struct(s.Ctx(), l)
			ss.Run()
		}
	}

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Exiting...")

}

//
//
//
func NewStratumV1Struct(ctx context.Context, l *protocol.ProtocolStruct) (n *StratumV1Struct) {
	ctx, cancel := context.WithCancel(ctx)
	ds := make(map[simple.ConnUniqueID]DstState)
	dd := make(map[simple.ConnUniqueID]*msgbus.Dest)
	id := fmt.Sprintf("MinerID:%d", <-MinerCountChan)
	defdest := contextlib.GetContextStruct(ctx).GetDest()
	if defdest == nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetDest() return nil")
	}
	miner := &msgbus.Miner{
		ID:                      msgbus.MinerID(id),
		Name:                    "",
		IP:                      "",
		MAC:                     "",
		State:                   "",
		Contract:                "",
		Dest:                    defdest.ID,
		InitialMeasuredHashRate: 0,
		CurrentHashRate:         0,
		CsMinerHandlerIgnore:    false,
	}

	n = &StratumV1Struct{
		ctx:                 ctx,
		cancel:              cancel,
		protocol:            l,
		minerRec:            miner,
		srcSubscribeRequest: &stratumRequest{},
		srcAuthRequest:      &stratumRequest{},
		srcState:            SrcStateNew,
		dstState:            ds,
		dstDest:             dd,
		switchToDestID:      "",
	}

	n.newMinerRecordPub(miner)

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

// ---------------------------------------------------------------------
//  StratumV1Struct Functions
// ---------------------------------------------------------------------

//
// Run() inialize the stratum running struct
//
func (s *StratumV1Struct) Run() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called ")

	cs := contextlib.GetContextStruct(s.Ctx())
	if cs == nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Struct not in CTX")
	}

	dst := cs.GetDest()
	if dst == nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Struct DST not defined")
	}

	s.protocol.Run()
	go s.goEvent()

	s.openDefaultConnection()

}

//
// openDefaultConnection()
// Start the event process to open up a defaut connection
//
func (s *StratumV1Struct) openDefaultConnection() (e error) {

	dest := contextlib.GetContextStruct(s.Ctx()).GetDest()
	if dest == nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelFatal, lumerinlib.FileLineFunc()+" Context Struct DST not defined")
		return errors.New("default Dest not defined")
	}

	s.switchToDestID = dest.ID
	s.protocol.Get(simple.DestMsg, simple.IDString(dest.ID))

	return nil
}

//
// goEvent()
//
func (s *StratumV1Struct) goEvent() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	simplechan := s.protocol.GetSimpleEventChan()

	contextlib.Logf(s.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" Simplechan:%v", simplechan)

	for event := range simplechan {

		// closed connection
		if event == nil {
			contextlib.Logf(s.Ctx(), contextlib.LevelFatal, lumerinlib.FileLineFunc()+"[Closing] event:%v", event)
			break
		}

		e := s.eventHandler(event)

		if e != nil {
			contextlib.Logf(s.Ctx(), contextlib.LevelFatal, lumerinlib.FileLineFunc()+"[Closing] eventHandler() returned error:%s", e)
			break
		}
	}
	s.Cancel()
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

	s.protocol.Unpub(simple.MinerMsg, simple.IDString(s.minerRec.ID))

	s.protocol.Cancel()

	if s.cancel == nil {
		contextlib.Logf(s.ctx, contextlib.LevelError, fmt.Sprint(lumerinlib.FileLineFunc()+" cancel func is nil, struct:%v", s))
		return
	}

	s.cancel()
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
func (s *StratumV1Struct) SetDsDestUid(uid simple.ConnUniqueID, dest *msgbus.Dest) {
	s.dstDest[uid] = dest
}

//
//
//
func (s *StratumV1Struct) GetDstDestUid(uid simple.ConnUniqueID) (dest *msgbus.Dest) {
	return s.dstDest[uid]
}

//
//
//
func (s *StratumV1Struct) GetDstUIDDestID(id msgbus.DestID) (uid simple.ConnUniqueID) {
	uid = -1
	for u, v := range s.dstDest {
		if v.ID == id {
			uid = u
			break
		}
	}
	return uid
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
//
func (s *StratumV1Struct) newMinerRecordPub(m *msgbus.Miner) {

	mcopy := *m
	_, e := s.protocol.Pub(simple.MinerMsg, simple.IDString(m.ID), mcopy)
	if e != nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Miner Pub() error:%s ", e)
	}
}

//
// swtichDest()
// Check if switchtoDestID is set
// Check if it is set to the current dest
// Switch active dest if not
//
func (s *StratumV1Struct) switchDest() {

	if s.switchToDestID == "" {
		return
	}

	currentUID := s.protocol.GetDefaultRouteUID()

	// is the next dest the current dest?
	if currentUID >= 0 && s.switchToDestID == s.dstDest[currentUID].ID {
		s.switchToDestID = ""
		return
	}

	newUID := s.GetDstUIDDestID(s.switchToDestID)
	if s.dstState[newUID] == DstStateStandBy {

		if currentUID >= 0 {
			s.dstState[currentUID] = DstStateStandBy
		}
		s.dstState[newUID] = DstStateRunning
		s.protocol.SetDefaultRouteUID(newUID)

		// Reset the switch to state
		s.switchToDestID = ""

		// Send set difficulty notice to SRC -  to reset it
		var params = make([]interface{}, 1)
		params[0] = 1.0
		n := &stratumNotice{
			ID:     nil,
			Method: string(SERVER_MINING_SET_DIFFICULTY),
			Params: params,
		}

		msg, e := n.createNoticeSetDifficultyMsg()
		if e != nil {
			contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createNoticeSetDifficulty() error: %s", e)
		}

		count, e := s.protocol.WriteSrc(msg)
		if e != nil {
			contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" WriteSrc() error: %s", e)
		}

		if count != len(msg) {
			contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" WriteSrc() count not equal to msg %d != %d", count, len(msg))

		}
	} else {
		contextlib.Logf(s.Ctx(), contextlib.LevelWarn, lumerinlib.FileLineFunc()+" next dest not in standby mode ")
	}

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
		svs.handleMsgBusEvent(msg)
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
