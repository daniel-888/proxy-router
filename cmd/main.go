package main

import (
	"fmt"

	"gitlab.com/TitanInd/lumerin/cmd/accountingmanager"
	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
	//"gitlab.com/TitanInd/lumerin/cmd/testmod"

	"gitlab.com/TitanInd/lumerin/cmd/config"
	"gitlab.com/TitanInd/lumerin/cmd/connectionmanager"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager"
	"gitlab.com/TitanInd/lumerin/cmd/localvalidator"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/walletmanager"
)

// -------------------------------------------
//
// Start up the modules one by one
// Config
// Logger
// MsgBus
// Connection Manager
// Scheduling Manager
// Contract Manager
//
// -------------------------------------------
func main() {
	var buyer bool = false

	// Need something better...
	done := make(chan int)

	buyerstr, err := config.ConfigGetVal(config.BuyerNode)
	if err != nil {
		panic(fmt.Sprintf("Getting Buynernode val failed: %s\n", err))
	}

	if buyerstr != "false" {
		buyer = true
	}

	disableconnection, err := config.ConfigGetVal(config.DisableConnection)
	if err != nil {
		panic(fmt.Sprintf("Getting Disable Connection val failed: %s\n", err))
	}
	disablecontract, err := config.ConfigGetVal(config.DisableContract)
	if err != nil {
		panic(fmt.Sprintf("Getting Disable Contract val failed: %s\n", err))
	}
	disableschedule, err := config.ConfigGetVal(config.DisableSchedule)
	if err != nil {
		panic(fmt.Sprintf("Getting Disable Schedule val failed: %s\n", err))
	}

	//
	// Fire up logger
	//
	// logging.Init(false)
	// defer logging.Cleanup()

	//
	// Fire up the Message Bus
	//
	ps := msgbus.New(10)

	//
	// Setup Default Dest
	//

	dest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetUrl: msgbus.DestNetUrl("stratum+tcp://127.0.0.1:33334/"),
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
	if disableconnection == "false" {

		cm, err := connectionmanager.New(ps)
		if err != nil {
			panic(fmt.Sprintf("connection manager failed:%s", err))
		}
		err = cm.Start()
		if err != nil {
			panic(fmt.Sprintf("connection manager failed to start:%s", err))
		}
	}

	//
	// Fire up schedule manager
	//
	if disableschedule == "false" {
		cs, err := connectionscheduler.New(ps)
		if err != nil {
			panic(fmt.Sprintf("schedule manager failed:%s", err))
		}
		err = cs.Start()
		if err != nil {
			panic(fmt.Sprintf("schedule manager failed to start:%s", err))
		}
	}

	//
	//Fire up contract manager
	//
	if disablecontract == "false" {
		var contractManagerConfig map[string]interface{}

		if buyer {
			contractManagerConfig, err = configurationmanager.LoadConfiguration("/Users/ryanbajollari/go/src/lumerin/cmd/configurationmanager/buyerconfig.json", "contractManager")
		} else {
			contractManagerConfig, err = configurationmanager.LoadConfiguration("/Users/ryanbajollari/go/src/lumerin/cmd/configurationmanager/sellerconfig.json", "contractManager")
		}
		if err != nil {
			panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
		}

		if err != nil {
			panic(fmt.Sprintf("contract manager failed:%s", err))
		}
		if buyer {
			var buyerCM contractmanager.BuyerContractManager
			err = contractmanager.Run(&buyerCM, ps, contractManagerConfig)
		} else {
			var sellerCM contractmanager.SellerContractManager
			err = contractmanager.Run(&sellerCM, ps, contractManagerConfig)
		}
		if err != nil {
			panic(fmt.Sprintf("contract manager failed to start:%s", err))
		}
	}

	//	ps.PubWait(msgbus.DestMsg, "destMsg01", msgbus.Dest{})
	//	ps.Sub(msgbus.DestMsg, "destMsg01", ech)
	//	ps.Set(msgbus.DestMsg, "destMsg01", dest)

	//	ps.Get(msgbus.DestMsg, "destMsg01", ech)
	//	ps.Get(msgbus.DestMsg, "", ech)

	//	ps.Set(msgbus.DestMsg, "destMsg01", dest)

	//	time.Sleep(5 * time.Second)

	// Need a better mechanism for running context

	// if false {
	// 	testmod.MinersTouchAll(ps)
	// }

	<-done

	return

	fmt.Println(accountingmanager.BoilerPlateFunc())
	//  fmt.Println(configurationmanager.BoilerPlateFunc())
	//	fmt.Println(connectionmanager.BoilerPlateFunc())
	// fmt.Println(connectionscheduler.BoilerPlateFunc())
	// fmt.Println(contractmanager.BoilerPlateFunc())
	// fmt.Println(externalapi.BoilerPlateFunc())
	fmt.Println(localvalidator.BoilerPlateFunc())
	// fmt.Println(logging.BoilerPlateFunc())
	fmt.Println(walletmanager.BoilerPlateFunc())
}
