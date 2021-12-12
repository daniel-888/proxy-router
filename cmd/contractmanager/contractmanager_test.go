package contractmanager

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"
	"crypto/ecdsa"
	"crypto/rand"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"

	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/clonefactory"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/implementation"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/ledger"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/lumerintoken"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager/contractartifacts/webfacing"
)

type TestSetup struct {
	rpcClient              *ethclient.Client
	nodeEthereumPrivateKey string
	nodeEthereumAccount    common.Address
	validatorAddress       common.Address
	proxyAddress           common.Address
	lumerinAddress         common.Address
	cloneFactoryAddress    common.Address
	ledgerAddress          common.Address
	webFacingAddress       common.Address
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

	time.Sleep(time.Millisecond * 700)
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

	time.Sleep(time.Millisecond * 700)
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

	time.Sleep(time.Millisecond * 700)
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

	time.Sleep(time.Millisecond * 700)
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

func encryptData(msg string, privkey string) []byte {
	message := []byte(msg)
	privateKey, err := crypto.HexToECDSA(privkey)
    if err != nil {
        log.Fatal(err)
    }
    publicKey := privateKey.Public()
	pubKey, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
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
	configaData, err := configurationmanager.LoadConfiguration("lumerinconfig.json", "contractManager")
	if err != nil {
		log.Fatal(err)
	}

	var client *ethclient.Client
	client, err = setUpClient(configaData["rpcClientAddress"].(string), common.HexToAddress(configaData["sellerEthereumAddress"].(string)))
	if err != nil {
		log.Fatal(err)
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
	ts.ledgerAddress = DeployContract(ts.rpcClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, constructorParams, "Ledger")
	fmt.Println("Ledger Contract Address: ", ts.ledgerAddress)

	constructorParams[3] = ts.ledgerAddress
	constructorParams[4] = ts.cloneFactoryAddress

	ts.webFacingAddress = DeployContract(ts.rpcClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, constructorParams, "WebFacing")
	fmt.Println("Web Facing Contract Address: ", ts.webFacingAddress)

	return ts
}

func TestSellerRoutine(t *testing.T) {
	ps := msgbus.New(10)
	ts := BeforeEach()
	var hashrateContractAddress [4]common.Address
	// hashrateContractAddress[0] = common.HexToAddress("0xa5e6cd816545c883bfa246e96bf7d3648d84d881")
	// hashrateContractAddress[1] = common.HexToAddress("0xbb05218023c62fe691bb78b3969eab50077b6a07")

	contractManagerConfig, err := configurationmanager.LoadConfiguration("lumerinconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	fmt.Println(contractManagerConfig)

	var cman SellerContractManager
	err = cman.init(ps, contractManagerConfig)
	if err != nil {
		panic(fmt.Sprintf("contract manager init failed:%s", err))
	}
	cman.cloneFactoryAddress = ts.cloneFactoryAddress
	cman.ledgerAddress = ts.ledgerAddress

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := subscribeToContractEvents(ts.rpcClient, ts.cloneFactoryAddress)
	go func() {
		i := 0
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				hashrateContractAddress[i] = common.HexToAddress(cfLog.Topics[1].Hex())
				fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddress[i].Hex())

				i++
				fmt.Println("Created Contract No: ", i)
			}
		}
	}()

	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.webFacingAddress, int(0), int(0), int(30), int(10), int(0))
	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.webFacingAddress, int(0), int(0), int(30), int(10), int(0))
	
	// 1 contract is running already at startup
	time.Sleep(time.Millisecond * 10000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.webFacingAddress, hashrateContractAddress[0], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
	time.Sleep(time.Millisecond * 10000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[0])
	err = cman.start()
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}

	//encryptedDest := encryptData("stratum+tcp://127.0.0.1:3333/testrig", cman.privateKey) 
	
	time.Sleep(time.Millisecond * 10000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.webFacingAddress, hashrateContractAddress[1], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
	time.Sleep(time.Millisecond * 10000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[1])
	time.Sleep(time.Millisecond * 10000)

	// test early closeout
	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.webFacingAddress, int(0), int(0), int(30), int(100), int(0))
	time.Sleep(time.Millisecond * 10000)

	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.webFacingAddress, hashrateContractAddress[2], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
	time.Sleep(time.Millisecond * 10000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[2])
	time.Sleep(time.Millisecond * 10000)
	var wg sync.WaitGroup
	wg.Add(1)
	setContractCloseOut(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[2], &wg, &cman.currentNonce)
	wg.Wait()

	time.Sleep(time.Millisecond * 10000)

	// test another creation event
	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.webFacingAddress, int(0), int(0), int(30), int(10), int(0))
	time.Sleep(time.Millisecond * 10000)

	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.webFacingAddress, hashrateContractAddress[3], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
	time.Sleep(time.Millisecond * 10000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[3])

	time.Sleep(time.Millisecond * 12000)

	// make new trasnaction to create new block in ganache instance
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

	waitchan := make(chan bool)
	<-waitchan
}

func TestBuyerRoutine(t *testing.T) {
	ps := msgbus.New(10)
	ts := BeforeEach()
	var hashrateContractAddress [3]common.Address

	contractManagerConfig, err := configurationmanager.LoadConfiguration("lumerinconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	var cman BuyerContractManager
	err = cman.init(ps, contractManagerConfig)
	if err != nil {
		panic(fmt.Sprintf("contract manager init failed:%s", err))
	}
	cman.webFacingAddress = ts.webFacingAddress
	cman.ledgerAddress = ts.ledgerAddress

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

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := subscribeToContractEvents(ts.rpcClient, ts.cloneFactoryAddress)
	go func() {
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

	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, int(0), int(0), int(30), int(10), int(0))
	//CreateHashrateContract(cman.rpcClient,cman.account,cman.privateKey,cman.webFacingAddress,int(0),int(0),int(30),int(10),int(0))

	// encryptedDest := encryptData("stratum+tcp://127.0.0.1:3333/testrig", cman.privateKey) 
	time.Sleep(time.Millisecond * 5000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[0], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
	//PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[1], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

	time.Sleep(time.Millisecond * 5000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[0])
	//SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[1])

	err = cman.start()
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}

	time.Sleep(time.Millisecond * 10000)

	// miner hashrate is updated
	miner1.CurrentHashRate = 20
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)

	time.Sleep(time.Millisecond * 10000)
	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, int(0), int(0), int(30), int(10), int(0))

	time.Sleep(time.Millisecond * 5000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[1], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

	time.Sleep(time.Millisecond * 5000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[1])

	time.Sleep(time.Millisecond * 5000)
	// new miner added
	miner3 := msgbus.Miner {
		ID:		msgbus.MinerID("MinerID03"),
		IP: 	"IpAddress3",
		CurrentHashRate:	20,
		State: msgbus.OnlineState,
	}
	ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner3.ID),miner3)

	time.Sleep(time.Millisecond * 5000)

	miner1.State = msgbus.OfflineState
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	miner2.State = msgbus.OfflineState
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	miner3.State = msgbus.OfflineState
	ps.Set(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)


	waitchan := make(chan bool)
	<-waitchan
}

func TestDeployment(t *testing.T) {
	BeforeEach()
}

func TestCreateHashrateContract(t *testing.T) {
	ps := msgbus.New(10)
	contractManagerConfig, err := configurationmanager.LoadConfiguration("../configurationmanager/sellerconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	fmt.Println(contractManagerConfig)

	var cman SellerContractManager

	cman.init(ps, contractManagerConfig)

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := subscribeToContractEvents(cman.rpcClient, cman.cloneFactoryAddress)
	go func() {
		for {
			select {
			case err := <-cfSub.Err():
				log.Fatal(err)
			case cfLog := <-cfLogs:
				hashrateContractAddress := common.HexToAddress(cfLog.Topics[1].Hex())
				fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddress.Hex())
			}
		}
	}()

	webFacingAddress := common.HexToAddress(contractManagerConfig["webFacingAddress"].(string))

	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, webFacingAddress, int(0), int(0), int(0), int(20), int(0))
	time.Sleep(time.Millisecond * 30000)
}

func TestPurchaseHashrateContract(t *testing.T) {
	ps := msgbus.New(10)
	contractManagerConfig, err := configurationmanager.LoadConfiguration("../configurationmanager/buyerconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	var cman BuyerContractManager
	err = cman.init(ps, contractManagerConfig)
	if err != nil {
		panic(fmt.Sprintf("contract manager failed:%s", err))
	}
	hashrateContractAddress := "0x853BEd8EE67871048fC16E6742fFaA7E01c16dCC"

	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, common.HexToAddress(hashrateContractAddress), cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
}

func TestFundHashrateContract(t *testing.T) {
	ps := msgbus.New(10)
	contractManagerConfig, err := configurationmanager.LoadConfiguration("../configurationmanager/buyerconfig.json", "contractManager")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}

	var cman BuyerContractManager
	err = cman.init(ps, contractManagerConfig)
	if err != nil {
		panic(fmt.Sprintf("contract manager failed:%s", err))
	}
	hashrateContractAddress := "0x853BEd8EE67871048fC16E6742fFaA7E01c16dCC"

	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, common.HexToAddress(hashrateContractAddress))
}