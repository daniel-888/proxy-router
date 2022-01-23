package contractmanager

import (
	"context"
	//"crypto/ecdsa"
	//"crypto/rand"
	//"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	//"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/lumerin/cmd/config"
	"gitlab.com/TitanInd/lumerin/lumerinlib"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/clonefactory"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/lumerintoken"
)

type TestSetup struct {
	ethClient              *ethclient.Client
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
	auth.GasLimit = uint64(6000000) // in units
	auth.GasPrice = gasPrice

	lmnAddress := constructorParams[0]
	validatorAddress := constructorParams[1]
	proxyAddress := constructorParams[2]

	switch contract {
	case "LumerinToken":
		address, _, _, err := lumerintoken.DeployLumerintoken(auth, client)
		if err != nil {
			log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
		}
		return address
	case "CloneFactory":
		address, _, _, err := clonefactory.DeployClonefactory(auth, client, lmnAddress, validatorAddress, proxyAddress)
		if err != nil {
			log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
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

func UpdateCipherText(client *ethclient.Client,
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
	fmt.Printf("Hashrate Contract %s Cipher Text Updated \n\n", contractAddress)
}

func createNewGanacheBlock(ts TestSetup, account common.Address, privateKey string, contractLength int, sleepTime int) {
	time.Sleep(time.Second * time.Duration(contractLength))
		
	nonce, err := ts.ethClient.PendingNonceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := ts.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(3000000)

	unsignedTx := types.NewTransaction(nonce, account, nil, gasLimit, gasPrice, nil)

	//Sign transaction
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, privateKeyECDSA)
	if err != nil {
		log.Fatal(err)
	}

	err = ts.ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Millisecond * time.Duration(sleepTime))
}

func BeforeEach() (ts TestSetup) {
	var constructorParams [5]common.Address
	configData, err := config.LoadConfiguration("contractManager")
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	mnemonic := configData["mnemonic"].(string)
	account, privateKey := hdWalletKeys(mnemonic, 0)
	ts.nodeEthereumAccount = account.Address
	ts.nodeEthereumPrivateKey = privateKey

	fmt.Println("Contract Manager account", ts.nodeEthereumAccount)
	fmt.Println("Contract Manager key", ts.nodeEthereumPrivateKey)

	var client *ethclient.Client
	client, err = setUpClient(configData["ethNodeAddr"].(string), ts.nodeEthereumAccount)
	if err != nil {
		log.Fatalf("Funcname::%s, Fileline::%s, Error::%v", lumerinlib.Funcname(), lumerinlib.FileLine(), err)
	}

	ts.ethClient = client
	ts.validatorAddress = common.HexToAddress(configData["validatorAddress"].(string)) // dummy address
	ts.proxyAddress = common.HexToAddress(configData["proxyAddress"].(string))         // dummy address
	ts.lumerinAddress = DeployContract(ts.ethClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, constructorParams, "LumerinToken")
	fmt.Println("Lumerin Token Contract Address: ", ts.lumerinAddress)

	constructorParams[0] = ts.lumerinAddress
	constructorParams[1] = ts.validatorAddress
	constructorParams[2] = ts.proxyAddress
	
	ts.cloneFactoryAddress = DeployContract(ts.ethClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, constructorParams, "CloneFactory")
	fmt.Println("Clone Factory Contract Address: ", ts.cloneFactoryAddress)

	return ts
}

// func TestDeployment(t *testing.T) {
// 	BeforeEach()
// }

// func TestCreateHashrateContract(t *testing.T) {
// 	ps := msgbus.New(10)
// 	var hashrateContractAddress common.Address
// 	contractManagerConfig, err := configurationmanager.LoadConfiguration("lumerinconfig.json", "contractManager")
// 	if err != nil {
// 		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
// 	}

// 	fmt.Println(contractManagerConfig)

// 	var cman SellerContractManager

// 	cman.init(ps, contractManagerConfigID)

// 	// subcribe to creation events emitted by clonefactory contract 
// 	cfLogs, cfSub, _ := subscribeToContractEvents(cman.ethClient, cman.cloneFactoryAddress)
// 	// create event signature to parse out creation event
// 	contractCreatedSig := []byte("contractCreated(address)")
// 	contractCreatedSigHash := crypto.Keccak256Hash(contractCreatedSig)
// 	go func() {
// 		for {
// 			select {
// 			case err := <-cfSub.Err():
// 				log.Fatal(err)
// 			case cfLog := <-cfLogs:
// 				if cfLog.Topics[0].Hex() == contractCreatedSigHash.Hex() {
// 					hashrateContractAddress = common.HexToAddress(cfLog.Topics[1].Hex())
// 					fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddress.Hex())
// 				}
// 			}
// 		}
// 	}()

// 	cloneFactoryAddress := common.HexToAddress(contractManagerConfig["cloneFactoryAddress"].(string))

// 	CreateHashrateContract(cman.ethClient, cman.account, cman.privateKey, cloneFactoryAddress, int(0), int(0), int(0), int(2000), cman.account)
// 	loop:
// 	for {
// 		if hashrateContractAddress != common.HexToAddress("0x0000000000000000000000000000000000000000") {
// 			break loop
// 		}
// 	}
// }

// func TestPurchaseHashrateContract(t *testing.T) {
// 	ps := msgbus.New(10)
// 	contractManagerConfig, err := configurationmanager.LoadConfiguration("lumerinconfig.json", "contractManager")
// 	if err != nil {
// 		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
// 	}

// 	var cman BuyerContractManager
// 	err = cman.init(ps, contractManagerConfig)
// 	if err != nil {
// 		panic(fmt.Sprintf("contract manager failed:%s", err))
// 	}
// 	hashrateContractAddress := "0x50a6c6c8eC06577A8258d5F86688d0045026e18e"

// 	PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, cman.cloneFactoryAddress, common.HexToAddress(hashrateContractAddress), cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
// }