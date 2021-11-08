package externalapi

import (
	"encoding/json"
	"testing"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func TestMsgBusDataAddedToApiRepos(t *testing.T) {
	ech := make(msgbus.EventChan)
	ps := msgbus.New(1)

	dest := msgbus.Dest{
		ID:   		"DestID01",
		NetHost:   	"127.0.0.1",
		NetPort: 	"80",
		NetProto:   "tcp",
	}
	config := msgbus.ConfigInfo{
		ID:          "ConfigID01",
		DefaultDest: "DestID01",
		Seller:      "SellerID01",
	}
	seller := msgbus.Seller{
		ID:                     "SellerID01",
		DefaultDest:            "DestID01",
		TotalAvailableHashRate: 0,
		UnusedHashRate:         0,
	}
	seller.NewContracts = map[msgbus.ContractID]bool{
		"0x85A256C5688D012263D5A79EE37E84FC35EC4524": true,
        "0x89921E8D51D22252D64EA34340A4161696887271": false,
        "0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3": true,
	}
	seller.CompleteContracts = map[msgbus.ContractID]bool{
		"0x50937C047DB93CB5C87F65B6EFFEA47D03DF0F7D": true,
        "0xFB610E4C269DA110C97B92F5F34EAA50E5F3D500": false,
        "0x397729E80F77BA09D930FE24E8D1FC74372E86D3": true,
	}   
    seller.ActiveContracts = map[msgbus.ContractID]bool{
		"0x9F252E1EC723AF6D96A36B4EB2B75A262291497C": true,
        "0xBB2EAAAAA9B08EC320FC984D7D19E28835DD94DD": false,
        "0x407E8A225658FEE384859874952E2BBC11E98B5C": true,
	}
	contract := msgbus.Contract{
		ID:				"ContractID01",
		State: 			msgbus.ContActiveState,
		Buyer: 			"Buyer ID01",
		Price: 			100,
		Limit: 			100,
		Speed: 			100,
		Length: 		100,
		ValidationFee:	100,
		StartingBlockTimestamp: 100,
	}
	
	miner := msgbus.Miner{
		ID:						"MinerID01",
		State: 					msgbus.OnlineState,
		Seller:   				"SellerID01",
		Dest:					"DestID01",	
		InitialMeasuredHashRate: 10000,
		CurrentHashRate:         9000,

	}
	connection := msgbus.Connection{
		ID:        				"ConnectionID01",
		Miner:    				"MinerID01",
		Dest:      				"DestID01",
		State:     				msgbus.ConnAuthState,
		TotalHash: 				10000,
		StartDate: 				time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC),
	}

	configRepo, connectionRepo, contractRepo, destRepo, minerRepo, sellerRepo := InitializeJSONRepos()

	go func(ech msgbus.EventChan) {
		for e := range ech {
			switch e.Msg {
			case msgbus.ConfigMsg:
				configRepo.AddConfigInfoFromMsgBus(config)
				if configRepo.ConfigInfoJSONs[0].ID != "ConfigID01" {
					t.Errorf("Failed to add Config to Repo")
				} 
			case msgbus.DestMsg:
				destRepo.AddDestFromMsgBus(dest)
				if destRepo.DestJSONs[0].ID != "DestID01" {
					t.Errorf("Failed to add Dest to Repo")
				} 
			case msgbus.SellerMsg:
				sellerRepo.AddSellerFromMsgBus(seller)
				if sellerRepo.SellerJSONs[0].ID != "SellerID01" {
					t.Errorf("Failed to add Seller to Repo")
				} 
			case msgbus.ContractMsg:
				contractRepo.AddContractFromMsgBus(contract)
				if contractRepo.ContractJSONs[0].ID != "ContractID01" {
					t.Errorf("Failed to add Contract to Repo")
				} 
			case msgbus.MinerMsg:
				minerRepo.AddMinerFromMsgBus(miner)
				if minerRepo.MinerJSONs[0].ID != "MinerID01" {
					t.Errorf("Failed to add Miner to Repo")
				} 
			case msgbus.ConnectionMsg:
				connectionRepo.AddConnectionFromMsgBus(connection)
				if connectionRepo.ConnectionJSONs[0].ID != "ConnectionID01" {
					t.Errorf("Failed to add Connection to Repo")
				} 
			default:
			
			} 
		}
	}(ech)

	ps.Pub(msgbus.ConfigMsg, "configMsg01", msgbus.ConfigInfo{})
	ps.Pub(msgbus.DestMsg, "destMsg01", msgbus.Dest{})
	ps.Pub(msgbus.SellerMsg, "sellerMsg01", msgbus.Seller{})
	ps.Pub(msgbus.ContractMsg, "contractMsg01", msgbus.Contract{})
	ps.Pub(msgbus.MinerMsg, "minerMsg01", msgbus.Miner{})
	ps.Pub(msgbus.ConnectionMsg, "connectionMsg01", msgbus.Connection{})

	ps.Sub(msgbus.ConfigMsg, "configMsg01", ech)
	ps.Sub(msgbus.DestMsg, "destMsg01", ech)
	ps.Sub(msgbus.SellerMsg, "sellerMsg01", ech)
	ps.Sub(msgbus.ContractMsg, "contractMsg01", ech)
	ps.Sub(msgbus.MinerMsg, "minerMsg01", ech)
	ps.Sub(msgbus.ConnectionMsg, "connectionMsg01", ech)

	ps.Set(msgbus.ConfigMsg, "configMsg01", config)
	ps.Set(msgbus.DestMsg, "destMsg01", dest)
	ps.Set(msgbus.SellerMsg, "sellerMsg01", seller)
	ps.Set(msgbus.ContractMsg, "contractMsg01", contract)
	ps.Set(msgbus.MinerMsg, "minerMsg01", miner)
	ps.Set(msgbus.ConnectionMsg, "connectionMsg01", connection)
}

func TestMockPOSTAddedToMsgBus(t *testing.T) {	
	// Mock POST Requests by declaring new JSON structures and adding them to api repos
	eaConfig,err := configurationmanager.LoadConfiguration("../configurationmanager/testconfig.json", "externalAPI")
	if err != nil {
		t.Errorf("LoadConfiguration returned error")
	}

	dest := eaConfig["dest"].(map[string]interface{})
	destMarshaled,_ := json.Marshal(dest)
	destJSON := msgdata.DestJSON {}
	json.Unmarshal(destMarshaled, &destJSON)

	config := eaConfig["config"].(map[string]interface{})
	configMarshaled,_ := json.Marshal(config)
	configJSON := msgdata.ConfigInfoJSON {}
	json.Unmarshal(configMarshaled, &configJSON)

	connection := eaConfig["connection"].(map[string]interface{})
	connectionMarshaled,_ := json.Marshal(connection)
	connectionJSON := msgdata.ConnectionJSON {}
	json.Unmarshal(connectionMarshaled, &connectionJSON)

	contract := eaConfig["contract"].(map[string]interface{})
	contractMarshaled,_ := json.Marshal(contract)
	contractJSON := msgdata.ContractJSON {}
	json.Unmarshal(contractMarshaled, &contractJSON)

	miner := eaConfig["miner"].(map[string]interface{})
	minerMarshaled,_ := json.Marshal(miner)
	minerJSON := msgdata.MinerJSON {}
	json.Unmarshal(minerMarshaled, &minerJSON)

	seller := eaConfig["seller"].(map[string]interface{})
	sellerMarshaled,_ := json.Marshal(seller)
	sellerJSON := msgdata.SellerJSON {}
	json.Unmarshal(sellerMarshaled, &sellerJSON)
	
	ech := make(msgbus.EventChan)
	ps := msgbus.New(1)

	configRepo, connectionRepo, contractRepo, destRepo, minerRepo, sellerRepo := InitializeJSONRepos()

	configRepo.AddConfigInfo(configJSON)
	connectionRepo.AddConnection(connectionJSON)
	contractRepo.AddContract(contractJSON)
	destRepo.AddDest(destJSON)
	minerRepo.AddMiner(minerJSON)
	sellerRepo.AddSeller(sellerJSON)

	var ConfigMSG msgbus.ConfigInfo
	var ConnectionMSG msgbus.Connection
	var ContractMSG msgbus.Contract
	var DestMSG msgbus.Dest
	var MinerMSG msgbus.Miner
	var SellerMSG msgbus.Seller

	configMSG := msgdata.ConvertConfigInfoJSONtoConfigInfoMSG(configRepo.ConfigInfoJSONs[0], ConfigMSG)
	connectionMSG := msgdata.ConvertConnectionJSONtoConnectionMSG(connectionRepo.ConnectionJSONs[0], ConnectionMSG)
	contractMSG := msgdata.ConvertContractJSONtoContractMSG(contractRepo.ContractJSONs[0], ContractMSG)
	destMSG := msgdata.ConvertDestJSONtoDestMSG(destRepo.DestJSONs[0], DestMSG)
	minerMSG := msgdata.ConvertMinerJSONtoMinerMSG(minerRepo.MinerJSONs[0], MinerMSG)
	sellerMSG := msgdata.ConvertSellerJSONtoSellerMSG(sellerRepo.SellerJSONs[0], SellerMSG)
	
	go func(ech msgbus.EventChan) {
		for e := range ech {
			if e.EventType == msgbus.GetEvent {
				switch e.Msg {
				case msgbus.ConfigMsg:
					if e.Data == nil {
						t.Errorf("Failed to add Config to message bus")
					} 
				case msgbus.DestMsg:
					if e.Data == nil {
						t.Errorf("Failed to add Dest to message bus")
					} 
				case msgbus.SellerMsg:
					if e.Data == nil {
						t.Errorf("Failed to add Seller to message bus")
					} 
				case msgbus.ContractMsg:
					if e.Data == nil {
						t.Errorf("Failed to add Contract to message bus")
					} 
				case msgbus.MinerMsg:
					if e.Data == nil {
						t.Errorf("Failed to add Miner to message bus")
					} 
				case msgbus.ConnectionMsg:
					if e.Data == nil {
						t.Errorf("Failed to add Connection to message bus")
					} 
				default:
				
				} 
			}
		}
	}(ech)

	ps.Pub(msgbus.ConfigMsg, msgbus.IDString(configMSG.ID), msgbus.ConfigInfo{})
	ps.Pub(msgbus.DestMsg, msgbus.IDString(destMSG.ID), msgbus.Dest{})
	ps.Pub(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), msgbus.Seller{})
	ps.Pub(msgbus.ContractMsg, msgbus.IDString(contractMSG.ID), msgbus.Contract{})
	ps.Pub(msgbus.MinerMsg, msgbus.IDString(minerMSG.ID), msgbus.Miner{})
	ps.Pub(msgbus.ConnectionMsg, msgbus.IDString(connectionMSG.ID), msgbus.Connection{})

	ps.Sub(msgbus.ConfigMsg, msgbus.IDString(configMSG.ID), ech)
	ps.Sub(msgbus.DestMsg, msgbus.IDString(destMSG.ID), ech)
	ps.Sub(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), ech)
	ps.Sub(msgbus.ContractMsg, msgbus.IDString(contractMSG.ID), ech)
	ps.Sub(msgbus.MinerMsg, msgbus.IDString(minerMSG.ID), ech)
	ps.Sub(msgbus.ConnectionMsg, msgbus.IDString(connectionMSG.ID), ech)

	ps.Set(msgbus.ConfigMsg, msgbus.IDString(configMSG.ID), configMSG)
	ps.Set(msgbus.DestMsg, msgbus.IDString(destMSG.ID), destMSG)
	ps.Set(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), sellerMSG)
	ps.Set(msgbus.ContractMsg, msgbus.IDString(contractMSG.ID), contractMSG)
	ps.Set(msgbus.MinerMsg, msgbus.IDString(minerMSG.ID), minerMSG)
	ps.Set(msgbus.ConnectionMsg, msgbus.IDString(connectionMSG.ID), connectionMSG)

	ps.Get(msgbus.ConfigMsg, msgbus.IDString(configMSG.ID), ech)
	ps.Get(msgbus.DestMsg, msgbus.IDString(destMSG.ID), ech)
	ps.Get(msgbus.SellerMsg, msgbus.IDString(sellerMSG.ID), ech)
	ps.Get(msgbus.ContractMsg, msgbus.IDString(contractMSG.ID), ech)
	ps.Get(msgbus.MinerMsg, msgbus.IDString(minerMSG.ID), ech)
	ps.Get(msgbus.ConnectionMsg, msgbus.IDString(connectionMSG.ID), ech)
}