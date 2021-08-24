package msgdata

import "errors"

// Struct of ConfigInfo parameters in JSON 
type ConfigInfoJSON struct {
	ID          		string `json:"ID"`
	DefaultConfigInfo 	string `json:"Default ConfigInfo"`
	Seller     		 	string `json:"Seller"`
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
func (r *ConfigInfoRepo) AddConfigInfo(dest ConfigInfoJSON) {
	r.ConfigInfoJSONs = append(r.ConfigInfoJSONs, dest)
}

//Update ConfigInfo Struct with specific ID and leave empty parameters unchanged
func (r *ConfigInfoRepo) UpdateConfigInfo(id string, newConfigInfo ConfigInfoJSON) error {
	for i,c := range r.ConfigInfoJSONs {
		if c.ID == id {
			if newConfigInfo.DefaultConfigInfo != "" {r.ConfigInfoJSONs[i].DefaultConfigInfo = newConfigInfo.DefaultConfigInfo}
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