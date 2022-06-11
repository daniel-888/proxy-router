package stratumv1

import (
	"context"
	"errors"
	"fmt"
	"net"
<<<<<<< HEAD
=======
	"strconv"
	"strings"
>>>>>>> pr-009

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

var ErrBadSrcState = errors.New("StratumV1: Bad Src State")
var ErrSrcReqNotSupported = errors.New("StratumV1: SRC Request Not Supported")
var ErrDstReqNotSupported = errors.New("StratumV1: DST Request Not Supported")
var ErrMaxRedialExceeded = errors.New("StratumV1: DST Maximum number of redials attempted")

type SrcState string
type DstState string

<<<<<<< HEAD
//
// New->Subscribed->Authorized->??
=======
type StratumConnectionScheduler string

const OnDemand StratumConnectionScheduler = "OnDemand"
const OnSubmit StratumConnectionScheduler = "OnSubmit"

//
// New->Subscribed->Authorized->Running??
>>>>>>> pr-009
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
const DstStateStandBy DstState = "stateStandBy"         // Active, but not the focus of the Src connection
const DstStateDialing DstState = "stateDialing"         // Inactive, initiating a connection
const DstStateRedialing DstState = "stateRedialing"     // Inactive, reinitiating a connection
const DstStateError DstState = "stateError"
const DstStateClosed DstState = "stateClosed"
<<<<<<< HEAD

const MaxRedials int = 2

type StratumV1ListenStruct struct {
	protocollisten *protocol.ProtocolListenStruct
=======
const DstStateNotFound DstState = "stateNotFound"

const MaxRedials int = 5

type StratumV1ListenStruct struct {
	protocollisten *protocol.ProtocolListenStruct
	scheduler      StratumConnectionScheduler
>>>>>>> pr-009
}

type StratumV1Struct struct {
	ctx                 context.Context
	cancel              func()
	protocol            *protocol.ProtocolStruct
	minerRec            *msgbus.Miner
<<<<<<< HEAD
	srcSubscribeRequest *stratumRequest // Copy of recieved Subscribe Request from Source
	srcAuthRequest      *stratumRequest // Copy of recieved Authorize Request from Source
	srcConfigure        *stratumRequest // Copy of recieved Configure Request from Source
	srcExtranonce       *stratumRequest // Copy of recieved Configure Request from Source
=======
	scheduler           StratumConnectionScheduler
	srcSubscribeRequest *stratumRequest // Copy of recieved Subscribe Request from Source
	srcAuthRequest      *stratumRequest // Copy of recieved Authorize Request from Source
	srcConfigure        *stratumRequest // Copy of recieved Configure Request from Source
	srcExtranonce       *stratumRequest // Copy of recieved Extranonce Request from Source
>>>>>>> pr-009
	srcState            SrcState
	dstState            map[simple.ConnUniqueID]DstState
	dstDest             map[simple.ConnUniqueID]*msgbus.Dest
	dstReDialCount      map[simple.ConnUniqueID]int
	dstExtranonce       map[simple.ConnUniqueID]string
	dstExtranonce2size  map[simple.ConnUniqueID]int
<<<<<<< HEAD
	dstDiff             map[simple.ConnUniqueID]int
	dstVersionMask      map[simple.ConnUniqueID]string
	dstLastMiningNotice map[simple.ConnUniqueID]*stratumNotice
	dstLastReqNotice    map[simple.ConnUniqueID]*stratumRequest
=======
	dstVersionMask      map[simple.ConnUniqueID]string
	dstLastSetDiff      map[simple.ConnUniqueID]int
	dstLastMiningNotice map[simple.ConnUniqueID]*stratumNotice
	dstLastReqNotify    map[simple.ConnUniqueID]*stratumRequest
>>>>>>> pr-009
	switchToDestID      msgbus.DestID

	// Add in stratum state information here
}

var MinerCountChan chan int

//
// init()
// initializes the DstCounter
//
func init() {
	MinerCountChan = make(chan int, 5)
<<<<<<< HEAD
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
=======
	lumerinlib.RunGoCounter(MinerCountChan)
>>>>>>> pr-009
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

	protocollisten, e := protocol.NewListen(ctx)
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" NewListen returned error:%s", e)
	} else {
		sls = &StratumV1ListenStruct{
			protocollisten: protocollisten,
<<<<<<< HEAD
=======
			scheduler:      OnDemand, // OnDemand is the Default
>>>>>>> pr-009
		}
	}

	return sls, e
}

//
//
//
<<<<<<< HEAD
=======
func (s *StratumV1ListenStruct) SetScheduler(scheduler StratumConnectionScheduler) {
	s.scheduler = scheduler
}

//
//
//
func (s *StratumV1ListenStruct) GetScheduler() (scheduler StratumConnectionScheduler) {
	scheduler = s.scheduler
	return scheduler
}

//
//
//
>>>>>>> pr-009
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
		case ps := <-protocolStructChan:
<<<<<<< HEAD
			ss := NewStratumV1Struct(s.Ctx(), ps)
			ss.Run()
=======
			ss := NewStratumV1Struct(s.Ctx(), ps, s.scheduler)
			if ss != nil {
				ss.Run()
			}
>>>>>>> pr-009
		}
	}

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Exiting...")

}

//
<<<<<<< HEAD
//
=======
// Used for testing
>>>>>>> pr-009
//
func (s *StratumV1ListenStruct) goListenAcceptOnce() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	protocolStructChan := s.protocollisten.GetAccept()
	select {
	case <-s.Ctx().Done():
		contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" context canceled")
	case ps := <-protocolStructChan:
<<<<<<< HEAD
		ss := NewStratumV1Struct(s.Ctx(), ps)
		ss.Run()
=======
		ss := NewStratumV1Struct(s.Ctx(), ps, s.scheduler)
		if ss != nil {
			ss.Run()
		}
>>>>>>> pr-009
	}

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Exiting...")

}

//
//
//
func (sls *StratumV1ListenStruct) Run() {

	//	contextlib.Logf(sls.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	sls.protocollisten.Run()
	go sls.goListenAccept()
}

//
//
//
func (sls *StratumV1ListenStruct) RunOnce() {

	//	contextlib.Logf(sls.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	sls.protocollisten.Run()
	go sls.goListenAcceptOnce()
}

//
//
//
func (s *StratumV1ListenStruct) Ctx() context.Context {
	if s == nil {
		panic(lumerinlib.FileLineFunc() + " nil pointer")
	}
	//contextlib.Logf(s.protocollisten.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")
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
//
//
<<<<<<< HEAD
func NewStratumV1Struct(ctx context.Context, ps *protocol.ProtocolStruct) (n *StratumV1Struct) {
=======
func NewStratumV1Struct(ctx context.Context, ps *protocol.ProtocolStruct, scheduler StratumConnectionScheduler) (n *StratumV1Struct) {
>>>>>>> pr-009
	ctx, cancel := context.WithCancel(ctx)
	ds := make(map[simple.ConnUniqueID]DstState)
	dd := make(map[simple.ConnUniqueID]*msgbus.Dest)
	rd := make(map[simple.ConnUniqueID]int)
	de := make(map[simple.ConnUniqueID]string)
	de2 := make(map[simple.ConnUniqueID]int)
<<<<<<< HEAD
	di := make(map[simple.ConnUniqueID]int)
=======
	lsd := make(map[simple.ConnUniqueID]int)
>>>>>>> pr-009
	vm := make(map[simple.ConnUniqueID]string)
	lmn := make(map[simple.ConnUniqueID]*stratumNotice)
	lrn := make(map[simple.ConnUniqueID]*stratumRequest)
	id := fmt.Sprintf("MinerID:%d", <-MinerCountChan)
	defdest := contextlib.GetContextStruct(ctx).GetDest()
	if defdest == nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetDest() return nil")
	}
<<<<<<< HEAD
	miner := &msgbus.Miner{
		ID:                      msgbus.MinerID(id),
		Name:                    "",
		IP:                      "",
		MAC:                     "",
		State:                   msgbus.OnlineState,
		Contract:                "",
		Dest:                    defdest.ID,
		InitialMeasuredHashRate: 0,
		CurrentHashRate:         0,
		CsMinerHandlerIgnore:    false,
=======

	addr, e := ps.GetSrcRemoteAddr()
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" GetSrcRemoteAddr() error:%s", e)
		return nil
	}

	addrstr := strings.Split(addr.String(), ":")
	ip := addrstr[0]
	port, e := strconv.Atoi(addrstr[1])
	if e != nil {
		port = 0
		contextlib.Logf(ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" strconv.Atoi() for str:%s error:%s", addrstr, e)
	}

	miner := &msgbus.Miner{
		ID:                      msgbus.MinerID(id),
		Name:                    "",
		IP:                      ip,
		Port:                    port,
		MAC:                     "", // Future... maybe
		State:                   msgbus.OnlineState,
		Contracts:               make(map[msgbus.ContractID]float64),
		Dest:                    defdest.ID,
		InitialMeasuredHashRate: 0,
		CurrentHashRate:         0,
>>>>>>> pr-009
	}

	n = &StratumV1Struct{
		ctx:                 ctx,
		cancel:              cancel,
		protocol:            ps,
		minerRec:            miner,
<<<<<<< HEAD
=======
		scheduler:           scheduler,
>>>>>>> pr-009
		srcSubscribeRequest: nil,
		srcAuthRequest:      nil,
		srcConfigure:        nil,
		srcExtranonce:       nil,
		srcState:            SrcStateNew,
		dstState:            ds,
		dstDest:             dd,
		dstReDialCount:      rd,
		dstExtranonce:       de,
		dstExtranonce2size:  de2,
<<<<<<< HEAD
		dstDiff:             di,
		dstVersionMask:      vm,
		dstLastMiningNotice: lmn,
		dstLastReqNotice:    lrn,
=======
		dstLastSetDiff:      lsd,
		dstVersionMask:      vm,
		dstLastMiningNotice: lmn,
		dstLastReqNotify:    lrn,
>>>>>>> pr-009
		switchToDestID:      "",
	}

	return n
}

//
<<<<<<< HEAD
=======
//
//
func (svs *StratumV1Struct) GetScheduler() (scheduler StratumConnectionScheduler) {
	scheduler = svs.scheduler
	return scheduler
}

//
>>>>>>> pr-009
// Run() inialize the stratum running struct
//
func (s *StratumV1Struct) Run() {

	//	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called ")

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

	// Moved to after recievin Subscribe
	// s.newMinerRecordPub()
	// s.openDefaultConnection()

}

//
// openDefaultConnection()
// Start the event process to open up a defaut connection
// Send GetDest()
// The GetDest Event will call the AsyncDial() for the dest
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

	//	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	simplechan := s.protocol.GetSimpleEventChan()

	// contextlib.Logf(s.Ctx(), contextlib.LevelInfo, lumerinlib.FileLineFunc()+" Simplechan:%v", simplechan)

	for event := range simplechan {

		// closed connection
		if event == nil {
<<<<<<< HEAD
			contextlib.Logf(s.Ctx(), contextlib.LevelFatal, lumerinlib.FileLineFunc()+"[Closing] event:%v", event)
=======
			contextlib.Logf(s.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+"[Closing] event:%v", event)
>>>>>>> pr-009
			break
		}

		e := s.eventHandler(event)

		if e != nil {
<<<<<<< HEAD
			contextlib.Logf(s.Ctx(), contextlib.LevelFatal, lumerinlib.FileLineFunc()+"[Closing] eventHandler() returned error:%s", e)
=======
			contextlib.Logf(s.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+"[Closing] eventHandler() returned error:%s", e)
>>>>>>> pr-009
			break
		}
	}
	s.Close()
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
//
//
func (s *StratumV1Struct) Close() {

	// Orderly shutdown of the system here

	s.protocol.Close()
	s.Cancel()

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
<<<<<<< HEAD
//
//
func (s *StratumV1Struct) CloseUid(uid simple.ConnUniqueID) {
	s.dstState[uid] = DstStateClosed
=======
// CloseUid()
//
// should the uid entry be removed at this point?
//
func (s *StratumV1Struct) CloseUid(uid simple.ConnUniqueID) {

	contextlib.Logf(s.ctx, contextlib.LevelInfo, fmt.Sprint(lumerinlib.FileLineFunc()+" UID:%d", uid))

	s.SetDstStateUid(uid, DstStateClosed)
>>>>>>> pr-009
	s.dstDest[uid] = nil
	s.protocol.CloseDst(uid)
}

//
//
//
func (s *StratumV1Struct) DstRedialUid(uid simple.ConnUniqueID) (e error) {
<<<<<<< HEAD
	s.dstState[uid] = DstStateRedialing
=======

	contextlib.Logf(s.ctx, contextlib.LevelInfo, fmt.Sprint(lumerinlib.FileLineFunc()+" UID:%d", uid))

	s.SetDstStateUid(uid, DstStateRedialing)
>>>>>>> pr-009
	s.dstReDialCount[uid]++
	if s.dstReDialCount[uid] > MaxRedials {
		s.CloseUid(uid)
		return ErrMaxRedialExceeded
	} else {
		s.protocol.AsyncReDial(uid)
	}
	return nil
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
<<<<<<< HEAD
	return s.dstState[uid]
=======
	v, ok := s.dstState[uid]
	if !ok {
		return DstStateNotFound
	}
	return v
>>>>>>> pr-009
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
func (s *StratumV1Struct) GetDstUsernameUid(uid simple.ConnUniqueID) (username string) {
	dest := s.dstDest[uid]
	username = dest.Username()
	return username
}

//
//
//
func (s *StratumV1Struct) GetDstPasswordUid(uid simple.ConnUniqueID) (password string) {
	dest := s.dstDest[uid]
	password = dest.Password()
	return password
}

//
//
//
func (s *StratumV1Struct) GetDstUIDDestID(id msgbus.DestID) (uid simple.ConnUniqueID) {
	uid = -1

	if s == nil {
		contextlib.Logf(s.ctx, contextlib.LevelError, fmt.Sprint(lumerinlib.FileLineFunc()+" stratum struct is nil"))
		return uid
	}

	if s.dstDest == nil {
		contextlib.Logf(s.ctx, contextlib.LevelError, fmt.Sprint(lumerinlib.FileLineFunc()+" stratum struct is nil"))
		return uid
	}

	if id == "" {
		contextlib.Logf(s.ctx, contextlib.LevelPanic, fmt.Sprint(lumerinlib.FileLineFunc()+" id is blank"))
	}

<<<<<<< HEAD
	if s.dstDest != nil {
		for u, v := range s.dstDest {
			if v == nil {
				contextlib.Logf(s.ctx, contextlib.LevelError, fmt.Sprint(lumerinlib.FileLineFunc()+" v is nil"))
				continue
			}

			if v.ID == id {
				uid = u
				break
			}
		}
	}
=======
	for u, v := range s.dstDest {
		if v == nil {
			contextlib.Logf(s.ctx, contextlib.LevelError, fmt.Sprint(lumerinlib.FileLineFunc()+" v is nil"))
			continue
		}

		if v.ID == id {
			uid = u
			break
		}
	}

>>>>>>> pr-009
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
func (s *StratumV1Struct) newMinerRecordPub() {

	miner := *s.minerRec
	rid, e := s.protocol.Pub(simple.MinerMsg, simple.IDString(miner.ID), &miner)
	if e != nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Miner Pub() error:%s RID:%d", e, rid)
	}
}

//
// swtichDest()
// Check if switchtoDestID is set
// Check if it is set to the current dest
// Switch active dest if not
//
func (s *StratumV1Struct) switchDest() {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called ")

	if s.switchToDestID == "" {
<<<<<<< HEAD
=======
		contextlib.Logf(s.Ctx(), contextlib.LevelInfo, fmt.Sprintf(lumerinlib.FileLineFunc()+" called with no designated next dest, ignoring"))
>>>>>>> pr-009
		return
	}

	currentUID, _ := s.protocol.GetDefaultRouteUID()
<<<<<<< HEAD

	//
	// Is the current Route the same as the new route?
	//
	if currentUID >= 0 {
		v, ok := s.dstDest[currentUID]

		if v == nil {
			contextlib.Logf(s.Ctx(), contextlib.LevelPanic, fmt.Sprintf(lumerinlib.FileLineFunc()+" dstDest[%d] ", currentUID))
		}

		if ok {
			if s.switchToDestID == v.ID {
				s.switchToDestID = ""
				if s.dstState[currentUID] != DstStateRunning {
					s.dstState[currentUID] = DstStateRunning
				}
				return
			}
		}
	}

	newUID := s.GetDstUIDDestID(s.switchToDestID)
	if s.dstState[newUID] == DstStateStandBy {

		contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Switch from UID:%d to UID:%d ", currentUID, newUID)

=======
	newUID := s.GetDstUIDDestID(s.switchToDestID)

	if newUID < 0 {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, fmt.Sprintf(lumerinlib.FileLineFunc()+" switchToDestID:%s has no UID ", s.switchToDestID))
		return // Because LevelPanic does not seem to be panicing like it should
	}

	if currentUID == newUID {
		contextlib.Logf(s.Ctx(), contextlib.LevelError, fmt.Sprintf(lumerinlib.FileLineFunc()+" new destis the current dest, skipping switch"))
		s.switchToDestID = ""
		return
	}

	//
	// Verify the state of the current route
	//
	if currentUID >= 0 {

		state := s.GetDstStateUid(currentUID)
		switch state {
		case DstStateRunning:
			v, ok := s.dstDest[currentUID]

			if v == nil {
				contextlib.Logf(s.Ctx(), contextlib.LevelPanic, fmt.Sprintf(lumerinlib.FileLineFunc()+" dstDest[%d] ", currentUID))
				panic("")
			}

			if ok {
				if s.switchToDestID == v.ID {
					s.switchToDestID = ""
					contextlib.Logf(s.Ctx(), contextlib.LevelWarn, fmt.Sprintf(lumerinlib.FileLineFunc()+" New Dest is current Dest: %s, UID:[%d] ", v.ID, currentUID))
					return
				}
			}

		case DstStateStandBy:
			contextlib.Logf(s.Ctx(), contextlib.LevelWarn, fmt.Sprintf(lumerinlib.FileLineFunc()+" UID:[%d] is in standby mode... huh? ", currentUID))
		case DstStateClosed:
			contextlib.Logf(s.Ctx(), contextlib.LevelWarn, fmt.Sprintf(lumerinlib.FileLineFunc()+" UID:[%d] is closed... huh? ", currentUID))
		default:
			contextlib.Logf(s.Ctx(), contextlib.LevelPanic, fmt.Sprintf(lumerinlib.FileLineFunc()+" UID:[%d] is in state:%s ", currentUID, state))
		}

	}

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Current:%d New:%d ", currentUID, newUID)

	state := s.GetDstStateUid(newUID)
	switch state {
	case DstStateStandBy:

		contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Switch from UID:%d to UID:%d ", currentUID, newUID)

		s.minerRec.Dest = s.switchToDestID

>>>>>>> pr-009
		if currentUID >= 0 {
			s.dstState[currentUID] = DstStateStandBy
		}
		s.dstState[newUID] = DstStateRunning
		s.protocol.SetDefaultRouteUID(newUID)

		//
		// Goose the miner to the correct Extranonce settings.
<<<<<<< HEAD
		//
		s.sendSetExtranonceNotice(newUID)
		s.sendSetDifficultyNotice(newUID)
		s.sendLastMiningNotice(newUID)
		s.sendLastReqNotice(newUID)

		// Reset the switch to state
		s.switchToDestID = ""

	} else {

		contextlib.Logf(s.Ctx(), contextlib.LevelError, fmt.Sprintf(lumerinlib.FileLineFunc()+" Ignore... next dest not in standby mode %s", s.dstState[newUID]))
=======
		// Then set the difficulty, the feed the last mining notice in
		//
		s.sendSetExtranonceNotice(newUID)
		s.sendLastSetDifficultyNotice(newUID)
		s.sendLastMiningNotice(newUID)
		s.sendLastReqNotify(newUID)

		// Reset the switchToState
		s.switchToDestID = ""

	case DstStateRunning:
		contextlib.Logf(s.Ctx(), contextlib.LevelWarn, fmt.Sprintf(lumerinlib.FileLineFunc()+" UID:%d already in RunningState", newUID))

	case DstStateRedialing:
		// Set switch event timer HERE say after a few seconds, and have it reset if another event takes its place?
		contextlib.Logf(s.Ctx(), contextlib.LevelInfo, fmt.Sprintf(lumerinlib.FileLineFunc()+" UID:%d Redialing", newUID))

	case DstStateClosed:
		contextlib.Logf(s.Ctx(), contextlib.LevelError, fmt.Sprintf(lumerinlib.FileLineFunc()+" UID:%d Redialing", newUID))

	default:
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, fmt.Sprintf(lumerinlib.FileLineFunc()+" UID:%d State:%s", newUID, state))
>>>>>>> pr-009
	}

}

//
// sendConfigure()
//
func (s *StratumV1Struct) sendConfigure() {

	if s.srcConfigure == nil {
		return
	}

	configure := s.srcConfigure

	msg, e := configure.createRequestMsg()
	if e != nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createRequestMsg() error:%e ", e)
	}

	LogJson(s.Ctx(), lumerinlib.FileLineFunc(), JSON_SEND_STOR2DST, msg)

	count, e := s.protocol.Write(msg)
	if e != nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" WriteSrc error:%s", e)
	}
	if count != len(msg) {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
	}
}

//
// sendExtranonoce()
//
func (s *StratumV1Struct) sendExtranonce() {

	if s.srcExtranonce == nil {
		return
	}

	extranonce := s.srcExtranonce

	msg, e := extranonce.createRequestMsg()
	if e != nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" createRequestMsg() error:%e ", e)
	}

	LogJson(s.Ctx(), lumerinlib.FileLineFunc(), JSON_SEND_STOR2DST, msg)

	count, e := s.protocol.Write(msg)
	if e != nil {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" WriteSrc error:%s", e)
	}
	if count != len(msg) {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
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

//
// sendSetExtranonceNotice()
//
func (svs *StratumV1Struct) sendSetExtranonceNotice(uid simple.ConnUniqueID) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" on  UID:%d", uid)

	_, ok := svs.dstExtranonce[uid]
	if !ok {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" dstExtranonce[%d] DNE ", uid)
	}
	_, ok = svs.dstExtranonce2size[uid]
	if !ok {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" dstExtranonce2size[%d] DNE ", uid)
	}

	msg, e := createSetExtranonceNoticeMsg(svs.dstExtranonce[uid], svs.dstExtranonce2size[uid])

	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createSetExtranonceNoticeMsg error:%s", e)
		return e
	}

	LogJson(svs.Ctx(), lumerinlib.FileLineFunc(), JSON_SEND_STOR2SRC, msg)

	count, e := svs.protocol.WriteSrc(msg)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Write error:%s", e)
		return e
	}

	if count != len(msg) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Write bad count:%d, %d", count, len(msg))
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
	}

	return e

}

//
<<<<<<< HEAD
// sendSetDifficultyNotice()
//
func (svs *StratumV1Struct) sendSetDifficultyNotice(uid simple.ConnUniqueID) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" on  UID:%d", uid)

	_, ok := svs.dstDiff[uid]
	if !ok {
		contextlib.Logf(svs.Ctx(), contextlib.LevelWarn, lumerinlib.FileLineFunc()+" dstDiff[%d] DNE ", uid)
		return nil
	}

	msg, e := createSetDifficultyNoticeMsg(svs.dstDiff[uid])
=======
// setLastSetDifficultyNotice()
//
func (svs *StratumV1Struct) setLastSetDifficultyNotice(uid simple.ConnUniqueID, n *stratumNotice) (e error) {

	if n.Method != string(SERVER_MINING_SET_DIFFICULTY) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" bad Method[%s]", n.Method)
	}

	diff, e := n.getSetDifficulty()
	svs.dstLastSetDiff[uid] = diff

	return e

}

//
// setLastReqSetDifficulty()
//
func (svs *StratumV1Struct) setLastReqSetDifficulty(uid simple.ConnUniqueID, r *stratumRequest) (e error) {

	if r.Method != string(SERVER_MINING_SET_DIFFICULTY) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" bad Method[%s]", r.Method)
	}

	diff, e := r.getSetDifficulty()
	svs.dstLastSetDiff[uid] = diff

	return e

}

//
// sendSetDifficultyNotice()
//
func (svs *StratumV1Struct) sendLastSetDifficultyNotice(uid simple.ConnUniqueID) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" on  UID:%d", uid)

	diff, ok := svs.dstLastSetDiff[uid]
	if !ok {
		contextlib.Logf(svs.Ctx(), contextlib.LevelWarn, lumerinlib.FileLineFunc()+" dstLastSetDiffNotice[%d] DNE ", uid)
		return nil
	}

	cs := contextlib.GetContextStruct(svs.Ctx())
	ps := cs.GetMsgBus()
	ps.SendValidateSetDiff(svs.Ctx(), svs.minerRec.ID, svs.dstDest[uid].ID, diff)

	msg, e := createSetDifficultyNoticeMsg(diff)
>>>>>>> pr-009

	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createSetDifficultyNoticeMsg() error:%s", e)
		return e
	}

	LogJson(svs.Ctx(), lumerinlib.FileLineFunc(), JSON_SEND_STOR2SRC, msg)

	count, e := svs.protocol.WriteSrc(msg)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Write error:%s", e)
		return e
	}

	if count != len(msg) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Write bad count:%d, %d", count, len(msg))
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
	}

	return e

}

//
<<<<<<< HEAD
=======
// setLastMiningNotice()
//
func (svs *StratumV1Struct) setLastMiningNotice(uid simple.ConnUniqueID, n *stratumNotice) (e error) {

	if n.Method != string(SERVER_MINING_NOTIFY) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" bad Method[%s]", n.Method)
	}

	svs.dstLastMiningNotice[uid] = n

	return e

}

//
>>>>>>> pr-009
// sendLastMiningNotice()
//
func (svs *StratumV1Struct) sendLastMiningNotice(uid simple.ConnUniqueID) (e error) {

	contextlib.Logf(svs.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" on  UID:%d", uid)

	_, ok := svs.dstLastMiningNotice[uid]
	if !ok {
<<<<<<< HEAD
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" dstLastMiningNotice[%d] DNE ", uid)
=======
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" dstLastMiningNotice[%d] DNE, skipping ", uid)
>>>>>>> pr-009
		return nil
	}

	if svs.dstLastMiningNotice[uid] == nil {
		return nil
	}

	notice := svs.dstLastMiningNotice[uid]
<<<<<<< HEAD
	svs.dstLastMiningNotice[uid] = nil
=======
	minerID := svs.minerRec.ID
	destID := svs.minerRec.Dest
	username := svs.dstDest[uid].Username()
	n := notice.Params.([]interface{})
	jobID := n[0].(string)
	prevblock := n[1].(string)
	gen1 := n[2].(string)
	gen2 := n[3].(string)
	merkel := n[4].([]interface{})
	version := n[5].(string)
	nbits := n[6].(string)
	ntime := n[7].(string)
	clean := n[8].(bool)

	cs := contextlib.GetContextStruct(svs.Ctx())
	ps := cs.GetMsgBus()
	ps.SendValidateNotify(svs.Ctx(), minerID, destID, username, jobID, prevblock, gen1, gen2, merkel, version, nbits, ntime, clean)
>>>>>>> pr-009

	msg, e := notice.createNoticeMsg()

	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createLastMiningNoticeMsg() error:%s", e)
		return e
	}

	LogJson(svs.Ctx(), lumerinlib.FileLineFunc(), JSON_SEND_STOR2SRC, msg)

	count, e := svs.protocol.WriteSrc(msg)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Write error:%s", e)
		return e
	}

	if count != len(msg) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Write bad count:%d, %d", count, len(msg))
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
	}

	return e

}

//
<<<<<<< HEAD
// sendLastReqNotice()
//
func (svs *StratumV1Struct) sendLastReqNotice(uid simple.ConnUniqueID) (e error) {

	_, ok := svs.dstLastReqNotice[uid]
	if !ok {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" dstLastReqNotice[%d] DNE ", uid)
		return nil
	}

	if svs.dstLastReqNotice[uid] == nil {
		return nil
	}

	request := svs.dstLastReqNotice[uid]
	svs.dstLastReqNotice[uid] = nil
=======
// setLastReqNotify()
//
func (svs *StratumV1Struct) setLastReqNotify(uid simple.ConnUniqueID, r *stratumRequest) (e error) {

	if r.Method != string(SERVER_MINING_NOTIFY) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" bad Method[%s]", r.Method)
	}

	svs.dstLastReqNotify[uid] = r

	return e

}

//
// sendLastReqNotify()
//
func (svs *StratumV1Struct) sendLastReqNotify(uid simple.ConnUniqueID) (e error) {

	_, ok := svs.dstLastReqNotify[uid]
	if !ok {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" dstLastReqNotify[%d] DNE ", uid)
		return nil
	}

	if svs.dstLastReqNotify[uid] == nil {
		return nil
	}

	request := svs.dstLastReqNotify[uid]
>>>>>>> pr-009

	msg, e := request.createRequestMsg()

	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" createRequestMsg() error:%s", e)
		return e
	}

	LogJson(svs.Ctx(), lumerinlib.FileLineFunc(), JSON_SEND_STOR2SRC, msg)

	count, e := svs.protocol.WriteSrc(msg)
	if e != nil {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Write error:%s", e)
		return e
	}

	if count != len(msg) {
		contextlib.Logf(svs.Ctx(), contextlib.LevelError, lumerinlib.FileLineFunc()+" Write bad count:%d, %d", count, len(msg))
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" WriteSrc bad count:%d, %d", count, len(msg))
	}

	return e
<<<<<<< HEAD

=======
>>>>>>> pr-009
}
