package maintest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
	"gitlab.com/TitanInd/lumerin/cmd/log"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol/stratumv1"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

type Config struct {
	BuyerNode           bool
	ListenIP            string
	ListenPort          string
	DefaultPoolAddr     string
	SchedulePassthrough bool
	LogFilePath         string
}

func LoadTestConfiguration(filePath string) (configs Config, err error) {
	var data map[string]interface{}
	currDir, _ := os.Getwd()
	defer os.Chdir(currDir)

	if err != nil {
		panic(fmt.Errorf("error retrieving config file variable: %s", err))
	}
	file := filepath.Base(filePath)
	filePath = filepath.Dir(filePath)
	os.Chdir(filePath)

	configFile, err := os.Open(file)
	if err != nil {
		return configs, err
	}
	defer configFile.Close()
	byteValue, _ := ioutil.ReadAll(configFile)

	err = json.Unmarshal(byteValue, &data)

	configData := data["config"].(map[string]interface{})
	configs.BuyerNode = configData["buyerNode"].(bool)

	connConfigData := data["connection"].(map[string]interface{})
	configs.ListenIP = connConfigData["listenIP"].(string)
	configs.ListenPort = connConfigData["listenPort"].(string)
	configs.DefaultPoolAddr = connConfigData["defaultPoolAddr"].(string)

	schedConfigData := data["schedule"].(map[string]interface{})
	configs.SchedulePassthrough = schedConfigData["passthrough"].(bool)

	logConfigData := data["logging"].(map[string]interface{})
	configs.LogFilePath = logConfigData["filePath"].(string)

	return configs, err
}

func SimMain(ps *msgbus.PubSub, l *log.Logger, configs Config) msgbus.DestID {
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
	srcStrat, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", configs.ListenIP, configs.ListenPort))
	if err != nil {
		lumerinlib.PanicHere("")
	}

	dstStrat, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", "127.0.0.1", "3334"))
	if err != nil {
		lumerinlib.PanicHere("")
	}

	stratum, err := stratumv1.NewListener(mainContext, srcStrat, dstStrat)
	if err != nil {
		panic(fmt.Sprintf("Stratum Protocol New() failed:%s", err))
	}

	stratum.Run()

	//
	// Fire up schedule manager
	//
	csched, err := connectionscheduler.New(&mainContext, &nodeOperator, configs.SchedulePassthrough)
	if err != nil {
		l.Logf(log.LevelPanic, "Schedule manager failed: %v", err)
	}
	err = csched.Start()
	if err != nil {
		l.Logf(log.LevelPanic, "Schedule manager to start: %v", err)
	}

	return dest.ID
}

func TestMain(t *testing.T) {
	configPath := "../../ganacheconfig.json"

	configs, err := LoadTestConfiguration(configPath)
	if err != nil {
		panic(fmt.Sprintf("Loading Config Failed: %s", err))
	}

	var sleepTime time.Duration = 3 * time.Second

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
	fmt.Print("\n\n/// Miner connecting to node ///\n\n\n")

	miner := msgbus.Miner{
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
	fmt.Print("\n\n/// Created contract found by lumerin node ///\n\n\n")

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

	miners, _ := ps.MinerGetAllWait()
	for _, v := range miners {
		miner, _ := ps.MinerGetWait(msgbus.MinerID(v))
		if miner.Contract != "" || miner.Dest != defaultDestID {
			t.Errorf("Miner contract and dest not set correctly")
		}
	}

	//
	// contract was purchased and target dest was inputed in it
	//
	fmt.Print("\n\n/// Contract was purchased and target dest was inputed in it ///\n\n\n")

	targetDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:55555/",
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)

	contract.State = msgbus.ContRunningState
	contract.Buyer = "BuyerID01"
	contract.StartingBlockTimestamp = 63637278298010
	contract.Dest = targetDest.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract.ID), contract)

	time.Sleep(sleepTime)

	miners, _ = ps.MinerGetAllWait()
	for _, v := range miners {
		miner, _ := ps.MinerGetWait(msgbus.MinerID(v))
		if miner.Contract != "ContractID01" || miner.Dest != targetDest.ID {
			t.Errorf("Miner contract and dest not set correctly")
		}
	}

	//
	// target dest was updated while contract running
	//
	fmt.Print("\n\n/// Target dest was updated while contract running ///\n\n\n")

	targetDest.NetUrl = "stratum+tcp://127.0.0.1:66666/"
	ps.SetWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)

	time.Sleep(sleepTime)

	miners, _ = ps.MinerGetAllWait()
	for _, v := range miners {
		miner, _ := ps.MinerGetWait(msgbus.MinerID(v))
		if miner.Contract != "ContractID01" || miner.Dest != targetDest.ID {
			t.Errorf("Miner contract and dest not set correctly")
		}
	}

	if targetDest.NetUrl != "stratum+tcp://127.0.0.1:66666/" {
		t.Errorf("Target dest was not updated")
	}
}
