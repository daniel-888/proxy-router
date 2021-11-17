package configurationmanager

import (
	"testing"

	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func TestLoadConfig(t *testing.T) {
	cmConfig,err := LoadConfiguration("testconfig.json", "configurationManager")
	if err != nil {
		t.Errorf("LoadConfiguration returned error")
	}

	if cmConfig["id"] != "ConfigID01" && cmConfig["destID"] != "DestID01" && cmConfig["seller"] != "SellerID01" {
		t.Errorf("Failed to correctly load configuration file")
	}
}

func TestLoadConfigFromServer(t *testing.T) {
	DownloadConfig("https://lumerin-node-configs.s3.amazonaws.com/config.json")
}

func TestLoadConfigToAPIandMsgBus(t *testing.T) {
	cmConfig,err := LoadConfiguration("testconfig.json", "configurationManager")
	if err != nil {
		t.Errorf("LoadConfiguration returned error")
	}

	ech := make(msgbus.EventChan)
	ps := msgbus.New(1)

	var config msgdata.ConfigInfoJSON
	config.ID = cmConfig["id"].(string)
	config.ID = cmConfig["destID"].(string)
	config.ID = cmConfig["seller"].(string)

	var seller msgdata.SellerJSON
	availableContractsMap := make(map[msgbus.ContractID]bool)
	for key,value := range cmConfig["availableContracts"].(map[string]interface{}){
		availableContractsMap[msgbus.ContractID(key)] = value.(bool)
	}
	activeContractsMap := make(map[msgbus.ContractID]bool)
	for key,value := range cmConfig["activeContracts"].(map[string]interface{}){
		activeContractsMap[msgbus.ContractID(key)] = value.(bool)
	}

	seller.ID = cmConfig["seller"].(string)
	seller.AvailableContracts = availableContractsMap
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

func TestLoadMsgBusFromConfig(t *testing.T) {
	ps,ech := LoadMsgBusFromConfig()

	ps.Get(msgbus.DestMsg, msgbus.IDString("DestID01"), ech)
	ps.Get(msgbus.SellerMsg, msgbus.IDString("SellerID01"), ech)
	ps.Get(msgbus.ConfigMsg, msgbus.IDString("ConfigID01"), ech)
}