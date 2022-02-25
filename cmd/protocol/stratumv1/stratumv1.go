package stratumv1

import (
	"context"
	"fmt"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

//
// Need a series of access functions
//
//
type connectionStruct struct {
	dst net.Addr
	id  simple.ConnUniqueID
}

type StratumV1ListenStruct struct {
	protocollisten *protocol.ProtocolListenStruct
}

type StratumV1Struct struct {
	// simpleproto simple.SimpleProtocolInterface
	// Other Protocol Structure here
	simple      *simple.SimpleStruct
	connections map[int]connectionStruct
}

//
// NewProtocol() called in the SIMPL layer when a new connection is recieved
// Stand up a default connection to the default destination
//
func NewProtocol(ss *simple.SimpleStruct) (s *StratumV1Struct) {

	s = &StratumV1Struct{
		simple:      ss,
		connections: make(map[int]connectionStruct, 0),
	}

	dst := ss.ctx.ContextValue(simple.SimpleDstAddrValue)

	ConnID, err := ss.Dial(dst)
	if err != nil {
		panic("")
	}

	return s
}

//
//
//
func New(ctx context.Context, mb *msgbus.PubSub, src net.Addr, dst net.Addr) (s *StratumV1ListenStruct, e error) {

	// Validate src and dst here

	ctx = context.WithValue(ctx, simple.SimpleProtocolValue, NewProtocol)
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
func (svs *StratumV1Struct) EventHandler(event *simple.SimpleEvent) {

	// Here is the access to the simple layer instance
	// simplestruct := svs.simple

	switch event.EventType {
	case simple.NoEvent:
		return

	case simple.MsgUpdateEvent:
		fallthrough
	case simple.MsgDeleteEvent:
		fallthrough
	case simple.MsgGetEvent:
		fallthrough
	case simple.MsgGetIndexEvent:
		fallthrough
	case simple.MsgSearchEvent:
		fallthrough
	case simple.MsgSearchIndexEvent:
		fallthrough
	case simple.MsgPublishEvent:
		fallthrough
	case simple.MsgUnpublishEvent:
		fallthrough
	case simple.MsgSubscribedEvent:
		fallthrough
	case simple.MsgUnsubscribedEvent:
		fallthrough
	case simple.MsgRemovedEvent:
		msg, ok := event.Data.(msgbus.Event)
		if !ok {
			lumerinlib.PanicHere(fmt.Sprintf(lumerinlib.FileLine()+" Event Data wrong Type:%t", event.Data))
		}
		// Error checking here  ev == megbus.Event
		svs.decodeMsgBusEvent(msg)
		return

	case simple.ConnReadEvent:
		// Error checking here event == connection event
		svs.handleConnReadEvent(event)
		return

	case simple.ConnEOFEvent:
		// Error checking here event == connection event
		svs.handleConnEOFEvent(event)
		return

	case simple.ConnErrorEvent:
		// Error checking here event == connection event
		svs.handleConnErrorEvent(event)
		return

	case simple.ErrorEvent:
		// Error checking here event == Error event
		svs.handleErrorEvent(event)
		return

	default:
		lumerinlib.PanicHere(fmt.Sprintf(lumerinlib.FileLine()+" Default Reached: Event Type:%s", string(event.EventType)))
	}

}

//
//
//
func (svs *StratumV1Struct) decodeMsgBusEvent(event msgbus.Event) {

	switch event.EventType {
	case msgbus.NoEvent:
		fmt.Sprintf(lumerinlib.Funcname() + " NoEvent received, returning\n")
		return
	case msgbus.UpdateEvent:
		svs.handleMsgUpdateEvent(event)
		return
	case msgbus.DeleteEvent:
		svs.handleMsgDeleteEvent(event)
		return
	case msgbus.GetEvent:
		svs.handleMsgGetEvent(event)
		return
	case msgbus.GetIndexEvent:
		svs.handleMsgIndexEvent(event)
		return
	case msgbus.SearchEvent:
		svs.handleMsgSearchEvent(event)
		return
	case msgbus.SearchIndexEvent:
		svs.handleMsgSearchIndexEvent(event)
		return
	case msgbus.PublishEvent:
		svs.handleMsgPublishEvent(event)
		return
	case msgbus.UnpublishEvent:
		svs.handleMsgUnpublishEvent(event)
		return
	case msgbus.SubscribedEvent:
		svs.handleMsgSubscribedEvent(event)
		return
	case msgbus.UnsubscribedEvent:
		svs.handleMsgUnsubscribedEvent(event)
		return
	case msgbus.RemovedEvent:
		svs.handleMsgRemovedEvent(event)
		return

	default:
		lumerinlib.PanicHere(fmt.Sprintf(lumerinlib.FileLine()+" Default Reached: Event Type:%s", string(event.EventType)))
	}
}
