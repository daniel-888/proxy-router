package configurationmanager

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"errors"

    "gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func LoadConfiguration(file string, pkg string) (map[string]interface{}, error) {
	var data map[string]interface{}
	configFile, err := os.Open(file)
	if err != nil {
		return data, err
	}
	defer configFile.Close()
	byteValue,_ := ioutil.ReadAll(configFile)

	err = json.Unmarshal(byteValue, &data)
	return data[pkg].(map[string]interface{}),err
}
 
func DownloadConfig(fullURLFile string) {
    fileURL, err := url.Parse(fullURLFile)
    if err != nil {
        log.Fatal(err)
    }
    path := fileURL.Path
    segments := strings.Split(path, "/")
    fileName := segments[len(segments)-1]
    
    var file *os.File
    if flag.Lookup("test.v") == nil { // called from main.go
        file, err = os.Create("./configurationmanager/" + fileName)
        if err != nil {
            log.Fatal(err)
        }
    } else { // called from test srcipt
        file, err = os.Create(fileName)
        if err != nil {
            log.Fatal(err)
        }
    }
    
    client := http.Client{
        CheckRedirect: func(r *http.Request, via []*http.Request) error {
            r.URL.Opaque = r.URL.Path
            return nil
        },
    }
  
    resp, err := client.Get(fullURLFile)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
 
    size, err := io.Copy(file, resp.Body)
	if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
 
    fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)
}

func CommandLine() string {
	configFile := flag.String("pathToConfig", "./configurationmanager/testconfig.json", "lumerin node configuration")
	flag.Parse()
	fmt.Printf("Loading %s\n", *configFile)
	
	return *configFile
}

func LoadMsgBusFromConfig() (*msgbus.PubSub, msgbus.EventChan) {
    var configFile string
    
    // If not being called from test script, run file specified in command line flag
    if flag.Lookup("test.v") == nil {
        configFile = CommandLine()
        
        // If file does not exist locally, check if it is available to download from remote server
        _,err := os.Stat(configFile)
        if errors.Is(err, os.ErrNotExist) {
            DownloadConfig("https://lumerin-node-configs.s3.amazonaws.com/config.json")
        }
    } else { // Function called from test script
        configFile = "testconfig.json"
    }
   
    configMap,err := LoadConfiguration(configFile, "connectionManager")
    if err != nil {
        log.Fatal(err)
    }

    ech := make(msgbus.EventChan)
	ps := msgbus.New(1)

    dest := msgbus.Dest{
        ID:         "DestID01",
        NetHost:    msgbus.DestNetHost(configMap["defaultPoolHost"].(string)),   
	    NetPort:    msgbus.DestNetPort(configMap["defaultPoolPort"].(string)),
	    NetProto:   msgbus.DestNetProto(configMap["defaultPoolProto"].(string)),
    }

    seller := msgbus.Seller{
        ID:             "SellerID01",
        DefaultDest:    dest.ID,
    }

    config := msgbus.ConfigInfo{
        ID:             "ConfigID01",
        DefaultDest:    dest.ID,
	    Seller:         seller.ID,
    }

    ps.Pub(msgbus.DestMsg, msgbus.IDString(dest.ID), msgbus.Dest{})
	ps.Pub(msgbus.SellerMsg, msgbus.IDString(seller.ID), msgbus.Seller{})
    ps.Pub(msgbus.ConfigMsg, msgbus.IDString(config.ID), msgbus.ConfigInfo{})

	ps.Sub(msgbus.DestMsg, msgbus.IDString(dest.ID), ech)
	ps.Sub(msgbus.SellerMsg, msgbus.IDString(seller.ID), ech)
    ps.Sub(msgbus.ConfigMsg, msgbus.IDString(config.ID), ech)

    ps.Set(msgbus.DestMsg, msgbus.IDString(dest.ID), dest)
	ps.Set(msgbus.SellerMsg, msgbus.IDString(seller.ID), seller)
    ps.Set(msgbus.ConfigMsg, msgbus.IDString(config.ID), config)

    return ps,ech
}