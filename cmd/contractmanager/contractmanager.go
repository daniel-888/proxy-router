package contractmanager

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"

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
	State			ContractState
	Price 			int
	Limit 			int
	Speed 			int	
	Length 			int
	Port 			int
	ValidationFee	int
	Buyer 			common.Address
	Seller 			common.Address
	IpAddr			string
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

	ipaddr, err := instance.Ipaddress(nil)
	if err != nil {
        log.Fatal(err)
    }
	contractValues.IpAddr = ipaddr

	return contractValues
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
	ContractMsg.Dest = msgbus.DestID(contractValues.IpAddr)
	ContractMsg.Price = contractValues.Price
	ContractMsg.Limit = contractValues.Limit
	ContractMsg.Speed = contractValues.Speed
	ContractMsg.Length = contractValues.Length
	ContractMsg.Port = contractValues.Port
	ContractMsg.ValidationFee = contractValues.ValidationFee

	return ContractMsg
}