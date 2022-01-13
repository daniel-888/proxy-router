package externalapi

import (
	"fmt"
	"testing"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func TestMsgBusDataAddedToApiRepos(t *testing.T) {
	ps := msgbus.New(10)

	dest := msgbus.Dest{
		ID:   		msgbus.DestID("DestID01"),
		NetUrl: 	msgbus.DestNetUrl("stratum+tcp://127.0.0.1:3334/"),	
	}
	seller := msgbus.Seller{
		ID:                    msgbus.SellerID("SellerID01"),
		DefaultDest:           dest.ID,
		TotalAvailableHashRate: 0,
		UnusedHashRate:         0,
	}
	buyer := msgbus.Buyer{
		ID: 			msgbus.BuyerID("BuyerID01"),
		DefaultDest: 	dest.ID,
	}
	contract := msgbus.Contract{
		IsSeller: 				true,	
		ID:						msgbus.ContractID("ContractID01"),
		State: 					msgbus.ContRunningState,
		Buyer: 					buyer.ID,
		Price: 					100,
		Limit: 					100,
		Speed: 					100,
		Length: 				100,
		StartingBlockTimestamp: 100,
		Dest:					dest.ID,
	}
	seller.Contracts = map[msgbus.ContractID]msgbus.ContractState{
		contract.ID: msgbus.ContRunningState,
	}
	buyer.Contracts = map[msgbus.ContractID]msgbus.ContractState{
		contract.ID: msgbus.ContRunningState,
	}
	config := msgbus.ConfigInfo{
		ID:          msgbus.ConfigID("ConfigID01"),
		DefaultDest: dest.ID,
		Seller:      seller.ID,
	}
	miner := msgbus.Miner{
		ID:						msgbus.MinerID("MinerID01"),
		State: 					msgbus.OnlineState,
		Seller:   				seller.ID,
		Dest:					dest.ID,	
		InitialMeasuredHashRate: 10000,
		CurrentHashRate:         9000,

	}
	connection := msgbus.Connection{
		ID:        				msgbus.ConnectionID("ConnectionID01"),
		Miner:    				miner.ID,
		Dest:      				dest.ID,
		State:     				msgbus.ConnAuthState,
		TotalHash: 				10000,
		StartDate: 				time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC),
	}

	var api APIRepos
	api.InitializeJSONRepos(ps)
	time.Sleep(time.Millisecond*1000)
	go api.RunAPI()
	
	fmt.Print("\n/// Publish Msgbus Msgs //\n\n")
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(dest.ID), msgbus.Dest{})
	time.Sleep(time.Millisecond*100)
	if api.Dest.DestJSONs[0].ID != string(dest.ID) {
		t.Errorf("Failed to add dest to API repo after publish from msgbus")
	}
	ps.PubWait(msgbus.BuyerMsg, msgbus.IDString(buyer.ID), msgbus.Buyer{})
	time.Sleep(time.Millisecond*100)
	if api.Buyer.BuyerJSONs[0].ID != string(buyer.ID) {
		t.Errorf("Failed to add buyer to API repo after publish from msgbus")
	}
	ps.PubWait(msgbus.SellerMsg, msgbus.IDString(seller.ID), msgbus.Seller{})
	time.Sleep(time.Millisecond*100)
	if api.Seller.SellerJSONs[0].ID != string(seller.ID) {
		t.Errorf("Failed to add seller to API repo after publish from msgbus")
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract.ID), msgbus.Contract{})
	time.Sleep(time.Millisecond*100)
	if api.Contract.ContractJSONs[0].ID != string(contract.ID) {
		t.Errorf("Failed to add contract to API repo after publish from msgbus")
	}
	ps.PubWait(msgbus.ConfigMsg, msgbus.IDString(config.ID), msgbus.ConfigInfo{})
	time.Sleep(time.Millisecond*100)
	if api.Config.ConfigInfoJSONs[0].ID != string(config.ID) {
		t.Errorf("Failed to add config to API repo after publish from msgbus")
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner.ID), msgbus.Miner{})
	time.Sleep(time.Millisecond*100)
	if api.Miner.MinerJSONs[0].ID != string(miner.ID) {
		t.Errorf("Failed to add miner to API repo after publish from msgbus")
	}
	ps.PubWait(msgbus.ConnectionMsg, msgbus.IDString(connection.ID), msgbus.Connection{})
	time.Sleep(time.Millisecond*100)
	if api.Connection.ConnectionJSONs[0].ID != string(connection.ID) {
		t.Errorf("Failed to add connection to API repo after publish from msgbus")
	}
	fmt.Print("\nAPI Repos::\n")
	fmt.Println("Dest Repo: ", api.Dest.DestJSONs)
	fmt.Println("Buyer Repo: ", api.Buyer.BuyerJSONs)
	fmt.Println("Seller Repo: ", api.Seller.SellerJSONs)
	fmt.Println("Contract Repo: ", api.Contract.ContractJSONs)
	fmt.Println("Config Repo: ", api.Config.ConfigInfoJSONs)
	fmt.Println("Miner Repo: ", api.Miner.MinerJSONs)
	fmt.Println("Connection Repo: ", api.Connection.ConnectionJSONs)

	fmt.Print("\n/// Update Msgbus Msgs //\n\n")
	ps.SetWait(msgbus.DestMsg, msgbus.IDString(dest.ID), dest)
	time.Sleep(time.Millisecond*100)
	if api.Dest.DestJSONs[0] != msgdata.ConvertDestMSGtoDestJSON(dest) {
		t.Errorf("Failed to update dest in API repo after update from msgbus")
	}
	ps.SetWait(msgbus.BuyerMsg, msgbus.IDString(buyer.ID), buyer)
	time.Sleep(time.Millisecond*100)
	if api.Buyer.BuyerJSONs[0].DefaultDest != msgdata.ConvertBuyerMSGtoBuyerJSON(buyer).DefaultDest {
		t.Errorf("Failed to update buyer in API repo after update from msgbus")
	}
	ps.SetWait(msgbus.SellerMsg, msgbus.IDString(seller.ID), seller)
	time.Sleep(time.Millisecond*100)
	if api.Seller.SellerJSONs[0].DefaultDest != msgdata.ConvertSellerMSGtoSellerJSON(seller).DefaultDest {
		t.Errorf("Failed to update seller in API repo after update from msgbus")
	}
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract.ID), contract)
	time.Sleep(time.Millisecond*100)
	if api.Contract.ContractJSONs[0] != msgdata.ConvertContractMSGtoContractJSON(contract) {
		t.Errorf("Failed to update contract in API repo after update from msgbus")
	}
	ps.SetWait(msgbus.ConfigMsg, msgbus.IDString(config.ID), config)
	time.Sleep(time.Millisecond*100)
	if api.Config.ConfigInfoJSONs[0] != msgdata.ConvertConfigInfoMSGtoConfigInfoJSON(config) {
		t.Errorf("Failed to update config in API repo after update from msgbus")
	}
	ps.SetWait(msgbus.MinerMsg, msgbus.IDString(miner.ID), miner)
	time.Sleep(time.Millisecond*100)
	if api.Miner.MinerJSONs[0] != msgdata.ConvertMinerMSGtoMinerJSON(miner) {
		t.Errorf("Failed to update miner in API repo after update from msgbus")
	}
	ps.SetWait(msgbus.ConnectionMsg, msgbus.IDString(connection.ID), connection)
	time.Sleep(time.Millisecond*100)
	if api.Connection.ConnectionJSONs[0] != msgdata.ConvertConnectionMSGtoConnectionJSON(connection) {
		t.Errorf("Failed to update connection in API repo after update from msgbus")
	}
	fmt.Print("\nAPI Repos::\n")
	fmt.Println("Dest Repo: ", api.Dest.DestJSONs)
	fmt.Println("Buyer Repo: ", api.Buyer.BuyerJSONs)
	fmt.Println("Seller Repo: ", api.Seller.SellerJSONs)
	fmt.Println("Contract Repo: ", api.Contract.ContractJSONs)
	fmt.Println("Config Repo: ", api.Config.ConfigInfoJSONs)
	fmt.Println("Miner Repo: ", api.Miner.MinerJSONs)
	fmt.Println("Connection Repo: ", api.Connection.ConnectionJSONs)

	time.Sleep(time.Minute/6)
	fmt.Print("\n/// UnPublish Msgbus Msgs //\n\n")
	ps.UnpubWait(msgbus.DestMsg, msgbus.IDString(dest.ID))
	time.Sleep(time.Millisecond*100)
	if len(api.Dest.DestJSONs) > 0 {
		t.Errorf("Failed to remove dest from API repo after unpublish from msgbus")
	}
	ps.UnpubWait(msgbus.BuyerMsg, msgbus.IDString(buyer.ID))
	time.Sleep(time.Millisecond*100)
	if len(api.Buyer.BuyerJSONs) > 0 {
		t.Errorf("Failed to remove buyer from API repo after unpublish from msgbus")
	}
	ps.UnpubWait(msgbus.SellerMsg, msgbus.IDString(seller.ID))
	time.Sleep(time.Millisecond*100)
	if len(api.Seller.SellerJSONs) > 0 {
		t.Errorf("Failed to remove seller from API repo after unpublish from msgbus")
	}
	ps.UnpubWait(msgbus.ContractMsg, msgbus.IDString(contract.ID))
	time.Sleep(time.Millisecond*100)
	if len(api.Contract.ContractJSONs) > 0 {
		t.Errorf("Failed to remove contract from API repo after unpublish from msgbus")
	}
	ps.UnpubWait(msgbus.ConfigMsg, msgbus.IDString(config.ID))
	time.Sleep(time.Millisecond*100)
	if len(api.Config.ConfigInfoJSONs) > 0 {
		t.Errorf("Failed to remove config from API repo after unpublish from msgbus")
	}
	ps.UnpubWait(msgbus.MinerMsg, msgbus.IDString(miner.ID))
	time.Sleep(time.Millisecond*100)
	if len(api.Miner.MinerJSONs) > 0 {
		t.Errorf("Failed to remove miner from API repo after unpublish from msgbus")
	}
	ps.UnpubWait(msgbus.ConnectionMsg, msgbus.IDString(connection.ID))
	time.Sleep(time.Millisecond*100)
	if len(api.Connection.ConnectionJSONs) > 0 {
		t.Errorf("Failed to remove connection from API repo after unpublish from msgbus")
	}
	fmt.Print("\nAPI Repos::\n")
	fmt.Println("Dest Repo: ", api.Dest.DestJSONs)
	fmt.Println("Buyer Repo: ", api.Buyer.BuyerJSONs)
	fmt.Println("Seller Repo: ", api.Seller.SellerJSONs)
	fmt.Println("Contract Repo: ", api.Contract.ContractJSONs)
	fmt.Println("Config Repo: ", api.Config.ConfigInfoJSONs)
	fmt.Println("Miner Repo: ", api.Miner.MinerJSONs)
	fmt.Println("Connection Repo: ", api.Connection.ConnectionJSONs)
	time.Sleep(time.Minute/6)
}
