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

	ctx, cancel := context.WithCancel(context.Background())

	ip := net.IPAddr{
		IP: net.IP(net.IPv4(127, 0, 0, 1)),
	}

	l, e := testListen(ctx, TCP, 12345, ip)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Listen() Failed: %s\n", e))
	}

	cancel()

	_, e = l.Accept()
	if e != nil {
		select {
		case <-ctx.Done():
			fmt.Printf(lumerinlib.FileLine()+" CTX Done(): %s\n", ctx.Err())
		default:
			t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Accept() Test Failed: %s", e))
		}
	}

}

//
//
//
func TestDial(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel
	var TestString = "This is a test string\n"

	ip := net.IPAddr{
		IP: net.IP(net.IPv4(127, 0, 0, 1)),
	}

	l, e := testListen(ctx, TCP, 12345, ip)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Listen() Failed: %s\n", e))
	}

	go goTestAcceptChannelEcho(l)

	s := testDial(ctx, "tcp", 12345, ip)

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

//
//
//
func testListen(ctx context.Context, network LumProto, port int, ip net.IPAddr) (l *LumerinListenStruct, e error) {

	return Listen(ctx, network, port, ip)
}

//
//
//
func testDial(ctx context.Context, network LumProto, port int, ip net.IPAddr) (s *LumerinSocketStruct) {

	s, e := Dial(ctx, "tcp", port, ip)
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
