package protocol

import (
	"context"
	"fmt"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

//
// Top layer protocol template functions that a new protocol will use to access the SIMPLe layer
//

type ProtocolListenStruct struct {
	ctx          context.Context
	cancel       func()
	simplelisten *simple.SimpleListenStruct
	accept       chan *ProtocolStruct
}

type ProtocolStruct struct {
	ctx       context.Context
	cancel    func()
	simple    *simple.SimpleStruct
	eventchan chan *simple.SimpleEvent
	srcconn   *ProtocolConnectionStruct
	dstconn   *ProtocolDstStruct
	msgbus    *ProtocolMsgBusStruct
	defRoute  int
}

//
// NewListen() Create a new ProtocolListenStruct
// Opens the default destination
//
func NewListen(ctx context.Context) (pls *ProtocolListenStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	//
	// Basic error checking, make sure that the ContextStruct is
	// filled out correctly
	//
	cs := contextlib.GetContextStruct(ctx)
	if cs == nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Structre not present")
	}

	// if cs.GetProtocol() == nil {
	// 	cs.Logf(contextlib.LevelPanic, "Context Protocol not defined")
	// }
	if cs.GetMsgBus() == nil {
		cs.Logf(contextlib.LevelPanic, "Context MsgBus not defined")
	}
	if cs.GetSrc() == nil {
		cs.Logf(contextlib.LevelPanic, "Context Src Addr not defined")
	}
	if cs.GetDst() == nil {
		cs.Logf(contextlib.LevelPanic, "Context Dst Addr not defined")
	}

	ctx, cancel := context.WithCancel(ctx)

	listenaddr := contextlib.GetSrc(ctx)

	sls, err := simple.NewListen(ctx, listenaddr)

	if err != nil {
		lumerinlib.PanicHere(fmt.Sprintf("Error:%s", err))
	}

	accept := make(chan *ProtocolStruct)
	pls = &ProtocolListenStruct{
		ctx:          ctx,
		cancel:       cancel,
		simplelisten: &sls,
		accept:       accept,
	}

	return pls, e
}

//
//
//
func (p *ProtocolListenStruct) GetAccept() <-chan *ProtocolStruct {
	return p.accept
}

//
// Ctx() returns the current context of the ProtocolListenStruct
//
func (p *ProtocolListenStruct) Ctx() context.Context {
	return p.ctx
}

//
// Cancel() calls the simple layer Cancel function on the SimpleListenStruct
//
func (pls *ProtocolListenStruct) Cancel() {

	contextlib.Logf(pls.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	pls.cancel()
}

//
// Run() calls the simple layer Run function on the SimpleListenStruct
//
func (pls *ProtocolListenStruct) Run() {

	contextlib.Logf(pls.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	pls.simplelisten.Run()
	go pls.goListenAccept()
}

//
// goAccept()
// Listens for the Accept event from the SIMPL layer
// The simpleStruct has already been inialized
//
func (pls *ProtocolListenStruct) goListenAccept() {

	contextlib.Logf(pls.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	newSimpleStructChan := pls.simplelisten.GetAccept()

FORLOOP:
	for {
		select {
		case <-pls.ctx.Done():
			break FORLOOP
		case newSimpleStruct := <-newSimpleStructChan:
			contextlib.Logf(pls.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" simplelisten.Accept() recieved")

			if newSimpleStruct == nil {
				contextlib.Logf(pls.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" pls.simplelisten.Accept() stopping")
				pls.Cancel()
				break FORLOOP
			}
			if nil == newSimpleStruct.Ctx() {
				contextlib.Logf(pls.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" newSimpleStruct empty CTX stopping")
				pls.Cancel()
				break FORLOOP
			}
			if nil == newSimpleStruct.ConnectionStruct {
				contextlib.Logf(pls.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" newSimpleStruct empty ConnectionStruct stopping")
				pls.Cancel()
				break FORLOOP
			}

			ps, e := NewProtocol(pls.ctx, newSimpleStruct)
			if e != nil {
				contextlib.Logf(pls.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" NewProtocol() error:%s", e)
			}
			if ps == nil {
				contextlib.Logf(pls.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" NewProtocol() returned nil, assuming closed")
				break
			}

			pls.accept <- ps
		}
	}
}

// --------------------------------------------
// ProtocolStruct functions
//

//
// NewProtocol() takes a simple struct and creates a ProtocolStruct, pulls the Src and Dst from the context
// and initiates a connection to the defualt Dst address
// This function is called from the layer above to initalize the common protocol functions, and enable
// access to the standard functions provided by this layer.
//
func NewProtocol(ctx context.Context, s *simple.SimpleStruct) (ps *ProtocolStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	dst := contextlib.GetDst(ctx)
	if dst == nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetDst() returned nil")
	}

	ps = NewProtocolStruct(ctx, s)

	// Fire up the default destination here
	index, e := ps.AsyncDial(dst)
	if index != 0 {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" AsyncDial returned non-zero for the default dest")
	}

	return ps, e
}

func NewProtocolStruct(ctx context.Context, s *simple.SimpleStruct) (n *ProtocolStruct) {

	ctx, cancel := context.WithCancel(s.Ctx())
	cs := contextlib.GetContextStruct(s.Ctx())
	src := cs.GetSrc()
	eventchan := make(chan *simple.SimpleEvent)
	pcs := NewProtocolConnectionStruct(ctx, src)
	pcs.SetState(ConnStateReady)

	n = &ProtocolStruct{
		ctx:       ctx,
		cancel:    cancel,
		simple:    s,
		eventchan: eventchan,
		srcconn:   pcs,
		dstconn: &ProtocolDstStruct{
			ctx:  ctx,
			conn: make(map[int]*ProtocolConnectionStruct),
		},
		msgbus: &ProtocolMsgBusStruct{},
	}

	return n
}

//
// Ctx() returns the context of the ProtocolStruct
//
func (ps *ProtocolStruct) Ctx() context.Context {
	if ps == nil {
		panic(lumerinlib.FileLineFunc() + "ProtocolStruct is nil")
	}
	return ps.ctx
}

//
// Cancel() calls the simple layer Cancel function on the SimpleListenStruct
//
func (ps *ProtocolStruct) Cancel() {
	if ps == nil {
		panic(lumerinlib.FileLineFunc() + "ProtocolStruct is nil")
	}
	contextlib.Logf(ps.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")
	ps.cancel()
}

//
// Run() calls the simple layer Run function on the SimpleListenStruct
//
func (ps *ProtocolStruct) Run() {
	if ps == nil {
		panic(lumerinlib.FileLineFunc() + "ProtocolStruct is nil")
	}
	contextlib.Logf(ps.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")
	ps.simple.Run()
}

//
//
//
func (ps *ProtocolStruct) GetSimpleEventChan() <-chan *simple.SimpleEvent {
	if ps == nil {
		panic(lumerinlib.FileLineFunc() + "ProtocolStruct is nil")
	}
	contextlib.Logf(ps.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")
	c := ps.simple.GetEventChan()
	if c == nil {
		contextlib.Logf(ps.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" GetEVentChan returned nil, ps:%v", ps)
	}
	return c
}

//
// AsyncDial()
// opens a new connection to the desitnation
//
//
func (ps *ProtocolStruct) AsyncDial(dst net.Addr) (index int, e error) {

	contextlib.Logf(ps.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	// Paranoid error checking here
	if ps == nil {
		panic(lumerinlib.FileLineFunc() + "ProtocolStruct is nil")
	}
	if ps.simple == nil {
		contextlib.Logf(ps.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" simple struct is nil")
	}
	if ps.simple.ConnectionStruct == nil {
		contextlib.Logf(ps.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" simple struct CoonectionStruct is nil")
	}

	index, e = ps.dstconn.NewProtocolDstStruct(dst)
	if e == nil {
		e = ps.simple.AsyncDial(index, dst)
	}

	return index, e
}

//
// SetDefaultRoute()
// Set the SIMPL layer default route
//
func (ps *ProtocolStruct) SetDefaultRouteIndex(index int) (e error) {

	slot, ok := ps.dstconn.conn[index]
	if !ok {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" bad index: %d", index)
	} else if ConnStateReady == slot.GetState() {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" connection state not ready: %s", slot.GetState())
	} else if ps.defRoute == index {
		contextlib.Logf(ps.ctx, contextlib.LevelWarn, lumerinlib.FileLineFunc()+" default route already set to index: %d", index)
	} else {
		uid := slot.GetUID()
		ps.simple.SetRoute(uid) // Are we going to keep this, or just assume that the protocol layer will know the default route?
		ps.defRoute = index
	}

	return e
}

//
// DstConn()
//
func (ps *ProtocolStruct) GetDstConn(index int) (pcs *ProtocolConnectionStruct, e error) {
	if ps.dstconn.conn[index] == nil {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+"Index:%d does not exist", index)
	} else {
		pcs = ps.dstconn.conn[index]
	}
	return pcs, e
}

//
// GetDefaultRouteIndex()
// get the  SIMPL layer default route
//
func (ps *ProtocolStruct) GetDefaultRouteIndex() int {
	return ps.defRoute
}

//
// Write() writes to the default route
//
func (ps *ProtocolStruct) Write(msg []byte) (count int, e error) {
	count = 0

	index := ps.defRoute
	state := ps.dstconn.conn[index].GetState()
	if ConnStateReady != state {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+" Connection state:%s", state)
	} else {
		count, e = ps.simple.Write(index, msg)
	}

	return count, e
}

//
// WriteSrc()
//
func (ps *ProtocolStruct) WriteSrc(msg []byte) (count int, e error) {

	count, e = ps.simple.Write(-1, msg)

	return count, e
}

//
// WriteDst()
//
func (ps *ProtocolStruct) WriteDst(index int, msg []byte) (count int, e error) {
	count, e = ps.simple.Write(index, msg)
	return count, e
}

//
// Pub() publishes data, and stores the request ID to match the Completion Event
//
func (ps *ProtocolStruct) Pub(msgtype simple.MsgType, id simple.ID, data simple.Data) (rID int, e error) {

	// rID, e = ps.simple.Pub(msgtype, id, data)
	e = ps.simple.Pub(msgtype, id, data)

	return 0, e
}

//
//
//
func (ps *ProtocolStruct) Unpub(msgtype simple.MsgType, id simple.ID) (rID int, e error) {

	return 0, nil
}

//
//
//
func (ps *ProtocolStruct) Sub(msgtype simple.MsgType, id simple.ID, eh func()) (rID int, e error) {

	return 0, nil
}

//
//
//
func (ps *ProtocolStruct) Unsub(msgtype simple.MsgType, id simple.ID, eh func()) (rID int, e error) {

	return 0, nil
}

//
//
//
func (ps *ProtocolStruct) Get(msgtype simple.MsgType, id simple.ID, eh func()) (rID int, e error) {

	return 0, nil
}

//
//
//
func (ps *ProtocolStruct) Set(msgtype simple.MsgType, id simple.ID, data interface{}) (rID int, e error) {

	return 0, nil
}

//
//
//
func (ps *ProtocolStruct) SearchIP(msgtype simple.MsgType, search string) (rID int, e error) {

	return 0, nil
}

//
//
//
func (ps *ProtocolStruct) SearchMac(msgtype simple.MsgType, search string) (rID int, e error) {

	return 0, nil
}

//
//
//
func (ps *ProtocolStruct) SearchName(msgtype simple.MsgType, search string) (rID int, e error) {

	return 0, nil
}
