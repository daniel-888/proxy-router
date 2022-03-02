package protocol

import (
	"context"
	"fmt"
	"testing"

	simple "gitlab.com/TitanInd/lumerin/cmd/lumerinnetwork/SIMPL"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func TestNewProto(t *testing.T) {

	ps := msgbus.New(1)
	src := lumerinlib.NewNetAddr(lumerinlib.TCP, "127.0.0.1:12345")
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, "127.0.0.1:12345")

	sc := simple.SimpleContextStruct{
		Protocol: newProtcolFunc,
		MsgBus:   ps,
		Src:      src,
		Dst:      dst,
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, simple.SimpleContext, sc)

	pls, e := NewListen(ctx)
	if e != nil {
		lumerinlib.PanicHere(fmt.Sprintf("New() problem:%s", e))
	}

	pls.Run()

	pls.Cancel()

}

//
//
//
func newProtcolFunc(ss *simple.SimpleStruct) chan *simple.SimpleEvent {

	sec := make(chan *simple.SimpleEvent)

	go func(sec chan *simple.SimpleEvent) {
		for {
			select {
			case e := <-sec:
				fmt.Printf("Event Recieved:%v", e)
			}
		}

	}(sec)

	return sec
}
