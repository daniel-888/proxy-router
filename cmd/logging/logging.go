package logging

import (
	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

func BoilerPlateFunc() (string, error) {
	msg := "Logging Package"
	return lumerinlib.BoilerPlateLibFunc(msg), nil // always returns no error
}