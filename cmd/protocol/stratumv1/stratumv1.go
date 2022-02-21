package stratumv1

import (
	"context"
	"net"

	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/simple"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

type StratumV1ListenStruct struct {
	protocollisten *protocol.ProtocolListenStruct
}

type StratumV1EventStruct struct {
}

//
//
//
func New(ctx context.Context, mb *msgbus.PubSub, src net.Addr, dst net.Addr) (s *StratumV1ListenStruct, e error) {

	// Validate src and dst here

	ctx = context.WithValue(ctx, simple.SimpleMsgBusValue, mb)
	ctx = context.WithValue(ctx, simple.SimpleSrcAddrValue, src)
	ctx = context.WithValue(ctx, simple.SimpleDstAddrValue, dst)
	ctx = context.WithValue(ctx, simple.SimpleEventHandler, &StratumV1EventStruct{})

	protocollisten, err := protocol.New(ctx)
	if err != nil {
		lumerinlib.PanicHere("")
	}

	s = &StratumV1ListenStruct{
		protocollisten: protocollisten,
	}

	return s, e
}

//
//
//
func (s *StratumV1ListenStruct) Cancel() {
	s.protocollisten.Cancel()
}

//
//
//
func (*StratumV1EventStruct) EventHandler(*simple.SimpleStruct) {

}
