package contractmanager

import (
	"fmt"
	"testing"
	"context"

	"github.com/ethereum/go-ethereum/common"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func TestPurchaseHashrateContract(t *testing.T) {
	ps := msgbus.New(10)
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

	var cman BuyerContractManager
	err = cman.init(&contractManagerCtx, ps, contractManagerConfigID)
	if err != nil {
		panic(fmt.Sprintf("contract manager failed:%s", err))
	}
	hashrateContractAddress := "0x50a6c6c8eC06577A8258d5F86688d0045026e18e"

	PurchaseHashrateContract(cman.ethClient, cman.account, cman.privateKey, cman.cloneFactoryAddress, common.HexToAddress(hashrateContractAddress), cman.account, "stratum+tcp://127.0.0.1:3333/testrig")
}