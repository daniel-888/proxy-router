package connectionscheduler

import (
	"context"
	"math"
	"sort"
	"sync"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/log"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/interfaces"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

const (
	HASHRATE_LIMIT = 20
	MIN_SLICE      = 0.10
)

type ConnectionScheduler struct {
	Ps                   *msgbus.PubSub
	Contracts            lumerinlib.ConcurrentMap
	ReadyMiners          lumerinlib.ConcurrentMap // miners with no contract
	BusyMiners           lumerinlib.ConcurrentMap // miners fulfilling a contract
	RunningContracts     []msgbus.ContractID
	ServiceContractChan  chan msgbus.ContractID
	wg                   sync.WaitGroup
	NodeOperator         msgbus.NodeOperator
	Passthrough          bool
	HashrateCalcLagTime  int
	Ctx                  context.Context
	connectionController interfaces.IConnectionController
}

// implement golang sort interface
type Miner struct {
	id           msgbus.MinerID
	hashrate     int
	slicePercent float64
}
type MinerList []Miner

func (m MinerList) Len() int           { return len(m) }
func (m MinerList) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MinerList) Less(i, j int) bool { return m[i].hashrate < m[j].hashrate }

func New(Ctx *context.Context, NodeOperator *msgbus.NodeOperator, Passthrough bool, HashrateCalcLagTime int, minerController interfaces.IConnectionController) (cs *ConnectionScheduler, err error) {
	ctxStruct := contextlib.GetContextStruct(*Ctx)
	cs = &ConnectionScheduler{
		Ps:                  ctxStruct.MsgBus,
		NodeOperator:        *NodeOperator,
		Passthrough:         Passthrough,
		HashrateCalcLagTime: HashrateCalcLagTime,
		Ctx:                 *Ctx,
	}
	cs.Contracts.M = make(map[string]interface{})
	cs.ReadyMiners.M = make(map[string]interface{})
	cs.BusyMiners.M = make(map[string]interface{})
	cs.ServiceContractChan = make(chan msgbus.ContractID, 5)
	cs.connectionController = minerController
	return cs, err
}

func (cs *ConnectionScheduler) Start() (err error) {
	contextlib.Logf(cs.Ctx, log.LevelInfo, "Connection Scheduler Starting")

	// Update connection scheduler with current contracts
	event, err := cs.Ps.GetWait(msgbus.ContractMsg, "")
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelError, "Failed to get all contract ids, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		return err
	}
	contracts := event.Data.(msgbus.IDIndex)
	for i := range contracts {
		event, err := cs.Ps.GetWait(msgbus.ContractMsg, contracts[i])
		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelError, "Failed to get contract, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
			return err
		}
		contract := event.Data.(msgbus.Contract)
		cs.Contracts.Set(string(contracts[i]), contract)
	}

	// Monitor Contract Events
	contractEventChan := msgbus.NewEventChan()
	_, err = cs.Ps.Sub(msgbus.ContractMsg, "", contractEventChan)
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelError, "Failed to subscribe to contract events, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		return err
	}
	go cs.ContractHandler(contractEventChan)

	// Start Contract Running Routine Manager
	go cs.RunningContractsManager()

	// Update connection scheduler with current miners in online state
	miners, err := cs.Ps.MinerGetAllWait()
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelError, "Failed to get all miner ids, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		return err
	}
	for i := range miners {
		miner, err := cs.Ps.MinerGetWait(miners[i])
		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelError, "Failed to get miner, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
			return err
		}

		dest, err := cs.Ps.DestGetWait(miner.Dest)

		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelError, "Failed to get miner destination, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		}

		contextlib.Logf(cs.Ctx, log.LevelInfo, "Adding connection... Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		connection, err := cs.connectionController.AddConnection(miner.IP, string(dest.NetUrl), string(miner.State), string(miner.ID))

		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelError, "Failed to add miner connection, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		}

		if miner.State == msgbus.OnlineState {

			if miner.Dest == cs.NodeOperator.DefaultDest {
				connection.SetAvailable(true)
				cs.ReadyMiners.Set(string(miners[i]), *miner)
			} else {
				connection.SetAvailable(false)
				cs.BusyMiners.Set(string(miners[i]), *miner)
			}
		}
	}

	// Monitor Miner Events
	minerEventChan := msgbus.NewEventChan()
	_, err = cs.Ps.Sub(msgbus.MinerMsg, "", minerEventChan)
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelError, "Failed to subscribe to miner events, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		return err
	}
	go cs.WatchMinerEvents(minerEventChan)

	return err
}

//------------------------------------------------------------------------
//
// Monitors new contract publish events, and then subscribes to the contracts
// Then monitors the update events on the contracts, and handles state changes
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) ContractHandler(ch msgbus.EventChan) {
	for {
		select {
		case <-cs.Ctx.Done():
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Cancelling current connection scheduler context: cancelling ContractHandler go routine")
			return

		case event := <-ch:
			id := msgbus.ContractID(event.ID)

			switch event.EventType {

			//
			// Publish Event
			//
			case msgbus.PublishEvent:
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Contract Publish Event: %v", event)
				contract := event.Data.(msgbus.Contract)

				if !cs.Contracts.Exists(string(id)) {
					cs.Contracts.Set(string(id), contract)
				} else {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.Funcname()+"Got Publish Event, but already had the ID: %v", event)
				}

				// pusblished contract is already running
				if contract.State == msgbus.ContRunningState {
					time.Sleep(time.Duration(cs.HashrateCalcLagTime) * time.Second)
					cs.ServiceContractChan <- id
				}

				//
				// Unpublish Event
				//
			case msgbus.UnpublishEvent:
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Contract Unpublish Event: %v", event)

				if cs.Contracts.Exists(string(id)) {
					cs.Contracts.Delete(string(id))
				} else {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.Funcname()+"Got Unsubscribe Event, but dont have the ID: %v", event)
				}

				//
				// Update Event
				//
			case msgbus.UpdateEvent:
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Contract Update Event: %v", event)

				if !cs.Contracts.Exists(string(id)) {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.Funcname()+"Got contract ID does not exist: %v", event)
				}

				// Update the current contract data
				currentContract := cs.Contracts.Get(string(id)).(msgbus.Contract)
				cs.Contracts.Set(string(id), event.Data.(msgbus.Contract))

				if currentContract.State == event.Data.(msgbus.Contract).State {
					contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Contract change with no state change: %v", event)
				} else {
					switch event.Data.(msgbus.Contract).State {
					case msgbus.ContAvailableState:
						contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Found Available Contract: %v", event)

					case msgbus.ContRunningState:
						contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Found Running Contract: %v", event)
						if currentContract.State != msgbus.ContRunningState {
							cs.ServiceContractChan <- id
						}

					default:
						contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.Funcname()+"Got bad state: %v", event)
					}
				}
			default:
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Contract Event: %v", event)
			}
		}
	}
}

func (cs *ConnectionScheduler) WatchMinerEvents(ch msgbus.EventChan) {
	for {
		select {
		case <-cs.Ctx.Done():
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Cancelling current connection scheduler context: cancelling RunningContractsManager go routine")
			return

		case event := <-ch:
			id := msgbus.MinerID(event.ID)

			var miner msgbus.Miner

			switch event.Data.(type) {
			case msgbus.Miner:
				miner = event.Data.(msgbus.Miner)
			case *msgbus.Miner:
				m := event.Data.(*msgbus.Miner)
				miner = *m
			}

			dest, err := cs.Ps.DestGetWait(miner.Dest)

			if err != nil {
				if miner.Dest == "" {
					dest = &msgbus.Dest{NetUrl: ""}
				}

				contextlib.Logf(cs.Ctx, log.LevelWarn, "Failed to get miner destination, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
			}

			contextlib.Logf(cs.Ctx, log.LevelInfo, "Getting or adding connection... Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
			connection, err := cs.connectionController.GetOrAddConnection(miner.IP, string(dest.NetUrl), string(miner.State), string(miner.ID))

			if err != nil {
				contextlib.Logf(cs.Ctx, log.LevelError, "Failed to get or add connection, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
			}

		loop:
			switch event.EventType {
			//
			// Publish Event
			//
			case msgbus.PublishEvent:

				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Publish Event: %v;  State: %v", event, miner.State)

				if miner.State != msgbus.OnlineState {
					contextlib.Logf(cs.Ctx, log.LevelInfo, "Removing connection... Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
					err := cs.connectionController.RemoveConnection(connection.GetId())

					if err != nil {
						contextlib.Logf(cs.Ctx, log.LevelError, "Failed to remove connection, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
					}

					break loop
				}

				// contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Publish Event: %v;  State: %v", event, miner.State)

				switch len(miner.Contracts) {
				case 0: // no contract
					if !cs.ReadyMiners.Exists(string(id)) {
						connection.SetAvailable(true)
						cs.ReadyMiners.Set(string(id), miner)
					} else {
						contextlib.Logf(cs.Ctx, log.LevelPanic, "Got PubEvent, but already had the ID: %v", event)
					}
				default:
					if !cs.BusyMiners.Exists(string(id)) {
						connection.SetAvailable(false)
					}
				}
				readyMiners := cs.ReadyMiners.GetAll()
				busyMiners := cs.BusyMiners.GetAll()
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners: %v", readyMiners)
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners: %v", busyMiners)

				//
				// Update Event
				//
			case msgbus.UpdateEvent:

				if miner.State != msgbus.OnlineState {
					cs.connectionController.RemoveConnection(string(id))
					cs.ReadyMiners.Delete(string(id))
					cs.BusyMiners.Delete(string(id))
					break loop
				}

				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Update Event: %v", event)

				switch len(miner.Contracts) {
				case 0: // no contract
					// Update the current miner data
					connection.SetAvailable(true)
					//cs.ReadyMiners.Set(string(id), miner)
				default:
					connection.SetAvailable(false)
					//cs.BusyMiners.Set(string(id), miner)
				}
				readyMiners := cs.ReadyMiners.GetAll()
				busyMiners := cs.BusyMiners.GetAll()
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners: %v", readyMiners)
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners: %v", busyMiners)

				//
				// Unpublish Event
				//
			case msgbus.UnpublishEvent:
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Unpublish/Unsubscribe Event: %v", event)
				cs.connectionController.RemoveConnection(connection.GetId())
				cs.ReadyMiners.Delete(string(id))
				cs.BusyMiners.Delete(string(id))

			default:
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Event: %v", event)
			}
		}
	}
}

//------------------------------------------------------------------------
//
//
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) RunningContractsManager() {
	for {
		select {
		case <-cs.Ctx.Done():
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Cancelling current connection scheduler context: cancelling RunningContractsManager go routine")
			return
		case <-cs.ServiceContractChan:
			contracts := cs.Contracts.GetAll()

			// Fill up ready and busy miners map
			miners, err := cs.Ps.MinerGetAllWait()
			if err != nil {
				contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
			}
			for i := range miners {
				miner, err := cs.Ps.MinerGetWait(miners[i])
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
				if miner.State == msgbus.OnlineState {
					if len(miner.Contracts) == 0 {
						cs.ReadyMiners.Set(string(miners[i]), *miner)
					} else {
						cs.BusyMiners.Set(string(miners[i]), *miner)
					}
				}
			}
			readyMiners := cs.ReadyMiners.GetAll()
			busyMiners := cs.BusyMiners.GetAll()
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners In Routine Manager: %v", readyMiners)
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners In Routine Manager: %v", busyMiners)

			for _, c := range contracts {
				if c.(msgbus.Contract).State == msgbus.ContRunningState {
					if cs.Passthrough {
						cs.wg.Add(1)
						go cs.ContractRunningPassthrough(c.(msgbus.Contract).ID)
					} else {
						cs.wg.Add(1)
						cs.RunningContracts = append(cs.RunningContracts, c.(msgbus.Contract).ID)
						go cs.ContractRunning(c.(msgbus.Contract).ID)
					}
				} else {
					// if in available state make sure no miners are servicing it
					miners := cs.Ps.MinersContainContract(c.(msgbus.Contract).ID)
					for _, v := range miners {
						cs.BusyMiners.Delete(string(v.ID))
						m, err := cs.Ps.MinerRemoveContractWait(v.ID, c.(msgbus.Contract).ID, cs.NodeOperator.DefaultDest)
						if err != nil {
							contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
						}
						newRunningContracts := []msgbus.ContractID{}
						for _, r := range cs.RunningContracts {
							if r != c.(msgbus.Contract).ID {
								newRunningContracts = append(newRunningContracts, r)
							}
						}
						cs.ReadyMiners.Set(string(m.ID), *m)
					}
				}
			}
			cs.wg.Wait()
		}
	}
}

//------------------------------------------------------------------------
//
// Direct all miners to running contract i.e. no miner hashrate calculations to service multiple contracts
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) ContractRunningPassthrough(contractId msgbus.ContractID) {
	defer cs.wg.Done()
	contextlib.Logf(cs.Ctx, log.LevelInfo, lumerinlib.FileLine()+"Contract Running in Passthrough Mode: %s", contractId)

	event, err := cs.Ps.GetWait(msgbus.ContractMsg, msgbus.IDString(contractId))
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", event)
	}
	contract := event.Data.(msgbus.Contract)
	destid := contract.Dest

	if destid == "" {
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"DestID is empty for Contract: %s", contractId)
	}

	// Find all of the online miners point them to new target
	miners := cs.ReadyMiners.GetAll()

	for _, m := range miners {
		_, err = cs.Ps.MinerSetContractWait(m.(msgbus.Miner).ID, contract.ID, 1, false)
		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelWarn, lumerinlib.FileLine()+"Error:%v", err)
		}
		miner, err := cs.Ps.MinerSetDestWait(m.(msgbus.Miner).ID, destid)
		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelWarn, lumerinlib.FileLine()+"Error:%v", err)
		}
		cs.ReadyMiners.Delete(string(miner.ID))
		cs.BusyMiners.Set(string(miner.ID), *miner)
	}

	time.Sleep(time.Second * time.Duration(cs.HashrateCalcLagTime))
	cs.ServiceContractChan <- contractId
}

//------------------------------------------------------------------------
//
// Search for optimal miner combination of online miners to point to running contract
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) ContractRunning(contractId msgbus.ContractID) {
	defer cs.wg.Done()
	contextlib.Logf(cs.Ctx, log.LevelInfo, lumerinlib.FileLine()+"Contract Running, ID: %s", contractId)

	event, err := cs.Ps.GetWait(msgbus.ContractMsg, msgbus.IDString(contractId))
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
	}
	contract := event.Data.(msgbus.Contract)

	//hashrateTolerance := float64(contract.Limit) / 100
	hashrateTolerance := float64(HASHRATE_LIMIT) / 100

	availableHashrate, _ := cs.calculateHashrateAvailability(contractId)

	MIN := int(float64(contract.Speed) - float64(contract.Speed)*hashrateTolerance)

	if availableHashrate >= MIN {
		cs.SetMinerTarget(contract)
	} else {
		contextlib.Logf(cs.Ctx, log.LevelWarn, "Not enough available hashrate to fulfill contract: %v", contract.ID)

		// free up busy miners with this contract id
		miners := cs.BusyMiners.GetAll()
		for _, v := range miners {
			if _, ok := v.(msgbus.Miner).Contracts[contractId]; ok {
				cs.BusyMiners.Delete(string(v.(msgbus.Miner).ID))
				m, err := cs.Ps.MinerRemoveContractWait(v.(msgbus.Miner).ID, contractId, cs.NodeOperator.DefaultDest)
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
				cs.ReadyMiners.Set(string(m.ID), *m)
			}
		}
		return
	}

	cs.ServiceContractChan <- contractId
}

func (cs *ConnectionScheduler) SetMinerTarget(contract msgbus.Contract) {
	contextlib.Logf(cs.Ctx, log.LevelInfo, "Setting Miner Target for Contract: %s", contract.ID)

	destid := contract.Dest
	promisedHashrate := contract.Speed
	hashrateTolerance := float64(HASHRATE_LIMIT) / 100

	// in buyer node point miner directly to the pool
	if cs.NodeOperator.IsBuyer {
		destid = cs.NodeOperator.DefaultDest
	}

	if destid == "" {
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"DestID is empty for Contract: %s", contract.ID)
	}

	// sort miners by hashrate from least to greatest
	sortedReadyMiners := cs.sortMinersByHashrate(contract.ID)
	contextlib.Logf(cs.Ctx, log.LevelInfo, "Sorted Miners for Contract %s: %v", contract.ID, sortedReadyMiners)

	// find all miner combinations that add up to promised hashrate
	minerCombinations := findSubsets(sortedReadyMiners, promisedHashrate, hashrateTolerance)
	if len(minerCombinations) == 0 {
		contextlib.Logf(cs.Ctx, log.LevelInfo, "Hashrate Value from Contract %s too small to create Valid Miner Combination ", contract.ID)
		return
	}

	contextlib.Logf(cs.Ctx, log.LevelInfo, "Valid Miner Combinations for Contract %s: %v", contract.ID, minerCombinations)

	// find best combination of miners
	minerCombination := bestCombination(minerCombinations, promisedHashrate)

	contextlib.Logf(cs.Ctx, log.LevelInfo, "Best Miner Combination for Contract %s: %v", contract.ID, minerCombination)

	// set contract and target destination for miners in optimal miner combination
	slicedMiners := []msgbus.Miner{}
	for _, v := range minerCombination {
		miner, err := cs.Ps.MinerGetWait(v.id)
		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
		}
		slicePercent := float64(v.hashrate) / float64(miner.CurrentHashRate)

		if slicePercent < 1 {
			_, err = cs.Ps.MinerSetContractWait(v.id, contract.ID, slicePercent, true)
			if err != nil {
				contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
			}
			slicedMiners = append(slicedMiners, *miner)
		} else {
			_, err = cs.Ps.MinerSetContractWait(v.id, contract.ID, slicePercent, false)
			if err != nil {
				contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
			}
			miner, err = cs.Ps.MinerSetDestWait(v.id, destid)
			if err != nil {
				contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
			}
			cs.ReadyMiners.Delete(string(miner.ID))
			cs.BusyMiners.Set(string(miner.ID), *miner)
		}
	}

	totalDuration := time.Second * time.Duration(cs.HashrateCalcLagTime)
	if len(slicedMiners) == 0 {
		contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners In Contract %s Set Target Func: %v", contract.ID, cs.ReadyMiners.M)
		contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners In Contract %s Set Target Func: %v", contract.ID, cs.ReadyMiners.M)
		time.Sleep(totalDuration)
	} else {
		totalDuration := time.Second * time.Duration(cs.HashrateCalcLagTime)
		for i, m := range slicedMiners {
			var durationPassed time.Duration
			for i, v := range m.Contracts {
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Switching Sliced Miner %s Dest to service Contract: %s", m.ID, i)
				event, err := cs.Ps.GetWait(msgbus.ContractMsg, msgbus.IDString(i))
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
				contract = event.Data.(msgbus.Contract)
				miner, err := cs.Ps.MinerSetDestWait(m.ID, contract.Dest)
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
				slicedDuration := time.Second * time.Duration(int(float64(cs.HashrateCalcLagTime)*v))
				cs.ReadyMiners.Delete(string(miner.ID))
				cs.BusyMiners.Set(string(miner.ID), *miner)
				readyMiners := cs.ReadyMiners.GetAll()
				busyMiners := cs.BusyMiners.GetAll()
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners In Contract %s Set Target Func: %v", contract.ID, readyMiners)
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners In Contract %s Set Target Func: %v", contract.ID, busyMiners)
				time.Sleep(slicedDuration)
				durationPassed += slicedDuration
			}
			if (i == len(slicedMiners)-1) && (durationPassed < totalDuration) {
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Switching Sliced Miner Dest to Default Dest, Miner: %s", m.ID)
				miner, err := cs.Ps.MinerSetDestWait(m.ID, cs.NodeOperator.DefaultDest)
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
				cs.ReadyMiners.Delete(string(miner.ID))
				cs.BusyMiners.Set(string(miner.ID), *miner)
				readyMiners := cs.ReadyMiners.GetAll()
				busyMiners := cs.BusyMiners.GetAll()
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners In Contract %s Set Target Func: %v", contract.ID, readyMiners)
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners In Contract %s Set Target Func: %v", contract.ID, busyMiners)
				time.Sleep(totalDuration - durationPassed)
			}
		}
	}
}

func (cs *ConnectionScheduler) calculateHashrateAvailability(id msgbus.ContractID) (availableHashrate int, contractHashrate int) {
	miners := cs.ReadyMiners.GetAll()
	for _, v := range miners {
		availableHashrate += v.(msgbus.Miner).CurrentHashRate
	}
	miners = cs.BusyMiners.GetAll()
	for _, v := range miners {
		if _, ok := v.(msgbus.Miner).Contracts[id]; ok {
			if !v.(msgbus.Miner).TimeSlice {
				contractHashrate += v.(msgbus.Miner).CurrentHashRate
			} else {
				contractHashrate += int(float64(v.(msgbus.Miner).CurrentHashRate) * v.(msgbus.Miner).Contracts[id])
			}
		}
		if v.(msgbus.Miner).TimeSlice {
			leftoverSlicePercent := cs.Ps.MinerSlicedUtilization(v.(msgbus.Miner).ID)
			availableHashrate += int(float64(v.(msgbus.Miner).CurrentHashRate) * leftoverSlicePercent)
		}
	}
	availableHashrate += contractHashrate

	contextlib.Logf(cs.Ctx, log.LevelInfo, "Available Hashrate for Contract %s: %v", id, availableHashrate)

	return availableHashrate, contractHashrate
}

func (cs *ConnectionScheduler) sortMinersByHashrate(contractId msgbus.ContractID) (m MinerList) {
	m = make(MinerList, 0)

	miners := cs.ReadyMiners.GetAll()
	for _, v := range miners {
		m = append(m, Miner{v.(msgbus.Miner).ID, v.(msgbus.Miner).CurrentHashRate, 0})
	}

	// include busy miners that are already associated with contract and sliced miners with extra contract space
	miners = cs.BusyMiners.GetAll()
	for _, v := range miners {
		if _, ok := v.(msgbus.Miner).Contracts[contractId]; ok {
			m = append(m, Miner{v.(msgbus.Miner).ID, v.(msgbus.Miner).CurrentHashRate, 0})
		} else if v.(msgbus.Miner).TimeSlice && (cs.Ps.MinerSlicedUtilization(v.(msgbus.Miner).ID) > 0) {
			leftoverSlicePercent := cs.Ps.MinerSlicedUtilization(v.(msgbus.Miner).ID)
			m = append(m, Miner{v.(msgbus.Miner).ID, int(float64(v.(msgbus.Miner).CurrentHashRate) * leftoverSlicePercent), 0})
		}
	}

	sort.Sort(m)
	return m
}

func sumSubsets(sortedMiners MinerList, n int, targetHashrate int, hashrateTolerance float64) (m MinerList, sum int) {
	// Create new array with size equal to sorted miners array to create binary array as per n(decimal number)
	x := make([]int, sortedMiners.Len())
	j := sortedMiners.Len() - 1

	// Convert the array into binary array
	for n > 0 {
		x[j] = n % 2
		n = n / 2
		j--
	}

	sumPrev := 0 // only return subsets where hashrate overflow is caused by 1 miner

	// Calculate the sum of this subset
	for i := range sortedMiners {
		if x[i] == 1 {
			sum += sortedMiners[i].hashrate
		}
		if i == len(sortedMiners)-1 {
			sumPrev = sum - sortedMiners[i].hashrate
		}
	}

	MIN := int(float64(targetHashrate) * (1 - hashrateTolerance))

	// if sum is within target hashrate bounds, subset was found
	if sum >= MIN && sumPrev < MIN{
		for i := range sortedMiners {
			if x[i] == 1 {
				m = append(m, sortedMiners[i])
			}
		}
		return m, sum
	}

	return nil,0
}

// find subsets of list of miners whose hashrate sum equal the target hashrate
func findSubsets(sortedMiners MinerList, targetHashrate int, hashrateTolerance float64) (minerCombinations []MinerList) {
	// Calculate total number of subsets
	tot := math.Pow(2, float64(sortedMiners.Len()))
	MAX := int(float64(targetHashrate) * (1 + hashrateTolerance))
	minerCombinationsSums := []int{}

	for i := 0; i < int(tot); i++ {
		m,s := sumSubsets(sortedMiners, i, targetHashrate, hashrateTolerance)
		if m != nil {
			minerCombinations = append(minerCombinations, m)
			minerCombinationsSums = append(minerCombinationsSums, s)
		}
	}

	if len(minerCombinations) == 0 {
		return []MinerList{}
	}

	for i,m := range minerCombinations {
		if minerCombinationsSums[i] > MAX { // need to slice miner
			sumPrev := minerCombinationsSums[i] - m[len(m) - 1].hashrate
			unslicedHashrate := m[len(m) - 1].hashrate
			slicedHashrate := targetHashrate - sumPrev
			if float64(slicedHashrate)/float64(unslicedHashrate) < MIN_SLICE {
				m[len(m) - 1].hashrate = int(float64(m[len(m) - 1].hashrate)*MIN_SLICE)
			} else {
				m[len(m) - 1].hashrate = slicedHashrate
			}
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
			if j == len(miners)-1 {
				totalHashRate += int(float64(miners[j].hashrate) * miners[j].slicePercent)
			} else {
				totalHashRate += miners[j].hashrate
			}
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
