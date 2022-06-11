package connectionscheduler

import (
	"context"
	"math"
	"sort"
<<<<<<< HEAD

	//"strconv"
	"sync"

	"gitlab.com/TitanInd/lumerin/cmd/log"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
=======
	"sync"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/log"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/interfaces"
>>>>>>> pr-009
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

<<<<<<< HEAD
=======
const (
	HASHRATE_LIMIT = 10
	MIN_SLICE      = 0.10
)

>>>>>>> pr-009
type ConnectionScheduler struct {
	Ps                   *msgbus.PubSub
	Contracts            lumerinlib.ConcurrentMap
	ReadyMiners          lumerinlib.ConcurrentMap // miners with no contract
	BusyMiners           lumerinlib.ConcurrentMap // miners fulfilling a contract
<<<<<<< HEAD
	NodeOperator         msgbus.NodeOperator
	MinerUpdatedChans    lumerinlib.ConcurrentMap
	ContractUpdatedChans lumerinlib.ConcurrentMap
	Passthrough          bool
	Ctx                  context.Context
=======
	RunningContracts     []msgbus.ContractID
	ServiceContractChan  chan msgbus.ContractID
	wg                   sync.WaitGroup
	NodeOperator         msgbus.NodeOperator
	Passthrough          bool
	HashrateCalcLagTime  int
	Ctx                  context.Context
	connectionController interfaces.IConnectionController
>>>>>>> pr-009
}

// implement golang sort interface
type Miner struct {
<<<<<<< HEAD
	id       msgbus.MinerID
	hashrate int
=======
	id           msgbus.MinerID
	hashrate     int
	slicePercent float64
>>>>>>> pr-009
}
type MinerList []Miner

func (m MinerList) Len() int           { return len(m) }
func (m MinerList) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MinerList) Less(i, j int) bool { return m[i].hashrate < m[j].hashrate }

<<<<<<< HEAD
func New(Ctx *context.Context, NodeOperator *msgbus.NodeOperator, Passthrough bool) (cs *ConnectionScheduler, err error) {
	ctxStruct := contextlib.GetContextStruct(*Ctx)
	cs = &ConnectionScheduler{
		Ps:           ctxStruct.MsgBus,
		NodeOperator: *NodeOperator,
		Passthrough:  Passthrough,
		Ctx:          *Ctx,
=======
func New(Ctx *context.Context, NodeOperator *msgbus.NodeOperator, Passthrough bool, HashrateCalcLagTime int, minerController interfaces.IConnectionController) (cs *ConnectionScheduler, err error) {
	ctxStruct := contextlib.GetContextStruct(*Ctx)
	cs = &ConnectionScheduler{
		Ps:                  ctxStruct.MsgBus,
		NodeOperator:        *NodeOperator,
		Passthrough:         Passthrough,
		HashrateCalcLagTime: HashrateCalcLagTime,
		Ctx:                 *Ctx,
>>>>>>> pr-009
	}
	cs.Contracts.M = make(map[string]interface{})
	cs.ReadyMiners.M = make(map[string]interface{})
	cs.BusyMiners.M = make(map[string]interface{})
<<<<<<< HEAD
	cs.MinerUpdatedChans.M = make(map[string]interface{})
	cs.ContractUpdatedChans.M = make(map[string]interface{})
=======
	cs.ServiceContractChan = make(chan msgbus.ContractID, 5)
	cs.connectionController = minerController
>>>>>>> pr-009
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

<<<<<<< HEAD
	// Monitor New Contracts
=======
	// Monitor Contract Events
>>>>>>> pr-009
	contractEventChan := msgbus.NewEventChan()
	_, err = cs.Ps.Sub(msgbus.ContractMsg, "", contractEventChan)
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelError, "Failed to subscribe to contract events, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		return err
	}
<<<<<<< HEAD
	go cs.goContractHandler(contractEventChan)
=======
	go cs.ContractHandler(contractEventChan)

	// Start Contract Running Routine Manager
	go cs.RunningContractsManager()
>>>>>>> pr-009

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
<<<<<<< HEAD
		if miner.State == msgbus.OnlineState {
			if miner.Dest == cs.NodeOperator.DefaultDest {
				cs.ReadyMiners.Set(string(miners[i]), *miner)
			} else {
=======

		dest, err := cs.Ps.DestGetWait(miner.Dest)

		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelError, "Failed to get miner destination, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		}

		contextlib.Logf(cs.Ctx, log.LevelInfo, "Adding connection... Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		connection, err := cs.connectionController.AddConnection(string(miner.ID), miner.IP, string(dest.NetUrl), string(miner.State))

		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelError, "Failed to add miner connection, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		}

		if miner.State == msgbus.OnlineState {

			if miner.Dest == cs.NodeOperator.DefaultDest {
				connection.SetAvailable(true)
				cs.ReadyMiners.Set(string(miners[i]), *miner)
			} else {
				connection.SetAvailable(false)
>>>>>>> pr-009
				cs.BusyMiners.Set(string(miners[i]), *miner)
			}
		}
	}

<<<<<<< HEAD
	// Monitor New OnlineMiners
=======
	// Monitor Miner Events
>>>>>>> pr-009
	minerEventChan := msgbus.NewEventChan()
	_, err = cs.Ps.Sub(msgbus.MinerMsg, "", minerEventChan)
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelError, "Failed to subscribe to miner events, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		return err
	}
<<<<<<< HEAD
	minerMux := &sync.Mutex{}
	go cs.goMinerHandler(minerEventChan, minerMux)
=======
	go cs.WatchMinerEvents(minerEventChan)
>>>>>>> pr-009

	return err
}

//------------------------------------------------------------------------
//
// Monitors new contract publish events, and then subscribes to the contracts
// Then monitors the update events on the contracts, and handles state changes
//
//------------------------------------------------------------------------
<<<<<<< HEAD
func (cs *ConnectionScheduler) goContractHandler(ch msgbus.EventChan) {
=======
func (cs *ConnectionScheduler) ContractHandler(ch msgbus.EventChan) {
>>>>>>> pr-009
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
<<<<<<< HEAD
					cs.MinerUpdatedChans.Set(string(id), make(chan bool, 5))
					cs.ContractUpdatedChans.Set(string(id), make(chan bool, 5))
					if cs.Passthrough {
						go cs.ContractRunningPassthrough(id)
					} else {
						go cs.ContractRunning(id)
					}
=======
					time.Sleep(time.Duration(cs.HashrateCalcLagTime) * time.Second)
					cs.ServiceContractChan <- id
>>>>>>> pr-009
				}

				//
				// Unpublish Event
				//
			case msgbus.UnpublishEvent:
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Contract Unpublish Event: %v", event)

				if cs.Contracts.Exists(string(id)) {
					cs.Contracts.Delete(string(id))
<<<<<<< HEAD

					cs.ContractUpdatedChans.Get(string(id)).(chan bool) <- true
					close(cs.MinerUpdatedChans.Get(string(id)).(chan bool))
					close(cs.ContractUpdatedChans.Get(string(id)).(chan bool))

					cs.MinerUpdatedChans.Delete(string(id))
					cs.ContractUpdatedChans.Delete(string(id))
=======
>>>>>>> pr-009
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
<<<<<<< HEAD
					// cs.ContractUpdatedChans.Get(string(id)).(chan bool) <- true
=======
>>>>>>> pr-009
				} else {
					switch event.Data.(msgbus.Contract).State {
					case msgbus.ContAvailableState:
						contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Found Available Contract: %v", event)
<<<<<<< HEAD
						if currentContract.State != msgbus.ContAvailableState {
							cs.ContractUpdatedChans.Get(string(id)).(chan bool) <- true
							close(cs.MinerUpdatedChans.Get(string(id)).(chan bool))
							close(cs.ContractUpdatedChans.Get(string(id)).(chan bool))

							cs.MinerUpdatedChans.Delete(string(id))
							cs.ContractUpdatedChans.Delete(string(id))
						}
=======
>>>>>>> pr-009

					case msgbus.ContRunningState:
						contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Found Running Contract: %v", event)
						if currentContract.State != msgbus.ContRunningState {
<<<<<<< HEAD
							cs.MinerUpdatedChans.Set(string(id), make(chan bool, 5))
							cs.ContractUpdatedChans.Set(string(id), make(chan bool, 5))
							if cs.Passthrough {
								go cs.ContractRunningPassthrough(id)
							} else {
								go cs.ContractRunning(id)
							}
=======
							cs.ServiceContractChan <- id
>>>>>>> pr-009
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

<<<<<<< HEAD
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
		case <-cs.Ctx.Done():
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Cancelling current connection scheduler context: cancelling MinerHandler go routine")
=======
func (cs *ConnectionScheduler) WatchMinerEvents(ch msgbus.EventChan) {
	for {
		select {
		case <-cs.Ctx.Done():
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Cancelling current connection scheduler context: cancelling RunningContractsManager go routine")
>>>>>>> pr-009
			return

		case event := <-ch:
			id := msgbus.MinerID(event.ID)

<<<<<<< HEAD
=======
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Miner event: %+v", event)

			var miner msgbus.Miner

			switch event.Data.(type) {
			case msgbus.Miner:
				miner = event.Data.(msgbus.Miner)
			case *msgbus.Miner:
				m := event.Data.(*msgbus.Miner)
				miner = *m
			}

			contextlib.Logf(cs.Ctx, log.LevelInfo, "Miner event data: %+v", miner)

>>>>>>> pr-009
		loop:
			switch event.EventType {
			//
			// Publish Event
			//
			case msgbus.PublishEvent:
<<<<<<< HEAD
				var miner msgbus.Miner

				switch event.Data.(type) {
				case msgbus.Miner:
					miner = event.Data.(msgbus.Miner)
				case *msgbus.Miner:
					m := event.Data.(*msgbus.Miner)
					miner = *m
				}

				if miner.State != msgbus.OnlineState {
					break loop
				}

				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Publish Event: %v", event)

				switch miner.Contract {
				case "": // no contract
					if !cs.ReadyMiners.Exists(string(id)) {
						cs.ReadyMiners.Set(string(id), miner)
						contracts := cs.Contracts.GetAll()
						for _, v := range contracts {
							if v.(msgbus.Contract).State == msgbus.ContRunningState {
								cs.MinerUpdatedChans.Get(string(v.(msgbus.Contract).ID)).(chan bool) <- true
							}
						}
=======
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Publish Event: %v", event)

				dest, err := cs.Ps.DestGetWait(miner.Dest)

				if err != nil {
					if miner.Dest == "" {
						dest = &msgbus.Dest{NetUrl: ""}
					}

					contextlib.Logf(cs.Ctx, log.LevelWarn, "Failed to get miner destination, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
				}

				connection, err := cs.connectionController.AddConnection(string(miner.ID), miner.IP, string(dest.NetUrl), string(miner.State))
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelError, "Cannot add connection: %v. Fileline: %s", err, lumerinlib.FileLine())
				}

				switch len(miner.Contracts) {
				case 0: // no contract
					if !cs.ReadyMiners.Exists(string(id)) {
						connection.SetAvailable(true)
						cs.ReadyMiners.Set(string(id), miner)
>>>>>>> pr-009
					} else {
						contextlib.Logf(cs.Ctx, log.LevelPanic, "Got PubEvent, but already had the ID: %v", event)
					}
				default:
					if !cs.BusyMiners.Exists(string(id)) {
<<<<<<< HEAD
						cs.BusyMiners.Set(string(id), miner)
						contracts := cs.Contracts.GetAll()
						for _, v := range contracts {
							if v.(msgbus.Contract).State == msgbus.ContRunningState {
								cs.MinerUpdatedChans.Get(string(v.(msgbus.Contract).ID)).(chan bool) <- true
							}
						}
					} else {
						contextlib.Logf(cs.Ctx, log.LevelPanic, "Got PubEvent, but already had the ID: %v", event)
					}
				}
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners: %v", cs.ReadyMiners.M)
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners: %v", cs.BusyMiners.M)
=======
						connection.SetAvailable(false)
					}
				}
				readyMiners := cs.ReadyMiners.GetAll()
				busyMiners := cs.BusyMiners.GetAll()
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners: %v", readyMiners)
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners: %v", busyMiners)
>>>>>>> pr-009

				//
				// Update Event
				//
			case msgbus.UpdateEvent:
<<<<<<< HEAD
				var miner msgbus.Miner

				switch event.Data.(type) {
				case msgbus.Miner:
					miner = event.Data.(msgbus.Miner)
				case *msgbus.Miner:
					m := event.Data.(*msgbus.Miner)
					miner = *m
				}

				if miner.State != msgbus.OnlineState {
					cs.BusyMiners.Delete(string(id))
					cs.ReadyMiners.Delete(string(id))
					break loop
				}

				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Update Event: %v", event)

				switch miner.Contract {
				case "": // no contract
					// Update the current miner data
					cs.BusyMiners.Delete(string(id))
					cs.ReadyMiners.Set(string(id), miner)
					contracts := cs.Contracts.GetAll()
					for _, v := range contracts {
						if v.(msgbus.Contract).State == msgbus.ContRunningState && !miner.CsMinerHandlerIgnore {
							cs.MinerUpdatedChans.Get(string(v.(msgbus.Contract).ID)).(chan bool) <- true
						}
					}
				default:
					// Update the current miner data
					cs.ReadyMiners.Delete(string(id))
					cs.BusyMiners.Set(string(id), miner)
					contracts := cs.Contracts.GetAll()
					for _, v := range contracts {
						if v.(msgbus.Contract).State == msgbus.ContRunningState && !miner.CsMinerHandlerIgnore {
							cs.MinerUpdatedChans.Get(string(v.(msgbus.Contract).ID)).(chan bool) <- true
						}
					}
				}
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners: %v", cs.ReadyMiners.M)
				contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners: %v", cs.BusyMiners.M)
=======
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Update Event: %v", event)

				connection := cs.connectionController.GetConnection(string(id))

				if miner.State != msgbus.OnlineState {
					cs.connectionController.RemoveConnection(string(id))
					cs.ReadyMiners.Delete(string(id))
					cs.BusyMiners.Delete(string(id))
					break loop
				}

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
>>>>>>> pr-009

				//
				// Unpublish Event
				//
			case msgbus.UnpublishEvent:
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Unpublish/Unsubscribe Event: %v", event)
<<<<<<< HEAD
				cs.BusyMiners.Delete(string(id))
				cs.ReadyMiners.Delete(string(id))
				contracts := cs.Contracts.GetAll()
				for _, v := range contracts {
					if v.(msgbus.Contract).State == msgbus.ContRunningState {
						cs.MinerUpdatedChans.Get(string(v.(msgbus.Contract).ID)).(chan bool) <- true
					}
				}
=======
				cs.connectionController.RemoveConnection(string(id))
				cs.ReadyMiners.Delete(string(id))
				cs.BusyMiners.Delete(string(id))
>>>>>>> pr-009

			default:
				contextlib.Logf(cs.Ctx, log.LevelTrace, lumerinlib.Funcname()+"Got Miner Event: %v", event)
			}
		}
	}
}

//------------------------------------------------------------------------
//
<<<<<<< HEAD
// Direct all miners to running contract i.e. miner hashrate calculations to service multiple contracts
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) ContractRunningPassthrough(contractId msgbus.ContractID) {
=======
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
						cs.RunningContracts = newRunningContracts
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
>>>>>>> pr-009
	contextlib.Logf(cs.Ctx, log.LevelInfo, lumerinlib.FileLine()+"Contract Running in Passthrough Mode: %s", contractId)

	event, err := cs.Ps.GetWait(msgbus.ContractMsg, msgbus.IDString(contractId))
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", event)
	}
	contract := event.Data.(msgbus.Contract)
	destid := contract.Dest

	if destid == "" {
<<<<<<< HEAD
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"DestID is empty")
	}

	// Find all of the online miners point them to new target
	miners, err := cs.Ps.MinerGetAllWait()
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
	}

	for _, v := range miners {
		err := cs.Ps.MinerSetContractWait(v, contract.ID, destid, true)
		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
		}
	}

	minerMapUpdated := cs.MinerUpdatedChans.Get(string(contractId)).(chan bool)
	contractUpdated := cs.ContractUpdatedChans.Get(string(contractId)).(chan bool)
	for {
		select {
		case <-cs.Ctx.Done():
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Cancelling current connection scheduler context: cancelling contract running Passthrough go routine for contract: %v", contract.ID)
			return

		case <-contractUpdated:
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Contract state switched to available: cancelling contract running go routine for contract: %v", contract.ID)

			// free up busy miners with this contract id
			miners := cs.BusyMiners.GetAll()
			for _, v := range miners {
				if v.(msgbus.Miner).Contract == contract.ID {
					err := cs.Ps.MinerRemoveContractWait(v.(msgbus.Miner).ID, cs.NodeOperator.DefaultDest, true)
					if err != nil {
						contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
					}
				}
			}
			return

		case <-minerMapUpdated:
			miners, err := cs.Ps.MinerGetAllWait()
			if err != nil {
				contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
			}
			for _, v := range miners {
				miner, err := cs.Ps.MinerGetWait(v)
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
				if miner.Contract == "" {
					err = cs.Ps.MinerSetContractWait(miner.ID, contract.ID, destid, true)
					if err != nil {
						contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
					}
				}
			}
		}
	}
=======
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
>>>>>>> pr-009
}

//------------------------------------------------------------------------
//
<<<<<<< HEAD
// Continously search for optimal miner combination of online miners to point to running contract
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) ContractRunning(contractId msgbus.ContractID) {
	contextlib.Logf(cs.Ctx, log.LevelInfo, lumerinlib.FileLine()+"Contract Running: %s", contractId)
=======
// Search for optimal miner combination of online miners to point to running contract
//
//------------------------------------------------------------------------
func (cs *ConnectionScheduler) ContractRunning(contractId msgbus.ContractID) {
	defer cs.wg.Done()
	contextlib.Logf(cs.Ctx, log.LevelInfo, lumerinlib.FileLine()+"Contract Running, ID: %s", contractId)
>>>>>>> pr-009

	event, err := cs.Ps.GetWait(msgbus.ContractMsg, msgbus.IDString(contractId))
	if err != nil {
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
	}
	contract := event.Data.(msgbus.Contract)
<<<<<<< HEAD
	hashrateTolerance := float64(contract.Limit) / 100
=======

	hashrateTolerance := float64(HASHRATE_LIMIT) / 100
>>>>>>> pr-009

	availableHashrate, _ := cs.calculateHashrateAvailability(contractId)

	MIN := int(float64(contract.Speed) - float64(contract.Speed)*hashrateTolerance)

	if availableHashrate >= MIN {
		cs.SetMinerTarget(contract)
	} else {
		contextlib.Logf(cs.Ctx, log.LevelWarn, "Not enough available hashrate to fulfill contract: %v", contract.ID)

		// free up busy miners with this contract id
		miners := cs.BusyMiners.GetAll()
		for _, v := range miners {
<<<<<<< HEAD
			if v.(msgbus.Miner).Contract == contract.ID {
				err := cs.Ps.MinerRemoveContractWait(v.(msgbus.Miner).ID, cs.NodeOperator.DefaultDest, true)
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
			}
		}
	}

	minerMapUpdated := cs.MinerUpdatedChans.Get(string(contractId)).(chan bool)
	contractUpdated := cs.ContractUpdatedChans.Get(string(contractId)).(chan bool)
	for {
		select {
		case <-cs.Ctx.Done():
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Cancelling current connection scheduler context: cancelling contract running go routine for contract: %v", contract.ID)
			return

		case <-contractUpdated:
			contextlib.Logf(cs.Ctx, log.LevelInfo, "Contract state switched to available: cancelling contract running go routine for contract: %v", contract.ID)

			// free up busy miners with this contract id
			miners := cs.BusyMiners.GetAll()
			for _, v := range miners {
				if v.(msgbus.Miner).Contract == contract.ID {
					err := cs.Ps.MinerRemoveContractWait(v.(msgbus.Miner).ID, cs.NodeOperator.DefaultDest, true)
					if err != nil {
						contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
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
				contextlib.Logf(cs.Ctx, log.LevelWarn, "Not enough available hashrate to fulfill contract: %v", contract.ID)
				// free up busy miners with this contract id
				miners := cs.BusyMiners.GetAll()
				for _, v := range miners {
					if v.(msgbus.Miner).Contract == contract.ID {
						err := cs.Ps.MinerRemoveContractWait(v.(msgbus.Miner).ID, cs.NodeOperator.DefaultDest, true)
						if err != nil {
							contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
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
	hashrateTolerance := float64(contract.Limit) / 100
=======
			if _, ok := v.(msgbus.Miner).Contracts[contractId]; ok {
				m, err := cs.Ps.MinerRemoveContractWait(v.(msgbus.Miner).ID, contractId, cs.NodeOperator.DefaultDest)
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
				if len(m.Contracts) == 0 {
					cs.BusyMiners.Delete(string(m.ID))
					cs.ReadyMiners.Set(string(m.ID), *m)
				} else {
					cs.BusyMiners.Set(string(m.ID), *m)
				}
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
>>>>>>> pr-009

	// in buyer node point miner directly to the pool
	if cs.NodeOperator.IsBuyer {
		destid = cs.NodeOperator.DefaultDest
	}

	if destid == "" {
<<<<<<< HEAD
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+" Error DestID is empty")
=======
		contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"DestID is empty for Contract: %s", contract.ID)
>>>>>>> pr-009
	}

	// sort miners by hashrate from least to greatest
	sortedReadyMiners := cs.sortMinersByHashrate(contract.ID)
<<<<<<< HEAD
	contextlib.Logf(cs.Ctx, log.LevelInfo, "Sorted Miners: %v", sortedReadyMiners)

	// find all miner combinations that add up to promised hashrate
	minerCombinations := findSubsets(sortedReadyMiners, promisedHashrate, hashrateTolerance)
	if minerCombinations == nil {
		contextlib.Logf(cs.Ctx, log.LevelWarn, "No valid miner combinations")

		// free up busy miners with this contract id
		miners := cs.BusyMiners.GetAll()
		for _, v := range miners {
			if v.(msgbus.Miner).Contract == contract.ID {
				err := cs.Ps.MinerRemoveContractWait(v.(msgbus.Miner).ID, cs.NodeOperator.DefaultDest, true)
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
			}
		}
		return
	}

	contextlib.Logf(cs.Ctx, log.LevelInfo, "Valid Miner Combinations: %v", minerCombinations)

	// find best combination of miners
	minerCombination := bestCombination(minerCombinations, promisedHashrate)
	contextlib.Logf(cs.Ctx, log.LevelInfo, "Best Miner Combination: %v", minerCombination)

	// set contract and target destination for miners in optimal miner combination
	for _, v := range minerCombination {
		err := cs.Ps.MinerSetContractWait(v.id, contract.ID, destid, true)
		if err != nil {
			contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
		}
	}

	// update busy miners map with new dests based on new miner combination i.e this function was called after an update to the miner map
	newBusyMinerMap := make(map[msgbus.MinerID]Miner)
	for _, v := range minerCombination {
		newBusyMinerMap[v.id] = v
	}
	miners := cs.BusyMiners.GetAll()
	for _, v := range miners {
		if _, ok := newBusyMinerMap[v.(msgbus.Miner).ID]; !ok {
			if v.(msgbus.Miner).Contract == contract.ID {
				err := cs.Ps.MinerRemoveContractWait(v.(msgbus.Miner).ID, cs.NodeOperator.DefaultDest, true)
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
=======
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
	contractStateChanged := false
	var durationPassed time.Duration
	if len(slicedMiners) == 0 {
		currentReadyMiners := cs.ReadyMiners.GetAll()
		currentBusyMiners := cs.BusyMiners.GetAll()
		contextlib.Logf(cs.Ctx, log.LevelInfo, "Ready Miners In Contract %s Set Target Func: %v", contract.ID, currentReadyMiners)
		contextlib.Logf(cs.Ctx, log.LevelInfo, "Busy Miners In Contract %s Set Target Func: %v", contract.ID, currentBusyMiners)
	loop1:
		for i := 0; i < 5; i++ {
			// check none of the miners were unpublished
			for _, v := range currentBusyMiners {
				if !cs.BusyMiners.Exists(string(v.(msgbus.Miner).ID)) {
					break loop1
				}
			}

			// check if contract went to available periodically
			event, err := cs.Ps.GetWait(msgbus.ContractMsg, msgbus.IDString(contract.ID))
			if err != nil {
				contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
			}
			contract = event.Data.(msgbus.Contract)
			if contract.State == msgbus.ContAvailableState {
				contractStateChanged = true
				break loop1
			}
			durationPassed += totalDuration / 5
			time.Sleep(totalDuration / 5)
		}
	} else {
	loop2:
		for i, m := range slicedMiners {
			for i, v := range m.Contracts {
				// check none of the miners were unpublished
				for _, v := range slicedMiners {
					if !cs.BusyMiners.Exists(string(v.ID)) {
						break loop2
					}
				}

				// check if contract went to available
				event, err := cs.Ps.GetWait(msgbus.ContractMsg, msgbus.IDString(contract.ID))
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
				contract = event.Data.(msgbus.Contract)
				if contract.State == msgbus.ContAvailableState {
					contractStateChanged = true
					break loop2
				}

				contextlib.Logf(cs.Ctx, log.LevelInfo, "Switching Sliced Miner %s Dest to service Contract: %s", m.ID, i)
				event, err = cs.Ps.GetWait(msgbus.ContractMsg, msgbus.IDString(i))
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
	if contractStateChanged {
		miners := cs.BusyMiners.GetAll()
		for _, v := range miners {
			if _, ok := v.(msgbus.Miner).Contracts[contract.ID]; ok {
				m, err := cs.Ps.MinerRemoveContractWait(v.(msgbus.Miner).ID, contract.ID, cs.NodeOperator.DefaultDest)
				if err != nil {
					contextlib.Logf(cs.Ctx, log.LevelPanic, lumerinlib.FileLine()+"Error:%v", err)
				}
				if len(m.Contracts) == 0 {
					cs.BusyMiners.Delete(string(m.ID))
					cs.ReadyMiners.Set(string(m.ID), *m)
				} else {
					cs.BusyMiners.Set(string(m.ID), *m)
				}
>>>>>>> pr-009
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
<<<<<<< HEAD
		if v.(msgbus.Miner).Contract == id {
			contractHashrate += v.(msgbus.Miner).CurrentHashRate
=======
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
>>>>>>> pr-009
		}
	}
	availableHashrate += contractHashrate

<<<<<<< HEAD
	contextlib.Logf(cs.Ctx, log.LevelInfo, "Available Hashrate: %v", availableHashrate)
=======
	contextlib.Logf(cs.Ctx, log.LevelInfo, "Available Hashrate for Contract %s: %v", id, availableHashrate)
>>>>>>> pr-009

	return availableHashrate, contractHashrate
}

func (cs *ConnectionScheduler) sortMinersByHashrate(contractId msgbus.ContractID) (m MinerList) {
	m = make(MinerList, 0)

	miners := cs.ReadyMiners.GetAll()
	for _, v := range miners {
<<<<<<< HEAD
		m = append(m, Miner{v.(msgbus.Miner).ID, v.(msgbus.Miner).CurrentHashRate})
	}

	// include busy miners that are already associated with contract
	miners = cs.BusyMiners.GetAll()
	for _, v := range miners {
		if v.(msgbus.Miner).Contract == contractId {
			m = append(m, Miner{v.(msgbus.Miner).ID, v.(msgbus.Miner).CurrentHashRate})
=======
		m = append(m, Miner{v.(msgbus.Miner).ID, v.(msgbus.Miner).CurrentHashRate, 1})
	}

	// include busy miners that are already associated with contract and sliced miners with extra contract space
	miners = cs.BusyMiners.GetAll()
	for _, v := range miners {
		if _, ok := v.(msgbus.Miner).Contracts[contractId]; ok {
			slicePercent := v.(msgbus.Miner).Contracts[contractId]
			m = append(m, Miner{v.(msgbus.Miner).ID, int(float64(v.(msgbus.Miner).CurrentHashRate) * slicePercent), slicePercent})
		} else if v.(msgbus.Miner).TimeSlice && (cs.Ps.MinerSlicedUtilization(v.(msgbus.Miner).ID) > 0) {
			leftoverSlicePercent := cs.Ps.MinerSlicedUtilization(v.(msgbus.Miner).ID)
			m = append(m, Miner{v.(msgbus.Miner).ID, int(float64(v.(msgbus.Miner).CurrentHashRate) * leftoverSlicePercent), leftoverSlicePercent})
>>>>>>> pr-009
		}
	}

	sort.Sort(m)
	return m
}

<<<<<<< HEAD
func sumSubsets(sortedMiners MinerList, n int, targetHashrate int, hashrateTolerance float64) (m MinerList) {
=======
func sumSubsets(sortedMiners MinerList, n int, targetHashrate int, hashrateTolerance float64) (m MinerList, sum int) {
>>>>>>> pr-009
	// Create new array with size equal to sorted miners array to create binary array as per n(decimal number)
	x := make([]int, sortedMiners.Len())
	j := sortedMiners.Len() - 1

	// Convert the array into binary array
	for n > 0 {
		x[j] = n % 2
		n = n / 2
		j--
	}

<<<<<<< HEAD
	sum := 0
=======
	sumPrev := 0 // only return subsets where hashrate overflow is caused by 1 miner
>>>>>>> pr-009

	// Calculate the sum of this subset
	for i := range sortedMiners {
		if x[i] == 1 {
			sum += sortedMiners[i].hashrate
		}
<<<<<<< HEAD
	}

	MIN := int(float64(targetHashrate) * (1 - hashrateTolerance))
	MAX := int(float64(targetHashrate) * (1 + hashrateTolerance))

	// if sum is within target hashrate bounds, subset was found
	if sum >= MIN && sum <= MAX {
=======
		if i == len(sortedMiners)-1 {
			sumPrev = sum - sortedMiners[i].hashrate
		}
	}

	MIN := int(float64(targetHashrate) * (1 - hashrateTolerance))

	// if sum is within target hashrate bounds, subset was found
	if sum >= MIN && sumPrev < MIN {
>>>>>>> pr-009
		for i := range sortedMiners {
			if x[i] == 1 {
				m = append(m, sortedMiners[i])
			}
		}
<<<<<<< HEAD
		return m
	}

	return nil
=======
		return m, sum
	}

	return nil, 0
>>>>>>> pr-009
}

// find subsets of list of miners whose hashrate sum equal the target hashrate
func findSubsets(sortedMiners MinerList, targetHashrate int, hashrateTolerance float64) (minerCombinations []MinerList) {
	// Calculate total number of subsets
	tot := math.Pow(2, float64(sortedMiners.Len()))
<<<<<<< HEAD

	for i := 0; i < int(tot); i++ {
		m := sumSubsets(sortedMiners, i, targetHashrate, hashrateTolerance)
		if m != nil {
			minerCombinations = append(minerCombinations, m)
		}
	}
=======
	MAX := int(float64(targetHashrate) * (1 + hashrateTolerance))
	minerCombinationsSums := []int{}

	for i := 0; i < int(tot); i++ {
		m, s := sumSubsets(sortedMiners, i, targetHashrate, hashrateTolerance)
		if m != nil {
			minerCombinations = append(minerCombinations, m)
			minerCombinationsSums = append(minerCombinationsSums, s)
		}
	}

	if len(minerCombinations) == 0 {
		return []MinerList{}
	}

	for i, m := range minerCombinations {
		if minerCombinationsSums[i] > MAX { // need to slice miner
			sumPrev := minerCombinationsSums[i] - m[len(m)-1].hashrate
			unslicedHashrate := m[len(m)-1].hashrate
			slicedHashrate := targetHashrate - sumPrev
			if float64(slicedHashrate)/float64(unslicedHashrate) < MIN_SLICE {
				m[len(m)-1].hashrate = int(float64(m[len(m)-1].hashrate) * MIN_SLICE)
				m[len(m)-1].slicePercent = MIN_SLICE
			} else {
				m[len(m)-1].hashrate = slicedHashrate
				m[len(m)-1].slicePercent = float64(slicedHashrate) / float64(unslicedHashrate)
			}
		}
	}

>>>>>>> pr-009
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
<<<<<<< HEAD
			totalHashRate += miners[j].hashrate
			num++
		}
=======
			if j == len(miners)-1 {
				totalHashRate += int(float64(miners[j].hashrate) * miners[j].slicePercent)
			} else {
				totalHashRate += miners[j].hashrate
			}
			num++
		}

>>>>>>> pr-009
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
