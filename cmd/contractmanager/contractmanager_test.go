package contractmanager

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/clonefactory"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/ledger"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/lumerintoken"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/webfacing"
)

type TestSetup struct {
	rpcClient *ethclient.Client
	contractManagerPrivateKey	string
	contractManagerAccount	common.Address
	validatorAddress	common.Address
	proxyAddress	common.Address
	lumerinAddress	common.Address
	cloneFactoryAddress	common.Address
	ledgerAddress	common.Address
	webFacingAddress	common.Address
}

func DeployContract(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	constructorParams [5]common.Address,
	contract string) common.Address {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}
	
	time.Sleep(time.Millisecond*700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nonce: ", nonce)

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

	lmnAddress := constructorParams[0]
	validatorAddress := constructorParams[1]
	proxyAddress := constructorParams[2]
	lAddress := constructorParams[3]
	cfAddress := constructorParams[4]

	switch contract {
	case "Ledger":
		address, _, _, err := ledger.DeployLedger(auth, client)
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

	time.Sleep(time.Millisecond*700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nonce: ", nonce)

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
	poolData string) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Millisecond*700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nonce: ", nonce)

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

	instance, err := webfacing.NewWebfacing(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.SetPurchaseContract(auth, _hashrateContract, _buyer, fromAddress, false, poolData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
	fmt.Printf("Hashrate Contract %s, was purchased by %s\n\n", _hashrateContract, _buyer)
}

func SetFundContract(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Millisecond*700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nonce: ", nonce)

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

	tx, err := instance.SetFundContract(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
	fmt.Printf("Hashrate Contract %s was funded\n\n", contractAddress)
}

func BeforeEach() (ts TestSetup) {
	var constructorParams [5]common.Address
	configaData, err := configurationmanager.LoadConfiguration("../configurationmanager/sellerconfig.json", "contractManager")
	if err != nil {
		log.Fatal(err)
	}

	ts.contractManagerAccount = common.HexToAddress(configaData["contractManagerAccount"].(string))
	ts.contractManagerPrivateKey = configaData["contractManagerPrivateKey"].(string)
	ts.rpcClient = SetUpClient(configaData["rpcClientAddress"].(string), common.HexToAddress(configaData["contractManagerAccount"].(string)))
	ts.validatorAddress = common.HexToAddress(configaData["validatorAddress"].(string)) // dummy address
	ts.proxyAddress = common.HexToAddress(configaData["proxyAddress"].(string))         // dummy address
	ts.lumerinAddress = DeployContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,constructorParams,"LumerinToken")
	fmt.Println("Lumerin Token Contract Address: ", ts.lumerinAddress)

	constructorParams[0] = ts.lumerinAddress
	constructorParams[1] = ts.validatorAddress
	constructorParams[2] = ts.proxyAddress

	ts.cloneFactoryAddress = DeployContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,constructorParams,"CloneFactory")
	fmt.Println("Clone Factory Contract Address: ", ts.cloneFactoryAddress)
	ts.ledgerAddress = DeployContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,constructorParams,"Ledger")
	fmt.Println("Ledger Contract Address: ", ts.ledgerAddress)

	constructorParams[3] = ts.ledgerAddress
	constructorParams[4] = ts.cloneFactoryAddress

	ts.webFacingAddress = DeployContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,constructorParams,"WebFacing")
	fmt.Println("Web Facing Contract Address: ", ts.webFacingAddress)

	return ts
} 

func TestLoadMsgBusAndAPIRepo(t *testing.T) {
	ts := BeforeEach()
	ech := make(msgbus.EventChan)
	ps := msgbus.New(1)
	_, _, contractRepo, _, _, _ := externalapi.InitializeJSONRepos()
	var contractMsgs [5]msgbus.Contract
	var hashrateContractAddresses [5]common.Address
	contractAddr := make(chan common.Address, 5)
	stop := make(chan bool)

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := SubscribeToContractEvents(ts.rpcClient, ts.cloneFactoryAddress)

	// confirm contract creation events were emitted by clonefactory contract and print out address of Hashrate contract created
	go func() {
		i := 0
		loop:
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				hashrateContractAddresses[i] = common.HexToAddress(cfLog.Topics[1].Hex())
				contractAddr<-hashrateContractAddresses[i]
				fmt.Printf("Log Block Number: %d\n", cfLog.BlockNumber)
				fmt.Printf("Log Index: %d\n", cfLog.Index)
				fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddresses[i].Hex())
				i++
			case <-stop:
				break loop //stop listening to events once 5 created contracts have been read
			}
		}
		close(cfLogs)
		cfSub.Unsubscribe()
	}()

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
	go func() {
		i := 0
		var address common.Address
		loop:
		for {
			address = <-contractAddr
			contractValues := ReadHashrateContract(ts.rpcClient, address)
			fmt.Printf("%+v\n", contractValues)
	
			if contractValues.State != 0 || contractValues.Price != int(i*5) || contractValues.Limit != int(i*10) || contractValues.Speed != int(i*20) ||
				contractValues.Length != int(i*40) || contractValues.Seller != ts.contractManagerAccount{
				t.Errorf("Read contract values not equal to expected values")
			}
			
			// push read in contract values into message bus contract struct
			contractMsgs[i] = CreateContractMsg(address, contractValues, true)
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
			i++
			if i == 5 { // all created contracts are read
				stop<-true // stop event reading routine and continue test
				break loop
			}
		}
	}()
	
	// create 5 new Hashrate contracts with arbitrary filled out parameters
	for i := 0; i < 5; i++ {
		CreateHashrateContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,ts.webFacingAddress,int(i*5),int(i*10),int(i*20),int(i*40),0)
	}

	<-stop
	close(contractAddr)
	close(stop)

	// subcribe to events emitted by webfacing to read purchase event
	wLogs, wSub := SubscribeToContractEvents(ts.rpcClient, ts.webFacingAddress)

	// purchase 1st created Hashrate contract to fill out rest of contract parameters and emit purchase event
	PurchaseHashrateContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,ts.webFacingAddress,
		hashrateContractAddresses[0], ts.contractManagerAccount, "IP Address")
	
	select {
	case err := <-wSub.Err():
		log.Fatal(err)

	case wLog := <-wLogs:
		fmt.Printf("Log Block Number: %d\n", wLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", wLog.Index)

		contractAddress := common.HexToAddress(string(wLog.Data))
		fmt.Printf("Address of Contract Bought: %s\n\n", contractAddress.Hex())
	}
	close(wLogs)
	wSub.Unsubscribe()

	// confirm purchase changed values in hashrate contract
	purchasedContractValues := ReadHashrateContract(ts.rpcClient, hashrateContractAddresses[0])
	if purchasedContractValues.State != 0 || purchasedContractValues.Price != int(0) || purchasedContractValues.Limit != int(0) || 
		purchasedContractValues.Speed != int(0) || purchasedContractValues.Length != int(0) || 
		purchasedContractValues.Buyer != ts.contractManagerAccount || purchasedContractValues.Seller != ts.contractManagerAccount {
			t.Errorf("Read contract values from purchased contract not equal to expected values")
	}

	miningPoolInfo := ReadMiningPoolInformation(ts.rpcClient, hashrateContractAddresses[0])
	if miningPoolInfo.IpAddress != "IP Address" || miningPoolInfo.Username != "" || miningPoolInfo.Port != "" {
		t.Errorf("Read contract values from purchased contract not equal to expected values")
	}

	// update msgbus struct and contract API repo for contract with new values
	purchasedContractMsg := CreateContractMsg(hashrateContractAddresses[0], purchasedContractValues, true)
	UpdateContractMsgMiningInfo(&purchasedContractMsg, miningPoolInfo)
	ps.Set(msgbus.ContractMsg, msgbus.IDString(purchasedContractMsg.ID), purchasedContractMsg)
	ps.Get(msgbus.ContractMsg, msgbus.IDString(purchasedContractMsg.ID), ech)

	contractJSON := msgdata.ConvertContractMSGtoContractJSON(purchasedContractMsg)
	contractRepo.UpdateContract(contractRepo.ContractJSONs[0].ID, contractJSON)
	fmt.Printf("API Contract Repo: %+v\n\n", contractRepo.ContractJSONs[0])
}	

func TestHashrateMonitoring(t *testing.T) {
	ts := BeforeEach()
	ps := msgbus.New(1)
	var contractMsgs [2]msgbus.Contract
	var hashrateContractAddresses [2]common.Address
	contractAddr := make(chan common.Address, 3)
	stopCF := make(chan bool)
	stopWF := make(chan bool)
	stop := make(chan bool)

	miner1 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID01"),
		IP: 	"IPAddress1",	
		CurrentHashRate:	30,
	}
	miner2 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID02"),
		IP: 	"IPAddress2",
		CurrentHashRate:	20,
	}
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner1.ID),msgbus.Miner{})
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner2.ID),msgbus.Miner{})
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := SubscribeToContractEvents(ts.rpcClient, ts.cloneFactoryAddress)
	
	go func() {
		i := 0
		loop:
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				hashrateContractAddresses[i] = common.HexToAddress(cfLog.Topics[1].Hex())
				contractAddr<-hashrateContractAddresses[i]
				fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddresses[i].Hex())
				i++
			case <-stopCF:
				break loop //stop listening to events
			}
		}
		close(cfLogs)
		cfSub.Unsubscribe()
	}()

	// subcribe to events emitted by webfacing contract to read contract purchase event
	wfLogs, wfSub := SubscribeToContractEvents(ts.rpcClient, ts.webFacingAddress)

	// confirm contract purchase events were emitted by WebFacing contract and print out address of buyer
	go func() {
		//var address common.Address
		loop:
		for {
			select {
			case err := <-wfSub.Err():
				log.Fatal(err)
			case wLog := <-wfLogs:
				//address = <-contractAddr
				contractAddress := common.HexToAddress(string(wLog.Data))
				fmt.Printf("Address of Contract Bought: %s\n\n", contractAddress.Hex())
				//contractAddrRead<-address
			case <-stopWF:
				break loop //stop listening to events
			}
		}
		close(wfLogs)
		wfSub.Unsubscribe()	
	}()

	// routine reads hashrate contract when created and purchases it
	go func() {
		i := 0
		var address common.Address
		var ipaddress string
		loop:
		for {
			address = <-contractAddr
			contractValues := ReadHashrateContract(ts.rpcClient, address)
			fmt.Printf("%+v\n", contractValues)
			
			// push read in contract values into message bus contract struct
			contractMsgs[i] = CreateContractMsg(address, contractValues, true)
			ps.Pub(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), msgbus.Contract{})
			ps.Set(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])
			
			// purchase new Hashrate contracts with different miner ip addresses		
			switch i {
			case 0:
				ipaddress = miner1.IP
			case 1:
				ipaddress = miner2.IP
			}
			PurchaseHashrateContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,ts.webFacingAddress,
				address,ts.contractManagerAccount,ipaddress)
			i++
			fmt.Println(i)
			if i == 2 {
				stopCF<-true // stop clonefactory event reading routine
				stopWF<-true // stop webfacing event reading routine
				stop<-true // continue to rest of test
				break loop
			}
		}
	}()

	// create 2 new Hashrate contracts with arbitrary filled out parameters
	for i := 0; i < 2; i++ {
		CreateHashrateContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,ts.webFacingAddress,int(0),int(0),int(30),int(0),int(0))		
	}

	<-stop
	close(contractAddr)
	close(stopCF)
	close(stopWF)

	// read mining pool info from purchased contracts
	miningPool1Info := ReadMiningPoolInformation(ts.rpcClient, hashrateContractAddresses[0])
	miningPool2Info := ReadMiningPoolInformation(ts.rpcClient, hashrateContractAddresses[1])

	// find Miner associated with IP Address set in contract purchase
	e1, err := ps.SearchIPWait(msgbus.MinerMsg, miningPool1Info.IpAddress)
	if err != nil {
		log.Fatal(err)
	}
	e2, err := ps.SearchIPWait(msgbus.MinerMsg, miningPool2Info.IpAddress)
	if err != nil {
		log.Fatal(err)
	}

	searchedMiner1 := e1.Data.(msgbus.IDIndex)
	searchedMiner2 := e2.Data.(msgbus.IDIndex)
	
	if !CloseOutMonitor(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,hashrateContractAddresses[0],searchedMiner1[0],contractMsgs[0],ps) {
		t.Errorf("Closeout monitor incorrectly closed out contract")
	}	
	if CloseOutMonitor(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,hashrateContractAddresses[1],searchedMiner2[0],contractMsgs[1],ps) {
		t.Errorf("Closeout monitor did not closeout contract that was not fulfilling requirements")
	}
}

func TestCreateUnsignedTransaction(t *testing.T) {
    ts := BeforeEach()
    var hashrateContractAddress common.Address

    // subcribe to events emitted by clonefactory contract to read contract creation event
    cfLogs, cfSub := SubscribeToContractEvents(ts.rpcClient, ts.cloneFactoryAddress)

    CreateHashrateContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,ts.webFacingAddress,int(0),int(0),int(30),int(0),int(0)) 

    select {
    case err := <-cfSub.Err():
        log.Fatal(err)
    case cfLog := <-cfLogs:
        hashrateContractAddress = common.HexToAddress(cfLog.Topics[1].Hex())
        fmt.Printf("Log Block Number: %d\n", cfLog.BlockNumber)
        fmt.Printf("Log Index: %d\n", cfLog.Index)
        fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddress.Hex())
    }

	PurchaseHashrateContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,ts.webFacingAddress,
		hashrateContractAddress,ts.contractManagerAccount,"IpAddress")

	contractValues := ReadHashrateContract(ts.rpcClient, hashrateContractAddress)
	fmt.Println("Contract State before closeout: ", contractValues.State)
	
    hashrateContractABI, err := abi.JSON(strings.NewReader(implementation.ImplementationABI))
    if err != nil {
        log.Fatal(err)
    }

    nonce, err := ts.rpcClient.PendingNonceAt(context.Background(), ts.contractManagerAccount)
    if err != nil {
        log.Fatal(err)
    }

    bytesData, err := hashrateContractABI.Pack("setContractCloseOut")
    if err != nil {
        log.Fatal(err)
    }

    gasPrice, err := ts.rpcClient.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    gasLimit := uint64(3000000)

    unsignedTx := types.NewTransaction(nonce, hashrateContractAddress, nil, gasLimit, gasPrice, bytesData)
    fmt.Printf("Unsigned Transaction: %+v\n", unsignedTx)
    fmt.Println("Unsigned Transaction Hash: ", unsignedTx.Hash())

	//Sign transaction
	privateKey, err := crypto.HexToECDSA(ts.contractManagerPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = ts.rpcClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	contractValues = ReadHashrateContract(ts.rpcClient, hashrateContractAddress)
	fmt.Println("Contract State after closeout: ", contractValues.State)
}
func TestSellerRoutine(t *testing.T) {
	ps := msgbus.New(10)
	ts := BeforeEach()
	var hashrateContractAddress [3]common.Address  
	// hashrateContractAddress[0] = common.HexToAddress("0xa5e6cd816545c883bfa246e96bf7d3648d84d881")
	// hashrateContractAddress[1] = common.HexToAddress("0xbb05218023c62fe691bb78b3969eab50077b6a07")
	
	contractmanagerConfig, err := configurationmanager.LoadConfiguration("../configurationmanager/sellerconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	cman, err := New(ps, contractmanagerConfig)
	if err != nil {
		panic(fmt.Sprintf("contract manager failed:%s", err))
	}
	cman.webFacingAddress = ts.webFacingAddress
	cman.cloneFactoryAddress = ts.cloneFactoryAddress
	cman.ledgerAddress = ts.ledgerAddress
	
	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := SubscribeToContractEvents(cman.rpcClient, cman.cloneFactoryAddress)
	go func () {
		i := 0
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				hashrateContractAddress[i] = common.HexToAddress(cfLog.Topics[1].Hex())
				fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddress[i].Hex())
				
				i++
			}
		}
	}()
	
	CreateHashrateContract(cman.rpcClient,cman.account,cman.privateKey,cman.webFacingAddress,int(0),int(0),int(30),int(10),int(0)) 
	CreateHashrateContract(cman.rpcClient,cman.account,cman.privateKey,cman.webFacingAddress,int(0),int(0),int(30),int(10),int(0)) 
	time.Sleep(time.Millisecond*10000)

	err = cman.StartSeller()
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}

	time.Sleep(time.Millisecond*10000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[0], cman.account, "IpAddress1|8888|ryan")
	time.Sleep(time.Millisecond*10000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[0])
	time.Sleep(time.Millisecond*10000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[1], cman.account, "IpAddress2|8888|ryan")
	time.Sleep(time.Millisecond*10000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[1])
	time.Sleep(time.Millisecond*10000)

	// test early closeout
	CreateHashrateContract(cman.rpcClient,cman.account,cman.privateKey,cman.webFacingAddress,int(0),int(0),int(30),int(100),int(0)) 
	time.Sleep(time.Millisecond*10000)

	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[2], cman.account, "IpAddress3|8888|ryan")
	time.Sleep(time.Millisecond*10000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[2])
	time.Sleep(time.Millisecond*10000)
	setContractCloseOut(cman.rpcClient,cman.account,cman.privateKey,hashrateContractAddress[2])
	waitchan := make(chan bool)
	<-waitchan
}

func TestBuyerRoutine(t *testing.T) {
	ps := msgbus.New(10)
	ts := BeforeEach()
	var hashrateContractAddress [3]common.Address  
	
	contractmanagerConfig, err := configurationmanager.LoadConfiguration("../configurationmanager/sellerconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	cman, err := New(ps, contractmanagerConfig)
	if err != nil {
		panic(fmt.Sprintf("contract manager failed:%s", err))
	}
	cman.webFacingAddress = ts.webFacingAddress
	cman.cloneFactoryAddress = ts.cloneFactoryAddress
	cman.ledgerAddress = ts.ledgerAddress

	miner1 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID01"),
		IP: 	"IpAddress1",	
		CurrentHashRate:	30,
	}
	miner2 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID02"),
		IP: 	"IpAddress2",
		CurrentHashRate:	20,
	}
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner1.ID),miner1)
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner2.ID),miner2)

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := SubscribeToContractEvents(cman.rpcClient, cman.cloneFactoryAddress)
	go func () {
		i := 0
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				hashrateContractAddress[i] = common.HexToAddress(cfLog.Topics[1].Hex())
				fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddress[i].Hex())
				
				i++
			}
		}
	}()

	CreateHashrateContract(cman.rpcClient,cman.account,cman.privateKey,cman.webFacingAddress,int(0),int(0),int(30),int(10),int(0)) 
	CreateHashrateContract(cman.rpcClient,cman.account,cman.privateKey,cman.webFacingAddress,int(0),int(0),int(30),int(10),int(0))

	time.Sleep(time.Millisecond*5000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[0], cman.account, "IpAddress1|8888|ryan")
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[1], cman.account, "IpAddress2|8888|ryan")

	time.Sleep(time.Millisecond*5000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[0])
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[1])
	
	err = cman.StartBuyer()
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}

	time.Sleep(time.Millisecond*20000)
	
	// miner hashrate fall below promised hashrate
	miner1.CurrentHashRate = 20
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	

	time.Sleep(time.Millisecond*5000)
	miner3 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID03"),
		IP: 	"IpAddress3",
		CurrentHashRate:	30,
	}
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner3.ID),miner3)

	time.Sleep(time.Millisecond*10000)
	CreateHashrateContract(cman.rpcClient,cman.account,cman.privateKey,cman.webFacingAddress,int(0),int(0),int(30),int(10),int(0)) 

	time.Sleep(time.Millisecond*5000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[2], cman.account, "IpAddress3|8888|ryan")

	time.Sleep(time.Millisecond*5000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[2])

	waitchan := make(chan bool)
	<-waitchan
}

func TestDeployment(t *testing.T) {
	ts := BeforeEach()
	CreateHashrateContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,ts.webFacingAddress,int(0),int(0),int(30),int(10),int(0)) 
	CreateHashrateContract(ts.rpcClient,ts.contractManagerAccount,ts.contractManagerPrivateKey,ts.webFacingAddress,int(0),int(0),int(30),int(10),int(0)) 
}