package msgdata

import "errors"

type IDString string
type ContractID IDString

//Struct of Seller parameters in JSON 
type SellerJSON struct {
	ID                     string `json:"ID"`
	DefaultSeller          string `json:"Default Seller"`
	TotalAvailableHashRate string `json:"Total Available Hashrate"`
	UnusedHashRate         string `json:"Unused Hash Rate"`
	NewContracts           []ContractID
	ReadyContracts         []ContractID
	ActiveContracts        []ContractID
}

//Struct that stores slice of all JSON Seller structs in Repo
type SellerRepo struct {
	SellerJSONs []SellerJSON
}

//Initialize Repo with empty slice of JSON Seller structs
func NewSeller() *SellerRepo {
	return &SellerRepo{}
}

//Return all Seller Structs in Repo
func (r *SellerRepo) GetAllSellers() []SellerJSON {
	return r.SellerJSONs
}

//Return Seller Struct by ID
func (r *SellerRepo) GetSeller(id string) (SellerJSON, error) {
	for i,d := range r.SellerJSONs {
		if d.ID == id {
			return r.SellerJSONs[i], nil
		}
	}
	return r.SellerJSONs[0], errors.New("ID not found")
}

//Add a new Seller Struct to to Repo
func (r *SellerRepo) AddSeller(dest SellerJSON) {
	r.SellerJSONs = append(r.SellerJSONs, dest)
}

//Update Seller Struct with specific ID and leave empty parameters unchanged
func (r *SellerRepo) UpdateSeller(id string, newSeller SellerJSON) error {
	for i,d := range r.SellerJSONs {
		if d.ID == id {
			if newSeller.DefaultSeller != "" {r.SellerJSONs[i].DefaultSeller = newSeller.DefaultSeller}
			if newSeller.TotalAvailableHashRate != "" {r.SellerJSONs[i].TotalAvailableHashRate = newSeller.TotalAvailableHashRate}
			if newSeller.UnusedHashRate != "" {r.SellerJSONs[i].UnusedHashRate = newSeller.UnusedHashRate}

			return nil
		}
	}
	return errors.New("ID not found")
}

//Delete Seller Struct with specific ID
func (r *SellerRepo) DeleteSeller(id string) error {
	for i,d := range r.SellerJSONs {
		if d.ID == id {
			r.SellerJSONs = append(r.SellerJSONs[:i], r.SellerJSONs[i+1:]...)

			return nil
		}
	}
	return errors.New("ID not found")
}