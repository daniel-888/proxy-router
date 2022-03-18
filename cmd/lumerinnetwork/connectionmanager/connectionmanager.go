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
const MaxDstSlots int = 16
const DefaultReadBufSize = 1024
const DefaultReadEventChanSize = 10

const SrcIdx int = -1
const DstIdx0 int = 0

var ErrConnMgrNoDefIndex = errors.New("CM: no default index")
var ErrConnMgrBadDefDest = errors.New("CM: bad default destination")
var ErrConnMgrIDXOutOfRange = errors.New("CM: index out of range")
var ErrConnReadNotReady = errors.New("CM: there is nothing to read")

//
// Listen Struct for new SRC connections coming in
//
type ConnectionListenStruct struct {
	listen *lumerinconnection.LumerinListenStruct
	ctx    context.Context
	cancel func()
	port   int
	addr   net.Addr
	accept chan *ConnectionStruct
}

//
// Struct for existing SRC connections and the associated outgoing DST connections
type ConnectionStruct struct {
	src      *lumerinconnection.LumerinSocketStruct
	dst      []*lumerinconnection.LumerinSocketStruct
	defidx   int
	ctx      context.Context
	cancel   func()
	readChan chan *ConnectionReadEvent // Fed by all of the go routines servising the read events
}

//
// Send a single Read Event back up the stack with the
// SRC or DST index ID // -1 == SRC, 0+ == DST
type ConnectionReadEvent struct {
	index int
	data  []byte
	count int
	err   error
}

func (c *ConnectionReadEvent) Index() int   { return c.index }
func (c *ConnectionReadEvent) Data() []byte { return c.data }
func (c *ConnectionReadEvent) Count() int   { return c.count }
func (c *ConnectionReadEvent) Err() error   { return c.err }

//
//
//
func Listen(ctx context.Context) (cls *ConnectionListenStruct, e error) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	ctx, cancel := context.WithCancel(ctx)
	_ = cancel

	addr := contextlib.GetSrc(ctx)

	l, e := lumerinconnection.Listen(ctx, addr)
	if e == nil {
		accept := make(chan *ConnectionStruct)
		cls = &ConnectionListenStruct{
			listen: l,
			ctx:    ctx,
			cancel: cancel,
			port:   0,
			addr:   addr,
			accept: accept,
		}
		go cls.goAccept()
	}

	return cls, e
}

//
//
//
func (cls *ConnectionListenStruct) getPort() (port int) {
	contextlib.Logf(cls.ctx, contextlib.LevelPanic, fmt.Sprint(lumerinlib.FileLineFunc()+" called"))
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
func (cls *ConnectionListenStruct) goAccept() {

	contextlib.Logf(cls.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	defer close(cls.accept)

FORLOOP:
	for {
		select {
		case <-cls.ctx.Done():
			break FORLOOP
		case l := <-cls.listen.Accept():
			ctx, cancel := context.WithCancel(cls.ctx)
			cs := &ConnectionStruct{
				src:      l,
				dst:      []*lumerinconnection.LumerinSocketStruct{},
				defidx:   0,
				ctx:      ctx,
				cancel:   cancel,
				readChan: make(chan *ConnectionReadEvent, DefaultReadEventChanSize),
			}

			cls.accept <- cs
			go cs.goRead(SrcIdx)
		}
	}

	contextlib.Logf(cls.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" Exiting...")
}

//
//
//
func (cls *ConnectionListenStruct) Accept() <-chan *ConnectionStruct {

	contextlib.Logf(cls.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLineFunc()+" called"))

	return cls.accept
}

//
//
//
func (cls *ConnectionListenStruct) Close() (e error) {
	contextlib.Logf(cls.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLineFunc()+" called"))
	return cls.listen.Close()
}

//
//
//
func (cls *ConnectionListenStruct) Cancel() {
	contextlib.Logf(cls.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLineFunc()+" closing down ConnectionListenStruct"))
	cls.cancel()
}

//
// func (cs *ConnectionStruct) goRead()
//
func (cs *ConnectionStruct) goRead(index int) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLineFunc()+" called on index: %d", index))

	defer close(cs.readChan)

	var l *lumerinconnection.LumerinSocketStruct

	if index < 0 {
		l = cs.src
	} else {
		l = cs.dst[index]
	}

	if l == nil {
		contextlib.Logf(cs.ctx, contextlib.LevelPanic, fmt.Sprint(lumerinlib.FileLineFunc()+" bad index:%d", index))
	}

	for {
		select {
		case <-cs.ctx.Done():
			return
		default:
		}

		data := make([]byte, DefaultReadBufSize)
		count, e := l.Read(data)
		data = data[:count]
		if e != nil {
			contextlib.Logf(cs.ctx, contextlib.LevelError, fmt.Sprint(lumerinlib.FileLineFunc()+" Read() on index:%d returned error:%s\n", index, e))
			cs.Cancel()
			return
		}
		if count == 0 {
			contextlib.Logf(cs.ctx, contextlib.LevelError, fmt.Sprint(lumerinlib.FileLineFunc()+" Read() on index:%d returned zero count\n", index))
			cs.Cancel()
			return
		}

		cre := &ConnectionReadEvent{
			index: index,
			data:  data,
			count: count,
			err:   e,
		}

		cs.readChan <- cre
	}
}

//
//
//
func (cs *ConnectionStruct) GetReadChan() <-chan *ConnectionReadEvent {
	if cs == nil {
		panic(lumerinlib.FileLineFunc() + " ConnectionStruct is nil")
	}
	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))
	return cs.readChan
}

//
//
//
func (cs *ConnectionStruct) Cancel() {
	cs.cancel()
}

//
// Dial() opens up a new dst connection and inserts it into the first avalable dst slot
// If this is the 0th slow, the default dst is set as well
//
// func (cs *ConnectionStruct) Dial(ctx context.Context, port int, ip net.IPAddr) (idx int, e error) {
func (cs *ConnectionStruct) Dial(addr net.Addr) (idx int, e error) {

	if cs == nil {
		panic("ConnectionStruct is nil...")
	}

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprint(lumerinlib.FileLineFunc()+" called"))

	idx = -1

	// dst, e := lumerinconnection.Dial(ctx, lumerinconnection.TCP, port, ip)
	dst, e := lumerinconnection.Dial(cs.ctx, addr)
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

	go cs.goRead(idx)

	return idx, nil
}

//
// ReDialIdx() will attempt to reconnect to the same dst, first checking the the line is closed
// It is used in case a connection is severed
//
func (cs *ConnectionStruct) ReDialIdx(idx int) (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelPanic, fmt.Sprintf(lumerinlib.FileLineFunc()+" Function Not Implemented Yet.."))

	return nil
}

//
// Close() will close out all src and dst connections via the cancel context function
//
func (cs *ConnectionStruct) Close() (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

	cs.cancel() // This should close all open src and dst connections

	return nil
}

//
//
//
func (cs *ConnectionStruct) SetRoute(idx int) (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

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
// AnyReadReady() checks all open connections to see if any are ready to read
//
//func (cs *ConnectionStruct) AnyReadReady() (r bool) {
//
//	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))
//
//	if cs.src != nil && cs.src.ReadReady() {
//		return true
//	}
//
//	for i := 0; i < len(cs.dst); i++ {
//		if cs.dst[i] != nil && cs.dst[i].ReadReady() {
//			return true
//		}
//	}
//
//	return false
//}

//
//
//
func (cs *ConnectionStruct) SrcGetSocket() (s *lumerinconnection.LumerinSocketStruct, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

	return cs.src, nil
}

//
//
//
//func (cs *ConnectionStruct) SrcReadReady() (r bool) {
//
//	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))
//
//	return cs.src.ReadReady()
//}

//
//
//
//func (cs *ConnectionStruct) SrcReadStruct() (c *ConnectionSrcDataStruct, e error) {
//
//	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))
//
//	if !cs.SrcReadReady() {
//		return nil, ErrConnReadNotReady
//	}
//
//	c = &ConnectionSrcDataStruct{
//		data: []byte{},
//	}
//
//	count, e := cs.SrcRead(c.data)
//	c.count = count
//	c.err = e
//
//	return c, nil
//}

//
//
//
func (cs *ConnectionStruct) SrcRead(buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

	return cs.src.Read(buf)
}

//
//
//
func (cs *ConnectionStruct) SrcWrite(buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

	return cs.src.Write(buf)

}

//
// SrcClose() calls (*CS) Close() to close everything down
//
func (cs *ConnectionStruct) SrcClose() (e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

	return cs.Close()
}

//
//
//
func (cs *ConnectionStruct) DstGetSocket() (s *lumerinconnection.LumerinSocketStruct, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

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
//func (cs *ConnectionStruct) DstReadReady() (r bool) {
//
//	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))
//
//	if cs.defidx >= 0 && cs.dst[cs.defidx] != nil && cs.dst[cs.defidx].ReadReady() {
//		return true
//	}
//	return false
//}

//
// Read the default Destination
//
//func (cs *ConnectionStruct) DstReadStruct() (d *ConnectionDstDataStruct, e error) {
//
//	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))
//
//	if !cs.DstReadReady() {
//		return nil, ErrConnReadNotReady
//	}
//
//	d = &ConnectionDstDataStruct{
//		index: cs.defidx,
//		data:  []byte{},
//		count: 0,
//		err:   nil,
//	}
//
//	count, e := cs.DstRead(d.data)
//	d.count = count
//	d.err = e
//
//	return d, nil
//}

//
//
//
func (cs *ConnectionStruct) DstRead(buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

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

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

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

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

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

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

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
//func (cs *ConnectionStruct) AnyIdxReadReady() (idx int, r bool) {
//
//	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))
//
//	for idx = 0; idx < len(cs.dst); idx++ {
//		if cs.dst[idx] != nil && cs.dst[idx].ReadReady() {
//			return idx, true
//		}
//	}
//	return -1, false
//
//}

//
//
//
//func (cs *ConnectionStruct) IdxReadReady(idx int) (r bool) {
//
//	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))
//
//	if idx >= 0 &&
//		idx < len(cs.dst) &&
//		cs.dst[idx] != nil &&
//		cs.dst[idx].ReadReady() {
//		return true
//	}
//	return false
//
//}

//
// Read the default Destination
//
//func (cs *ConnectionStruct) IdxReadStruct(idx int) (d *ConnectionDstDataStruct, e error) {
//
//	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))
//
//	if !cs.IdxReadReady(idx) {
//		return nil, ErrConnReadNotReady
//	}
//
//	d = &ConnectionDstDataStruct{
//		index: idx,
//		data:  []byte{},
//		count: 0,
//		err:   nil,
//	}
//
//	count, e := cs.IdxRead(idx, d.data)
//	d.count = count
//	d.err = e
//
//	return d, nil
//}

//
//
//
func (cs *ConnectionStruct) IdxRead(idx int, buf []byte) (count int, e error) {

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

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

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

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

	contextlib.Logf(cs.ctx, contextlib.LevelTrace, fmt.Sprintf(lumerinlib.FileLineFunc()+" called"))

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
