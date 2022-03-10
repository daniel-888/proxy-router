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
	srcconn   *ProtocolConnectionStruct
	dstconn   *ProtocolDstStruct
	msgbus    *ProtocolMsgBusStruct
}

//
// NewListen() Create a new ProtocolListenStruct
// Opens the default destination
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
	sls, err := simple.New(ctx, listenaddr, func() {})
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
		for {
			select {
			case <-pls.ctx.Done():
				return
			case newSimpleStruct := <-pls.simplelisten.Accept():
				newSimpleStruct.Run(pls.ctx)
			}
		}
	}()
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
		srcconn: &ProtocolConnectionStruct{
			Addr: src,
			Id:   0,
		},
		dstconn: &ProtocolDstStruct{
			conn: make(map[int]*ProtocolConnectionStruct),
		},
		msgbus: &ProtocolMsgBusStruct{},
	}

	// Open the default connection
	index, e := ps.OpenConn(dst)
	// First connection should be zero
	if index != 0 {
		contextlib.Logf(s.Ctx(), contextlib.LevelPanic, lumerinlib.FileLine()+" called")
	}
	e = ps.SetDefaultRoute(index)

	return pls, e
}

//
// Ctx() returns the context of the ProtocolStruct
//
func (ps *ProtocolStruct) Ctx() context.Context {
	return ps.ctx
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
// OpenConn()
// opens a new connection to the desitnation and returns the index of it
//
func (ps *ProtocolStruct) OpenConn(dst net.Addr) (index int, e error) {

	contextlib.Logf(ps.ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	//
	// Call simple layer to dial up a connection to the dst
	//
	id, e := ps.simple.Dial(dst)
	if e != nil {
		return -1, e
	}

	//
	// Add new connection to the ProtocolDstStruct
	//
	return ps.dstconn.addConn(dst, id)
}

//
// SetDefaultRoute()
// Set the SIMPL layer default route
//
func (ps *ProtocolStruct) SetDefaultRoute(index int) (e error) {

	id, e := ps.dstconn.getConnID(index)
	if e != nil {
		return e
	}

	// Set the default route to the first route
	// e := ps.simple.SetRoute(id)
	// if e != nil {}

	ps.simple.SetRoute(id)

	return nil
}

//
// GetDefaultRoute()
// get the  SIMPL layer default route
//
func (ps *ProtocolStruct) GetDefaultRoute() (index int, e error) {

	// Set the default route to the first route
	ps.simple.GetRoute()
	var id int = 0
	//id, e := ps.simple.GetRoute()
	//if e != nil {
	//	return -1, e
	//}

	index, e = ps.dstconn.getConnIndex(simple.ConnUniqueID(id))

	return index, e

}

//
// Write()
//
func (ps *ProtocolStruct) Write(msg []byte) (count int, e error) {
	count = 0
	// l := len(msg)

	index, e := ps.GetDefaultRoute()
	if e != nil {
		return -1, e
	}

	id, e := ps.dstconn.getConnID(index)
	if e != nil {
		return -1, e
	}

	ps.simple.Write(id, msg)

	// count, e = ps.simple.Write(id, msg)
	// if e != nil {
	// 	return count, e
	// }
	// if l != count {
	// 	return count, fmt.Errorf(lumerinlib.FileLine() + " full msg length was not written")
	// }

	return count, e
}

//
// WriteSrc()
//
func (ps *ProtocolStruct) WriteSrc(msg []byte) (count int, e error) {
	count = 0

	return count, e
}

//
// WriteDst()
//
func (ps *ProtocolStruct) WriteDst(index int, msg []byte) (count int, e error) {
	count = 0

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
