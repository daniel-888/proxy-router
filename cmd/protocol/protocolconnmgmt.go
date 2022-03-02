package protocol

import (
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
)

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
	conn map[int]ProtocolConnectionStruct
}

//
//
//
func (p *ProtocolDstStruct) openConn(dst net.Addr) (index int, e error) {

	return 0, nil
}

//
//
//
func (p *ProtocolDstStruct) closeConn(index int) (e error) {

	return nil
}
