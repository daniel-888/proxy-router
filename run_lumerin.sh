#!/bin/sh

lumerin \
	--buyer="false" \
	--network="ropsten" \
	--mnemonic="" \
	--ethnodeaddress="" \
	--claimfunds="false" \
	--accountindex="0" \
	--listenip="127.0.0.1" \
	--listenport=3334 \
	--configfile="./ropstenconfig.json" \
	--configdownload="" \
	--logfile="/tmp/lumerin1.log" \
	--loglevel=4 \
	--defaultpooladdr="stratum+tcp://127.0.0.1:33334/" \
	--disableconnection="false" \
	--disablecontract="false" \
	--disableschedule="false" \
	--disableapi="false"


