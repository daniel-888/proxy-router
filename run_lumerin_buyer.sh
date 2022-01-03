#!/bin/bash

cmd/lumerin_amd64/lumerin \
	--defaultpooladdr="stratum+tcp://127.0.0.1:33334/" \
	--disableconnection="false" \
	--disablecontract="true" \
	--disableschedule="false" \
	--buyer="true" \
	--configfile="cmd/configurationmanager/buyerconfig.json" \
	--listenip="127.0.0.1" \
	--listenport=3333 \
	--ethurl="wss://10.112.0.13:8545" \
  	--logfile="/tmp/lumerin2.log" \
	--loglevel=4 \


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
