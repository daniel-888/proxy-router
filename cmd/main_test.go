package main

import (
	"testing"
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/config"
	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
	"gitlab.com/TitanInd/lumerin/cmd/log"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol/stratumv1"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

func SimMain(ps *msgbus.PubSub, l *log.Logger, configs config.ConfigRead ) msgbus.DestID {
	mainContext := context.Background()

	//
	// Add the various Context variables here
	// msgbus, logger, defailt listen address, defalt desitnation address
	//

	src := lumerinlib.NewNetAddr(lumerinlib.TCP, configs.ListenIP+":"+configs.ListenPort)
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, configs.DefaultPoolAddr)

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
		NetUrl: msgbus.DestNetUrl(configs.DefaultPoolAddr),
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
		IsBuyer:     configs.BuyerNode,
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
	if !configs.DisableStratumv1 {

		src, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", configs.ListenIP, configs.ListenPort))
		if err != nil {
			lumerinlib.PanicHere("")
		}

		dst, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", "127.0.0.1", "3334"))
		if err != nil {
			lumerinlib.PanicHere("")
		}

		stratum, err := stratumv1.NewListener(mainContext, ps, src, dst)
		if err != nil {
			panic(fmt.Sprintf("Stratum Protocol New() failed:%s", err))
		}

		stratum.Run()

	}

	//
	// Fire up schedule manager
	//
	if !configs.DisableSchedule {
		cs, err := connectionscheduler.New(&mainContext, &nodeOperator, configs.SchedulePassthrough)
		if err != nil {
			l.Logf(log.LevelPanic, "Schedule manager failed: %v", err)
		}
		err = cs.Start()
		if err != nil {
			l.Logf(log.LevelPanic, "Schedule manager to start: %v", err)
		}
	}

	return dest.ID
}

func TestMain(t *testing.T) {
	os.Args[1] = "-configfile=../ganacheconfig.json"
	config.Init()
	configs := config.ReadConfigs()

	var sleepTime time.Duration = 3*time.Second

	l := log.New()
	
	logFile, err := os.OpenFile(configs.LogFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		l.Logf(log.LevelFatal, "error opening log file: %v", err)
	}
	defer logFile.Close()
	l.SetFormat(log.FormatJSON).SetOutput(logFile)

	ps := msgbus.New(10, l)

	defaultDestID := SimMain(ps, l, configs)

	//
	// miner connecting to lumerin node
	//
	miner := msgbus.Miner {
		ID:                   msgbus.MinerID("MinerID01"),
		IP:                   "IpAddress1",
		State:                msgbus.OnlineState,
		Dest:                 defaultDestID,
		CsMinerHandlerIgnore: false,
	}

	time.Sleep(sleepTime)

	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner.ID), miner)

	//
	// seller created contract found by lumerin node
	//
	contract := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID01"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    100,
		Length:   1000,
	}

	time.Sleep(sleepTime)

	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract.ID), contract)

	miners,_ := ps.MinerGetAllWait()
	for _,v := range miners {
		miner,_ := ps.MinerGetWait(msgbus.MinerID(v))
		if miner.Contract != "" || miner.Dest != defaultDestID {
			t.Errorf("Miner contract and dest not set correctly")
		}
	}

	//
	// contract was purchased and target dest was inputed in it
	//
	targetDest := msgbus.Dest {
		ID: msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:55555/",
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)

	contract.State = msgbus.ContRunningState
	contract.Buyer = "BuyerID01"
	contract.StartingBlockTimestamp = 63637278298010
	contract.Dest = targetDest.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract.ID), contract)

	time.Sleep(sleepTime)

	miners,_ = ps.MinerGetAllWait()
	for _,v := range miners {
		miner,_ := ps.MinerGetWait(msgbus.MinerID(v))
		if miner.Contract != "ContractID01" || miner.Dest != targetDest.ID {
			t.Errorf("Miner contract and dest not set correctly")
		}
	}


	//
	// target dest was updated while contract running
	//
	targetDest.NetUrl = "stratum+tcp://127.0.0.1:66666/"
	ps.SetWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)

	time.Sleep(sleepTime)

	miners,_ = ps.MinerGetAllWait()
	for _,v := range miners {
		miner,_ := ps.MinerGetWait(msgbus.MinerID(v))
		if miner.Contract != "ContractID01" || miner.Dest != targetDest.ID {
			t.Errorf("Miner contract and dest not set correctly")
		}
	}

	if targetDest.NetUrl != "stratum+tcp://127.0.0.1:66666/" {
		t.Errorf("Target dest was not updated")
	}
}