package configurationmanager

import (
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func BoilerPlateFunc() (string, error) {
	msg := "Configuration Manager Package"
	return lumerinlib.BoilerPlateLibFunc(msg), nil // always returns no error
}