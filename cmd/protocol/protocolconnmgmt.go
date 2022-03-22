package protocol

import (
	"context"
	"errors"
	"net"

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
	uID    int
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
func (p *ProtocolConnectionStruct) GetUID() int {
	return p.uID
}

//
// SetUID()
//
func (p *ProtocolConnectionStruct) SetUID(uid int) {
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

//
// goOpenConn()
// Dial new connection to Dst, and then handle the start up sequence
//
func (pcs *ProtocolConnectionStruct) goOpenConn(ps *ProtocolStruct) {

	contextlib.Logf(pcs.ctx, contextlib.LevelTrace, lumerinlib.FileLineFunc()+" called")

	state := pcs.GetState()
	if state != ConnStateNew {
		contextlib.Logf(pcs.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" getState() returned:%s, not %s", state, ConnStateNew)
	}

	addr := pcs.GetAddr()

	UniqueID, e := ps.simple.Dial(addr)
	if e != nil {
		pcs.SetState(ConnStateError)
		pcs.SetError(e)
	} else {
		pcs.SetState(ConnStateReady)
		pcs.SetUID(UniqueID)
	}

	// Open Handler Here

}
