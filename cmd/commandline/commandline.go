package commandline

import (
	"flag"
)

//
// Format is
// -x single letter flag
// --xword word flag
//
//
// Connection Manager
// --listenip=127.0.0.1
// --listenport=3333
//
// Etherium Node (contract Manager)
// --ethip=127.0.0.1
// --ethport=7545
//
// config file
// --configfile=lumerinconfig.json
//
//

var EthIP string
var EthPort string
var ListenIP string
var ListenPort string
var ConfigFile string

func init() {
	const (
		ethip      = "ethip"
		ethport    = "ethport"
		listenip   = "listenip"
		listenport = "listenport"
		configfile = "configfile"
	)

	flag.StringVar(&EthIP, ethip, "127.0.0.1", "Etherium IP to connect to")
	flag.StringVar(&EthPort, ethport, "7545", "Etherium Port to listen on")
	flag.StringVar(&ListenIP, listenip, "127.0.0.1", "Local IP to listen on")
	flag.StringVar(&ListenPort, listenport, "3333", "Local Port to listen on")
	flag.StringVar(&ConfigFile, configfile, "./lumerinconfig.json", "Config File Location")

}

func Parse() {

	flag.Parse()

}
