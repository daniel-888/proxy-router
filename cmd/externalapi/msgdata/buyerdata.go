package msgdata

import (
	"errors"
	"fmt"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

//Struct of Buyer parameters in JSON 
type BuyerJSON struct {
	ID          string 										`json:"id"`
	DefaultDest string 										`json:"defaultDest"`
	Contracts	map[msgbus.ContractID]msgbus.ContractState	`json:"contracts"`
}

//Struct that stores slice of all JSON Buyer structs in Repo
type BuyerRepo struct {
	BuyerJSONs []BuyerJSON
	Ps          *msgbus.PubSub
}

//Initialize Repo with empty slice of JSON Buyer structs
func NewBuyer(ps *msgbus.PubSub) *BuyerRepo {
	return &BuyerRepo{
		BuyerJSONs:	[]BuyerJSON{},
		Ps:				ps,
	}
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
func (r *BuyerRepo) AddBuyerFromMsgBus(buyerID msgbus.BuyerID, buyer msgbus.Buyer) {
	var buyerJSON BuyerJSON

	buyerJSON.ID = string(buyerID)
	buyerJSON.DefaultDest = string(buyer.DefaultDest)
	buyerJSON.Contracts = buyer.Contracts
	
	r.BuyerJSONs = append(r.BuyerJSONs, buyerJSON)
}

//Update Buyer Struct with specific ID and leave empty parameters unchanged
func (r *BuyerRepo) UpdateBuyer(id string, newBuyer BuyerJSON) error {
	for i,d := range r.BuyerJSONs {
		fmt.Println(d.ID)
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

//Subscribe to events for buyer msgs on msgbus to update API repos with data
func (r *BuyerRepo) SubscribeToBuyerMsgBus() {
	buyerCh := r.Ps.NewEventChan()
	
	// add existing buyers to api repo
	event, err := r.Ps.GetWait(msgbus.BuyerMsg, "")
	if err != nil {
		panic(fmt.Sprintf("Getting Buyers Failed: %s", err))
	}
	buyers := event.Data.(msgbus.IDIndex)
	if len(buyers) > 0 {
		for i := range buyers {
			event, err = r.Ps.GetWait(msgbus.BuyerMsg, msgbus.IDString(buyers[i]))
			if err != nil {
				panic(fmt.Sprintf("Getting Buyer Failed: %s", err))
			}
			buyer := event.Data.(msgbus.Buyer)
			r.AddBuyerFromMsgBus(msgbus.BuyerID(buyers[i]), buyer)
		}
	}

	event, err = r.Ps.SubWait(msgbus.BuyerMsg, "", buyerCh)
	if err != nil {
		panic(fmt.Sprintf("SubWait failed: %s\n", err))
	}
	if event.EventType != msgbus.SubscribedEvent {
		panic(fmt.Sprintf("Wrong event type %v\n", event))
	}

	for event = range buyerCh {
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
			buyerID := msgbus.BuyerID(event.ID)
			buyer := event.Data.(msgbus.Buyer)
			r.AddBuyerFromMsgBus(buyerID, buyer)
			
			//
			// Delete/Unpublish Event
			//
		case msgbus.DeleteEvent:
			fallthrough
		case msgbus.UnpublishEvent:
			fmt.Printf(lumerinlib.Funcname()+" Delete/Unpublish Event: %v\n", event)
			buyerID := msgbus.BuyerID(event.ID)
			r.DeleteBuyer(string(buyerID))

			//
			// Update Event
			//
		case msgbus.UpdateEvent:
			fmt.Printf(lumerinlib.Funcname()+" Update Event: %v\n", event)
			buyerID := msgbus.BuyerID(event.ID)
			buyer := event.Data.(msgbus.Buyer)
			buyerJSON := ConvertBuyerMSGtoBuyerJSON(buyer)
			r.UpdateBuyer(string(buyerID), buyerJSON)
			
			//
			// Rut Row...
			//
		default:
			fmt.Printf(lumerinlib.Funcname()+" Got Event: %v\n", event)
		}
	}
}

func ConvertBuyerJSONtoBuyerMSG(buyer BuyerJSON) msgbus.Buyer {
	var msg msgbus.Buyer
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