package msgdata

import (
	"errors"
	"strconv"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

// Struct of Connection parameters in JSON
type ConnectionJSON struct {
	ID        	string `json:"ID"`
	Miner     	string `json:"Miner"`
	Dest  		string `json:"Dest"`
	State     	string `json:"State"`
	TotalHash 	string `json:"Total Hash"`
	StartDate 	string `json:"Start Date"`
}

//Struct that stores slice of all JSON Connection structs in Repo
type ConnectionRepo struct {
	ConnectionJSONs []ConnectionJSON
}

//Initialize Repo with empty slice of JSON Connection structs
func NewConnection() *ConnectionRepo {
	return &ConnectionRepo{}
}

//Return all Connection Structs in Repo
func (r *ConnectionRepo) GetAllConnections() []ConnectionJSON {
	return r.ConnectionJSONs
}

//Return Connection Struct by ID
func (r *ConnectionRepo) GetConnection(id string) (ConnectionJSON, error) {
	for i,c := range r.ConnectionJSONs {
		if c.ID == id {
			return r.ConnectionJSONs[i], nil
		}
	}
	return r.ConnectionJSONs[0], errors.New("ID not found")
}

//Add a new Connection Struct to to Repo
func (r *ConnectionRepo) AddConnection(conn ConnectionJSON) {
	r.ConnectionJSONs = append(r.ConnectionJSONs, conn)
}

//Converts Connection struct from msgbus to JSON struct and adds it to Repo
func (r *ConnectionRepo) AddConnectionFromMsgBus(conn msgbus.Connection) {
	var connJSON ConnectionJSON

	connJSON.ID = string(conn.ID)
	connJSON.Miner = string(conn.Miner)
	connJSON.Dest = string(conn.Dest)
	connJSON.State = 	string(conn.State)
	connJSON.TotalHash = strconv.Itoa(conn.TotalHash)
	connJSON.StartDate = conn.StartDate.String()
	
	r.ConnectionJSONs = append(r.ConnectionJSONs, connJSON)
}

//Update Connection Struct with specific ID and leave empty parameters unchanged
func (r *ConnectionRepo) UpdateConnection(id string, newConnection ConnectionJSON) error {
	for i,c := range r.ConnectionJSONs {
		if c.ID == id {
			if newConnection.Miner != "" {r.ConnectionJSONs[i].Miner = newConnection.Miner}
			if newConnection.Dest != "" {r.ConnectionJSONs[i].Dest = newConnection.Dest}
			if newConnection.State != "" {r.ConnectionJSONs[i].State = newConnection.State}
			if newConnection.TotalHash != "" {r.ConnectionJSONs[i].TotalHash = newConnection.TotalHash}
			if newConnection.StartDate != "" {r.ConnectionJSONs[i].StartDate = newConnection.StartDate}

			return nil
		}
	}
	return errors.New("ID not found")
}

//Delete Connection Struct with specific ID
func (r *ConnectionRepo) DeleteConnection(id string) error {
	for i,c := range r.ConnectionJSONs {
		if c.ID == id {
			r.ConnectionJSONs = append(r.ConnectionJSONs[:i], r.ConnectionJSONs[i+1:]...)

			return nil
		}
	}
	return errors.New("ID not found")
}

func ConvertConnectionJSONtoConnectionMSG(conn ConnectionJSON, msg msgbus.Connection) msgbus.Connection {	
	msg.ID = msgbus.ConnectionID(conn.ID)
	msg.Miner = msgbus.MinerID(conn.Miner)
	msg.Dest = msgbus.DestID(conn.Dest)
	msg.State = msgbus.ConnectionState(conn.State)
	msg.TotalHash,_ = strconv.Atoi(conn.TotalHash)
	msg.StartDate,_ = time.Parse(conn.StartDate, "000000")

	return msg	
}