package main

import (
	"gitlab.com/TitanInd/lumerin/cmd/externalapi"
)


func main () {
	_,_,_,_,_,_ = externalapi.InitializeJSONRepos()
	externalapi.RunAPI()
}

