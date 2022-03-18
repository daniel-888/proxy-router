package lumerinconnection

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/sockettcp"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

//
// Lumerinconnection is a package that can interpret connection requests to see if
// there is already a Lumerin Protocol Trunk in place with the target node
// and redirect the connection across the trunk instead of a new direct connection.
//
//

const LumerinAcceptChannelLen int = 2
const LumerinReadChannelLen int = 10

var ErrLumConListenClosed = errors.New("Lumerin Connection Listen Socket closed")

type LumProto string

const TCP LumProto = "tcp"
const TCP4 LumProto = "tcp4"
const TCP6 LumProto = "tcp6"
const UDP LumProto = "udp"
const UDP4 LumProto = "udp4"
const UDP6 LumProto = "udp6"
const TRUNK LumProto = "trunk"
const TCPTRUNK LumProto = "tcptrunk"
const UDPTRUNK LumProto = "udptrunk"
const ANYAVAILABLE LumProto = "anyavailable"

//
// This will contain a regular socket or virtual socket structure
//
type LumerinListenStruct struct {
	ctx      context.Context
	cancel   func()
	listener interface{}
	accept   chan *LumerinSocketStruct
}

type LumerinSocketStruct struct {
	ctx    context.Context
	socket interface{}
}

//
// Do this here or at a lower layer?
//
type LumerinListenStatusStruct struct {
	connectionCount int
}

type LumerinConnectionStatusStruct struct {
	bytesRead    int
	bytesWritten int
	countRead    int
	countWrite   int
}

var ErrLumSocClosed = errors.New("lumerin socket: virt socket closed")

//
// Setup listening on the port/IP and on the Lumerin Port
// -- setsup the listening routine with cancel context
//
// Needs to be replaced with net.Addr
//
// func Listen(ctx context.Context, p LumProto, port int, ip net.IPAddr) (l *LumerinListenStruct, e error) {
func Listen(ctx context.Context, addr net.Addr) (l *LumerinListenStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	ctx, cancel := context.WithCancel(ctx)

	proto := addr.Network()
	ipaddr := addr.String()

	// Lots of error checking here

	lumproto := LumProto(proto)

	switch lumproto {
	case TCP:
		fallthrough
	case TCP4:
		fallthrough
	case TCP6:
		tcp, e := sockettcp.Listen(ctx, proto, ipaddr)
		if e == nil {
			accept := make(chan *LumerinSocketStruct)
			l = &LumerinListenStruct{
				ctx:      ctx,
				cancel:   cancel,
				listener: tcp,
				accept:   accept,
			}

			go l.goAccept()
		}

	case UDP:
		fallthrough
	case UDP4:
		fallthrough
	case UDP6:
		fallthrough
	case TRUNK:
		fallthrough
	case TCPTRUNK:
		fallthrough
	case UDPTRUNK:
		fallthrough
	case ANYAVAILABLE:
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Protocol not implemented:%s", string(lumproto))

	default:
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Proto:'%s' not supported\n", lumproto)
	}

	return l, e
}

//
// close() internal function to check to see if the listen socket has been canceled
//
func (ll *LumerinListenStruct) closed() bool {

	contextlib.Logf(ll.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	select {
	case <-ll.ctx.Done():
		return true
	default:
		return false
	}
}

//
// reads the acceptChan for new connections, or the channel closure
//
func (ll *LumerinListenStruct) goAccept() {

	contextlib.Logf(ll.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	defer close(ll.accept)

	for !ll.closed() {

		switch ll.listener.(type) {
		case *sockettcp.ListenTCPStruct:
			tcpchan := ll.listener.(*sockettcp.ListenTCPStruct).Accept()
			soc := <-tcpchan
			if soc == nil {
				contextlib.Logf(ll.ctx, contextlib.LevelWarn, lumerinlib.FileLineFunc()+" Accept() returned nil, assumed closed")
				break
			} else {
				lci := &LumerinSocketStruct{
					ctx:    ll.ctx,
					socket: soc,
				}
				ll.accept <- lci
			}

		default:
			contextlib.Logf(ll.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Default reached, type: %T", ll.listener)
		}

	}

	contextlib.Logf(ll.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Exiting...")
}

//
//
//
func (ll *LumerinListenStruct) Accept() <-chan *LumerinSocketStruct {

	contextlib.Logf(ll.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	return ll.accept

}

//
//
//
func (ll *LumerinListenStruct) Cancel() {
	contextlib.Logf(ll.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")
	ll.cancel()
}

//
//
//
func (ll *LumerinListenStruct) Close() (e error) {

	contextlib.Logf(ll.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch ll.listener.(type) {
	case *sockettcp.ListenTCPStruct:
		ll.listener.(*sockettcp.ListenTCPStruct).Close()
	default:
		contextlib.Logf(ll.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Default reached, type: %T", ll.listener)
	}
	return e
}

// ---------------------------------------------------------------------------------
//

//
//
//
func Dial(ctx context.Context, addr net.Addr) (lci *LumerinSocketStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	proto := addr.Network()
	ipaddr := addr.String()

	var split []string
	split = strings.Split(ipaddr, ":")
	if len(split) < 2 {
		contextlib.Logf(ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Split() returned len:%d:", len(split))
		return nil, fmt.Errorf(lumerinlib.FileLineFunc()+" Split returned len %d", len(split))
	}
	host := split[0]
	port := split[1]
	if port == "0" {
		contextlib.Logf(ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" port == 0")
		return nil, fmt.Errorf(lumerinlib.FileLineFunc() + " port == 0")
	}
	if host == "" {
		contextlib.Logf(ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" host == ''")
		return nil, fmt.Errorf(lumerinlib.FileLineFunc() + " host == ''")
	}

	// Lots of error checking here

	lumproto := LumProto(proto)

	switch lumproto {
	case TCP:
		fallthrough
	case TCP4:
		fallthrough
	case TCP6:
		var tcp *sockettcp.SocketTCPStruct
		tcp, e = sockettcp.Dial(ctx, string(lumproto), ipaddr)
		if e != nil {
			contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Dial() error returned: %s", e)
			return nil, e
		}
		lci = &LumerinSocketStruct{
			ctx:    ctx,
			socket: tcp,
		}

	case UDP:
		fallthrough
	case UDP4:
		fallthrough
	case UDP6:
		fallthrough
	case TRUNK:
		fallthrough
	case TCPTRUNK:
		fallthrough
	case UDPTRUNK:
		fallthrough
	case ANYAVAILABLE:
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Protocol not implemented yet: %s", lumproto)

	default:
		contextlib.Logf(ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Default reached, lumproto: %s", lumproto)
	}

	return lci, e
}

//
//
//
// func (l *LumerinSocketStruct) ReadReady() <-chan bool {
// 	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")
// 	return l.socket.(*sockettcp.SocketTCPStruct).ReadReady()
// }

//
//
//
func (l *LumerinSocketStruct) Read(buf []byte) (count int, e error) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		count, e = l.socket.(*sockettcp.SocketTCPStruct).Read(buf)
		if e != nil {
			contextlib.Logf(l.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Read() returned error: %s", e)
		}
	default:
		contextlib.Logf(l.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Default reached, type: %T", l.socket)
	}

	return count, e
}

//
//
//
func (l *LumerinSocketStruct) Write(buf []byte) (count int, e error) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		count, e = l.socket.(*sockettcp.SocketTCPStruct).Write(buf)
		if e != nil {
			contextlib.Logf(l.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Write() returned error: %s", e)
		}
	default:
		contextlib.Logf(l.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Default reached, type: %T", l.socket)
	}

	return count, e
}

//
//
//
func (l *LumerinSocketStruct) Status() (stat LumerinConnectionStatusStruct, e error) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		_, e = l.socket.(*sockettcp.SocketTCPStruct).Status()

		if e != nil {
			contextlib.Logf(l.ctx, contextlib.LevelError, lumerinlib.FileLineFunc()+" Status() returned error: %s", e)
		}
		// Fill in here
		stat = LumerinConnectionStatusStruct{}

	default:
		contextlib.Logf(l.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Default reached, type: %T", l.socket)
	}

	return stat, e
}

//
//
//
func (l *LumerinSocketStruct) Close() (e error) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		l.socket.(*sockettcp.SocketTCPStruct).Close()
	default:
		contextlib.Logf(l.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Default reached, type: %T", l.socket)
	}

	return e
}
