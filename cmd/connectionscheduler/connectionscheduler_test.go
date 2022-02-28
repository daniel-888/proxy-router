package connectionscheduler

import (
	"testing"
	"context"
	"fmt"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func TestSellerConnectionScheduler(t *testing.T) {
	ps := msgbus.New(10)
	mainCtx := context.Background()

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
		ID: msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
		DefaultDest: defaultDest.ID,
		IsBuyer: false,
	}

	cs, err := New(&mainCtx, ps, &nodeOperator)
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed:%s", err))
	}
	err = cs.Start()
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed to start:%s", err))
	}

	//
	// 1 miner and 1 contract with hashrate within 10% tolerance
	//
	miner1Hashrate := 100
	miner1 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID01"),
		IP: 	"IpAddress1",
		CurrentHashRate:	miner1Hashrate - 5, // 95
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner1.ID),miner1)

	time.Sleep(time.Second*2)

	contract1 := msgbus.Contract {
		IsSeller: true,
		ID: msgbus.ContractID("ContractID01"),
		State: msgbus.ContAvailableState,
		Price: 0,
		Limit: 0,
		Speed: miner1Hashrate,
	}
	ps.PubWait(msgbus.ContractMsg,msgbus.IDString(contract1.ID),contract1)

	time.Sleep(time.Second*2)

	contract1.State = msgbus.ContRunningState
	contract1.Buyer = "buyer"
	contract1.Dest = "stratum+tcp://127.0.0.1:44444/"

	ps.SetWait(msgbus.ContractMsg,msgbus.IDString(contract1.ID),contract1)

	time.Sleep(time.Second*2)
	
	if cs.ReadyMiners.Exists(miner1.ID) {
		t.Errorf("Miner 1 was not removed from ready miners map")
	}
	if !cs.BusyMiners.Exists(miner1.ID) {
		t.Errorf("Miner 1 was not moved to busy miners map")
	}

	miner1GET,err := ps.MinerGetWait(miner1.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 1:%s", err))
	}
	if (miner1GET.Contract != contract1.ID || miner1GET.Dest != contract1.Dest) {
		t.Errorf("Scheduler did not update miner 1 with new dest and contract in msgbus")
	}

	//
	// miner 1 updated to fall out of tolerance range
	//
	event,_ = ps.GetWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID))
	miner1 = event.Data.(msgbus.Miner)
	miner1.CsMinerHandlerIgnore = false
	miner1.CurrentHashRate = 50
	ps.SetWait(msgbus.MinerMsg,msgbus.IDString(miner1.ID),miner1)
	time.Sleep(time.Second*2)

	if !cs.ReadyMiners.Exists(miner1.ID) {
		t.Errorf("Miner 1 was not moved back to ready miners map")
	}
	if cs.BusyMiners.Exists(miner1.ID) {
		t.Errorf("Miner 1 was not removed from busy miners map")
	}

	miner1GET,err = ps.MinerGetWait(miner1.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 1:%s", err))
	}
	if (miner1GET.Contract != "" || miner1GET.Dest != nodeOperator.DefaultDest) {
		t.Errorf("Scheduler did not update miner 1 with default dest and empty contract param")
	}
	contract1.State = msgbus.ContAvailableState
	contract1.Buyer = ""
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(time.Second*2)

	//
	// Publish multiple miners and find best combination to point to running contract dest
	//
	miner2 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID02"),
		IP: 	"IpAddress2",
		CurrentHashRate:	35,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner3 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID03"),
		IP: 	"IpAddress3",
		CurrentHashRate:	72,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner4 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID04"),
		IP: 	"IpAddress4",
		CurrentHashRate:	16,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner5 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID05"),
		IP: 	"IpAddress5",
		CurrentHashRate:	88,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner6 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID06"),
		IP: 	"IpAddress6",
		CurrentHashRate:	27,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner2.ID),miner2)
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner3.ID),miner3)
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner4.ID),miner4)
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner5.ID),miner5)
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner6.ID),miner6)

	contract1.State = msgbus.ContRunningState
	contract1.Buyer = "buyer"
	contract1.Dest = "stratum+tcp://127.0.0.1:44444/"

	time.Sleep(time.Second*3)

	ps.SetWait(msgbus.ContractMsg,msgbus.IDString(contract1.ID),contract1)

	time.Sleep(time.Second*3)

	// Best miner combo should be miner 3 and 6
	correctReadyMiners := []msgbus.Miner{miner1, miner2, miner4, miner5}
	correctBusyMiners := []msgbus.Miner{miner3, miner6}

	for _,v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(v.ID) {
			t.Errorf("Ready miners map not correct")
		}
	}
	for _,v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(v.ID) {
			t.Errorf("Busy miners map not correct")
		}
	}

	miner3GET,err := ps.MinerGetWait(miner3.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 3:%s", err))
	}
	miner6GET,err := ps.MinerGetWait(miner6.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 6:%s", err))
	}
	if (miner3GET.Contract != contract1.ID || miner3GET.Dest != contract1.Dest) {
		t.Errorf("Scheduler did not update miner 3 with new dest and contract in msgbus")
	}
	if (miner6GET.Contract != contract1.ID || miner6GET.Dest != contract1.Dest) {
		t.Errorf("Scheduler did not update miner 6 with new dest and contract in msgbus")
	}

	//
	// Publish new miner and update another that creates new best combination
	//
	event,_ = ps.GetWait(msgbus.MinerMsg, msgbus.IDString(miner5.ID))
	miner5 = event.Data.(msgbus.Miner)
	miner5.CsMinerHandlerIgnore = false
	miner5.CurrentHashRate = 80
	miner7 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID07"),
		IP: 	"IpAddress7",
		CurrentHashRate:	20,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		//NodeOperator: nodeOperator.ID,
	}
	ps.SetWait(msgbus.MinerMsg,msgbus.IDString(miner5.ID),miner5)
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner7.ID),miner7)

	time.Sleep(time.Second*3)

	// Best miner combo should be miner 5 and 7
	correctReadyMiners = []msgbus.Miner{miner1, miner2, miner3, miner4, miner6}
	correctBusyMiners = []msgbus.Miner{miner5, miner7}

	for _,v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(v.ID) {
			t.Errorf("Ready miners map not correct")
		}
	}
	for _,v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(v.ID) {
			t.Errorf("Busy miners map not correct")
		}
	}

	miner5GET,err := ps.MinerGetWait(miner5.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 5:%s", err))
	}
	miner7GET,err := ps.MinerGetWait(miner7.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 7:%s", err))
	}
	if (miner5GET.Contract != contract1.ID || miner7GET.Dest != contract1.Dest) {
		t.Errorf("Scheduler did not update miner 5 with new dest and contract in msgbus")
	}
	if (miner7GET.Contract != contract1.ID || miner7GET.Dest != contract1.Dest) {
		t.Errorf("Scheduler did not update miner 7 with new dest and contract in msgbus")
	}

	miner3GET,err = ps.MinerGetWait(miner3.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 3:%s", err))
	}
	if (miner3GET.Contract != "" || miner3GET.Dest != nodeOperator.DefaultDest) {
		t.Errorf("Scheduler did not update miner 3 with default dest and empty contract param")
	}
	miner6GET,err = ps.MinerGetWait(miner6.ID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get miner 6:%s", err))
	}
	if (miner6GET.Contract != "" || miner6GET.Dest != nodeOperator.DefaultDest) {
		t.Errorf("Scheduler did not update miner 6 with default dest and empty contract param")
	}

	fmt.Println("Ready Miners: ", cs.ReadyMiners.m)
	fmt.Println("Busy Miners: ", cs.BusyMiners.m)
	time.Sleep(time.Second*2)

	//
	// Another contract is created and purchased with different dest
	//
	contract2 := msgbus.Contract {
		IsSeller: true,
		ID: msgbus.ContractID("ContractID02"),
		State: msgbus.ContAvailableState,
		Price: 0,
		Limit: 0,
		Speed: 52,
	}
	ps.PubWait(msgbus.ContractMsg,msgbus.IDString(contract2.ID),contract2)

	time.Sleep(time.Second*2)

	contract2.State = msgbus.ContRunningState
	contract2.Buyer = "buyer"
	contract2.Dest = "stratum+tcp://127.0.0.1:55555/"

	ps.SetWait(msgbus.ContractMsg,msgbus.IDString(contract2.ID),contract2)

	time.Sleep(time.Second*2)

	// Best miner combo should be miner 5 and 7 for contract 1 and miner 2 and 4 for contract 2
	correctReadyMiners = []msgbus.Miner{miner1, miner3, miner6}
	correctBusyMiners = []msgbus.Miner{miner2, miner4, miner5, miner7}

	for _,v := range correctReadyMiners {
		if !cs.ReadyMiners.Exists(v.ID) {
			t.Errorf("Ready miners map not correct")
		}
	}
	for _,v := range correctBusyMiners {
		if !cs.BusyMiners.Exists(v.ID) {
			t.Errorf("Busy miners map not correct")
		}
	}

	fmt.Println("Ready Miners: ", cs.ReadyMiners.m)
	fmt.Println("Busy Miners: ", cs.BusyMiners.m)
	time.Sleep(time.Second*2)
}

func TestBuyerConnectionScheduler(t *testing.T) {
	ps := msgbus.New(10)
	mainCtx := context.Background()

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
		ID: msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
		DefaultDest: defaultDest.ID,
		IsBuyer: true,
	}

	cs, err := New(&mainCtx, ps, &nodeOperator)
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed:%s", err))
	}
	err = cs.Start()
	if err != nil {
		panic(fmt.Sprintf("schedule manager failed to start:%s", err))
	}

	//
	// Publish multiple miners with varying hashrate
	//
	miner1 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID01"),
		IP: 	"IpAddress1",
		CurrentHashRate:	27,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner2 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID02"),
		IP: 	"IpAddress2",
		CurrentHashRate:	35,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner3 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID03"),
		IP: 	"IpAddress3",
		CurrentHashRate:	72,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner4 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID04"),
		IP: 	"IpAddress4",
		CurrentHashRate:	16,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	miner5 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID05"),
		IP: 	"IpAddress5",
		CurrentHashRate:	88,
		State: msgbus.OnlineState,
		Dest: defaultDest.ID,
		CsMinerHandlerIgnore: false,
	}
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner1.ID),miner1)
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner2.ID),miner2)
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner3.ID),miner3)
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner4.ID),miner4)
	ps.PubWait(msgbus.MinerMsg,msgbus.IDString(miner5.ID),miner5)

	time.Sleep(time.Second*2)

	contract1 := msgbus.Contract {
		IsSeller: false,
		ID: msgbus.ContractID("ContractID01"),
		State: msgbus.ContRunningState,
		Price: 0,
		Limit: 0,
		Speed: 100,
	}
	ps.PubWait(msgbus.ContractMsg,msgbus.IDString(contract1.ID),contract1)

	time.Sleep(time.Second*2)

	contract2 := msgbus.Contract {
		IsSeller: false,
		ID: msgbus.ContractID("ContractID02"),
		State: msgbus.ContRunningState,
		Price: 0,
		Limit: 0,
		Speed: 100,
	}
	ps.PubWait(msgbus.ContractMsg,msgbus.IDString(contract2.ID),contract2)

	time.Sleep(time.Second*2)

	contract3 := msgbus.Contract {
		IsSeller: false,
		ID: msgbus.ContractID("ContractID03"),
		State: msgbus.ContRunningState,
		Price: 0,
		Limit: 0,
		Speed: 100,
	}
	ps.PubWait(msgbus.ContractMsg,msgbus.IDString(contract3.ID),contract3)

	time.Sleep(time.Second*2)

	miner1.CurrentHashRate = 2
	miner2.CurrentHashRate = 2
	miner3.CurrentHashRate = 2
	miner4.CurrentHashRate = 2
	miner5.CurrentHashRate = 2
	ps.SetWait(msgbus.MinerMsg,msgbus.IDString(miner1.ID),miner1)
	ps.SetWait(msgbus.MinerMsg,msgbus.IDString(miner2.ID),miner2)
	ps.SetWait(msgbus.MinerMsg,msgbus.IDString(miner3.ID),miner3)
	ps.SetWait(msgbus.MinerMsg,msgbus.IDString(miner4.ID),miner4)
	ps.SetWait(msgbus.MinerMsg,msgbus.IDString(miner5.ID),miner5)

	time.Sleep(time.Second*5)
}