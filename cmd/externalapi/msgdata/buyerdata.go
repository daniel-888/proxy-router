package msgdata

import (
	"errors"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

//Struct of Buyer parameters in JSON 
type BuyerJSON struct {
	ID          string 										`json:"id"`
	DefaultDest string 										`json:"destID"`
	Contracts	map[msgbus.ContractID]msgbus.ContractState	`json:"contracts"`
}

//Struct that stores slice of all JSON Buyer structs in Repo
type BuyerRepo struct {
	BuyerJSONs []BuyerJSON
}

//Initialize Repo with empty slice of JSON Buyer structs
func NewBuyer() *BuyerRepo {
	return &BuyerRepo{}
}

//Return all Buyer Structs in Repo
func (r *BuyerRepo) GetAllBuyers() []BuyerJSON {
	return r.BuyerJSONs
}

//Return Buyer Struct by ID
func (r *BuyerRepo) GetBuyer(id string) (BuyerJSON, error) {
	for i,d := range r.BuyerJSONs {
		if d.ID == id {
			return r.BuyerJSONs[i], nil
		}
	}
	return r.BuyerJSONs[0], errors.New("ID not found")
}

//Add a new Buyer Struct to to Repo
func (r *BuyerRepo) AddBuyer(buyer BuyerJSON) {
	r.BuyerJSONs = append(r.BuyerJSONs, buyer)
}

//Converts Buyer struct from msgbus to JSON struct and adds it to Repo
func (r *BuyerRepo) AddBuyerFromMsgBus(buyer msgbus.Buyer) {
	var buyerJSON BuyerJSON

	buyerJSON.ID = string(buyer.ID)
	buyerJSON.DefaultDest = string(buyer.DefaultDest)
	buyerJSON.Contracts = buyer.Contracts
	
	r.BuyerJSONs = append(r.BuyerJSONs, buyerJSON)
}

//Update Buyer Struct with specific ID and leave empty parameters unchanged
func (r *BuyerRepo) UpdateBuyer(id string, newBuyer BuyerJSON) error {
	for i,d := range r.BuyerJSONs {
		if d.ID == id {
			if newBuyer.DefaultDest != "" {r.BuyerJSONs[i].DefaultDest = newBuyer.DefaultDest}
			if newBuyer.Contracts != nil {r.BuyerJSONs[i].Contracts = newBuyer.Contracts}
			return nil
		}
	}
	return errors.New("ID not found")
}

//Delete Buyer Struct with specific ID
func (r *BuyerRepo) DeleteBuyer(id string) error {
	for i,d := range r.BuyerJSONs {
		if d.ID == id {
			r.BuyerJSONs = append(r.BuyerJSONs[:i], r.BuyerJSONs[i+1:]...)

			return nil
		}
	}
	return errors.New("ID not found")
}

func ConvertBuyerJSONtoBuyerMSG(buyer BuyerJSON, msg msgbus.Buyer) msgbus.Buyer {
	msg.ID = msgbus.BuyerID(buyer.ID)
	msg.DefaultDest = msgbus.DestID(buyer.DefaultDest)
	msg.Contracts = buyer.Contracts

	return msg	
}

func ConvertBuyerMSGtoBuyerJSON(msg msgbus.Buyer) (buyer BuyerJSON) {
	buyer.ID = string(msg.ID)
	buyer.DefaultDest = string(msg.DefaultDest)
	buyer.Contracts = msg.Contracts

	return buyer	
}