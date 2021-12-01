package testmod

import (
	"fmt"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func MinersTouchAll(ps *msgbus.PubSub) {

	for {
		time.Sleep(60 * time.Second)

		event, err := ps.GetWait(msgbus.MinerMsg, "")
		if err != nil {
			panic(fmt.Sprintf(lumerinlib.FileLine()+" Error:%s\n", err))
		}

		if event.EventType != msgbus.GetIndexEvent {
			panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v\n", event))
		}

		if len(event.Data.(msgbus.IDIndex)) == 0 {
			fmt.Printf(lumerinlib.FileLine() + " No miners are online\n")
		} else {
			for _, i := range event.Data.(msgbus.IDIndex) {
				minerevent, err := ps.GetWait(msgbus.MinerMsg, i)
				if err != nil {
					panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", minerevent))
				}

				minerdata := minerevent.Data
				fmt.Printf("Miner Touching record for: %s\n", minerdata.(msgbus.Miner).ID)
				setevent, err := ps.SetWait(msgbus.MinerMsg, i, minerdata)
				if err != nil {
					panic(fmt.Sprint(lumerinlib.FileLine()+"Error:%v", setevent))
				}
			}
		}
	}

}
