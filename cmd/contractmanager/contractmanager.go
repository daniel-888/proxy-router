package contractmanager

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/ledger"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/webfacing"
)

const (
	AvailableState uint8 = 0
	ActiveState    uint8 = 1
	RunningState   uint8 = 2
	CompleteState  uint8 = 3
)

type hashrateContractValues struct {
	State                  uint8
	Price                  int
	Limit                  int
	Speed                  int
	Length                 int
	ValidationFee          int
	StartingBlockTimestamp int
	Buyer                  common.Address
	Seller                 common.Address
}

type nonce struct {
	mutex	sync.Mutex
	nonce	uint64
}

type ContractManager interface {
	start() error
	init(ps *msgbus.PubSub, cmConfig map[string]interface{}) (err error)
	setupExistingContracts()
	readContracts() []common.Address
	watchHashrateContract(addr msgbus.ContractID, hrLogs chan types.Log, hrSub ethereum.Subscription)
	closeOutMonitor(contractMsg msgbus.Contract)
}

type BuyerContractManager struct {
	ps                  	*msgbus.PubSub
	rpcClient           	*ethclient.Client
	webFacingAddress    	common.Address
	ledgerAddress       	common.Address
	account             	common.Address
	privateKey          	string
	currentNonce			nonce
	msg						msgbus.Buyer
	api						externalapi.APIRepos
	miners					map[msgbus.MinerID]msgbus.Miner
	hashrateUpdateChans		map[msgbus.ContractID]chan bool
}	

type SellerContractManager struct {
	ps                  	*msgbus.PubSub
	rpcClient           	*ethclient.Client
	cloneFactoryAddress 	common.Address
	ledgerAddress       	common.Address
	account             	common.Address
	privateKey          	string
	currentNonce			nonce
	msg						msgbus.Seller
	api						externalapi.APIRepos
}

func Run(contractManager ContractManager, ps *msgbus.PubSub, cmConfig map[string]interface{}) (err error) {
	err = contractManager.init(ps, cmConfig)
	if err != nil {
		return err
	}
	err = contractManager.start()
	if err != nil {
		return err
	}

	return nil
}

func (seller *SellerContractManager) init(ps *msgbus.PubSub, cmConfig map[string]interface{}) (err error) {
	var client *ethclient.Client
	client, err = setUpClient(cmConfig["rpcClientAddress"].(string), common.HexToAddress(cmConfig["sellerEthereumAddress"].(string)))
	if err != nil {
		log.Fatal(err)
	}
	seller.ps = ps
	seller.rpcClient = client
	seller.cloneFactoryAddress = common.HexToAddress(cmConfig["cloneFactoryAddress"].(string))
	seller.ledgerAddress = common.HexToAddress(cmConfig["ledgerAddress"].(string))
	seller.account = common.HexToAddress(cmConfig["sellerEthereumAddress"].(string))
	seller.privateKey = cmConfig["sellerEthereumPrivateKey"].(string)
	seller.api.InitializeJSONRepos()

	availableContractsMap := make(map[msgbus.ContractID]bool)
	activeContractsMap := make(map[msgbus.ContractID]bool)
	runningContractsMap := make(map[msgbus.ContractID]bool)
	completeContractsMap := make(map[msgbus.ContractID]bool)

	sellerMsg := msgbus.Seller {
		ID: msgbus.SellerID(seller.account.Hex()),
		AvailableContracts:	availableContractsMap,
		ActiveContracts: 	activeContractsMap,
		RunningContracts: 	runningContractsMap,
		CompleteContracts: 	completeContractsMap,
	}

	seller.msg = sellerMsg

	return err
}

func (seller *SellerContractManager) start() error {
	// go seller.api.RunAPI()

	seller.setupExistingContracts()

	// routine for listensing to contract creation events that will update seller msg with new contracts and load new contract onto msgbus
	cfLogs, cfSub := subscribeToContractEvents(seller.rpcClient, seller.cloneFactoryAddress)
	go seller.watchContractCreation(cfLogs, cfSub)

	// routine starts routines for seller's contracts that monitors contract purchase, close, and cancel events
	go func() {
		// start routines for existing contracts
		for addr := range seller.msg.AvailableContracts {
			hrLogs, hrSub := subscribeToContractEvents(seller.rpcClient, common.HexToAddress(string(addr)))
			go seller.watchHashrateContract(addr, hrLogs, hrSub)
		}

		// monitor new contracts getting created and start hashrate conrtract monitor routine when they are created
		contractEventChan := seller.ps.NewEventChan()
		err := seller.ps.Sub(msgbus.ContractMsg, "", contractEventChan)
		if err != nil {
			log.Fatal(err)
		}
		for {	
			event := <-contractEventChan
			if event.EventType == msgbus.PublishEvent {
				newContract := event.Data.(msgbus.Contract)
				if newContract.State == msgbus.ContAvailableState {
					addr := common.HexToAddress(string(newContract.ID))
					hrLogs, hrSub := subscribeToContractEvents(seller.rpcClient, addr)
					go seller.watchHashrateContract(msgbus.ContractID(addr.Hex()), hrLogs, hrSub)
				}
			}
		}
	}()
	return nil
}

func (seller *SellerContractManager) setupExistingContracts() {
	var contractValues []hashrateContractValues
	var contractMsgs []msgbus.Contract

	sellerContracts := seller.readContracts()
	fmt.Println("Existing Seller Contracts: ", sellerContracts)
	for i := range sellerContracts {
		contractValues = append(contractValues, readHashrateContract(seller.rpcClient, sellerContracts[i]))
		contractMsgs = append(contractMsgs, createContractMsg(sellerContracts[i], contractValues[i], true))
		seller.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])
		seller.api.Contract.AddContractFromMsgBus(contractMsgs[i])

		seller.msg.AvailableContracts[msgbus.ContractID(sellerContracts[i].Hex())] = false
		seller.msg.ActiveContracts[msgbus.ContractID(sellerContracts[i].Hex())] = false
		seller.msg.RunningContracts[msgbus.ContractID(sellerContracts[i].Hex())] = false
		seller.msg.CompleteContracts[msgbus.ContractID(sellerContracts[i].Hex())] = false

		switch contractValues[i].State {
		case AvailableState:
			seller.msg.AvailableContracts[msgbus.ContractID(sellerContracts[i].Hex())] = true
		case ActiveState:
			seller.msg.ActiveContracts[msgbus.ContractID(sellerContracts[i].Hex())] = true
		case RunningState:
			seller.msg.RunningContracts[msgbus.ContractID(sellerContracts[i].Hex())] = true
		case CompleteState:
			seller.msg.CompleteContracts[msgbus.ContractID(sellerContracts[i].Hex())] = true
		}
	}

	seller.ps.PubWait(msgbus.SellerMsg, msgbus.IDString(seller.msg.ID), seller.msg)
	seller.api.Seller.AddSellerFromMsgBus(seller.msg)
}

func (seller *SellerContractManager) readContracts() []common.Address {
	var sellerContractAddresses []common.Address
	var hashrateContractInstance *implementation.Implementation
	var hashrateContractSeller common.Address

	instance, err := ledger.NewLedger(seller.ledgerAddress, seller.rpcClient)
	if err != nil {
		log.Fatal(err)
	}

	hashrateContractAddresses, err := instance.GetListOfContractsLedger(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// parse existing hashrate contracts for ones that belong to seller
	for i := range hashrateContractAddresses {
		hashrateContractInstance, err = implementation.NewImplementation(hashrateContractAddresses[i], seller.rpcClient)
		if err != nil {
			log.Fatal(err)
		}
		hashrateContractSeller, err = hashrateContractInstance.Seller(nil)
		if err != nil {
			log.Fatal(err)
		}
		if hashrateContractSeller == seller.account {
			sellerContractAddresses = append(sellerContractAddresses, hashrateContractAddresses[i])
		}
	}

	return sellerContractAddresses
}

func (seller *SellerContractManager) watchContractCreation(cfLogs chan types.Log, cfSub ethereum.Subscription) {
	defer close(cfLogs)
	defer cfSub.Unsubscribe()
	for {
		select {
		case err := <-cfSub.Err():
			log.Fatal(err)
		case cfLog := <-cfLogs:
			address := common.HexToAddress(cfLog.Topics[1].Hex())
			// check if contract created belongs to seller
			hashrateContractInstance, err := implementation.NewImplementation(address, seller.rpcClient)
			if err != nil {
				log.Fatal(err)
			}
			hashrateContractSeller, err := hashrateContractInstance.Seller(nil)
			if err != nil {
				log.Fatal(err)
			}
			if hashrateContractSeller == seller.account {
				fmt.Printf("Address of created Hashrate Contract: %s\n\n", address.Hex())

				createdContractValues := readHashrateContract(seller.rpcClient, address)
				createdContractMsg := createContractMsg(address, createdContractValues, true)
				seller.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(address.Hex()), createdContractMsg)
				seller.api.Contract.AddContractFromMsgBus(createdContractMsg)

				seller.msg.AvailableContracts[msgbus.ContractID(address.Hex())] = true
				seller.msg.ActiveContracts[msgbus.ContractID(address.Hex())] = false
				seller.msg.RunningContracts[msgbus.ContractID(address.Hex())] = false
				seller.msg.CompleteContracts[msgbus.ContractID(address.Hex())] = false
				seller.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(seller.msg.ID), seller.msg)
				seller.api.Seller.UpdateSeller(string(seller.msg.ID), msgdata.ConvertSellerMSGtoSellerJSON(seller.msg))
			}
		}
	}
}

func (seller *SellerContractManager) watchHashrateContract(addr msgbus.ContractID, hrLogs chan types.Log, hrSub ethereum.Subscription) {
	contractEventChan := seller.ps.NewEventChan()

	// check if contract is already in the running state and needs to be monitored for closeout
	event, err := seller.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
	if err != nil {
		panic(fmt.Sprintf("Getting Hashrate Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Getting Hashrate Contract Failed: %s", event.Err))
	}
	hashrateContractMsg := event.Data.(msgbus.Contract)
	if hashrateContractMsg.State == msgbus.ContRunningState {
		go seller.closeOutMonitor(hashrateContractMsg)
	}

	// create event signatures to parse out which event was being emitted from hashrate contract
	contractPurchasedSig := []byte("contractPurchased(address)")
	contractClosedSig := []byte("contractClosed(address)")
	contractFundedSig := []byte("contractFunded(address)")
	contractPurchasedSigHash := crypto.Keccak256Hash(contractPurchasedSig)
	contractClosedSigHash := crypto.Keccak256Hash(contractClosedSig)
	contractFundedSigHash := crypto.Keccak256Hash(contractFundedSig)

	// to decode event data
	implementationAbi, err := abi.JSON(strings.NewReader(string(implementation.ImplementationABI)))
	if err != nil {
		log.Fatal(err)
	}
	purchasedEvent := struct {
		Buyer common.Address
	}{}

	// routine monitoring and acting upon events emmited by hashrate contract
	go func() {
		defer close(hrLogs)
		defer hrSub.Unsubscribe()
		for {
			select {
			case err := <-hrSub.Err():
				log.Fatal(err)
			case hLog := <-hrLogs:
				switch hLog.Topics[0].Hex() {
				case contractPurchasedSigHash.Hex():
					fmt.Printf("Address of purchased Hashrate Contract : %s\n\n", addr)
					err := implementationAbi.UnpackIntoInterface(&purchasedEvent, "contractPurchased", hLog.Data)
					if err != nil {
						log.Fatal(err)
					}

					destUrl := readDestUrl(seller.rpcClient, common.HexToAddress(string(addr)), seller.privateKey)
					destMsg := msgbus.Dest{
						ID:     msgbus.DestID(msgbus.GetRandomIDString()),
						NetUrl: msgbus.DestNetUrl(destUrl),
					}
					seller.ps.PubWait(msgbus.DestMsg, msgbus.IDString(destMsg.ID), destMsg)
					seller.api.Dest.AddDestFromMsgBus(destMsg)
				
					event, err := seller.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
					if err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
					}
					if event.Err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
					}
					contractMsg := event.Data.(msgbus.Contract)
					contractMsg.Dest = destMsg.ID
					contractMsg.State = msgbus.ContActiveState
					contractMsg.Buyer = msgbus.BuyerID(purchasedEvent.Buyer.Hex())
					seller.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(addr), contractMsg)
					seller.api.Contract.UpdateContract(string(addr), msgdata.ConvertContractMSGtoContractJSON(contractMsg))

					seller.msg.AvailableContracts[addr] = false
					seller.msg.ActiveContracts[addr] = true
					seller.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(seller.msg.ID), seller.msg)
					seller.api.Seller.UpdateSeller(string(seller.msg.ID), msgdata.ConvertSellerMSGtoSellerJSON(seller.msg))

				case contractFundedSigHash.Hex():
					fmt.Printf("Address of funded Hashrate Contract : %s\n\n", addr)

					contractValues := readHashrateContract(seller.rpcClient, common.HexToAddress(string(addr)))
					contractMsg := createContractMsg(common.HexToAddress(string(addr)), contractValues, true)
					seller.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(addr), contractMsg)
					seller.api.Contract.UpdateContract(string(addr), msgdata.ConvertContractMSGtoContractJSON(contractMsg))

					seller.msg.ActiveContracts[addr] = false
					seller.msg.RunningContracts[addr] = true
					seller.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(seller.msg.ID), seller.msg)

				case contractClosedSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Closed \n\n", addr)

					closedContractValues := readHashrateContract(seller.rpcClient, common.HexToAddress(string(addr)))
					closedContractMsg := createContractMsg(common.HexToAddress(string(addr)), closedContractValues, true)
					seller.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(closedContractMsg.ID), closedContractMsg)
					
					seller.msg.RunningContracts[addr] = false
					seller.msg.CompleteContracts[addr] = true
					seller.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(seller.msg.ID), seller.msg)
					seller.api.Seller.UpdateSeller(string(seller.msg.ID), msgdata.ConvertSellerMSGtoSellerJSON(seller.msg))
				}
			}
		}
	}()

	err = seller.ps.Sub(msgbus.ContractMsg, msgbus.IDString(addr), contractEventChan)
	if err != nil {
		log.Fatal(err)
	}
	// once contract is running, closeout after length of contract has passed if it was not closed out early
	for {
		event := <-contractEventChan
		if event.EventType == msgbus.UpdateEvent {
			runningContractMsg := event.Data.(msgbus.Contract)
			if runningContractMsg.State == msgbus.ContRunningState {
				// run routine for each running contract to check if contract length has passed and contract should be closed out
				go seller.closeOutMonitor(runningContractMsg)
			}
		}		
	}
}

func (seller *SellerContractManager) closeOutMonitor(contractMsg msgbus.Contract) {
	contractFinishedTimestamp := contractMsg.StartingBlockTimestamp + contractMsg.Length

	// subscribe to latest block headers
	headers := make(chan *types.Header)
	sub, err := seller.rpcClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	defer close(headers)
	defer sub.Unsubscribe()

	loop:
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			// get latest block from header
			block, err := seller.rpcClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			// check if contract length has passed
			if block.Time() >= uint64(contractFinishedTimestamp) {
				// if contract was not already closed early, close out here
				contractValues := readHashrateContract(seller.rpcClient, common.HexToAddress(string(contractMsg.ID)))
				if contractValues.State == RunningState {
					var wg sync.WaitGroup
					wg.Add(1)
					setContractCloseOut(seller.rpcClient, seller.account, seller.privateKey, common.HexToAddress(string(contractMsg.ID)), &wg, &seller.currentNonce)
					wg.Wait()
				}
				break loop
			}
		}
	}
}

func (buyer *BuyerContractManager) init(ps *msgbus.PubSub, cmConfig map[string]interface{}) (err error) {
	var client *ethclient.Client
	client, err = setUpClient(cmConfig["rpcClientAddress"].(string), common.HexToAddress(cmConfig["buyerEthereumAddress"].(string)))
	if err != nil {
		log.Fatal(err)
	}
	buyer.ps = ps
	buyer.rpcClient = client
	buyer.webFacingAddress = common.HexToAddress(cmConfig["webFacingAddress"].(string))
	buyer.ledgerAddress = common.HexToAddress(cmConfig["ledgerAddress"].(string))
	buyer.account = common.HexToAddress(cmConfig["buyerEthereumAddress"].(string))
	buyer.privateKey = cmConfig["buyerEthereumPrivateKey"].(string)
	buyer.api.InitializeJSONRepos()

	activeContractsMap := make(map[msgbus.ContractID]bool)
	runningContractsMap := make(map[msgbus.ContractID]bool)
	completeContractsMap := make(map[msgbus.ContractID]bool)

	buyerMsg := msgbus.Buyer {
		ID: msgbus.BuyerID(buyer.account.Hex()),
		ActiveContracts: 	activeContractsMap,
		RunningContracts: 	runningContractsMap,
		CompleteContracts: 	completeContractsMap,
	}

	buyer.msg = buyerMsg

	buyer.miners = make(map[msgbus.MinerID]msgbus.Miner)

	buyer.hashrateUpdateChans = make(map[msgbus.ContractID]chan bool)

	return err
}

func (buyer *BuyerContractManager) start() error {
	// go buyer.api.RunAPI()

	buyer.setupExistingContracts()

	// update buyer node with current miners
	miners, err := buyer.ps.MinerGetAllWait()
	if err != nil {
		log.Fatal(err)
	}
	for i := range miners {
		miner, err := buyer.ps.MinerGetWait(miners[i])
		if err != nil {
			log.Fatal(err)
		}
		buyer.miners[msgbus.MinerID(miners[i])] = *miner
	}

	// check hashrate everytime miner msgs are published, updated, deleted
	minerEventChan := buyer.ps.NewEventChan()
	event, err := buyer.ps.SubWait(msgbus.MinerMsg, msgbus.IDString(""), minerEventChan)
	if err != nil {
		log.Fatal(err)
	}
	if event.EventType != msgbus.SubscribedEvent {
		panic(fmt.Sprintf(" Wrong event type %v\n", event))
	}
	go buyer.minerMonitor(minerEventChan)

	// subcribe to events emitted by webfacing contract to read contract purchase event
	wfLogs, wfSub := subscribeToContractEvents(buyer.rpcClient, buyer.webFacingAddress)

	// routine for listensing to contract purchase events to update buyer with new contracts they purchased
	go buyer.watchContractPurchase(wfLogs, wfSub)

	// routine starts routines for buyers's contracts that monitors contract running and close events
	go func() {
		// start watch hashrate contract for existing running contracts
		for addr := range buyer.msg.ActiveContracts {
			hrLogs, hrSub := subscribeToContractEvents(buyer.rpcClient, common.HexToAddress(string(addr)))
			go buyer.watchHashrateContract(addr, hrLogs, hrSub)
		}

		// monitor new contracts getting purchased and start watch hashrate conrtract routine when they are purchased
		contractEventChan := buyer.ps.NewEventChan()
		err := buyer.ps.Sub(msgbus.ContractMsg, "", contractEventChan)
		if err != nil {
			log.Fatal(err)
		}
		for {	
			event := <-contractEventChan
			if event.EventType == msgbus.PublishEvent {
				newContract := event.Data.(msgbus.Contract)
				fmt.Println("New Contract: ", newContract)
				if newContract.State == msgbus.ContAvailableState {
					addr := common.HexToAddress(string(newContract.ID))
					hrLogs, hrSub := subscribeToContractEvents(buyer.rpcClient, addr)
					go buyer.watchHashrateContract(msgbus.ContractID(addr.Hex()), hrLogs, hrSub)
				}
			}
		}
	}()
	return nil
}

func (buyer *BuyerContractManager) setupExistingContracts() {
	var contractValues []hashrateContractValues
	var contractMsgs []msgbus.Contract

	buyerContracts := buyer.readContracts()
	fmt.Println("Existing Buyer Contracts: ", buyerContracts)
	for i := range buyerContracts {
		contractValues = append(contractValues, readHashrateContract(buyer.rpcClient, buyerContracts[i]))
		contractMsgs = append(contractMsgs, createContractMsg(buyerContracts[i], contractValues[i], false))
		buyer.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])
		buyer.api.Contract.AddContractFromMsgBus(contractMsgs[i])

		buyer.msg.ActiveContracts[msgbus.ContractID(buyerContracts[i].Hex())] = false
		buyer.msg.RunningContracts[msgbus.ContractID(buyerContracts[i].Hex())] = false
		buyer.msg.CompleteContracts[msgbus.ContractID(buyerContracts[i].Hex())] = false

		switch contractValues[i].State {
		case ActiveState:
			buyer.msg.ActiveContracts[msgbus.ContractID(buyerContracts[i].Hex())] = true
		case RunningState:
			buyer.msg.RunningContracts[msgbus.ContractID(buyerContracts[i].Hex())] = true
			buyer.hashrateUpdateChans[msgbus.ContractID(buyerContracts[i].Hex())] = make(chan bool)
		case CompleteState:
			buyer.msg.CompleteContracts[msgbus.ContractID(buyerContracts[i].Hex())] = true
		}
	}

	buyer.ps.PubWait(msgbus.BuyerMsg, msgbus.IDString(buyer.msg.ID), buyer.msg)
	buyer.api.Buyer.AddBuyerFromMsgBus(buyer.msg)
}

func (buyer *BuyerContractManager) readContracts() []common.Address {
	var buyerContractAddresses []common.Address
	var hashrateContractInstance *implementation.Implementation
	var hashrateContractBuyer common.Address

	instance, err := ledger.NewLedger(buyer.ledgerAddress, buyer.rpcClient)
	if err != nil {
		log.Fatal(err)
	}

	hashrateContractAddresses, err := instance.GetListOfContractsLedger(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// parse existing hashrate contracts for ones that belong to buyer
	for i := range hashrateContractAddresses {
		hashrateContractInstance, err = implementation.NewImplementation(hashrateContractAddresses[i], buyer.rpcClient)
		if err != nil {
			log.Fatal(err)
		}
		hashrateContractBuyer, err = hashrateContractInstance.Buyer(nil)
		if err != nil {
			log.Fatal(err)
		}
		if hashrateContractBuyer == buyer.account {
			buyerContractAddresses = append(buyerContractAddresses, hashrateContractAddresses[i])
		}
	}

	return buyerContractAddresses
}

func (buyer *BuyerContractManager) watchContractPurchase(wfLogs chan types.Log, wfSub ethereum.Subscription) {
	// to decode event data
	webFacingAbi, err := abi.JSON(strings.NewReader(string(webfacing.WebfacingABI)))
	if err != nil {
		log.Fatal(err)
	}
	purchasedEvent := struct {
		Contract common.Address
	}{}
	
	defer close(wfLogs)
	defer wfSub.Unsubscribe()
	
	for {
		select {
		case err := <-wfSub.Err():
			log.Fatal(err)
		case wfLog := <-wfLogs:
			err := webFacingAbi.UnpackIntoInterface(&purchasedEvent, "contractPurchase", wfLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			contractAddress := purchasedEvent.Contract
			contractValues := readHashrateContract(buyer.rpcClient, contractAddress)
			fmt.Println("Contract Values: ", contractValues)
			if contractValues.Buyer == buyer.account {
				fmt.Printf("Address of purchased Hashrate Contract : %s\n\n", contractAddress.Hex())
				
				destUrl := readDestUrl(buyer.rpcClient, common.HexToAddress(string(contractAddress.Hex())), buyer.privateKey)
				destMsg := msgbus.Dest{
					ID:     msgbus.DestID(msgbus.GetRandomIDString()),
					NetUrl: msgbus.DestNetUrl(destUrl),
				}
				buyer.ps.PubWait(msgbus.DestMsg, msgbus.IDString(destMsg.ID), destMsg)
				buyer.api.Dest.AddDestFromMsgBus(destMsg)

				contractMsg := createContractMsg(contractAddress, contractValues, false)
				contractMsg.Dest = destMsg.ID
				buyer.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsg.ID), contractMsg)
				buyer.api.Contract.AddContractFromMsgBus(contractMsg)

				buyer.msg.ActiveContracts[msgbus.ContractID(contractAddress.Hex())] = true
				buyer.msg.RunningContracts[msgbus.ContractID(contractAddress.Hex())] = false
				buyer.msg.CompleteContracts[msgbus.ContractID(contractAddress.Hex())] = false
				buyer.ps.SetWait(msgbus.BuyerMsg, msgbus.IDString(buyer.msg.ID), buyer.msg)
				buyer.api.Buyer.UpdateBuyer(string(buyer.msg.ID), msgdata.ConvertBuyerMSGtoBuyerJSON(buyer.msg))
			}
		}
	}
}

func (buyer *BuyerContractManager) watchHashrateContract(addr msgbus.ContractID, hrLogs chan types.Log, hrSub ethereum.Subscription) {
	contractEventChan := buyer.ps.NewEventChan()

	// check if contract is already in the running state and needs to be monitored for closeout
	event, err := buyer.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
	if err != nil {
		panic(fmt.Sprintf("Getting Hashrate Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Getting Hashrate Contract Failed: %s", event.Err))
	}
	hashrateContractMsg := event.Data.(msgbus.Contract)
	if hashrateContractMsg.State == msgbus.ContRunningState {
		fmt.Println("starting closeout monitor for: ", hashrateContractMsg.ID)
		go buyer.closeOutMonitor(hashrateContractMsg)
	}

	// create event signatures to parse out which event was being emitted from hashrate contract
	contractClosedSig := []byte("contractClosed(address)")
	contractFundedSig := []byte("contractFunded(address)")
	contractClosedSigHash := crypto.Keccak256Hash(contractClosedSig)
	contractFundedSigHash := crypto.Keccak256Hash(contractFundedSig)

	// routine monitoring and acting upon events emmited by hashrate contract
	go func() {
		defer close(hrLogs)
		defer hrSub.Unsubscribe()
		for {
			select {
			case err := <-hrSub.Err():
				log.Fatal(err)
			case hLog := <-hrLogs:
				switch hLog.Topics[0].Hex() {
				case contractFundedSigHash.Hex():
					fmt.Printf("Address of funded Hashrate Contract : %s\n\n", addr)

					contractValues := readHashrateContract(buyer.rpcClient, common.HexToAddress(string(addr)))
					contractMsg := createContractMsg(common.HexToAddress(string(addr)), contractValues, true)
					buyer.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(addr), contractMsg)
					buyer.api.Contract.UpdateContract(string(addr), msgdata.ConvertContractMSGtoContractJSON(contractMsg))

					buyer.msg.ActiveContracts[addr] = false
					buyer.msg.RunningContracts[addr] = true
					buyer.hashrateUpdateChans[addr] = make(chan bool)
					buyer.ps.SetWait(msgbus.BuyerMsg, msgbus.IDString(buyer.msg.ID), buyer.msg)
					buyer.api.Buyer.UpdateBuyer(string(buyer.msg.ID), msgdata.ConvertBuyerMSGtoBuyerJSON(buyer.msg))
					
				case contractClosedSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Closed \n\n", addr)

					closedContractValues := readHashrateContract(buyer.rpcClient, common.HexToAddress(string(addr)))
					closedContractMsg := createContractMsg(common.HexToAddress(string(addr)), closedContractValues, true)
					buyer.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(closedContractMsg.ID), closedContractMsg)
					buyer.api.Contract.UpdateContract(string(addr), msgdata.ConvertContractMSGtoContractJSON(closedContractMsg))

					buyer.msg.RunningContracts[addr] = false
					buyer.msg.CompleteContracts[addr] = true
					delete(buyer.hashrateUpdateChans, addr)
					buyer.ps.SetWait(msgbus.BuyerMsg, msgbus.IDString(buyer.msg.ID), buyer.msg)
					buyer.api.Buyer.UpdateBuyer(string(buyer.msg.ID), msgdata.ConvertBuyerMSGtoBuyerJSON(buyer.msg))
				}
			}
		}
	}()

	var contractEvent msgbus.Event
	err = buyer.ps.Sub(msgbus.ContractMsg, msgbus.IDString(addr), contractEventChan)
	if err != nil {
		log.Fatal(err)
	}
	// once contract is running, closeout after length of contract has passed if it was not closed out early
	for {
		contractEvent = <-contractEventChan
		if contractEvent.EventType == msgbus.UpdateEvent {
			runningContractMsg := contractEvent.Data.(msgbus.Contract)
			if runningContractMsg.State == msgbus.ContRunningState {
				// run routine for each running contract to check if contract length has passed and contract should be closed out
				fmt.Println("starting closeout monitor for: ", runningContractMsg.ID)
				go buyer.closeOutMonitor(runningContractMsg)
			}
		}		
	}
}

func (buyer *BuyerContractManager) minerMonitor(ch msgbus.EventChan) {
	// subscribe channel to existing miners 
	for miner := range buyer.miners {
		e1, err := buyer.ps.SubWait(msgbus.MinerMsg, msgbus.IDString(miner), ch)
		if err != nil {
			panic(fmt.Sprintf("SubWait failed: %s\n", err))
		}
		if e1.EventType != msgbus.SubscribedEvent {
			panic(fmt.Sprintf("Wrong event type %v\n", e1))
		}
	}
	for {
		event := <-ch
		id := msgbus.MinerID(event.ID)

		switch event.EventType {

			//
			// Publish Event
			//
		case msgbus.PublishEvent:
			// Is this a new miner?

			fmt.Printf("Got PublishEvent: %v\n", event)

			if _, ok := buyer.miners[id]; !ok {
				buyer.miners[id] = event.Data.(msgbus.Miner)

				// let closeout monitor for running contracts know about new miner
				for addr := range buyer.hashrateUpdateChans {
					buyer.hashrateUpdateChans[addr]<-true
				}
				
				//
				// Use the existing channel to monitor
				//
				e1, err := buyer.ps.SubWait(msgbus.MinerMsg, event.ID, ch)
				if err != nil {
					panic(fmt.Sprintf("SubWait failed: %s\n", err))
				}
				if e1.EventType != msgbus.SubscribedEvent {
					panic(fmt.Sprintf("Wrong event type %v\n", e1))
				}
			} else {
				panic(fmt.Sprintf("Got PubEvent, but already had the ID: %v\n", event))
			}

			//
			// Delete/Unsubscribe Event
			//
		case msgbus.DeleteEvent:
			fallthrough
		case msgbus.UnsubscribedEvent:
			fmt.Printf("Miner Delete/Unsubscribe Event:%v\n", event)

			if _, ok := buyer.miners[id]; ok {
				delete(buyer.miners, id)
			} else {
				panic(fmt.Sprintf("Got UnsubscribeEvent, but dont have the ID: %v\n", event))
			}

			// let closeout monitor for running contracts know about new miner
			for addr := range buyer.hashrateUpdateChans {
				buyer.hashrateUpdateChans[addr]<-true
			}
			//
			// Update Event
			//
		case msgbus.UpdateEvent:
			if _, ok := buyer.miners[id]; !ok {
				panic(fmt.Sprintf("Got Miner ID does not exist: %v\n", event))
			}

			// Update the current miner data
			buyer.miners[id] = event.Data.(msgbus.Miner)
			
			// let closeout monitor for running contracts know about new miner
			for addr := range buyer.hashrateUpdateChans {
				buyer.hashrateUpdateChans[addr]<-true
			}
		default:
			fmt.Printf("Got Event: %v\n", event)
		}
	}
}

func (buyer *BuyerContractManager) closeOutMonitor(contractMsg msgbus.Contract) {
	addr := common.HexToAddress(string(contractMsg.ID))
	
	// check current hashrate being delivered to node
	if !buyer.checkHashRate(addr) {
		// hashrate is not greater than 0
		return
	}

	// check hashrate when miners are published, updated, deleted
	loop:
	for {
		<-buyer.hashrateUpdateChans[contractMsg.ID]
		//fmt.Println("Got miner update")
		if !buyer.checkHashRate(addr) {
			// hashrate is not greater than 0
			break loop
		}
	}
}

func (buyer *BuyerContractManager) checkHashRate(addr common.Address) bool {
	// check for miners in the online state and add up total hashrate being delivered to node
	totalHashRate := 0
	for i := range buyer.miners {
		if buyer.miners[i].State == msgbus.OnlineState {
			totalHashRate += buyer.miners[i].CurrentHashRate
		}
	}
	
	fmt.Println("Total Hashrate: ", totalHashRate)
	if totalHashRate == 0 {
		log.Printf("Closing out contract %s for not meeting hashrate requirements\n", addr.Hex())
		var wg sync.WaitGroup
		wg.Add(1)
		setContractCloseOut(buyer.rpcClient, buyer.account, buyer.privateKey, addr, &wg, &buyer.currentNonce)
		wg.Wait()
		return false
	}

	log.Printf("Hashrate promised by Hashrate Contract Address: %s is being fulfilled", addr.Hex())
	return true
}

func setUpClient(clientAddress string, contractManagerAccount common.Address) (client *ethclient.Client, err error) {
	client, err = ethclient.Dial(clientAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to rpc client at %v\n", clientAddress)

	var balance *big.Int
	balance, err = client.BalanceAt(context.Background(), contractManagerAccount, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("Balance of contract manager account:", ethValue, "ETH")

	return client, err
}

func subscribeToContractEvents(client *ethclient.Client, contractAddress common.Address) (chan types.Log, ethereum.Subscription) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	return logs, sub
}

func readHashrateContract(client *ethclient.Client, contractAddress common.Address) hashrateContractValues {
	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	var contractValues hashrateContractValues

	state, err := instance.ContractState(nil)
	if err != nil {
		log.Fatal(err)
	}
	contractValues.State = state

	price, err := instance.Price(nil)
	if err != nil {
		log.Fatal(err)
	}
	contractValues.Price = int(price.Int64())

	limit, err := instance.Limit(nil)
	if err != nil {
		log.Fatal(err)
	}
	contractValues.Limit = int(limit.Int64())

	speed, err := instance.Speed(nil)
	if err != nil {
		log.Fatal(err)
	}
	contractValues.Speed = int(speed.Int64())

	length, err := instance.Length(nil)
	if err != nil {
		log.Fatal(err)
	}
	contractValues.Length = int(length.Int64())

	validationFee, err := instance.ValidationFee(nil)
	if err != nil {
		log.Fatal(err)
	}
	contractValues.ValidationFee = int(validationFee.Int64())

	startingBlockTimestamp, err := instance.StartingBlockTimestamp(nil)
	if err != nil {
		log.Fatal(err)
	}
	contractValues.StartingBlockTimestamp = int(startingBlockTimestamp.Int64())

	buyer, err := instance.Buyer(nil)
	if err != nil {
		log.Fatal(err)
	}
	contractValues.Buyer = buyer

	seller, err := instance.Seller(nil)
	if err != nil {
		log.Fatal(err)
	}
	contractValues.Seller = seller

	return contractValues
}

func readDestUrl(client *ethclient.Client, contractAddress common.Address, privateKeyString string) string {
	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Getting Dest url from contract %s\n\n", contractAddress)

	encryptedDestUrl, err := instance.EncryptedPoolData(nil)
	if err != nil {
		log.Fatal(err)
	}

	/*
	Decryption Logic:

	destUrlBytes := []byte(encryptedDestUrl)
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyECIES := ecies.ImportECDSA(privateKey)
	decryptedDestUrlBytes, err := privateKeyECIES.Decrypt(destUrlBytes, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	decryptedDestUrl := string(decryptedDestUrlBytes)

	return decryptedDestUrl

	*/
	return encryptedDestUrl
}

func setContractCloseOut(client *ethclient.Client, fromAddress common.Address, privateKeyString string, contractAddress common.Address, wg *sync.WaitGroup, currentNonce *nonce) {
	defer wg.Done()
	defer currentNonce.mutex.Unlock()
	
	currentNonce.mutex.Lock()

	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000) // in units
	auth.Value = big.NewInt(0)      // in wei

	currentNonce.nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(currentNonce.nonce))

	tx, err := instance.SetContractCloseOut(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
	fmt.Println("Closing Out Contract: ", contractAddress)
}

func createContractMsg(contractAddress common.Address, contractValues hashrateContractValues, isSeller bool) msgbus.Contract {
	convertToMsgBusState := map[uint8]msgbus.ContractState{
		AvailableState: msgbus.ContAvailableState,
		ActiveState:    msgbus.ContActiveState,
		RunningState:   msgbus.ContRunningState,
		CompleteState:  msgbus.ContCompleteState,
	}

	var contractMsg msgbus.Contract
	contractMsg.IsSeller = isSeller
	contractMsg.ID = msgbus.ContractID(contractAddress.Hex())
	contractMsg.State = convertToMsgBusState[contractValues.State]
	contractMsg.Buyer = msgbus.BuyerID(contractValues.Buyer.Hex())
	contractMsg.Price = contractValues.Price
	contractMsg.Limit = contractValues.Limit
	contractMsg.Speed = contractValues.Speed
	contractMsg.Length = contractValues.Length
	contractMsg.ValidationFee = contractValues.ValidationFee
	contractMsg.StartingBlockTimestamp = contractValues.StartingBlockTimestamp

	return contractMsg
}