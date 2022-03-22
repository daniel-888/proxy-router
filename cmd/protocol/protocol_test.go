package protocol

import (
	"context"
	"fmt"
	"testing"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/sockettcp"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

type newProtocolFunc func(*simple.SimpleStruct) chan *simple.SimpleEvent
type newProtocolStruct struct {
	funcptr newProtocolFunc
}

var port int = 12345
var ip string = "127.0.0.1"

func TestNewProto(t *testing.T) {
	pls := newListen(t)
	pls.Run()
	pls.Cancel()
	select {
	case <-pls.ctx.Done():
	}

}

//
// TestNewConnection()
// Test a new Listener, connects to it, pushes data though it, then validates it.
//
func TestNewConnection(t *testing.T) {

	var testString = "This is a test string\n"

	pls := newListen(t)

	pls.Run()

	s, e := connect(t, pls.Ctx())
	if e != nil {
		t.Errorf(lumerinlib.FileLineFunc()+" error:%s", e)
	}

	count, e := s.Write([]byte(testString))
	if e != nil {
		t.Errorf(lumerinlib.FileLineFunc()+" error:%s", e)
	}
	if count != len(testString) {
		t.Errorf(lumerinlib.FileLineFunc()+" count is wrong,sent:%d, recv:%d", len(testString), count)
	}

	// Need to read the data from the event handler here

	pls.Cancel()
	select {
	case <-pls.Ctx().Done():
		return
	}

}

//
// TestConnectionDial()
// test setting up a connection, and then dialing into it
// send data thorugh the connection and validates that it is recieved
//
func TestConnectionDial(t *testing.T) {

}

//
// newProtocolConnection
// This function provides a window into creating a new ProtocolStruct instances
// it creates the instance, and sends back an event channel to send events to
//
func newProtcolConnection(ss *simple.SimpleStruct) {

	contextlib.Logf(ss.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	i := ss.Ctx().Value(contextlib.ContextKey)
	cs, ok := i.(contextlib.ContextStruct)
	if !ok {
		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Struct not in CTX")
	}

	dst := cs.GetDst()
	if dst == nil {
		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Struct DST not defined")
	}

	// inialize a new ProtocolStruct to gain access to the standard protocol functions
	// The default Dst should be opened when this returns
	pls, err := NewProtocol(ss.Ctx(), ss)
	if err != nil {
		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Create NewProtocol() failed: %s", err)
	}

	pls.Run()

	go pls.goEvent()

	// return the event handler channel to the caller (the simple layer accept() function )
	// return pls.Event()

}

func (p *ProtocolStruct) goEvent() {

	contextlib.Logf(p.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	for event := range p.GetSimpleEventChan() {
		contextlib.Logf(p.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Got Event %v", event)
	}
}

//
//
//
func newListen(t *testing.T) (pls *ProtocolListenStruct) {

	ps := msgbus.New(1, nil)
	addr := fmt.Sprintf("%s:%d", ip, port)
	src := lumerinlib.NewNetAddr(lumerinlib.TCP, addr)
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, addr)

	var new = &newProtocolStruct{
		funcptr: NewProtocolFunc,
	}

	ctx := context.Background()
	cs := &contextlib.ContextStruct{}
	cs.SetMsgBus(ps)
	cs.SetSrc(src)
	cs.SetDst(dst)
	cs.SetProtocol(new)
	ctx = context.WithValue(ctx, contextlib.ContextKey, cs)

	pls, e := NewListen(ctx)
	if e != nil {
		t.Errorf(lumerinlib.FileLineFunc()+" error:%s", e)
	}

	return pls
}

//
//
//
func connect(t *testing.T, ctx context.Context) (*sockettcp.SocketTCPStruct, error) {
	_ = t
	return sockettcp.Dial(ctx, "tcp", fmt.Sprintf("%s:%d", ip, port))
}

//
//
//
func (n *newProtocolStruct) NewProtocol(ss *simple.SimpleStruct) chan *simple.SimpleEvent {
	return n.funcptr(ss)
}

//
//
//
func NewProtocolFunc(ss *simple.SimpleStruct) chan *simple.SimpleEvent {

	contextlib.Logf(ss.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")
	return make(chan *simple.SimpleEvent)

	//	i := ss.Ctx().Value(contextlib.ContextKey)
	//	cs, ok := i.(contextlib.ContextStruct)
	//	if !ok {
	//		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Struct not in CTX")
	//	}
	//
	//	dst := cs.GetDst()
	//	if dst == nil {
	//		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Struct DST not defined")
	//	}
	//
	//	// inialize a new ProtocolStruct to gain access to the standard protocol functions
	//	// The default Dst should be opened when this returns
	//	pls, err := protocol.NewProtocol(ss)
	//	if err != nil {
	//		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Create NewProtocol() failed: %s", err)
	//	}
	//
	//	svs := &StratumV1Struct{
	//		protocol:            pls,
	//		minerRec:            nil,
	//		srcSubscribeRequest: nil,
	//		srcAuthRequest:      nil,
	//		// Fill in other state information here
	//	}
	//
	//	// Launch the event handler
	//	go svs.goEvent()
	//
	//	// return the event handler channel to the caller (the simple layer accept() function )
	//	return svs.protocol.Event()
}
