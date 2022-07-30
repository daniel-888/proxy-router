package localvalidator

/*
comment by josh to see if git can see my ssh key
*/
import (
	"github.com/daniel-888/proxy-router/lumerinlib"
)

func BoilerPlateFunc() (string, error) {
	msg := "Local Validator Package"
	return lumerinlib.BoilerPlateLibFunc(msg), nil // always returns no error
}
