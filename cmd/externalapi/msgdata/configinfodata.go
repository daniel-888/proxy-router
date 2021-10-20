package msgdata

import (
	"errors"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

// Struct of ConfigInfo parameters in JSON
type ConfigInfoJSON struct {
	ID          		string `json:"id"`
	DefaultDest 		string `json:"defaultDest"`
	Seller     		 	string `json:"seller"`
}

//Struct that stores slice of all JSON ConfigInfo structs in Repo
type ConfigInfoRepo struct {
	ConfigInfoJSONs []ConfigInfoJSON
}

//Initialize Repo with empty slice of JSON ConfigInfo structs
func NewConfigInfo() *ConfigInfoRepo {
	return &ConfigInfoRepo{}
}

//Return all ConfigInfo Structs in Repo
func (r *ConfigInfoRepo) GetAllConfigInfos() []ConfigInfoJSON {
	return r.ConfigInfoJSONs
}

//Return ConfigInfo Struct by ID
func (r *ConfigInfoRepo) GetConfigInfo(id string) (ConfigInfoJSON, error) {
	for i,c := range r.ConfigInfoJSONs {
		if c.ID == id {
			return r.ConfigInfoJSONs[i], nil
		}
	}
	return r.ConfigInfoJSONs[0], errors.New("ID not found")
}

//Add a new ConfigInfo Struct to to Repo
func (r *ConfigInfoRepo) AddConfigInfo(conf ConfigInfoJSON) {
	r.ConfigInfoJSONs = append(r.ConfigInfoJSONs, conf)
}

//Converts ConfigInfo struct from msgbus to JSON struct and adds it to Repo
func (r *ConfigInfoRepo) AddConfigInfoFromMsgBus(conf msgbus.ConfigInfo) {
	var confJSON ConfigInfoJSON
	
	confJSON.ID = string(conf.ID)
	confJSON.DefaultDest = string(conf.DefaultDest)
	confJSON.Seller = string(conf.Seller)
	
	r.ConfigInfoJSONs = append(r.ConfigInfoJSONs, confJSON)
}

//Update ConfigInfo Struct with specific ID and leave empty parameters unchanged
func (r *ConfigInfoRepo) UpdateConfigInfo(id string, newConfigInfo ConfigInfoJSON) error {
	for i,c := range r.ConfigInfoJSONs {
		if c.ID == id {
			if newConfigInfo.DefaultDest != "" {r.ConfigInfoJSONs[i].DefaultDest = newConfigInfo.DefaultDest}
			if newConfigInfo.Seller != "" {r.ConfigInfoJSONs[i].Seller = newConfigInfo.Seller}

			return nil
		}
	}
	return errors.New("ID not found")
}

//Delete ConfigInfo Struct with specific ID
func (r *ConfigInfoRepo) DeleteConfigInfo(id string) error {
	for i,c := range r.ConfigInfoJSONs {
		if c.ID == id {
			r.ConfigInfoJSONs = append(r.ConfigInfoJSONs[:i], r.ConfigInfoJSONs[i+1:]...)

			return nil
		}
	}
	return errors.New("ID not found")
}

func ConvertConfigInfoJSONtoConfigInfoMSG(conf ConfigInfoJSON, msg msgbus.ConfigInfo) msgbus.ConfigInfo {
	msg.ID = msgbus.ConfigID(conf.ID)
	msg.DefaultDest = msgbus.DestID(conf.DefaultDest)
	msg.Seller = msgbus.SellerID(conf.Seller)

	return msg	
}