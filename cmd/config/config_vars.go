package config

//
// Define all configuration variables here
//

type ConfigConst string

const (
	ConfigHelp                       ConfigConst = "ConfigHelp"
	ConfigContractEthURL             ConfigConst = "ConfigContractEthURL"
	ConfigContractManagerAcct        ConfigConst = "ConfigContractManagerAcct"
	ConfigContractMinShareAmtPerMin  ConfigConst = "ConfigContractMinShareAmtPerMin"
	ConfigContractMinShareAvePerHour ConfigConst = "ConfigContractMinShareAvePerHour"
	ConfigContractShareDropTolerance ConfigConst = "ConfigContractShareDropTolerance"
	ConfigConnectionListenIP         ConfigConst = "ConfigConnectionListenIP"
	ConfigConnectionListenPort       ConfigConst = "ConfigConnectionListenPort"
	ConfigConfigFilePath             ConfigConst = "ConfigConfigFilePath"
	ConfigConfigDownloadPath         ConfigConst = "ConfigConfigDownloadPath"
	ConfigLogFilePath                ConfigConst = "ConfigLogFilePath"
	ConfigLogLevel                   ConfigConst = "ConfigLogLevel"
)

// Config Structure
type configitem struct {
	flagname   string
	flagusage  string
	envname    string
	configname string
	defval     string
	configval  *string
	envval     *string
	flagval    *string
}

//
// Define all Configuration constants that can be read in from the command line and the environment
//
var ConfigMap = map[ConfigConst]configitem{
	ConfigHelp: {
		flagname:   "help",
		flagusage:  "Display The help Screen",
		envname:    "",
		configname: "",
		defval:     "",
		configval:  nil,
		envval:     nil,
		flagval:    nil,
	},
	ConfigConnectionListenIP: {
		flagname:   "listenip",
		flagusage:  "IP to listen on",
		envname:    "LISTENIP",
		configname: "connect.listenip",
		defval:     "127.0.0.1",
		configval:  nil,
		envval:     nil,
		flagval:    nil,
	},
	ConfigConnectionListenPort: {
		flagname:   "listenport",
		flagusage:  "Connection Port to listen on",
		envname:    "LISTENPORT",
		configname: "connect.listenport",
		defval:     "3333",
		configval:  nil,
		envval:     nil,
		flagval:    nil,
	},
	ConfigContractEthURL: {
		flagname:   "ethurl",
		flagusage:  "GETH Node URL",
		envname:    "ETHURL",
		configname: "contract.ethurl",
		defval:     "wss://127.0.0.1:7545",
		configval:  nil,
		envval:     nil,
		flagval:    nil,
	},
	ConfigContractManagerAcct: {
		flagname:   "contractmanager",
		flagusage:  "Contract Manager Account ID",
		envname:    "CONTRACTMANAGER",
		configname: "contract.manager",
		defval:     "",
		configval:  nil,
		envval:     nil,
		flagval:    nil,
	},
	ConfigContractMinShareAmtPerMin: {
		flagname:   "",
		flagusage:  "",
		envname:    "",
		configname: "contractManager.MinShareAmtPerMin",
		defval:     "10",
		configval:  nil,
		envval:     nil,
		flagval:    nil,
	},
	ConfigContractMinShareAvePerHour: {
		flagname:   "",
		flagusage:  "",
		envname:    "",
		configname: "contractManager.MinShareAvePerHour",
		defval:     "10",
		configval:  nil,
		envval:     nil,
		flagval:    nil,
	},

	ConfigContractShareDropTolerance: {
		flagname:   "",
		flagusage:  "",
		envname:    "",
		configname: "contractManager.ShareDropTolerance",
		defval:     "10",
		configval:  nil,
		envval:     nil,
		flagval:    nil,
	},
	ConfigConfigFilePath: {
		flagname:  "configfile",
		flagusage: "Configuration File Path",
		envname:   "CONFIGFILE",
		defval:    "lumerinconfig.json",
		configval: nil,
		envval:    nil,
		flagval:   nil,
	},
	ConfigConfigDownloadPath: {
		flagname:  "configdownload",
		flagusage: "Configuration Download Path",
		envname:   "CONFIGDOWNLOAD",
		defval:    "",
		configval: nil,
		envval:    nil,
		flagval:   nil,
	},
	ConfigLogLevel: {
		flagname:  "loglevel",
		flagusage: "Logging level",
		envname:   "LOGLEVEL",
		defval:    "",
		configval: nil,
		envval:    nil,
		flagval:   nil,
	},
	ConfigLogFilePath: {
		flagname:  "logfile",
		flagusage: "Log File Path",
		envname:   "LOGFILEPATH",
		defval:    "lumerin.log",
		configval: nil,
		envval:    nil,
		flagval:   nil,
	},
}
