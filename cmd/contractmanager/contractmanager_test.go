package contractmanager

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/clonefactory"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/ledger"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/webfacing"
)

func DeployContract(client *ethclient.Client, 
					fromAddress common.Address, 
					privateKeyString string, 
					constructorParams[] common.Address, 
					contract string) common.Address {
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
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice

	switch contract {
	case "Ledger":
		address, _, _, err := ledger.DeployLedger(auth, client)
		if err != nil {
        	log.Fatal(err)
    	}
		return address	
	case "CloneFactory":
		address, _, _, err := clonefactory.DeployClonefactory(auth, client)
		if err != nil {
        	log.Fatal(err)
    	}
		return address	
	case "WebFacing":
		param1 := constructorParams[0]
		param2 := constructorParams[1]
		param3 := constructorParams[2]
		param4 := constructorParams[3]
		address, _, _, err := webfacing.DeployWebfacing(auth, client, param1, param2, param3, param4)
    	if err != nil {
        	log.Fatal(err)
    	}
		return address	
	}

	address := common.HexToAddress("0x0")
	return address
}

func CreateHashrateContract(client *ethclient.Client, 
							fromAddress common.Address, 
							privateKeyString string, 
							contractAddress common.Address,
							_price uint,
							_limit uint,
							_speed uint,
							_length uint) {
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
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(3000000) // in units
    auth.GasPrice = gasPrice

    instance, err := webfacing.NewWebfacing(contractAddress, client)
    if err != nil {
        log.Fatal(err)
    }

	price := big.NewInt(int64(_price))
	limit := big.NewInt(int64(_limit))
	speed := big.NewInt(int64(_speed))
	length := big.NewInt(int64(_length))
    tx, err := instance.SetCreateRentalContract(auth, price, limit, speed, length)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func PurchaseHashrateContract (client *ethclient.Client, 
						fromAddress common.Address, 
						privateKeyString string, 
						contractAddress common.Address,
						_ipaddress string,
						_username string,
						_password string) {
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
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.SetPurchaseContract(auth, _ipaddress, _username, _password)
	if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func TestContractManager(t *testing.T) {
	ech := make(msgbus.EventChan)
	ps := msgbus.New(1)
	_, _, contractRepo, _, _, _ := externalapi.InitializeJSONRepos()
	var contractMsgs [5]msgbus.Contract
	var hashrateContractAddresses [5]common.Address
	
	// account from aws ganache instance
	contractManagerAccount := common.HexToAddress("0xf408f04F9b7691f7174FA2bb73ad6d45fD5d3CBe")
	contractManagerPrivateKey := "47b65307d0d654fd4f786b908c04af8fface7710fc998b37d219de19c39ee58c"
	rpcClient := "ws://3.217.127.193:8545"
	client := SetUpClient(contractManagerAccount, rpcClient)

	// deploy Clone Factory and Ledger contracts
	cfAddress := DeployContract(client, contractManagerAccount, contractManagerPrivateKey, nil, "CloneFactory")
	fmt.Println("Clone Factory Contract Address: ", cfAddress)
	lAddress := DeployContract(client, contractManagerAccount, contractManagerPrivateKey, nil, "Ledger")
	fmt.Println("Ledger Contract Address: ", lAddress)

	wfConstructorParams := []common.Address {
		lAddress, // Ledger address 
		cfAddress, // Clone Factory address
		common.HexToAddress("0x85A256C5688D012263D5A79EE37E84FC35EC4524"), // contract manager address (dummy address)
		common.HexToAddress("0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3"), // proxy address (dummy address)
	}
	
	// deploy WebFacing contract
	wfAddress := DeployContract(client, contractManagerAccount, contractManagerPrivateKey, wfConstructorParams, "WebFacing")
	fmt.Println("WebFacing Contract Address: ", wfAddress)

	// subcribe to events emitted by WebFacing contract
	wfLogs,wfSub := SubscribeToContractEvents(client, wfAddress)

	// confirm 5 contract creation events were emitted by WebFacing contract and print out address of Hashrate contract created
	for i:=0; i < 5; i++ {
		// create new Hashrate contracts with arbitrary filled out parameters
		CreateHashrateContract(client, contractManagerAccount, contractManagerPrivateKey, wfAddress, uint(i*5), uint(i*10), uint(i*20), uint(i*40))

		select {
		case err := <-wfSub.Err():
			log.Fatal(err)

		case vLog := <-wfLogs:
			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", vLog.Index)

			hashrateContractAddresses[i] = common.HexToAddress(vLog.Topics[1].Hex())
			fmt.Printf("Address of created Hashrate Contract: %s\n", hashrateContractAddresses[i].Hex())
		}
	}

	// subcribe to events emitted by 3rd Hashrate contract
	hLogs,hSub := SubscribeToContractEvents(client, hashrateContractAddresses[2])

	// Purchase 3rd created Hashrate contract to fill out rest of contract parameters and emit purchase event
	PurchaseHashrateContract(client, contractManagerAccount, contractManagerPrivateKey, hashrateContractAddresses[2], "IP Address", "Username", "Password")

	select {
	case err := <-hSub.Err():
		log.Fatal(err)

	case vLog := <-hLogs:
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		buyerAddress := common.HexToAddress(string(vLog.Data))
		fmt.Printf("Address of Buyer: %s\n", buyerAddress.Hex())
	}

	// check data was added to message bus
	go func(ech msgbus.EventChan) {
		for e := range ech {
			if e.EventType == msgbus.GetEvent {
				if e.Data == nil {
					t.Errorf("Failed to add Contract to message bus")
				} 
			}
		}
	}(ech)
	
	// read values from created Hashrate contracts and confirm they are the same as initialization parameters and 3rd contract was updated
	for i:=0; i < 5; i++ {
		contractValues := ReadHashrateContract(client,hashrateContractAddresses[i])
		fmt.Printf("State: %d\n Price: %d\n Limit: %d\n Speed: %d\n Length: %d\n Port: %d\n Validation Fee: %d\n Buyer: %s\n Seller: %s\n IpAddress: %s\n Username: %s\n Password: %s\n",
		contractValues.State,contractValues.Price,contractValues.Limit,contractValues.Speed,contractValues.Length,contractValues.Port,contractValues.ValidationFee,
		contractValues.Buyer.Hex(),contractValues.Seller.Hex(),contractValues.IpAddr,contractValues.Username,contractValues.Password) 
		
		if i == 2 { // purchased Hashrate contract
			if contractValues.State != 2 && contractValues.Price != uint(i*5) && contractValues.Limit != uint(i*10) && contractValues.Speed != uint(i*20) && 
			contractValues.Length != uint(i*40) && contractValues.Port != 0 && contractValues.Seller != contractManagerAccount && contractValues.IpAddr != "IP Address" && 
			contractValues.Username != "Username" && contractValues.Password != "Password" {
				t.Errorf("Read contract values not equal to expected values")
			} 
		} else { // rest of unpurchased Hashrate contract 
			if contractValues.State != 0 && contractValues.Price != uint(i*5) && contractValues.Limit != uint(i*10) && contractValues.Speed != uint(i*20) && 
			contractValues.Length != uint(i*40) && contractValues.Port != 0 && contractValues.Seller != contractManagerAccount && contractValues.IpAddr != "" && 
			contractValues.Username != "" && contractValues.Password != "" {
				t.Errorf("Read contract values not equal to expected values")
			} 
		}
		
		// push read in contract values into message bus contract struct
		contractMsgs[i] = CreateContractMsg(hashrateContractAddresses[i], contractValues)
		ps.Pub(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), msgbus.Contract{})
		ps.Sub(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), ech)
		ps.Set(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])
		ps.Get(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), ech)

		// push read in contract values into API repo
		contractRepo.AddContractFromMsgBus(contractMsgs[i])

		// confirm contract values were pushed to API
		fmt.Printf("API Contract Repo: %+v\n", contractRepo.ContractJSONs[i])
		if len(contractRepo.ContractJSONs) != i + 1 {
			t.Errorf("Contract struct not added")
		} 
	}
}