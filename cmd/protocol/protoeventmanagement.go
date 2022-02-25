package protocol

import (
	"context"
	"fmt"
	"net"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

//
// Top layer protocol template functions that a new protocol will use to access the SIMPLe layer
//

type ProtocolConnectionStruct struct {
	Addr net.Addr
	Id   simple.ConnUniqueID
}

type ProtocolListenStruct struct {
	ctx          context.Context
	cancel       func()
	simplelisten *simple.SimpleListenStruct
}

//type ProtocolInterface interface {
//	EventHandler(*simple.SimpleStruct)
//}

//
// New() Create a new ProtocolListenStruct
//
func New(ctx context.Context) (pls *ProtocolListenStruct, e error) {

	var ok bool

	//
	// Basic error checking, make sure that the SimpleContextStruct is
	// filled out correctly
	//
	sc := ctx.Value(simple.SimpleContext)
	if sc == nil {
		lumerinlib.PanicHere("")
	}

	sc, ok = sc.(simple.SimpleContextStruct)
	if !ok {
		lumerinlib.PanicHere("")
	}
	if sc.(simple.SimpleContextStruct).Protocol == nil {
		lumerinlib.PanicHere("")
	}
	if sc.(simple.SimpleContextStruct).MsgBus == nil {
		lumerinlib.PanicHere("")
	}
	if sc.(simple.SimpleContextStruct).Dst == nil {
		lumerinlib.PanicHere("")
	}
	if sc.(simple.SimpleContextStruct).Src == nil {
		lumerinlib.PanicHere("")
	}

	ctx, cancel := context.WithCancel(ctx)

	listenaddr := sc.(simple.SimpleContextStruct).Src
	sls, err := simple.New(ctx, listenaddr)
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
