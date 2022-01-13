package msgdata

import (
	"errors"
	"fmt"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

//Struct of Miner parameters in JSON
type MinerJSON struct {
	ID                      string	`json:"id"`
	Name					string 	`json:"name"`
	IP						string 	`json:"ip"`
	MAC						string 	`json:"mac"`
	State                   string 	`json:"state"`
	Seller                  string 	`json:"seller"`
	Dest                   	string 	`json:"dest"`
	InitialMeasuredHashRate int 	`json:"initialMeasuredHashRate"`
	CurrentHashRate         int 	`json:"currentHashRate"`
}

//Struct that stores slice of all JSON Miner structs in Repo
type MinerRepo struct {
	MinerJSONs []MinerJSON
	ps          *msgbus.PubSub
}

//Initialize Repo with empty slice of JSON Miner structs
func NewMiner(ps *msgbus.PubSub) *MinerRepo {
	return &MinerRepo{
		MinerJSONs: []MinerJSON{},
		ps:			ps,
	}
}

//Return all Miner Structs in Repo
func (r *MinerRepo) GetAllMiners() []MinerJSON {
	return r.MinerJSONs
}

//Return Miner Struct by ID
func (r *MinerRepo) GetMiner(id string) (MinerJSON, error) {
	for i,m := range r.MinerJSONs {
		if m.ID == id {
			return r.MinerJSONs[i], nil
		}
	}
	return r.MinerJSONs[0], errors.New("ID not found")
}

//Add a new Miner Struct to to Repo
func (r *MinerRepo) AddMiner(miner MinerJSON) {
	r.MinerJSONs = append(r.MinerJSONs, miner)
}

//Converts Miner struct from msgbus to JSON struct and adds it to Repo
func (r *MinerRepo) AddMinerFromMsgBus(minerID msgbus.MinerID, miner msgbus.Miner) {
	var minerJSON MinerJSON

	minerJSON.ID = string(minerID)
	minerJSON.State = string(miner.State)
	minerJSON.Seller = string(miner.Seller)
	minerJSON.Dest = string(miner.Dest)
	minerJSON.InitialMeasuredHashRate = miner.InitialMeasuredHashRate
	minerJSON.CurrentHashRate = miner.CurrentHashRate
	
	r.MinerJSONs = append(r.MinerJSONs, minerJSON)
}

//Update Miner Struct with specific ID and leave empty parameters unchanged
func (r *MinerRepo) UpdateMiner(id string, newMiner MinerJSON) error {
	for i,m := range r.MinerJSONs {
		if m.ID == id {
			if newMiner.State != "" {r.MinerJSONs[i].State = newMiner.State}
			if newMiner.Seller != "" {r.MinerJSONs[i].Seller = newMiner.Seller}
			if newMiner.Dest != "" {r.MinerJSONs[i].Dest = newMiner.Dest}
			if newMiner.InitialMeasuredHashRate != 0 {r.MinerJSONs[i].InitialMeasuredHashRate = newMiner.InitialMeasuredHashRate}
			if newMiner.CurrentHashRate != 0 {r.MinerJSONs[i].CurrentHashRate = newMiner.CurrentHashRate}

			return nil
		}
	}
	return errors.New("ID not found")
}

//Delete Miner Struct with specific ID
func (r *MinerRepo) DeleteMiner(id string) error {
	for i,m := range r.MinerJSONs {
		if m.ID == id {
			r.MinerJSONs = append(r.MinerJSONs[:i], r.MinerJSONs[i+1:]...)

			return nil
		}
	}
	return errors.New("ID not found")
}

//Subscribe to events for miner msgs on msgbus to update API repo with data
func (r *MinerRepo) SubscribeToMinerMsgBus() {
	minerCh := r.ps.NewEventChan()
	
	// add existing miners to api repo
	miners, err := r.ps.MinerGetAllWait()
	if err != nil {
		panic(fmt.Sprintf("Getting Miners Failed: %s", err))
	}
	if len(miners) > 0 {
		for i := range miners {
			miner, err := r.ps.MinerGetWait(miners[i])
			if err != nil {
				panic(fmt.Sprintf("Getting Miner Failed: %s", err))
			}
			r.AddMinerFromMsgBus(miners[i], *miner)
		}
	}

	event, err := r.ps.SubWait(msgbus.MinerMsg, "", minerCh)
	if err != nil {
		panic(fmt.Sprintf("SubWait failed: %s\n", err))
	}
	if event.EventType != msgbus.SubscribedEvent {
		panic(fmt.Sprintf("Wrong event type %v\n", event))
	}

	for event = range minerCh {
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
			minerID := msgbus.MinerID(event.ID)
			miner := event.Data.(msgbus.Miner)
			r.AddMinerFromMsgBus(minerID, miner)
			
			//
			// Delete/Unpublish Event
			//
		case msgbus.DeleteEvent:
			fallthrough
		case msgbus.UnpublishEvent:
			fmt.Printf(lumerinlib.Funcname()+" Delete/Unpublish Event: %v\n", event)
			minerID := msgbus.MinerID(event.ID)
			r.DeleteMiner(string(minerID))

			//
			// Update Event
			//
		case msgbus.UpdateEvent:
			fmt.Printf(lumerinlib.Funcname()+" Update Event: %v\n", event)
			minerID := msgbus.MinerID(event.ID)
			miner := event.Data.(msgbus.Miner)
			minerJSON := ConvertMinerMSGtoMinerJSON(miner)
			r.UpdateMiner(string(minerID), minerJSON)
			
			//
			// Rut Row...
			//
		default:
			fmt.Printf(lumerinlib.Funcname()+" Got Event: %v\n", event)
		}
	}
}

func ConvertMinerJSONtoMinerMSG(miner MinerJSON, msg msgbus.Miner) msgbus.Miner {
	msg.ID = msgbus.MinerID(miner.ID)
	msg.State = msgbus.MinerState(miner.State)
	msg.Seller = msgbus.SellerID(miner.Seller)
	msg.Dest = msgbus.DestID(miner.Dest)
	msg.InitialMeasuredHashRate = miner.InitialMeasuredHashRate
	msg.CurrentHashRate = miner.CurrentHashRate

	return msg	
}

func ConvertMinerMSGtoMinerJSON(msg msgbus.Miner) (miner MinerJSON) {
	miner.ID = string(msg.ID)
	miner.Name = msg.Name
	miner.IP = msg.IP
	miner.MAC = msg.MAC
	miner.State = string(msg.State)
	miner.Seller = string(msg.Seller)
	miner.Dest = string(msg.Dest)
	miner.InitialMeasuredHashRate = msg.InitialMeasuredHashRate
	miner.CurrentHashRate = msg.CurrentHashRate

	return miner	
}