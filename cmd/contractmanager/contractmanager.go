package contractmanager

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
)

type ContractState uint8

const (
	Available	ContractState = 0
	Active		ContractState = 1
	Running		ContractState = 2
	Complete	ContractState = 3
)

type HashrateContractValues struct {
	State					ContractState
	Price 					int
	Limit 					int
	Speed 					int	
	Length 					int
	Port 					int
	ValidationFee			int
	StartingBlockTimestamp	int
	Buyer 					common.Address
	Seller 					common.Address
	IpAddress				string
	Username				string
	Password				string
} 

type MiningPoolInformation struct {
	IpAddress	string 
	Username	string
	Password	string
}

type ThresholdParams struct {
	MinShareAmtPerMin	int
	MinShareAvgPerHour	int
	ShareDropTolerance	int
}

type ContractManager struct {
	ps 					*msgbus.PubSub
	rpcClient 			*ethclient.Client
	cloneFactoryAddress	common.Address
	webFacingAddress	common.Address
	account				common.Address
	privateKey			string
}

func New(ps *msgbus.PubSub, cmConfig map[string]interface{}) (cm *ContractManager, err error) {
	cm = &ContractManager{
		ps: ps,
		rpcClient: SetUpClient("../configurationmanager/testconfig.json"),
		cloneFactoryAddress: common.HexToAddress(cmConfig["cloneFactoryAddress"].(string)),
		webFacingAddress: common.HexToAddress(cmConfig["webFacingAddress"].(string)),
		account: common.HexToAddress(cmConfig["contractManagerAccount"].(string)),
		privateKey: cmConfig["contractManagerPrivateKey"].(string),
	}
	return cm, err
}

func SetUpClient(configPath string) *ethclient.Client {
	configaData, err := configurationmanager.LoadConfiguration(configPath, "contractManager")
	if err != nil {
		log.Fatal(err)
	}
	clientAddress := configaData["rpcClientAddress"].(string)
	contractManagerAccount := common.HexToAddress(configaData["contractManagerAccount"].(string))

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
	contractValues.State = ContractState(state)

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

	port, err := instance.Port(nil)
	if err != nil {
        log.Fatal(err)
    }
	contractValues.Port = int(port.Int64())	

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

/*
	TODO: Mining pool info will be encrypted moving forward so decryption logic will need to be implemented
*/
func ReadMiningPoolInformation(client *ethclient.Client, fromAddress common.Address, contractAddress common.Address) MiningPoolInformation {
	var auth *bind.CallOpts = new(bind.CallOpts)
	auth.From = fromAddress
	BlockNumber,err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.BlockNumber = big.NewInt(int64(BlockNumber))

	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Address %s getting mining pool info from contract %s\n\n", fromAddress, contractAddress)

	ipaddress,username,password,err := instance.GetMiningPoolInformation(auth)
	if err != nil {
		log.Fatal(err)
	}

	miningPoolInfo := MiningPoolInformation{
		IpAddress: ipaddress,
		Username: username,
		Password: password,
	}

	return miningPoolInfo
}

func setContractCloseOut(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
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
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.SetContractCloseOut(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
}

func CreateContractMsg(contractAddress common.Address, contractValues HashrateContractValues) msgbus.Contract {
	contractStateToString := map[ContractState]msgbus.ContractState {
		Available:	"Available",
		Active: 	"Active",
		Running:	"Running",
		Complete:	"Complete",
	}

	var contractMsg msgbus.Contract
	contractMsg.ID = msgbus.ContractID(contractAddress.Hex())
	contractMsg.State = contractStateToString[contractValues.State]
	contractMsg.Buyer = msgbus.BuyerID(contractValues.Buyer.Hex())
	contractMsg.Price = contractValues.Price
	contractMsg.Limit = contractValues.Limit
	contractMsg.Speed = contractValues.Speed
	contractMsg.Length = contractValues.Length
	contractMsg.Port = contractValues.Port
	contractMsg.ValidationFee = contractValues.ValidationFee
	contractMsg.StartingBlockTimestamp = contractValues.StartingBlockTimestamp

	return contractMsg
}

func UpdateContractMsgMiningInfo(contractMsg *msgbus.Contract, miningPoolInfo MiningPoolInformation) {
	contractMsg.IpAddress = miningPoolInfo.IpAddress
	contractMsg.Username = miningPoolInfo.Username
	contractMsg.Password = miningPoolInfo.Password
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
	ech := make(msgbus.EventChan)
	hashRate := make(chan int)
	ps.Sub(msgbus.MinerMsg, minerID, ech)
	ps.Get(msgbus.MinerMsg, minerID, ech)

	go readMinerHashrate(ech, hashRate)

	currentHashRate := <-hashRate
	if currentHashRate < contractMsg.Speed {
		log.Printf("Closing out contract %s for not meeting hashrate requirements\n", contractAddress.Hex())
		setContractCloseOut(client,fromAddress,privateKeyString,contractAddress)
		return true
	}

	log.Println("Contract hashrate is being fulfilled")

	return false
}

func readMinerHashrate(ech msgbus.EventChan, currentHashRate chan int){
	for e := range ech {
		if e.EventType == msgbus.GetEvent {
			minerInfo := e.Data.(msgbus.Miner)
			currentHashRate <- minerInfo.CurrentHashRate
		}
	}
}

func (cm *ContractManager) Start() error {
	contractCreatedAddr := make(chan common.Address)
	contractPurchasedAddr := make(chan common.Address)
	updateContractAddr := make(chan common.Address)
	closeoutContractAddr := make(chan common.Address)
	stop := make(chan bool)

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := SubscribeToContractEvents(cm.rpcClient, cm.cloneFactoryAddress)

	// routine for listensing to contract creation events 
	go func() {
		loop:
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				address := common.HexToAddress(cfLog.Topics[1].Hex())
				contractCreatedAddr<-address
				fmt.Printf("Log Block Number: %d\n", cfLog.BlockNumber)
				fmt.Printf("Log Index: %d\n", cfLog.Index)
				fmt.Printf("Address of created Hashrate Contract: %s\n\n", address.Hex())
			case <-stop:
				break loop //stop listening to events
			}
		}
		close(cfLogs)
		cfSub.Unsubscribe()
	}()

	// routine for listensing to events from hashrate contract
	go func() {
		// create event signatures to parse out which event was being emitted
		contractPurchasedSig := []byte("contractPurchased(address)")
		contractClosedSig := []byte("contractClosed()")
		contractCanceledSig := []byte("contractCanceled()")
		contractPurchasedSigHash := crypto.Keccak256Hash(contractPurchasedSig)
		contractClosedSigHash := crypto.Keccak256Hash(contractClosedSig)
		contractCanceledSigHash := crypto.Keccak256Hash(contractCanceledSig)

		// subcribe to events emitted by hashrate contract
		hashrateContractAddress := <-contractCreatedAddr
		hrLogs, hrSub := SubscribeToContractEvents(cm.rpcClient, hashrateContractAddress)
		//var address common.Address
		loop:
		for {
			select {
			case err := <-hrSub.Err():
				log.Fatal(err)
			case hLog := <-hrLogs:
				switch hLog.Topics[0].Hex() {
				case contractPurchasedSigHash.Hex():
					address := common.HexToAddress(string(hLog.Data))
					contractPurchasedAddr<-hashrateContractAddress
					fmt.Printf("Address of Contract Bought: %s\n\n", address.Hex())
				case contractClosedSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Closed", hashrateContractAddress)
				case contractCanceledSigHash.Hex():
					fmt.Printf("Hashrate Contract %s Cancelled", hashrateContractAddress)
				}
			case <-stop:
				break loop //stop listening to events
			}
		}
		close(hrLogs)
		hrSub.Unsubscribe()	
	}()

	// read values from created Hashrate contracts and publish them to msgbus
	go func() {
		loop:
		for {
			select {
			case address := <-contractPurchasedAddr:
				contractValues := ReadHashrateContract(cm.rpcClient, address)
				
				// push read in contract values into message bus contract struct
				contractMsg := CreateContractMsg(address, contractValues)
				event, err := cm.ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contractMsg.ID), contractMsg)
				if err != nil {
					panic(fmt.Sprintf("Adding New Contract Failed: %s", err))
				}
				if event.Err != nil {
					panic(fmt.Sprintf("Adding New Contract Failed: %s", event.Err))
				}
				updateContractAddr<-address

			case <-stop:
				break loop
			}
		}
	}()

	// update contract in msgbus with mining pool info upon purchase
	go func() {
		loop:
		for {
			select {
			case address := <-updateContractAddr:
				miningPoolInfo := ReadMiningPoolInformation(cm.rpcClient, cm.account, address)
				
				event,err := cm.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(address.Hex()))
				if err != nil {
					panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
				}
				if event.Err != nil {
					panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
				}
				
				contractMsg := event.Data.(msgbus.Contract)
				UpdateContractMsgMiningInfo(&contractMsg, miningPoolInfo)

				event,err = cm.ps.SetWait(msgbus.ContractMsg, msgbus.IDString(address.Hex()), contractMsg)
				if err != nil {
					panic(fmt.Sprintf("Updating Purchased Contract Failed: %s", err))
				}
				if event.Err != nil {
					panic(fmt.Sprintf("Updating Purchased Contract Failed: %s", event.Err))
				}
				closeoutContractAddr<-address
			case <-stop:
				break loop
			}
		}
	}()

	// run closeout monitor for purchased contract
	go func() {
		loop:
		for {
			select{
			case address := <-closeoutContractAddr:
				go func() {
					loop2:
					for {
						event,err := cm.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(address.Hex()))
						if err != nil {
							panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
						}
						if event.Err != nil {
							panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
						}

						contractMsg := event.Data.(msgbus.Contract)
						// find Miner associated with IP Address set in contract
						event, err = cm.ps.SearchIPWait(msgbus.MinerMsg, contractMsg.IpAddress)
						if err != nil {
							panic(fmt.Sprintf("Searching For Miner Specified In Contract Failed: %s", err))
						}
						if event.Err != nil {
							panic(fmt.Sprintf("Searching For Miner Specified In Contract Failed: %s", event.Err))
						}
						searchedMiner := event.Data.(msgbus.IDIndex)
						closed := CloseOutMonitor(cm.rpcClient,cm.account,cm.privateKey,address,searchedMiner[0],contractMsg,cm.ps)
						if closed {
							break loop2
						}
					}
				}()
			case <-stop:
				break loop
			}
		}
	}()

	return nil
}