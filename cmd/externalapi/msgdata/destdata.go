package msgdata

import (
	"errors"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

// Struct of Dest parameters in JSON
type DestJSON struct {
	ID			string	`json:"id"`
	NetUrl 		string	`json:"netUrl"`
}

//Struct that stores slice of all JSON Dest structs in Repo
type DestRepo struct {
	DestJSONs []DestJSON
}

//Initialize Repo with empty slice of JSON Dest structs
func NewDest() *DestRepo {
	return &DestRepo{
		DestJSONs: []DestJSON{},
	}
}

//Return all Dest Structs in Repo
func (r *DestRepo) GetAllDests() []DestJSON {
	return r.DestJSONs
}

//Return Dest Struct by ID
func (r *DestRepo) GetDest(id string) (DestJSON, error) {
	for i,d := range r.DestJSONs {
		if d.ID == id {
			return r.DestJSONs[i], nil
		}
	}
	return r.DestJSONs[0], errors.New("ID not found")
}

//Add a new Dest Struct to to Repo
func (r *DestRepo) AddDest(dest DestJSON) {
	r.DestJSONs = append(r.DestJSONs, dest)
}

//Converts Dest struct from msgbus to JSON struct and adds it to Repo
func (r *DestRepo) AddDestFromMsgBus(dest msgbus.Dest) {
	var destJSON DestJSON
	
	destJSON.ID = string(dest.ID)
	destJSON.NetUrl = string(dest.NetUrl)
	
	r.DestJSONs = append(r.DestJSONs, destJSON)
}

//Update Dest Struct with specific ID and leave empty parameters unchanged
func (r *DestRepo) UpdateDest(id string, newDest DestJSON) error {
	for i,d := range r.DestJSONs {
		if d.ID == id {
			if newDest.NetUrl != "" {r.DestJSONs[i].NetUrl = newDest.NetUrl}

			return nil
		}
	}
	return errors.New("ID not found")
}

//Delete Dest Struct with specific ID
func (r *DestRepo) DeleteDest(id string) error {
	for i,d := range r.DestJSONs {
		if d.ID == id {
			r.DestJSONs = append(r.DestJSONs[:i], r.DestJSONs[i+1:]...)

			return nil
		}
	}
	return errors.New("ID not found")
}

func ConvertDestJSONtoDestMSG(dest DestJSON, msg msgbus.Dest) msgbus.Dest {
	msg.ID = msgbus.DestID(dest.ID)
	msg.NetUrl = msgbus.DestNetUrl(dest.NetUrl)

	return msg	
}