package msgdata

import (
	"errors"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

//Struct of Seller parameters in JSON 
type SellerJSON struct {
	ID                     	string 										`json:"id"`
	DefaultDest          	string 										`json:"destID"`
	TotalAvailableHashRate 	int 										`json:"totalAvailableHashrate"`
	UnusedHashRate         	int 										`json:"unusedHashRate"`
	Contracts				map[msgbus.ContractID]msgbus.ContractState	`json:"contracts"`
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
func (r *SellerRepo) AddSeller(seller SellerJSON) {
	r.SellerJSONs = append(r.SellerJSONs, seller)
}

//Converts Seller struct from msgbus to JSON struct and adds it to Repo
func (r *SellerRepo) AddSellerFromMsgBus(seller msgbus.Seller) {
	var sellerJSON SellerJSON

	sellerJSON.ID = string(seller.ID)
	sellerJSON.DefaultDest = string(seller.DefaultDest)
	sellerJSON.TotalAvailableHashRate = seller.TotalAvailableHashRate
	sellerJSON.UnusedHashRate = seller.UnusedHashRate
	sellerJSON.Contracts = seller.Contracts
	
	r.SellerJSONs = append(r.SellerJSONs, sellerJSON)
}

//Update Seller Struct with specific ID and leave empty parameters unchanged
func (r *SellerRepo) UpdateSeller(id string, newSeller SellerJSON) error {
	for i,d := range r.SellerJSONs {
		if d.ID == id {
			if newSeller.DefaultDest != "" {r.SellerJSONs[i].DefaultDest = newSeller.DefaultDest}
			if newSeller.TotalAvailableHashRate != 0 {r.SellerJSONs[i].TotalAvailableHashRate = newSeller.TotalAvailableHashRate}
			if newSeller.UnusedHashRate != 0 {r.SellerJSONs[i].UnusedHashRate = newSeller.UnusedHashRate}
			if newSeller.Contracts != nil {r.SellerJSONs[i].Contracts = newSeller.Contracts}
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

func ConvertSellerJSONtoSellerMSG(seller SellerJSON, msg msgbus.Seller) msgbus.Seller {
	msg.ID = msgbus.SellerID(seller.ID)
	msg.DefaultDest = msgbus.DestID(seller.DefaultDest)
	msg.TotalAvailableHashRate = seller.TotalAvailableHashRate
	msg.UnusedHashRate = seller.UnusedHashRate
	msg.Contracts = seller.Contracts

	return msg	
}

func ConvertSellerMSGtoSellerJSON(msg msgbus.Seller) (seller SellerJSON) {
	seller.ID = string(msg.ID)
	seller.DefaultDest = string(msg.DefaultDest)
	seller.TotalAvailableHashRate = msg.TotalAvailableHashRate
	seller.UnusedHashRate = msg.UnusedHashRate
	seller.Contracts = msg.Contracts

	return seller	
}