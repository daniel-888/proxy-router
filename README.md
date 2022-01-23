# Lumerin

Lumerin Node

# Run
1. Clone Repo and cd into "cmd" directory 
2. In terminal, type "go build -i -o $GOPATH/bin/lumerin" [Enter] to create lumerin executable in bin folder of your gopath 
3. Edit flag parameters in "run_lumerin.sh" for your specified run requirments
4. If using a json config, create a configuration file using the template provided in "lumerinconfig.json" and set the --configfile flag inside "run_lumerin.sh" to the relative path of your json config file 
5. In terminal, type "./run_lumerin.sh"
