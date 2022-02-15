package contractmanager

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"sync"
	"time"

	//"encoding/hex"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	//"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/clonefactory"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
)

const (
	AvailableState uint8 = 0
	RunningState   uint8 = 1
)

const HASHRATE_TOLERANCE = .10 

type hashrateContractValues struct {
	State                  uint8
	Price                  int
	Limit                  int
	Speed                  int
	Length                 int
	StartingBlockTimestamp int
	Buyer                  common.Address
	Seller                 common.Address
}

type nonce struct {
	mutex sync.Mutex
	nonce uint64
}

type ContractManager interface {
	start() (err error)
	init(ctx *context.Context, ps *msgbus.PubSub, contractManagerConfigID msgbus.IDString, nodeOperatorMsg *msgbus.NodeOperator) (err error)
	setupExistingContracts() (err error)
	readContracts() ([]common.Address, error)
	watchHashrateContract(addr msgbus.ContractID, hrLogs chan types.Log, hrSub ethereum.Subscription)
}

type SellerContractManager struct {
	ps                  *msgbus.PubSub
	ethClient           *ethclient.Client
	cloneFactoryAddress common.Address
	account             common.Address
	privateKey          string
	claimFunds          bool
	currentNonce        nonce
	nodeOperator        msgbus.NodeOperator
	ctx 				context.Context
}

type BuyerContractManager struct {
	ps                  *msgbus.PubSub
	ethClient           *ethclient.Client
	cloneFactoryAddress common.Address
	account             common.Address
	privateKey          string
	currentNonce        nonce
	nodeOperator        msgbus.NodeOperator
	ctx 				context.Context
}

func Run(ctx *context.Context, contractManager ContractManager, ps *msgbus.PubSub, contractManagerConfigID msgbus.IDString, nodeOperatorMsg *msgbus.NodeOperator) (err error) {
	contractManagerCtx, contractManagerCancel := context.WithCancel(*ctx)
	go newConfigMonitor(ctx, contractManagerCancel, contractManager, ps, contractManagerConfigID, nodeOperatorMsg)

	err = contractManager.init(&contractManagerCtx, ps, contractManagerConfigID, nodeOperatorMsg)
	if err != nil {
		return err
	}
	err = contractManager.start()
	if err != nil {
		return err
	}
	
	return err
}

func newConfigMonitor(ctx *context.Context, cancel context.CancelFunc, contractManager ContractManager, ps *msgbus.PubSub, contractManagerConfigID msgbus.IDString, nodeOperatorMsg *msgbus.NodeOperator) {
	contractConfigCh := ps.NewEventChan()
	event, err := ps.SubWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractConfigCh)
	if err != nil {
		panic(fmt.Sprintf("SubWait failed: %s\n", err))
	}
	if event.EventType != msgbus.SubscribedEvent {
		panic(fmt.Sprintf("Wrong event type %v\n", event))
	}

	for event = range contractConfigCh {
		if event.EventType == msgbus.UpdateEvent {
			fmt.Printf("Updated Contract Manager Configuration: Restarting Contract Manager: %v\n", event)
			cancel()
			err = Run(ctx, contractManager, ps, contractManagerConfigID, nodeOperatorMsg)
			if err != nil {
				panic(fmt.Sprintf("contract manager failed to run:%s", err))
			}
			return
		}
	}
}

func (seller *SellerContractManager) init(ctx *context.Context, ps *msgbus.PubSub, contractManagerConfigID msgbus.IDString, nodeOperatorMsg *msgbus.NodeOperator) (err error) {
	seller.ctx = *ctx
	
	event, err := ps.GetWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID) 
	if err != nil {
		return err
	}
	contractManagerConfig := event.Data.(msgbus.ContractManagerConfig)
	seller.claimFunds = contractManagerConfig.ClaimFunds 
	ethNodeAddr := contractManagerConfig.EthNodeAddr 
	mnemonic := contractManagerConfig.Mnemonic
	accountIndex := contractManagerConfig.AccountIndex

	account,privateKey := hdWalletKeys(mnemonic, accountIndex)
	seller.account = account.Address
	seller.privateKey = privateKey
	
	var client *ethclient.Client
	client, err = setUpClient(ethNodeAddr, seller.account)
	if err != nil {
		return err
	}
	seller.ps = ps
	seller.ethClient = client
	seller.cloneFactoryAddress = common.HexToAddress(contractManagerConfig.CloneFactoryAddress)
	
	seller.nodeOperator = *nodeOperatorMsg
	seller.nodeOperator.EthereumAccount = seller.account.Hex()

	if seller.nodeOperator.Contracts == nil {
		seller.nodeOperator.Contracts = make(map[msgbus.ContractID]msgbus.ContractState)
	}
	
	return err
}

func (seller *SellerContractManager) start() (err error) {
	err = seller.setupExistingContracts()
	if err != nil {
		return err
	}

	// routine for listensing to contract creation events that will update seller msg with new contracts and load new contract onto msgbus
	cfLogs, cfSub, err := subscribeToContractEvents(seller.ethClient, seller.cloneFactoryAddress)
	if err != nil {
		return err
	}
	go seller.watchContractCreation(cfLogs, cfSub)

	// routine starts routines for seller's contracts that monitors contract purchase, close, and cancel events
	go func() {
		// start routines for existing contracts
		for addr := range seller.nodeOperator.Contracts {
			hrLogs, hrSub, err := subscribeToContractEvents(seller.ethClient, common.HexToAddress(string(addr)))
			if err != nil {
				panic(fmt.Sprintf("Failed to subscribe to events on hashrate contract %s, Fileline::%s, Error::%v\n", addr, lumerinlib.FileLine(), err))
			}
			go seller.watchHashrateContract(addr, hrLogs, hrSub)
		}

		// monitor new contracts getting created and start hashrate conrtract monitor routine when they are created
		contractEventChan := seller.ps.NewEventChan()
		err = seller.ps.Sub(msgbus.ContractMsg, "", contractEventChan)
		if err != nil {
			panic(fmt.Sprintf("Failed to subscribe to contract events on msgbus, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err))
		}
		for {
			select {
			case <-seller.ctx.Done():
				fmt.Println("Cancelling current contract manager context: cancelling start routine")
				return
			case event := <-contractEventChan:
				if event.EventType == msgbus.PublishEvent {
					newContract := event.Data.(msgbus.Contract)
					if newContract.State == msgbus.ContAvailableState {
						addr := common.HexToAddress(string(newContract.ID))
						hrLogs, hrSub, err := subscribeToContractEvents(seller.ethClient, addr)
						if err != nil {
							panic(fmt.Sprintf("Failed to subscribe to events on hashrate contract %s, Fileline::%s, Error::%v\n", newContract.ID, lumerinlib.FileLine(), err))
						}
						go seller.watchHashrateContract(msgbus.ContractID(addr.Hex()), hrLogs, hrSub)
					}
				}
			}
		}
	}()
	return err
}

func (seller *SellerContractManager) setupExistingContracts() (err error) {
	var contractValues []hashrateContractValues
	var contractMsgs []msgbus.Contract

	sellerContracts, err := seller.readContracts()
	if err != nil {
		return err
	}
	fmt.Println("Existing Seller Contracts: ", sellerContracts)
	
	for i := range sellerContracts {
		id := msgbus.ContractID(sellerContracts[i].Hex())
		if _, ok := seller.nodeOperator.Contracts[id]; !ok {
			contract, err := readHashrateContract(seller.ethClient, sellerContracts[i])
			if err != nil {
				return err
			}
			contractValues = append(contractValues, contract)
			contractMsgs = append(contractMsgs, createContractMsg(sellerContracts[i], contractValues[i], true))
			seller.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])
	
			seller.nodeOperator.Contracts[msgbus.ContractID(sellerContracts[i].Hex())] = msgbus.ContAvailableState
	
			if contractValues[i].State == RunningState {
				seller.nodeOperator.Contracts[msgbus.ContractID(sellerContracts[i].Hex())] = msgbus.ContRunningState
			}
		}
	}

	seller.ps.SetWait(msgbus.NodeOperatorMsg, msgbus.IDString(seller.nodeOperator.ID), seller.nodeOperator)
	
	return err
}

func (seller *SellerContractManager) readContracts() ([]common.Address, error) {
	var sellerContractAddresses []common.Address
	var hashrateContractInstance *implementation.Implementation
	var hashrateContractSeller common.Address

	instance, err := clonefactory.NewClonefactory(seller.cloneFactoryAddress, seller.ethClient)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return sellerContractAddresses, err
	}

	hashrateContractAddresses, err := instance.GetContractList(&bind.CallOpts{})
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return sellerContractAddresses, err
	}

	// parse existing hashrate contracts for ones that belong to seller
	for i := range hashrateContractAddresses {
		hashrateContractInstance, err = implementation.NewImplementation(hashrateContractAddresses[i], seller.ethClient)
		if err != nil {
			log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
			return sellerContractAddresses, err
		}
		hashrateContractSeller, err = hashrateContractInstance.Seller(nil)
		if err != nil {
			log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
			return sellerContractAddresses, err
		}
		if hashrateContractSeller == seller.account {
			sellerContractAddresses = append(sellerContractAddresses, hashrateContractAddresses[i])
		}
	}

	return sellerContractAddresses, err
}

func (seller *SellerContractManager) watchContractCreation(cfLogs chan types.Log, cfSub ethereum.Subscription) {
	defer close(cfLogs)
	defer cfSub.Unsubscribe()

	// create event signature to parse out creation event
	contractCreatedSig := []byte("contractCreated(address)")
	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
	for {
		select {
		case err := <-cfSub.Err():
			panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
		case <-seller.ctx.Done():
			fmt.Println("Cancelling current contract manager context: cancelling watchContractCreation go routine")
			return
		case cfLog := <-cfLogs:
			if cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex() {
				address := common.HexToAddress(cfLog.Topics[1].Hex())
				// check if contract created belongs to seller
				hashrateContractInstance, err := implementation.NewImplementation(address, seller.ethClient)
				if err != nil {
					panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
				}
				hashrateContractSeller, err := hashrateContractInstance.Seller(nil)
				if err != nil {
					panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
				}
				if hashrateContractSeller == seller.account {
					fmt.Printf("Address of created Hashrate Contract: %s\n\n", address.Hex())

					createdContractValues, err := readHashrateContract(seller.ethClient, address)
					if err != nil {
						panic(fmt.Sprintf("Reading hashrate contract failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
					}
					createdContractMsg := createContractMsg(address, createdContractValues, true)
					seller.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(address.Hex()), createdContractMsg)

					seller.nodeOperator.Contracts[msgbus.ContractID(address.Hex())] = msgbus.ContAvailableState

					seller.ps.SetWait(msgbus.NodeOperatorMsg, msgbus.IDString(seller.nodeOperator.ID), seller.nodeOperator)
				}
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
	contractClosedSig := []byte("contractClosed()")
	purchaseInfoUpdatedSig := []byte("purchaseInfoUpdated()")
	cipherTextUpdatedSig := []byte("cipherTextUpdated(string)")
	contractPurchasedSigHash := crypto.Keccak256Hash(contractPurchasedSig)
	contractClosedSigHash := crypto.Keccak256Hash(contractClosedSig)
	purchaseInfoUpdatedSigHash := crypto.Keccak256Hash(purchaseInfoUpdatedSig)
	cipherTextUpdatedSigHash := crypto.Keccak256Hash(cipherTextUpdatedSig)

	// routine monitoring and acting upon events emmited by hashrate contract
	go func() {
		defer close(hrLogs)
		defer hrSub.Unsubscribe()
		for {
			select {
			case err := <-hrSub.Err():
				panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
			case <-seller.ctx.Done():
				fmt.Println("Cancelling current contract manager context: cancelling watchHashrateContract go routine")
				return
			case hLog := <-hrLogs:
				switch hLog.Topics[0].Hex() {
				case contractPurchasedSigHash.Hex():
					buyer := common.HexToAddress(hLog.Topics[1].Hex())
					fmt.Printf("%s purchased Hashrate Contract: %s\n\n", buyer.Hex(), addr)

					destUrl, err := readDestUrl(seller.ethClient, common.HexToAddress(string(addr)), seller.privateKey)
					if err != nil {
						panic(fmt.Sprintf("Reading dest url failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
					}
					destMsg := msgbus.Dest{
						ID:     msgbus.DestID(msgbus.GetRandomIDString()),
						NetUrl: msgbus.DestNetUrl(destUrl),
					}
					seller.ps.PubWait(msgbus.DestMsg, msgbus.IDString(destMsg.ID), destMsg)

					event, err := seller.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
					if err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
					}
					if event.Err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
					}
					contractValues, err := readHashrateContract(seller.ethClient, common.HexToAddress(string(addr)))
					if err != nil {
						panic(fmt.Sprintf("Reading hashrate contract failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
					}
					contractMsg := createContractMsg(common.HexToAddress(string(addr)), contractValues, true)
					contractMsg.Dest = destMsg.ID
					contractMsg.State = msgbus.ContRunningState
					contractMsg.Buyer = string(buyer.Hex())
					seller.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(addr), contractMsg)

					seller.nodeOperator.Contracts[addr] = msgbus.ContRunningState
					seller.ps.SetWait(msgbus.NodeOperatorMsg, msgbus.IDString(seller.nodeOperator.ID), seller.nodeOperator)

				case cipherTextUpdatedSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Cipher Text Updated \n\n", addr)

					event, err := seller.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
					if err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
					}
					if event.Err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
					}
					contractMsg := event.Data.(msgbus.Contract)
					event, err = seller.ps.GetWait(msgbus.DestMsg, msgbus.IDString(contractMsg.Dest))
					if err != nil {
						panic(fmt.Sprintf("Getting Dest Failed: %s", err))
					}
					if event.Err != nil {
						panic(fmt.Sprintf("Getting Dest Failed: %s", event.Err))
					}
					destMsg := event.Data.(msgbus.Dest)

					destUrl, err := readDestUrl(seller.ethClient, common.HexToAddress(string(addr)), seller.privateKey)
					if err != nil {
						panic(fmt.Sprintf("Reading dest url failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
					}
					destMsg.NetUrl = msgbus.DestNetUrl(destUrl)
					seller.ps.SetWait(msgbus.DestMsg, msgbus.IDString(destMsg.ID), destMsg)

				case contractClosedSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Closed \n\n", addr)

					event, err := seller.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
					if err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
					}
					if event.Err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
					}
					contractMsg := event.Data.(msgbus.Contract)
					contractMsg.State = msgbus.ContAvailableState
					contractMsg.Buyer = ""
					seller.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contractMsg.ID), contractMsg)

					seller.nodeOperator.Contracts[addr] = msgbus.ContAvailableState
					seller.ps.SetWait(msgbus.NodeOperatorMsg, msgbus.IDString(seller.nodeOperator.ID), seller.nodeOperator)

				case purchaseInfoUpdatedSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Purchase Info Updated \n\n", addr)

					event, err := seller.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
					if err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
					}
					if event.Err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
					}
					contractMsg := event.Data.(msgbus.Contract)

					updatedContractValues, err := readHashrateContract(seller.ethClient, common.HexToAddress(string(addr)))
					if err != nil {
						panic(fmt.Sprintf("Reading hashrate contract failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
					}
					updateContractMsg(&contractMsg, updatedContractValues)
					seller.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contractMsg.ID), contractMsg)
				}
			}
		}
	}()

	err = seller.ps.Sub(msgbus.ContractMsg, msgbus.IDString(addr), contractEventChan)
	if err != nil {
		panic(fmt.Sprintf("Subscribing to Contract Failed: %s", err))
	}
	// once contract is running, closeout after length of contract has passed if it was not closed out early
	for {
		select {
		case <-seller.ctx.Done():
			fmt.Println("Cancelling current contract manager context: cancelling watchHashrateContract go routine")
			return
		case event := <-contractEventChan:
			if event.EventType == msgbus.UpdateEvent {
				runningContractMsg := event.Data.(msgbus.Contract)
				if runningContractMsg.State == msgbus.ContRunningState {
					// run routine for each running contract to check if contract length has passed and contract should be closed out
					go seller.closeOutMonitor(runningContractMsg)
				}
			}
		}
	}
}

func (seller *SellerContractManager) closeOutMonitor(contractMsg msgbus.Contract) {
	contractFinishedTimestamp := contractMsg.StartingBlockTimestamp + contractMsg.Length

	// subscribe to latest block headers
	headers := make(chan *types.Header)
	sub, err := seller.ethClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
	}
	defer close(headers)
	defer sub.Unsubscribe()

	loop:
	for {
		select {
		case err := <-sub.Err():
			panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
		case <-seller.ctx.Done():
			fmt.Println("Cancelling current contract manager context: cancelling closeout monitor go routine")
			return
		case header := <-headers:
			// get latest block from header
			block, err := seller.ethClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
			}

			// check if contract length has passed
			if block.Time() >= uint64(contractFinishedTimestamp) {
				var closeOutType uint

				// seller only wants to closeout
				closeOutType = 2
				// seller wants to claim funds with closeout
				if seller.claimFunds {
					closeOutType = 3
				}

				// if contract was not already closed early, close out here
				contractValues, err := readHashrateContract(seller.ethClient, common.HexToAddress(string(contractMsg.ID)))
				if err != nil {
					panic(fmt.Sprintf("Reading hashrate contract failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
				}
				if contractValues.State == RunningState {
					var wg sync.WaitGroup
					wg.Add(1)
					err = setContractCloseOut(seller.ethClient, seller.account, seller.privateKey, common.HexToAddress(string(contractMsg.ID)), &wg, &seller.currentNonce, closeOutType)
					if err != nil {
						panic(fmt.Sprintf("Contract Close Out failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
					}
					wg.Wait()
				}
				break loop
			}
		}
	}
}

func (buyer *BuyerContractManager) init(ctx *context.Context, ps *msgbus.PubSub, contractManagerConfigID msgbus.IDString, nodeOperatorMsg *msgbus.NodeOperator) (err error) {
	buyer.ctx = *ctx
	
	event, err := ps.GetWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID) 
	if err != nil {
		return err
	}
	contractManagerConfig := event.Data.(msgbus.ContractManagerConfig)
	ethNodeAddr := contractManagerConfig.EthNodeAddr 
	mnemonic := contractManagerConfig.Mnemonic
	accountIndex := contractManagerConfig.AccountIndex

	account,privateKey := hdWalletKeys(mnemonic, accountIndex)
	buyer.account = account.Address
	buyer.privateKey = privateKey
	
	var client *ethclient.Client
	client, err = setUpClient(ethNodeAddr, buyer.account)
	if err != nil {
		return err
	}
	buyer.ps = ps
	buyer.ethClient = client
	buyer.cloneFactoryAddress = common.HexToAddress(contractManagerConfig.CloneFactoryAddress)
	
	buyer.nodeOperator = *nodeOperatorMsg
	buyer.nodeOperator.EthereumAccount = buyer.account.Hex()

	if buyer.nodeOperator.Contracts == nil {
		buyer.nodeOperator.Contracts = make(map[msgbus.ContractID]msgbus.ContractState)
	}

	return err
}

func (buyer *BuyerContractManager) start() (err error) {
	err = buyer.setupExistingContracts()
	if err != nil {
		return err
	}

	// subcribe to events emitted by clonefactory contract to read contract purchase event
	cfLogs, cfSub, err := subscribeToContractEvents(buyer.ethClient, buyer.cloneFactoryAddress)
	if err != nil {
		return err
	}

	// routine for listensing to contract purchase events to update buyer with new contracts they purchased
	go buyer.watchContractPurchase(cfLogs, cfSub)

	// miner event channel for miner monitor that checks miner publishes, updates, and deletes
	minerEventChan := buyer.ps.NewEventChan()
	err = buyer.ps.Sub(msgbus.MinerMsg, msgbus.IDString(""), minerEventChan)
	if err != nil {
		panic(fmt.Sprintf("Failed to subscribe to miner events on msgbus, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err))
	}

	// routine starts routines for buyers's contracts that monitors contract running and close events
	go func() {
		// start watch hashrate contract for existing running contracts
		for addr := range buyer.nodeOperator.Contracts {
			hrLogs, hrSub, err := subscribeToContractEvents(buyer.ethClient, common.HexToAddress(string(addr)))
			if err != nil {
				panic(fmt.Sprintf("Failed to subscribe to events on hashrate contract %s, Fileline::%s, Error::%s\n", addr, lumerinlib.FileLine(), err))
			}
			go buyer.watchHashrateContract(addr, hrLogs, hrSub)

			contractEventChan := buyer.ps.NewEventChan()
			err = buyer.ps.Sub(msgbus.ContractMsg, msgbus.IDString(addr), contractEventChan)
			if err != nil {
				panic(fmt.Sprintf("Failed to subscribe to contract %s events, Fileline::%s, Error::%s\n", addr, lumerinlib.FileLine(), err))
			}
			go buyer.closeOutMonitor(minerEventChan, contractEventChan, addr)
		}

		// monitor new contracts getting purchased and start watch hashrate conrtract routine when they are purchased
		contractEventChan := buyer.ps.NewEventChan()
		err := buyer.ps.Sub(msgbus.ContractMsg, "", contractEventChan)
		if err != nil {
			panic(fmt.Sprintf("Failed to subscribe to contract events on msgbus, Fileline::%s, Error::%s\n", lumerinlib.FileLine(), err))
		}
		for {
			select {
			case <-buyer.ctx.Done():
				fmt.Println("Cancelling current contract manager context: cancelling start routine")
				return
			case event := <-contractEventChan:
				if event.EventType == msgbus.PublishEvent {
					newContract := event.Data.(msgbus.Contract)
					addr := common.HexToAddress(string(newContract.ID))
					hrLogs, hrSub, err := subscribeToContractEvents(buyer.ethClient, addr)
					if err != nil {
						panic(fmt.Sprintf("Failed to subscribe to events on hashrate contract %s, Fileline::%s, Error::%s\n", addr, lumerinlib.FileLine(), err))
					}
					go buyer.watchHashrateContract(msgbus.ContractID(addr.Hex()), hrLogs, hrSub)

					newContractEventChan := buyer.ps.NewEventChan()
					err = buyer.ps.Sub(msgbus.ContractMsg, msgbus.IDString(newContract.ID), newContractEventChan)
					if err != nil {
						panic(fmt.Sprintf("Failed to subscribe to contract %s events, Fileline::%s, Error::%s\n", newContract.ID, lumerinlib.FileLine(), err))
					}
					go buyer.closeOutMonitor(minerEventChan, newContractEventChan, newContract.ID)
				}
			}
		}
	}()
	return nil
}

func (buyer *BuyerContractManager) setupExistingContracts() (err error) {
	var contractValues []hashrateContractValues
	var contractMsgs []msgbus.Contract
	var nodeOperatorUpdated bool

	buyerContracts, err := buyer.readContracts()
	if err != nil {
		return err
	}
	fmt.Println("Existing Buyer Contracts: ", buyerContracts)

	for i := range buyerContracts {
		id := msgbus.ContractID(buyerContracts[i].Hex())
		if _, ok := buyer.nodeOperator.Contracts[id]; !ok {
			contract, err := readHashrateContract(buyer.ethClient, buyerContracts[i])
			if err != nil {
				return err
			}
			contractValues = append(contractValues, contract)
			contractMsgs = append(contractMsgs, createContractMsg(buyerContracts[i], contractValues[i], false))
			buyer.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])
	
			buyer.nodeOperator.Contracts[msgbus.ContractID(buyerContracts[i].Hex())] = msgbus.ContRunningState
			nodeOperatorUpdated = true
		}	
	}

	if nodeOperatorUpdated {
		buyer.ps.PubWait(msgbus.NodeOperatorMsg, msgbus.IDString(buyer.nodeOperator.ID), buyer.nodeOperator)
	}

	return err
}

func (buyer *BuyerContractManager) readContracts() ([]common.Address, error) {
	var buyerContractAddresses []common.Address
	var hashrateContractInstance *implementation.Implementation
	var hashrateContractBuyer common.Address

	instance, err := clonefactory.NewClonefactory(buyer.cloneFactoryAddress, buyer.ethClient)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return buyerContractAddresses, err
	}

	hashrateContractAddresses, err := instance.GetContractList(&bind.CallOpts{})
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return buyerContractAddresses, err
	}

	// parse existing hashrate contracts for ones that belong to buyer
	for i := range hashrateContractAddresses {
		hashrateContractInstance, err = implementation.NewImplementation(hashrateContractAddresses[i], buyer.ethClient)
		if err != nil {
			log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
			return buyerContractAddresses, err
		}
		hashrateContractBuyer, err = hashrateContractInstance.Buyer(nil)
		if err != nil {
			log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
			return buyerContractAddresses, err
		}
		if hashrateContractBuyer == buyer.account {
			buyerContractAddresses = append(buyerContractAddresses, hashrateContractAddresses[i])
		}
	}

	return buyerContractAddresses, err
}

func (buyer *BuyerContractManager) watchContractPurchase(cfLogs chan types.Log, cfSub ethereum.Subscription) {
	defer close(cfLogs)
	defer cfSub.Unsubscribe()

	// create event signature to parse out purchase event
	clonefactoryContractPurchasedSig := []byte("clonefactoryContractPurchased(address)")
	clonefactoryContractPurchasedSigHash := crypto.Keccak256Hash(clonefactoryContractPurchasedSig)

	for {
		select {
		case <-buyer.ctx.Done():
			fmt.Println("Cancelling current contract manager context: cancelling watchContractPurchase routine")
			return
		case err := <-cfSub.Err():
			panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
		case cfLog := <-cfLogs:
			if cfLog.Topics[0].Hex() == clonefactoryContractPurchasedSigHash.Hex() {
				address := common.HexToAddress(cfLog.Topics[1].Hex())
				// check if contract was purchased by buyer
				hashrateContractInstance, err := implementation.NewImplementation(address, buyer.ethClient)
				if err != nil {
					panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
				}
				hashrateContractBuyer, err := hashrateContractInstance.Buyer(nil)
				if err != nil {
					panic(fmt.Sprintf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err))
				}
				if hashrateContractBuyer == buyer.account {
					fmt.Printf("Address of purchased Hashrate Contract : %s\n\n", address.Hex())

					destUrl, err := readDestUrl(buyer.ethClient, common.HexToAddress(string(address.Hex())), buyer.privateKey)
					if err != nil {
						panic(fmt.Sprintf("Reading dest url failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
					}
					destMsg := msgbus.Dest{
						ID:     msgbus.DestID(msgbus.GetRandomIDString()),
						NetUrl: msgbus.DestNetUrl(destUrl),
					}
					buyer.ps.PubWait(msgbus.DestMsg, msgbus.IDString(destMsg.ID), destMsg)

					purchasedContractValues, err := readHashrateContract(buyer.ethClient, address)
					if err != nil {
						panic(fmt.Sprintf("Reading hashrate contract failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
					}
					contractMsg := createContractMsg(address, purchasedContractValues, false)
					contractMsg.Dest = destMsg.ID
					contractMsg.State = msgbus.ContRunningState
					buyer.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsg.ID), contractMsg)

					buyer.nodeOperator.Contracts[contractMsg.ID] = msgbus.ContRunningState
					buyer.ps.SetWait(msgbus.NodeOperatorMsg, msgbus.IDString(buyer.nodeOperator.ID), buyer.nodeOperator)
				}
			}
		}
	}
}

func (buyer *BuyerContractManager) watchHashrateContract(addr msgbus.ContractID, hrLogs chan types.Log, hrSub ethereum.Subscription) {
	defer close(hrLogs)
	defer hrSub.Unsubscribe()

	// create event signatures to parse out which event was being emitted from hashrate contract
	contractClosedSig := []byte("contractClosed()")
	purchaseInfoUpdatedSig := []byte("purchaseInfoUpdated()")
	cipherTextUpdatedSig := []byte("cipherTextUpdated(string)")
	contractClosedSigHash := crypto.Keccak256Hash(contractClosedSig)
	purchaseInfoUpdatedSigHash := crypto.Keccak256Hash(purchaseInfoUpdatedSig)
	cipherTextUpdatedSigHash := crypto.Keccak256Hash(cipherTextUpdatedSig)

	// monitor events emmited by hashrate contract
	for {
		select {
		case <-buyer.ctx.Done():
			fmt.Println("Cancelling current contract manager context: cancelling watchHashrateContract go routine")
			return
		case err := <-hrSub.Err():
			log.Fatal(err)
		case hLog := <-hrLogs:
			switch hLog.Topics[0].Hex() {
			case contractClosedSigHash.Hex():
				fmt.Printf("Hashrate Contract %s Closed \n\n", addr)

				buyer.ps.Unpub(msgbus.ContractMsg, msgbus.IDString(addr))

				delete(buyer.nodeOperator.Contracts, addr)
				buyer.ps.SetWait(msgbus.NodeOperatorMsg, msgbus.IDString(buyer.nodeOperator.ID), buyer.nodeOperator)

			case purchaseInfoUpdatedSigHash.Hex():
				fmt.Printf("Hashrate Contract %s Purchase Info Updated \n\n", addr)

				event, err := buyer.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
				if err != nil {
					panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
				}
				if event.Err != nil {
					panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
				}
				contractMsg := event.Data.(msgbus.Contract)

				updatedContractValues, err := readHashrateContract(buyer.ethClient, common.HexToAddress(string(addr)))
				if err != nil {
					panic(fmt.Sprintf("Reading hashrate contract failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
				}
				updateContractMsg(&contractMsg, updatedContractValues)
				buyer.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contractMsg.ID), contractMsg)

			case cipherTextUpdatedSigHash.Hex():
				fmt.Printf("Hashrate Contract %s Cipher Text Updated \n\n", addr)

				event, err := buyer.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
				if err != nil {
					panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
				}
				if event.Err != nil {
					panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
				}
				contractMsg := event.Data.(msgbus.Contract)
				event, err = buyer.ps.GetWait(msgbus.DestMsg, msgbus.IDString(contractMsg.Dest))
				if err != nil {
					panic(fmt.Sprintf("Getting Dest Failed: %s", err))
				}
				if event.Err != nil {
					panic(fmt.Sprintf("Getting Dest Failed: %s", event.Err))
				}
				destMsg := event.Data.(msgbus.Dest)

				destUrl, err := readDestUrl(buyer.ethClient, common.HexToAddress(string(addr)), buyer.privateKey)
				if err != nil {
					panic(fmt.Sprintf("Reading dest url failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
				}
				destMsg.NetUrl = msgbus.DestNetUrl(destUrl)
				buyer.ps.SetWait(msgbus.DestMsg, msgbus.IDString(destMsg.ID), destMsg)
			}
		}
	}
}

func (buyer *BuyerContractManager) closeOutMonitor(minerCh msgbus.EventChan, contractCh msgbus.EventChan, contractId msgbus.ContractID) {
	for {
		select {
		case <-buyer.ctx.Done():
			fmt.Println("Cancelling current contract manager context: cancelling closeOutMonitor go routine")
			return
		case event := <-minerCh:
			if event.EventType == msgbus.PublishEvent || event.EventType == msgbus.UpdateEvent || event.EventType == msgbus.UnpublishEvent{
				// check hashrate is being fulfilled for all running contracts
				time.Sleep(time.Second * 10) // give buffer time for total hashrate to adjust to multiple updates
				contractClosed := buyer.checkHashRate(contractId)
				if contractClosed {
					return
				}
			}
		case event := <-contractCh:
			if event.EventType == msgbus.UnpublishEvent {
				return
			}
			if event.EventType == msgbus.PublishEvent || event.EventType == msgbus.UpdateEvent {
				// check hashrate is being fulfilled after contract update
				contractClosed := buyer.checkHashRate(contractId)
				if contractClosed {
					return
				}
			}
		}	
	}
}

func (buyer *BuyerContractManager) checkHashRate(contractId msgbus.ContractID) bool {
	// check for miners delivering hashrate for this contract
	totalHashrate := 0
	event,err := buyer.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(contractId))
	if err != nil {
		panic(fmt.Sprintf("Getting Hashrate Contract Failed: %s", err))
	}
	contract := event.Data.(msgbus.Contract)
	miners,err := buyer.ps.MinerGetAllWait()
	if err != nil {
		panic(fmt.Sprintf("Failed to get all miners, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err))
	}

	var miner *msgbus.Miner
	for i := range miners {
		miner,err = buyer.ps.MinerGetWait(miners[i]) 
		if err != nil {
			panic(fmt.Sprintf("Failed to get miner, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err))
		}
		if miner.Contract == contractId {
			totalHashrate += miner.CurrentHashRate
		}
	}

	promisedHashrateMin := int(float32(contract.Speed)*(1 - HASHRATE_TOLERANCE))

	fmt.Printf("Hashrate being sent to contract %s: %d\n", contractId, totalHashrate)
	if totalHashrate <= promisedHashrateMin {
		log.Printf("Closing out contract %s for not meeting hashrate requirements\n", contractId)
		var wg sync.WaitGroup
		wg.Add(1)
		err := setContractCloseOut(buyer.ethClient, buyer.account, buyer.privateKey, common.HexToAddress(string(contractId)), &wg, &buyer.currentNonce, 0)
		if err != nil {
			panic(fmt.Sprintf("Contract Close Out failed, Fileline::%s, Error::%v", lumerinlib.FileLine(), err))
		}
		wg.Wait()
		return true
	}

	log.Printf("Hashrate promised by contract %s is being fulfilled", contractId)
	return false
}

func hdWalletKeys(mnemonic string, accountIndex int) (accounts.Account, string) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/" + fmt.Sprint(accountIndex))
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	return account, privateKey
}


func setUpClient(clientAddress string, contractManagerAccount common.Address) (client *ethclient.Client, err error) {
	client, err = ethclient.Dial(clientAddress)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return client, err
	}

	fmt.Printf("Connected to rpc client at %v\n", clientAddress)

	var balance *big.Int
	balance, err = client.BalanceAt(context.Background(), contractManagerAccount, nil)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return client, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("Balance of contract manager account:", ethValue, "ETH")

	return client, err
}

func subscribeToContractEvents(client *ethclient.Client, contractAddress common.Address) (chan types.Log, ethereum.Subscription, error) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return logs, sub, err
	}

	return logs, sub, err
}

func readHashrateContract(client *ethclient.Client, contractAddress common.Address) (hashrateContractValues, error) {
	var contractValues hashrateContractValues

	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return contractValues, err
	}
	
	state,price,limit,speed,length,startingBlockTimestamp,buyer,seller,_,err := instance.GetPublicVariables(&bind.CallOpts{})
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return contractValues, err
	}
	contractValues.State = state
	contractValues.Price = int(price.Int64())
	contractValues.Limit = int(limit.Int64())
	contractValues.Speed = int(speed.Int64())
	contractValues.Length = int(length.Int64())
	contractValues.StartingBlockTimestamp = int(startingBlockTimestamp.Int64())
	contractValues.Buyer = buyer
	contractValues.Seller = seller

	return contractValues, err
}

func readDestUrl(client *ethclient.Client, contractAddress common.Address, privateKeyString string) (string, error) {
	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return "", err
	}

	fmt.Printf("Getting Dest url from contract %s\n\n", contractAddress)

	encryptedDestUrl, err := instance.EncryptedPoolData(nil)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return "", err
	}

	/*
	// Decryption Logic
	destUrlBytes,_ := hex.DecodeString(encryptedDestUrl)
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return "", err
	}
	privateKeyECIES := ecies.ImportECDSA(privateKey)
	decryptedDestUrlBytes, err := privateKeyECIES.Decrypt(destUrlBytes, nil, nil)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return "", err
	}
	decryptedDestUrl := string(decryptedDestUrlBytes)

	return decryptedDestUrl, err
	*/
	return encryptedDestUrl, err
}

func setContractCloseOut(client *ethclient.Client, fromAddress common.Address, privateKeyString string, contractAddress common.Address, wg *sync.WaitGroup, currentNonce *nonce, closeOutType uint) error {
	defer wg.Done()
	defer currentNonce.mutex.Unlock()

	currentNonce.mutex.Lock()

	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return err
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return err
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000) // in units
	auth.Value = big.NewInt(0)      // in wei

	currentNonce.nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return err
	}
	auth.Nonce = big.NewInt(int64(currentNonce.nonce))

	tx, err := instance.SetContractCloseOut(auth, big.NewInt(int64(closeOutType)))
	if err != nil {
		log.Printf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		return err
	}

	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
	fmt.Println("Closing Out Contract: ", contractAddress)
	return err
}

func createContractMsg(contractAddress common.Address, contractValues hashrateContractValues, isSeller bool) msgbus.Contract {
	convertToMsgBusState := map[uint8]msgbus.ContractState{
		AvailableState: msgbus.ContAvailableState,
		RunningState:   msgbus.ContRunningState,
	}

	var contractMsg msgbus.Contract
	contractMsg.IsSeller = isSeller
	contractMsg.ID = msgbus.ContractID(contractAddress.Hex())
	contractMsg.State = convertToMsgBusState[contractValues.State]
	contractMsg.Buyer = string(contractValues.Buyer.Hex())
	contractMsg.Price = contractValues.Price
	contractMsg.Limit = contractValues.Limit
	contractMsg.Speed = contractValues.Speed
	contractMsg.Length = contractValues.Length
	contractMsg.StartingBlockTimestamp = contractValues.StartingBlockTimestamp

	return contractMsg
}

func updateContractMsg(contractMsg *msgbus.Contract, contractValues hashrateContractValues) {
	contractMsg.Price = contractValues.Price
	contractMsg.Limit = contractValues.Limit
	contractMsg.Speed = contractValues.Speed
	contractMsg.Length = contractValues.Length
}