package msgbus

import (
	"fmt"
	"testing"
)

func TestBoilerPlateFunc(t *testing.T) {
	ech := make(EventChan)

	config := ConfigInfo{
		ID:          	"ConfigID01",
		DefaultDest: 	"DestID01",
		NodeOperator:	"NOID01",
	}
	dest := Dest{
		ID:       DestID(DEFAULT_DEST_ID),
		NetUrl:   DestNetUrl("stratum+tcp://127.0.0.1:3334/"),
	}
	nodeOp := NodeOperator{
		ID:                     "SellerID01",
		DefaultDest:            "DestID01",
		TotalAvailableHashRate: 0,
		UnusedHashRate:         0,
		Contracts:        make(map[ContractID]ContractState),
	}
	contract := Contract{}
	miner := Miner{}
	connection := Connection{}

	ps := New(1)
	//if msg != "Accounting Manager Package" && err != nil {
	//	t.Fatalf("Test Failed")
	//}

	go func(ech EventChan) {
		for e := range ech {
			fmt.Printf("Read Chan: %+v\n", e)
		}

		fmt.Printf("Closed Read Chan\n")

	}(ech)

	ps.Pub(ConfigMsg, "configMsg01", ConfigInfo{})
	ps.Pub(DestMsg, "destMsg01", Dest{})
	ps.Pub(NodeOperatorMsg, "sellerMsg01", NodeOperator{})
	ps.Pub(ContractMsg, "contractMsg01", Contract{})
	ps.Pub(MinerMsg, "minerMsg01", Miner{})
	ps.Pub(ConnectionMsg, "connectionMsg01", Connection{})

	ps.Sub(ConfigMsg, "configMsg01", ech)
	ps.Sub(DestMsg, "destMsg01", ech)
	ps.Sub(NodeOperatorMsg, "sellerMsg01", ech)
	ps.Sub(ContractMsg, "contractMsg01", ech)
	ps.Sub(MinerMsg, "minerMsg01", ech)
	ps.Sub(ConnectionMsg, "connectionMsg01", ech)

	ps.Set(ConfigMsg, "configMsg01", config)
	ps.Set(DestMsg, "destMsg01", dest)
	ps.Set(NodeOperatorMsg, "sellerMsg01", nodeOp)
	ps.Set(ContractMsg, "contractMsg01", contract)
	ps.Set(MinerMsg, "minerMsg01", miner)
	ps.Set(ConnectionMsg, "connectionMsg01", connection)

	ps.Get(ConfigMsg, "", ech)
	ps.Get(DestMsg, "", ech)
	ps.Get(NodeOperatorMsg, "", ech)
	ps.Get(ContractMsg, "", ech)
	ps.Get(MinerMsg, "", ech)
	ps.Get(ConnectionMsg, "", ech)

	ps.Get(ConfigMsg, "configMsg01", ech)
	ps.Get(DestMsg, "destMsg01", ech)
	ps.Get(NodeOperatorMsg, "sellerMsg01", ech)
	ps.Get(ContractMsg, "contractMsg01", ech)
	ps.Get(MinerMsg, "minerMsg01", ech)
	ps.Get(ConnectionMsg, "connectionMsg01", ech)

}
