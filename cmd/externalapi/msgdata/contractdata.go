package msgdata

import (
	"errors"
	"strconv"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

// Struct of Contract parameters in JSON 
type ContractJSON struct {
	ID              	string 	`json:"ID"`
	State           	string 	`json:"State"`
	Buyer				string 	`json:"Buyer"`
	Dest				string	`json:"Dest"`
	CommitedHashRate	string 	`json:"Commited Hash Rate"`
	TargetHashRate   	string 	`json:"Target Hash Rate"`
	CurrentHashRate  	string 	`json:"Current Hash Rate"`
	Tolerance        	string 	`json:"Tolerance"`
	Penalty          	string 	`json:"Penalty"`
	Priority         	string 	`json:"Priority"`
	StartDate        	string 	`json:"Start Date"`
	EndDate          	string 	`json:"End Date"`
}

//Struct that stores slice of all JSON Contract structs in Repo
type ContractRepo struct {
	ContractJSONs []ContractJSON
}

//Initialize Repo with empty slice of JSON Contract structs
func NewContract() *ContractRepo {
	return &ContractRepo{}
}

//Return all Contract Structs in Repo
func (r *ContractRepo) GetAllContracts() []ContractJSON {
	return r.ContractJSONs
}

//Return Contract Struct by ID
func (r *ContractRepo) GetContract(id string) (ContractJSON, error) {
	for i,c := range r.ContractJSONs {
		if c.ID == id {
			return r.ContractJSONs[i], nil
		}
	}
	return r.ContractJSONs[0], errors.New("ID not found")
}

//Add a new Contract Struct to to Repo
func (r *ContractRepo) AddContract(contract ContractJSON) {
	r.ContractJSONs = append(r.ContractJSONs, contract)
}

//Converts Contract struct from msgbus to JSON struct and adds it to Repo
func (r *ContractRepo) AddContractFromMsgBus(contract msgbus.Contract) {
	var contractJSON ContractJSON

	contractJSON.ID = string(contract.ID)
	contractJSON.State = string(contract.State)
	contractJSON.Buyer = string(contract.Buyer)
	contractJSON.Dest = string(contract.Dest)
	contractJSON.CommitedHashRate = strconv.Itoa(contract.CommitedHashRate)
	contractJSON.TargetHashRate = strconv.Itoa(contract.TargetHashRate)
	contractJSON.CurrentHashRate = strconv.Itoa(contract.CurrentHashRate)
	contractJSON.Tolerance = strconv.Itoa(contract.Tolerance)
	contractJSON.Penalty = strconv.Itoa(contract.Penalty)
	contractJSON.Priority = strconv.Itoa(contract.Priority)
	contractJSON.StartDate = contract.StartDate.String()
	contractJSON.EndDate = contract.EndDate.String()
	
	r.ContractJSONs = append(r.ContractJSONs, contractJSON)
}

//Update Contract Struct with specific ID and leave empty parameters unchanged
func (r *ContractRepo) UpdateContract(id string, newContract ContractJSON) error {
	for i,c := range r.ContractJSONs {
		if c.ID == id {
			if newContract.State != "" {r.ContractJSONs[i].State = newContract.State}
			if newContract.Buyer != "" {r.ContractJSONs[i].Buyer = newContract.Buyer}
			if newContract.Dest != "" {r.ContractJSONs[i].Dest = newContract.Dest}
			if newContract.CommitedHashRate != "" {r.ContractJSONs[i].CommitedHashRate = newContract.CommitedHashRate}
			if newContract.CurrentHashRate != "" {r.ContractJSONs[i].CurrentHashRate = newContract.CurrentHashRate}
			if newContract.Tolerance != "" {r.ContractJSONs[i].Tolerance = newContract.Tolerance}
			if newContract.Priority != "" {r.ContractJSONs[i].Priority = newContract.Priority}
			if newContract.StartDate != "" {r.ContractJSONs[i].StartDate = newContract.StartDate}
			if newContract.EndDate != "" {r.ContractJSONs[i].EndDate = newContract.EndDate}

			return nil
		}
	}
	return errors.New("ID not found")
}

//Delete Contract Struct with specific ID
func (r *ContractRepo) DeleteContract(id string) error {
	for i,c := range r.ContractJSONs {
		if c.ID == id {
			r.ContractJSONs = append(r.ContractJSONs[:i], r.ContractJSONs[i+1:]...)

			return nil
		}
	}
	return errors.New("ID not found")
}

func ConvertContractJSONtoContractMSG(contract ContractJSON, msg msgbus.Contract) msgbus.Contract {
	msg.ID = msgbus.ContractID(contract.ID)
	msg.State = msgbus.ContractState(contract.State)
	msg.Buyer = msgbus.BuyerID(contract.Buyer)
	msg.Dest = msgbus.DestID(contract.Dest)
	msg.CommitedHashRate,_ = strconv.Atoi(contract.CommitedHashRate)
	msg.TargetHashRate,_ = strconv.Atoi(contract.TargetHashRate)
	msg.CurrentHashRate,_ = strconv.Atoi(contract.CurrentHashRate)
	msg.Tolerance,_ = strconv.Atoi(contract.Tolerance)
	msg.Penalty,_ = strconv.Atoi(contract.Penalty)
	msg.Priority,_ = strconv.Atoi(contract.Priority)
	msg.StartDate,_ = time.Parse(contract.StartDate, "000000")
	msg.EndDate,_ = time.Parse(contract.EndDate, "000000")

	return msg	
}