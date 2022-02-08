# Lumerin

Lumerin Node (aka ProxyRouter)

# Setup
1. Install Go
2. Install Ganache
3. Clone repo
4. `cd` into `cmd` directory
5. Run `go build -o $GOPATH/bin/lumerin` // builds binary
6. `cd` into `cmd/contractmanager`
7. Run `go test -run Deployment` // will deploy contracts to Ganache
8. Edit `lumerinconfig.json`<br/>
    a) "mnemonic" will be generated in Ganache<br/>
    b) "cloneFactoryAddress" will be generated from Deployment test<br/>
    c) "cloneFactoryAddress" will be generated from Deployment test<br/>
9. Edit `run_lumerin.sh` (optional)
10. Run `./run_lumerin.sh`
