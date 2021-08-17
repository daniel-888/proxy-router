package main

import (
	"fmt"
	"gitlab.com/TitanInd/lumerin/cmd/accountingmanager"
	"gitlab.com/TitanInd/lumerin/cmd/configurationmanager"
	"gitlab.com/TitanInd/lumerin/cmd/connectionmanager"
	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
	"gitlab.com/TitanInd/lumerin/cmd/contractmanager"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
	"gitlab.com/TitanInd/lumerin/cmd/localvalidator"
	"gitlab.com/TitanInd/lumerin/cmd/logging"
	"gitlab.com/TitanInd/lumerin/cmd/walletmanager"
)

func main () {
	fmt.Println(accountingmanager.BoilerPlateFunc())
	fmt.Println(configurationmanager.BoilerPlateFunc())
	fmt.Println(connectionmanager.BoilerPlateFunc())
	fmt.Println(connectionscheduler.BoilerPlateFunc())
	fmt.Println(contractmanager.BoilerPlateFunc())
	fmt.Println(externalapi.BoilerPlateFunc())
	fmt.Println(localvalidator.BoilerPlateFunc())
	fmt.Println(logging.BoilerPlateFunc())
	fmt.Println(walletmanager.BoilerPlateFunc())
}

