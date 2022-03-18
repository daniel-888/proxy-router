package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/config"
	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/log"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol/stratumv1"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

// -------------------------------------------
//
// Start up the modules one by one
// Config
// Log
// MsgBus
// Connection Manager
// Scheduling Manager
// Contract Manager
//
// -------------------------------------------
func main() {
	l := log.New()

	logFile, err := os.OpenFile(config.MustGet(config.ConfigLogFilePath), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		l.Logf(log.LevelFatal, "error opening log file: %v", err)
	}
	defer logFile.Close()

	l.SetFormat(log.FormatJSON).SetOutput(logFile)

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

	disablecontract, err := config.ConfigGetVal(config.DisableContract)
	if err != nil {
		panic(fmt.Sprintf("Getting Disable Contract val failed: %s\n", err))
	}
	disableschedule, err := config.ConfigGetVal(config.DisableSchedule)
	if err != nil {
		panic(fmt.Sprintf("Getting Disable Schedule val failed: %s\n", err))
	}
	disablestratumv1, err := config.ConfigGetVal(config.DisableStratumv1)
	if err != nil {
		panic(fmt.Sprintf("Getting Disable Schedule val failed: %s\n", err))
	}

	listenport, err := config.ConfigGetVal(config.ConfigConnectionListenPort)
	if err != nil {
		panic(fmt.Sprintf("Getting Listen Port val failed: %s\n", err))
	}

	listenip, err := config.ConfigGetVal(config.ConfigConnectionListenIP)
	if err != nil {
		panic(fmt.Sprintf("Getting Listen IP val failed: %s\n", err))
	}

	defaultpooladdr, err := config.ConfigGetVal(config.DefaultPoolAddr)
	if err != nil {
		panic(fmt.Sprintf("Getting Default Pool Address/URL: %s\n", err))
	}

	disableapi, err := config.ConfigGetVal(config.DisableAPI)
	if err != nil {
		panic(fmt.Sprintf("Getting Disable API val failed: %s\n", err))
	}

	//
	// Fire up logger
	//
	// log := log.New()

	//
	// Fire up the Message Bus
	//
	ps := msgbus.New(10, l)

	//
	// Add the various Context variables here
	// msgbus, logger, defailt listen address, defalt desitnation address
	//

	src := lumerinlib.NewNetAddr(lumerinlib.TCP, listenip+":"+listenport)
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, defaultpooladdr)

	//
	// the proro argument (#1) gets set in the Protocol sus-system
	//
	cs := contextlib.NewContextStruct(nil, ps, l, src, dst)

	//
	//  All of the various needed subsystem values get passed into the context here.
	//
	mainContext = context.WithValue(mainContext, contextlib.ContextKey, cs)

	//
	// Setup Default Dest
	//
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
		ID:          msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
		IsBuyer:     buyer,
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
	// Fire up the StratumV1 Potocol
	//
	if disablestratumv1 == "false" {

		src, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", listenip, listenport))
		if err != nil {
			lumerinlib.PanicHere("")
		}

		dst, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", "127.0.0.1", "3334"))
		if err != nil {
			lumerinlib.PanicHere("")
		}

		stratum, err := stratumv1.NewListener(mainContext, src, dst)
		if err != nil {
			panic(fmt.Sprintf("Stratum Protocol New() failed:%s", err))
		}

		stratum.Run()

	}

	//
	// Fire up schedule manager
	//
	if disableschedule == "false" {
		cs, err := connectionscheduler.New(&mainContext, &nodeOperator)
		if err != nil {
			l.Logf(log.LevelPanic, "Schedule manager failed: %v", err)
		}
		err = cs.Start()
		if err != nil {
			l.Logf(log.LevelPanic, "Schedule manager to start: %v", err)
		}
	}

	//
	// Fire up contract manager
	//
	if disablecontract == "false" {
		var contractManagerConfig msgbus.ContractManagerConfig
		contractManagerConfigID := msgbus.GetRandomIDString()

		network, err := config.ConfigGetVal(config.ConfigContractNetwork)
		if err != nil {
			l.Logf(log.LevelPanic, "Getting network val failed: %v", err)
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
			fmt.Println("Connecting to Custom Network")
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
				l.Logf(log.LevelPanic, "Failed to load contract manager configuration: %v", err)
			}

			contractManagerConfig.Mnemonic = contractManagerConfigFile["mnemonic"].(string)
			contractManagerConfig.AccountIndex = int(contractManagerConfigFile["accountIndex"].(float64))
			contractManagerConfig.EthNodeAddr = contractManagerConfigFile["ethNodeAddr"].(string)
			contractManagerConfig.ClaimFunds = contractManagerConfigFile["claimFunds"].(bool)
			contractManagerConfig.TimeThreshold = int(contractManagerConfigFile["timeThreshold"].(float64))
			contractManagerConfig.CloneFactoryAddress = contractManagerConfigFile["cloneFactoryAddress"].(string)
			contractManagerConfig.LumerinTokenAddress = contractManagerConfigFile["lumerinTokenAddress"].(string)
			contractManagerConfig.ValidatorAddress = contractManagerConfigFile["validatorAddress"].(string)
			contractManagerConfig.ProxyAddress = contractManagerConfigFile["proxyAddress"].(string)
		} else {
			contractManagerConfig.Mnemonic, err = config.ConfigGetVal(config.ConfigContractMnemonic)
			if err != nil {
				l.Logf(log.LevelPanic, "Getting mnemonic val failed: %v", err)
			}

			accountIndexStr, err := config.ConfigGetVal(config.ConfigContractAccountIndex)
			if err != nil {
				l.Logf(log.LevelPanic, "Getting account index val failed: %v", err)
			}
			contractManagerConfig.AccountIndex, err = strconv.Atoi(accountIndexStr)
			if err != nil {
				l.Logf(log.LevelPanic, "Converting account index string to int failed: %v", err)
			}

			contractManagerConfig.EthNodeAddr, err = config.ConfigGetVal(config.ConfigContractEthereumNodeAddress)
			if err != nil {
				l.Logf(log.LevelPanic, "Getting ethereum node address val failed: %v", err)
			}
			contractManagerConfig.ClaimFunds = false
			claimFundsStr, err := config.ConfigGetVal(config.ConfigContractClaimFunds)
			if err != nil {
				l.Logf(log.LevelPanic, "Getting claim funds val failed: %v", err)
			}
			if claimFundsStr == "true" {
				contractManagerConfig.ClaimFunds = true
			}

			timeThresholdStr, err := config.ConfigGetVal(config.ConfigContractTimeThreshold)
			if err != nil {
				l.Logf(log.LevelPanic, "Getting time threshold val failed: %v", err)
			}
			contractManagerConfig.TimeThreshold, err = strconv.Atoi(timeThresholdStr)
			if err != nil {
				l.Logf(log.LevelPanic, "Converting time threshold string to int failed: %v", err)
			}
		}

		// Publish Contract Manager Config to MsgBus
		ps.PubWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractManagerConfig)

		if buyer {
			var buyerCM contractmanager.BuyerContractManager
			err = contractmanager.Run(&mainContext, &buyerCM, contractManagerConfigID, &nodeOperator)
		} else {
			var sellerCM contractmanager.SellerContractManager
			err = contractmanager.Run(&mainContext, &sellerCM, contractManagerConfigID, &nodeOperator)
		}
		if err != nil {
			l.Logf(log.LevelPanic, "Contract manager failed to run: %v", err)
		}
	}

	//
	//Fire up external api
	//
	if disableapi == "false" {
		var api externalapi.APIRepos
		api.InitializeJSONRepos(ps)
		time.Sleep(time.Millisecond * 2000)
		port := config.MustGet(config.ConfigRESTPort)
		go api.RunAPI(port, l)
	}

	select {
	case <-sigInt:
		fmt.Println("Signal Interupt: Cancelling all contexts and shuting down program")
		mainCancel()
	case <-mainContext.Done():
		time.Sleep(time.Second * 5)
		signal.Stop(sigInt)
		return
	}
}
