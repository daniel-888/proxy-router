package protocol

import (
	"context"
	"net"

	"gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/simple"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

//
// Top layer protocol template functions that a new protocol will use to access the SIMPLe layer
//

type ProtocolListenStruct struct {
	ctx          context.Context
	cancel       func()
	simplelisten *simple.SimpleListenStruct
}

type ProtocolInterface interface {
	EventHandler(*simple.SimpleStruct)
}

//
// New() Create a new ProtocolListenStruct
//
func New(ctx context.Context, newprotofunc simple.NewProtoFunc) (pls *ProtocolListenStruct, e error) {

	var ok bool

	eh := ctx.Value(simple.SimpleEventHandler)
	if eh == nil {
		lumerinlib.PanicHere("")
	}

	mb := ctx.Value(simple.SimpleMsgBusValue)
	_, ok = mb.(*msgbus.PubSub)
	if !ok {
		lumerinlib.PanicHere("")
	}

	dst := ctx.Value(simple.SimpleDstAddrValue)
	_, ok = dst.(net.Addr)
	if !ok {
		lumerinlib.PanicHere("")
	}

	listen := ctx.Value(simple.SimpleSrcAddrValue)
	_, ok = listen.(net.Addr)
	if !ok {
		lumerinlib.PanicHere("")
	}

	ctx, cancel := context.WithCancel(ctx)

	sls, err := simple.Listen(ctx, listen.(net.Addr), newprotofunc)
	if err != nil {
		lumerinlib.PanicHere("")
	}

	pls = &ProtocolListenStruct{
		ctx:          ctx,
		cancel:       cancel,
		simplelisten: sls,
	}

	return pls, e
}

//
// Cancel() calls the simple layer Cancel function on the SimpleListenStruct
//
func (pls *ProtocolListenStruct) Cancel() {
	pls.cancel()
}

//
// Run() calls the simple layer Run function on the SimpleListenStruct
//
func (pls *ProtocolListenStruct) Run() {
	pls.simplelisten.Run()
	go pls.goAccept()
}

//
// goAccept()
//
func (pls *ProtocolListenStruct) goAccept() {

	go func() {
		select {
		case <-pls.ctx.Done():
			return
		case accept := <-pls.simplelisten.Accept():
			accept.Run()
		}
	}()
	pls.simplelisten.Run()
}
