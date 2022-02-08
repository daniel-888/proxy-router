package main

import (
	"fmt"
	"strconv"
	"time"
	"context"
	"os"
	"os/signal"

	"gitlab.com/TitanInd/lumerin/cmd/config"
	"gitlab.com/TitanInd/lumerin/cmd/connectionmanager"
	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
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
	mainContext, mainCancel := context.WithCancel(context.Background())
	sigInt := make(chan os.Signal, 1) 
	signal.Notify(sigInt, os.Interrupt)

	var buyer bool = false

	configFile, err := config.ConfigGetVal(config.ConfigConfigFilePath)
	if err != nil {
		panic(fmt.Sprintf("Getting Config File val failed: %s\n", err))
	}

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
	disableapi, err := config.ConfigGetVal(config.DisableAPI)
	if err != nil {
		panic(fmt.Sprintf("Getting Disable API val failed: %s\n", err))
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
	defaultpooladdr, err := config.ConfigGetVal(config.DefaultPoolAddr)
	if err != nil {
		panic(fmt.Sprintf("Getting Default Pool Address/URL: %s\n", err))
	}
	dest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetUrl: msgbus.DestNetUrl(defaultpooladdr),
	}

	event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(msgbus.DEFAULT_DEST_ID), dest)
	if err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", event.Err))
	}

	//
	// Setup Node Operator Msg
	//
	nodeOperator := msgbus.NodeOperator{
		ID: msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
		DefaultDest: dest.ID,
	}
	event, err = ps.PubWait(msgbus.NodeOperatorMsg, msgbus.IDString(nodeOperator.ID), nodeOperator)
	if err != nil {
		panic(fmt.Sprintf("Adding Node Operator Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Node Operator Failed: %s", event.Err))
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
		cs, err := connectionscheduler.New(&mainContext, ps, &nodeOperator)
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
		var contractManagerConfig msgbus.ContractManagerConfig
		contractManagerConfigID := msgbus.GetRandomIDString()

		network, err := config.ConfigGetVal(config.ConfigContractNetwork)
		if err != nil {
			panic(fmt.Sprintf("Getting network val failed %s\n", err))
		}

		switch network {
		case "ropsten":
			contractManagerConfig.CloneFactoryAddress = "0x15BdE7774F4A69A7d1fdb66CE94CDF26FCd8F45e"
			contractManagerConfig.LumerinTokenAddress = "0x84E00a18a36dFa31560aC216da1A9bef2164647D"
			contractManagerConfig.ValidatorAddress = "0x508CD3988E2b4B8f1d243b961a855347349f6F63"
			contractManagerConfig.ProxyAddress = "0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3"
			fmt.Println("Connecting to Ropsten Network")
		case "custom":
			contractManagerConfig.CloneFactoryAddress = "0x15BdE7774F4A69A7d1fdb66CE94CDF26FCd8F45e"
			contractManagerConfig.LumerinTokenAddress = "0x84E00a18a36dFa31560aC216da1A9bef2164647D"
			contractManagerConfig.ValidatorAddress = "0x508CD3988E2b4B8f1d243b961a855347349f6F63"
			contractManagerConfig.ProxyAddress = "0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3"
		case "mainnet":
			contractManagerConfig.CloneFactoryAddress = "0x15BdE7774F4A69A7d1fdb66CE94CDF26FCd8F45e"
			contractManagerConfig.LumerinTokenAddress = "0x84E00a18a36dFa31560aC216da1A9bef2164647D"
			contractManagerConfig.ValidatorAddress = "0x508CD3988E2b4B8f1d243b961a855347349f6F63"
			contractManagerConfig.ProxyAddress = "0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3"
			fmt.Println("Connecting to Main Network")
		default:
			panic(fmt.Sprintln("Invalid network input (must be ropsten, custom, or mainnet)"))
		}

		if configFile != "" { // if a config file was specified use it instead of flag params
			contractManagerConfigFile, err := config.LoadConfiguration("contractManager")
			if err != nil {
				panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
			}

			contractManagerConfig.Mnemonic = contractManagerConfigFile["mnemonic"].(string)
			contractManagerConfig.AccountIndex = int(contractManagerConfigFile["accountIndex"].(float64))
			contractManagerConfig.EthNodeAddr = contractManagerConfigFile["ethNodeAddr"].(string)
			contractManagerConfig.ClaimFunds = contractManagerConfigFile["claimFunds"].(bool)
			contractManagerConfig.CloneFactoryAddress = contractManagerConfigFile["cloneFactoryAddress"].(string)
			contractManagerConfig.LumerinTokenAddress = contractManagerConfigFile["lumerinTokenAddress"].(string)
			contractManagerConfig.ValidatorAddress = contractManagerConfigFile["validatorAddress"].(string)
			contractManagerConfig.ProxyAddress = contractManagerConfigFile["proxyAddress"].(string)
		} else {
			contractManagerConfig.Mnemonic, err = config.ConfigGetVal(config.ConfigContractMnemonic)
			if err != nil {
				panic(fmt.Sprintf("Getting mnemonic val failed: %s\n", err))
			}
		
			accountIndexStr, err := config.ConfigGetVal(config.ConfigContractAccountIndex)
			if err != nil {
				panic(fmt.Sprintf("Getting account index val failed: %s\n", err))
			}
			contractManagerConfig.AccountIndex,err = strconv.Atoi(accountIndexStr)
			if err != nil {
				panic(fmt.Sprintf("Converting account index string to int failed: %s\n", err))
			}
	
			contractManagerConfig.EthNodeAddr, err = config.ConfigGetVal(config.ConfigContractEthereumNodeAddress)
			if err != nil {
				panic(fmt.Sprintf("Getting ethereum node address val failed: %s\n", err))
			}
			contractManagerConfig.ClaimFunds = false
			claimFundsStr, err := config.ConfigGetVal(config.ConfigContractClaimFunds)
			if err != nil {
				panic(fmt.Sprintf("Getting claim funds val failed: %s\n", err))
			}
			if claimFundsStr == "true" {
				contractManagerConfig.ClaimFunds = true
			}
		}
		
		// Publish Contract Manager Config to MsgBus
		ps.PubWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractManagerConfig)

		if buyer {
			var buyerCM contractmanager.BuyerContractManager
			err = contractmanager.Run(&mainContext, &buyerCM, ps, contractManagerConfigID, &nodeOperator)
		} else {
			var sellerCM contractmanager.SellerContractManager
			err = contractmanager.Run(&mainContext, &sellerCM, ps, contractManagerConfigID, &nodeOperator)
		}
		if err != nil {
			panic(fmt.Sprintf("contract manager failed to run:%s", err))
		}
	}

	//
	//Fire up external api
	//
	if disableapi == "false" {
		var api externalapi.APIRepos
		api.InitializeJSONRepos(ps)
		time.Sleep(time.Millisecond*2000)
		go api.RunAPI()
	}

	select {
	case <-sigInt:
		fmt.Println("Signal Interupt: Cancelling all contexts and shuting down program")
		mainCancel()
	case <-mainContext.Done():
		time.Sleep(time.Second*5)
		signal.Stop(sigInt)
		return
	}
}