package configurationmanager

import (
	"testing"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
)

func TestLoadConfig (t *testing.T) {
	config,_ := LoadConfiguration("exampleconfig.json")
	if config.ID != "ConfigID01" && config.DefaultDest != "DestID01" && config.Seller != "SellerID01" {
		t.Errorf("Failed to correctly load exampleconfig.json")
	}
}

func TestLoadConfigToAPIandMsgBus (t *testing.T) {
	config,_ := LoadConfiguration("exampleconfig.json")
	ech := make(msgbus.EventChan)
	ps := msgbus.New(1)
	configRepo,_,_,_,_,_ := externalapi.InitializeJSONRepos()
	configRepo.AddConfigInfo(config)
	var ConfigMSG msgbus.ConfigInfo
	configMSG := msgdata.ConvertConfigInfoJSONtoConfigInfoMSG(configRepo.ConfigInfoJSONs[0], ConfigMSG)

	go func(ech msgbus.EventChan) {
		for e := range ech {
			if e.EventType == msgbus.GetEvent {
				if e.Data == nil {
					t.Errorf("Failed to add Config to message bus")
				} 
			}
		}
	}(ech)

	ps.Pub(msgbus.ConfigMsg, msgbus.IDString(config.ID), msgbus.ConfigInfo{})
	ps.Sub(msgbus.ConfigMsg, msgbus.IDString(config.ID), ech)
	ps.Set(msgbus.ConfigMsg, msgbus.IDString(config.ID), configMSG)
	ps.Get(msgbus.ConfigMsg, msgbus.IDString(config.ID), ech)
}
