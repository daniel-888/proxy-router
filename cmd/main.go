package main

import (
	"fmt"

	"gitlab.com/TitanInd/lumerin/cmd/accountingmanager"
	// "gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/connectionmanager"
	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
	// "gitlab.com/TitanInd/lumerin/cmd/contractmanager"
	// "gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/localvalidator"
	"gitlab.com/TitanInd/lumerin/cmd/logging"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/walletmanager"
)

func main() {

	done := make(chan int)

	//
	// Fire up the Message Bus
	//
	ps := msgbus.New(10)

	//
	// Setup Default Dest
	//

	dest := msgbus.Dest{
		ID:       msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetProto: msgbus.DestNetProto("tcp"),
		NetHost:  msgbus.DestNetHost("127.0.0.1"),
		NetPort:  msgbus.DestNetPort("3334"),
	}

	event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(msgbus.DEFAULT_DEST_ID), dest)
	if err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", event.Err))
	}

	//
	// Fire up the connection Manager
	//
	cm, err := connectionmanager.New(ps)
	if err != nil {
		panic(fmt.Sprintf("connection manager failed:%s", err))
	}
	err = cm.Start()
	if err != nil {
		panic(fmt.Sprintf("connection manager failed to start:%s", err))
	}

	//	ps.PubWait(msgbus.DestMsg, "destMsg01", msgbus.Dest{})
	//	ps.Sub(msgbus.DestMsg, "destMsg01", ech)
	//	ps.Set(msgbus.DestMsg, "destMsg01", dest)

	//	ps.Get(msgbus.DestMsg, "destMsg01", ech)
	//	ps.Get(msgbus.DestMsg, "", ech)

	//	ps.Set(msgbus.DestMsg, "destMsg01", dest)

	//	time.Sleep(5 * time.Second)

	<-done
	return

	fmt.Println(accountingmanager.BoilerPlateFunc())
	//  fmt.Println(configurationmanager.BoilerPlateFunc())
	//	fmt.Println(connectionmanager.BoilerPlateFunc())
	fmt.Println(connectionscheduler.BoilerPlateFunc())
	// fmt.Println(contractmanager.BoilerPlateFunc())
	// fmt.Println(externalapi.BoilerPlateFunc())
	fmt.Println(localvalidator.BoilerPlateFunc())
	fmt.Println(logging.BoilerPlateFunc())
	fmt.Println(walletmanager.BoilerPlateFunc())
}


