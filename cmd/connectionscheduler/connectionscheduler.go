package connectionscheduler

import (
	"context"
	"fmt"
	"log"
	"math"
	"sort"
	"sync"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

const HASHRATE_TOLERANCE = .10

type ContractsMap struct {
	sync.RWMutex
	c map[msgbus.ContractID]msgbus.Contract
}

type MinersMap struct {
	sync.RWMutex
	m map[msgbus.MinerID]msgbus.Miner
}

type UpdateChansMap struct {
	sync.RWMutex
	u map[msgbus.ContractID]chan bool
}

type ConnectionScheduler struct {
	ps                  *msgbus.PubSub
	Contracts           ContractsMap
	ReadyMiners         MinersMap // miners with no contract
	BusyMiners          MinersMap // miners fulfilling a contract
	nodeOperator        msgbus.NodeOperator
	minerUpdatedChans   UpdateChansMap
	contractClosedChans UpdateChansMap
	ctx                 context.Context
}

// implement golang sort interface
type Miner struct {
	id       msgbus.MinerID
	hashrate int
}
type MinerList []Miner

func (m MinerList) Len() int           { return len(m) }
func (m MinerList) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MinerList) Less(i, j int) bool { return m[i].hashrate < m[j].hashrate }

// Get, set, exists, and delete methods for thread safe maps
func (r *ContractsMap) Get(key msgbus.ContractID) msgbus.Contract {
	r.RLock()
	defer r.RUnlock()
	return r.c[key]
}

func (r *ContractsMap) GetAll() (contracts []msgbus.Contract) {
	r.RLock()
	defer r.RUnlock()
	for _, v := range r.c {
		contracts = append(contracts, v)
	}
	return contracts
}

func (r *ContractsMap) Set(key msgbus.ContractID, val msgbus.Contract) {
	r.Lock()
	defer r.Unlock()
	r.c[key] = val
}

func (r *ContractsMap) Exists(key msgbus.ContractID) bool {
	r.RLock()
	defer r.RUnlock()
	_, ok := r.c[key]
	return ok
}

func (r *ContractsMap) Delete(key msgbus.ContractID) {
	r.Lock()
	defer r.Unlock()
	delete(r.c, key)
}

func (r *MinersMap) Get(key msgbus.MinerID) msgbus.Miner {
	r.RLock()
	defer r.RUnlock()
	return r.m[key]
}

func (r *MinersMap) GetAll() (miners []msgbus.Miner) {
	r.RLock()
	defer r.RUnlock()
	for _, v := range r.m {
		miners = append(miners, v)
	}
	return miners
}

func (r *MinersMap) Set(key msgbus.MinerID, val msgbus.Miner) {
	r.Lock()
	defer r.Unlock()
	r.m[key] = val
}

func (r *MinersMap) Exists(key msgbus.MinerID) bool {
	r.RLock()
	defer r.RUnlock()
	_, ok := r.m[key]
	return ok
}

func (r *MinersMap) Delete(key msgbus.MinerID) {
	r.Lock()
	defer r.Unlock()
	delete(r.m, key)
}

func (r *UpdateChansMap) Get(key msgbus.ContractID) chan bool {
	r.RLock()
	defer r.RUnlock()
	return r.u[key]
}

func (r *UpdateChansMap) Set(key msgbus.ContractID, val chan bool) {
	r.Lock()
	defer r.Unlock()
	r.u[key] = val
}

func (r *UpdateChansMap) Exists(key msgbus.ContractID) bool {
	r.RLock()
	defer r.RUnlock()
	_, ok := r.u[key]
	return ok
}

func (r *UpdateChansMap) Delete(key msgbus.ContractID) {
	r.Lock()
	defer r.Unlock()
	delete(r.u, key)
}

//------------------------------------------
//
//------------------------------------------
func New(ctx *context.Context, ps *msgbus.PubSub, nodeOperator *msgbus.NodeOperator) (cs *ConnectionScheduler, err error) {
	cs = &ConnectionScheduler{
		ps:           ps,
		nodeOperator: *nodeOperator,
		ctx:          *ctx,
	}
	cs.Contracts.c = make(map[msgbus.ContractID]msgbus.Contract)
	cs.ReadyMiners.m = make(map[msgbus.MinerID]msgbus.Miner)
	cs.BusyMiners.m = make(map[msgbus.MinerID]msgbus.Miner)
	cs.minerUpdatedChans.u = make(map[msgbus.ContractID]chan bool)
	cs.contractClosedChans.u = make(map[msgbus.ContractID]chan bool)
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
		log.Printf("Failed to get all contract ids, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err)
		return err
	}
	contracts := event.Data.(msgbus.IDIndex)
	for i := range contracts {
		event, err := cs.ps.GetWait(msgbus.ContractMsg, contracts[i])
		if err != nil {
			log.Printf("Failed to get contract, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err)
			return err
		}
		contract := event.Data.(msgbus.Contract)
		cs.Contracts.Set(msgbus.ContractID(contracts[i]), contract)
	}

	// Monitor New Contracts
	contractEventChan := msgbus.NewEventChan()
	_, err = cs.ps.Sub(msgbus.ContractMsg, "", contractEventChan)
	if err != nil {
		log.Printf("Failed to subscribe to contract events, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err)
		return err
	}
	go cs.goContractHandler(contractEventChan)

	// Update connection scheduler with current miners in online state
	miners, err := cs.ps.MinerGetAllWait()
	if err != nil {
		log.Printf("Failed to get all miner ids, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err)
		return err
	}
	for i := range miners {
		miner, err := cs.ps.MinerGetWait(miners[i])
		if err != nil {
			log.Printf("Failed to get miner, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err)
			return err
		}
		if miner.State == msgbus.OnlineState {
			if miner.Dest == cs.nodeOperator.DefaultDest {
				cs.ReadyMiners.Set(msgbus.MinerID(miners[i]), *miner)
			} else {
				cs.BusyMiners.Set(msgbus.MinerID(miners[i]), *miner)
			}
		}
	}

	// Monitor New OnlineMiners
	minerEventChan := msgbus.NewEventChan()
	_, err = cs.ps.Sub(msgbus.MinerMsg, "", minerEventChan)
	if err != nil {
		log.Printf("Failed to subscribe to miner events, Fileline::%s, Error::%v\n", lumerinlib.FileLine(), err)
		return err
	}
	minerMux := &sync.Mutex{}
	go cs.goMinerHandler(minerEventChan, minerMux)

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
	for {
		select {
		case <-cs.ctx.Done():
			fmt.Println("Cancelling current connection scheduler context: cancelling ContractHandler go routine")
			return

		case event := <-ch:
			id := msgbus.ContractID(event.ID)

			switch event.EventType {

			//
			// Publish Event
			//
			case msgbus.PublishEvent:
				fmt.Printf(lumerinlib.Funcname()+" PublishEvent: %v\n", event)
				contract := event.Data.(msgbus.Contract)

				if !cs.Contracts.Exists(id) {
					cs.Contracts.Set(id, contract)
				} else {
					panic(fmt.Sprintf(lumerinlib.FileLine()+" got PubEvent, but already had the ID: %v\n", event))
				}

				// pusblished contract is already running
				if contract.State == msgbus.ContRunningState {
					cs.minerUpdatedChans.Set(id, make(chan bool, 5))
					cs.contractClosedChans.Set(id, make(chan bool, 5))
					go cs.ContractRunning(id)
				}

				//
				// Delete/Unsubscribe Event
				//
			case msgbus.DeleteEvent:
				fallthrough
			case msgbus.UnsubscribedEvent:
				fmt.Printf(lumerinlib.Funcname()+" Contract Delete/Unsubscribe Event: %v\n", event)

				if cs.Contracts.Exists(id) {
					cs.Contracts.Delete(id)

					cs.contractClosedChans.Get(id) <- true
					close(cs.minerUpdatedChans.Get(id))
					close(cs.contractClosedChans.Get(id))

					cs.minerUpdatedChans.Delete(id)
					cs.contractClosedChans.Delete(id)
				} else {
					panic(fmt.Sprintf(lumerinlib.FileLine()+" got UnsubscribeEvent, but dont have the ID: %v\n", event))
				}

				//
				// Update Event
				//
			case msgbus.UpdateEvent:
				fmt.Printf(lumerinlib.Funcname()+" UpdateEvent: %v\n", event)

				if !cs.Contracts.Exists(id) {
					panic(fmt.Sprintf(lumerinlib.FileLine()+" got contract ID does not exist: %v\n", event))
				}

				// Update the current contract data
				currentContract := cs.Contracts.Get(id)
				cs.Contracts.Set(id, event.Data.(msgbus.Contract))

				if currentContract.State == event.Data.(msgbus.Contract).State {
					fmt.Printf(lumerinlib.FileLine()+" got contract change with no state change: %v\n", event)
				} else {
					switch event.Data.(msgbus.Contract).State {
					case msgbus.ContAvailableState:
						fmt.Printf(lumerinlib.FileLine()+" Found Available Contract: %v\n", event)
						if currentContract.State != msgbus.ContAvailableState {
							cs.contractClosedChans.Get(id) <- true
							close(cs.minerUpdatedChans.Get(id))
							close(cs.contractClosedChans.Get(id))

							cs.minerUpdatedChans.Delete(id)
							cs.contractClosedChans.Delete(id)
						}

					case msgbus.ContRunningState:
						fmt.Printf(lumerinlib.FileLine()+" Found Running Contract: %v\n", event)

						if currentContract.State != msgbus.ContRunningState {
							cs.minerUpdatedChans.Set(id, make(chan bool, 5))
							cs.contractClosedChans.Set(id, make(chan bool, 5))
							go cs.ContractRunning(id)
						}

					default:
						panic(fmt.Sprintf(lumerinlib.FileLine()+" got bad State: %v\n", event))
					}
				}
			}
		}
	}
}

//------------------------------------------------------------------------
//
// Monitors new miner publish events, and then subscribes to the miners
// Then monitors the update events on the miners, and keeps track of their
// hashrate
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) goMinerHandler(ch msgbus.EventChan, mux *sync.Mutex) {
	for {
		select {
		case <-cs.ctx.Done():
			fmt.Println("Cancelling current connection scheduler context: cancelling MinerHandler go routine")
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

				switch miner.Contract {
				case "": // no contract
					if !cs.ReadyMiners.Exists(id) {
						cs.ReadyMiners.Set(id, miner)
						contracts := cs.Contracts.GetAll()
						for _, v := range contracts {
							if v.State == msgbus.ContRunningState {
								cs.minerUpdatedChans.Get(v.ID) <- true
							}
						}
					} else {
						panic(fmt.Sprintf("Got PubEvent, but already had the ID: %v\n", event))
					}
				default:
					if !cs.BusyMiners.Exists(id) {
						cs.BusyMiners.Set(id, miner)
						contracts := cs.Contracts.GetAll()
						for _, v := range contracts {
							if v.State == msgbus.ContRunningState {
								cs.minerUpdatedChans.Get(v.ID) <- true
							}
						}
					} else {
						panic(fmt.Sprintf("Got PubEvent, but already had the ID: %v\n", event))
					}
				}
				fmt.Println("Ready Miners: ", cs.ReadyMiners.m)
				fmt.Println("Busy Miners: ", cs.BusyMiners.m)

				//
				// Update Event
				//
			case msgbus.UpdateEvent:
				miner := event.Data.(msgbus.Miner)

				if miner.State != msgbus.OnlineState {
					cs.BusyMiners.Delete(id)
					cs.ReadyMiners.Delete(id)
					break loop
				}

				fmt.Printf("Miner Update Event:%v\n", event)

				switch miner.Contract {
				case "": // no contract
					// Update the current miner data
					cs.BusyMiners.Delete(id)
					cs.ReadyMiners.Set(id, miner)
					contracts := cs.Contracts.GetAll()
					for _, v := range contracts {
						if v.State == msgbus.ContRunningState && !miner.CsMinerHandlerIgnore {
							cs.minerUpdatedChans.Get(v.ID) <- true
						}
					}
				default:
					// Update the current miner data
					cs.ReadyMiners.Delete(id)
					cs.BusyMiners.Set(id, miner)
					contracts := cs.Contracts.GetAll()
					for _, v := range contracts {
						if v.State == msgbus.ContRunningState && !miner.CsMinerHandlerIgnore {
							cs.minerUpdatedChans.Get(v.ID) <- true
						}
					}
				}
				fmt.Println("Ready Miners: ", cs.ReadyMiners.m)
				fmt.Println("Busy Miners: ", cs.BusyMiners.m)

				//
				// Unpublish Event
				//
			case msgbus.UnpublishEvent:
				fmt.Printf("Miner Unpublish/Unsubscribe Event:%v\n", event)
				cs.BusyMiners.Delete(id)
				cs.ReadyMiners.Delete(id)
				contracts := cs.Contracts.GetAll()
				for _, v := range contracts {
					if v.State == msgbus.ContRunningState {
						cs.minerUpdatedChans.Get(v.ID) <- true
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

	event, err := cs.ps.GetWait(msgbus.ContractMsg, msgbus.IDString(contractId))
	if err != nil {
		panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", event))
	}
	contract := event.Data.(msgbus.Contract)

	availableHashrate, _ := cs.calculateHashrateAvailability(contractId)

	MIN := int(float32(contract.Speed) - float32(contract.Speed)*HASHRATE_TOLERANCE)

	if availableHashrate >= MIN {
		cs.SetMinerTarget(contract)
	} else {
		fmt.Println("Not enough available hashrate to fulfill contract: ", contract.ID)
		// free up busy miners with this contract id
		miners := cs.BusyMiners.GetAll()
		for _, v := range miners {
			if v.Contract == contract.ID {
				err := cs.ps.MinerRemoveContractWait(v.ID, cs.nodeOperator.DefaultDest, true)
				if err != nil {
					panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
				}
			}
		}
	}

	minerMapUpdated := cs.minerUpdatedChans.Get(contractId)
	contractClosed := cs.contractClosedChans.Get(contractId)
	for {
		select {
		case <-cs.ctx.Done():
			fmt.Println("Cancelling current connection scheduler context: cancelling contract running go routine for contract: ", contract.ID)
			return

		case <-contractClosed:
			fmt.Println("Contract state switched to available: cancelling contract running go routine for contract: ", contract.ID)

			// free up busy miners with this contract id
			miners := cs.BusyMiners.GetAll()
			for _, v := range miners {
				if v.Contract == contract.ID {
					err := cs.ps.MinerRemoveContractWait(v.ID, cs.nodeOperator.DefaultDest, true)
					if err != nil {
						panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
					}
				}
			}
			return

		case <-minerMapUpdated:
			//availableHashrate, contractHashrate := cs.calculateHashrateAvailability(contractId)
			availableHashrate, _ := cs.calculateHashrateAvailability(contractId)

			if availableHashrate >= MIN {
				//if contractHashrate < MIN {
				cs.SetMinerTarget(contract)
				//}
			} else {
				fmt.Println("Not enough available hashrate to fulfill contract: ", contract.ID)
				// free up busy miners with this contract id
				miners := cs.BusyMiners.GetAll()
				for _, v := range miners {
					if v.Contract == contract.ID {
						err := cs.ps.MinerRemoveContractWait(v.ID, cs.nodeOperator.DefaultDest, true)
						if err != nil {
							panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
						}
					}
				}
			}
		}
	}
}

func (cs *ConnectionScheduler) SetMinerTarget(contract msgbus.Contract) {
	destid := contract.Dest
	promisedHashrate := contract.Speed

	// in buyer node point miner directly to the pool
	if cs.nodeOperator.IsBuyer {
		destid = cs.nodeOperator.DefaultDest
	}

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
		miners := cs.BusyMiners.GetAll()
		for _, v := range miners {
			if v.Contract == contract.ID {
				err := cs.ps.MinerRemoveContractWait(v.ID, cs.nodeOperator.DefaultDest, true)
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

	// set contract and target destination for miners in optimal miner combination
	for _, v := range minerCombination {
		err := cs.ps.MinerSetContractWait(v.id, contract.ID, destid, true)
		if err != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
		}
	}

	// update busy miners map with new dests based on new miner combination i.e this function was called after an update to the miner map
	newBusyMinerMap := make(map[msgbus.MinerID]Miner)
	for _, v := range minerCombination {
		newBusyMinerMap[v.id] = v
	}
	miners := cs.BusyMiners.GetAll()
	for _, v := range miners {
		if _, ok := newBusyMinerMap[v.ID]; !ok {
			if v.Contract == contract.ID {
				err := cs.ps.MinerRemoveContractWait(v.ID, cs.nodeOperator.DefaultDest, true)
				if err != nil {
					panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
				}
			}
		}
	}
}

func (cs *ConnectionScheduler) calculateHashrateAvailability(id msgbus.ContractID) (availableHashrate int, contractHashrate int) {
	miners := cs.ReadyMiners.GetAll()
	for _, v := range miners {
		availableHashrate += v.CurrentHashRate
	}
	miners = cs.BusyMiners.GetAll()
	for _, v := range miners {
		if v.Contract == id {
			contractHashrate += v.CurrentHashRate
		}
	}
	availableHashrate += contractHashrate

	fmt.Println("Available Hashrate", availableHashrate)
	return availableHashrate, contractHashrate
}

func (cs *ConnectionScheduler) sortMinersByHashrate(contractId msgbus.ContractID) (m MinerList) {
	m = make(MinerList, 0)

	miners := cs.ReadyMiners.GetAll()
	for _, v := range miners {
		m = append(m, Miner{v.ID, v.CurrentHashRate})
	}

	// include busy miners that are already associated with contract
	miners = cs.BusyMiners.GetAll()
	for _, v := range miners {
		if v.Contract == contractId {
			m = append(m, Miner{v.ID, v.CurrentHashRate})
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
		n = n / 2
		j--
	}

	sum := 0

	// Calculate the sum of this subset
	for i := range sortedMiners {
		if x[i] == 1 {
			sum += sortedMiners[i].hashrate
		}
	}

	MIN := int(float32(targetHashrate) * (1 - HASHRATE_TOLERANCE))
	MAX := int(float32(targetHashrate) * (1 + HASHRATE_TOLERANCE))

	// if sum is within target hashrate bounds, subset was found
	if sum >= MIN && sum <= MAX {
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
func findSubsets(sortedMiners MinerList, targetHashrate int) (minerCombinations []MinerList) {
	// Calculate total number of subsets
	tot := math.Pow(2, float64(sortedMiners.Len()))

	for i := 0; i < int(tot); i++ {
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
