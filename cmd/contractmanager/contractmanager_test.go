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

	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/clonefactory"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/ledger"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/webfacing"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/lumerintoken"
)

func DeployContract(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	constructorParams [5]common.Address,
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	lmnAddress := constructorParams[0]
	validatorAddress := constructorParams[1]
	proxyAddress := constructorParams[2]
	lAddress := constructorParams[3]
	cfAddress := constructorParams[4]

	switch contract {
	case "Ledger":
		address, _, _, err := ledger.DeployLedger(auth, client, validatorAddress)
		if err != nil {
			log.Fatal(err)
		}
		return address
	case "LumerinToken":
		address, _, _, err := lumerintoken.DeployLumerintoken(auth, client)
		if err != nil {
			log.Fatal(err)
		}
		return address
	case "CloneFactory":
		address, _, _, err := clonefactory.DeployClonefactory(auth, client, lmnAddress, validatorAddress, proxyAddress)
		if err != nil {
			log.Fatal(err)
		}
		return address
	case "WebFacing":
		address, _, _, err := webfacing.DeployWebfacing(auth, client, lAddress, cfAddress)
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
	_price int,
	_limit int,
	_speed int,
	_length int,
	_validationFee int) {
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

	instance, err := webfacing.NewWebfacing(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	price := big.NewInt(int64(_price))
	limit := big.NewInt(int64(_limit))
	speed := big.NewInt(int64(_speed))
	length := big.NewInt(int64(_length))
	validationFee := big.NewInt(int64(_validationFee))
	tx, err := instance.SetCreateRentalContract(auth, price, limit, speed, length, validationFee)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
}

func PurchaseHashrateContract(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address,
	_hashrateContract common.Address,
	_buyer common.Address,
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
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := webfacing.NewWebfacing(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.SetPurchaseContract(auth, _hashrateContract, _buyer, fromAddress,false,_ipaddress, _username, _password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
}

func SetFundHashrateContract(client *ethclient.Client,
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

	tx, err := instance.SetFundContract(auth)
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
	var constructorParams [5]common.Address 

	cmConfig, err := configurationmanager.LoadConfiguration("../configurationmanager/testconfig.json", "contractManager")
	if err != nil {
		log.Fatal(err)
	}

	// account from aws ganache instance
	contractManagerAccount := common.HexToAddress(cmConfig["contractManagerAccount"].(string))
	contractManagerPrivateKey := cmConfig["contractManagerPrivateKey"].(string)
	rpcClient := cmConfig["rpcClientAddress"].(string)
	client := SetUpClient(contractManagerAccount, rpcClient)

	validatorAddress := common.HexToAddress(cmConfig["validatorAddress"].(string)) // dummy address
	proxyAddress := common.HexToAddress(cmConfig["proxyAddress"].(string))         // dummy address

	// deploy Lumerin Token
	lmnAddress := DeployContract(client, contractManagerAccount, contractManagerPrivateKey, constructorParams, "LumerinToken")

	constructorParams[0] = lmnAddress
	constructorParams[1] = validatorAddress
	constructorParams[2] = proxyAddress

	// deploy Clone Factory and Ledger contracts
	cfAddress := DeployContract(client, contractManagerAccount, contractManagerPrivateKey, constructorParams, "CloneFactory")
	fmt.Println("Clone Factory Contract Address: ", cfAddress)
	lAddress := DeployContract(client, contractManagerAccount, contractManagerPrivateKey, constructorParams, "Ledger")
	fmt.Println("Ledger Contract Address: ", lAddress)

	constructorParams[3] = lAddress
	constructorParams[4] = cfAddress

	// deploy WebFacing contract
	wfAddress := DeployContract(client, contractManagerAccount, contractManagerPrivateKey, constructorParams, "WebFacing")
	fmt.Println("WebFacing Contract Address: ", wfAddress)

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := SubscribeToContractEvents(client, cfAddress)

	// confirm 5 contract creation events were emitted by WebFacing contract and print out address of Hashrate contract created
	for i := 0; i < 5; i++ {
		// create new Hashrate contracts with arbitrary filled out parameters
		CreateHashrateContract(client, contractManagerAccount, contractManagerPrivateKey, wfAddress, int(i*5), int(i*10), int(i*20), int(i*40), 30)

		select {
		case err := <-cfSub.Err():
			log.Fatal(err)

		case cfLog := <-cfLogs:
			fmt.Printf("Log Block Number: %d\n", cfLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", cfLog.Index)

			hashrateContractAddresses[i] = common.HexToAddress(cfLog.Topics[1].Hex())
			fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddresses[i].Hex())
		}
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
	for i := 0; i < 5; i++ {
		contractValues := ReadHashrateContract(client, hashrateContractAddresses[i])
		fmt.Printf("%+v\n", contractValues)

		if contractValues.State != 0 || contractValues.Price != int(i*5) || contractValues.Limit != int(i*10) || contractValues.Speed != int(i*20) ||
			contractValues.Length != int(i*40) || contractValues.Port != 0 || contractValues.Seller != contractManagerAccount{
			t.Errorf("Read contract values not equal to expected values")
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
		fmt.Printf("API Contract Repo: %+v\n\n", contractRepo.ContractJSONs[i])
		if len(contractRepo.ContractJSONs) != i+1 {
			t.Errorf("Contract struct not added")
		}
	}

	// subcribe to events emitted by webfacing to read purchase event
	wLogs, wSub := SubscribeToContractEvents(client, wfAddress)

	// account from ganache instance
	buyerAddress := common.HexToAddress("0x25c1230C7EFC00cFd2fcAA3a44f30948853824bc") 
	// purchase 3rd created Hashrate contract to fill out rest of contract parameters and emit purchase event
	PurchaseHashrateContract(client, contractManagerAccount, contractManagerPrivateKey, wfAddress,
		hashrateContractAddresses[2], buyerAddress, "IP Address", "Username", "Password")
	// SetFundHashrateContract(client, contractManagerAccount, contractManagerPrivateKey, hashrateContractAddresses[2])

	select {
	case err := <-wSub.Err():
		log.Fatal(err)

	case wLog := <-wLogs:
		fmt.Printf("Log Block Number: %d\n", wLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", wLog.Index)

		buyerAddress := common.HexToAddress(string(wLog.Data))
		fmt.Printf("Address of Buyer: %s\n\n", buyerAddress.Hex())
	}

	// confirm purchase changed values in hashrate contract
	purchasedContractValues := ReadHashrateContract(client, hashrateContractAddresses[2])
	if purchasedContractValues.State != 0 || purchasedContractValues.Price != int(10) || purchasedContractValues.Limit != int(20) || 
		purchasedContractValues.Speed != int(40) || purchasedContractValues.Length != int(80) || purchasedContractValues.Port != 0 || purchasedContractValues.Buyer != buyerAddress ||
		purchasedContractValues.Seller != contractManagerAccount {
			t.Errorf("Read contract values from purchased contract not equal to expected values")
	}

	// update msgbus struct and contract API repo for contract with new values
	purchasedContractMsg := CreateContractMsg(hashrateContractAddresses[2], purchasedContractValues)
	ps.Set(msgbus.ContractMsg, msgbus.IDString(purchasedContractMsg.ID), purchasedContractMsg)
	ps.Get(msgbus.ContractMsg, msgbus.IDString(purchasedContractMsg.ID), ech)

	contractJSON := msgdata.ConvertContractMSGtoContractJSON(purchasedContractMsg)
	contractRepo.UpdateContract(contractRepo.ContractJSONs[2].ID, contractJSON)
	fmt.Printf("API Contract Repo: %+v\n\n", contractRepo.ContractJSONs[2])

	// Contract Manager Receives message from validator and checks to see contract is meeting threshold requirements
	// thresholds := DefineThresholdParams("../configurationmanager/testconfig.json")
	// vmsg := make(chan ValidatorMsg)

	// go func (){
	// 	valmsg := ValidatorMsg {
	// 		ShareAmtPerMin:		5,
	// 		ShareAvgPerHour:	5,
	// 		ShareDrop:			5,
	// 		HashesCompleted:	10,
	// 		ContractFulfilled:	false,
	// 	} 
	// 	vmsg<-valmsg
	// }()
	
	// CloseOutMonitor(client, contractManagerAccount, contractManagerPrivateKey, hashrateContractAddresses[2], vmsg, thresholds)

}	