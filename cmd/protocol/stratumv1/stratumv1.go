package stratumv1

import (
	"context"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

type StratumV1ListenStruct struct {
	protocollisten *protocol.ProtocolListenStruct
}

type StratumV1Struct struct {
	simpleproto simple.SimpleProtocolInterface
	// Other Protocol Structure here
}

func NewProtocol() (s *StratumV1Struct) {
	s = &StratumV1Struct{}
	return s
}

//
//
//
func New(ctx context.Context, mb *msgbus.PubSub, src net.Addr, dst net.Addr) (s *StratumV1ListenStruct, e error) {

	// Validate src and dst here

	ctx = context.WithValue(ctx, simple.SimpleMsgBusValue, mb)
	ctx = context.WithValue(ctx, simple.SimpleSrcAddrValue, src)
	ctx = context.WithValue(ctx, simple.SimpleDstAddrValue, dst)

	// var newprotointerface interface{} = NewProtocol

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
func (s *StratumV1ListenStruct) Run() {
	s.protocollisten.Run()
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
// Event Handler
func (*StratumV1Struct) EventHandler(event *simple.SimpleEvent) {

	switch event.EventType {
	case simple.NoEvent:
		return

	case simple.MsgUpdateEvent:
		return

	case simple.MsgDeleteEvent:
		return

	case simple.MsgGetEvent:
		return

	case simple.MsgGetIndexEvent:
		return

	case simple.MsgSearchEvent:
		return

	case simple.MsgSearchIndexEvent:
		return

	case simple.MsgPublishEvent:
		return

	case simple.MsgUnpublishEvent:
		return

	case simple.MsgSubscribedEvent:
		return

	case simple.MsgUnsubscribedEvent:
		return

	case simple.MsgRemovedEvent:
		return

	case simple.ConnReadEvent:
		return

	case simple.ConnEOFEvent:
		return

	case simple.ConnErrorEvent:
		return

	case simple.ErrorEvent:
		return

	default:
		panic("")
	}

}
