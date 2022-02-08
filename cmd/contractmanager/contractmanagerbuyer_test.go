package contractmanager

import (
	//"crypto/ecdsa"
	//"crypto/rand"
	//"errors"
	"fmt"
	"log"
	"testing"
	"time"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	//"github.com/ethereum/go-ethereum/crypto/ecies"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func TestBuyerRoutine(t *testing.T) {
	ps := msgbus.New(10)
	ts := BeforeEach()
	var hashrateContractAddress [3]common.Address
	var purchasedHashrateContractAddress [3]common.Address
	mainCtx := context.Background()
	contractManagerCtx, contractManagerCancel := context.WithCancel(mainCtx)
	var contractManagerConfig msgbus.ContractManagerConfig
	contractManagerConfigID := msgbus.GetRandomIDString()

	contractLength := 10000 

	contractManagerConfigFile, err := LoadTestConfiguration("contractManager", "../../ganacheconfig.json")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	contractManagerConfig.Mnemonic = contractManagerConfigFile["mnemonic"].(string)
	contractManagerConfig.AccountIndex = int(contractManagerConfigFile["accountIndex"].(float64))
	contractManagerConfig.EthNodeAddr = contractManagerConfigFile["ethNodeAddr"].(string)
	contractManagerConfig.CloneFactoryAddress = ts.cloneFactoryAddress.Hex()

	sleepTime := 5000 // 5000 ms sleeptime in ganache
	if contractManagerConfig.EthNodeAddr != "ws://127.0.0.1:7545" {
		sleepTime = 20000 // 20000 ms on testnet
	}

	account, privateKey := hdWalletKeys(contractManagerConfig.Mnemonic, contractManagerConfig.AccountIndex + 1)
	sellerAddress := account.Address
	sellerPrivateKey := privateKey
	fmt.Println("Seller account", sellerAddress)
	fmt.Println("Seller private key", sellerPrivateKey)

	ps.PubWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractManagerConfig)

	nodeOperator := msgbus.NodeOperator{
		ID: msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
	}
	var cman BuyerContractManager
	go newConfigMonitor(&contractManagerCtx, contractManagerCancel, &cman, ps, contractManagerConfigID, &nodeOperator)
	err = cman.init(&contractManagerCtx, ps, contractManagerConfigID, &nodeOperator)
	if err != nil {
		panic(fmt.Sprintf("contract manager init failed:%s", err))
	}

	miner1 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID01"),
		IP: 	"IpAddress1",
		CurrentHashRate:	30,
		State: msgbus.OnlineState,
	}
	miner2 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID02"),
		IP: 	"IpAddress2",
		CurrentHashRate:	20,
		State: msgbus.OnlineState,
	}
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner1.ID),miner1)
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner2.ID),miner2)

	// subcribe to creation events emitted by clonefactory contract 
	cfLogs, cfSub, _ := subscribeToContractEvents(ts.ethClient, ts.cloneFactoryAddress)
	// create event signature to parse out creation event
	contractCreatedSig := []byte("contractCreated(address)")
	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
	clonefactoryContractPurchasedSig := []byte("clonefactoryContractPurchased(address)")
	clonefactoryContractPurchasedSigHash := crypto.Keccak256Hash(clonefactoryContractPurchasedSig)
	go func() {
		i := 0
		j := 0
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
			case cfLog := <-cfLogs:
				switch {
				case cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex():
					hashrateContractAddress[i] = common.HexToAddress(cfLog.Topics[1].Hex())
					fmt.Printf("Address of created Hashrate Contract %d: %s\n\n", i + 1, hashrateContractAddress[i].Hex())
					i++
				
				case cfLog.Topics[0].Hex() == clonefactoryContractPurchasedSigHash.Hex():
					purchasedHashrateContractAddress[j] = common.HexToAddress(cfLog.Topics[1].Hex())
					fmt.Printf("Address of purchased Hashrate Contract %d: %s\n\n", j + 1, purchasedHashrateContractAddress[j].Hex())
					j++
				}
			}
		}
	}()

	//
	// test startup with 1 running contract and 1 availabe contract
	//
	CreateHashrateContract(cman.ethClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(0), int(1000), int(contractLength), cman.account)
	CreateHashrateContract(cman.ethClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(0), int(1000), int(contractLength), cman.account)

	// wait until created hashrate contract was found before continuing 
	loop1:
	for {
		if hashrateContractAddress[0] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop1	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[0], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing
	loop2:
	for {
		if purchasedHashrateContractAddress[0] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop2	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	
	err = cman.start()
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}
	
	// check contract manager found miners in msgbus 
	if cman.miners[miner1.ID] != miner1 {
		t.Errorf("MinerID01 not found by contract manager")
	}
	if cman.miners[miner2.ID] != miner2 {
		t.Errorf("MinerID02 not found by contract manager")
	}

	// contract manager sees existing contracts and states are correct
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 1 was not found or is not in correct state")
	}
	if _,ok := cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] ; ok {
		t.Errorf("Contract 2 was found by buyer node while in the available state")
	}

	// contract should be removed from buyer node when set back to available
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
	// 	time.Sleep(time.Millisecond * time.Duration(sleepTime)) // length of transaction
	// 	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())]; ok {
	// 		t.Errorf("Contract 1 did not close out correctly")
	// 	}
	// }()

	// contract manager should updated states
	// wait until created hashrate contract was found before continuing 
	loop3:
	for {
		if hashrateContractAddress[1] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop3	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[1], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing
	loop4:
	for {
		if purchasedHashrateContractAddress[1] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop4	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 2 is not in correct state")
	}

	// contract should be removed from buyer node when set back to available
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
	// 	time.Sleep(time.Millisecond * time.Duration(sleepTime)) // length of transaction
	// 	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())]; ok {
	// 		t.Errorf("Contract 2 did not close out correctly")
	// 	}
	// }()

	/*
	//
	// Test early closeout from seller
	//
	CreateHashrateContract(cman.ethClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength*10), cman.account)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] ; ok {
		t.Errorf("Contract 3 was found by buyer node while in the available state")
	}

	PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[2], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 3 is not in correct state")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	setContractCloseOut(cman.ethClient, sellerAddress, sellerPrivateKey, hashrateContractAddress[2], &wg, &cman.currentNonce, 0)
	wg.Wait()
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())]; ok {
		t.Errorf("Contract 3 did not close out correctly")
	}
	*/

	//
	// Test contract creation, purchasing, and target dest being updated while node is running
	//
	CreateHashrateContract(cman.ethClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(0), int(1000), int(contractLength), cman.account)

	loop5:
	for {
		if hashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop5	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if _,ok := cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] ; ok {
		t.Errorf("Contract 4 was found by buyer node while in the available state")
	}
	PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[2], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing
	loop6:
	for {
		if purchasedHashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop6	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 4 is not in correct state")
	}

	UpdateCipherText(cman.ethClient, cman.account, cman.privateKey, hashrateContractAddress[2], "stratum+tcp://127.0.0.1:3333/updated")
	time.Sleep(time.Millisecond * time.Duration(sleepTime*2))
	// check dest msg with associated contract was updated in msgbus
	event, err := cman.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(hashrateContractAddress[2].Hex()))
	if err != nil {
		panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
	}
	contractMsg := event.Data.(msgbus.Contract)
	event, err = cman.ps.GetWait(msgbus.DestMsg, msgbus.IDString(contractMsg.Dest))
	if err != nil {
		panic(fmt.Sprintf("Getting Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Getting Dest Failed: %s", event.Err))
	}
	destMsg := event.Data.(msgbus.Dest)
	if destMsg.NetUrl != "stratum+tcp://127.0.0.1:3333/updated" {
		t.Errorf("Contract 3's target dest was not updated")
	}

	//
	// Test miner hashrate being updated
	//
	// miner 1's hashrate is updated
	miner1.CurrentHashRate = 20
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	
	//check contract manager got update to miner
	if cman.miners[miner1.ID].CurrentHashRate != miner1.CurrentHashRate {
		t.Errorf("Contract Manager did not receive update to MinerID01 hashrate")
	}

	// new miner published
	miner3 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID03"),
		IP: 	"IpAddress3",
		CurrentHashRate:	20,
		State: msgbus.OnlineState,
	}
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner3.ID),miner3)
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))

	//check contract manager saw new miner and updated total hashrate
	if cman.miners[miner3.ID].CurrentHashRate != miner3.CurrentHashRate {
		t.Errorf("Contract Manager did not see new Miner 3")
	}

	// miner 2 deleted
	ps.UnpubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID))
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if _,ok := cman.miners[miner2.ID]; ok {
		t.Errorf("Contract Manager did not remove miner 2 after it was deleted from msgbus")
	}

	//
	// Test miners are set to offline state so running contracts should close out
	//
	miner1.State = msgbus.OfflineState
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	miner3.State = msgbus.OfflineState
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))

	// check contracts map is empty now
	if len(cman.nodeOperator.Contracts) != 0 {
		t.Errorf("Contracts did not closeout after all miners were set to offline")
	}

	//
	// test contract manager config updated
	//
	contractManagerConfig.AccountIndex = 1 
	ps.SetWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractManagerConfig)
	time.Sleep(time.Second * 3)
	newAccount,_ := hdWalletKeys(contractManagerConfig.Mnemonic, 1)
	if cman.account != newAccount.Address {
		t.Errorf("Contract manager's configuration was not updated after msgbus update")
	}
}