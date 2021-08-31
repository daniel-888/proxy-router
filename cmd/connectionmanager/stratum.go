package connectionmanager

import (
	"encoding/json"
	"fmt"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

type stratumMethods string
type stratumStates string

const (
	StratumNew        stratumStates = "StratumNew"
	StratumSubscribed stratumStates = "StratumSubscribed"
	StratumAuthorized stratumStates = "StratumAuthorized"
)

const (
	MINING_CONFIGURE  stratumMethods = "mining.configure"
	MINING_SUBSCRIBE  stratumMethods = "mining.subscribe"
	MINING_AUTHORIZE  stratumMethods = "mining.authorize"
	MINING_SET_TARGET stratumMethods = "mining.set_target"
	MINING_NOTIFY     stratumMethods = "mining.notify"
	MINING_SUBMIT     stratumMethods = "mining.submit"
)

type request struct {
	ID     int      `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"parms"`
}

type responce struct {
	ID     int    `json:"id"`
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

//------------------------------------------------------
//
//------------------------------------------------------
func getRequestMsg(msg []byte) (*request, error) {

	fmt.Printf("Stratum Request: %s\n", msg)

	r := request{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error Unmarshaling Request Err:%s\n", err)
		return nil, err
	}

	return &r, err
}

//------------------------------------------------------
//
//------------------------------------------------------
func getResponceMsg(msg []byte) (*responce, error) {

	fmt.Printf("Stratum Responce: %s\n", msg)

	r := responce{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error Unmarshaling Responce Err:%s\n", err)
		return nil, err
	}

	return &r, err
}
