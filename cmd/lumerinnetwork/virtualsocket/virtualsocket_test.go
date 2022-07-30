package connectionmanager

import (
	"fmt"
	"testing"

	"github.com/daniel-888/proxy-router/lumerinlib"
)

func TestListenOpen(t *testing.T) {

	_, e := Listen(12345)
	if e != nil {
		t.Fatalf(fmt.Sprintf(lumerinlib.FileLine()+"Listen() Test Failed: %s", e))
	}
}
