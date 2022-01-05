package msgdata

import (
	"errors"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

// Struct of Contract parameters in JSON 
type ContractJSON struct {
	ID               		string 	`json:"id"`
	State            		string 	`json:"state"`
	Buyer			 		string 	`json:"buyer"`
	Price			 		int		`json:"price"`
	Limit			 		int		`json:"limit"`
	Speed			 		int		`json:"speed"`
	Length        	 		int		`json:"length"`
	Port          	 		int		`json:"port"`
	StartingBlockTimestamp	int		`json:"startingBlockTimestamp"`
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
	contractJSON.Price = contract.Price 
	contractJSON.Limit = contract.Limit
	contractJSON.Speed = contract.Speed
	contractJSON.Length = contract.Length
	contractJSON.StartingBlockTimestamp = contract.StartingBlockTimestamp
	
	r.ContractJSONs = append(r.ContractJSONs, contractJSON)
}

//Update Contract Struct with specific ID and leave empty parameters unchanged
func (r *ContractRepo) UpdateContract(id string, newContract ContractJSON) error {
	for i,c := range r.ContractJSONs {
		if c.ID == id {
			if newContract.State != "" {r.ContractJSONs[i].State = newContract.State}
			if newContract.Buyer != "" {r.ContractJSONs[i].Buyer = newContract.Buyer}
			if newContract.Price != 0 {r.ContractJSONs[i].Price = newContract.Price}
			if newContract.Limit != 0 {r.ContractJSONs[i].Limit = newContract.Limit}
			if newContract.Speed != 0 {r.ContractJSONs[i].Speed = newContract.Speed}
			if newContract.Length != 0 {r.ContractJSONs[i].Length = newContract.Length}
			if newContract.Port != 0 {r.ContractJSONs[i].Port = newContract.Port}
			if newContract.StartingBlockTimestamp != 0 {r.ContractJSONs[i].StartingBlockTimestamp = newContract.StartingBlockTimestamp}

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
	msg.Price = contract.Price 
	msg.Limit = contract.Limit
	msg.Speed = contract.Speed
	msg.Length = contract.Length
	msg.StartingBlockTimestamp = contract.StartingBlockTimestamp

	return msg	
}

func ConvertContractMSGtoContractJSON(msg msgbus.Contract) (contract ContractJSON) {
	contract.ID = string(msg.ID)
	contract.State = string(msg.State)
	contract.Buyer = string(msg.Buyer)
	contract.Price = msg.Price
	contract.Limit = msg.Limit
	contract.Speed = msg.Speed
	contract.Length = msg.Length
	contract.StartingBlockTimestamp = msg.StartingBlockTimestamp

	return contract	
}