package msgbus

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/log"
)

func TestBoilerPlateFunc(t *testing.T) {
	eventChan := make(EventChan)

	config := ConfigInfo{
		ID:           "ConfigID01",
		DefaultDest:  "DestID01",
		NodeOperator: "NOID01",
	}
	dest := Dest{
		ID:     DestID(DEFAULT_DEST_ID),
		NetUrl: DestNetUrl("stratum+tcp://127.0.0.1:3334/"),
	}
	nodeOp := NodeOperator{
		ID:                     "SellerID01",
		DefaultDest:            "DestID01",
		TotalAvailableHashRate: 0,
		UnusedHashRate:         0,
		Contracts:              make(map[ContractID]ContractState),
	}
	contract := Contract{}
	miner := Miner{}
	connection := Connection{}
	l := log.New()

	ps := New(1, l)
	if _, err := ps.Shutdown(); err != nil {
		time.Sleep(time.Second * 5)
		t.Error(err)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("NO ERROR!!! YAY!")
	return
	go func(eventChan EventChan) {
		for event := range eventChan {
			fmt.Printf("Read Chan: %+v\n", event)
		}

		fmt.Printf("Closed Read Chan\n")

	}(eventChan)
	defer close(eventChan)

	pubSetParams := []struct {
		msg  MsgType
		id   IDString
		data interface{}
	}{
		{ConfigMsg, "configMsg01", config},
		{DestMsg, "destMsg01", dest},
		{NodeOperatorMsg, "sellerMsg01", nodeOp},
		{ContractMsg, "contractMsg01", contract},
		{MinerMsg, "minerMsg01", miner},
		{ConnectionMsg, "connectionMsg01", connection},
	}

	for _, params := range pubSetParams {
		if _, err := ps.Pub(params.msg, params.id, params.data); err != nil {
			t.Errorf("trying to pub: %v", err)
		}
	}

	subParams := []struct {
		msg MsgType
		id  IDString
		ch  EventChan
	}{
		{ConfigMsg, "configMsg01", eventChan},
		{DestMsg, "destMsg01", eventChan},
		{NodeOperatorMsg, "sellerMsg01", eventChan},
		{ContractMsg, "contractMsg01", eventChan},
		{MinerMsg, "minerMsg01", eventChan},
		{ConnectionMsg, "connectionMsg01", eventChan},
	}

	for _, params := range subParams {
		if _, err := ps.Sub(params.msg, params.id, params.ch); err != nil {
			t.Errorf("trying to sub: %v", err)
		}
	}

	for _, params := range pubSetParams {
		if _, err := ps.Set(params.msg, params.id, params.data); err != nil {
			t.Errorf("trying to set: %v", err)
		}
	}

	getParams := []struct {
		msg MsgType
		id  IDString
		ch  EventChan
	}{
		{ConfigMsg, "", eventChan},
		{DestMsg, "", eventChan},
		{NodeOperatorMsg, "", eventChan},
		{ContractMsg, "", eventChan},
		{MinerMsg, "", eventChan},
		{ConnectionMsg, "", eventChan},
		{ConfigMsg, "configMsg01", eventChan},
		{DestMsg, "destMsg01", eventChan},
		{NodeOperatorMsg, "sellerMsg01", eventChan},
		{ContractMsg, "contractMsg01", eventChan},
		{MinerMsg, "minerMsg01", eventChan},
		{ConnectionMsg, "connectionMsg01", eventChan},
	}

	for _, params := range getParams {
		if _, err := ps.Get(params.msg, params.id, params.ch); err != nil {
			t.Errorf("trying to get: %v", err)
		}
	}
}

func TestGetRandomIDString(t *testing.T) {
	requiredRegex := `^[0-9a-fA-F]{8}\-[0-9a-fA-F]{8}\-[0-9a-fA-F]{8}\-[0-9a-fA-F]{8}$`
	regex, err := regexp.Compile(requiredRegex)
	if err != nil {
		t.Errorf("compiling regex: %v", err)
	}

	// run 100 tests
	for i := 0; i < 100; i++ {
		testID := GetRandomIDString()

		if matched := regex.Match([]byte(testID)); !matched {
			t.Errorf("GetRandomIDString returned an incorrectly formatted string: %v", testID)
		}
	}
}
