package protocol

import (
	"errors"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
)

var ErrProtoConnNoDstIndex = errors.New("DST index not found")
var ErrProtoConnNoID = errors.New("Connection ID not found")

//
// Protocol encapulates the non upper layer protocol specific items
// such as the connection management, which should be the same for all
// It also encapsulates the management functions and tracking of outstanding
// Actions waiting on events
//
// Communications, MsgBus
//
//

type ProtocolConnectionStruct struct {
	Addr net.Addr
	Id   simple.ConnUniqueID
}

type ProtocolDstStruct struct {
	conn map[int]*ProtocolConnectionStruct
}

//
// addConn() Finds and existing instance index or returnes a new index to a new entry
//
func (p *ProtocolDstStruct) addConn(dst net.Addr, id simple.ConnUniqueID) (index int, e error) {

	l := len(p.conn)

	for i := 0; i < l; i++ {
		if p.conn[i].Addr != dst {
			continue
		}
		if p.conn[i].Id != id {
			continue
		}

		return i, e
	}

	p.conn[l] = &ProtocolConnectionStruct{
		Addr: dst,
		Id:   id,
	}

	// Find next open slot and insert it

	return l, nil
}

//
//
//
func (p *ProtocolDstStruct) getConnIndex(id simple.ConnUniqueID) (index int, e error) {

	l := len(p.conn)

	for i := 0; i < l; i++ {
		if p.conn[i].Id == id {
			return l, e
		}
	}

	return -1, ErrProtoConnNoID

}

//
//
//
func (p *ProtocolDstStruct) getConnID(index int) (id simple.ConnUniqueID, e error) {
	var ok bool
	if _, ok = p.conn[index]; !ok {
		e = ErrProtoConnNoDstIndex
	} else {
		id = p.conn[index].Id
	}

	return id, e

}

//
//
//
func (p *ProtocolDstStruct) getConnAddr(index int) (addr net.Addr, e error) {
	var ok bool
	if _, ok = p.conn[index]; !ok {
		e = ErrProtoConnNoDstIndex
	} else {
		addr = p.conn[index].Addr
	}

	return addr, e

}
