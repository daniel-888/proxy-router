package config

import (
	"flag"
	"fmt"
	"os"
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

func init() {

	//
	// Read in command line flags
	//
	for i, v := range ConfigMap {
		if v.flagname != "" {
			v.flagval = flag.String(v.flagname, "", v.flagusage)
			fmt.Printf("Val:'%s' set to:'%s', default:'%s'\n", v.flagname, *v.flagval, v.defval)
			ConfigMap[i] = v
		}
	}

	Init()

	//
	// Read in environmental variables
	//
	for _, v := range ConfigMap {
		i := os.Getenv(v.envname)
		if i != "" {
			v.envval = &i
		}
	}

}

func Init() {

	flag.Parse()

}

func ConfigGetFlagName(cc ConfigConst) (v string, e error) {
	if _, ok := ConfigMap[cc]; ok {
		v = ConfigMap[cc].flagname
	} else {
		e = fmt.Errorf("undefined config constant: %s", cc)
	}
	return
}

func ConfigGetEnvName(cc ConfigConst) (v string, e error) {
	if _, ok := ConfigMap[cc]; ok {
		v = ConfigMap[cc].envname
	} else {
		e = fmt.Errorf("undefined config constant: %s", cc)
	}
	return
}

func ConfigGetConfigName(cc ConfigConst) (v string, e error) {
	if _, ok := ConfigMap[cc]; ok {
		v = ConfigMap[cc].configname
	} else {
		e = fmt.Errorf("undefined config constant: %s", cc)
	}
	return
}

func ConfigGetVal(cc ConfigConst) (v string, e error) {
	if val, ok := ConfigMap[cc]; ok {
		if val.flagval != nil && *val.flagval != "" {
			v = *val.flagval
		} else if val.envval != nil {
			v = *val.envval
		} else if val.configval != nil {
			v = *val.configval
		} else {
			v = val.defval
		}
	} else {
		e = fmt.Errorf("undefined config constant: %s", cc)
	}

	return
}

// Local File takes precidence over remote download config
func LoadConfiguration() (e error) {
	e = nil
	filename, err := ConfigGetVal(ConfigConfigFilePath)
	if err != nil {
		panic(fmt.Errorf("error retrieving config file variable: %s", err))
	}

	_ = filename

	// Skip for now
	// Get the code out the door
	//

	return e
}
