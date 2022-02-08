package connectionscheduler

import (
	"context"
	"fmt"
	"math"
	"sort"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

const HASHRATE_TOLERANCE = .10 

type ConnectionScheduler struct {
	ps        			*msgbus.PubSub
	Contracts 			map[msgbus.ContractID]msgbus.Contract
	ReadyMiners			map[msgbus.MinerID]msgbus.Miner // miners pointed to default pool
	BusyMiners			map[msgbus.MinerID]msgbus.Miner // miners pointed to target dest set by contract
	nodeOperator    	msgbus.NodeOperator
	minerUpdatedChans	map[msgbus.ContractID]chan bool
	contractClosedChans	map[msgbus.ContractID]chan bool
	ctx		  			context.Context
}

// implement golang sort interface
type Miner struct {
	id			msgbus.MinerID
	hashrate	int
}
type MinerList []Miner
func (m MinerList) Len() int	{ return len(m) }
func (m MinerList) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MinerList) Less(i, j int) bool { return m[i].hashrate < m[j].hashrate }

//------------------------------------------
//
//------------------------------------------
func New(ctx *context.Context, ps *msgbus.PubSub, nodeOperator *msgbus.NodeOperator) (cs *ConnectionScheduler, err error) {
	cs = &ConnectionScheduler{
		ps: ps,
		nodeOperator: *nodeOperator,
		ctx: *ctx,
	}
	cs.Contracts = make(map[msgbus.ContractID]msgbus.Contract)
	cs.ReadyMiners = make(map[msgbus.MinerID]msgbus.Miner)
	cs.BusyMiners = make(map[msgbus.MinerID]msgbus.Miner)
	cs.minerUpdatedChans = make(map[msgbus.ContractID]chan bool)
	cs.contractClosedChans = make(map[msgbus.ContractID]chan bool)
	return cs, err
}

//------------------------------------------
//
//------------------------------------------
func (cs *ConnectionScheduler) Start() (err error) {
	fmt.Printf("Connection Scheduler Starting\n")

	// Update connection scheduler with current contracts
	event, err := cs.ps.GetWait(msgbus.ContractMsg, "")
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
	}
	contracts := event.Data.(msgbus.IDIndex)
	for i := range contracts {
		event, err := cs.ps.GetWait(msgbus.ContractMsg, contracts[i])
		if err != nil {
			panic(fmt.Sprintf("Failed to get miner, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err))
		}
		contract := event.Data.(msgbus.Contract)
		cs.Contracts[msgbus.ContractID(contracts[i])] = contract
	}

	// Monitor New Contracts
	contractEventChan := cs.ps.NewEventChan()
	err = cs.ps.Sub(msgbus.ContractMsg, "", contractEventChan)
	if err != nil {
		return err
	}
	go cs.goContractHandler(contractEventChan)

	// Update connection scheduler with current miners in online state
	miners, err := cs.ps.MinerGetAllWait()
	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
	}
	for i := range miners {
		miner, err := cs.ps.MinerGetWait(miners[i])
		if err != nil {
			panic(fmt.Sprintf("Failed to get miner, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err))
		}
		if miner.State == msgbus.OnlineState {
			if miner.Dest == cs.nodeOperator.DefaultDest {
				cs.ReadyMiners[msgbus.MinerID(miners[i])] = *miner
			} else {
				cs.BusyMiners[msgbus.MinerID(miners[i])] = *miner
			}
		}
	}
	//cs.calculateHashrateAvailability()

	// Monitor New OnlineMiners
	minerEventChan := cs.ps.NewEventChan()
	err = cs.ps.Sub(msgbus.MinerMsg, "", minerEventChan)
	if err != nil {
		return err
	}
	go cs.goMinerHandler(minerEventChan)

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
					if currentContract.State != msgbus.ContAvailableState {
						cs.contractClosedChans[id]<-true
						close(cs.minerUpdatedChans[id])
						close(cs.contractClosedChans[id])
						delete(cs.minerUpdatedChans, id)
						delete(cs.contractClosedChans, id)
					}

				case msgbus.ContRunningState:
					fmt.Printf(lumerinlib.FileLine()+" Found Running Contract: %v\n", event)

					if currentContract.State != msgbus.ContRunningState {
						cs.minerUpdatedChans[id] = make(chan bool, 5)
						cs.contractClosedChans[id] = make(chan bool, 5)
						go cs.ContractRunning(id)
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
// Monitors new miner publish events, and then subscribes to the miners
// Then monitors the update events on the miners, and keeps track of their
// hashrate
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) goMinerHandler(ch msgbus.EventChan) {
	for {
		select {
		case <-cs.ctx.Done():
			fmt.Println("Cancelling current connection scheduler context: cancelling minerMonitor go routine")
			return
		case event := <-ch:
			id := msgbus.MinerID(event.ID)

			loop:
			switch event.EventType {
				//
				// Publish Event
				//
			case msgbus.PublishEvent:
				fmt.Printf("Got PublishEvent: %v\n", event)
				miner := event.Data.(msgbus.Miner)

				if miner.State != msgbus.OnlineState {
					break loop
				}

				switch miner.Dest {
				case cs.nodeOperator.DefaultDest:
					if _, ok := cs.ReadyMiners[id]; !ok {
						cs.ReadyMiners[id] = miner
						for _,v := range cs.Contracts {
							if v.State == msgbus.ContRunningState {
								cs.minerUpdatedChans[v.ID]<-true
							}
						}
					} else {
						panic(fmt.Sprintf("Got PubEvent, but already had the ID: %v\n", event))
					}
				default:
					if _, ok := cs.BusyMiners[id]; !ok {
						cs.BusyMiners[id] = miner
						for _,v := range cs.Contracts {
							if v.State == msgbus.ContRunningState {
								cs.minerUpdatedChans[v.ID]<-true
							}
						}
					} else {
						panic(fmt.Sprintf("Got PubEvent, but already had the ID: %v\n", event))
					}
				}
				fmt.Println("Ready Miners: ", cs.ReadyMiners)
				fmt.Println("Busy Miners: ", cs.BusyMiners)
	
				//
				// Update Event
				//
			case msgbus.UpdateEvent:
				miner := event.Data.(msgbus.Miner)

				if miner.State != msgbus.OnlineState {
					delete(cs.BusyMiners, id)
					delete(cs.ReadyMiners, id)
					break loop
				}

				fmt.Printf("Miner Update Event:%v\n", event)
				
				switch miner.Dest {
				case cs.nodeOperator.DefaultDest:
					// Update the current miner data
					delete(cs.BusyMiners, id)
					cs.ReadyMiners[id] = miner
					for _,v := range cs.Contracts {
						if v.State == msgbus.ContRunningState && !miner.CsMinerHandlerIgnore{
							cs.minerUpdatedChans[v.ID]<-true
						}
					}
				default:
					// Update the current miner data
					delete(cs.ReadyMiners, id)
					cs.BusyMiners[id] = miner
					for _,v := range cs.Contracts {
						if v.State == msgbus.ContRunningState && !miner.CsMinerHandlerIgnore{
							cs.minerUpdatedChans[v.ID]<-true
						}
					}
				}
				fmt.Println("Ready Miners: ", cs.ReadyMiners)
				fmt.Println("Busy Miners: ", cs.BusyMiners)
				
				//
				// Unpublish Event
				//
			case msgbus.UnpublishEvent:
				fmt.Printf("Miner Unpublish/Unsubscribe Event:%v\n", event)
				delete(cs.ReadyMiners, id)
				delete(cs.BusyMiners, id)
				for _,v := range cs.Contracts {
					if v.State == msgbus.ContRunningState {
						cs.minerUpdatedChans[v.ID]<-true
					}
				}

			default:
				fmt.Printf("Got Event: %v\n", event)
			}
		}	
	}
}

//------------------------------------------------------------------------
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) ContractRunning(contractId msgbus.ContractID) {

	fmt.Printf(lumerinlib.FileLine()+" Contract Running: %s\n", contractId)

	// Calculate the new Target
	event, err := cs.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(contractId))
	if err != nil {
		panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", event))
	}
	contract := event.Data.(msgbus.Contract)

	availableHashrate := cs.calculateHashrateAvailability(contractId)
	
	MIN := int(float32(contract.Speed) - float32(contract.Speed)*HASHRATE_TOLERANCE) 

	if availableHashrate >= MIN  {
		cs.SetMinerTarget(contract)
	}

	// contractEventChan := cs.ps.NewEventChan()
	// err = cs.ps.Sub(msgbus.ContractMsg, msgbus.IDString(id), contractEventChan)
	// if err != nil {
	// 	panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", event))
	// }

	// minerEventChan := cs.ps.NewEventChan()
	// err = cs.ps.Sub(msgbus.MinerMsg, "", minerEventChan)
	// if err != nil {
	// 	panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", event))
	// }

	minerMapUpdated := cs.minerUpdatedChans[contractId]
	contractClosed := cs.contractClosedChans[contractId]
	for {
		select {
		case <-cs.ctx.Done():
			fmt.Println("Cancelling current connection scheduler context: cancelling contract running go routine for contract: ", contract.ID)
			return
		case <-contractClosed:
			fmt.Println("Contract state switched to complete: cancelling contract running go routine for contract: ", contract.ID)
			return
		case <-minerMapUpdated:
			availableHashrate = cs.calculateHashrateAvailability(contractId)

			if availableHashrate <= MIN {
				fmt.Println("Not enough available hashrate to fulfill contract: ", contract.ID)
				// free up busy miners with this contract id
				for k, v := range cs.BusyMiners {
					if v.Contract == contract.ID {
						err := cs.ps.MinerRemoveContractWait(k, cs.nodeOperator.DefaultDest, true)
						if err != nil {
							panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
						}
					}
				}
				return
			}
			cs.SetMinerTarget(contract)
		}	
	}
}

func (cs *ConnectionScheduler) SetMinerTarget(contract msgbus.Contract) {
	destid := contract.Dest
	promisedHashrate := contract.Speed

	if destid == "" {
		panic(fmt.Sprint(lumerinlib.FileLine() + " Error DestID is empty"))
	}

	// sort miners by hashrate from least to greatest
	sortedReadyMiners := cs.sortMinersByHashrate(contract.ID)
	fmt.Println("Sorted Miners: ", sortedReadyMiners)

	// find all miner combinations that add up to promised hashrate
	minerCombinations := findSubsets(sortedReadyMiners, promisedHashrate)
	if minerCombinations == nil {
		fmt.Println("No valid miner combinations")
		// free up busy miners with this contract id
		for k, v := range cs.BusyMiners {
			if v.Contract == contract.ID {
				err := cs.ps.MinerRemoveContractWait(k, cs.nodeOperator.DefaultDest, true)
				if err != nil {
					panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
				}
			}
		}
		return
	}
	fmt.Println("Valid Miner Combinations: ", minerCombinations)

	// find best combination of miners
	minerCombination := bestCombination(minerCombinations, promisedHashrate)
	fmt.Println("Best Miner Combination: ", minerCombination)


	// set target destination for miners in optimal miner combination
	for _,v := range minerCombination {
		err := cs.ps.MinerSetContractWait(v.id, contract.ID, true)
		if err != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
		}
	}

	// update busy miners map with new dests based on new miner combination
	newBusyMinerMap := make(map[msgbus.MinerID]Miner)
	for _,v := range minerCombination {
		newBusyMinerMap[v.id] = v
	}
	for k,v := range cs.BusyMiners {
		if _,ok := newBusyMinerMap[v.ID]; !ok {
			if v.Contract == contract.ID {
				err := cs.ps.MinerRemoveContractWait(k, cs.nodeOperator.DefaultDest, true)
				if err != nil {
					panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
				}
			}	
		}
	}
}

func (cs *ConnectionScheduler) calculateHashrateAvailability(id msgbus.ContractID) (availableHashrate int) {
	for _,v := range cs.ReadyMiners {
		availableHashrate += v.CurrentHashRate
	}
	for _,v := range cs.BusyMiners {
		if v.Contract == id {
			availableHashrate += v.CurrentHashRate
		}
	}
	fmt.Println("Available Hashrate", availableHashrate)
	return availableHashrate
}

func (cs *ConnectionScheduler) sortMinersByHashrate(contractId msgbus.ContractID) (m MinerList){
	m = make(MinerList, 0)

	for k, v := range cs.ReadyMiners {
		m = append(m, Miner{k, v.CurrentHashRate})
	}

	// include busy miners that are already associated with contract
	for k, v := range cs.BusyMiners {
		if v.Contract == contractId {
			m = append(m, Miner{k, v.CurrentHashRate})
		}
	}

	sort.Sort(m)
	return m
}

func sumSubsets(sortedMiners MinerList, n int, targetHashrate int) (m MinerList) {
	// Create new array with size equal to sorted miners array to create binary array as per n(decimal number)
    x := make([]int, sortedMiners.Len())
    j := sortedMiners.Len() - 1

	// Convert the array into binary array
	for n > 0 {
		x[j] = n % 2
		n = n/2
		j--
	}

	sum := 0

	// Calculate the sum of this subset
	for i := range sortedMiners {
		if x[i] == 1 {
			sum += sortedMiners[i].hashrate
		}
	}

	MIN := int(float32(targetHashrate) - float32(targetHashrate)*HASHRATE_TOLERANCE) 
	MAX := int(float32(targetHashrate) + float32(targetHashrate)*HASHRATE_TOLERANCE) 

	// if sum is within target hashrate bounds, subset was found
	if (sum >= MIN && sum <= MAX) {
		for i := range sortedMiners {
			if x[i] == 1 {
				m = append(m, sortedMiners[i])
			}
		}
		return m
	}

	return nil
}

// find subsets of list of miners whose hashrate sum equal the target hashrate
func findSubsets(sortedMiners MinerList, targetHashrate int) (minerCombinations []MinerList){
	// Calculate total number of subsets
	tot := math.Pow(2, float64(sortedMiners.Len()))

	for i:=0; i<int(tot); i++ {
		m := sumSubsets(sortedMiners, i, targetHashrate)
		if m != nil {
			minerCombinations = append(minerCombinations, m)
		}
	}
	return minerCombinations
}

func bestCombination(minerCombinations []MinerList, targetHashrate int) MinerList {
	hashrates := make([]int, len(minerCombinations))
	numMiners := make([]int, len(minerCombinations))

	// find hashrate and number of miners in each combination
	for i := range minerCombinations {
		miners := minerCombinations[i]
		totalHashRate := 0
		num := 0
		for j := range miners {
			totalHashRate += miners[j].hashrate
			num++
		}
		hashrates[i] = totalHashRate
		numMiners[i] = num
	}

	// find combination closest to target hashrate
	index := 0
	for i := range hashrates {
		res1 := math.Abs(float64(targetHashrate) - float64(hashrates[index]))
		res2 := math.Abs(float64(targetHashrate) - float64(hashrates[i]))
		if res1 > res2 {
			index = i
		}
	}

	// if duplicate exists choose the one with the least number of miners
	newIndex := index
	for i := range hashrates {
		if hashrates[i] == hashrates[index] && numMiners[i] < numMiners[newIndex] {
			newIndex = i
		}	
	}

	return minerCombinations[newIndex]
}