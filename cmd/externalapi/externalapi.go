package externalapi

import (
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func BoilerPlateFunc() (string, error) {
	msg := "External API Package"
	return lumerinlib.BoilerPlateLibFunc(msg), nil // always returns no error
}