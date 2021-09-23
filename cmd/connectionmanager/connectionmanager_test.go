package connectionmanager

import (
	"fmt"
	"testing"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func TestBoilerPlateFunc(t *testing.T) {

	// waitchan := make(chan int)

	ps := msgbus.New(1)

	// cm, err := connectionmanager.New(ps)
	cm, err := New(ps)

	if err != nil {
		panic("connection manager fialed")
	}

	err = cm.Start()
	if err != nil {
		panic("connection manager failed to start")
	}

	fmt.Println("Break point")
	//<-waitchan

}

