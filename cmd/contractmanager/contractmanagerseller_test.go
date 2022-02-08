package contractmanager

import (
	//"crypto/ecdsa"
	//"crypto/rand"
	//"errors"
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	//"github.com/ethereum/go-ethereum/crypto/ecies"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func TestSellerRoutine(t *testing.T) {
	ps := msgbus.New(10)
	ts := BeforeEach()
	var hashrateContractAddress [4]common.Address
	var purchasedHashrateContractAddress [4]common.Address
	mainCtx := context.Background()
	contractManagerCtx, contractManagerCancel := context.WithCancel(mainCtx)
	var contractManagerConfig msgbus.ContractManagerConfig
	contractManagerConfigID := msgbus.GetRandomIDString()

	// encrpted cipher text generated from node code using buyer's public key
	//encryptedDest := "04d9b65eada6828aad11f7956e92a5afaa46718e95c2229b21b371c3c6e317bad00018d15f2cedb6400d2156a3cc1c3360b7f747d5ab7e72926937776fc133ae5b9ada0e1d95b57f29b917220a92ed28ff1f57301b6688f7e5ef4ae87015508aefb7156aba0de5cc25d65d1f11a7d3c75330d54d045ebc22231af70fb1aa02b38a6cf93b34a974076db109433ba4191171b2292885"
	//updateEncryptedDest := "049de7772c44fd044bab5600d878a60d14bcd43276888b84e6ea461ed7b7befa06fb2a3eb6c9d8cd065f17fd5744aac7e1ad90d3d1d9da37d42cbc090d813febdef2b6a8d9038d6b5f2023610f64b8837afe3fa1cb7d92977658604848c66d99bfac4ad8596833ae3645a8f05ca6122e246791150f05a3bcf29efd1e33fbb774182acd9c7a7dcfa6b5c1184e2ce8384b4123541abb"

	contractManagerConfigFile, err := LoadTestConfiguration("contractManager", "../../ganacheconfig.json")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	contractManagerConfig.Mnemonic = contractManagerConfigFile["mnemonic"].(string)
	contractManagerConfig.AccountIndex = int(contractManagerConfigFile["accountIndex"].(float64))
	contractManagerConfig.EthNodeAddr = contractManagerConfigFile["ethNodeAddr"].(string)
	contractManagerConfig.ClaimFunds = contractManagerConfigFile["claimFunds"].(bool)
	contractManagerConfig.CloneFactoryAddress = ts.cloneFactoryAddress.Hex()

	contractLength := 15 // 15 s on ganache
	sleepTime := 5000 // 5000 ms sleeptime in ganache
	if contractManagerConfig.EthNodeAddr != "ws://127.0.0.1:7545" {
		contractLength = 30 // 60 s on ropsten
		sleepTime = 30000 // 30000 ms on testnet
	}
	account, privateKey := hdWalletKeys(contractManagerConfig.Mnemonic, contractManagerConfig.AccountIndex + 1)
	buyerAddress := account.Address
	buyerPrivateKey := privateKey

	fmt.Println("Buyer account", buyerAddress)
	fmt.Println("Buyer private key", buyerPrivateKey)

	ps.PubWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractManagerConfig)

	nodeOperator := msgbus.NodeOperator{
		ID: msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
	}
	var cman SellerContractManager
	go newConfigMonitor(&contractManagerCtx, contractManagerCancel, &cman, ps, contractManagerConfigID, &nodeOperator)
	err = cman.init(&contractManagerCtx, ps, contractManagerConfigID, &nodeOperator)
	if err != nil {
		panic(fmt.Sprintf("contract manager init failed:%s", err))
	}

	// subcribe to creation events emitted by clonefactory contract
	cfLogs, cfSub, _ := subscribeToContractEvents(ts.ethClient, ts.cloneFactoryAddress)
	// create event signature to parse out creation and purchase event
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
	CreateHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength), buyerAddress)
	CreateHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength), buyerAddress)
	
	// wait until created hashrate contract was found before continuing 
	loop1:
	for {
		if hashrateContractAddress[0] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop1	
		}
	}
	// PurchaseHashrateContract(cman.ethClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[0], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")
	PurchaseHashrateContract(cman.ethClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[0], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")

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
	
	// contract manager sees existing contracts and states are correct
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 1 was not found or is not in correct state")
	}
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] != msgbus.ContAvailableState {
		t.Errorf("Contract 2 was not found or is not in correct state")
	}

	// contract should be back to available after length has expired
	// if network is ganache, create a new transaction so a new block is created 
	if contractManagerConfig.EthNodeAddr == "ws://127.0.0.1:7545" {
		createNewGanacheBlock(ts, cman.account, cman.privateKey, contractLength, sleepTime)
		if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 1 did not close out correctly")
		}
	} else {
		time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
		time.Sleep(time.Millisecond * time.Duration(sleepTime*2)) // length of transaction
		if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 1 did not close out correctly")
		}
	}

	//
	// test purchase available contract and closeout after length
	//
	fmt.Print("\n\n/// Purchase Available Contract and Closeout After Length ///\n\n\n")
	// contract manager should updated state
	// wait until created hashrate contract was found before continuing 
	loop3:
	for {
		if hashrateContractAddress[1] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop3	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	// PurchaseHashrateContract(cman.ethClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[1], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")
	PurchaseHashrateContract(cman.ethClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[1], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")

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

	// contract should be back to available after length has expired
	// if network is ganache, create a new transaction so a new block is created 
	if contractManagerConfig.EthNodeAddr == "ws://127.0.0.1:7545" {
		createNewGanacheBlock(ts, cman.account, cman.privateKey, contractLength, sleepTime)
		if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 2 did not close out correctly")
		}
	} else {
		time.Sleep(time.Second * time.Duration(contractLength*2)) // length of contract
		time.Sleep(time.Millisecond * time.Duration(sleepTime*2)) // length of transaction
		if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 2 did not close out correctly")
		}
	}
		
	//
	// test early closeout from buyer
	//
	fmt.Print("\n\n/// Early Closeout Frome Buyer ///\n\n\n")
	CreateHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength*10), buyerAddress)

	// wait until created hashrate contract was found before continuing 
	loop5:
	for {
		if hashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop5	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContAvailableState {
		t.Errorf("Contract 3 was not found or is not in correct state")
	}
	// PurchaseHashrateContract(cman.ethClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[2], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")
	PurchaseHashrateContract(cman.ethClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[2], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing 
	loop6:
	for {
		if purchasedHashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop6	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 3 is not in correct state")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Print("Closeout From Buyer: ")
	setContractCloseOut(cman.ethClient, buyerAddress, buyerPrivateKey, hashrateContractAddress[2], &wg, &cman.currentNonce, 0)
	wg.Wait()
	time.Sleep(time.Millisecond * time.Duration(sleepTime*4))
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContAvailableState {
		t.Errorf("Contract 3 did not close out correctly")
	}

	//
	// test contract creation and going through full length with update made to target dest info from buyer while node is running
	//
	fmt.Print("\n\n/// Update Made To Target Dest By Buyer While Contract Is Running ///\n\n\n")
	CreateHashrateContract(cman.ethClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength), buyerAddress)
	
	// wait until created hashrate contract was found before continuing 
	loop7:
	for {
		if hashrateContractAddress[3] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop7	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[3].Hex())] != msgbus.ContAvailableState {
		t.Errorf("Contract 4 was not found or is not in correct state")
	}

	// PurchaseHashrateContract(cman.ethClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[3], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")
	PurchaseHashrateContract(cman.ethClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[3], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing 
	loop8:
	for {
		if purchasedHashrateContractAddress[3] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop8	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[3].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 4 is not in correct state")
	}

	UpdateCipherText(cman.ethClient, buyerAddress, buyerPrivateKey, hashrateContractAddress[3], "stratum+tcp://127.0.0.1:3333/updated")
	time.Sleep(time.Millisecond * time.Duration(sleepTime*2))
	// check dest msg with associated contract was updated in msgbus
	event, err := cman.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(hashrateContractAddress[3].Hex()))
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
		t.Errorf("Contract 4's target dest was not updated")
	}

	// if network is ganache, create a new transaction so a new block is created 
	if contractManagerConfig.EthNodeAddr == "ws://127.0.0.1:7545" {
		createNewGanacheBlock(ts, cman.account, cman.privateKey, contractLength, sleepTime)
		if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[3].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 4 did not close out correctly")
		}
	} else {
		time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
		time.Sleep(time.Millisecond * time.Duration(sleepTime)) // length of transaction
		if cman.nodeOperator.Contracts[msgbus.ContractID(hashrateContractAddress[3].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 4 did not close out correctly")
		}
	}

	//
	// test contract manager config updated
	//
	contractManagerConfig.ClaimFunds = false 
	ps.SetWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractManagerConfig)
	time.Sleep(time.Second * 3)
	if cman.claimFunds != false {
		t.Errorf("Contract manager's configuration was not updated after msgbus update")
	}

	//
	// test seller updated purchase info
	//
	// UpdatePurchaseInformation(cman.ethClient, cman.account, cman.privateKey, hashrateContractAddress[3], int(10), int(10), int(50), int(contractLength*10))
	// time.Sleep(time.Millisecond * time.Duration(sleepTime*2))
	// // check purchase information in contract was updated in msgbus
	// event, err = cman.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(hashrateContractAddress[3].Hex()))
	// if err != nil {
	// 	panic(fmt.Sprintf("Getting Contract Failed: %s", err))
	// }
	// if event.Err != nil {
	// 	panic(fmt.Sprintf("Getting Contract Failed: %s", event.Err))
	// }
	// contractMsg = event.Data.(msgbus.Contract)
	// if contractMsg.Limit != 10 || contractMsg.Speed != 50 {
	// 	t.Errorf("Contract 4's purchase info was not updated")
	// } 
}