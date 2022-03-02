package protocol

import (
	"context"
	"fmt"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
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

	//
	// Basic error checking, make sure that the SimpleContextStruct is
	// filled out correctly
	//
	sc := ctx.Value(simple.SimpleContext)
	if sc == nil {
		lumerinlib.PanicHere("")
	}

	sc, ok = sc.(simple.SimpleContextStruct)
	if !ok {
		lumerinlib.PanicHere("")
	}
	if sc.(simple.SimpleContextStruct).Protocol == nil {
		lumerinlib.PanicHere("")
	}
	if sc.(simple.SimpleContextStruct).MsgBus == nil {
		lumerinlib.PanicHere("")
	}
	if sc.(simple.SimpleContextStruct).Dst == nil {
		lumerinlib.PanicHere("")
	}
	if sc.(simple.SimpleContextStruct).Src == nil {
		lumerinlib.PanicHere("")
	}

	ctx, cancel := context.WithCancel(ctx)

	listenaddr := sc.(simple.SimpleContextStruct).Src
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
// Cancel() calls the simple layer Cancel function on the SimpleListenStruct
//
func (pls *ProtocolListenStruct) Cancel() {
	pls.cancel()
}

//
// Run() calls the simple layer Run function on the SimpleListenStruct
//
func (pls *ProtocolListenStruct) Run() {
	go pls.goAccept()
}

//
// goAccept()
//
func (pls *ProtocolListenStruct) goAccept() {

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

	ctx, cancel := context.WithCancel(s.Ctx())
	eventchan := make(chan *simple.SimpleEvent)
	scs := s.Ctx().Value(simple.SimpleContext)
	src := scs.(simple.SimpleContextStruct).Src
	dst := scs.(simple.SimpleContextStruct).Dst

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
// Cancel() calls the simple layer Cancel function on the SimpleListenStruct
//
func (ps *ProtocolStruct) Cancel() {
	ps.cancel()
}

//
// Run() calls the simple layer Run function on the SimpleListenStruct
//
func (ps *ProtocolStruct) Run() {
	// go ps.goAccept()
}

//
//
//
func (ps *ProtocolStruct) Event() chan *simple.SimpleEvent {
	return ps.eventchan
}

//
//
//p
func (ps *ProtocolStruct) OpenConn(dst net.Addr) (index int, e error) {
	return ps.dstconn.openConn(dst)
}
