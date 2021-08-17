package connectionmanager

import (
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func BoilerPlateFunc() (string, error) {
	msg := "Connection Manager Package"
	return lumerinlib.BoilerPlateLibFunc(msg), nil // always returns no error
}