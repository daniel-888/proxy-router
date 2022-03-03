package lumerinconnection

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"testing"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

//
//
//
func TestSetupListenCancel(t *testing.T) {

	ctx := context.Background()

	testaddr := &testAddr{
		network: "tcp",
		ipaddr:  "127.0.0.1:12345",
	}

	l, e := testListen(ctx, testaddr)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Listen() Failed: %s\n", e))
	}

	l.Cancel()

	_, e = l.Accept()
	if e != nil {
		select {
		case <-ctx.Done():
			fmt.Printf(lumerinlib.FileLine()+" CTX Done(): %s\n", ctx.Err())
		default:
			fmt.Printf(lumerinlib.FileLine()+"Accept() Test Passed: %s\n", e)
		}
	} else {
		t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Accept() Failed: %s\n", e))
	}

}

//
//
//
func TestDial(t *testing.T) {

	ctx := context.Background()
	var TestString = "This is a test string\n"

	testaddr := &testAddr{
		network: "tcp",
		ipaddr:  "127.0.0.1:12346",
	}

	l, e := testListen(ctx, testaddr)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Listen() Failed: %s\n", e))
	}

	go goTestAcceptChannelEcho(l)

	s := testDial(ctx, testaddr)

	fmt.Printf(lumerinlib.FileLine() + " Dial completed\n")

	writeb := []byte(TestString)
	writecount, e := s.Write(writeb)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Write() Test Failed: %s\n", e))
	}
	if writecount != len(writeb) {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Write() Test Failed: %s\n", e))
	}

	fmt.Printf(lumerinlib.FileLine() + " Write() completed\n")

	reader := bufio.NewReader(s)
	readbuf, e := reader.ReadBytes('\n')
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" RedBytes() Test Failed: %s\n", e))
	}
	readcount := len(readbuf)
	if readcount != writecount {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+"Count Test Failed read: %d, write: %d\n", readcount, writecount))
	}

}

// ---------------------------------------------------------------------------------

//
//
//
func testListen(ctx context.Context, addr net.Addr) (l *LumerinListenStruct, e error) {

	return Listen(ctx, addr)
}

//
//
//
// func testDial(ctx context.Context, network LumProto, port int, ip net.IPAddr) (s *LumerinSocketStruct) {
func testDial(ctx context.Context, addr net.Addr) (s *LumerinSocketStruct) {

	// s, e := Dial(ctx, network, port, ip)
	s, e := Dial(ctx, addr)
	if e != nil {
		fmt.Printf(lumerinlib.FileLine()+" Dial Test Failed: %s\n", e)
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Dial Test Failed: %s\n", e))
	}

	return s
}

//
//
//
func goTestAcceptChannelEcho(l *LumerinListenStruct) {

	fmt.Printf(lumerinlib.FileLine() + " Waiting on Connection\n")

	s, e := l.Accept()

	if e != nil {
		fmt.Printf(lumerinlib.FileLine()+" Socket Accept() Failed: %s\n", e)
		l.Close()
		return
	}

	fmt.Printf(lumerinlib.FileLine() + " Connection Accepted\n")

	for {
		buf := make([]byte, 2)
		fmt.Printf(lumerinlib.FileLine() + " Read()ing\n")
		readcount, e := s.Read(buf)
		if e == io.EOF {
			fmt.Printf(lumerinlib.FileLine()+" Read() EOF count:%d\n", readcount)
			return
		}
		if e != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+" Read Failed: %s\n", e))
		}

		if readcount != 0 {
			buf = buf[:readcount]
			writecount, e := s.Write(buf)
			if e != nil {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" write Failed: %s\n", e))
			}
			if writecount == 0 {
				panic(fmt.Sprintf(lumerinlib.FileLine() + " write Failed: Zero bytes written\n"))
			}
		}
	}
}

type testAddr struct {
	ipaddr  string
	network string
}

func (t *testAddr) Network() string {
	return t.network
}

func (t *testAddr) String() string {
	return t.ipaddr
}
