package configurationmanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadConfiguration(file string, pkg string) (map[string]interface{}, error) {
	var data map[string]interface{}
	configfile, err := os.Open(file)
	if err != nil {
		return data, err
	}
	defer configfile.Close()
	byteValue,_ := ioutil.ReadAll(configfile)

	err = json.Unmarshal([]byte(byteValue), &data)
	return data[pkg].(map[string]interface{}),err
}