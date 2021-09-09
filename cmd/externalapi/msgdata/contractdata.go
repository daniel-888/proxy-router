package msgdata

import (
	"errors"
	//"strconv"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

// Struct of Contract parameters in JSON 
type ContractJSON struct {
	ID              	string 	`json:"id"`
	State           	string 	`json:"state"`
	Buyer				string 	`json:"buyer"`
	Dest				string	`json:"dest"`
	CommitedHashRate	int 	`json:"commitedHashRate"`
	TargetHashRate   	int 	`json:"targetHashRate"`
	CurrentHashRate  	int 	`json:"currentHashRate"`
	Tolerance        	int 	`json:"tolerance"`
	Penalty          	int 	`json:"penalty"`
	Priority         	int 	`json:"priority"`
	StartDate        	string 	`json:"startDate"`
	EndDate          	string 	`json:"endDate"`
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
	contractJSON.CommitedHashRate = contract.CommitedHashRate //strconv.Itoa(contract.CommitedHashRate)
	contractJSON.TargetHashRate = contract.TargetHashRate
	contractJSON.CurrentHashRate = contract.CurrentHashRate
	contractJSON.Tolerance = contract.Tolerance
	contractJSON.Penalty = contract.Penalty
	contractJSON.Priority = contract.Priority
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
			if newContract.CommitedHashRate != 0 {r.ContractJSONs[i].CommitedHashRate = newContract.CommitedHashRate}
			if newContract.CurrentHashRate != 0 {r.ContractJSONs[i].CurrentHashRate = newContract.CurrentHashRate}
			if newContract.Tolerance != 0 {r.ContractJSONs[i].Tolerance = newContract.Tolerance}
			if newContract.Priority != 0 {r.ContractJSONs[i].Priority = newContract.Priority}
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
	msg.CommitedHashRate = contract.CommitedHashRate //strconv.Atoi(contract.CommitedHashRate)
	msg.TargetHashRate = contract.TargetHashRate
	msg.CurrentHashRate = contract.CurrentHashRate
	msg.Tolerance = contract.Tolerance
	msg.Penalty = contract.Penalty
	msg.Priority = contract.Priority
	msg.StartDate,_ = time.Parse(contract.StartDate, "000000")
	msg.EndDate,_ = time.Parse(contract.EndDate, "000000")

	return msg	
}