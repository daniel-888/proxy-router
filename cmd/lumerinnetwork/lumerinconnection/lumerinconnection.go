package lumerinconnection

import (
	"context"
	"errors"
	"fmt"
	"net"

	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/sockettcp"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

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

//
// This will contain a regular socket or virtual socket structure
//
type lumerinListenStruct struct {
	listener interface{}
}

type lumerinSocketStruct struct {
	socket interface{}
}

//
// Interface to either a socket, or virtual socket connection
//
//type LumerinConnectionStruct struct {
//	connType     LumProto
//	socketStruct interface{}
//	ctx          context.Context
//	cancel       func()
//	status       LumerinConnectionStatusStruct
//}

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

//var ErrLumSocTargetNotAvailable = errors.New("lumerin socket: target not available")
//var ErrLumSocTargetClosed = errors.New("lumerin socket: target closed")
//var ErrLumSocTargetNotResponding = errors.New("lumerin socket: Target not responding")
//var ErrLumSocTargetRejecting = errors.New("lumerin socket: Target Rejecting")
//var ErrLumSocListenAddrBusy = errors.New("lumerin socket: Listen Address Busy")

//
// Setup listening on the port/IP and on the Lumerin Port
// -- setsup the listening routine with cancel context
//
func Listen(ctx context.Context, p LumProto, port int, ip net.IPAddr) (l *lumerinListenStruct, e error) {

	// Error Checking here

	// ctx, cancel := context.WithCancel(ctx)

	ipaddr := fmt.Sprintf("%s:%d", ip.String(), port)

	// Parse different kinds of listeners here

	switch p {
	case TCP:
		fallthrough
	case TCP4:
		fallthrough
	case TCP6:
		tcp, e := sockettcp.Listen(ctx, string(p), ipaddr)
		if e == nil {
			l = &lumerinListenStruct{
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
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Protocol not implemented:%s", string(p)))

	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+":"+lumerinlib.Funcname()+" Proto:'%s' not supported\n", p))

	}

	return l, e
}

//
// reads the acceptChan for new connections, or the channel closure
//
func (ll *lumerinListenStruct) Accept() (lci *lumerinSocketStruct, e error) {

	switch ll.listener.(type) {
	case *sockettcp.ListenTCPStruct:
		tcp := ll.listener.(*sockettcp.ListenTCPStruct)
		soc, e := tcp.Accept()
		if e == nil {
			lci = &lumerinSocketStruct{
				socket: soc,
			}
		}
	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+":"+lumerinlib.Funcname()+" Type:'%T' not supported\n", ll.listener))
	}
	return lci, e
}

func (ll *lumerinListenStruct) Close() (e error) {

	switch ll.listener.(type) {
	case sockettcp.ListenTCPStruct:
		tcp := ll.listener.(sockettcp.ListenTCPStruct)
		e = tcp.Close()
	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+":"+lumerinlib.Funcname()+" Type:'%T' not supported\n", ll.listener))
	}
	return e
}

//
//
//
func Dial(ctx context.Context, p LumProto, addr string) (lci *lumerinSocketStruct, e error) {

	switch p {
	case TCP:
		fallthrough
	case TCP4:
		fallthrough
	case TCP6:
		var tcp *sockettcp.SocketTCPStruct
		tcp, e = sockettcp.Dial(ctx, string(p), addr)
		if e == nil {
			lci = &lumerinSocketStruct{
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
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Protocol not implemented:%s", string(p)))

	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+":"+lumerinlib.Funcname()+" Proto:'%s' not supported\n", p))
	}

	return lci, e
}

//
//
//
func (l *lumerinSocketStruct) ReadReady() (ready bool) {

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		ready = l.socket.(*sockettcp.SocketTCPStruct).ReadReady()
	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+":"+lumerinlib.Funcname()+" Type:'%T' not supported\n", l.socket))
	}

	return ready
}

//
//
//
func (l *lumerinSocketStruct) Read(buf []byte) (int, error) {

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		return l.socket.(*sockettcp.SocketTCPStruct).Read(buf)
	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+":"+lumerinlib.Funcname()+" Type:'%T' not supported\n", l.socket))
	}
}

//
//
//
func (l *lumerinSocketStruct) Write(buf []byte) (count int, e error) {

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		count, e = l.socket.(*sockettcp.SocketTCPStruct).Write(buf)
	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+":"+lumerinlib.Funcname()+" Type:'%T' not supported\n", l.socket))
	}

	return count, e
}

//
//
//
func (l *lumerinSocketStruct) Status() (stat LumerinConnectionStatusStruct, e error) {

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		_, e = l.socket.(*sockettcp.SocketTCPStruct).Status()

		// Fill in here
		stat = LumerinConnectionStatusStruct{}

	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+":"+lumerinlib.Funcname()+" Type:'%T' not supported\n", l.socket))
	}

	return stat, e
}

//
//
//
func (l *lumerinSocketStruct) Close() (e error) {

	switch l.socket.(type) {
	case *sockettcp.SocketTCPStruct:
		e = l.socket.(*sockettcp.SocketTCPStruct).Close()
	default:
		panic(fmt.Sprintf(lumerinlib.FileLine()+":"+lumerinlib.Funcname()+" Type:'%T' not supported\n", l.socket))
	}

	return e
}
