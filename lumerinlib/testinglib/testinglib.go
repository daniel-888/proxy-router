package testinglib

import (
	"math/rand"
)

var basePort int = 50000
var localhost string = "127.0.0.1"

//
//
//
func GetRandPort() (port int) {
	port = rand.Intn(10000) + basePort
	return port
}
