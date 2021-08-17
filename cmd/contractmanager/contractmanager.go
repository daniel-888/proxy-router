package contractmanager

import (
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func BoilerPlateFunc() (string, error) {
	msg := "Contract Manager Package"
	return lumerinlib.BoilerPlateLibFunc(msg), nil // always returns no error
}