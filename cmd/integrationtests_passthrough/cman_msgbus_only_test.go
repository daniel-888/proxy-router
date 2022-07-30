package integrationtestspass

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/daniel-888/proxy-router/cmd/log"
	"github.com/daniel-888/proxy-router/cmd/msgbus"
	"github.com/daniel-888/proxy-router/cmd/protocol/stratumv1"
	"github.com/daniel-888/proxy-router/lumerinlib"
	contextlib "github.com/daniel-888/proxy-router/lumerinlib/context"
)

type LocalConfig struct {
	BuyerNode           bool
	ListenIP            string
	ListenPort          string
	DefaultPoolAddr     string
	SchedulePassthrough bool
	LogFilePath         string
}

var miners []msgbus.DestID

func TestConnMgr(t *testing.T) {

	var sleepTime time.Duration = 30 * time.Second

	miners = make([]msgbus.DestID, 0)

	//
	// Load configuration
	//
	configPath := "../../lumerinconfig-connmgr-msgbus-test.json"
	configs, err := LoadTestConfiguration(configPath)
	if err != nil {
		panic(fmt.Sprintf("Loading Config Failed: %s", err))
	}

	//
	// Setup logging
	//
	l := log.New()
	l.SetLevel(log.LevelTrace)
	// l.SetLevel(log.LevelDebug)

	// logFile, err := os.OpenFile("/tmp/logfile", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	// logFile, err := os.OpenFile("/dev/stdout", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	// logFile, err := os.OpenFile(configs.LogFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	// logFile, err := os.OpenFile(configs.LogFilePath, os.O_WRONLY, 0666)
	if err != nil {
		l.Logf(log.LevelFatal, "error opening log file: %v", err)
	}
	// defer logFile.Close()
	// l.SetFormat(log.FormatJSON).SetOutput(logFile)
	// l.SetFormat(log.FormatJSON).SetOutput(logFile)
	// l.SetFormat(log.FormatJSON).SetOutput(os.Stdout)
	l.SetOutput(os.Stdout)
	// l.SetOutput(os.Stdout)

	l.Logf(log.LevelTrace, "Starting Logfile")

	//
	// Setup MsgBus
	//
	//ps := msgbus.New(10, l)
	ps := msgbus.New(10, l)

	mainContext := context.Background()

	defaultdst := "stratum+tcp://seanmcadam.switcher0:@mining.pool.titan.io:4242"
	seconddst := "stratum+tcp://seanmcadam.switcher0:@btc.f2pool.com:3333/"
	thirddst := "stratum+tcp://seanmcadam.switcher0:@ss.antpool.com:3333/"
	fourthdst := "stratum+tcp://seanmcadam.switcher0:@ss.antpool.com:3333/"
	fifthdst := "stratum+tcp://seanmcadam.switcher0:@us-east.stratum.slushpool.com:3333/"

	//defaultdst := "stratum+tcp://seanmcadam.switcher0:@pooltesta1.sbx.lumerin.io:4242/"
	//seconddst := "stratum+tcp://seanmcadam.switcher1:@pooltesta1.sbx.lumerin.io:4242/"

	//defaultdst := "stratum+tcp://seanmcadam.switcher0:@localhost:33335/"
	//seconddst := "stratum+tcp://seanmcadam.switcher1:@localhost:33335/"

	src := lumerinlib.NewNetAddr(lumerinlib.TCP, configs.ListenIP+":"+configs.ListenPort)

	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, defaultdst)
	//dst2 := lumerinlib.NewNetAddr(lumerinlib.TCP, seconddst)
	//dst3 := lumerinlib.NewNetAddr(lumerinlib.TCP, thirddst)
	//dst4 := lumerinlib.NewNetAddr(lumerinlib.TCP, fourthdst)
	//dst5 := lumerinlib.NewNetAddr(lumerinlib.TCP, fifthdst)

	cs := contextlib.NewContextStruct(nil, ps, l, src, dst)

	mainContext = context.WithValue(mainContext, contextlib.ContextKey, cs)

	defaultDest := &msgbus.Dest{
		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetUrl: msgbus.DestNetUrl(defaultdst),
	}

	secondDest := &msgbus.Dest{
		ID:     msgbus.DestID("SecondDest"),
		NetUrl: msgbus.DestNetUrl(seconddst),
	}

	thirdDest := &msgbus.Dest{
		ID:     msgbus.DestID("ThirdDest"),
		NetUrl: msgbus.DestNetUrl(thirddst),
	}

	fourthDest := &msgbus.Dest{
		ID:     msgbus.DestID("FourthDest"),
		NetUrl: msgbus.DestNetUrl(fourthdst),
	}

	fifthDest := &msgbus.Dest{
		ID:     msgbus.DestID("FifthDest"),
		NetUrl: msgbus.DestNetUrl(fifthdst),
	}

	publishRecord(ps, defaultDest)
	publishRecord(ps, secondDest)
	publishRecord(ps, thirdDest)
	publishRecord(ps, fourthDest)
	publishRecord(ps, fifthDest)

	miners = append(miners, secondDest.ID)
	miners = append(miners, thirdDest.ID)
	miners = append(miners, fourthDest.ID)
	miners = append(miners, fifthDest.ID)
	miners = append(miners, defaultDest.ID)

	//
	// Publish Default Dest record
	//
	//event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(msgbus.DEFAULT_DEST_ID), defaultDest)
	//if err != nil {
	//	panic(fmt.Sprintf("Adding Default Dest Failed: %s", err))
	//}
	//if event.Err != nil {
	//	panic(fmt.Sprintf("Adding Default Dest Failed: %s", event.Err))
	//}

	//
	// Publish alternate pool destination
	//
	//newTargetDest := msgbus.Dest{
	//	ID:     msgbus.DestID(msgbus.GetRandomIDString()),
	//	NetUrl: msgbus.DestNetUrl(seconddst),
	//}
	//ps.PubWait(msgbus.DestMsg, msgbus.IDString(newTargetDest.ID), newTargetDest)

	srcStrat, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", configs.ListenIP, configs.ListenPort))
	if err != nil {
		lumerinlib.PanicHere("")
	}

	stratum, err := stratumv1.NewListener(mainContext, srcStrat, defaultDest)
	if err != nil {
		panic(fmt.Sprintf("Stratum Protocol New() failed:%s", err))
	}

	stratum.RunOnce()
	// stratum.Run()

	//
	// Sleep while the system stablizes
	//
	time.Sleep(sleepTime)

	for {
		for u, v := range miners {
			fmt.Printf("\n***********\nPoint to[%d] %s\n************\n", u, v)
			setMinerDest(ps, v)
			time.Sleep(sleepTime)
		}
	}
}

//
//
//
func LoadTestConfiguration(filePath string) (configs LocalConfig, err error) {
	var data map[string]interface{}
	currDir, _ := os.Getwd()
	defer os.Chdir(currDir)

	if err != nil {
		panic(fmt.Errorf("error retrieving config file variable: %s", err))
	}
	file := filepath.Base(filePath)
	filePath = filepath.Dir(filePath)
	os.Chdir(filePath)

	configFile, err := os.Open(file)
	if err != nil {
		return configs, err
	}
	defer configFile.Close()
	byteValue, _ := ioutil.ReadAll(configFile)

	err = json.Unmarshal(byteValue, &data)

	configData := data["config"].(map[string]interface{})
	configs.BuyerNode = configData["buyerNode"].(bool)

	connConfigData := data["connection"].(map[string]interface{})
	configs.ListenIP = connConfigData["listenIP"].(string)
	configs.ListenPort = connConfigData["listenPort"].(string)
	configs.DefaultPoolAddr = connConfigData["defaultPoolAddr"].(string)

	schedConfigData := data["schedule"].(map[string]interface{})
	configs.SchedulePassthrough = schedConfigData["passthrough"].(bool)

	logConfigData := data["logging"].(map[string]interface{})
	configs.LogFilePath = logConfigData["filePath"].(string)

	return configs, err
}

//
// Publish Dest record
//
func publishRecord(ps *msgbus.PubSub, d *msgbus.Dest) {

	event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(d.ID), d)
	if err != nil {
		panic(fmt.Sprintf("Adding Default Dest:%s Failed: %s", d.ID, err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Default Dest:%s Failed: %s", d.ID, event.Err))
	}
}

//
// Set Miner
//
func setMiner(ps *msgbus.PubSub, miner msgbus.Miner) {

	e := ps.MinerSetWait(miner)
	if e != nil {
		panic(fmt.Sprintf("MinerSetWait() error:%s on %s", e, miner.ID))
	}
}

//
// Set Miner Dest
//
func setMinerDest(ps *msgbus.PubSub, dest msgbus.DestID) {
	miners, _ := ps.MinerGetAllWait()
	for _, v := range miners {
		minerptr, _ := ps.MinerGetWait(msgbus.MinerID(v))
		if minerptr != nil {
			miner := *minerptr
			miner.Dest = dest
			setMiner(ps, miner)
		}
	}
}
