package connectionscheduler

import (
	"fmt"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

type ConnectionScheduler struct {
	ps        *msgbus.PubSub
	Contracts map[msgbus.ContractID]msgbus.Contract
}

//------------------------------------------
//
//------------------------------------------
func New(ps *msgbus.PubSub) (cs *ConnectionScheduler, err error) {
	cs = &ConnectionScheduler{
		ps: ps,
	}
	return cs, err
}

//------------------------------------------
//
//------------------------------------------
func (cs *ConnectionScheduler) Start() (err error) {

	fmt.Printf("Connection Scheduler Starting\n")

	cs.Contracts = make(map[msgbus.ContractID]msgbus.Contract)

	// Monitor New Contracts
	contractEventChan := cs.ps.NewEventChan()
	err = cs.ps.Sub(msgbus.ContractMsg, "", contractEventChan)
	if err != nil {
		return err
	}

	go cs.goContractHandler(contractEventChan)

	fmt.Printf("Connection Scheduler Started\n")

	return err
}

//------------------------------------------------------------------------
//
// Monitors new contract publish events, and then subscribes to the contracts
// Then monitors the update events on the contracts, and handles state changes
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) goContractHandler(ch msgbus.EventChan) {

	for event := range ch {

		id := msgbus.ContractID(event.ID)

		switch event.EventType {
		//
		// Subscribe Event
		//
		case msgbus.SubscribedEvent:
			fmt.Printf(lumerinlib.Funcname()+" subscribed:%v\n", event)

			//
			// Publish Event
			//
		case msgbus.PublishEvent:
			// Is this a new contract?

			fmt.Printf(lumerinlib.Funcname()+" PublishEvent: %v\n", event)

			if _, ok := cs.Contracts[id]; !ok {
				cs.Contracts[id] = event.Data.(msgbus.Contract)

				//
				// If the contract is a buyer, contractscheduler can ignore it.
				// Use the existing channel to monitor
				//
				if cs.Contracts[id].IsSeller {
					e1, err := cs.ps.SubWait(msgbus.ContractMsg, event.ID, ch)
					if err != nil {
						panic(fmt.Sprintf(lumerinlib.FileLine()+" SubWait failed: %s\n", err))
					}
					if e1.EventType != msgbus.SubscribedEvent {
						panic(fmt.Sprintf(lumerinlib.FileLine()+" Wrong event type %v\n", e1))
					}
				}

			} else {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" got PubEvent, but already had the ID: %v\n", event))
			}

			//
			// Delete/Unsubscribe Event
			//
		case msgbus.DeleteEvent:
			fallthrough
		case msgbus.UnsubscribedEvent:
			fmt.Printf(lumerinlib.Funcname()+" Contract Delete/Unsubscribe Event: %v\n", event)

			if _, ok := cs.Contracts[id]; ok {
				delete(cs.Contracts, id)
			} else {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" got UnsubscribeEvent, but dont have the ID: %v\n", event))
			}

			//
			// Update Event
			//
		case msgbus.UpdateEvent:

			fmt.Printf(lumerinlib.Funcname()+" UpdateEvent: %v\n", event)

			if _, ok := cs.Contracts[id]; !ok {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" got contract ID does not exist: %v\n", event))
			}

			// Update the current contract data
			currentContract := cs.Contracts[id]
			cs.Contracts[id] = event.Data.(msgbus.Contract)

			if currentContract.State == event.Data.(msgbus.Contract).State {
				fmt.Printf(lumerinlib.FileLine()+" got contract change with no state change: %v\n", event)
			} else {
				switch event.Data.(msgbus.Contract).State {
				case msgbus.ContAvailableState:
					fmt.Printf(lumerinlib.FileLine()+" Found Available Contract: %v\n", event)

				case msgbus.ContRunningState:
					fmt.Printf(lumerinlib.FileLine()+" Found Running Contract: %v\n", event)

					if currentContract.State != msgbus.ContRunningState {
						cs.ContractRunning(id)
					}

					// Set Contract to running, and rework all of the miners

				default:
					panic(fmt.Sprintf(lumerinlib.FileLine()+" got bad State: %v\n", event))
				}

			}

			//
			// Rut Row...
			//
		default:
			panic(fmt.Sprintf(lumerinlib.FileLine()+" got Event: %v\n", event))
		}

	}

	fmt.Printf(lumerinlib.Funcname() + " Exiting\n")

}

//------------------------------------------------------------------------
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) ContractRunning(id msgbus.ContractID) {

	fmt.Printf(lumerinlib.FileLine()+" Contract Running: %s\n", id)

	// Calculate the new Target

	contract, err := cs.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(id))
	if err != nil {
		panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", contract))
	}

	destid := contract.Data.(msgbus.Contract).Dest

	if destid == "" {
		panic(fmt.Sprint(lumerinlib.FileLine() + " Error DestID is empty"))
	}

	// Find all of the online miners point them to new target
	miners, err := cs.ps.MinerGetAllWait()

	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
	}

	for _, v := range miners {
		err := cs.ps.MinerSetDestWait(v, destid)
		if err != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
		}
	}
}
