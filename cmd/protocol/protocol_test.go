package protocol

import (
	"context"
	"fmt"
	"testing"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

func TestNewProto(t *testing.T) {

	ps := msgbus.New(1, nil)
	src := lumerinlib.NewNetAddr(lumerinlib.TCP, "127.0.0.1:12345")
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, "127.0.0.1:12345")

	ctx := context.Background()
	cs := &contextlib.ContextStruct{}
	cs.SetMsgBus(ps)
	cs.SetSrc(src)
	cs.SetDst(dst)
	cs.SetProtocol(newProtcolFunc)
	ctx = context.WithValue(ctx, contextlib.ContextKey, cs)

	pls, e := NewListen(ctx)
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, fmt.Sprintf("New() problem:%s", e))
	}

	pls.Run()

	pls.Cancel()

}

func TestOpenConn(t *testing.T) {

	ps := msgbus.New(1, nil)
	src := lumerinlib.NewNetAddr(lumerinlib.TCP, "127.0.0.1:12345")
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, "127.0.0.1:12345")

	ctx := context.Background()
	cs := &contextlib.ContextStruct{}
	cs.SetMsgBus(ps)
	cs.SetSrc(src)
	cs.SetDst(dst)
	cs.SetProtocol(newProtcolFunc)
	ctx = context.WithValue(ctx, contextlib.ContextKey, cs)
	simple := &simple.SimpleStruct{
		ctx: ctx,
	}

	proto := &ProtocolStruct{
		ctx:    ctx,
		simple: simple,
	}

	index, e := proto.OpenConn(dst)
	if e != nil {
		contextlib.Logf(ctx, contextlib.LevelPanic, fmt.Sprintf("New() problem:%s", e))
	}
	_ = index

}

//
//
//
func newProtcolFunc(ss *simple.SimpleStruct) chan *simple.SimpleEvent {

	sec := make(chan *simple.SimpleEvent)

	go func(sec chan *simple.SimpleEvent) {

		for e := range sec {
			fmt.Printf("Event Recieved:%v", e)
		}

	}(sec)

	return sec
}
