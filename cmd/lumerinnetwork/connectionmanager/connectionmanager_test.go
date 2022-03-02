package connectionmanager

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"testing"

	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/lumerinconnection"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

var TestString = "This is a test string\n"

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

	// defer l.Close()
	l.Cancel()

	_, e = l.Accept()
	if e != nil {
		select {
		case <-ctx.Done():
			fmt.Printf(lumerinlib.FileLine()+" CTX Done(): %s\n", ctx.Err())
		default:
			fmt.Printf(lumerinlib.FileLine()+"Accept() OK: returned error: %s\n", e)
		}
	} else {
		t.Fatalf(fmt.Sprintf(lumerinlib.FileLine() + "Accept() Test Failed no error returned"))
	}

}

//
//
//
func TestSrcDial(t *testing.T) {

	ctx := context.Background()

	testaddr := &testAddr{
		network: "tcp",
		ipaddr:  "127.0.0.1:12346",
	}

	l, e := testListen(ctx, testaddr)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Listen() Failed: %s\n", e))
	}
	defer l.Cancel()

	go goTestAcceptChannelEcho(l)

	//
	// Dial (using lumerinconnection) the listener, write test data, recieve same test data
	//
	s, e := lumerinconnection.Dial(ctx, testaddr)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Dial() Failed: %s\n", e))
	}

	defer s.Close()

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
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" ReadBytes() Test Failed: %s\n", e))
	}
	readcount := len(readbuf)
	if readcount != writecount {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+"Count Test Failed read: %d, write: %d\n", readcount, writecount))
	}

}

//
//
//
func TestSrcDefDstDial(t *testing.T) {

	ctx := context.Background()

	testaddr := &testAddr{
		network: "tcp",
		ipaddr:  "127.0.0.1:12347",
	}

	l, e := testListen(ctx, testaddr)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Listen() Failed: %s\n", e))
	}

	defer l.Close()

	s := testSetupEchoConnection(t, l)

	defer s.Close()

	writeb := []byte(TestString)
	writecount, e := s.DstWrite(writeb)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" DstWrite() Test Failed: %s\n", e))
	}
	if writecount != len(writeb) {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" DstWrite() Test Failed: %s\n", e))
	}

	fmt.Printf(lumerinlib.FileLine() + " DstWrite() completed\n")

	soc, e := s.DstGetSocket()
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" DstGetSocket() Test Failed: %s\n", e))
	}
	reader := bufio.NewReader(soc)
	readbuf, e := reader.ReadBytes('\n')
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" ReadBytes() Test Failed: %s\n", e))
	}
	readcount := len(readbuf)
	if readcount != writecount {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+"Count Test Failed read: %d, write: %d\n", readcount, writecount))
	}

}

//
//
//
func TestSrcIdxDstDial(t *testing.T) {

	ctx := context.Background()

	testaddr := &testAddr{
		network: "tcp",
		ipaddr:  "127.0.0.1:12348",
	}

	l, e := testListen(ctx, testaddr)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Listen() Failed: %s\n", e))
	}

	defer l.Close()

	s := testSetupEchoConnection(t, l)

	defer s.Close()

	writeb := []byte(TestString)
	writecount, e := s.IdxWrite(0, writeb)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" IdxWrite() Test Failed: %s\n", e))
	}
	if writecount != len(writeb) {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" IdxWrite() Test Failed: %s\n", e))
	}

	fmt.Printf(lumerinlib.FileLine() + " IdxWrite() completed\n")

	soc, e := s.IdxGetSocket(0)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" DstGetSocket() Test Failed: %s\n", e))
	}
	defer soc.Close()
	reader := bufio.NewReader(soc)
	readbuf, e := reader.ReadBytes('\n')
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" ReadBytes() Test Failed: %s\n", e))
	}
	readcount := len(readbuf)
	if readcount != writecount {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+"Count Test Failed read: %d, write: %d\n", readcount, writecount))
	}

}

// ---------------------------------------------------------------------------------------------------

//
//
// func testListen(ctx context.Context, port int, ip net.IPAddr) (l *ConnectionListenStruct, e error) {
func testListen(ctx context.Context, addr net.Addr) (l *ConnectionListenStruct, e error) {
	return Listen(ctx, addr)
}

//
//
//
func testAcceptChannelEcho(l *ConnectionListenStruct) (s *ConnectionStruct) {

	fmt.Printf(lumerinlib.FileLine() + " Waiting on Connection\n")

	s, e := l.Accept()

	if e != nil {
		fmt.Printf(lumerinlib.FileLine()+" Socket Accept() Failed: %s\n", e)
		l.Close()
		return
	}

	return s
}

//
//
//
func testSetupEchoConnection(t *testing.T, l *ConnectionListenStruct) (cs *ConnectionStruct) {

	fmt.Printf(lumerinlib.FileLine() + " Waiting on Connection\n")

	lss, e := lumerinconnection.Dial(l.ctx, l.addr)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Dial() Failed: %s\n", e))
	}

	cs, e = l.Accept()

	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Accept() Failed: %s\n", e))
	}

	fmt.Printf(lumerinlib.FileLine() + " Connection Accepted\n")

	cs.dst = append(cs.dst, lss)
	cs.defidx = 0

	go cs.goSrcChannelEcho()

	return cs
}

//
//
//
func goTestAcceptChannelEcho(l *ConnectionListenStruct) {

	fmt.Printf(lumerinlib.FileLine() + " Waiting on Connection\n")

	s, e := l.Accept()

	if e != nil {
		fmt.Printf(lumerinlib.FileLine()+" Socket Accept() Failed: %s\n", e)
		l.Close()
		return
	}

	fmt.Printf(lumerinlib.FileLine() + " Connection Accepted\n")

	s.goSrcChannelEcho()
}

//
//
//
func (s *ConnectionStruct) goSrcChannelEcho() {

	fmt.Printf(lumerinlib.FileLine() + " SRC Echo\n")

	for {
		buf := make([]byte, 2)
		fmt.Printf(lumerinlib.FileLine() + " Read()ing\n")
		readcount, e := s.SrcRead(buf)
		if e == io.EOF {
			fmt.Printf(lumerinlib.FileLine()+" Read() EOF count:%d\n", readcount)
			return
		}
		if e != nil {
			select {
			case <-s.ctx.Done():
				return
			default:
				panic(fmt.Sprintf(lumerinlib.FileLine()+" Read Failed: %s\n", e))
			}
		}

		if readcount != 0 {
			buf = buf[:readcount]
			writecount, e := s.SrcWrite(buf)
			if e != nil {
				select {
				case <-s.ctx.Done():
					return
				default:
					panic(fmt.Sprintf(lumerinlib.FileLine()+" write Failed: %s\n", e))
				}
			}
			if writecount == 0 {
				select {
				case <-s.ctx.Done():
					return
				default:
					panic(fmt.Sprintf(lumerinlib.FileLine() + " write Failed: Zero bytes written\n"))
				}
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
