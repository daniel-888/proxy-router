package main

import (
	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
)

func main () {
	config,connection,contract,dest,miner,seller := externalapi.InitializeJSONRepos()
	externalapi.RunAPI(config,connection,contract,dest,miner,seller)
}

