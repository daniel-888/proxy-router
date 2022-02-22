package simple

import (
	"context"
	"net"

	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/connectionmanager"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

type NewProtoFunc func() *SimpleProtocolInterface
type ContextValue string
type EventType string

const SimpleMsgBusValue ContextValue = "MSGBUS"
const SimpleSrcAddrValue ContextValue = "SRCADDR"
const SimpleDstAddrValue ContextValue = "DSTADDR"

const NoEvent EventType = "noevent"
const MsgUpdateEvent EventType = "msgupdate"
const MsgDeleteEvent EventType = "msgdelete"
const MsgGetEvent EventType = "msgget"
const MsgGetIndexEvent EventType = "msgindex"
const MsgSearchEvent EventType = "msgsearch"
const MsgSearchIndexEvent EventType = "msgsearchindex"
const MsgPublishEvent EventType = "msgpublish"
const MsgUnpublishEvent EventType = "msgunpublish"
const MsgSubscribedEvent EventType = "msgsubscribe"
const MsgUnsubscribedEvent EventType = "msgunsubscribe"
const MsgRemovedEvent EventType = "msgremoved"
const ConnReadEvent EventType = "connread"
const ConnEOFEvent EventType = "conneof"
const ConnErrorEvent EventType = "connerror"
const ErrorEvent EventType = "error"

//
// Place holder for the SIMPLe layer code
//
type SimpleListenStruct struct {
	ctx              context.Context
	cancel           func()
	accept           chan *SimpleStruct
	newprotocol      NewProtoFunc
	connectionlisten *connectionmanager.ConnectionListenStruct
}

type SimpleStruct struct {
	ctx        context.Context
	cancel     func()
	protocol   *SimpleProtocolInterface
	connection connectionmanager.ConnectionStruct
}

//
// The Data interfaces need to be defined
//
type SimpleEvent struct {
	Event EventType
	Data  interface{}
}

type SimpleProtocolInterface interface {
	EventHandler(*SimpleEvent)
}

//
// Listen()
// Create a new connection manager to listen
// Connect to the Msgbus
//
func Listen(ctx context.Context, listen net.Addr, newproto NewProtoFunc) (sls *SimpleListenStruct, e error) {

	ctx, cancel := context.WithCancel(ctx)

	cls, err := connectionmanager.Listen(ctx, listen)
	if err != nil {
		lumerinlib.PanicHere("")
	}

	sls = &SimpleListenStruct{
		ctx:              ctx,
		cancel:           cancel,
		accept:           make(chan *SimpleStruct),
		connectionlisten: cls,
		newprotocol:      newproto,
	}

	return sls, e
}

//
// Run the listener
//
func (sls *SimpleListenStruct) Run() {
	sls.connectionlisten.Run()
	go sls.goAccept()

}

//
// Cancel the listener
//
func (sls *SimpleListenStruct) Cancel() {
	sls.cancel()
}

//
// Return the simple struct via a channel
//
func (sls *SimpleListenStruct) Accept() <-chan *SimpleStruct {
	return sls.accept
}

//
// Return the simple struct via a channel
//
func (sls *SimpleListenStruct) goAccept() {

	select {
	case sls.ctx.Done():
		return
	case accept := <-sls.accept:
		accept.Run()
	}

}

func (*SimpleStruct) Run() {

}
