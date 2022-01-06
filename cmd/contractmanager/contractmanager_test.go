package contractmanager

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/clonefactory"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/lumerintoken"
)

type TestSetup struct {
	rpcClient              *ethclient.Client
	nodeEthereumPrivateKey string
	nodeEthereumAccount    common.Address
	validatorAddress       common.Address
	proxyAddress           common.Address
	lumerinAddress         common.Address
	cloneFactoryAddress    common.Address
}

func DeployContract(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	constructorParams [5]common.Address,
	contract string) common.Address {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	time.Sleep(time.Millisecond * 700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	fmt.Println("Nonce: ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	lmnAddress := constructorParams[0]
	validatorAddress := constructorParams[1]
	proxyAddress := constructorParams[2]

	switch contract {
	case "Ledger":

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
	_validator common.Address) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	time.Sleep(time.Millisecond * 700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	fmt.Println("Nonce: ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := clonefactory.NewClonefactory(contractAddress, client)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	price := big.NewInt(int64(_price))
	limit := big.NewInt(int64(_limit))
	speed := big.NewInt(int64(_speed))
	length := big.NewInt(int64(_length))
	tx, err := instance.SetCreateNewRentalContract(auth, price, limit, speed, length, _validator)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
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
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	time.Sleep(time.Millisecond * 700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	fmt.Println("Nonce: ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := clonefactory.NewClonefactory(contractAddress, client)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	tx, err := instance.SetPurchaseRentalContract(auth, _hashrateContract, poolData)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
	fmt.Printf("Hashrate Contract %s, was purchased by %s\n\n", _hashrateContract, _buyer)
}

func UpdatePurchaseInformation(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address,
	_price int,
	_limit int,
	_speed int,
	_length int) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	time.Sleep(time.Millisecond * 700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	fmt.Println("Nonce: ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	price := big.NewInt(int64(_price))
	limit := big.NewInt(int64(_limit))
	speed := big.NewInt(int64(_speed))
	length := big.NewInt(int64(_length))
	closeOutType := big.NewInt(int64(3)) 
	tx, err := instance.SetUpdatePurchaseInformation(auth, price, limit, speed, length, closeOutType)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
	fmt.Printf("Hashrate Contract %s purchase info was updated\n\n", contractAddress)
}

func UpdateMiningInformation(client *ethclient.Client,
	fromAddress common.Address,
	privateKeyString string,
	contractAddress common.Address,
	_newEncryptedPoolData string) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	time.Sleep(time.Millisecond * 700)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	fmt.Println("Nonce: ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	instance, err := implementation.NewImplementation(contractAddress, client)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	tx, err := instance.SetUpdateMiningInformation(auth, _newEncryptedPoolData)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}
	fmt.Printf("tx sent: %s\n\n", tx.Hash().Hex())
	fmt.Printf("Hashrate Contract %s mining info was updated\n\n", contractAddress)
}

func encryptData(msg string, privkey string) []byte {
	message := []byte(msg)
	privateKey, err := crypto.HexToECDSA(privkey)
    if err != nil {
        log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
    }
    publicKey := privateKey.Public()
	pubKey, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
		err = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
    }
	pubKeyECIES := ecies.ImportECDSAPublic(pubKey)
	encryptData, err := ecies.Encrypt(rand.Reader, pubKeyECIES, message, nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	return encryptData
}

func BeforeEach() (ts TestSetup) {
	var constructorParams [5]common.Address
	configaData, err := configurationmanager.LoadConfiguration("ganachetestconfig.json", "contractManager")
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	var client *ethclient.Client
	client, err = setUpClient(configaData["rpcClientAddress"].(string), common.HexToAddress(configaData["sellerEthereumAddress"].(string)))
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	ts.nodeEthereumAccount = common.HexToAddress(configaData["sellerEthereumAddress"].(string))
	ts.nodeEthereumPrivateKey = configaData["sellerEthereumPrivateKey"].(string)
	ts.rpcClient = client
	ts.validatorAddress = common.HexToAddress(configaData["validatorAddress"].(string)) // dummy address
	ts.proxyAddress = common.HexToAddress(configaData["proxyAddress"].(string))         // dummy address
	ts.lumerinAddress = DeployContract(ts.rpcClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, constructorParams, "LumerinToken")
	fmt.Println("Lumerin Token Contract Address: ", ts.lumerinAddress)

	constructorParams[0] = ts.lumerinAddress
	constructorParams[1] = ts.validatorAddress
	constructorParams[2] = ts.proxyAddress
	
	ts.cloneFactoryAddress = DeployContract(ts.rpcClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, constructorParams, "CloneFactory")
	fmt.Println("Clone Factory Contract Address: ", ts.cloneFactoryAddress)

	return ts
}

func TestSellerRoutine(t *testing.T) {
	ps := msgbus.New(10)
	ts := BeforeEach()
	var hashrateContractAddress [4]common.Address
	var purchasedHashrateContractAddress [4]common.Address
	//encryptedDest := encryptData("stratum+tcp://127.0.0.1:3333/testrig", cman.privateKey) 

	contractManagerConfig, err := configurationmanager.LoadConfiguration("ganachetestconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	contractLength := 15 // 15 s on ganache
	sleepTime := 5000 // 5000 ms sleeptime in ganache
	if contractManagerConfig["rpcClientAddress"].(string) != "ws://127.0.0.1:7545" {
		contractLength = 60 // 60 s on ropsten
		sleepTime = 30000 // 30000 ms on testnet
	}
	buyerAddress := common.HexToAddress(contractManagerConfig["buyerEthereumAddress"].(string))
	buyerPrivateKey := contractManagerConfig["buyerEthereumPrivateKey"].(string)

	var cman SellerContractManager
	err = cman.init(ps, contractManagerConfig)
	if err != nil {
		panic(fmt.Sprintf("contract manager init failed:%s", err))
	}
	cman.cloneFactoryAddress = ts.cloneFactoryAddress

	// subcribe to creation events emitted by clonefactory contract
	cfLogs, cfSub, _ := subscribeToContractEvents(ts.rpcClient, ts.cloneFactoryAddress)
	// create event signature to parse out creation and purchase event
	contractCreatedSig := []byte("contractCreated(address)")
	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
	clonefactoryContractPurchasedSig := []byte("clonefactoryContractPurchased(address)")
	clonefactoryContractPurchasedSigHash := crypto.Keccak256Hash(clonefactoryContractPurchasedSig)
	go func() {
		i := 0
		j := 0
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
			case cfLog := <-cfLogs:
				switch {
				case cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex():
					hashrateContractAddress[i] = common.HexToAddress(cfLog.Topics[1].Hex())
					fmt.Printf("Address of created Hashrate Contract %d: %s\n\n", i + 1, hashrateContractAddress[i].Hex())
					i++
				
				case cfLog.Topics[0].Hex() == clonefactoryContractPurchasedSigHash.Hex():
					purchasedHashrateContractAddress[j] = common.HexToAddress(cfLog.Topics[1].Hex())
					fmt.Printf("Address of purchased Hashrate Contract %d: %s\n\n", j + 1, purchasedHashrateContractAddress[j].Hex())
					j++
				}
			}
		}
	}()
	
	//
	// test startup with 1 running contract and 1 availabe contract
	//
	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength), buyerAddress)
	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength), buyerAddress)
	
	// wait until created hashrate contract was found before continuing 
	loop1:
	for {
		if hashrateContractAddress[0] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop1	
		}
	}
	PurchaseHashrateContract(cman.rpcClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[0], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing 
	loop2:
	for {
		if purchasedHashrateContractAddress[0] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop2	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))

	err = cman.start()
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}
	
	// contract manager sees existing contracts and states are correct
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 1 was not found or is not in correct state")
	}
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] != msgbus.ContAvailableState {
		t.Errorf("Contract 2 was not found or is not in correct state")
	}

	// contract should be back to available after length has expired
	go func() {
		time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
		time.Sleep(time.Second * time.Duration(sleepTime)) // length of transaction
		if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 1 did not close out correctly")
		}
	}()
	
	// contract manager should updated state
	// wait until created hashrate contract was found before continuing 
	loop3:
	for {
		if hashrateContractAddress[1] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop3	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	PurchaseHashrateContract(cman.rpcClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[1], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing 
	loop4:
	for {
		if purchasedHashrateContractAddress[1] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop4	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 2 is not in correct state")
	}

	// contract should be back to available after length has expired
	go func() {
		time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
		time.Sleep(time.Second * time.Duration(sleepTime)) // length of transaction
		if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 2 did not close out correctly")
		}
	}()

	//
	// test early closeout from buyer
	//
	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength*10), buyerAddress)

	// wait until created hashrate contract was found before continuing 
	loop5:
	for {
		if hashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop5	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContAvailableState {
		t.Errorf("Contract 3 was not found or is not in correct state")
	}
	PurchaseHashrateContract(cman.rpcClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[2], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing 
	loop6:
	for {
		if purchasedHashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop6	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 3 is not in correct state")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	setContractCloseOut(cman.rpcClient, buyerAddress, buyerPrivateKey, hashrateContractAddress[2], &wg, &cman.currentNonce, 0)
	wg.Wait()
	time.Sleep(time.Millisecond * time.Duration(sleepTime*2))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContAvailableState {
		t.Errorf("Contract 3 did not close out correctly")
	}

	//
	// test contract creation and going through full length with update made to mining info from buyer while node is running
	//
	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength), buyerAddress)
	
	// wait until created hashrate contract was found before continuing 
	loop7:
	for {
		if hashrateContractAddress[3] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop7	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[3].Hex())] != msgbus.ContAvailableState {
		t.Errorf("Contract 4 was not found or is not in correct state")
	}

	PurchaseHashrateContract(cman.rpcClient, buyerAddress, buyerPrivateKey, ts.cloneFactoryAddress, hashrateContractAddress[3], buyerAddress, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing 
	loop8:
	for {
		if purchasedHashrateContractAddress[3] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop8	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[3].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 4 is not in correct state")
	}

	UpdateMiningInformation(cman.rpcClient, buyerAddress, buyerPrivateKey, hashrateContractAddress[3], "stratum+tcp://127.0.0.1:3333/updated")
	time.Sleep(time.Millisecond * time.Duration(sleepTime*2))
	// check dest msg with associated contract was updated in msgbus
	event, err := cman.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(hashrateContractAddress[3].Hex()))
	if err != nil {
		panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
	}
	contractMsg := event.Data.(msgbus.Contract)
	event, err = cman.ps.GetWait(msgbus.DestMsg, msgbus.IDString(contractMsg.Dest))
	if err != nil {
		panic(fmt.Sprintf("Getting Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Getting Dest Failed: %s", event.Err))
	}
	destMsg := event.Data.(msgbus.Dest)
	if destMsg.NetUrl != "stratum+tcp://127.0.0.1:3333/updated" {
		t.Errorf("Contract 4's target dest was not updated")
	}

	// if network is ganache, create a new transaction so a new block is created 
	if contractManagerConfig["rpcClientAddress"].(string) == "ws://127.0.0.1:7545" {
		time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
		
		nonce, err := ts.rpcClient.PendingNonceAt(context.Background(), cman.account)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := ts.rpcClient.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		gasLimit := uint64(3000000)

		unsignedTx := types.NewTransaction(nonce, cman.account, nil, gasLimit, gasPrice, nil)
		fmt.Printf("Unsigned Transaction: %+v\n", unsignedTx)
		fmt.Println("Unsigned Transaction Hash: ", unsignedTx.Hash())
	
		//Sign transaction
		privateKey, err := crypto.HexToECDSA(cman.privateKey)
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

		time.Sleep(time.Millisecond * time.Duration(sleepTime)) // length of transaction
		if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[3].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 4 did not close out correctly")
		}
	} else {
		time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
		time.Sleep(time.Millisecond * time.Duration(sleepTime)) // length of transaction
		if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[3].Hex())] != msgbus.ContAvailableState {
			t.Errorf("Contract 4 did not close out correctly")
		}
	}

	//
	// test seller updated purchase info
	//
	UpdatePurchaseInformation(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[3], int(10), int(10), int(50), int(contractLength*10))
	time.Sleep(time.Millisecond * time.Duration(sleepTime*2))
	// check purchase information in contract was updated in msgbus
	event, err = cman.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(hashrateContractAddress[3].Hex()))
	if err != nil {
		panic(fmt.Sprintf("Getting Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Getting Contract Failed: %s", event.Err))
	}
	contractMsg = event.Data.(msgbus.Contract)
	if contractMsg.Limit != 10 || contractMsg.Speed != 50 {
		t.Errorf("Contract 4's purchase info was not updated")
	} 
}

func TestBuyerRoutine(t *testing.T) {
	ps := msgbus.New(10)
	ts := BeforeEach()
	var hashrateContractAddress [3]common.Address
	var purchasedHashrateContractAddress [3]common.Address

	contractLength := 10000 
	// encryptedDest := encryptData("stratum+tcp://127.0.0.1:3333/testrig", cman.privateKey) 

	contractManagerConfig, err := configurationmanager.LoadConfiguration("ganachetestconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	sleepTime := 5000 // 5000 ms sleeptime in ganache
	if contractManagerConfig["rpcClientAddress"].(string) != "ws://127.0.0.1:7545" {
		sleepTime = 20000 // 20000 ms on testnet
	}
	sellerAddress := common.HexToAddress(contractManagerConfig["sellerEthereumAddress"].(string))
	sellerPrivateKey := contractManagerConfig["sellerEthereumPrivateKey"].(string)

	var cman BuyerContractManager
	err = cman.init(ps, contractManagerConfig)
	if err != nil {
		panic(fmt.Sprintf("contract manager init failed:%s", err))
	}
	cman.cloneFactoryAddress = ts.cloneFactoryAddress

	miner1 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID01"),
		IP: 	"IpAddress1",
		CurrentHashRate:	30,
		State: msgbus.OnlineState,
	}
	miner2 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID02"),
		IP: 	"IpAddress2",
		CurrentHashRate:	20,
		State: msgbus.OnlineState,
	}
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner1.ID),miner1)
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner2.ID),miner2)

	// subcribe to creation events emitted by clonefactory contract 
	cfLogs, cfSub, _ := subscribeToContractEvents(ts.rpcClient, ts.cloneFactoryAddress)
	// create event signature to parse out creation event
	contractCreatedSig := []byte("contractCreated(address)")
	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
	clonefactoryContractPurchasedSig := []byte("clonefactoryContractPurchased(address)")
	clonefactoryContractPurchasedSigHash := crypto.Keccak256Hash(clonefactoryContractPurchasedSig)
	go func() {
		i := 0
		j := 0
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
			case cfLog := <-cfLogs:
				switch {
				case cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex():
					hashrateContractAddress[i] = common.HexToAddress(cfLog.Topics[1].Hex())
					fmt.Printf("Address of created Hashrate Contract %d: %s\n\n", i + 1, hashrateContractAddress[i].Hex())
					i++
				
				case cfLog.Topics[0].Hex() == clonefactoryContractPurchasedSigHash.Hex():
					purchasedHashrateContractAddress[j] = common.HexToAddress(cfLog.Topics[1].Hex())
					fmt.Printf("Address of purchased Hashrate Contract %d: %s\n\n", j + 1, purchasedHashrateContractAddress[j].Hex())
					j++
				}
			}
		}
	}()

	//
	// test startup with 1 running contract and 1 availabe contract
	//
	CreateHashrateContract(cman.rpcClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(0), int(1000), int(contractLength), cman.account)
	CreateHashrateContract(cman.rpcClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(0), int(1000), int(contractLength), cman.account)

	// wait until created hashrate contract was found before continuing 
	loop1:
	for {
		if hashrateContractAddress[0] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop1	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[0], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing
	loop2:
	for {
		if purchasedHashrateContractAddress[0] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop2	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	
	err = cman.start()
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}
	
	// check contract manager found miners in msgbus 
	if cman.miners[miner1.ID] != miner1 {
		t.Errorf("MinerID01 not found by contract manager")
	}
	if cman.miners[miner2.ID] != miner2 {
		t.Errorf("MinerID02 not found by contract manager")
	}

	// contract manager sees existing contracts and states are correct
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 1 was not found or is not in correct state")
	}
	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] ; ok {
		t.Errorf("Contract 2 was found by buyer node while in the available state")
	}

	// contract should be removed from buyer node when set back to available
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
	// 	time.Sleep(time.Millisecond * time.Duration(sleepTime)) // length of transaction
	// 	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[0].Hex())]; ok {
	// 		t.Errorf("Contract 1 did not close out correctly")
	// 	}
	// }()

	// contract manager should updated states
	// wait until created hashrate contract was found before continuing 
	loop3:
	for {
		if hashrateContractAddress[1] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop3	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[1], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing
	loop4:
	for {
		if purchasedHashrateContractAddress[1] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop4	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 2 is not in correct state")
	}

	// contract should be removed from buyer node when set back to available
	// go func() {
	// 	time.Sleep(time.Second * time.Duration(contractLength)) // length of contract
	// 	time.Sleep(time.Millisecond * time.Duration(sleepTime)) // length of transaction
	// 	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[1].Hex())]; ok {
	// 		t.Errorf("Contract 2 did not close out correctly")
	// 	}
	// }()

	/*
	//
	// Test early closeout from seller
	//
	CreateHashrateContract(cman.rpcClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(0), int(30), int(contractLength*10), cman.account)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] ; ok {
		t.Errorf("Contract 3 was found by buyer node while in the available state")
	}

	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[2], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 3 is not in correct state")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	setContractCloseOut(cman.rpcClient, sellerAddress, sellerPrivateKey, hashrateContractAddress[2], &wg, &cman.currentNonce, 0)
	wg.Wait()
	time.Sleep(time.Millisecond * time.Duration(sleepTime))
	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())]; ok {
		t.Errorf("Contract 3 did not close out correctly")
	}
	*/

	//
	// Test contract creation, purchasing, and target dest being updated while node is running
	//
	CreateHashrateContract(cman.rpcClient, sellerAddress, sellerPrivateKey, ts.cloneFactoryAddress, int(0), int(0), int(1000), int(contractLength), cman.account)

	loop5:
	for {
		if hashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop5	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if _,ok := cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] ; ok {
		t.Errorf("Contract 4 was found by buyer node while in the available state")
	}
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.cloneFactoryAddress, hashrateContractAddress[2], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

	// wait until hashrate contract was purchased before continuing
	loop6:
	for {
		if purchasedHashrateContractAddress[2] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop6	
		}
	}
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if cman.msg.Contracts[msgbus.ContractID(hashrateContractAddress[2].Hex())] != msgbus.ContRunningState {
		t.Errorf("Contract 4 is not in correct state")
	}

	UpdateMiningInformation(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[2], "stratum+tcp://127.0.0.1:3333/updated")
	time.Sleep(time.Millisecond * time.Duration(sleepTime*2))
	// check dest msg with associated contract was updated in msgbus
	event, err := cman.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(hashrateContractAddress[2].Hex()))
	if err != nil {
		panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Getting Purchased Contract Failed: %s", event.Err))
	}
	contractMsg := event.Data.(msgbus.Contract)
	event, err = cman.ps.GetWait(msgbus.DestMsg, msgbus.IDString(contractMsg.Dest))
	if err != nil {
		panic(fmt.Sprintf("Getting Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Getting Dest Failed: %s", event.Err))
	}
	destMsg := event.Data.(msgbus.Dest)
	if destMsg.NetUrl != "stratum+tcp://127.0.0.1:3333/updated" {
		t.Errorf("Contract 3's target dest was not updated")
	}

	//
	// Test miner hashrate being updated
	//
	// miner 1's hashrate is updated
	miner1.CurrentHashRate = 20
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	
	//check contract manager got update to miner
	if cman.miners[miner1.ID].CurrentHashRate != miner1.CurrentHashRate {
		t.Errorf("Contract Manager did not receive update to MinerID01 hashrate")
	}

	// new miner published
	miner3 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID03"),
		IP: 	"IpAddress3",
		CurrentHashRate:	20,
		State: msgbus.OnlineState,
	}
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner3.ID),miner3)
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))

	//check contract manager saw new miner and updated total hashrate
	if cman.miners[miner3.ID].CurrentHashRate != miner3.CurrentHashRate {
		t.Errorf("Contract Manager did not see new Miner 3")
	}

	// miner 2 deleted
	ps.UnpubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID))
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	if _,ok := cman.miners[miner2.ID]; ok {
		t.Errorf("Contract Manager did not remove miner 2 after it was deleted from msgbus")
	}

	//
	// Test miners are set to offline state so running contracts should close out
	//
	miner1.State = msgbus.OfflineState
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	time.Sleep(time.Millisecond * time.Duration(sleepTime/5))
	miner3.State = msgbus.OfflineState
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
	time.Sleep(time.Millisecond * time.Duration(sleepTime))

	// check contracts map is empty now
	if len(cman.msg.Contracts) != 0 {
		t.Errorf("Contracts did not closeout after all miners were set to offline")
	}
}

func TestDeployment(t *testing.T) {
	BeforeEach()
}

func TestCreateHashrateContract(t *testing.T) {
	ps := msgbus.New(10)
	var hashrateContractAddress common.Address
	contractManagerConfig, err := configurationmanager.LoadConfiguration("ropstentestconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	fmt.Println(contractManagerConfig)

	var cman SellerContractManager

	cman.init(ps, contractManagerConfig)

	// subcribe to creation events emitted by clonefactory contract 
	cfLogs, cfSub, _ := subscribeToContractEvents(cman.rpcClient, cman.cloneFactoryAddress)
	// create event signature to parse out creation event
	contractCreatedSig := []byte("contractCreated(address)")
	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
	go func() {
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				if cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex() {
					hashrateContractAddress = common.HexToAddress(cfLog.Topics[1].Hex())
					fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddress.Hex())
				}
			}
		}
	}()

	cloneFactoryAddress := common.HexToAddress(contractManagerConfig["cloneFactoryAddress"].(string))

	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cloneFactoryAddress, int(0), int(0), int(0), int(20), cman.account)
	loop:
	for {
		if hashrateContractAddress != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop
		}
	}
}

func TestPurchaseHashrateContract(t *testing.T) {
	ps := msgbus.New(10)
	contractManagerConfig, err := configurationmanager.LoadConfiguration("ropstentestconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	var cman BuyerContractManager
	err = cman.init(ps, contractManagerConfig)
	if err != nil {
		panic(fmt.Sprintf("contract manager failed:%s", err))
	}
	hashrateContractAddress := "0x50a6c6c8eC06577A8258d5F86688d0045026e18e"

	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.cloneFactoryAddress, common.HexToAddress(hashrateContractAddress), cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
}
