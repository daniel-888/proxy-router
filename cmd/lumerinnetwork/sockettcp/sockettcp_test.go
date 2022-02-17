package sockettcp

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"testing"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

var TestString = "This is the the test string\n"

func TestTCPSetupTestCancel(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	// addr := "127.0.0.1:12345"
	addr := ":12345"

	l, e := Listen(ctx, "tcp", addr)
	if e != nil {
		t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Listen() Test Failed: %s", e))
	}

	cancel()

	_, e = l.Accept()
	fmt.Printf(lumerinlib.FileLine()+" Accept() Returned:%s\n", e)

	if e != nil {
		select {
		case <-ctx.Done():
			fmt.Printf(lumerinlib.FileLine()+" CTX Done(): %s\n", ctx.Err())
		default:
			t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Accept() Test Failed: %s", e))
		}
	} else {
		t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Accept() Test Failed: %s", e))

	}

}

func TestTCPListenAddr(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	addr := ":55667"

	l, e := Listen(ctx, "tcp4", addr)
	if e != nil {
		t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Listen() Test Failed: network:%s", e))
	}

	netaddr, e := l.Addr()
	if e != nil {
		t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Addr() Test Failed: network:%s", e))
	}

	fmt.Printf(lumerinlib.FileLine()+" Addr() Returned %s:%s\n", netaddr.Network(), netaddr.String())

	if netaddr.Network() != "tcp" {
		t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Addr() Test Failed: network:%s", netaddr.Network()))
	}

	if netaddr.String() != "0.0.0.0:55667" {
		t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Addr() Test Failed: address:%s", netaddr.String()))
	}

	cancel()

}

func TestTCPSetupListener(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel

	addr := "127.0.0.1:12345"

	l, e := Listen(ctx, "tcp", addr)
	if e != nil {
		t.Fatalf("Listen() Test Failed")
	}

	go goTestAcceptChannelEcho(l)

	fmt.Printf(lumerinlib.FileLine() + " Dialing\n")

	s, e := Dial(ctx, "tcp", addr)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Dial Test Failed: %s", e))
	}

	fmt.Printf(lumerinlib.FileLine()+" Dial completed L:%s R:%s\n", s.LocalAddrString(), s.RemoteAddrString())

	writeb := []byte(TestString)

	writecount, e := s.Write(writeb)
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Read() Test Failed: %s\n", e))
	}
	if writecount != len(writeb) {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" Read() Test Failed: %s\n", e))
	}

	reader := bufio.NewReader(s)
	readbuf, e := reader.ReadBytes('\n')
	if e != nil {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+" ReadBytes() Test Failed: %s\n", e))
	}
	readcount := len(readbuf)
	if readcount != writecount {
		t.Fatal(fmt.Errorf(lumerinlib.FileLine()+"Count Test Failed read: %d, write: %d\n", readcount, writecount))
	}

	fmt.Printf("Done\n")
}

//
//
//
func goTestAcceptChannelEcho(l *ListenTCPStruct) {

	fmt.Printf(lumerinlib.FileLine() + " Waiting on Connection\n")

	s, e := l.Accept()

	if e != nil {
		fmt.Printf(lumerinlib.FileLine()+" Socket Accept() Failed: %s\n", e)
		l.Close()
		return
	}

	fmt.Printf(lumerinlib.FileLine()+" Accept() complete L:%s R:%s\n", s.LocalAddrString(), s.RemoteAddrString())

	for {
		buf := make([]byte, 2048)
		readcount, e := s.Read(buf)
		if e == io.EOF {
			fmt.Printf(lumerinlib.FileLine()+" Read() EOF count:%d\n", readcount)
			return
		}
		if e != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+" Read Failed: %s\n", e))
		}
		fmt.Printf(lumerinlib.FileLine()+" Read() done count:%d\n", readcount)

		if readcount != 0 {
			buf = buf[:readcount]
			writecount, e := s.Write(buf)
			if e != nil {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" write Failed: %s\n", e))
			}
			if writecount == 0 {
				panic(fmt.Sprintf(lumerinlib.FileLine() + " write Failed: Zero bytes written\n"))
			}

			fmt.Printf(lumerinlib.FileLine()+" Write() done count:%d\n", writecount)
		}
	}
}
