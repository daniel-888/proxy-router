package sockettcp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"regexp"
	"strconv"
	"time"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

//
// Socket TCP
// Manages TCP socket connections acting as standard IO Read/Write package
//
//

const TCPAcceptChannelLen int = 2
const TCPReadChannelLen int = 10
const TCPReadReadyChannelLen int = 1
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
	socket    net.Conn
	ctx       context.Context
	cancel    func()
	readready chan bool
	readchan  chan readStruct
	readbuf   []byte // Way point for read data from the socket
	status    SocketStatusStruct
}

type readStruct struct {
	err error
	buf []byte
}

//
//
//
func (r *readStruct) Err() error {
	return r.err
}

//
//
//
func (r *readStruct) Buf() []byte {
	return r.buf
}

//
// Opens a listening socket on the port and IP address of the local system
//
func Listen(ctx context.Context, network string, addr string) (l *ListenTCPStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

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

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	<-l.ctx.Done()
	e := l.Close()
	if e != nil {
		contextlib.Logf(l.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" close returned error")
	}
}

//
// goAccept() go routine to accept connections and return new socket structs to the Accept() function
//
func (l *ListenTCPStruct) goAccept() {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	defer close(l.accept)

	for !l.closed() {
		conn, e := l.listener.Accept()

		if e != nil {
			select {
			case <-l.ctx.Done():
				contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" closed")
			default:
				contextlib.Logf(l.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Accept() returned error:%s", e)
			}
			return
		}

		if conn == nil {
			contextlib.Logf(l.ctx, contextlib.LevelWarn, lumerinlib.FileLineFunc()+" Accept() returned empty connection")
			return
		}

		newsoc := createNewSocket(l.ctx, conn)

		l.accept <- newsoc
	}

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" exiting")
}

//
// Blocking call to Accept and open a new connection
//
func (l *ListenTCPStruct) Accept() <-chan *SocketTCPStruct {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	return l.accept

}

//
// close() internal function to check to see if the listen socket has been canceled
//
func (l *ListenTCPStruct) closed() bool {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

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

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	return l.listener.Close()
}

//
//
//
func (l *ListenTCPStruct) Cancel() {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	l.cancel()
}

//
// Returns current status of the Listener
//
func (l *ListenTCPStruct) Addr() (addr net.Addr, e error) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	addr = l.listener.Addr()

	return addr, e
}

//
// Returns address of the Listener
//
func (l *ListenTCPStruct) Status() (ltss ListenerStatusStruct, e error) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	ltss = l.status

	return ltss, e
}

//
// Dial() creates a new TCP connection to the target address
// or returns an error
//
func Dial(ctx context.Context, network string, addr string) (s *SocketTCPStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

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

	//
	// CONTEXT Funkyness Here
	//
	cs := contextlib.GetContextStruct(ctx)
	// dialctx, cancel := context.WithCancel(ctx)
	// dialctx, cancel = context.WithTimeout(dialctx, time.Minute)
	dialctx, cancel := context.WithTimeout(ctx, time.Minute)
	dialctx = context.WithValue(dialctx, contextlib.ContextKey, cs)

	var conn net.Conn
	var d net.Dialer
	conn, e = d.DialContext(dialctx, network, addr)

	if e != nil {
		cancel()
	} else {
		ctx = context.WithValue(ctx, contextlib.ContextKey, cs)
		s = createNewSocket(ctx, conn)
	}

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Complete")

	return s, e
}

//
// Go Routine to listen to the context for cancel
//
func createNewSocket(ctx context.Context, conn net.Conn) (soc *SocketTCPStruct) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	rr := make(chan bool, TCPReadReadyChannelLen)
	rc := make(chan readStruct, TCPReadChannelLen)
	ctx, cancel := contextlib.CreateNewContext(ctx)
	soc = &SocketTCPStruct{
		socket:    conn,
		ctx:       ctx,
		cancel:    cancel,
		readready: rr,
		readchan:  rc,
		readbuf:   make([]byte, 0, TCPReadBufferSize),
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

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	<-s.ctx.Done()
	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" shutting down socket")
	e := s.Close()
	if e != nil {
		contextlib.Logf(s.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Close() returned Error:%s", e)
	}

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" exiting")
}

//
// close() internal function to check to see if the socket has been canceled
//
func (s *SocketTCPStruct) closed() (ret bool) {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	select {
	case <-s.ctx.Done():
		ret = true
	default:
		ret = false
	}

	return ret
}

//
// goRead() go routine to read the socket, buffer it and send it to the Read() function
//
func (s *SocketTCPStruct) goRead() {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	defer close(s.readchan)
	defer close(s.readready)

	for !s.closed() {

		contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" enter for loop")

		var e error
		buf := make([]byte, TCPReadBufferSize)
		readcount, e := s.socket.Read(buf)

		contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+"Socket Read buf count:%d", readcount)

		if e != nil {
			if e == io.EOF {
				s.Close()
			} else {
				select {
				case <-s.ctx.Done():
					e = ErrSocTCPClosed
				default:
					contextlib.Logf(s.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Read() returned error: %s", e)
				}
			}
		} else if readcount == 0 {
			e = fmt.Errorf("soc.goRead() returned 0 bytes")
		}

		r := readStruct{
			err: e,
			buf: buf[:readcount],
		}

		s.readchan <- r
		if len(s.readready) == 0 {
			s.readready <- true
		}
	}

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" exiting")

}

//
// Readready() Non-blocking call to see if a call to Read() would block or not
//
func (s *SocketTCPStruct) ReadReady() <-chan bool {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	return s.readready

}

//
// Read()
// Blocks on getting a read back from readchan.
//
func (s *SocketTCPStruct) Read(buf []byte) (count int, e error) {

	count = 0

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	if cap(buf) == 0 {
		contextlib.Logf(s.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" buf is zero lenth")
		return 0, fmt.Errorf(lumerinlib.FileLineFunc() + " buffer lenth is 0")
	}

	if s.closed() {
		return 0, ErrSocTCPClosed
	}

	readbufsize := len(s.readbuf)
	readchansize := len(s.readchan)
	// NAND - only run the loop if these two condtions are met.
	// Dont go into the loop if readbuf > 0 and readchansize is 0
	// because there is no incoming reads, and there is data in the buffer
	if !(readbufsize > 0 && readchansize == 0) {

		// s.readbuf is a local storage for reads to buffer incoming data
		// The local buffer will reach a high water mark and just start returning data
		// The buf will be filled or the s.readbuf will be emptied

	FORLOOP:
		for {
			select {
			// Exit, we are done
			case <-s.ctx.Done():
				count = 0
				e = s.ctx.Err()
				break FORLOOP

			case r := <-s.readchan:
				// Exit, Socket returned an error
				if r.Err() != nil {
					count = 0
					e = r.Err()
					break FORLOOP
				}

				s.readbuf = append(s.readbuf, r.buf...)

				if len(s.readbuf) > TCPReadBufferSize {
					break FORLOOP
				}

				readbufsize = len(s.readbuf)
				readchansize = len(s.readchan)
				if readbufsize > 0 && readchansize == 0 {
					break FORLOOP
				}
			}
		}
	}

	if e == nil {

		// Drain s.readbuf into the buf

		if len(s.readbuf) > 0 {
			count = copy(buf, s.readbuf)
			s.readbuf = s.readbuf[count:]
		}

		s.status.bytesRead += count
		s.status.countRead++
	}

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" return count:%d", count)

	return count, e
}

//
//
//
func (s *SocketTCPStruct) Write(buf []byte) (count int, e error) {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	if len(buf) == 0 {
		panic(fmt.Errorf(lumerinlib.FileLineFunc() + " Write() buffer lenth is zero"))
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

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

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

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	s.cancel() // For good measure

	return s.socket.Close()
}

//
// Returns the local address of the socket
//
func (s *SocketTCPStruct) LocalAddrString() string {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	return s.socket.LocalAddr().String()
}

//
// Returns the remote address of the socket
//
func (s *SocketTCPStruct) RemoteAddrString() string {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	return s.socket.RemoteAddr().String()
}

//
// Returns the local address of the socket
//
func (l *ListenTCPStruct) LocalAddr() (host string, port int, e error) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	addr := l.listener.Addr().String()
	host, port, e = getAddr(l.ctx, addr)
	return
}

//
// Returns the local address of the socket
//
func (s *SocketTCPStruct) LocalAddr() (host string, port int, e error) {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	return getAddr(s.ctx, s.socket.LocalAddr().String())
}

//
// Returns the local address of the socket
//
func (s *SocketTCPStruct) RemoteAddr() (host string, port int, e error) {

	contextlib.Logf(s.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	return getAddr(s.ctx, s.socket.RemoteAddr().String())
}

//
// Returns the local address of the socket
//
func getAddr(ctx context.Context, addr string) (host string, port int, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	regex := regexp.MustCompile("^(\\[*[a-fA-F0-9:]+\\]*):(\\d+)$")

	regexret := regex.FindStringSubmatch(addr)
	_ = regexret

	host = regexret[1]
	portstr := regexret[2]

	port, e = strconv.Atoi(portstr)

	return host, port, e
}
