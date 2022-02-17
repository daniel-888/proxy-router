package simple

import (
	"context"
	"net"
)

const SimpleEventHandler string = "EVENTHANDLER"
const SimpleMsgBusValue string = "MSGBUS"
const SimpleSrcAddrValue string = "SRCADDR"
const SimpleDstAddrValue string = "DSTADDR"

//
// Place holder for the SIMPLe layer code
//

type SimpleListenStruct struct {
	ctx    context.Context
	cancel func()
	accept chan *SimpleStruct
}

type SimpleStruct struct {
	ctx      context.Context
	cancel   func()
	protocol interface{} // Set in the protocol accept function
}

//
//
//
func New(ctx context.Context, listen net.Addr) (*SimpleListenStruct, error) {

	return nil, nil
}

func (*SimpleListenStruct) Run() {

}

func (*SimpleListenStruct) Cancel() {

}

func (*SimpleListenStruct) Accept() <-chan *SimpleStruct {

	return nil
}

func (*SimpleStruct) Run() {

}
