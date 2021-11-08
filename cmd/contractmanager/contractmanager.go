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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/ledger"
)

const (
	NewState		uint8 = 0
	ReadyState		uint8 = 1
	ActiveState		uint8 = 2
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

func New(ps *msgbus.PubSub, cmConfig map[string]interface{}) (cm *ContractManager, err error) {
	cm = &ContractManager{
		ps: ps,
		walletAddress: common.HexToAddress(cmConfig["nodeWalletAddress"].(string)),
		rpcClient: SetUpClient(cmConfig["rpcClientAddress"].(string), common.HexToAddress(cmConfig["contractManagerAccount"].(string))),
		cloneFactoryAddress: common.HexToAddress(cmConfig["cloneFactoryAddress"].(string)),
		webFacingAddress: common.HexToAddress(cmConfig["webFacingAddress"].(string)),
		ledgerAddress: common.HexToAddress(cmConfig["ledgerAddress"].(string)),
		account: common.HexToAddress(cmConfig["contractManagerAccount"].(string)),
		privateKey: cmConfig["contractManagerPrivateKey"].(string),
	}
	return cm, err
}

func SetUpClient(clientAddress string, contractManagerAccount common.Address) *ethclient.Client {
	client, err := ethclient.Dial(clientAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to rpc client at %v\n", clientAddress)

	balance, err := client.BalanceAt(context.Background(), contractManagerAccount, nil) 
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println("Balance of contract manager account:", ethValue, "ETH") 

	return client
}

func SubscribeToContractEvents(client *ethclient.Client, contractAddress common.Address) (chan types.Log, ethereum.Subscription) {
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

func ReadHashrateContract(client *ethclient.Client, contractAddress common.Address) HashrateContractValues {
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

func ReadSellerContracts(client *ethclient.Client, contractAddress common.Address, sellerAddress common.Address) []common.Address{
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

/*
	TODO: Mining pool info will be encrypted moving forward so decryption logic will need to be implemented
*/
func ReadMiningPoolInformation(client *ethclient.Client, contractAddress common.Address) MiningPoolInformation {
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

func CreateContractMsg(contractAddress common.Address, contractValues HashrateContractValues, isSeller bool) msgbus.Contract {
	convertToMsgBusState := map[uint8]msgbus.ContractState {
		NewState:		msgbus.ContNewState,
		ReadyState: 	msgbus.ContReadyState,
		ActiveState:	msgbus.ContActiveState,
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

func UpdateContractMsgMiningInfo(contractMsg *msgbus.Contract, miningPoolInfo MiningPoolInformation) {
	contractMsg.IpAddress = miningPoolInfo.IpAddress
	contractMsg.Username = miningPoolInfo.Username
	contractMsg.Port = miningPoolInfo.Port
}

func DefineThresholdParams(configFilePath string) ThresholdParams {
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

/*
	TODO: Implement similar closeout monitor that checks threshold parameters are being met
*/
func CloseOutMonitor(client *ethclient.Client,
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

	return false
}

func (cm *ContractManager) StartSeller() error {
	newSellerContractMap := make(map[msgbus.ContractID]bool)
	readySellerContractMap := make(map[msgbus.ContractID]bool)
	activeSellerContractMap := make(map[msgbus.ContractID]bool)
	completeSellerContractMap := make(map[msgbus.ContractID]bool)
	contractCreated := make(chan common.Address)
	var contractValues []HashrateContractValues
	var contractMsgs []msgbus.Contract

	sellerAddresses := ReadSellerContracts(cm.rpcClient,cm.ledgerAddress,cm.walletAddress)
	fmt.Println("Existing Seller Contracts: ", sellerAddresses)
	for i := range sellerAddresses {
		newSellerContractMap[msgbus.ContractID(sellerAddresses[i].Hex())] = true
		readySellerContractMap[msgbus.ContractID(sellerAddresses[i].Hex())] = false
		activeSellerContractMap[msgbus.ContractID(sellerAddresses[i].Hex())] = false
		completeSellerContractMap[msgbus.ContractID(sellerAddresses[i].Hex())] = false
		contractValues = append(contractValues, ReadHashrateContract(cm.rpcClient,sellerAddresses[i]))
		contractMsgs = append(contractMsgs, CreateContractMsg(sellerAddresses[i], contractValues[i], true))
		cm.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])
	}

	sellerMSG := msgbus.Seller {
		ID: msgbus.SellerID(cm.walletAddress.Hex()),
		NewContracts: newSellerContractMap,
		ReadyContracts: readySellerContractMap,
		ActiveContracts: activeSellerContractMap,
		CompleteContracts: completeSellerContractMap,
	}

	cm.ps.PubWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := SubscribeToContractEvents(cm.rpcClient, cm.cloneFactoryAddress)

	// routine for listensing to contract creation events that will update seller msg with new contracts and load new contract onto msgbus
	go func() {
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				address := common.HexToAddress(cfLog.Topics[1].Hex())
				newSellerContractMap[msgbus.ContractID(address.Hex())] = true
				readySellerContractMap[msgbus.ContractID(address.Hex())] = false
				activeSellerContractMap[msgbus.ContractID(address.Hex())] = false
				completeSellerContractMap[msgbus.ContractID(address.Hex())] = false
				sellerMSG.NewContracts = newSellerContractMap
				sellerMSG.ReadyContracts = readySellerContractMap
				sellerMSG.ActiveContracts = activeSellerContractMap
				sellerMSG.CompleteContracts = completeSellerContractMap
				fmt.Printf("Log Block Number: %d\n", cfLog.BlockNumber)
				fmt.Printf("Log Index: %d\n", cfLog.Index)
				fmt.Printf("Address of created Hashrate Contract: %s\n\n", address.Hex())
				cm.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)
				createdContractValues := ReadHashrateContract(cm.rpcClient, address)
				createdContractMsg := CreateContractMsg(address, createdContractValues, true)
				cm.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(address.Hex()), createdContractMsg)
				contractCreated<-address
			}
		}
	}()

	// routine starts routines for seller's contracts that monitors contract purchase, close, and cancel events 
	go func() {
		// start routines for existing contracts
		for addr := range newSellerContractMap {
			// subscribe to events coming off of hashrate contract
			hrLogs, hrSub := SubscribeToContractEvents(cm.rpcClient, common.HexToAddress(string(addr)))
			go hashrateContractMonitor(addr, hrLogs, hrSub, cm, newSellerContractMap, readySellerContractMap, activeSellerContractMap, completeSellerContractMap, sellerMSG) 
		}
		// monitor new contracts getting created and start routines when they are created
		for {
			addr := <-contractCreated
			hrLogs, hrSub := SubscribeToContractEvents(cm.rpcClient, addr)
			go hashrateContractMonitor(msgbus.ContractID(addr.Hex()), hrLogs, hrSub, cm, newSellerContractMap, readySellerContractMap, activeSellerContractMap, completeSellerContractMap, sellerMSG) 
		}
	}()

	return nil
}

func hashrateContractMonitor(addr msgbus.ContractID, hrLogs chan types.Log, hrSub ethereum.Subscription, cm *ContractManager, newSellerContractMap map[msgbus.ContractID]bool, 
	readySellerContractMap map[msgbus.ContractID]bool, activeSellerContractMap map[msgbus.ContractID]bool, completeSellerContractMap map[msgbus.ContractID]bool, sellerMSG msgbus.Seller) {
	runningContractAddr := make(chan msgbus.ContractID, 10)

	// create event signatures to parse out which event was being emitted from hashrate contract
	contractPurchasedSig := []byte("contractPurchased(address)")
	contractClosedSig := []byte("contractClosed()")
	contractCanceledSig := []byte("contractCanceled()")
	contractPurchasedSigHash := crypto.Keccak256Hash(contractPurchasedSig)
	contractClosedSigHash := crypto.Keccak256Hash(contractClosedSig)
	contractCanceledSigHash := crypto.Keccak256Hash(contractCanceledSig)
	
	go func() {
		for {
			select {
			case err := <-hrSub.Err():
				log.Fatal(err)
			case hLog := <-hrLogs:
				switch hLog.Topics[0].Hex() {
				case contractPurchasedSigHash.Hex():
					fmt.Printf("Address of Contract Bought: %s\n\n", addr)

					// update contract state in msgbus to ready and get mining pool info
					newSellerContractMap[addr] = false
					readySellerContractMap[addr] = true
					sellerMSG.NewContracts = newSellerContractMap
					sellerMSG.ReadyContracts = readySellerContractMap
					cm.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)
					event, err := cm.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(addr))
					if err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
					}
					if event.Err != nil {
						panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
					}
					contractMsg := event.Data.(msgbus.Contract)
					contractMsg.State = msgbus.ContReadyState
					contractMsg.Buyer = msgbus.BuyerID(string(hLog.Data))
					miningPoolInfo := ReadMiningPoolInformation(cm.rpcClient, common.HexToAddress(string(addr)))
					contractMsg.IpAddress = miningPoolInfo.IpAddress
					contractMsg.Username = miningPoolInfo.Username
					contractMsg.Port = miningPoolInfo.Port
					cm.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(addr), contractMsg)

					// contract state will not be running until funded so do not continue until it is in the running state		
					hashrateContractInstance,err := implementation.NewImplementation(common.HexToAddress(string(addr)), cm.rpcClient)
					if err != nil {
						log.Fatal(err)
					}
					for {
						contractState,err := hashrateContractInstance.ContractState(nil)
						if err != nil {
							log.Fatal(err)
						}
						if contractState == ActiveState {
							break
						}
					}

					// update contract state in msgbus to active
					readySellerContractMap[addr] = false
					activeSellerContractMap[addr] = true
					sellerMSG.ReadyContracts = readySellerContractMap
					sellerMSG.ActiveContracts = activeSellerContractMap
					cm.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)
					contractMsg.State = msgbus.ContActiveState
					cm.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(addr), contractMsg)
					runningContractAddr<-addr
				case contractClosedSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Closed", addr)
				case contractCanceledSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Cancelled", addr)
				}
			}
		}
	}()
	
	// once contract is running, closeout after length of contract has passed and updated msgbus with new state
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
			setContractCloseOut(cm.rpcClient, cm.account, cm.privateKey, common.HexToAddress(string(contractMsg.ID)))
			closedContractValues := ReadHashrateContract(cm.rpcClient, common.HexToAddress(string(contractMsg.ID)))
			closedContractMsg := CreateContractMsg(common.HexToAddress(string(contractMsg.ID)), closedContractValues, true)
			cm.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contractMsg.ID), closedContractMsg)
			activeSellerContractMap[contractMsg.ID] = false
			completeSellerContractMap[contractMsg.ID] = true
			sellerMSG.ActiveContracts = activeSellerContractMap
			sellerMSG.CompleteContracts = completeSellerContractMap
			cm.ps.SetWait(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)
		}(contractMsg)
	}
}


func (cm *ContractManager) StartBuyer() error {

	return nil
}