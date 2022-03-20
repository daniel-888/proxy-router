package protocol

import (
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
	Addr  net.Addr
	State ConnectionState
	UID   int
	Err   error
}

type ProtocolDstStruct struct {
	conn map[int]*ProtocolConnectionStruct
}

//
// New() finds or creates an open slot and starts to open a connection to the dst address
//
func (p *ProtocolDstStruct) New(dst net.Addr) (index int, e error) {

	pcs := &ProtocolConnectionStruct{
		Addr:  dst,
		UID:   -1,
		State: ConnStateNew,
		Err:   nil,
	}

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
//
//
func (p *ProtocolConnectionStruct) getUID() int {
	return p.UID
}

//
//
//
func (p *ProtocolConnectionStruct) setUID(uid int) {
	p.UID = uid
}

//
//
//
func (p *ProtocolConnectionStruct) getAddr() net.Addr {
	return p.Addr
}

//
//
//
func (p *ProtocolConnectionStruct) getState() ConnectionState {
	return p.State
}

//
//
//
func (p *ProtocolConnectionStruct) setState(s ConnectionState) {
	p.State = s
}

//
//
//
func (p *ProtocolConnectionStruct) getError() error {
	return p.Err
}

//
//
//
func (p *ProtocolConnectionStruct) setError(e error) {
	p.Err = e
}

//
//
//
func (ps *ProtocolStruct) goOpenConn(index int) {

	slot, ok := ps.dstconn.conn[index]
	if !ok {
		contextlib.Logf(ps.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" Dst slot:%d does not exist", index)
	}

	state := slot.getState()
	if state != ConnStateNew {
		contextlib.Logf(ps.ctx, contextlib.LevelPanic, lumerinlib.FileLineFunc()+" getState() returned:%s, not %s", state, ConnStateNew)
	}

	addr := slot.getAddr()

	UniqueID, e := ps.simple.Dial(addr)
	if e != nil {
		slot.setState(ConnStateError)
		slot.setError(e)
	} else {
		slot.setState(ConnStateReady)
		slot.setUID(UniqueID)
	}

}
