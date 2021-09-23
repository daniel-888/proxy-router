package contractmanager

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
	// "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	// "gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts"
)
func BoilerPlateFunc() (string, error) {
	msg := "Contract Manager Package"
	return lumerinlib.BoilerPlateLibFunc(msg), nil // always returns no error
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

// func SubscribeToContractCreatedEvent(client *ethclient.Client, contractAddress common.Address) {
//     query := ethereum.FilterQuery{
//         Addresses: []common.Address{contractAddress},
//     }

//     logs := make(chan types.Log)
//     sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
//     if err != nil {
//         log.Fatal(err)
//     }

// 	for {	
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal(err)
// 		case vLog := <-logs:
// 			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
//         	fmt.Printf("Log Index: %d\n", vLog.Index)

// 			hashrateContractAddress := common.HexToAddress(vLog.Topics[1].Hex())
// 			fmt.Printf("Address of created Hashrate Contract: %s\n", hashrateContractAddress.Hex())
// 		}
// 	}
// }

// func ReadHashrateContract(client *ethclient.Client, contractAddress common.Address) (uint, uint, uint, uint) {
// 	instance, err := implementation.NewImplementation(common.Address, client)
//     if err != nil {
//         log.Fatal(err)
//     }

// 	price, err := instance.contractCost
// 	if err != nil {
//         log.Fatal(err)
//     }
// 	limit, err := instance.limit
// 	if err != nil {
//         log.Fatal(err)
//     }
// 	speed, err := instance.speed
// 	if err != nil {
//         log.Fatal(err)
//     }
// 	length, err := instance.length
// 	if err != nil {
//         log.Fatal(err)
//     }

// 	return price,limit,speed,length
// }
