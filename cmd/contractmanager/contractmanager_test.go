package contractmanager

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"

	//"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"

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

func BeforeEach() (ts TestSetup) {
	var constructorParams [5]common.Address
	configaData, err := configurationmanager.LoadConfiguration("../configurationmanager/sellerconfig.json", "contractManager")
	if err != nil {
		log.Fatal(err)
	}

	var client *ethclient.Client
	client, err = setUpClient(configaData["rpcClientAddress"].(string), common.HexToAddress(configaData["nodeEthereumAddress"].(string)))
	if err != nil {
		log.Fatal(err)
	}

	ts.nodeEthereumAccount = common.HexToAddress(configaData["nodeEthereumAddress"].(string))
	ts.nodeEthereumPrivateKey = configaData["nodeEthereumPrivateKey"].(string)
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

func TestLoadMsgBusAndAPIRepo(t *testing.T) {
	ts := BeforeEach()
	ech := make(msgbus.EventChan)
	ps := msgbus.New(1)
	var api externalapi.APIRepos
	api.InitializeJSONRepos()
	var contractMsgs [5]msgbus.Contract
	var hashrateContractAddresses [5]common.Address
	contractAddr := make(chan common.Address, 5)
	stop := make(chan bool)

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := subscribeToContractEvents(ts.rpcClient, ts.cloneFactoryAddress)

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
				contractAddr <- hashrateContractAddresses[i]
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
			contractValues := readHashrateContract(ts.rpcClient, address)
			fmt.Printf("%+v\n", contractValues)

			if contractValues.State != 0 || contractValues.Price != int(i*5) || contractValues.Limit != int(i*10) || contractValues.Speed != int(i*20) ||
				contractValues.Length != int(i*40) || contractValues.Seller != ts.nodeEthereumAccount {
				t.Errorf("Read contract values not equal to expected values")
			}

			// push read in contract values into message bus contract struct
			contractMsgs[i] = createContractMsg(address, contractValues, true)
			ps.Pub(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), msgbus.Contract{})
			ps.Sub(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), ech)
			ps.Set(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), contractMsgs[i])
			ps.Get(msgbus.ContractMsg, msgbus.IDString(contractMsgs[i].ID), ech)

			// push read in contract values into API repo
			api.Contract.AddContractFromMsgBus(contractMsgs[i])

			// confirm contract values were pushed to API
			fmt.Printf("API Contract Repo: %+v\n\n", api.Contract.ContractJSONs[i])
			if len(api.Contract.ContractJSONs) != i+1 {
				t.Errorf("Contract struct not added")
			}
			i++
			if i == 5 { // all created contracts are read
				stop <- true // stop event reading routine and continue test
				break loop
			}
		}
	}()

	// create 5 new Hashrate contracts with arbitrary filled out parameters
	for i := 0; i < 5; i++ {
		CreateHashrateContract(ts.rpcClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, ts.webFacingAddress, int(i*5), int(i*10), int(i*20), int(i*40), 0)
	}

	<-stop
	close(contractAddr)
	close(stop)

	// subcribe to events emitted by webfacing to read purchase event
	wLogs, wSub := subscribeToContractEvents(ts.rpcClient, ts.webFacingAddress)

	// purchase 1st created Hashrate contract to fill out rest of contract parameters and emit purchase event
	PurchaseHashrateContract(ts.rpcClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, ts.webFacingAddress,
		hashrateContractAddresses[0], ts.nodeEthereumAccount, "IP Address")

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
	purchasedContractValues := readHashrateContract(ts.rpcClient, hashrateContractAddresses[0])
	if purchasedContractValues.State != 0 || purchasedContractValues.Price != int(0) || purchasedContractValues.Limit != int(0) ||
		purchasedContractValues.Speed != int(0) || purchasedContractValues.Length != int(0) ||
		purchasedContractValues.Buyer != ts.nodeEthereumAccount || purchasedContractValues.Seller != ts.nodeEthereumAccount {
		t.Errorf("Read contract values from purchased contract not equal to expected values")
	}

	destUrl := readDestUrl(ts.rpcClient, hashrateContractAddresses[0])
	if destUrl != "IP Address" {
		t.Errorf("Read contract values from purchased contract not equal to expected values")
	}
	destMsg := msgbus.Dest{
		ID:     msgbus.DestID("DestID1"),
		NetUrl: msgbus.DestNetUrl(destUrl),
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(destMsg.ID), destMsg)

	// update msgbus struct and contract API repo for contract with new values
	purchasedContractMsg := createContractMsg(hashrateContractAddresses[0], purchasedContractValues, true)
	purchasedContractMsg.Dest = destMsg.ID
	ps.Set(msgbus.ContractMsg, msgbus.IDString(purchasedContractMsg.ID), purchasedContractMsg)
	ps.Get(msgbus.ContractMsg, msgbus.IDString(purchasedContractMsg.ID), ech)

	contractJSON := msgdata.ConvertContractMSGtoContractJSON(purchasedContractMsg)
	api.Contract.UpdateContract(api.Contract.ContractJSONs[0].ID, contractJSON)
	fmt.Printf("API Contract Repo: %+v\n\n", api.Contract.ContractJSONs[0])
}

func TestCreateUnsignedTransaction(t *testing.T) {
	ts := BeforeEach()
	var hashrateContractAddress common.Address

	// subcribe to events emitted by clonefactory contract to read contract creation event
	cfLogs, cfSub := subscribeToContractEvents(ts.rpcClient, ts.cloneFactoryAddress)

	CreateHashrateContract(ts.rpcClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, ts.webFacingAddress, int(0), int(0), int(30), int(0), int(0))

	select {
	case err := <-cfSub.Err():
		log.Fatal(err)
	case cfLog := <-cfLogs:
		hashrateContractAddress = common.HexToAddress(cfLog.Topics[1].Hex())
		fmt.Printf("Log Block Number: %d\n", cfLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", cfLog.Index)
		fmt.Printf("Address of created Hashrate Contract: %s\n\n", hashrateContractAddress.Hex())
	}

	PurchaseHashrateContract(ts.rpcClient, ts.nodeEthereumAccount, ts.nodeEthereumPrivateKey, ts.webFacingAddress,
		hashrateContractAddress, ts.nodeEthereumAccount, "IpAddress")

	contractValues := readHashrateContract(ts.rpcClient, hashrateContractAddress)
	fmt.Println("Contract State before closeout: ", contractValues.State)

	hashrateContractABI, err := abi.JSON(strings.NewReader(implementation.ImplementationABI))
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := ts.rpcClient.PendingNonceAt(context.Background(), ts.nodeEthereumAccount)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := ts.rpcClient.NetworkID(context.Background())
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
	
	privateKey, err := crypto.HexToECDSA(ts.nodeEthereumPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	signer := types.NewEIP155Signer(chainID)

	unsignedTxHash := signer.Hash(unsignedTx)
	fmt.Println("Unsigned Transaction Hash: ", unsignedTxHash)

	sig, err := crypto.Sign(unsignedTxHash[:], privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Signature: ", sig)

	/*
		Create raw transaction hash from signed tx object
	*/
	signedTx,err := unsignedTx.WithSignature(signer, sig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Signed Tx: ", signedTx)

	rawTxBytes,err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Raw tx bytes: ", rawTxBytes)

	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Println("Raw tx hex: ", rawTxHex)

	/*
		Decode raw trasaction hash into tx object to send transaction
	*/
	decodedRawTxBytes,err := hex.DecodeString(rawTxHex)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decoded raw tx bytes: ", decodedRawTxBytes)

	tx := new(types.Transaction)
	rlp.DecodeBytes(decodedRawTxBytes, &tx)
	fmt.Println("Decoded Transaction: ", tx)

	err = ts.rpcClient.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	contractValues = readHashrateContract(ts.rpcClient, hashrateContractAddress)
	fmt.Println("Contract State after closeout: ", contractValues.State)
}
func TestSellerRoutine(t *testing.T) {
	ps := msgbus.New(10)
	ts := BeforeEach()
	var hashrateContractAddress [4]common.Address
	// hashrateContractAddress[0] = common.HexToAddress("0xa5e6cd816545c883bfa246e96bf7d3648d84d881")
	// hashrateContractAddress[1] = common.HexToAddress("0xbb05218023c62fe691bb78b3969eab50077b6a07")

	contractManagerConfig, err := configurationmanager.LoadConfiguration("../configurationmanager/sellerconfig.json", "contractManager")
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
	time.Sleep(time.Millisecond * 10000)

	err = cman.start()
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}

	time.Sleep(time.Millisecond * 10000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, ts.webFacingAddress, hashrateContractAddress[0], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
	time.Sleep(time.Millisecond * 10000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[0])
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
	setContractCloseOut(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[2])

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

	contractManagerConfig, err := configurationmanager.LoadConfiguration("../configurationmanager/sellerconfig.json", "contractManager")
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

	// miner1 := msgbus.Miner {
	// 	ID:		msgbus.MinerID("MinerID01"),
	// 	IP: 	"IpAddress1",
	// 	CurrentHashRate:	30,
	// 	State: msgbus.OnlineState,
	// }
	// miner2 := msgbus.Miner {
	// 	ID:		msgbus.MinerID("MinerID02"),
	// 	IP: 	"IpAddress2",
	// 	CurrentHashRate:	20,
	// 	State: msgbus.OnlineState,
	// }
	// ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner1.ID),miner1)
	// ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner2.ID),miner2)

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

	time.Sleep(time.Millisecond * 5000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[0], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
	//PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[1], cman.account, "IpAddress2|8888|ryan")

	time.Sleep(time.Millisecond * 5000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[0])
	//SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[1])

	err = cman.start()
	if err != nil {
		panic(fmt.Sprintf("contract manager failed to start:%s", err))
	}

	time.Sleep(time.Millisecond * 20000)

	// miner hashrate fall below promised hashrate
	// miner1.CurrentHashRate = 20
	// ps.Set(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)

	time.Sleep(time.Millisecond * 5000)
	// miner3 := msgbus.Miner {
	// 	ID:		msgbus.MinerID("MinerID03"),
	// 	IP: 	"IpAddress3",
	// 	CurrentHashRate:	30,
	// 	State: msgbus.OnlineState,
	// }
	// ps.Pub(msgbus.MinerMsg,msgbus.IDString(miner3.ID),miner3)

	time.Sleep(time.Millisecond * 10000)
	CreateHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, int(0), int(0), int(30), int(10), int(0))

	time.Sleep(time.Millisecond * 5000)
	PurchaseHashrateContract(cman.rpcClient, cman.account, cman.privateKey, cman.webFacingAddress, hashrateContractAddress[1], cman.account, "stratum+tcp://127.0.0.1:3333/testrig")

	time.Sleep(time.Millisecond * 5000)
	SetFundContract(cman.rpcClient, cman.account, cman.privateKey, hashrateContractAddress[1])

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