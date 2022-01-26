package contractmanager

import (
	"fmt"
	"log"
	"testing"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func TestCreateHashrateContract(t *testing.T) {
	ps := msgbus.New(10)
	var hashrateContractAddress common.Address
	contractManagerCtx := context.Background()
	var contractManagerConfig msgbus.ContractManagerConfig
	contractManagerConfigID := msgbus.GetRandomIDString()

	contractManagerConfigFile, err := LoadTestConfiguration("contractManager", "../../ropstenconfig.json")
	if err != nil {
		panic(fmt.Sprintf("failed to load contract manager configuration:%s", err))
	}
	contractManagerConfig.Mnemonic = contractManagerConfigFile["mnemonic"].(string)
	contractManagerConfig.AccountIndex = int(contractManagerConfigFile["accountIndex"].(float64))
	contractManagerConfig.EthNodeAddr = contractManagerConfigFile["ethNodeAddr"].(string)
	contractManagerConfig.ClaimFunds = contractManagerConfigFile["claimFunds"].(bool)
	contractManagerConfig.CloneFactoryAddress = contractManagerConfigFile["cloneFactoryAddress"].(string)
	
	ps.PubWait(msgbus.ContractManagerConfigMsg, contractManagerConfigID, contractManagerConfig)

	var cman SellerContractManager
	cman.init(&contractManagerCtx, ps, contractManagerConfigID)

	// subcribe to creation events emitted by clonefactory contract 
	cfLogs, cfSub, _ := subscribeToContractEvents(cman.ethClient, cman.cloneFactoryAddress)
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

	cloneFactoryAddress := common.HexToAddress(contractManagerConfig.CloneFactoryAddress)

	CreateHashrateContract(cman.ethClient, cman.account, cman.privateKey, cloneFactoryAddress, int(0), int(0), int(0), int(2000), cman.account)
	loop:
	for {
		if hashrateContractAddress != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			break loop
		}
	}
}