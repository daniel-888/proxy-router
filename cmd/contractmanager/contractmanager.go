package contractmanager

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/ledger"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/webfacing"
)

type ContractManager struct {
	ps 					*msgbus.PubSub
	rpcClient 			*ethclient.Client
	walletAddress		common.Address
	cloneFactoryAddress	common.Address
	webFacingAddress	common.Address
	ledgerAddress		common.Address
	account				common.Address
	privateKey			string
}

const (
	AvailableState	uint8 = 0
	ActiveState		uint8 = 1
	RunningState	uint8 = 2
	CompleteState	uint8 = 3
)

type HashrateContractValues struct {
	State					uint8
	Price 					int
	Limit 					int
	Speed 					int	
	Length 					int
	ValidationFee			int
	StartingBlockTimestamp	int
	Buyer 					common.Address
	Seller 					common.Address
} 

type MiningPoolInformation struct {
	IpAddress	string 
	Port 		string
	Username	string
}

type ThresholdParams struct {
	MinShareAmtPerMin	int
	MinShareAvgPerHour	int
	ShareDropTolerance	int
}

func New(ps *msgbus.PubSub, cmConfig map[string]interface{}) (cm *ContractManager, err error) {
	var client *ethclient.Client
	client, err = setUpClient(cmConfig["rpcClientAddress"].(string), common.HexToAddress(cmConfig["contractManagerAccount"].(string)))
	if err != nil {
		log.Fatal(err)
	}
	cm = &ContractManager{
		ps: ps,
		walletAddress: common.HexToAddress(cmConfig["nodeWalletAddress"].(string)),
		rpcClient: client,
		cloneFactoryAddress: common.HexToAddress(cmConfig["cloneFactoryAddress"].(string)),
		webFacingAddress: common.HexToAddress(cmConfig["webFacingAddress"].(string)),
		ledgerAddress: common.HexToAddress(cmConfig["ledgerAddress"].(string)),
		account: common.HexToAddress(cmConfig["contractManagerAccount"].(string)),
		privateKey: cmConfig["contractManagerPrivateKey"].(string),
	}
	return cm, err
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

func readHashrateContract(client *ethclient.Client, contractAddress common.Address) HashrateContractValues {
	instance, err := implementation.NewImplementation(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

	var contractValues HashrateContractValues

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

func readSellerContracts(client *ethclient.Client, contractAddress common.Address, sellerAddress common.Address) []common.Address{
	var sellerContractAddresses []common.Address
	var hashrateContractInstance *implementation.Implementation
	var hashrateContractSeller common.Address

	instance, err := ledger.NewLedger(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

	hashrateContractAddresses, err := instance.GetListOfContractsLedger(&bind.CallOpts{})
	if err != nil {
        log.Fatal(err)
    }

	// parse existing hashrate contracts for ones that belong to seller
	for i := range hashrateContractAddresses {
		hashrateContractInstance, err = implementation.NewImplementation(hashrateContractAddresses[i], client)
		if err != nil {
			log.Fatal(err)
		}
		hashrateContractSeller, err = hashrateContractInstance.Seller(nil)
		if err != nil {
			log.Fatal(err)
		}
		if hashrateContractSeller == sellerAddress {
			sellerContractAddresses = append(sellerContractAddresses, hashrateContractAddresses[i])
		}
	}

	return sellerContractAddresses
}

func readBuyerContracts(client *ethclient.Client, contractAddress common.Address, buyerAddress common.Address) []common.Address{
	var buyerContractAddresses []common.Address
	var hashrateContractInstance *implementation.Implementation
	var hashrateContractBuyer common.Address

	instance, err := ledger.NewLedger(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

	hashrateContractAddresses, err := instance.GetListOfContractsLedger(&bind.CallOpts{})
	if err != nil {
        log.Fatal(err)
    }

	// parse existing hashrate contracts for ones that belong to buyer
	for i := range hashrateContractAddresses {
		hashrateContractInstance, err = implementation.NewImplementation(hashrateContractAddresses[i], client)
		if err != nil {
			log.Fatal(err)
		}
		hashrateContractBuyer, err = hashrateContractInstance.Buyer(nil)
		if err != nil {
			log.Fatal(err)
		}
		if hashrateContractBuyer == buyerAddress {
			buyerContractAddresses = append(buyerContractAddresses, hashrateContractAddresses[i])
		}
	}

	return buyerContractAddresses
}

/*
	TODO: Mining pool info will be encrypted moving forward so decryption logic will need to be implemented
*/
func readMiningPoolInformation(client *ethclient.Client, contractAddress common.Address) MiningPoolInformation {
	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Getting mining pool info from contract %s\n\n", contractAddress)

	poolData, err := instance.EncryptedPoolData(nil)
	if err != nil {
        log.Fatal(err)
    }
	poolDataSplit := strings.Split(poolData, "|")

	miningPoolInfo := MiningPoolInformation{
		IpAddress: poolDataSplit[0],
		Port: poolDataSplit[1],
		Username: poolDataSplit[2],
	}

	return miningPoolInfo
}

func setContractCloseOut(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address) {
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

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	tx, err := instance.SetContractCloseOut(auth)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
	fmt.Println("Closing Out Contract: ", contractAddress)
}

func createContractMsg(contractAddress common.Address, contractValues HashrateContractValues, isSeller bool) msgbus.Contract {
	convertToMsgBusState := map[uint8]msgbus.ContractState {
		AvailableState:	msgbus.ContAvailableState,
		ActiveState: 	msgbus.ContActiveState,
		RunningState:	msgbus.ContRunningState,
		CompleteState:	msgbus.ContCompleteState,
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

func updateContractMsgMiningInfo(contractMsg *msgbus.Contract, miningPoolInfo MiningPoolInformation) {
	contractMsg.IpAddress = miningPoolInfo.IpAddress
	contractMsg.Username = miningPoolInfo.Username
	contractMsg.Port = miningPoolInfo.Port
}

func defineThresholdParams(configFilePath string) ThresholdParams {
	var tParams ThresholdParams
	configParams,err := configurationmanager.LoadConfiguration(configFilePath, "contractManager")
	if err != nil {
        log.Fatal(err)
    }

	tParams.MinShareAmtPerMin = int(configParams["minShareAmtPerMin"].(float64))
	tParams.MinShareAvgPerHour = int(configParams["minShareAvgPerHour"].(float64))
	tParams.ShareDropTolerance = int(configParams["shareDropTolerance"].(float64))

	return tParams
}

func hashrateContractMonitor(addr msgbus.ContractID, hrLogs chan types.Log, hrSub ethereum.Subscription, cm *ContractManager, availableSellerContractsMap map[msgbus.ContractID]bool, 
	activeSellerContractsMap map[msgbus.ContractID]bool, runningSellerContractsMap map[msgbus.ContractID]bool, completeSellerContractsMap map[msgbus.ContractID]bool, sellerMSG msgbus.Seller) {
	runningContractAddr := make(chan msgbus.ContractID, 10)

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

					// update contract state in msgbus to active and get mining pool info
					availableSellerContractsMap[addr] = false
					activeSellerContractsMap[addr] = true
					sellerMSG.AvailableContracts = availableSellerContractsMap
					sellerMSG.ActiveContracts = activeSellerContractsMap
					cm.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)
					event, err := cm.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
					if err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
					}
					if event.Err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
					}
					contractMsg := event.Data.(msgbus.Contract)
					contractMsg.State = msgbus.ContActiveState
					contractMsg.Buyer = msgbus.BuyerID(purchasedEvent.Buyer.Hex())
					miningPoolInfo := readMiningPoolInformation(cm.rpcClient, common.HexToAddress(string(addr)))
					updateContractMsgMiningInfo(&contractMsg, miningPoolInfo)
					cm.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(addr), contractMsg)
					
				case contractFundedSigHash.Hex():
					fmt.Printf("Address of funded Hashrate Contract : %s\n\n", addr)
		
					// update contract state in msgbus to running and broadcast to closeout routine that contract is running
					activeSellerContractsMap[addr] = false
					runningSellerContractsMap[addr] = true
					sellerMSG.ActiveContracts = activeSellerContractsMap
					sellerMSG.RunningContracts = runningSellerContractsMap
					cm.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)
					event, err := cm.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
					if err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
					}
					if event.Err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
					}
					contractMsg := event.Data.(msgbus.Contract)
					contractMsg.State = msgbus.ContRunningState
					cm.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(addr), contractMsg)
					runningContractAddr<-addr

				case contractClosedSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Closed \n\n", addr)
					closedContractValues := readHashrateContract(cm.rpcClient, common.HexToAddress(string(addr)))
					closedContractMsg := createContractMsg(common.HexToAddress(string(addr)), closedContractValues, true)
					cm.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(closedContractMsg.ID), closedContractMsg)
					runningSellerContractsMap[closedContractMsg.ID] = false
					completeSellerContractsMap[closedContractMsg.ID] = true
					sellerMSG.RunningContracts = runningSellerContractsMap
					sellerMSG.CompleteContracts = completeSellerContractsMap
					cm.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)
				}
			}
		}
	}()
	
	// once contract is running, closeout after length of contract has passed if it was not closed out early
	for {
		address := <-runningContractAddr
		event,err := cm.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(address))
		if err != nil {
			panic(fmt.Sprintf("Getting Running Contract Failed: %s", err))
		}
		if event.Err != nil {
			panic(fmt.Sprintf("Getting Running Contract Failed: %s", event.Err))
		}
		contractMsg := event.Data.(msgbus.Contract)
		go func(contractMsg msgbus.Contract) {
			contractLength := contractMsg.Length
			time.Sleep(time.Second*time.Duration(contractLength))
			// if contract was not already closed early, close out here
			contractValues := readHashrateContract(cm.rpcClient, common.HexToAddress(string(address)))
			if contractValues.State == RunningState {
				setContractCloseOut(cm.rpcClient, cm.account, cm.privateKey, common.HexToAddress(string(contractMsg.ID)))
			}
		}(contractMsg)
	}
}

/*
	TODO: Implement similar closeout monitor that checks threshold parameters are being met
*/
func closeOutMonitor(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address, 
	minerID msgbus.IDString,
	contractMsg msgbus.Contract,
	ps *msgbus.PubSub) bool	{

	event,_ := ps.GetWait(msgbus.MinerMsg, minerID)
	minerInfo := event.Data.(msgbus.Miner)

	currentHashRate := minerInfo.CurrentHashRate
	if currentHashRate < contractMsg.Speed {
		log.Printf("Closing out contract %s for not meeting hashrate requirements\n", contractAddress.Hex())
		setContractCloseOut(client,fromAddress,privateKeyString,contractAddress)
		return true
	}

	log.Println("Contract hashrate is being fulfilled")

	time.Sleep(time.Millisecond*5000)
	return false
}

func (cm *ContractManager) StartSeller() error {
	availableSellerContractsMap := make(map[msgbus.ContractID]bool)
	activeSellerContractsMap := make(map[msgbus.ContractID]bool)
	runningSellerContractsMap := make(map[msgbus.ContractID]bool)
	completeSellerContractsMap := make(map[msgbus.ContractID]bool)
	contractCreated := make(chan common.Address)
	var contractValues []HashrateContractValues
	var contractMsgs []msgbus.Contract

	sellerContracts := readSellerContracts(cm.rpcClient,cm.ledgerAddress,cm.walletAddress)
	fmt.Println("Existing Seller Contracts: ", sellerContracts)
	for i := range sellerContracts {
		availableSellerContractsMap[msgbus.ContractID(sellerContracts[i].Hex())] = true
		activeSellerContractsMap[msgbus.ContractID(sellerContracts[i].Hex())] = false
		runningSellerContractsMap[msgbus.ContractID(sellerContracts[i].Hex())] = false
		completeSellerContractsMap[msgbus.ContractID(sellerContracts[i].Hex())] = false
		contractValues = append(contractValues, readHashrateContract(cm.rpcClient,sellerContracts[i]))
		contractMsgs = append(contractMsgs, createContractMsg(sellerContracts[i], contractValues[i], true))
		cm.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])
	}

	sellerMSG := msgbus.Seller {
		ID: msgbus.SellerID(cm.walletAddress.Hex()),
		AvailableContracts:	availableSellerContractsMap,
		ActiveContracts: 	activeSellerContractsMap,
		RunningContracts: 	activeSellerContractsMap,
		CompleteContracts: 	completeSellerContractsMap,
	}

	cm.ps.PubWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := subscribeToContractEvents(cm.rpcClient, cm.cloneFactoryAddress)

	// routine for listensing to contract creation events that will update seller msg with new contracts and load new contract onto msgbus
	go func() {
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				address := common.HexToAddress(cfLog.Topics[1].Hex())
				// check if contract created belongs to seller
				hashrateContractInstance, err := implementation.NewImplementation(address, cm.rpcClient)
				if err != nil {
					log.Fatal(err)
				}
				hashrateContractSeller, err := hashrateContractInstance.Seller(nil)
				if err != nil {
					log.Fatal(err)
				}
				if hashrateContractSeller == cm.walletAddress {
					availableSellerContractsMap[msgbus.ContractID(address.Hex())] = true
					activeSellerContractsMap[msgbus.ContractID(address.Hex())] = false
					runningSellerContractsMap[msgbus.ContractID(address.Hex())] = false
					completeSellerContractsMap[msgbus.ContractID(address.Hex())] = false
					sellerMSG.ActiveContracts = availableSellerContractsMap
					sellerMSG.ActiveContracts = activeSellerContractsMap
					sellerMSG.RunningContracts = runningSellerContractsMap
					sellerMSG.CompleteContracts = completeSellerContractsMap
					fmt.Printf("Log Block Number: %d\n", cfLog.BlockNumber)
					fmt.Printf("Log Index: %d\n", cfLog.Index)
					fmt.Printf("Address of created Hashrate Contract: %s\n\n", address.Hex())
					cm.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)
					createdContractValues := readHashrateContract(cm.rpcClient, address)
					createdContractMsg := createContractMsg(address, createdContractValues, true)
					cm.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(address.Hex()), createdContractMsg)
					contractCreated<-address
				}
			}
		}
	}()

	// routine starts routines for seller's contracts that monitors contract purchase, close, and cancel events 
	go func() {
		// start routines for existing contracts
		for addr := range availableSellerContractsMap {
			// subscribe to events coming off of hashrate contract
			hrLogs, hrSub := subscribeToContractEvents(cm.rpcClient, common.HexToAddress(string(addr)))
			go hashrateContractMonitor(addr, hrLogs, hrSub, cm, availableSellerContractsMap, activeSellerContractsMap, runningSellerContractsMap, completeSellerContractsMap, sellerMSG) 
		}
		// monitor new contracts getting created and start routines when they are created
		for {
			addr := <-contractCreated
			hrLogs, hrSub := subscribeToContractEvents(cm.rpcClient, addr)
			go hashrateContractMonitor(msgbus.ContractID(addr.Hex()), hrLogs, hrSub, cm, availableSellerContractsMap, activeSellerContractsMap, runningSellerContractsMap, completeSellerContractsMap, sellerMSG) 
		}
	}()

	return nil
}

func (cm *ContractManager) StartBuyer() error {
	activeBuyerContractsMap := make(map[msgbus.ContractID]bool)
	runningBuyerContractsMap := make(map[msgbus.ContractID]bool)
	completeBuyerContractsMap := make(map[msgbus.ContractID]bool)
	purchasedContractAddr := make(chan common.Address)
	runningContractAddr := make(chan msgbus.ContractID, 10)
	var contractValues []HashrateContractValues
	var contractMsgs []msgbus.Contract
	var runningContracts []common.Address

	buyerContracts := readBuyerContracts(cm.rpcClient, cm.ledgerAddress, cm.walletAddress)
	fmt.Println("Existing Buyer Contracts: ", buyerContracts)
	for i := range buyerContracts {
		contractValues = append(contractValues, readHashrateContract(cm.rpcClient,buyerContracts[i]))
		contractMsgs = append(contractMsgs, createContractMsg(buyerContracts[i], contractValues[i], false))
		cm.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])

		switch contractValues[i].State {
		case ActiveState:
			activeBuyerContractsMap[msgbus.ContractID(buyerContracts[i].Hex())] = true
			runningBuyerContractsMap[msgbus.ContractID(buyerContracts[i].Hex())] = false
			completeBuyerContractsMap[msgbus.ContractID(buyerContracts[i].Hex())] = false
		case RunningState:
			activeBuyerContractsMap[msgbus.ContractID(buyerContracts[i].Hex())] = false
			runningBuyerContractsMap[msgbus.ContractID(buyerContracts[i].Hex())] = true
			completeBuyerContractsMap[msgbus.ContractID(buyerContracts[i].Hex())] = false
			runningContracts = append(runningContracts, buyerContracts[i])
		case CompleteState:
			activeBuyerContractsMap[msgbus.ContractID(buyerContracts[i].Hex())] = false
			runningBuyerContractsMap[msgbus.ContractID(buyerContracts[i].Hex())] = false
			completeBuyerContractsMap[msgbus.ContractID(buyerContracts[i].Hex())] = true
		}
	}

	buyerMSG := msgbus.Buyer {
		ID: msgbus.BuyerID(cm.walletAddress.Hex()),
		ActiveContracts: 	activeBuyerContractsMap,
		RunningContracts: 	runningBuyerContractsMap,
		CompleteContracts: 	completeBuyerContractsMap,
	}

	cm.ps.PubWait(msgbus.BuyerMsg, msgbus.IDString(buyerMSG.ID), buyerMSG)

	// monitor hashrate coming from miner associated with running contracts found at startup
	for i := range runningContracts {
		// get contract msg
		event1, err := cm.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(runningContracts[i].Hex()))
		if err != nil {
			panic(fmt.Sprintf("Getting Running Contract Failed: %s", err))
		}
		if event1.Err != nil {
			panic(fmt.Sprintf("Getting Running Contract Failed: %s", err))
		}
		contractMsg := event1.Data.(msgbus.Contract)

		// get miner based on ip address on contract
		miningPoolInfo := readMiningPoolInformation(cm.rpcClient, runningContracts[i])
		event2, err := cm.ps.SearchIPWait(msgbus.MinerMsg, miningPoolInfo.IpAddress)
		if err != nil {
			panic(fmt.Sprintf("Search for miner with IP Address %s Failed: %s", miningPoolInfo.IpAddress, err))
		}
		if event2.Err != nil {
			panic(fmt.Sprintf("Search for miner with IP Address %s Failed: %s", miningPoolInfo.IpAddress, err))
		}
		minerID := event2.Data.(msgbus.IDIndex)

		go func() {
			for {
				isClosed := closeOutMonitor(cm.rpcClient, cm.account, cm.privateKey, common.HexToAddress(string(contractMsg.ID)), minerID[0], contractMsg, cm.ps)
				contractValues := readHashrateContract(cm.rpcClient, common.HexToAddress(string(contractMsg.ID)))
				if contractValues.State == CompleteState || isClosed {
					runningBuyerContractsMap[contractMsg.ID] = false
					completeBuyerContractsMap[contractMsg.ID] = true
					buyerMSG.RunningContracts = runningBuyerContractsMap
					buyerMSG.CompleteContracts = completeBuyerContractsMap
					cm.ps.SetWait(msgbus.BuyerMsg, msgbus.IDString(buyerMSG.ID), buyerMSG)

					closedContractMsg := createContractMsg(common.HexToAddress(string(contractMsg.ID)), contractValues, false)
					cm.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(closedContractMsg.ID), closedContractMsg)
					break
				}
			}
		}()
	}	

	// subcribe to events emitted by webfacing contract to read contract purchase event
	wfLogs, wfSub := subscribeToContractEvents(cm.rpcClient, cm.webFacingAddress)
	// to decode event data
	webFacingAbi, err := abi.JSON(strings.NewReader(string(webfacing.WebfacingABI)))
	if err != nil {
        log.Fatal(err)
    }
	purchasedEvent := struct {
		Contract common.Address
	}{}

	// routine for listensing to contract purchase events to update buyer with new contracts they purchased
	go func() {
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
				fmt.Printf("Address of purchased Hashrate Contract : %s\n\n", contractAddress.Hex())
				contractValues := readHashrateContract(cm.rpcClient, contractAddress)
				if contractValues.Buyer == cm.walletAddress {
					fmt.Printf("Address of purchased Hashrate Contract : %s\n\n", contractAddress.Hex())
					contractMsg := createContractMsg(contractAddress, contractValues, false)
					miningPoolInfo := readMiningPoolInformation(cm.rpcClient, contractAddress)
					updateContractMsgMiningInfo(&contractMsg, miningPoolInfo)
					cm.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsg.ID), contractMsg)
					
					activeBuyerContractsMap[msgbus.ContractID(contractAddress.Hex())] = true
					runningBuyerContractsMap[msgbus.ContractID(contractAddress.Hex())] = false
					completeBuyerContractsMap[msgbus.ContractID(contractAddress.Hex())] = false
					buyerMSG.ActiveContracts = activeBuyerContractsMap
					buyerMSG.RunningContracts = runningBuyerContractsMap
					buyerMSG.CompleteContracts = completeBuyerContractsMap
					cm.ps.SetWait(msgbus.BuyerMsg, msgbus.IDString(buyerMSG.ID), buyerMSG)
					purchasedContractAddr<-contractAddress

					
				}
			}
		}
	}()

	// routine listens to purchased hashrate contract until it is funded to update it to running state
	go func () {
		for {
			address := <-purchasedContractAddr
			// get contract msg
			event, err := cm.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(address.Hex()))
			if err != nil {
				panic(fmt.Sprintf("Getting Running Contract Failed: %s", err))
			}
			if event.Err != nil {
				panic(fmt.Sprintf("Getting Running Contract Failed: %s", err))
			}
			contractMsg := event.Data.(msgbus.Contract)

			// subcribe to events emitted by hashrate contract to be notified when contract is funded
			hrLogs, hrSub := subscribeToContractEvents(cm.rpcClient, address)
	
			// create event signature to parse out contractFunded event
			contractFundedSig := []byte("contractFunded(address)")
			contractFundedSigHash := crypto.Keccak256Hash(contractFundedSig)
	
			go func() {
				for {
					select {
					case err := <-hrSub.Err():
						log.Fatal(err)
					case hrLog := <-hrLogs:
						if hrLog.Topics[0] == contractFundedSigHash {
							// update contract state in msgbus to active
							activeBuyerContractsMap[msgbus.ContractID(address.Hex())] = false
							runningBuyerContractsMap[msgbus.ContractID(address.Hex())] = true
							buyerMSG.ActiveContracts = activeBuyerContractsMap
							buyerMSG.RunningContracts = runningBuyerContractsMap
							cm.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(buyerMSG.ID), buyerMSG)
							contractMsg.State = msgbus.ContRunningState
							cm.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsg.ID), contractMsg)
							runningContractAddr<-msgbus.ContractID(address.Hex())
						}
					}
				}
			}()
		}
	}()

	// routine to monitor hashrate coming from miner associated with new running contracts
	go func() {
		for {
			address := <-runningContractAddr
			// get contract msg
			event1, err := cm.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(address))
			if err != nil {
				panic(fmt.Sprintf("Getting Running Contract Failed: %s", err))
			}
			if event1.Err != nil {
				panic(fmt.Sprintf("Getting Running Contract Failed: %s", err))
			}
			contractMsg := event1.Data.(msgbus.Contract)
	
			// get miner based on ip address on contract
			miningPoolInfo := readMiningPoolInformation(cm.rpcClient, common.HexToAddress(string(address)))
			event2, err := cm.ps.SearchIPWait(msgbus.MinerMsg, miningPoolInfo.IpAddress)
			if err != nil {
				panic(fmt.Sprintf("Search for miner with IP Address %s Failed: %s", miningPoolInfo.IpAddress, err))
			}
			if event2.Err != nil {
				panic(fmt.Sprintf("Search for miner with IP Address %s Failed: %s", miningPoolInfo.IpAddress, err))
			}
			minerID := event2.Data.(msgbus.IDIndex)
	
			go func() {
				for {
					isClosed := closeOutMonitor(cm.rpcClient, cm.account, cm.privateKey, common.HexToAddress(string(contractMsg.ID)), minerID[0], contractMsg, cm.ps)
					contractValues := readHashrateContract(cm.rpcClient, common.HexToAddress(string(contractMsg.ID)))
					if contractValues.State == CompleteState || isClosed {
						runningBuyerContractsMap[contractMsg.ID] = false
						completeBuyerContractsMap[contractMsg.ID] = true
						buyerMSG.RunningContracts = runningBuyerContractsMap
						buyerMSG.CompleteContracts = completeBuyerContractsMap
						cm.ps.SetWait(msgbus.BuyerMsg, msgbus.IDString(buyerMSG.ID), buyerMSG)
	
						closedContractMsg := createContractMsg(common.HexToAddress(string(contractMsg.ID)), contractValues, false)
						cm.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(closedContractMsg.ID), closedContractMsg)
						break
					}
				}
			}()
		}
	}()

	return nil
}