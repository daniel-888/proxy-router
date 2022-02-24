package protocol

import (
	"context"
	"fmt"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

type ContextValue string

const SimpleMsgBusValue ContextValue = "MSGBUS"
const SimpleSrcAddrValue ContextValue = "SRCADDR"
const SimpleDstAddrValue ContextValue = "DSTADDR"

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
func New(ctx context.Context) (pls *ProtocolListenStruct, e error) {

	var ok bool

	// eh := ctx.Value(simple.SimpleEventHandler)
	// if eh == nil {
	// 	lumerinlib.PanicHere("")
	// }

	mb := ctx.Value(SimpleMsgBusValue)
	_, ok = mb.(*msgbus.PubSub)
	if !ok {
		lumerinlib.PanicHere("Missing SimpleMsgBusValue")
	}

	dst := ctx.Value(SimpleDstAddrValue)
	_, ok = dst.(net.Addr)
	if !ok {
		lumerinlib.PanicHere("Missing SimpleDstAddrValue")
	}

	listen := ctx.Value(SimpleSrcAddrValue)
	_, ok = listen.(net.Addr)
	if !ok {
		lumerinlib.PanicHere("Missing SimpleSrcAddrValue")
	}

	ctx, cancel := context.WithCancel(ctx)

	sls, err := simple.New(ctx, listen.(net.Addr))
	if err != nil {
		lumerinlib.PanicHere(fmt.Sprintf("Error:%s", err))
	}

	pls = &ProtocolListenStruct{
		ctx:          ctx,
		cancel:       cancel,
		simplelisten: &sls,
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
		case newSimpleStruct := <-pls.simplelisten.Accept():
			newSimpleStruct.Run(pls.ctx)
		}
	}()
	pls.simplelisten.Run()
}
