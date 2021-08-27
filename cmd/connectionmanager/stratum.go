package connectionmanager

import (
	"encoding/json"
	"fmt"
)

type stratumMethods string
type stratumState int

const (
	StratumNew stratumState = iota
	StratumSubscribed
	StratumAuthorized
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

func getRequestMsg(msg []byte) (*request, error) {

	fmt.Printf("Request: %s", msg)

	r := request{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		fmt.Printf("Unmarshal request error:%s\n", err)
		fmt.Printf("Unmarshal request MSG:%s\n", msg)
	}

	return &r, err
}

func getResponceMsg(msg []byte) (*responce, error) {

	fmt.Printf("Responce: %s", msg)

	r := responce{}
	err := json.Unmarshal(msg, &r)
	if err != nil {
		fmt.Printf("Unmarshal responce error:%s\n", err)
		fmt.Printf("Unmarshal responce MSG:%s\n", msg)
	}
	return &r, err
}
