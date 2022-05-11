package integrationtests

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol/stratumv1"
	"gitlab.com/TitanInd/lumerin/cmd/validator/validator"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

type ValDisabledConfig struct {
	BuyerNode           bool
	ListenIP            string
	ListenPort          string
	DefaultPoolAddr     string
	SchedulePassthrough bool
	LogFilePath         string
}

func LoadValDisabledTestConfiguration(filePath string) (configs ValDisabledConfig, err error) {
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

func ValDisabledSimMain(ps *msgbus.PubSub, configs ValDisabledConfig, hashrateCalcLagTime time.Duration) msgbus.DestID {
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
	cs := contextlib.NewContextStruct(nil, ps, nil, src, dst)

	//
	//  All of the various needed subsystem values get passed into the context here.
	//
	mainContext = context.WithValue(mainContext, contextlib.ContextKey, cs)

	//
	// Setup Default Dest
	//
	dest := &msgbus.Dest{
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

	stratum, err := stratumv1.NewListener(mainContext, srcStrat, dest)
	if err != nil {
		panic(fmt.Sprintf("Stratum Protocol New() failed:%s", err))
	}

	stratum.Run()

	//
	// Fire up schedule manager
	//
	csched, err := connectionscheduler.New(&mainContext, &nodeOperator, false, int(hashrateCalcLagTime))
	if err != nil {
		panic(fmt.Sprintf("Schedule manager failed: %v", err))
	}
	err = csched.Start()
	if err != nil {
		panic(fmt.Sprintf("Schedule manager failed to start: %v", err))
	}

	//
	// Fire up validator
	//
	v := validator.MakeNewValidator(&mainContext)
	err = v.Start()
	if err != nil {
		panic(fmt.Sprintf("Validator failed to start: %v", err))
	}
	
	return dest.ID
}

func TestValDisabled(t *testing.T) {
	configPath := "../../lumerinconfig.json"

	configs, err := LoadValDisabledTestConfiguration(configPath)
	if err != nil {
		panic(fmt.Sprintf("Loading Config Failed: %s", err))
	}

	var sleepTime time.Duration = 10 * time.Second

	ps := msgbus.New(10, nil)

	var hashrateCalcLagTime time.Duration = 20
	var reAdjustmentTime time.Duration = 3

	defaultDestID := ValDisabledSimMain(ps, configs, hashrateCalcLagTime)

	time.Sleep(sleepTime*10)

	fmt.Print("\n\n/// Multiple miners connecting to node ///\n\n\n")

	miner1 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID01"),
		IP:                   "IpAddress1",
		CurrentHashRate:      0,
		State:                msgbus.OnlineState,
		Dest:                 defaultDestID,
		Contracts: 			  make(map[msgbus.ContractID]bool),	
	}
	miner2 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID02"),
		IP:                   "IpAddress2",
		CurrentHashRate:      0,
		State:                msgbus.OnlineState,
		Dest:                 defaultDestID,
		Contracts: 			  make(map[msgbus.ContractID]bool),
	}
	miner3 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID03"),
		IP:                   "IpAddress3",
		CurrentHashRate:      0,
		State:                msgbus.OnlineState,
		Dest:                 defaultDestID,
		Contracts: 			  make(map[msgbus.ContractID]bool),
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// Validator getting submits and updating miner hashrates ///\n\n\n")

	for i:=0;i<10;i++ {
		jobId1 := "JobId1-" + strconv.FormatInt(int64(i), 10)
		jobId2 := "JobId2-" + strconv.FormatInt(int64(i), 10)
		jobId3 := "JobId3-" + strconv.FormatInt(int64(i), 10)
		miningSubmit1 := msgbus.Submit{
			ID: msgbus.SubmitID(jobId1),
			Miner: miner1.ID,
		}
		miningSubmit2 := msgbus.Submit{
			ID: msgbus.SubmitID(jobId2),
			Miner: miner2.ID,
		}
		miningSubmit3 := msgbus.Submit{
			ID: msgbus.SubmitID(jobId3),
			Miner: miner3.ID,
		}
		ps.PubWait(msgbus.SubmitMsg, msgbus.IDString(jobId1), miningSubmit1)
		ps.PubWait(msgbus.SubmitMsg, msgbus.IDString(jobId2), miningSubmit2)
		ps.PubWait(msgbus.SubmitMsg, msgbus.IDString(jobId3), miningSubmit3)
		time.Sleep((time.Second * hashrateCalcLagTime)/10)
	}

	m1,_ := ps.MinerGetWait(miner1.ID)
	m2,_ := ps.MinerGetWait(miner2.ID)
	m3,_ := ps.MinerGetWait(miner3.ID)

	fmt.Println("Miner 1 Current Hashrate: ", m1.CurrentHashRate)
	fmt.Println("Miner 2 Current Hashrate: ", m2.CurrentHashRate)
	fmt.Println("Miner 3 Current Hashrate: ", m3.CurrentHashRate)
	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// 2 New available contracts found ///\n\n\n")

	contract1 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID01"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    150,
	}
	contract2 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID02"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    150,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(time.Second * 2)
	
	fmt.Print("\n\n/// Contract 1 purchased and now running ///\n\n\n")

	targetDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:55555/",
	}
	event,_ := ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)
	targetDestID := msgbus.DestID(event.ID)

	contract1.State = msgbus.ContRunningState
	contract1.Buyer = "buyer1"
	contract1.Dest = targetDest.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(reAdjustmentTime*time.Second)

	var slicedMiner *msgbus.Miner
	var fullMiner1 *msgbus.Miner
	var fullMiner2 *msgbus.Miner
	minerIDs,_ := ps.MinerGetAllWait()
	for _,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		if miner.TimeSlice {
			slicedMiner = miner
		} else if miner.Contracts[contract1.ID] && !miner.TimeSlice{
			fullMiner1 = miner
		} else {
			fullMiner2 = miner
		}
	}
	if slicedMiner.Dest != targetDestID {
		t.Errorf("Sliced miner dest field incorrect")
	}

	if fullMiner1.Dest != targetDestID {
		t.Errorf("Full miner dest field incorrect")
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	slicedMiner,_ = ps.MinerGetWait(slicedMiner.ID)
	fullMiner1,_ = ps.MinerGetWait(fullMiner1.ID)
	if slicedMiner.Dest != defaultDestID {
		t.Errorf("Sliced miner dest field incorrect")
	}

	if fullMiner1.Dest != targetDestID {
		t.Errorf("Full miner dest field incorrect")
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	fmt.Print("\n\n/// Contract 2 purchased and now running ///\n\n\n")

	targetDest2 := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:66666/",
	}
	event,_ = ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest2.ID), targetDest2)
	targetDest2ID := msgbus.DestID(event.ID)

	contract2.State = msgbus.ContRunningState
	contract2.Buyer = "buyer2"
	contract2.Dest = targetDest2.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	slicedMiner,_ = ps.MinerGetWait(slicedMiner.ID)
	fullMiner1,_ = ps.MinerGetWait(fullMiner1.ID)
	fullMiner2,_ = ps.MinerGetWait(fullMiner2.ID)
	if slicedMiner.Dest != targetDestID {
		t.Errorf("Sliced miner dest field incorrect, Dest in Miner: %s", slicedMiner.Dest)
	}

	if fullMiner1.Dest != targetDestID {
		t.Errorf("Full miner 1 dest field incorrect, Dest in Miner: %s", fullMiner1.Dest)
	}

	if fullMiner2.Dest != targetDest2ID {
		t.Errorf("Full miner 2 dest field incorrect, Dest in Miner: %s", fullMiner2.Dest)
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	slicedMiner,_ = ps.MinerGetWait(slicedMiner.ID)
	fullMiner1,_ = ps.MinerGetWait(fullMiner1.ID)
	fullMiner2,_ = ps.MinerGetWait(fullMiner2.ID)

	if slicedMiner.Dest != targetDest2ID {
		t.Errorf("Sliced miner dest field incorrect, Dest in Miner: %s", slicedMiner.Dest)
	}

	if fullMiner1.Dest != targetDestID {
		t.Errorf("Full miner 1 dest field incorrect, Dest in Miner: %s", fullMiner1.Dest)
	}

	if fullMiner2.Dest != targetDest2ID {
		t.Errorf("Full miner 2 dest field incorrect, Dest in Miner: %s", fullMiner2.Dest)
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	fmt.Print("\n\n/// Contract 1 closes out ///\n\n\n")

	contract1.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	slicedMiner,_ = ps.MinerGetWait(slicedMiner.ID)
	fullMiner1,_ = ps.MinerGetWait(fullMiner1.ID)
	fullMiner2,_ = ps.MinerGetWait(fullMiner2.ID)

	if slicedMiner.Dest != targetDest2ID {
		t.Errorf("Sliced miner dest field incorrect, Dest in Miner: %s", slicedMiner.Dest)
	}

	if fullMiner1.Dest != defaultDestID {
		t.Errorf("Full miner 1 dest field incorrect, Dest in Miner: %s", fullMiner1.Dest)
	}

	if fullMiner2.Dest != targetDest2ID {
		t.Errorf("Full miner 2 dest field incorrect, Dest in Miner: %s", fullMiner2.Dest)
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	fmt.Print("\n\n/// Contract 2 closes out ///\n\n\n")

	contract2.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(reAdjustmentTime*time.Second)


	time.Sleep(40*time.Second)
}