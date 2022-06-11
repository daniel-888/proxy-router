package connectionscheduler

import (
	"context"
	"fmt"
	"testing"
	"time"
<<<<<<< HEAD
	
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

func TestSellerConnectionScheduler(t *testing.T) {
	ps := msgbus.New(10, nil)
=======

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/log"
	"gitlab.com/TitanInd/lumerin/connections"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

func TestPassthroughConnectionScheduler(t *testing.T) {
	l := log.New()
	ps := msgbus.New(10, l)
>>>>>>> pr-009

	ctxStruct := contextlib.NewContextStruct(nil, ps, nil, nil, nil)
	mainCtx := context.WithValue(context.Background(), contextlib.ContextKey, ctxStruct)

	defaultpooladdr := "stratum+tcp://127.0.0.1:33334/"
	defaultDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetUrl: msgbus.DestNetUrl(defaultpooladdr),
	}
	event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(msgbus.DEFAULT_DEST_ID), defaultDest)
	if err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", event.Err))
	}

	nodeOperator := msgbus.NodeOperator{
		ID:          msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
		DefaultDest: defaultDest.ID,
		IsBuyer:     false,
	}

<<<<<<< HEAD
	cs, err := New(&mainCtx, &nodeOperator, false)
=======
	cs, err := New(&mainCtx, &nodeOperator, true, 2, connections.CreateConnectionCollection())
>>>>>>> pr-009
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed:%s", err))
	}
	err = cs.Start()
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed to start:%s", err))
	}

<<<<<<< HEAD
	//
	// 1 miner and 1 contract with hashrate within 10% tolerance
	//
	fmt.Print("\n\n/// 1 miner and 1 contract with hashrate within 10% tolerance ///\n\n\n")
	miner1Hashrate := 100
	miner1 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID01"),
		IP:                   "IpAddress1",
		CurrentHashRate:      miner1Hashrate - 5, // 95
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)

	time.Sleep(time.Second * 2)

=======
	fmt.Print("\n\n/// Multiple miners connecting to node ///\n\n\n")

	miner1 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID01"),
		IP:              "IpAddress1",
		CurrentHashRate: 27,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	miner2 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID02"),
		IP:              "IpAddress2",
		CurrentHashRate: 35,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	miner3 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID03"),
		IP:              "IpAddress3",
		CurrentHashRate: 72,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// New available contract found ///\n\n\n")

>>>>>>> pr-009
	contract1 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID01"),
		State:    msgbus.ContAvailableState,
<<<<<<< HEAD
		Price:    0,
		Limit:    10,
		Speed:    miner1Hashrate,
=======
		Price:    10,
		Limit:    10,
		Speed:    100,
>>>>>>> pr-009
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)

	time.Sleep(time.Second * 2)

<<<<<<< HEAD
	contract1.State = msgbus.ContRunningState
	contract1.Buyer = "buyer"
	contract1.Dest = "stratum+tcp://127.0.0.1:44444/"

	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)

	time.Sleep(time.Second * 2)

	if cs.ReadyMiners.Exists(string(miner1.ID)) {
		t.Errorf("Miner 1 was not removed from ready miners map")
	}
	if !cs.BusyMiners.Exists(string(miner1.ID)) {
		t.Errorf("Miner 1 was not moved to busy miners map")
	}

	miner1GET, err := ps.MinerGetWait(miner1.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 1:%s", err))
	}
	if miner1GET.Contract != contract1.ID || miner1GET.Dest != contract1.Dest {
		t.Errorf("Scheduler did not update miner 1 with new dest and contract in msgbus")
	}

	//
	// miner 1 updated to fall out of tolerance range
	//
	fmt.Print("\n\n/// miner 1 updated to fall out of tolerance range ///\n\n\n")
	event, _ = ps.GetWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID))
	miner1 = event.Data.(msgbus.Miner)
	miner1.CsMinerHandlerIgnore = false
	miner1.CurrentHashRate = 50
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	time.Sleep(time.Second * 2)

	if !cs.ReadyMiners.Exists(string(miner1.ID)) {
		t.Errorf("Miner 1 was not moved back to ready miners map")
	}
	if cs.BusyMiners.Exists(string(miner1.ID)) {
		t.Errorf("Miner 1 was not removed from busy miners map")
	}

	miner1GET, err = ps.MinerGetWait(miner1.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 1:%s", err))
	}
	if miner1GET.Contract != "" || miner1GET.Dest != nodeOperator.DefaultDest {
		t.Errorf("Scheduler did not update miner 1 with default dest and empty contract param")
	}
	contract1.State = msgbus.ContAvailableState
	contract1.Buyer = ""
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(time.Second * 2)

	//
	// Publish multiple miners and find best combination to point to running contract dest
	//
	fmt.Print("\n\n/// Publish multiple miners and find best combination to point to running contract dest ///\n\n\n")
	miner2 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID02"),
		IP:                   "IpAddress2",
		CurrentHashRate:      35,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner3 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID03"),
		IP:                   "IpAddress3",
		CurrentHashRate:      72,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner4 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID04"),
		IP:                   "IpAddress4",
		CurrentHashRate:      16,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner5 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID05"),
		IP:                   "IpAddress5",
		CurrentHashRate:      88,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner6 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID06"),
		IP:                   "IpAddress6",
		CurrentHashRate:      27,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner4.ID), miner4)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner5.ID), miner5)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner6.ID), miner6)

	contract1.State = msgbus.ContRunningState
	contract1.Buyer = "buyer"
	contract1.Dest = "stratum+tcp://127.0.0.1:44444/"

	time.Sleep(time.Second * 3)

	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)

	time.Sleep(time.Second * 3)

	// Best miner combo should be miner 3 and 6
	correctReadyMiners := []msgbus.Miner{miner1, miner2, miner4, miner5}
	correctBusyMiners := []msgbus.Miner{miner3, miner6}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
	}

	miner3GET, err := ps.MinerGetWait(miner3.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 3:%s", err))
	}
	miner6GET, err := ps.MinerGetWait(miner6.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 6:%s", err))
	}
	if miner3GET.Contract != contract1.ID || miner3GET.Dest != contract1.Dest {
		t.Errorf("Scheduler did not update miner 3 with new dest and contract in msgbus")
	}
	if miner6GET.Contract != contract1.ID || miner6GET.Dest != contract1.Dest {
		t.Errorf("Scheduler did not update miner 6 with new dest and contract in msgbus")
	}

	//
	// Publish new miner and update another that creates new best combination
	//
	fmt.Print("\n\n/// Publish new miner and update another that creates new best combination ///\n\n\n")
	event, _ = ps.GetWait(msgbus.MinerMsg, msgbus.IDString(miner5.ID))
	miner5 = event.Data.(msgbus.Miner)
	miner5.CsMinerHandlerIgnore = false
	miner5.CurrentHashRate = 80
	miner7 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID07"),
		IP:              "IpAddress7",
		CurrentHashRate: 20,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		//NodeOperator: nodeOperator.ID,
	}
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner5.ID), miner5)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner7.ID), miner7)

	time.Sleep(time.Second * 3)

	// Best miner combo should be miner 5 and 7
	correctReadyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4, miner6}
	correctBusyMiners = []msgbus.Miner{miner5, miner7}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
	}

	miner5GET, err := ps.MinerGetWait(miner5.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 5:%s", err))
	}
	miner7GET, err := ps.MinerGetWait(miner7.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 7:%s", err))
	}
	if miner5GET.Contract != contract1.ID || miner7GET.Dest != contract1.Dest {
		t.Errorf("Scheduler did not update miner 5 with new dest and contract in msgbus")
	}
	if miner7GET.Contract != contract1.ID || miner7GET.Dest != contract1.Dest {
		t.Errorf("Scheduler did not update miner 7 with new dest and contract in msgbus")
	}

	miner3GET, err = ps.MinerGetWait(miner3.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 3:%s", err))
	}
	if miner3GET.Contract != "" || miner3GET.Dest != nodeOperator.DefaultDest {
		t.Errorf("Scheduler did not update miner 3 with default dest and empty contract param")
	}
	miner6GET, err = ps.MinerGetWait(miner6.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 6:%s", err))
	}
	if miner6GET.Contract != "" || miner6GET.Dest != nodeOperator.DefaultDest {
		t.Errorf("Scheduler did not update miner 6 with default dest and empty contract param")
	}

	//fmt.Println("Ready Miners: ", cs.ReadyMiners.M)
	//fmt.Println("Busy Miners: ", cs.BusyMiners.M)
	time.Sleep(time.Second * 2)

	//
	// Another contract is created and purchased with different dest
	//
	fmt.Print("\n\n/// Another contract is created and purchased with different dest ///\n\n\n")
=======
	fmt.Print("\n\n/// Contract purchased and now running ///\n\n\n")

	targetDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:55555/",
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)

	contract1.State = msgbus.ContRunningState
	contract1.Buyer = "buyer"
	contract1.Dest = targetDest.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(time.Second * 2)

	miners, _ := ps.MinerGetAllWait()

	for _, m := range miners {
		miner, _ := ps.MinerGetWait(m)
		if _,ok :=miner.Contracts[contract1.ID]; !ok || miner.Dest != contract1.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// New miner connected ///\n\n\n")

	miner4 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID04"),
		IP:              "IpAddress4",
		CurrentHashRate: 88,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner4.ID), miner4)
	time.Sleep(time.Second * 5)

	miner, _ := ps.MinerGetWait(miner4.ID)
	if _,ok := miner.Contracts[contract1.ID]; !ok || miner.Dest != contract1.Dest {
		t.Errorf("Miner contract and dest field incorrect")
	}

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// Contract closes out ///\n\n\n")

	contract1.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(time.Second * 5)

	miners, _ = ps.MinerGetAllWait()

	for _, m := range miners {
		miner, _ := ps.MinerGetWait(m)
		if len(miner.Contracts) != 0 || miner.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)
}

/*
On Ropsten, demonstrate one seller, three miners, two contracts (valued at 1+partial miner capacity),
purchased by two separate buyers concurrently and can show purchased hashrate consistently for the duration of both contracts
*/
func TestTimeSlicing(t *testing.T) {
	l := log.New()
	ps := msgbus.New(10, l)

	ctxStruct := contextlib.NewContextStruct(nil, ps, nil, nil, nil)
	mainCtx := context.WithValue(context.Background(), contextlib.ContextKey, ctxStruct)

	defaultpooladdr := "stratum+tcp://127.0.0.1:33334/"
	defaultDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetUrl: msgbus.DestNetUrl(defaultpooladdr),
	}
	event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(msgbus.DEFAULT_DEST_ID), defaultDest)
	if err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", event.Err))
	}

	nodeOperator := msgbus.NodeOperator{
		ID:          msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
		DefaultDest: defaultDest.ID,
		IsBuyer:     false,
	}

	var hashrateCalcLagTime time.Duration = 20
	var reAdjustmentTime time.Duration = 3

	cs, err := New(&mainCtx, &nodeOperator, false, int(hashrateCalcLagTime), connections.CreateConnectionCollection())
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed:%s", err))
	}
	err = cs.Start()
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed to start:%s", err))
	}

	fmt.Print("\n\n/// Multiple miners connecting to node ///\n\n\n")

	miner1 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID01"),
		IP:              "IpAddress1",
		CurrentHashRate: 0,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	miner2 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID02"),
		IP:              "IpAddress2",
		CurrentHashRate: 0,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	miner3 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID03"),
		IP:              "IpAddress3",
		CurrentHashRate: 0,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)

	time.Sleep(time.Second * hashrateCalcLagTime)

	fmt.Print("\n\n/// Validator updated miner hashrates ///\n\n\n")

	miner1.CurrentHashRate = 100
	miner2.CurrentHashRate = 100
	miner3.CurrentHashRate = 100
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
	time.Sleep(time.Second * 1)

	fmt.Print("\n\n/// 2 New available contracts found ///\n\n\n")

	contract1 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID01"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    150,
	}
>>>>>>> pr-009
	contract2 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID02"),
		State:    msgbus.ContAvailableState,
<<<<<<< HEAD
		Price:    0,
		Limit:    10,
		Speed:    52,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)

	time.Sleep(time.Second * 2)

	contract2.State = msgbus.ContRunningState
	contract2.Buyer = "buyer"
	contract2.Dest = "stratum+tcp://127.0.0.1:55555/"

	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)

	time.Sleep(time.Second * 2)

	// Best miner combo should be miner 5 and 7 for contract 1 and miner 2 and 4 for contract 2
	correctReadyMiners = []msgbus.Miner{miner1, miner3, miner6}
	correctBusyMiners = []msgbus.Miner{miner2, miner4, miner5, miner7}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
	}

	//fmt.Println("Ready Miners: ", cs.ReadyMiners.M)
	//fmt.Println("Busy Miners: ", cs.BusyMiners.M)
	time.Sleep(time.Second * 2)
}

func TestBuyerConnectionScheduler(t *testing.T) {
	ps := msgbus.New(10, nil)
=======
		Price:    10,
		Limit:    10,
		Speed:    150,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)

	fmt.Print("\n\n/// Contract 1 purchased and now running ///\n\n\n")

	targetDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:55555/",
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)

	contract1.State = msgbus.ContRunningState
	contract1.Buyer = "buyer1"
	contract1.Dest = targetDest.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 1--\n")
	minerIDs,_ := ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 2--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	fmt.Print("\n\n/// Contract 2 purchased and now running ///\n\n\n")

	targetDest2 := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:66666/",
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest2.ID), targetDest2)

	contract2.State = msgbus.ContRunningState
	contract2.Buyer = "buyer2"
	contract2.Dest = targetDest2.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 1--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 2--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	fmt.Print("\n\n/// Contract 1 closes out ///\n\n\n")

	contract1.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	minerIDs,_ = ps.MinerGetAllWait()
	fmt.Println()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	fmt.Print("\n\n/// Contract 2 closes out ///\n\n\n")

	contract2.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	minerIDs,_ = ps.MinerGetAllWait()
	fmt.Println()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()
}

func TestMultiTimeSlicing(t *testing.T) {
	l := log.New()
	ps := msgbus.New(10, l)
>>>>>>> pr-009

	ctxStruct := contextlib.NewContextStruct(nil, ps, nil, nil, nil)
	mainCtx := context.WithValue(context.Background(), contextlib.ContextKey, ctxStruct)

	defaultpooladdr := "stratum+tcp://127.0.0.1:33334/"
	defaultDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetUrl: msgbus.DestNetUrl(defaultpooladdr),
	}
	event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(msgbus.DEFAULT_DEST_ID), defaultDest)
	if err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", event.Err))
	}

	nodeOperator := msgbus.NodeOperator{
		ID:          msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
		DefaultDest: defaultDest.ID,
<<<<<<< HEAD
		IsBuyer:     true,
	}

	cs, err := New(&mainCtx, &nodeOperator, false)
=======
		IsBuyer:     false,
	}

	var hashrateCalcLagTime time.Duration = 20
	var reAdjustmentTime time.Duration = 3

	cs, err := New(&mainCtx, &nodeOperator, false, int(hashrateCalcLagTime), connections.CreateConnectionCollection())
>>>>>>> pr-009
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed:%s", err))
	}
	err = cs.Start()
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed to start:%s", err))
	}

<<<<<<< HEAD
	//
	// Publish multiple miners with varying hashrate
	//
	miner1 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID01"),
		IP:                   "IpAddress1",
		CurrentHashRate:      27,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner2 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID02"),
		IP:                   "IpAddress2",
		CurrentHashRate:      35,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner3 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID03"),
		IP:                   "IpAddress3",
		CurrentHashRate:      72,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner4 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID04"),
		IP:                   "IpAddress4",
		CurrentHashRate:      16,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner5 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID05"),
		IP:                   "IpAddress5",
		CurrentHashRate:      88,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
=======

	fmt.Print("\n\n/// Multiple miners connecting to node ///\n\n\n")

	miner1 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID01"),
		IP:              "IpAddress1",
		CurrentHashRate: 0,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	miner2 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID02"),
		IP:              "IpAddress2",
		CurrentHashRate: 0,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	miner3 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID03"),
		IP:              "IpAddress3",
		CurrentHashRate: 0,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
>>>>>>> pr-009
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
<<<<<<< HEAD
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner4.ID), miner4)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner5.ID), miner5)

	time.Sleep(time.Second * 2)

	contract1 := msgbus.Contract{
		IsSeller: false,
		ID:       msgbus.ContractID("ContractID01"),
		State:    msgbus.ContRunningState,
		Price:    0,
		Limit:    10,
		Speed:    100,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)

	time.Sleep(time.Second * 2)

	contract2 := msgbus.Contract{
		IsSeller: false,
		ID:       msgbus.ContractID("ContractID02"),
		State:    msgbus.ContRunningState,
		Price:    0,
		Limit:    10,
		Speed:    100,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)

	time.Sleep(time.Second * 2)

	contract3 := msgbus.Contract{
		IsSeller: false,
		ID:       msgbus.ContractID("ContractID03"),
		State:    msgbus.ContRunningState,
		Price:    0,
		Limit:    0,
		Speed:    100,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract3.ID), contract3)

	time.Sleep(time.Second * 2)

	miner1.CurrentHashRate = 2
	miner2.CurrentHashRate = 2
	miner3.CurrentHashRate = 2
	miner4.CurrentHashRate = 2
	miner5.CurrentHashRate = 2
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner4.ID), miner4)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner5.ID), miner5)

	time.Sleep(time.Second * 5)
}

func TestPassthroughConnectionScheduler(t *testing.T) {
	ps := msgbus.New(10, nil)
=======

	time.Sleep(time.Second * hashrateCalcLagTime)


	fmt.Print("\n\n/// Validator updated miner hashrates ///\n\n\n")

	miner1.CurrentHashRate = 100
	miner2.CurrentHashRate = 100
	miner3.CurrentHashRate = 100
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
	time.Sleep(time.Second * 1)


	fmt.Print("\n\n/// 2 New available contracts found ///\n\n\n")

	contract1 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID01"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    175,
	}
	contract2 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID02"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    125,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)


	fmt.Print("\n\n/// Contract 1 purchased and now running ///\n\n\n")

	targetDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:55555/",
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)

	contract1.State = msgbus.ContRunningState
	contract1.Buyer = "buyer1"
	contract1.Dest = targetDest.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 1--\n")
	minerIDs,_ := ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((3*hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 2--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)


	fmt.Print("\n\n/// Contract 2 purchased and now running ///\n\n\n")

	targetDest2 := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:66666/",
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest2.ID), targetDest2)

	contract2.State = msgbus.ContRunningState
	contract2.Buyer = "buyer2"
	contract2.Dest = targetDest2.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 1--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((3*hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 2--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)


	fmt.Print("\n\n/// Contract 1 closes out ///\n\n\n")

	contract1.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	minerIDs,_ = ps.MinerGetAllWait()
	fmt.Println()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)


	fmt.Print("\n\n/// Contract 2 closes out ///\n\n\n")

	contract2.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	minerIDs,_ = ps.MinerGetAllWait()
	fmt.Println()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()
}

func TestEdgeCases (t *testing.T) {
	l := log.New()
	ps := msgbus.New(10, l)
>>>>>>> pr-009

	ctxStruct := contextlib.NewContextStruct(nil, ps, nil, nil, nil)
	mainCtx := context.WithValue(context.Background(), contextlib.ContextKey, ctxStruct)

	defaultpooladdr := "stratum+tcp://127.0.0.1:33334/"
	defaultDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetUrl: msgbus.DestNetUrl(defaultpooladdr),
	}
	event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(msgbus.DEFAULT_DEST_ID), defaultDest)
	if err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", event.Err))
	}

	nodeOperator := msgbus.NodeOperator{
		ID:          msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
		DefaultDest: defaultDest.ID,
<<<<<<< HEAD
		IsBuyer:     true,
	}

	cs, err := New(&mainCtx, &nodeOperator, true)
=======
		IsBuyer:     false,
	}

	var hashrateCalcLagTime time.Duration = 20
	var reAdjustmentTime time.Duration = 3

	cs, err := New(&mainCtx, &nodeOperator, false, int(hashrateCalcLagTime), connections.CreateConnectionCollection())
>>>>>>> pr-009
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed:%s", err))
	}
	err = cs.Start()
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed to start:%s", err))
	}

<<<<<<< HEAD
	fmt.Print("\n\n/// Multiple miners connecting to node ///\n\n\n")

	miner1 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID01"),
		IP:                   "IpAddress1",
		CurrentHashRate:      27,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner2 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID02"),
		IP:                   "IpAddress2",
		CurrentHashRate:      35,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner3 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID03"),
		IP:                   "IpAddress3",
		CurrentHashRate:      72,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner4 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID04"),
		IP:                   "IpAddress4",
		CurrentHashRate:      16,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner5 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID05"),
		IP:                   "IpAddress5",
		CurrentHashRate:      88,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
=======

	fmt.Print("\n\n/// Multiple miners connecting to node ///\n\n\n")

	miner1 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID01"),
		IP:              "IpAddress1",
		CurrentHashRate: 0,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	miner2 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID02"),
		IP:              "IpAddress2",
		CurrentHashRate: 0,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
	}
	miner3 := msgbus.Miner{
		ID:              msgbus.MinerID("MinerID03"),
		IP:              "IpAddress3",
		CurrentHashRate: 0,
		State:           msgbus.OnlineState,
		Dest:            defaultDest.ID,
		Contracts:       make(map[msgbus.ContractID]float64),
>>>>>>> pr-009
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
<<<<<<< HEAD
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner4.ID), miner4)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner5.ID), miner5)

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// New available contract found ///\n\n\n")
=======

	time.Sleep(time.Second * hashrateCalcLagTime)


	fmt.Print("\n\n/// Validator updated miner hashrates ///\n\n\n")

	miner1.CurrentHashRate = 200
	miner2.CurrentHashRate = 140
	miner3.CurrentHashRate = 140
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
	time.Sleep(time.Second * 1)


	fmt.Print("\n\n/// 2 New available contracts found ///\n\n\n")
>>>>>>> pr-009

	contract1 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID01"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
<<<<<<< HEAD
		Speed:    100,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)

	correctReadyMiners := []msgbus.Miner{miner1, miner2, miner3, miner4, miner5}
	correctBusyMiners := []msgbus.Miner{}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
		if v.Contract != "" || v.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
		if v.Contract != contract1.ID || v.Dest != contract1.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// Contract purchased and now running ///\n\n\n")
=======
		Speed:    270,
	}
	contract2 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID02"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    110,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)


	fmt.Print("\n\n/// Contract 1 purchased and now running ///\n\n\n")
>>>>>>> pr-009

	targetDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:55555/",
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)

	contract1.State = msgbus.ContRunningState
<<<<<<< HEAD
	contract1.Buyer = "buyer"
	contract1.Dest = targetDest.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(time.Second * 2)

	correctReadyMiners = []msgbus.Miner{}
	correctBusyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4, miner5}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != "" || miner.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != contract1.ID || miner.Dest != contract1.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// New miner connected ///\n\n\n")

	miner6 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID06"),
		IP:                   "IpAddress6",
		CurrentHashRate:      88,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner6.ID), miner6)
	time.Sleep(time.Second * 2)

	correctReadyMiners = []msgbus.Miner{}
	correctBusyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4, miner5, miner6}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != "" || miner.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != contract1.ID || miner.Dest != contract1.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// Contract closes out ///\n\n\n")

	contract1.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(time.Second * 2)

	correctReadyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4, miner5, miner6}
	correctBusyMiners = []msgbus.Miner{}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != "" || miner.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != contract1.ID || miner.Dest != contract1.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// New available contract found ///\n\n\n")

	contract2 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID02"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    100,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)

	correctReadyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4, miner5, miner6}
	correctBusyMiners = []msgbus.Miner{}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
		if v.Contract != "" || v.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
		if v.Contract != contract2.ID || v.Dest != contract2.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// Contract purchased and now running ///\n\n\n")

	contract2.State = msgbus.ContRunningState
	contract2.Buyer = "buyer"
	contract2.Dest = "stratum+tcp://127.0.0.1:55555/"
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(time.Second * 2)

	correctReadyMiners = []msgbus.Miner{}
	correctBusyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4, miner5, miner6}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != "" || miner.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != contract2.ID || miner.Dest != contract2.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// Few miners disconnect ///\n\n\n")

	miner5.State = msgbus.OfflineState
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner5.ID), miner5)
	ps.UnpubWait(msgbus.MinerMsg, msgbus.IDString(miner6.ID))
	time.Sleep(time.Second * 2)

	correctReadyMiners = []msgbus.Miner{}
	correctBusyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != "" || miner.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != contract2.ID || miner.Dest != contract2.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// Contract Target Dest updated ///\n\n\n")

	targetDest.NetUrl = "stratum+tcp://127.0.0.1:66666/"
	ps.SetWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)

	time.Sleep(time.Second * 2)

	correctReadyMiners = []msgbus.Miner{}
	correctBusyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != "" || miner.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != contract2.ID || miner.Dest != contract2.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// Contract Closed Out ///\n\n\n")

	contract2.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(time.Second * 2)

	correctReadyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4}
	correctBusyMiners = []msgbus.Miner{}

	for _, v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(string(v.ID)) {
			t.Errorf("Ready miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != "" || miner.Dest != nodeOperator.DefaultDest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}
	for _, v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(string(v.ID)) {
			t.Errorf("Busy miners map not correct")
		}
		miner, _ := ps.MinerGetWait(v.ID)
		if miner.Contract != contract2.ID || miner.Dest != contract2.Dest {
			t.Errorf("Miner contract and dest field incorrect")
		}
	}

	time.Sleep(time.Second * 2)
}
=======
	contract1.Buyer = "buyer1"
	contract1.Dest = targetDest.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 1--\n")
	minerIDs,_ := ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((3*hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 2--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)


	fmt.Print("\n\n/// Miner disconnects and reconnects ///\n\n\n")

	time.Sleep(reAdjustmentTime*time.Second)

	ps.UnpubWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID))
	time.Sleep(reAdjustmentTime*time.Second)
	
	miner3.ID = msgbus.MinerID("MinerID03-Reconnect")
	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 1--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((3*hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 2--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)


	fmt.Print("\n\n/// Contract 2 purchased and now running ///\n\n\n")

	targetDest2 := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:66666/",
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest2.ID), targetDest2)

	contract2.State = msgbus.ContRunningState
	contract2.Buyer = "buyer2"
	contract2.Dest = targetDest2.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 1--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((3*hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	fmt.Print("\n--Time Slice 2--\n")
	minerIDs,_ = ps.MinerGetAllWait()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)


	fmt.Print("\n\n/// Contract 1 closes out in middle of slice period ///\n\n\n")

	time.Sleep(reAdjustmentTime*time.Second)
	contract1.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	minerIDs,_ = ps.MinerGetAllWait()
	fmt.Println()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()

	time.Sleep((hashrateCalcLagTime/4 - reAdjustmentTime)*time.Second)


	fmt.Print("\n\n/// Contract 2 closes out ///\n\n\n")

	contract2.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	minerIDs,_ = ps.MinerGetAllWait()
	fmt.Println()
	for i,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		fmt.Printf("Miner%d: %v\n", i, miner)
	}
	fmt.Println()
}
>>>>>>> pr-009
