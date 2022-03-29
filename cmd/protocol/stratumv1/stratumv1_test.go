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

var test = "This is a test string\n"

var testAuthorizeMsg = `{"id":2,"method":"mining.authorize","params":["testrig.worker1",""]}`

//
// The basics, open up a listening connection...
//
func TestNewProto(t *testing.T) {

	ctx := context.Background()

	localport := getRandPort()
	addr := fmt.Sprintf("%s:%d", ip, localport)
	ctx = newContextStruct(ctx, addr, addr)

	sls := newStratumConnection(t, ctx, addr, addr)
	sls.Run()
	sls.Cancel()

}

func TestNewStratumConnection(t *testing.T) {
	ctx := context.Background()

	localport := getRandPort()
	addr := fmt.Sprintf("%s:%d", ip, localport)
	ctx = newContextStruct(ctx, addr, addr)

	sls := newStratumConnection(t, ctx, addr, addr)
	sls.Run()

	// connect to listening port with socket
	s, e := sockettcp.Dial(ctx, "tcp", addr)
	if e != nil {
		t.Errorf("sockettcp.Dial() returned error:%s", e)
	}
	if s == nil {
		t.Errorf("sockettcp.Dial() returned nil")
	}

	sls.Cancel()
}

//
//
//
func TestNewSrcConnection(t *testing.T) {

	ctx := context.Background()
	ctx = newContextStruct(ctx, "", "")

	//
	// Fakes a destination listener for stratum to connect to
	// DSTPORT
	listener, dstport, e := socketListener(ctx)
	if e != nil {
		t.Fatalf("socketListener failed: %s", e)
	}

	t.Logf("TestNewConnection() Listen port: %d", dstport)

	if listener.Done() {
		t.Fatalf("fakelistener is closed")
	}

	//
	// Local Node Listener port localport
	//
	localport := getRandPort()
	srcaddr := fmt.Sprintf("%s:%d", ip, localport)
	dstaddr := fmt.Sprintf("%s:%d", ip, dstport)

	// ctx = newContextStruct(ctx, srcaddr, dstaddr)

	// Open the actual stratumV1 test connection
	sls := newStratumConnection(t, ctx, srcaddr, dstaddr)
	sls.Run()

	// Trigger a new instance of Stratum
	stratumsrc, e := sockettcp.Dial(ctx, "tcp", srcaddr)
	if e != nil {
		t.Errorf("sockettcp.Dial() returned error:%s", e)
	}
	if stratumsrc == nil {
		t.Errorf("sockettcp.Dial() returned nil")
	}
	_ = stratumsrc

	count, e := stratumsrc.Write([]byte(testAuthorizeMsg))
	if e != nil {
		t.Errorf("sockettcp.Write() returned Error:%s", e)
	}
	if count != len(testString) {
		t.Errorf("sockettcp.Write() returned bad count:%d, != %d", count, len(testString))
	}

	buf := make([]byte, 64)
	count, e = stratumsrc.Read(buf)
	if e != nil {
		t.Errorf("sockettcp.Read() returned Error:%s", e)
	}
	if count == 0 {
		t.Errorf("sockettcp.Read() returned zero count")
	}
	buf = buf[:count]

	t.Logf(" src Read():%s", buf)

	stratumsrc.Close()

	//testconn, e := connect(t, ctx, srcaddr)
	//if e != nil {
	//	t.Errorf("connect() error:%s", e)
	//}

	// Accept the expected dest connection (BLOCKING)
	// should happen after the connect to the node
	// FIX HERE
	//dstsoc := <-listener.Accept()
	//if dstsoc == nil {
	//	t.Errorf("Accept() return nil")
	//}
	//t.Logf("Accepted connection")

	// HERE, this socket is closed
	// Push data into the Stratum connection
	// count, e := testconn.Write([]byte(testString))
	// if e != nil {
	// 	t.Errorf("Write() error:%s", e)
	// }
	// if count != len(testString) {
	// 	t.Errorf("Write() error:%s", e)
	// }
	// t.Logf("Wrote Test String into connection")

	// Need Read Relay here  stratum.Read -> default dest Write, should be event handler

	// Loof for the data on the  default dst connection
	// rr := dstsoc.ReadReady()
	//var dstbuf []byte = make([]byte, 1024)
	//dstmsgcount, e := dstsoc.Read(dstbuf)
	//if e != nil {
	//	t.Errorf("bad dest message error:%s", e)
	//}
	//if dstmsgcount != len(testString) {
	//	t.Errorf("bad dest message lenth() error")
	//}

	<-sls.Ctx().Done()
	t.Logf("Geeting ready to leave")
	sls.Cancel()

}

// ---------------------------------------------------------------------------
//
//
func newContextStruct(ctx context.Context, srcstr string, dstID msgbus.DestID) (ret context.Context) {

	cs := &contextlib.ContextStruct{}

	// cs.SetLog(log.New())
	cs.SetLog(nil)

	cs.SetMsgBus(msgbus.New(1, nil))

	if srcstr != "" {
		src := lumerinlib.NewNetAddr(lumerinlib.TCP, srcstr)
		cs.SetSrc(src)

	}
	if dststr != "" {
		dst := lumerinlib.NewNetAddr(lumerinlib.TCP, dststr)
		cs.SetDstID(dst)
	}

	//
	// This is the only place that SetProtocol is called

	// if len(proto) > 0 {
	// 	cs.SetProtocol(proto[0])
	// } else {
	// 	var new = &newStratumV1Struct{
	// 		funcptr: NewStratumV1, // Set the default new function
	// 	}
	// 	cs.SetProtocol(new)
	// }

	ret = contextlib.SetContextStruct(ctx, cs)
	return ret
}

//
// newConnection()
//
func newStratumConnection(t *testing.T, ctx context.Context, src string, dst string) (sls *StratumV1ListenStruct) {

	cs := contextlib.GetContextStruct(ctx)
	srcaddr := cs.GetSrc()
	dstaddr := cs.GetDst()

	sls, err := NewListener(ctx, srcaddr, dstaddr)

	sls.Run()

	if err != nil {
		t.Errorf("NewListner() returned error:%s", err)
	}
	if sls == nil {
		t.Errorf("NewListner() returned nil")

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
func testNewStratumV1(ss *simple.SimpleStruct) {

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
	pls, err := protocol.NewProtocol(ss.Ctx(), ss)
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
func socketListener(ctx context.Context) (l *sockettcp.ListenTCPStruct, port int, e error) {

	port = getRandPort()
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	l, e = sockettcp.NewListen(ctx, "tcp", addr)
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Create sockettcp.Listen() errored:%s", e)
	}

	l.Run()

	_, port, e = l.LocalAddr()
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Create sockettcp.Listen() errored:%s", e)
	}

	return l, port, e
}
