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
}

type ProtocolStruct struct {
	ctx       context.Context
	cancel    func()
	simple    *simple.SimpleStruct
	eventchan chan *simple.SimpleEvent
	srcconn   ProtocolConnectionStruct
	dstconn   ProtocolDstStruct
	msgbus    ProtocolMsgBusStruct
}

//
// New() Create a new ProtocolListenStruct
//
func NewListen(ctx context.Context) (pls *ProtocolListenStruct, e error) {
	var ok bool

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	//
	// Basic error checking, make sure that the ContextStruct is
	// filled out correctly
	//
	c := ctx.Value(contextlib.ContextKey)
	if c == nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLine()+" called")
	}

	cs, ok := c.(*contextlib.ContextStruct)
	if !ok {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLine()+" Context Structre not correct")
	}
	if cs.GetProtocol() == nil {
		cs.Logf(contextlib.LevelPanic, "Context Protocol not defined")
	}
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
	sls, err := simple.New(ctx, listenaddr)
	if err != nil {
		lumerinlib.PanicHere(fmt.Sprintf("Error:%s", err))
	}

	pls = &ProtocolListenStruct{
		ctx:          ctx,
		cancel:       cancel,
		simplelisten: &sls,
	}

	return pls, e
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

	contextlib.Logf(pls.ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	pls.cancel()
}

//
// Run() calls the simple layer Run function on the SimpleListenStruct
//
func (pls *ProtocolListenStruct) Run() {

	contextlib.Logf(pls.ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	go pls.goAccept()
}

//
// goAccept()
//
func (pls *ProtocolListenStruct) goAccept() {

	contextlib.Logf(pls.ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	go func() {
		select {
		case <-pls.ctx.Done():
			return
		case newSimpleStruct := <-pls.simplelisten.Accept():
			newSimpleStruct.Run(pls.ctx)
		}
	}()
	pls.simplelisten.Run()
}

// --------------------------------------------

//
//
//
func NewProtocol(s *simple.SimpleStruct) (pls *ProtocolStruct, e error) {

	contextlib.Logf(s.Ctx(), contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	ctx, cancel := context.WithCancel(s.Ctx())
	eventchan := make(chan *simple.SimpleEvent)
	cs := s.Ctx().Value(contextlib.ContextKey).(contextlib.ContextStruct)
	src := cs.GetSrc()
	dst := cs.GetDst()

	ps := &ProtocolStruct{
		ctx:       ctx,
		cancel:    cancel,
		simple:    s,
		eventchan: eventchan,
		srcconn: ProtocolConnectionStruct{
			Addr: src,
			Id:   0,
		},
		dstconn: ProtocolDstStruct{
			conn: make(map[int]ProtocolConnectionStruct, 0),
		},
		msgbus: ProtocolMsgBusStruct{},
	}

	_, e = ps.dstconn.openConn(dst)

	return pls, e
}

//
// Ctx() gets the context of the ProtocolStruct
//
func (p *ProtocolStruct) Ctx() context.Context {
	return p.ctx
}

//
// Cancel() calls the simple layer Cancel function on the SimpleListenStruct
//
func (ps *ProtocolStruct) Cancel() {

	contextlib.Logf(ps.ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	ps.cancel()
}

//
// Run() calls the simple layer Run function on the SimpleListenStruct
//
func (ps *ProtocolStruct) Run() {

	contextlib.Logf(ps.ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")
	// go ps.goAccept()
}

//
//
//
func (ps *ProtocolStruct) Event() chan *simple.SimpleEvent {

	contextlib.Logf(ps.ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	return ps.eventchan
}

//
//
//p
func (ps *ProtocolStruct) OpenConn(dst net.Addr) (index int, e error) {

	contextlib.Logf(ps.ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	return ps.dstconn.openConn(dst)
}
