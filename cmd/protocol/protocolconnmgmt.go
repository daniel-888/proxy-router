package protocol

import (
	"context"
	"errors"
	"fmt"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

var ErrProtoConnNoDstIndex = errors.New("DST index not found")
var ErrProtoConnNoID = errors.New("Connection ID not found")

type ConnectionState string

const ConnStateNew ConnectionState = "New"
const ConnStateInit ConnectionState = "Init"
const ConnStateReady ConnectionState = "Ready"
const ConnStateError ConnectionState = "Error"

//
// Protocol encapulates the non upper layer protocol specific items
// such as the connection management, which should be the same for all
// It also encapsulates the management functions and tracking of outstanding
// Actions waiting on events
//
// Communications, MsgBus
//
//

// Addr stores the address of the connection
// State shows the current state of the conneciton
// UID is the unique ID from the SIMPL layer
// Err shows the error state if it is in the error state
type ProtocolConnectionStruct struct {
	ctx    context.Context
	cancel func()
	addr   net.Addr
	state  ConnectionState
	uID    simple.ConnUniqueID
	err    error
}

type ProtocolDstStruct struct {
	ctx  context.Context
	conn map[int]*ProtocolConnectionStruct
}

//
//
//
func NewProtocolConnectionStruct(ctx context.Context, addr net.Addr) (pcs *ProtocolConnectionStruct) {

	contextlib.Logf(ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	ctx, cancel := context.WithCancel(ctx)
	pcs = &ProtocolConnectionStruct{
		ctx:    ctx,
		cancel: cancel,
		addr:   addr,
		uID:    -1,
		state:  ConnStateNew,
		err:    nil,
	}

	return pcs
}

//
// New() finds or creates an open slot and starts to open a connection to the dst address
//
func (p *ProtocolDstStruct) NewProtocolDstStruct(dst net.Addr) (index int, e error) {

	contextlib.Logf(p.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	pcs := NewProtocolConnectionStruct(p.ctx, dst)

	var i int
	length := len(p.conn)
	for i = 0; i < length; i++ {
		if p.conn[i] == nil {
			p.conn[i] = pcs
		}
	}
	if i == length {
		p.conn[i] = pcs
	}

	return i, e
}

//
// These should use an Interface.... circle back to this later
//

// ---------------------------------------------------
// *ProtocolDstStruct)

//
// Cancel()
//
func (p *ProtocolDstStruct) Cancel(index int) (e error) {
	if p.conn[index] == nil {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+"Index:%d does not exist", index)
	} else {
		p.conn[index].cancel()
	}

	return e
}

//
// GetUID()
//
func (p *ProtocolDstStruct) GetUID(index int) (uid simple.ConnUniqueID, e error) {
	if p.conn[index] == nil {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+"Index:%d does not exist", index)
	} else {
		uid = p.conn[index].uID
	}
	return uid, e
}

//
// SetUID()
//
func (p *ProtocolDstStruct) SetUID(index int, uid simple.ConnUniqueID) (e error) {
	if p.conn[index] == nil {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+"Index:%d does not exist", index)
	} else {
		p.conn[index].uID = uid
	}

	return e

}

//
// GetAddr()
//
func (p *ProtocolDstStruct) GetAddr(index int) (addr net.Addr, e error) {
	if p.conn[index] == nil {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+"Index:%d does not exist", index)
	} else {
		addr = p.conn[index].addr
	}

	return addr, e
}

//
// GetState()
//
func (p *ProtocolDstStruct) GetState(index int) (state ConnectionState, e error) {
	if p.conn[index] == nil {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+"Index:%d does not exist", index)
	} else {
		state = p.conn[index].state
	}

	return state, e
}

//
// SetState()
//
func (p *ProtocolDstStruct) SetState(index int, s ConnectionState) (e error) {
	if p.conn[index] == nil {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+"Index:%d does not exist", index)
	} else {
		p.conn[index].state = s
	}

	return e
}

//
//
//
func (p *ProtocolDstStruct) GetError(index int) (err error, e error) {
	if p.conn[index] == nil {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+"Index:%d does not exist", index)
	} else {
		err = p.conn[index].err
	}

	return err, e
}

//
//
//
func (p *ProtocolDstStruct) SetError(index int, err error) (e error) {
	if p.conn[index] == nil {
		e = fmt.Errorf(lumerinlib.FileLineFunc()+"Index:%d does not exist", index)
	} else {
		p.conn[index].err = err
	}

	return e
}

// ---------------------------------------------------
// *ProtocolConnectionStruct)

//
// Ctx()
//
func (p *ProtocolConnectionStruct) Ctx() context.Context {
	return p.ctx
}

//
// Cancel()
//
func (p *ProtocolConnectionStruct) Cancel() {
	p.cancel()
}

//
// GetUID()
//
func (p *ProtocolConnectionStruct) GetUID() simple.ConnUniqueID {
	return p.uID
}

//
// SetUID()
//
func (p *ProtocolConnectionStruct) SetUID(uid simple.ConnUniqueID) {
	p.uID = uid
}

//
// GetAddr()
//
func (p *ProtocolConnectionStruct) GetAddr() net.Addr {
	return p.addr
}

//
// GetState()
//
func (p *ProtocolConnectionStruct) GetState() ConnectionState {
	return p.state
}

//
// SetState()
//
func (p *ProtocolConnectionStruct) SetState(s ConnectionState) {
	p.state = s
}

//
//
//
func (p *ProtocolConnectionStruct) GetError() error {
	return p.err
}

//
//
//
func (p *ProtocolConnectionStruct) SetError(e error) {
	p.err = e
}
