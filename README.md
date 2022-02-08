# Lumerin

Lumerin Node (aka ProxyRouter)

# Setup
1. Install Go
2. Install Ganache
3. Clone repo
4. Copy `lumerinconfig.example.json` to `lumerinconfig.json`
5. `cd` into `cmd` directory
6. Run `go build -o $GOPATH/bin/lumerin` // builds binary
7. `cd` into `cmd/contractmanager`
8. Run `go test -run Deployment` // will deploy contracts to Ganache
9. Edit `lumerinconfig.json`<br/>
    a) "mnemonic" will be generated in Ganache<br/>
    b) "cloneFactoryAddress" will be generated from Deployment test<br/>
    c) "cloneFactoryAddress" will be generated from Deployment test<br/>
10. Edit `run_lumerin.sh` (optional)
11. Run `./run_lumerin.sh`
