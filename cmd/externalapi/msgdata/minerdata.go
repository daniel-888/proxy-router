package msgdata

import (
	"errors"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

//Struct of Miner parameters in JSON 
type MinerJSON struct {
	ID                      string	`json:"id"`
	State                   string 	`json:"state"`
	Seller                  string 	`json:"seller"`
	Dest                   	string 	`json:"dest"`
	InitialMeasuredHashRate int 	`json:"initialMeasuredHashRate"`
	CurrentHashRate         int 	`json:"currentHashRate"`
}

//Struct that stores slice of all JSON Miner structs in Repo
type MinerRepo struct {
	MinerJSONs []MinerJSON
}

//Initialize Repo with empty slice of JSON Miner structs
func NewMiner() *MinerRepo {
	return &MinerRepo{}
}

//Return all Miner Structs in Repo
func (r *MinerRepo) GetAllMiners() []MinerJSON {
	return r.MinerJSONs
}

//Return Miner Struct by ID
func (r *MinerRepo) GetMiner(id string) (MinerJSON, error) {
	for i,m := range r.MinerJSONs {
		if m.ID == id {
			return r.MinerJSONs[i], nil
		}
	}
	return r.MinerJSONs[0], errors.New("ID not found")
}

//Add a new Miner Struct to to Repo
func (r *MinerRepo) AddMiner(miner MinerJSON) {
	r.MinerJSONs = append(r.MinerJSONs, miner)
}

//Converts Miner struct from msgbus to JSON struct and adds it to Repo
func (r *MinerRepo) AddMinerFromMsgBus(miner msgbus.Miner) {
	var minerJSON MinerJSON

	minerJSON.ID = string(miner.ID)
	minerJSON.State = string(miner.State)
	minerJSON.Seller = string(miner.Seller)
	minerJSON.Dest = string(miner.Dest)
	minerJSON.InitialMeasuredHashRate = miner.InitialMeasuredHashRate
	minerJSON.CurrentHashRate = miner.CurrentHashRate
	
	r.MinerJSONs = append(r.MinerJSONs, minerJSON)
}

//Update Miner Struct with specific ID and leave empty parameters unchanged
func (r *MinerRepo) UpdateMiner(id string, newMiner MinerJSON) error {
	for i,m := range r.MinerJSONs {
		if m.ID == id {
			if newMiner.State != "" {r.MinerJSONs[i].State = newMiner.State}
			if newMiner.Seller != "" {r.MinerJSONs[i].Seller = newMiner.Seller}
			if newMiner.Dest != "" {r.MinerJSONs[i].Dest = newMiner.Dest}
			if newMiner.InitialMeasuredHashRate != 0 {r.MinerJSONs[i].InitialMeasuredHashRate = newMiner.InitialMeasuredHashRate}
			if newMiner.CurrentHashRate != 0 {r.MinerJSONs[i].CurrentHashRate = newMiner.CurrentHashRate}

			return nil
		}
	}
	return errors.New("ID not found")
}

//Delete Miner Struct with specific ID
func (r *MinerRepo) DeleteMiner(id string) error {
	for i,m := range r.MinerJSONs {
		if m.ID == id {
			r.MinerJSONs = append(r.MinerJSONs[:i], r.MinerJSONs[i+1:]...)

			return nil
		}
	}
	return errors.New("ID not found")
}

func ConvertMinerJSONtoMinerMSG(miner MinerJSON, msg msgbus.Miner) msgbus.Miner {
	msg.ID = msgbus.MinerID(miner.ID)
	msg.State = msgbus.MinerState(miner.State)
	msg.Seller = msgbus.SellerID(miner.Seller)
	msg.Dest = msgbus.DestID(miner.Dest)
	msg.InitialMeasuredHashRate = miner.InitialMeasuredHashRate
	msg.CurrentHashRate = miner.CurrentHashRate

	return msg	
}