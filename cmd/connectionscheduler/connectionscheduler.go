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
// func (cs *ConnectionScheduler) delContract(channel []msgbus.EventChan, c msgbus.EventChan) (result []msgbus.EventChan) {
//
//	result = channel
//
//	for i, echan := range channel {
//		if c == echan {
//			length := len(channel)
//			if length == 1 {
//				result = channel[:0]
//			} else if i == length {
//				result = channel[:length-1]
//			} else {
//				result[i] = channel[length-1]
//				result = channel[:length-1]
//			}
//		}
//	}
//	return result
//}

//------------------------------------------
//
//------------------------------------------
//func (cs *ConnectionScheduler) addContract(channel []msgbus.EventChan, c msgbus.EventChan) (result []msgbus.EventChan) {
//	result = append(channel, c)
//	return result
//}

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

	fmt.Printf(" Connection Scheduler Starting\n")

	cs.Contracts = make(map[msgbus.ContractID]msgbus.Contract)

	// Monitor New Contracts
	contractEventChan := cs.ps.NewEventChan()
	err = cs.ps.Sub(msgbus.ContractMsg, "", contractEventChan)
	if err != nil {
		return err
	}

	go cs.newContractHandler(contractEventChan)

	fmt.Printf("Connection Scheduler Started\n")

	return err
}

//------------------------------------------
//
//------------------------------------------
func (cs *ConnectionScheduler) newContractHandler(ch msgbus.EventChan) {

	for event := range ch {

		id := msgbus.ContractID(event.ID)

		switch event.EventType {
		case msgbus.SubscribedEvent:
			fmt.Printf("Contract subscribed:%v\n", event)
			continue

		case msgbus.PublishEvent:
			// Is this a new contract?
			if _, ok := cs.Contracts[id]; !ok {
				cs.Contracts[id] = event.Data.(msgbus.Contract)

				e1, err := cs.ps.SubWait(msgbus.ContractMsg, event.ID, ch)
				if err != nil {
					panic(fmt.Sprintf("SubWait failed: %s", err))
				}
				if e1.EventType != msgbus.SubscribedEvent {
					panic("Wrong event type")
				}

			} else {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" got PubEvent, but already had the ID: %v", event))
			}

		case msgbus.DeleteEvent:
			fallthrough
		case msgbus.UnsubscribedEvent:
			fmt.Printf("Contract Event:%v\n", event)

			if _, ok := cs.Contracts[id]; ok {
				delete(cs.Contracts, id)
			} else {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" got UnsubscribeEvent, but dont have the ID: %v", event))
			}

		case msgbus.UpdateEvent:
			if _, ok := cs.Contracts[id]; !ok {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" got contract ID does not exist: %v", event))
			}

			// Update the current contract data
			currentContract := cs.Contracts[id]
			cs.Contracts[id] = event.Data.(msgbus.Contract)

			if currentContract.State != event.Data.(msgbus.Contract).State {
				switch event.Data.(msgbus.Contract).State {
				case msgbus.ContAvailableState:
					fmt.Sprintf(lumerinlib.FileLine()+" Found Available Contract: %v", event)

				case msgbus.ContActiveState:
					fmt.Sprintf(lumerinlib.FileLine()+" Found Active Contract: %v", event)

				case msgbus.ContRunningState:
					if currentContract.State != msgbus.ContRunningState {
						cs.ContractRunning(id)
					}

					// Set Contract to running, and rework all of the miners
				case msgbus.ContCompleteState:
					if currentContract.State != msgbus.ContCompleteState {
						cs.ContractComplete(id)
					}

					// Set Contract to Complete, and reset all the miners
				default:
					panic(fmt.Sprintf(lumerinlib.FileLine()+" got bad State: %v", event))
				}

			} else {
				panic(fmt.Sprintf(lumerinlib.FileLine()+" got contract change with now state change: %v", event))

			}

		default:
			panic(fmt.Sprintf(lumerinlib.FileLine()+" got Event: %v", event))
		}

	}

	fmt.Printf(lumerinlib.Funcname() + " Exiting\n")

}

func (cs *ConnectionScheduler) ContractRunning(id msgbus.ContractID) {
	// Calculate the new Target

	return

	// Find all of the online miners point them to new target
	event, err := cs.ps.GetWait(msgbus.MinerMsg, "")
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s", err))
	}

	if event.EventType != msgbus.GetIndexEvent {
		panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", event))
	}

	//	if 0 < len(event.Data.(msgbus.IDIndex)) {
	//		for _, i := range event.Data.(msgbus.IDIndex) {
	//			e, err := cs.ps.GetWait(msgbus.MinerMsg, msgbus.MinerID(i))
	//			if err != nil {
	//				panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", event))
	//			}
	//
	//		}
	//	}

}

func (cs *ConnectionScheduler) ContractComplete(id msgbus.ContractID) {

	return
}
