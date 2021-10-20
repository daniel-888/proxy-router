package contractmanager

import (
	"context"
	"fmt"
	"log"
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
} 

type ThresholdParams struct {
	MinShareAmtPerMin	int
	MinShareAvgPerHour	int
	ShareDropTolerance	int
}

type ValidatorMsg struct {
	ShareAmtPerMin		int		`json:"shareAmtPerMin"`
	ShareAvgPerHour		int		`json:"shareAvgPerHour"`
	ShareDrop			int		`json:"shareDrop"`
	HashesCompleted		int		`json:"hashesCompleted"`
	ContractFulfilled 	bool	`json:"contractFulfilled"`
}

type ValidatedParams struct {
	ShareAmtPerMin		int
	ShareAvgPerHour		int
	ShareDrop			int
	HashesCompleted		int
	ContractFulfilled 	bool
}

func SetUpClient(account common.Address, rpcClient string) *ethclient.Client {
	client, err := ethclient.Dial(rpcClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to rpc client at %v\n", rpcClient)

    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Account Balance of %v: %v\n",account.Hex(),balance) 

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

func SetContractCloseOut(client *ethclient.Client,
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

	auth := bind.NewKeyedTransactor(privateKey)
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

func CreateContractMsg (contractAddress common.Address, contractValues HashrateContractValues) msgbus.Contract {
	contractStateToString := map[ContractState]msgbus.ContractState {
		Available:	"Available",
		Active: 	"Active",
		Running:	"Running",
		Complete:	"Complete",
	}

	var ContractMsg msgbus.Contract
	ContractMsg.ID = msgbus.ContractID(contractAddress.Hex())
	ContractMsg.State = contractStateToString[contractValues.State]
	ContractMsg.Buyer = msgbus.BuyerID(contractValues.Buyer.Hex())
	ContractMsg.Price = contractValues.Price
	ContractMsg.Limit = contractValues.Limit
	ContractMsg.Speed = contractValues.Speed
	ContractMsg.Length = contractValues.Length
	ContractMsg.Port = contractValues.Port
	ContractMsg.ValidationFee = contractValues.ValidationFee
	ContractMsg.StartingBlockTimestamp = contractValues.StartingBlockTimestamp

	return ContractMsg
}

func DefineThresholdParams(configFilePath string) ThresholdParams {
	var tParams ThresholdParams
	configParams,err := configurationmanager.LoadConfiguration(configFilePath, "contractManager")
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println("here1")
	tParams.MinShareAmtPerMin = int(configParams["minShareAmtPerMin"].(float64))
	tParams.MinShareAvgPerHour = int(configParams["minShareAvgPerHour"].(float64))
	tParams.ShareDropTolerance = int(configParams["shareDropTolerance"].(float64))
	fmt.Println("here2")
	return tParams
}

func consumeValidatorMsg(contractAddress common.Address, msgs <-chan ValidatorMsg, params chan<- ValidatedParams) {
	for msg := range msgs {
		log.Printf("Contract Address: %s\n", contractAddress.Hex())
		params <- ValidatedParams{
			ShareAmtPerMin: int(msg.ShareAmtPerMin),
			ShareAvgPerHour: int(msg.ShareAvgPerHour),
			ShareDrop: int(msg.ShareDrop),
			HashesCompleted: int(msg.HashesCompleted),
			ContractFulfilled: bool(msg.ContractFulfilled),
		}
	}
}

func CloseOutMonitor(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address, 
	msgs chan ValidatorMsg, 
	thresholds ThresholdParams) {
	params := make(chan ValidatedParams)
	go consumeValidatorMsg(contractAddress, msgs, params)

	param := <-params
	if param.ShareAmtPerMin < thresholds.MinShareAmtPerMin || param.ShareAvgPerHour < thresholds.MinShareAvgPerHour || param.ShareDrop > thresholds.ShareDropTolerance {
		log.Printf("Closing out contract %s for not meeting threshold requirements\n", contractAddress.Hex())
		SetContractCloseOut(client,fromAddress,privateKeyString,contractAddress)
		return
	}
	if param.ContractFulfilled {
		log.Printf("Closing out contract %s for fulfilling requirements\n", contractAddress.Hex())
		SetContractCloseOut(client,fromAddress,privateKeyString,contractAddress)
		return
	}

	log.Printf("Contract at %s has %d hashes completed\n", contractAddress.Hex(), param.HashesCompleted)
}