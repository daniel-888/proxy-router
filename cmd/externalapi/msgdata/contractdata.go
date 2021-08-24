package msgdata

import "errors"

// Struct of Contract parameters in JSON 
type ContractJSON struct {
	ID               string `json:"ID"`
	State            string `json:"State"`
	Buyer            string `json:"Buyer"`
	Contract         string `json:"Contract"`
	CommitedHashRate string `json:"Commited Hash Rate"`
	TargetHashRate   string `json:"Target Hash Rate"`
	CurrentHashRate  string `json:"Current Hash Rate"`
	Tolerance        string `json:"Tolerance"`
	Penalty          string `json:"Penalty"`
	Priority         string `json:"Priority"`
	StartDate        string `json:"Start Date"`
	EndDate          string `json:"End Date"`
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
func (r *ContractRepo) AddContract(dest ContractJSON) {
	r.ContractJSONs = append(r.ContractJSONs, dest)
}

//Update Contract Struct with specific ID and leave empty parameters unchanged
func (r *ContractRepo) UpdateContract(id string, newContract ContractJSON) error {
	for i,c := range r.ContractJSONs {
		if c.ID == id {
			if newContract.State != "" {r.ContractJSONs[i].State = newContract.State}
			if newContract.Buyer != "" {r.ContractJSONs[i].Buyer = newContract.Buyer}
			if newContract.Contract != "" {r.ContractJSONs[i].Contract = newContract.Contract}
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