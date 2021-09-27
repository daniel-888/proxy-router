package contractmanager

import (
// 	"fmt"
"testing"
// 	"context"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/crypto"

// 	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts"
)

func TestBoilerPlateFunc(t *testing.T) {
	msg, err := BoilerPlateFunc()
	if msg != "Contract Manager Package" && err != nil {
		t.Fatalf("Test Failed")
	}
}

// func DeployContract(client *ethclient.Client, 
// 					fromAddress common.Address, 
// 					privateKeyString string, 
// 					constructorParams[] common.Address, 
// 					contract string) common.Address {
// 	privateKey, err := crypto.HexToECDSA(privateKeyString)
//     if err != nil {
//         log.Fatal(err)
//     }
	
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//     if err != nil {
//         log.Fatal(err)
//     }

//     gasPrice, err := client.SuggestGasPrice(context.Background())
//     if err != nil {
//         log.Fatal(err)
//     }

// 	fmt.Println(gasPrice)

// 	auth := bind.NewKeyedTransactor(privateKey)
//     auth.Nonce = big.NewInt(int64(nonce))
//     auth.Value = big.NewInt(0)     // in wei
//     auth.GasLimit = uint64(3000000) // in units
//     auth.GasPrice = gasPrice

// 	switch contract {
// 	case "Ledger":
// 	case "CloneFactory":
// 	case "Validator":
// 	case "Proxy":
// 	case "WebFacing":
// 		param1 := constructorParams[0]
// 		param2 := constructorParams[1]
// 		param3 := constructorParams[2]
// 		param4 := constructorParams[3]
// 		address, _, _, err := webfacing.DeployWebFacing(auth, client, param1, param2, param3, param4)
//     	if err != nil {
//         	log.Fatal(err)
//     	}
// 		return address
// 	default:
// 		address := common.HexToAddress("0x0")
// 		return address
// 	}
// }

// func CreateHashrateContract(client *ethclient.Client, 
// 							fromAddress common.Address, 
// 							privateKeyString string, 
// 							contractAddress common.Address,
// 							price uint,
// 							limit uint,
// 							speed uint,
// 							length uint) {
// 	privateKey, err := crypto.HexToECDSA(privateKeyString)
//     if err != nil {
//         log.Fatal(err)
//     }
	
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//     if err != nil {
//         log.Fatal(err)
//     }

//     gasPrice, err := client.SuggestGasPrice(context.Background())
//     if err != nil {
//         log.Fatal(err)
//     }

// 	fmt.Println(gasPrice)

// 	auth := bind.NewKeyedTransactor(privateKey)
//     auth.Nonce = big.NewInt(int64(nonce))
//     auth.Value = big.NewInt(0)     // in wei
//     auth.GasLimit = uint64(3000000) // in units
//     auth.GasPrice = gasPrice

//     instance, err := webfacing.NewWebFacing(contractAddress, client)
//     if err != nil {
//         log.Fatal(err)
//     }

//     tx, err := instance.setCreateRentalContract(auth, price, limit, speed, length)
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
// }
// func TestBoilerPlateFunc(t *testing.T) {
// 	contractManagerAccount := common.HexToAddress("0xf408f04F9b7691f7174FA2bb73ad6d45fD5d3CBe")
// 	contractManagerPrivateKey := "47b65307d0d654fd4f786b908c04af8fface7710fc998b37d219de19c39ee58c"
// 	rpcClient := "ws://3.217.127.193:8545"
// 	constructorParams := []common.Address {
// 		common.HexToAddress("0xf408f04F9b7691f7174FA2bb73ad6d45fD5d3CBe"), // ledger address
// 		common.HexToAddress("0xf408f04F9b7691f7174FA2bb73ad6d45fD5d3CBe"), // clone factory address
// 		common.HexToAddress("0xf408f04F9b7691f7174FA2bb73ad6d45fD5d3CBe"), // validator address
// 		common.HexToAddress("0xf408f04F9b7691f7174FA2bb73ad6d45fD5d3CBe"), // proxy address
// 	}

// 	client := SetUpClient(contractManagerAccount, rpcClient)
// 	contractAddress := DeployContract(client, contractManagerAccount, contractManagerPrivateKey, constructorParams, "WebFacing")
// 	fmt.Println(contractAddress)
	
// 	SubscribeToContractCreatedEvent(client, contractAddress)
// }