package sockettcp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

//
// Socket TCP
// Manages TCP socket connections acting as standard IO Read/Write package
//
//

const TCPAcceptChannelLen int = 2
const TCPReadChannelLen int = 10
const TCPReadBufferSize int = 1024

var ErrSocTCPClosed = errors.New("socket TCP: socket closed")
var ErrSocTCPBadNetwork = errors.New("socket TCP: bad network protocol")

//var ErrSocTCPReadCopyUnderRun = errors.New("socke TCP: Read() Copy Under Run")
//var ErrSocTCPTargetNotResponding = errors.New("socke TCP: Target not responding")
//var ErrSocTCPTargetRejecting = errors.New("socke TCP: Target Rejecting")
//var ErrSocTCPListenAddrBusy = errors.New("socke TCP: Listen Port Address Busy")
//var ErrSocTCPIPAddrBusy = errors.New("socket TCP: Listen IP Address Busy")

//
// ----------------
//

type SocketStatusStruct struct {
	bytesRead    int
	bytesWritten int
	countRead    int
	countWrite   int
}

type ListenerStatusStruct struct {
	connectionCount int
}

//
// ----------------
//

type ListenTCPStruct struct {
	listener net.Listener
	ctx      context.Context
	cancel   func()
	accept   chan *SocketTCPStruct
	status   ListenerStatusStruct
}

type SocketTCPStruct struct {
	socket   net.Conn
	ctx      context.Context
	cancel   func()
	readchan chan readStruct
	readbuf  []byte
	status   SocketStatusStruct
}

type readStruct struct {
	err error
	buf []byte
}

func (r *readStruct) Err() error {
	return r.err
}

func (r *readStruct) Buf() []byte {
	return r.buf
}

//
// Opens a listening socket on the port and IP address of the local system
//
func Listen(ctx context.Context, network string, addr string) (l *ListenTCPStruct, e error) {

	//
	// network can be tcp, tcp4 or tcp6
	//
	switch network {
	case "tcp":
	case "tcp4":
	case "tcp6":
	default:
		return l, ErrSocTCPBadNetwork
	}

	ctx, cancel := context.WithCancel(ctx)
	c := cancel

	listenconfig := net.ListenConfig{}
	listener, e := listenconfig.Listen(ctx, network, addr)

	if e != nil {
		return l, e
	}

	accept := make(chan *SocketTCPStruct, TCPAcceptChannelLen)

	l = &ListenTCPStruct{
		listener: listener,
		ctx:      ctx,
		cancel:   c,
		accept:   accept,
		status: ListenerStatusStruct{
			connectionCount: 0,
		},
	}

	//
	// Function to listen to the context for cancel
	//

	go l.goWaitOnCancel()

	//
	// Go routine to run accept() on the socket and pass the connection back
	// via the channel.  The accept function can be canceled via the context
	//
	go l.goAccept()

	return l, e
}

//
// goWaitOnCancel() Go Routine to listen to the context for cancel and close the socket
//
func (l *ListenTCPStruct) goWaitOnCancel() {
	<-l.ctx.Done()
	e := l.Close()
	if e != nil {
		fmt.Printf(lumerinlib.Funcname()+"Close() returned error %s\n", e)
	}
}

//
// goAccept() go routine to accept connections and return new socket structs to the Accept() function
//
func (l *ListenTCPStruct) goAccept() {
	defer close(l.accept)

	for !l.closed() {
		conn, e := l.listener.Accept()

		if e != nil {
			select {
			case <-l.ctx.Done():
				fmt.Printf("soc.Accpet() Closed\n")
			default:
				fmt.Printf("soc.Accpet() returned error: %s\n", e)
			}
			return
		}

		if conn == nil {
			fmt.Printf("soc.Accpet() returned empty connection\n")
			return
		}

		newsoc := createNewSocket(l.ctx, conn)

		l.accept <- newsoc
	}

}

//
// Blocking call to Accept and open a new connection
//
func (l *ListenTCPStruct) Accept() (s *SocketTCPStruct, e error) {

	select {
	case s := <-l.accept:
		l.status.connectionCount++
		return s, e
	case <-l.ctx.Done():
		return s, ErrSocTCPClosed
	}

}

//
// close() internal function to check to see if the listen socket has been canceled
//
func (l *ListenTCPStruct) closed() bool {
	select {
	case <-l.ctx.Done():
		return true
	default:
		return false
	}
}

//
// Closes down a listening Socket
//
func (l *ListenTCPStruct) Close() error {
	return l.listener.Close()
}

//
//
//
func (l *ListenTCPStruct) Cancel() {
	l.cancel()
}

//
// Returns current status of the Listener
//
func (l *ListenTCPStruct) Addr() (addr net.Addr, e error) {

	addr = l.listener.Addr()

	return addr, e
}

//
// Returns address of the Listener
//
func (l *ListenTCPStruct) Status() (ltss ListenerStatusStruct, e error) {

	ltss = l.status

	return ltss, e
}

//
// Dial() creates a new TCP connection to the target address
// or returns an error
//
func Dial(ctx context.Context, network string, addr string) (s *SocketTCPStruct, e error) {
	var d net.Dialer

	fmt.Printf(lumerinlib.Funcname() + " enter func\n")

	//
	// network can be tcp, tcp4 or tcp6
	//
	switch network {
	case "tcp":
	case "tcp4":
	case "tcp6":
	default:
		return s, ErrSocTCPBadNetwork
	}

	// Error: Review the address Here

	dialctx, cancel := context.WithTimeout(ctx, time.Minute)

	var conn net.Conn
	conn, e = d.DialContext(dialctx, network, addr)

	if e != nil {
		cancel()
		return s, e
	}

	s = createNewSocket(ctx, conn)

	fmt.Printf(lumerinlib.Funcname() + " Finished\n")

	return s, e
}

//
// Go Routine to listen to the context for cancel
//
func createNewSocket(ctx context.Context, conn net.Conn) (soc *SocketTCPStruct) {

	rc := make(chan readStruct, TCPReadChannelLen)
	ctx, cancel := context.WithCancel(ctx)
	soc = &SocketTCPStruct{
		socket:   conn,
		ctx:      ctx,
		cancel:   cancel,
		readchan: rc,
		readbuf:  make([]byte, 0),
		status: SocketStatusStruct{
			bytesRead:    0,
			bytesWritten: 0,
			countRead:    0,
			countWrite:   0,
		},
	}

	go soc.goWaitOnCancel()
	go soc.goRead()

	return soc
}

//
// Go Routine to listen to the context for cancel
//
func (s *SocketTCPStruct) goWaitOnCancel() {

	fmt.Printf(lumerinlib.Funcname() + " enter func\n")

	<-s.ctx.Done()
	log.Println("... shutting down socket")
	e := s.Close()
	if e != nil {
		fmt.Printf(lumerinlib.FileLine()+" Close() returned error %s\n", e)
	}

	fmt.Printf(lumerinlib.Funcname() + " exit func\n")
}

//
// close() internal function to check to see if the socket has been canceled
//
func (s *SocketTCPStruct) closed() bool {
	select {
	case <-s.ctx.Done():
		return true
	default:
		return false
	}
}

//
// goRead() go routine to read the socket, buffer it and send it to the Read() function
//
func (s *SocketTCPStruct) goRead() {
	defer close(s.readchan)

	fmt.Printf(lumerinlib.Funcname() + " enter func\n")

	for !s.closed() {

		fmt.Printf(lumerinlib.Funcname() + " enter for loop\n")

		var e error
		buf := make([]byte, TCPReadBufferSize)
		readcount, e := s.socket.Read(buf)

		fmt.Printf(lumerinlib.Funcname()+" Read() count:%d\n", readcount)

		if e != nil {
			var err error = nil
			if e == io.EOF {
				err = e
				s.Close()
			} else {
				select {
				case <-s.ctx.Done():
					err = ErrSocTCPClosed
				default:
					err = fmt.Errorf("soc.goRead() returned error: %s", e)
				}
			}

			r := readStruct{
				err: err,
				buf: buf[:readcount],
			}
			s.readchan <- r
			return
		}

		if readcount == 0 {
			e = fmt.Errorf("soc.goRead() returned 0 bytes")
		}

		r := readStruct{
			err: e,
			buf: buf[:readcount],
		}
		s.readchan <- r
	}

	fmt.Printf(lumerinlib.Funcname() + " exit func\n")

}

//
// Readready() Non-blocking call to see if a call to Read() would block or not
//
func (s *SocketTCPStruct) ReadReady() (ready bool) {

	if s.closed() {
		return false
	}

	if len(s.readchan) != 0 || len(s.readbuf) != 0 {
		return true
	}

	return false
}

//
// Read()
// Manages the read buffer, adding to it from the readchannel if the channel is ready
// and filling the return buffer with what ever will fit from the read buffer.
// The subroutine will block if there is nothing in the readchannel and the read buffer.
//
func (s *SocketTCPStruct) Read(buf []byte) (count int, e error) {

	if cap(buf) == 0 {
		panic(fmt.Errorf(lumerinlib.FileLine() + " Read() buffer capacity is zero"))
	}

	if s.closed() {
		return 0, ErrSocTCPClosed
	}

	if !s.ReadReady() {
		select {
		case <-s.ctx.Done():
			count = 0
			e = s.ctx.Err()
			return count, e

		case r := <-s.readchan:
			if r.Err() != nil {
				count = 0
				e = r.Err()
				return count, e
			} else {
				s.readbuf = append(s.readbuf, r.buf...)
			}
		}
	}

FORLOOP:
	for {
		select {
		case <-s.ctx.Done():
			count = 0
			e = s.ctx.Err()
			return count, e

		case r := <-s.readchan:
			if r.Err() != nil {
				count = 0
				e = r.Err()
				return count, e
			} else {
				s.readbuf = append(s.readbuf, r.buf...)
			}

		default:
			break FORLOOP
		}
	}

	bufcap := cap(buf)
	readbuflen := len(s.readbuf)

	if readbuflen == 0 {
		return 0, nil
	}

	if readbuflen > bufcap {
		count = copy(buf, s.readbuf)
		s.readbuf = s.readbuf[count:]
	} else {
		count = copy(buf, s.readbuf)
		s.readbuf = s.readbuf[:0]
	}

	s.status.bytesRead += count
	s.status.countRead++

	return count, e
}

//
//
//
func (s *SocketTCPStruct) Write(buf []byte) (count int, e error) {

	if len(buf) == 0 {
		panic(fmt.Errorf(lumerinlib.FileLine() + " Write() buffer lenth is zero"))
	}

	if s.closed() {
		return 0, ErrSocTCPClosed
	}

	count, e = s.socket.Write(buf)

	if e == nil {
		s.status.bytesWritten += count
		s.status.countWrite++
	}

	return count, e
}

//
//
//
func (s *SocketTCPStruct) Status() (ss SocketStatusStruct, e error) {

	if s.closed() {
		e = ErrSocTCPClosed
	}

	ss = s.status
	return ss, e
}

//
//
//
func (s *SocketTCPStruct) Close() error {

	s.cancel() // For good measure

	return s.socket.Close()
}

//
// Returns the local address of the socket
//
func (s *SocketTCPStruct) LocalAddrString() string {
	return s.socket.LocalAddr().String()
}

//
// Returns the remote address of the socket
//
func (s *SocketTCPStruct) RemoteAddrString() string {
	return s.socket.RemoteAddr().String()
}
