package stratumv1

import (
	"context"
	"fmt"
	"testing"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func TestNewProto(t *testing.T) {

	ps := msgbus.New(1, nil)
	src := lumerinlib.NewNetAddr(lumerinlib.TCP, "127.0.0.1:12345")
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, "127.0.0.1:12345")

	ctx := context.Background()

	sls, err := New(ctx, ps, src, dst)
	if err != nil {
		lumerinlib.PanicHere(fmt.Sprintf("New() problem:%s", err))
	}

	sls.Run()
	sls.Cancel()

}
