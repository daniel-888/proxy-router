package testmod

import (
	"fmt"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func CreateContract(ps *msgbus.PubSub) {

	id := msgbus.GetRandomIDString()
	contract := msgbus.Contract{
		ID:    msgbus.ContractID(id),
		State: msgbus.ContAvailableState,
	}

	event, err := ps.PubWait(msgbus.ContractMsg, msgbus.IDString(id), contract)
	if err != nil {
		panic(fmt.Sprintf("Adding Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Contract Failed: %s", event.Err))
	}

	time.Sleep(1 * time.Second)

	contract.State = msgbus.ContAvailableState

	event, err = ps.SetWait(msgbus.ContractMsg, msgbus.IDString(id), contract)
	if err != nil {
		panic(fmt.Sprintf("Adding Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Contract Failed: %s", event.Err))
	}

	time.Sleep(1 * time.Second)

	contract.State = msgbus.ContRunningState

	event, err = ps.SetWait(msgbus.ContractMsg, msgbus.IDString(id), contract)
	if err != nil {
		panic(fmt.Sprintf("Adding Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Contract Failed: %s", event.Err))
	}

	time.Sleep(1 * time.Second)

	contract.State = msgbus.ContCompleteState

	event, err = ps.SetWait(msgbus.ContractMsg, msgbus.IDString(id), contract)
	if err != nil {
		panic(fmt.Sprintf("Adding Contract Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Contract Failed: %s", event.Err))
	}

}
