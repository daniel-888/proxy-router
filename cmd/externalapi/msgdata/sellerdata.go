package msgdata

import (
	"errors"
	"fmt"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

//Struct of Seller parameters in JSON 
type SellerJSON struct {
	ID                     	string 										`json:"id"`
	DefaultDest          	string 										`json:"defaultDest"`
	TotalAvailableHashRate 	int 										`json:"totalAvailableHashrate"`
	UnusedHashRate         	int 										`json:"unusedHashRate"`
	Contracts				map[msgbus.ContractID]msgbus.ContractState	`json:"contracts"`
}

//Struct that stores slice of all JSON Seller structs in Repo
type SellerRepo struct {
	SellerJSONs []SellerJSON
	Ps          *msgbus.PubSub
}

//Initialize Repo with empty slice of JSON Seller structs
func NewSeller(ps *msgbus.PubSub) *SellerRepo {
	return &SellerRepo{
		SellerJSONs:	[]SellerJSON{},
		Ps:				ps,
	}
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
func (r *SellerRepo) AddSellerFromMsgBus(sellerID msgbus.SellerID, seller msgbus.Seller) {
	var sellerJSON SellerJSON

	sellerJSON.ID = string(sellerID)
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

//Subscribe to events for seller msgs on msgbus to update API repos with data
func (r *SellerRepo) SubscribeToSellerMsgBus() {
	sellerCh := r.Ps.NewEventChan()
	
	// add existing sellers to api repo
	event, err := r.Ps.GetWait(msgbus.SellerMsg, "")
	if err != nil {
		panic(fmt.Sprintf("Getting Sellers Failed: %s", err))
	}
	sellers := event.Data.(msgbus.IDIndex)
	if len(sellers) > 0 {
		for i := range sellers {
			event, err = r.Ps.GetWait(msgbus.SellerMsg, msgbus.IDString(sellers[i]))
			if err != nil {
				panic(fmt.Sprintf("Getting Seller Failed: %s", err))
			}
			seller := event.Data.(msgbus.Seller)
			r.AddSellerFromMsgBus(msgbus.SellerID(sellers[i]), seller)
		}
	}
	
	event, err = r.Ps.SubWait(msgbus.SellerMsg, "", sellerCh)
	if err != nil {
		panic(fmt.Sprintf("SubWait failed: %s\n", err))
	}
	if event.EventType != msgbus.SubscribedEvent {
		panic(fmt.Sprintf("Wrong event type %v\n", event))
	}

	for event = range sellerCh {
		loop:
		switch event.EventType {
		//
		// Subscribe Event
		//
		case msgbus.SubscribedEvent:
			fmt.Printf(lumerinlib.Funcname()+" Subscribe Event: %v\n", event)

			//
			// Publish Event
			//
		case msgbus.PublishEvent:
			fmt.Printf(lumerinlib.Funcname()+" Publish Event: %v\n", event)
			sellerID := msgbus.SellerID(event.ID)

			// do not push to api repo if it already exists
			for i := range r.SellerJSONs {
				if r.SellerJSONs[i].ID == string(sellerID) {
					break loop
				}
			}
			seller := event.Data.(msgbus.Seller)
			r.AddSellerFromMsgBus(sellerID, seller)
			
			//
			// Delete/Unpublish Event
			//
		case msgbus.DeleteEvent:
			fallthrough
		case msgbus.UnpublishEvent:
			fmt.Printf(lumerinlib.Funcname()+" Delete/Unpublish Event: %v\n", event)
			sellerID := msgbus.SellerID(event.ID)
			r.DeleteSeller(string(sellerID))

			//
			// Update Event
			//
		case msgbus.UpdateEvent:
			fmt.Printf(lumerinlib.Funcname()+" Update Event: %v\n", event)
			sellerID := msgbus.SellerID(event.ID)
			seller := event.Data.(msgbus.Seller)
			sellerJSON := ConvertSellerMSGtoSellerJSON(seller)
			r.UpdateSeller(string(sellerID), sellerJSON)
			
			//
			// Rut Row...
			//
		default:
			fmt.Printf(lumerinlib.Funcname()+" Got Event: %v\n", event)
		}
	}
}

func ConvertSellerJSONtoSellerMSG(seller SellerJSON) msgbus.Seller {
	var msg msgbus.Seller

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