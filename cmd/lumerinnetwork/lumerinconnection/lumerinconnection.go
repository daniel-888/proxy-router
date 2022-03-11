package lumerinconnection

import (
	"context"
	"errors"
	"fmt"
	"net"

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
	listener interface{}
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
			l = &LumerinListenStruct{
				ctx:      ctx,
				listener: tcp,
			}
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
// reads the acceptChan for new connections, or the channel closure
//
func (ll *LumerinListenStruct) Accept() (lci *LumerinSocketStruct, e error) {

	contextlib.Logf(ll.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch ll.listener.(type) {
	case *sockettcp.ListenTCPStruct:
		tcp := ll.listener.(*sockettcp.ListenTCPStruct)
		var soc *sockettcp.SocketTCPStruct
		soc, e = tcp.Accept()
		if e == nil {
			lci = &LumerinSocketStruct{
				ctx:    ll.ctx,
				socket: soc,
			}
		}
	default:
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Type:'%T' not supported\n", ll.listener))
	}
	return lci, e
}

//
//
//
func (ll *LumerinListenStruct) Cancel() {

	contextlib.Logf(ll.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch ll.listener.(type) {
	case *sockettcp.ListenTCPStruct:
		tcp := ll.listener.(*sockettcp.ListenTCPStruct)
		tcp.Cancel()
	default:
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Type:'%T' not supported\n", ll.listener))
	}
}

//
//
//
func (ll *LumerinListenStruct) Close() (e error) {

	contextlib.Logf(ll.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch ll.listener.(type) {
	case *sockettcp.ListenTCPStruct:
		tcp := ll.listener.(*sockettcp.ListenTCPStruct)
		e = tcp.Close()
	default:
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Type:'%T' not supported\n", ll.listener))
	}
	return e
}

//
//
//
// func Dial(ctx context.Context, p LumProto, port int, ip net.IPAddr) (lci *LumerinSocketStruct, e error) {
func Dial(ctx context.Context, addr net.Addr) (lci *LumerinSocketStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

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
		var tcp *sockettcp.SocketTCPStruct
		tcp, e = sockettcp.Dial(ctx, string(lumproto), ipaddr)
		if e == nil {
			lci = &LumerinSocketStruct{
				ctx:    ctx,
				socket: tcp,
			}
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
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Protocol not implemented:%s", string(lumproto)))

	default:
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Proto:'%s' not supported\n", lumproto))
	}

	return lci, e
}

//
//
//
func (l *LumerinSocketStruct) ReadReady() (ready bool) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		ready = l.socket.(*sockettcp.SocketTCPStruct).ReadReady()
	default:
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Type:'%T' not supported\n", l.socket))
	}

	return ready
}

//
//
//
func (l *LumerinSocketStruct) Read(buf []byte) (int, error) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		return l.socket.(*sockettcp.SocketTCPStruct).Read(buf)
	default:
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Type:'%T' not supported\n", l.socket))
	}
}

//
//
//
func (l *LumerinSocketStruct) Write(buf []byte) (count int, e error) {

	contextlib.Logf(l.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		count, e = l.socket.(*sockettcp.SocketTCPStruct).Write(buf)
	default:
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Type:'%T' not supported\n", l.socket))
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

		// Fill in here
		stat = LumerinConnectionStatusStruct{}

	default:
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Type:'%T' not supported\n", l.socket))
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
		e = l.socket.(*sockettcp.SocketTCPStruct).Close()
	default:
		panic(fmt.Sprintf(lumerinlib.FileLineFunc()+" Type:'%T' not supported\n", l.socket))
	}

	return e
}
