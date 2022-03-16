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

	ctx := context.Background()

	localport := getRandPort()
	addr := fmt.Sprintf("%s:%d", ip, localport)

	sls := newConnection(t, ctx, addr, addr)
	sls.Run()
	sls.Cancel()

}

func TestNewConnection(t *testing.T) {

	ctx := context.Background()

	// Local Node Listener port
	// Listen on the Default destination address port
	localport := getRandPort()
	srcaddr := fmt.Sprintf("%s:%d", ip, localport)
	fakeListener, dstport := fakeListener(ctx)
	dstaddr := fmt.Sprintf("%s:%d", ip, dstport)

	// Open the actual stratumV1 test connection
	sls := newConnection(t, ctx, srcaddr, dstaddr)
	sls.Run()

	// Run a incoming test connection to Stratum Listen port
	cs := contextlib.GetContextStruct(sls.Ctx())
	ctx = contextlib.SetContextStruct(ctx, cs)
	testconn, e := connect(t, ctx, srcaddr)
	if e != nil {
		t.Errorf("connect() error:%s", e)
	}

	// Accept the expected dest connection (BLOCKING)
	// should happen after the connect to the node
	// FIX HERE
	dstsoc := <-fakeListener.Accept()
	if e != nil {
		t.Errorf("Accept() error:%s", e)
	}
	t.Logf("Accepted connection")

	// Push data into the Stratum connection
	count, e := testconn.Write([]byte(testString))
	if e != nil {
		t.Errorf("Write() error:%s", e)
	}
	if count != len(testString) {
		t.Errorf("Write() error:%s", e)
	}
	t.Logf("Wrote Test String into connection")

	// Need Read Relay here  stratum.Read -> default dest Write, should be event handler

	// Loof for the data on the  default dst connection
	// rr := dstsoc.ReadReady()
	var dstbuf []byte = make([]byte, 1024)
	dstmsgcount, e := dstsoc.Read(dstbuf)
	if e != nil {
		t.Errorf("bad dest message error:%s", e)
	}
	if dstmsgcount != len(testString) {
		t.Errorf("bad dest message lenth() error")
	}

	sls.Cancel()
	<-sls.Ctx().Done()

}

// ---------------------------------------------------------------------------
//
//

//
// newConnection()
//
func newConnection(t *testing.T, ctx context.Context, srcstr string, dststr string) (sls *StratumV1ListenStruct) {

	ps := msgbus.New(1, nil)
	src := lumerinlib.NewNetAddr(lumerinlib.TCP, srcstr)
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, dststr)

	var new = &newStratumV1Struct{
		funcptr: testStratumV1,
	}

	sls, err := NewListener(ctx, ps, src, dst, new)

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

	if ss == nil {
		panic(lumerinlib.FileLineFunc() + " nil SimpleStruct")
	}
	if ss.ConnectionStruct == nil {
		panic(lumerinlib.FileLineFunc() + " nil SimpleStruct.ConnectionStruct")
	}

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

	ss.Run()

}

//
//
//
func getRandPort() (port int) {
	port = rand.Intn(10000) + basePort
	return port
}

//
//
//
func fakeListener(ctx context.Context) (l *sockettcp.ListenTCPStruct, port int) {

	port = getRandPort()
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	l, e := sockettcp.Listen(ctx, "tcp", addr)
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Create sockettcp.Listen() errored:%s", e)
	}

	_, port, e = l.LocalAddr()
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Create sockettcp.Listen() errored:%s", e)
	}

	return l, port
}
