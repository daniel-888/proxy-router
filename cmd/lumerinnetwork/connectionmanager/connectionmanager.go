package connectionmanager

import (
	"context"
	"errors"
	"fmt"
	"net"

	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/lumerinconnection"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

const DefaultDstSlots int = 8

var ErrConnMgrNoDefIndex = errors.New("CM: no default index")
var ErrConnMgrBadDefDest = errors.New("CM: bad default destination")
var ErrConnMgrIDXOutOfRange = errors.New("CM: index out of range")

//
// Listen Struct for new SRC connections coming in
//
type ConnectionListenStruct struct {
	listen *lumerinconnection.LumerinListenStruct
	ctx    context.Context
	cancel func()
	port   int
	addr   net.Addr
}

//
// Struct for existing SRC connections and the associated outgoing DST connections
type ConnectionStruct struct {
	src    *lumerinconnection.LumerinSocketStruct
	dst    []*lumerinconnection.LumerinSocketStruct
	defidx int
	ctx    context.Context
	cancel func()
}

//
//
//
func Listen(ctx context.Context) (cls *ConnectionListenStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLine()+" called")

	ctx, cancel := context.WithCancel(ctx)
	_ = cancel

	addr := contextlib.GetSrc(ctx)

	l, e := lumerinconnection.Listen(ctx, addr)
	if e == nil {
		cls = &ConnectionListenStruct{
			listen: l,
			ctx:    ctx,
			cancel: cancel,
			addr:   addr,
		}
	}

	return cls, e
}

//
//
//
func (cls *ConnectionListenStruct) getPort() (port int) {
	contextlib.Logf(cls.ctx, contextlib.LevelPanic, fmt.Sprint(lumerinlib.FileLine()+" called"))
	return 0
}

//
//
//
func (cls *ConnectionListenStruct) getIp() net.Addr {
	return cls.addr
}

//
//
//
func (cls *ConnectionListenStruct) Accept() (cs *ConnectionStruct, e error) {

	contextlib.Logf(cls.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLine()+" called"))

	lci, e := cls.listen.Accept()
	if e != nil {
		return cs, e
	}

	cs = cls.newConnectionStruct(lci)
	return cs, e
}

//
//
//
func (cls *ConnectionListenStruct) Close() (e error) {
	contextlib.Logf(cls.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLine()+" called"))
	return cls.listen.Close()
}

//
//
//
func (cls *ConnectionListenStruct) Cancel() {
	contextlib.Logf(cls.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLine()+" called"))
	cls.cancel()
}

//
//
//
func (cls *ConnectionListenStruct) newConnectionStruct(srclss *lumerinconnection.LumerinSocketStruct) (cs *ConnectionStruct) {

	contextlib.Logf(cls.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLine()+" called"))

	ctx, cancel := context.WithCancel(cls.ctx)
	dstarrlss := make([]*lumerinconnection.LumerinSocketStruct, 0, DefaultDstSlots)

	cs = &ConnectionStruct{
		src:    srclss,
		dst:    dstarrlss,
		defidx: -1,
		ctx:    ctx,
		cancel: cancel,
	}

	return cs
}

//
// Dial() opens up a new dst connection and inserts it into the first avalable dst slot
// If this is the 0th slow, the default dst is set as well
//
// func (cs *ConnectionStruct) Dial(ctx context.Context, port int, ip net.IPAddr) (idx int, e error) {
func (cs *ConnectionStruct) Dial(ctx context.Context, addr net.Addr) (idx int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLine()+" called"))

	idx = -1

	// dst, e := lumerinconnection.Dial(ctx, lumerinconnection.TCP, port, ip)
	dst, e := lumerinconnection.Dial(ctx, addr)
	if e != nil {
		return idx, e
	}

	// find next available dst slot

	for i := 0; i < len(cs.dst); i++ {
		if cs.dst[i] == nil {
			idx = i
			cs.dst[i] = dst
			if cs.defidx < 0 {
				cs.defidx = i
			}

			return idx, e
		}
	}

	cs.dst = append(cs.dst, dst)
	idx = len(cs.dst) - 1

	return idx, nil
}

//
// ReDialIdx() will attempt to reconnect to the same dst, first checking the the line is closed
// It is used in case a connection is severed
//
func (cs *ConnectionStruct) ReDialIdx(idx int) (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelPanic, fmt.Sprintf(lumerinlib.FileLine()+" Function Not Implemented Yet.."))

	return nil
}

//
// Close() will close out all src and dst connections via the cancel context function
//
func (cs *ConnectionStruct) Close() (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	cs.cancel() // This should close all open src and dst connections

	return nil
}

//
//
//
func (cs *ConnectionStruct) SetRoute(idx int) (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if idx < 0 || idx >= len(cs.dst) {
		e = ErrConnMgrIDXOutOfRange
		return e
	}

	if cs.dst[idx] == nil {
		e = ErrConnMgrBadDefDest
		return e
	}

	cs.defidx = idx

	return nil

}

//
// ReadReady() checks all open connections to see if any are ready to read
//
func (cs *ConnectionStruct) ReadReady() (r bool) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if cs.src != nil && cs.src.ReadReady() {
		return true
	}

	for i := 0; i < len(cs.dst); i++ {
		if cs.dst[i] != nil && cs.dst[i].ReadReady() {
			return true
		}
	}

	return false
}

//
//
//
func (cs *ConnectionStruct) SrcGetSocket() (s *lumerinconnection.LumerinSocketStruct, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	return cs.src, nil
}

//
//
//
func (cs *ConnectionStruct) SrcReadReady() (r bool) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	return cs.src.ReadReady()
}

//
//
//
func (cs *ConnectionStruct) SrcRead(buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	return cs.src.Read(buf)
}

//
//
//
func (cs *ConnectionStruct) SrcWrite(buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	return cs.src.Write(buf)

}

//
// SrcClose() calls (*CS) Close() to close everything down
//
func (cs *ConnectionStruct) SrcClose() (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	return cs.Close()
}

//
//
//
func (cs *ConnectionStruct) DstGetSocket() (s *lumerinconnection.LumerinSocketStruct, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if cs.defidx < 0 {
		e = ErrConnMgrNoDefIndex
		return nil, e
	}

	if cs.dst[cs.defidx] == nil {
		e = ErrConnMgrBadDefDest
		return nil, e

	}
	return cs.dst[cs.defidx], e
}

//
//
//
func (cs *ConnectionStruct) DstReadReady() (r bool) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if cs.defidx >= 0 && cs.dst[cs.defidx] != nil && cs.dst[cs.defidx].ReadReady() {
		return true
	}
	return false
}

//
//
//
func (cs *ConnectionStruct) DstRead(buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if cs.defidx < 0 {
		e = ErrConnMgrNoDefIndex
		return 0, e
	}

	if cs.dst[cs.defidx] == nil {
		e = ErrConnMgrBadDefDest
		return 0, e

	}

	return cs.dst[cs.defidx].Read(buf)

}

//
//
//
func (cs *ConnectionStruct) DstWrite(buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if cs.defidx < 0 {
		e = ErrConnMgrNoDefIndex
		return 0, e
	}

	if cs.dst[cs.defidx] == nil {
		e = ErrConnMgrBadDefDest
		return 0, e

	}

	return cs.dst[cs.defidx].Write(buf)

}

//
//
//
func (cs *ConnectionStruct) DstClose() (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if cs.defidx < 0 {
		e = ErrConnMgrNoDefIndex
		return e
	}

	if cs.dst[cs.defidx] == nil {
		e = ErrConnMgrBadDefDest
		return e

	}

	return cs.dst[cs.defidx].Close()

}

//
//
//
func (cs *ConnectionStruct) IdxGetSocket(idx int) (s *lumerinconnection.LumerinSocketStruct, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if cs.defidx < 0 {
		e = ErrConnMgrNoDefIndex
		return nil, e
	}

	if idx >= len(cs.dst) {
		e = ErrConnMgrIDXOutOfRange
		return nil, e
	}

	return cs.dst[idx], e
}

//
//
//
func (cs *ConnectionStruct) IdxReadReady(idx int) (r bool) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if idx >= 0 &&
		idx < len(cs.dst) &&
		cs.dst[idx] != nil &&
		cs.dst[idx].ReadReady() {
		return true
	}
	return false

}

//
//
//
func (cs *ConnectionStruct) IdxRead(idx int, buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if idx < 0 {
		e = ErrConnMgrNoDefIndex
		return 0, e
	}

	if idx >= len(cs.dst) {
		e = ErrConnMgrIDXOutOfRange
		return 0, e
	}

	if cs.dst[idx] == nil {
		e = ErrConnMgrBadDefDest
		return 0, e

	}

	return cs.dst[idx].Read(buf)

}

//
//
//
func (cs *ConnectionStruct) IdxWrite(idx int, buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if idx < 0 {
		e = ErrConnMgrNoDefIndex
		return 0, e
	}

	if idx >= len(cs.dst) {
		e = ErrConnMgrIDXOutOfRange
		return 0, e
	}

	if cs.dst[idx] == nil {
		e = ErrConnMgrBadDefDest
		return 0, e

	}

	return cs.dst[idx].Write(buf)

}

//
//
//
func (cs *ConnectionStruct) IdxClose(idx int) (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLine()+" called"))

	if idx < 0 {
		e = ErrConnMgrNoDefIndex
		return e
	}

	if idx >= len(cs.dst) {
		e = ErrConnMgrIDXOutOfRange
		return e
	}

	if cs.dst[idx] == nil {
		e = ErrConnMgrBadDefDest
		return e

	}

	e = cs.dst[idx].Close()
	cs.dst[idx] = nil

	return e
}
