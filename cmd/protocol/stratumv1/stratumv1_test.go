package stratumv1

import (
	"context"
	"fmt"
	"testing"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/sockettcp"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

var port int = 3334
var ip string = "127.0.0.1"
var testString = "This is a test string\n"

//
//
//
func TestNewProto(t *testing.T) {

	sls := newConnection(t)
	sls.Run()
	sls.Cancel()

}

func TestNewConnection(t *testing.T) {

	sls := newConnection(t)
	sls.Run()

	s, e := connect(t, sls.Ctx())
	if e != nil {
		t.Errorf("connect() error:%s", e)
	}

	count, e := s.Write([]byte(testString))
	if e != nil {
		t.Errorf("Write() error:%s", e)
	}
	if count != len(testString) {
		t.Errorf("Write() error:%s", e)
	}

	sls.Cancel()

}

//
//
//

//
// newConnection()
//
func newConnection(t *testing.T) (sls *StratumV1ListenStruct) {

	addr := fmt.Sprintf("%s:%d", ip, port)
	ps := msgbus.New(1, nil)
	src := lumerinlib.NewNetAddr(lumerinlib.TCP, addr)
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, addr)

	ctx := context.Background()

	sls, err := New(ctx, ps, src, dst, StratumV1Func)
	if err != nil {
		t.Errorf("New() returne error:%s", err)
	}

	return sls

}

//
//
//
func connect(t *testing.T, ctx context.Context) (s *sockettcp.SocketTCPStruct, e error) {
	s, e = sockettcp.Dial(ctx, "tcp", fmt.Sprintf("%s:%d", ip, port))
	if e != nil {
		t.Errorf("Dial() returned error:%s", e)
	}

	return s, e
}

//
//
func StratumV1Func(ss *simple.SimpleStruct) chan *simple.SimpleEvent {

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
	pls, err := protocol.NewProtocol(ss)
	if err != nil {
		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Create NewProtocol() failed: %s", err)
	}

	svs := &StratumV1Struct{
		protocol:            pls,
		minerRec:            nil,
		srcSubscribeRequest: nil,
		srcAuthRequest:      nil,
		// Fill in other state information here
	}

	// Launch the event handler
	go svs.goEvent()

	// return the event handler channel to the caller (the simple layer accept() function )
	return svs.protocol.Event()
}
