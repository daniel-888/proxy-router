package configurationmanager

import (
	"encoding/json"
	"os"

	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
)

func LoadConfiguration(file string) (msgdata.ConfigInfoJSON, error) {
	var config msgdata.ConfigInfoJSON
	configfile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configfile)
	err = jsonParser.Decode(&config)
	return config,err
}