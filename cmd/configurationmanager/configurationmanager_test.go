package configurationmanager

import (
	"testing"

	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)
func TestLoadConfig (t *testing.T) {
	configMap,err := LoadConfiguration("exampleconfig.json")
	if err != nil {
		t.Errorf("LoadConfiguration returned error")
	}
	if configMap["id"] != "ConfigID01" && configMap["destID"] != "DestID01" && configMap["seller"] != "SellerID01" {
		t.Errorf("Failed to correctly load exampleconfig.json")
	}
}

func TestLoadConfigToAPIandMsgBus (t *testing.T) {
	configMap,err := LoadConfiguration("exampleconfig.json")
	if err != nil {
		t.Errorf("LoadConfiguration returned error")
	}
	ech := make(msgbus.EventChan)
	ps := msgbus.New(1)

	var config msgdata.ConfigInfoJSON
	config.ID = configMap["id"].(string)
	config.ID = configMap["destID"].(string)
	config.ID = configMap["seller"].(string)

	var seller msgdata.SellerJSON
	newContractsMap := make(map[msgbus.ContractID]bool)
	for key,value := range configMap["newContracts"].(map[string]interface{}){
		newContractsMap[msgbus.ContractID(key)] = value.(bool)
	}
	readyContractsMap := make(map[msgbus.ContractID]bool)
	for key,value := range configMap["readyContracts"].(map[string]interface{}){
		readyContractsMap[msgbus.ContractID(key)] = value.(bool)
	}
	activeContractsMap := make(map[msgbus.ContractID]bool)
	for key,value := range configMap["activeContracts"].(map[string]interface{}){
		activeContractsMap[msgbus.ContractID(key)] = value.(bool)
	}

	seller.ID = configMap["seller"].(string)
	seller.NewContracts = newContractsMap
	seller.ReadyContracts = readyContractsMap
	seller.ActiveContracts = activeContractsMap

	configRepo,_,_,_,_,sellerRepo := externalapi.InitializeJSONRepos()
	configRepo.AddConfigInfo(config)
	var ConfigMSG msgbus.ConfigInfo
	configMSG := msgdata.ConvertConfigInfoJSONtoConfigInfoMSG(configRepo.ConfigInfoJSONs[0], ConfigMSG)
	sellerRepo.AddSeller(seller)
	var SellerMSG msgbus.Seller
	sellerMSG := msgdata.ConvertSellerJSONtoSellerMSG(sellerRepo.SellerJSONs[0], SellerMSG)

	go func(ech msgbus.EventChan) {
		for e := range ech {
			if e.EventType == msgbus.GetEvent {
				switch e.Msg {
				case msgbus.ConfigMsg:
					if e.Data == nil {
						t.Errorf("Failed to add Config to message bus")
					} 
				case msgbus.SellerMsg:
					if e.Data == nil {
						t.Errorf("Failed to add Seller to message bus")
					} 
				default:
	
				}
			}
		}
	}(ech)

	ps.Pub(msgbus.ConfigMsg, msgbus.IDString(config.ID), msgbus.ConfigInfo{})
	ps.Pub(msgbus.SellerMsg, msgbus.IDString(seller.ID), msgbus.Seller{})

	ps.Sub(msgbus.ConfigMsg, msgbus.IDString(config.ID), ech)
	ps.Sub(msgbus.SellerMsg, msgbus.IDString(seller.ID), ech)

	ps.Set(msgbus.ConfigMsg, msgbus.IDString(config.ID), configMSG)
	ps.Set(msgbus.SellerMsg, msgbus.IDString(seller.ID), sellerMSG)
	
	ps.Get(msgbus.ConfigMsg, msgbus.IDString(config.ID), ech)
	ps.Get(msgbus.SellerMsg, msgbus.IDString(seller.ID), ech)
}
