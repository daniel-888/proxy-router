package stratumv1

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/sockettcp"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

var basePort int = 50000
var ip string = "127.0.0.1"
var testString = "This is a test string\n"

//
//
//
func TestNewProto(t *testing.T) {

	localport := getRandPort()
	addr := fmt.Sprintf("%s:%d", ip, localport)

	sls := newConnection(t, addr)
	sls.Run()
	sls.Cancel()

}

func TestNewConnection(t *testing.T) {

	localport := getRandPort()
	addr := fmt.Sprintf("%s:%d", ip, localport)

	sls := newConnection(t, addr)
	sls.Run()

	s, e := connect(t, sls.Ctx(), addr)
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

	<-sls.Ctx().Done()
	sls.Cancel()

}

// ---------------------------------------------------------------------------
//
//

//
// newConnection()
//
func newConnection(t *testing.T, addr string) (sls *StratumV1ListenStruct) {

	ps := msgbus.New(1, nil)
	src := lumerinlib.NewNetAddr(lumerinlib.TCP, addr)
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, addr)

	ctx := context.Background()

	var new = &newStratumV1Struct{
		funcptr: testStratumV1,
	}

	sls, err := NewListener(ctx, ps, src, dst, new)

	// sls, err := NewListener(ctx, ps, src, dst, StratumV1Func)
	if err != nil {
		t.Errorf("NewListner() returned error:%s", err)
	}

	return sls

}

//
//
//
func connect(t *testing.T, ctx context.Context, addr string) (s *sockettcp.SocketTCPStruct, e error) {
	s, e = sockettcp.Dial(ctx, "tcp", addr)
	if e != nil {
		t.Errorf("Dial() returned error:%s", e)
	}

	return s, e
}

//
//
//
func testStratumV1(ss *simple.SimpleStruct) {

	contextlib.Logf(ss.Ctx(), contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Called")

	cs := contextlib.GetContextStruct(ss.Ctx())

	if cs == nil {
		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Context Structre not correct")
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
	if pls == nil {
		contextlib.Logf(ss.Ctx(), contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Create NewProtocol() failed - no pointer returned")
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

}

//
//
//
func getRandPort() (port int) {
	port = rand.Intn(10000) + basePort
	return port
}
