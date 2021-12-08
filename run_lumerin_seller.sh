#!/bin/bash

cmd/lumerin_amd64/lumerin \
	--defaultpooladdr="stratum+tcp://127.0.0.1:33334/" \
	--buyer="false" \
	--configfile="cmd/configurationmanager/sellerconfig.json" \
	--listenip="127.0.0.1" \
	--listenport=3334 \
	--ethurl="wss://10.112.0.13:8545" \
  	--logfile="/tmp/lumerin1.log" \
	--loglevel=4 \
	--disableconnection="false" \
	--disablecontract="false" \
	--disableschedule="false" \

#	--configfile="lumerinconfig.json" \

#
#  -configfile string
#        Configuration File Path (default "default")
#  -contractmanager string
#        Contract Manager Account ID (default "default")
#  -ethurl string
#        GETH Node URL (default "default")
#  -help string
#        Display The help Screen (default "default")
#  -listenip string
#        IP to listen on (default "default")
#  -listenport string
#        Connection Port to listen on (default "default")
#  -logfile string
#        Log File Path (default "default")
#  -loglevel string
#        Logging level (default "default")
#
