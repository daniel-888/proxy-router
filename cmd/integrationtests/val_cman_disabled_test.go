package integrationtests

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

	"gitlab.com/TitanInd/lumerin/cmd/connectionscheduler"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/cmd/protocol/stratumv1"
	"gitlab.com/TitanInd/lumerin/cmd/validator/validator"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

type ValDisabledConfig struct {
	BuyerNode           bool
	ListenIP            string
	ListenPort          string
	DefaultPoolAddr     string
	SchedulePassthrough bool
	LogFilePath         string
}

func LoadValDisabledTestConfiguration(filePath string) (configs ValDisabledConfig, err error) {
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

func ValDisabledSimMain(ps *msgbus.PubSub, configs ValDisabledConfig, hashrateCalcLagTime time.Duration) msgbus.DestID {
	mainContext := context.Background()

	//
	// Add the various Context variables here
	// msgbus, logger, defailt listen address, defalt desitnation address
	//

	src := lumerinlib.NewNetAddr(lumerinlib.TCP, configs.ListenIP+":"+configs.ListenPort)
	dst := lumerinlib.NewNetAddr(lumerinlib.TCP, configs.DefaultPoolAddr)

	//
	// the proro argument (#1) gets set in the Protocol sus-system
	//
	cs := contextlib.NewContextStruct(nil, ps, nil, src, dst)

	//
	//  All of the various needed subsystem values get passed into the context here.
	//
	mainContext = context.WithValue(mainContext, contextlib.ContextKey, cs)

	//
	// Setup Default Dest
	//
	dest := &msgbus.Dest{
		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetUrl: msgbus.DestNetUrl(configs.DefaultPoolAddr),
	}

	event, err := ps.PubWait(msgbus.DestMsg, msgbus.IDString(msgbus.DEFAULT_DEST_ID), dest)
	if err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Default Dest Failed: %s", event.Err))
	}

	//
	// Setup Node Operator Msg
	//
	nodeOperator := msgbus.NodeOperator{
		ID:          msgbus.NodeOperatorID(msgbus.GetRandomIDString()),
		IsBuyer:     configs.BuyerNode,
		DefaultDest: dest.ID,
	}
	event, err = ps.PubWait(msgbus.NodeOperatorMsg, msgbus.IDString(nodeOperator.ID), nodeOperator)
	if err != nil {
		panic(fmt.Sprintf("Adding Node Operator Failed: %s", err))
	}
	if event.Err != nil {
		panic(fmt.Sprintf("Adding Node Operator Failed: %s", event.Err))
	}

	//
	// Fire up the StratumV1 Potocol
	//
	srcStrat, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", configs.ListenIP, configs.ListenPort))
	if err != nil {
		lumerinlib.PanicHere("")
	}

	stratum, err := stratumv1.NewListener(mainContext, srcStrat, dest)
	if err != nil {
		panic(fmt.Sprintf("Stratum Protocol New() failed:%s", err))
	}

	stratum.Run()

	//
	// Fire up schedule manager
	//
	csched, err := connectionscheduler.New(&mainContext, &nodeOperator, false, int(hashrateCalcLagTime))
	if err != nil {
		panic(fmt.Sprintf("Schedule manager failed: %v", err))
	}
	err = csched.Start()
	if err != nil {
		panic(fmt.Sprintf("Schedule manager failed to start: %v", err))
	}

	//
	// Fire up validator
	//
	v := validator.MakeNewValidator(&mainContext)
	err = v.Start()
	if err != nil {
		panic(fmt.Sprintf("Validator failed to start: %v", err))
	}
	
	return dest.ID
}

func TestValDisabled(t *testing.T) {
	configPath := "../../lumerinconfig.json"

	configs, err := LoadValDisabledTestConfiguration(configPath)
	if err != nil {
		panic(fmt.Sprintf("Loading Config Failed: %s", err))
	}

	var sleepTime time.Duration = 10 * time.Second

	ps := msgbus.New(10, nil)

	var hashrateCalcLagTime time.Duration = 20
	var reAdjustmentTime time.Duration = 3

	defaultDestID := ValDisabledSimMain(ps, configs, hashrateCalcLagTime)

	time.Sleep(sleepTime*10)

	fmt.Print("\n\n/// Multiple miners connecting to node ///\n\n\n")

	miner1 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID01"),
		IP:                   "IpAddress1",
		CurrentHashRate:      0,
		State:                msgbus.OnlineState,
		Dest:                 defaultDestID,
		Contracts: 			  make(map[msgbus.ContractID]bool),	
	}
	miner2 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID02"),
		IP:                   "IpAddress2",
		CurrentHashRate:      0,
		State:                msgbus.OnlineState,
		Dest:                 defaultDestID,
		Contracts: 			  make(map[msgbus.ContractID]bool),
	}
	miner3 := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID03"),
		IP:                   "IpAddress3",
		CurrentHashRate:      0,
		State:                msgbus.OnlineState,
		Dest:                 defaultDestID,
		Contracts: 			  make(map[msgbus.ContractID]bool),
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner1.ID), miner1)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner2.ID), miner2)
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner3.ID), miner3)
	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// Validator getting submits and updating miner hashrates ///\n\n\n")

	/*
	{"params": ["prod.s9x8", "d73b189a", "4900020000000000", "61e6f630", "70010699"], "id": 19809, "method": "mining.submit"}
	{"params": ["prod.s9x8", "d73b189a", "40d0020000000000", "61e6f630", "c38a8042"], "id": 19810, "method": "mining.submit"}
	{"params": ["prod.s9x8", "d73b189a", "d9e9020000000000", "61e6f630", "11745e4a"], "id": 19811, "method": "mining.submit"}
	{"id":6190,"jsonrpc":"2.0","method":"mining.notify","params":["616c4a28","17c2c0507d5b4f32aa1ca39d82b83f16dfbf75d000093fd60000000000000000","01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2803cbf90a00046cf6e6610c","0a746974616e2f6a74677261737369650affffffff024aa07e26000000001976a9143f1ad6ada38343e55cc4c332657e30a71c86b66188ac0000000000000000266a24aa21a9ed5617dd59c856a6ae1c00b6df9c4ed26727616d4d7b59f3eaacd16d810ff6dd3400000000",[ "e03d3ffb98db39658948c3e7e612b18d60a863b981b76d4fe17717d77818e4a3", "a3bc1f07702d7d17411fa3a7d686efc4ac6e45d7b3f3d5f6091194afcf5ce9ab", "5c60807c2e560c0ca85edc301136a4f3f442dc738fc2896cebd0088d7273d7ab", "62248c534cc4cf906f63ce71b3662bb986430c563225d7106ab1f14791ca6d90", "f957f25e5fac2ee6e1de76b3272ad5ddc751414ef2bc1d67fdcb50484eea9be7", "300c97cc0ce4179ef67dd0d974fd7f70bcb7f4d71a8b675f7f2955c780d94ecf", "243796dcef2ee48f1a5132c42abb9ab11ae47ddf87f77797390ef0cf0cdd3a3b", "5e26cbaa9f32657aac5be852e9c560e429f80019fda2da39fee69685086325f2", "df44bb117da9c9c79919d78a16255715fd4c62aa03e6b67c4667907ac897e15a", "94a8c79e827851ebb6f043e393db73111e754a7b0b55fb0c83c6cbc6235f1603", "1db100d99f37b95d429df28a11f0bef873dc63c8510787e2c9be199662236f06", "680aa7cd5a2007a9f2ae2a76ed2fb2e3231c8b6cd3ccaeea951a089a91912223" ],"20000000","170b8c8b","61e6f66c",false]}
	{"id":5896,"jsonrpc":"2.0","method":"mining.notify","params":["783647bc","17c2c0507d5b4f32aa1ca39d82b83f16dfbf75d000093fd60000000000000000","01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2803cbf90a00046cf6e6610c","0a746974616e2f6a74677261737369650affffffff024aa07e26000000001976a9143f1ad6ada38343e55cc4c332657e30a71c86b66188ac0000000000000000266a24aa21a9ed5617dd59c856a6ae1c00b6df9c4ed26727616d4d7b59f3eaacd16d810ff6dd3400000000",[ "e03d3ffb98db39658948c3e7e612b18d60a863b981b76d4fe17717d77818e4a3", "a3bc1f07702d7d17411fa3a7d686efc4ac6e45d7b3f3d5f6091194afcf5ce9ab", "5c60807c2e560c0ca85edc301136a4f3f442dc738fc2896cebd0088d7273d7ab", "62248c534cc4cf906f63ce71b3662bb986430c563225d7106ab1f14791ca6d90", "f957f25e5fac2ee6e1de76b3272ad5ddc751414ef2bc1d67fdcb50484eea9be7", "300c97cc0ce4179ef67dd0d974fd7f70bcb7f4d71a8b675f7f2955c780d94ecf", "243796dcef2ee48f1a5132c42abb9ab11ae47ddf87f77797390ef0cf0cdd3a3b", "5e26cbaa9f32657aac5be852e9c560e429f80019fda2da39fee69685086325f2", "df44bb117da9c9c79919d78a16255715fd4c62aa03e6b67c4667907ac897e15a", "94a8c79e827851ebb6f043e393db73111e754a7b0b55fb0c83c6cbc6235f1603", "1db100d99f37b95d429df28a11f0bef873dc63c8510787e2c9be199662236f06", "680aa7cd5a2007a9f2ae2a76ed2fb2e3231c8b6cd3ccaeea951a089a91912223" ],"20000000","170b8c8b","61e6f66c",false]}
	{"params": ["stage.s9x211", "783647bc", "8372000000000000", "61e6f66c", "0a3f74a7"], "id": 16801, "method": "mining.submit"}
	{"params": ["prod.s9x8", "616c4a28", "5a7a010000000000", "61e6f66c", "e6b732f5"], "id": 19812, "method": "mining.submit"}
	{"params": ["prod.s9x8", "616c4a28", "77f9020000000000", "61e6f66c", "d83d2cf9"], "id": 19813, "method": "mining.submit"}
	{"params": ["prod.s9x8", "616c4a28", "5035030000000000", "61e6f66c", "602849db"], "id": 19814, "method": "mining.submit"}
	{"id":6191,"jsonrpc":"2.0","method":"mining.notify","params":["42bd6b64","17c2c0507d5b4f32aa1ca39d82b83f16dfbf75d000093fd60000000000000000","01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2803cbf90a0004a8f6e6610c","0a746974616e2f6a74677261737369650affffffff02fe6f8126000000001976a9143f1ad6ada38343e55cc4c332657e30a71c86b66188ac0000000000000000266a24aa21a9edabebc22b545ae710e5ef8dc110c77870c5589a282567f36786a677b23cd0c8c800000000",[ "e03d3ffb98db39658948c3e7e612b18d60a863b981b76d4fe17717d77818e4a3", "a3bc1f07702d7d17411fa3a7d686efc4ac6e45d7b3f3d5f6091194afcf5ce9ab", "5c60807c2e560c0ca85edc301136a4f3f442dc738fc2896cebd0088d7273d7ab", "62248c534cc4cf906f63ce71b3662bb986430c563225d7106ab1f14791ca6d90", "f957f25e5fac2ee6e1de76b3272ad5ddc751414ef2bc1d67fdcb50484eea9be7", "ae2c5fb4cb6d2613fada24bf9eb731c176a3391cc1fe0262eb497ec4275779d2", "db033c650ce7e18c493019116aab554a2685082e27ec77aa7bf834c2da787c0b", "bb32f4b07a04676807e36c901c83d81a26529766caf2a0e611fa1c1f1b00f15d", "a5e29f9d83c401b4d0271b593ca2288f69ab79f1641e6f17a30f5cbb5141c30e", "0eb80a25f031588ccca0f9d246beabb24955c1da6c1402d8bd5b4f82ba6420a2", "be403c71eeb1bda016a246ff6a4ae2784cf8746142f17218563447e44f0251a1", "5690bd3e9f645f2f7b37d7532cb832f37a173d171bbeb54ea8b81e5ced39da99" ],"20000000","170b8c8b","61e6f6a8",false]}
	{"id":5897,"jsonrpc":"2.0","method":"mining.notify","params":["9e845ebf","17c2c0507d5b4f32aa1ca39d82b83f16dfbf75d000093fd60000000000000000","01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2803cbf90a0004a8f6e6610c","0a746974616e2f6a74677261737369650affffffff02fe6f8126000000001976a9143f1ad6ada38343e55cc4c332657e30a71c86b66188ac0000000000000000266a24aa21a9edabebc22b545ae710e5ef8dc110c77870c5589a282567f36786a677b23cd0c8c800000000",[ "e03d3ffb98db39658948c3e7e612b18d60a863b981b76d4fe17717d77818e4a3", "a3bc1f07702d7d17411fa3a7d686efc4ac6e45d7b3f3d5f6091194afcf5ce9ab", "5c60807c2e560c0ca85edc301136a4f3f442dc738fc2896cebd0088d7273d7ab", "62248c534cc4cf906f63ce71b3662bb986430c563225d7106ab1f14791ca6d90", "f957f25e5fac2ee6e1de76b3272ad5ddc751414ef2bc1d67fdcb50484eea9be7", "ae2c5fb4cb6d2613fada24bf9eb731c176a3391cc1fe0262eb497ec4275779d2", "db033c650ce7e18c493019116aab554a2685082e27ec77aa7bf834c2da787c0b", "bb32f4b07a04676807e36c901c83d81a26529766caf2a0e611fa1c1f1b00f15d", "a5e29f9d83c401b4d0271b593ca2288f69ab79f1641e6f17a30f5cbb5141c30e", "0eb80a25f031588ccca0f9d246beabb24955c1da6c1402d8bd5b4f82ba6420a2", "be403c71eeb1bda016a246ff6a4ae2784cf8746142f17218563447e44f0251a1", "5690bd3e9f645f2f7b37d7532cb832f37a173d171bbeb54ea8b81e5ced39da99" ],"20000000","170b8c8b","61e6f6a8",false]}
	*/


	notifyJobIds := [3]string{"e16730ab","3fc7779e","57ef4334"}
	notifyPrevBlocks := [3]string{"aa89eee9a63a5da9c451dd242e03c93147950a770006f2140000000000000000","aa89eee9a63a5da9c451dd242e03c93147950a770006f2140000000000000000","aa89eee9a63a5da9c451dd242e03c93147950a770006f2140000000000000000"}
	notifyGen1s := [3]string{"01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2103883c0b0004ab0480620c","01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2103883c0b0004e70480620c","01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2103883c0b0004230580620c"}
	notifyGen2s := [3]string{"5f546974616e2e696f5fffffffff0253854226000000001976a914288913831abe556f331ed13f73c98126f7a03a7588ac0000000000000000266a24aa21a9ed13927c1a804523a2ac252945392de21e7981c70afd5a0112a831c4c3d792b0c700000000", "5f546974616e2e696f5fffffffff0296114b26000000001976a914288913831abe556f331ed13f73c98126f7a03a7588ac0000000000000000266a24aa21a9ede9ed4f2cf9294336e3ac300e9f06d049be2d99c77d3fcffb6b0ba9b934f7179100000000", "5f546974616e2e696f5fffffffff02e54d4f26000000001976a914288913831abe556f331ed13f73c98126f7a03a7588ac0000000000000000266a24aa21a9ed2001039a0fee2bcfb931c32d15847418198c1a033599e630d8dc27bbc22c03b700000000"}
	notifyMerkless := [][]interface{}{{"e03d3ffb98db39658948c3e7e612b18d60a863b981b76d4fe17717d77818e4a3", "a3bc1f07702d7d17411fa3a7d686efc4ac6e45d7b3f3d5f6091194afcf5ce9ab", "5c60807c2e560c0ca85edc301136a4f3f442dc738fc2896cebd0088d7273d7ab", "62248c534cc4cf906f63ce71b3662bb986430c563225d7106ab1f14791ca6d90", "f957f25e5fac2ee6e1de76b3272ad5ddc751414ef2bc1d67fdcb50484eea9be7", "300c97cc0ce4179ef67dd0d974fd7f70bcb7f4d71a8b675f7f2955c780d94ecf", "243796dcef2ee48f1a5132c42abb9ab11ae47ddf87f77797390ef0cf0cdd3a3b", "5e26cbaa9f32657aac5be852e9c560e429f80019fda2da39fee69685086325f2", "df44bb117da9c9c79919d78a16255715fd4c62aa03e6b67c4667907ac897e15a", "94a8c79e827851ebb6f043e393db73111e754a7b0b55fb0c83c6cbc6235f1603", "1db100d99f37b95d429df28a11f0bef873dc63c8510787e2c9be199662236f06", "680aa7cd5a2007a9f2ae2a76ed2fb2e3231c8b6cd3ccaeea951a089a91912223"},
	{"90770a84646512101cadb195ee5e676388a004ea394cc7913051075dab2c27e6","3a34ac6e1042213afc32c759a773933ab4f3c915089d0d73577d10c8d19be586","ad8657af7f58c7711a0bc65ffa28585fb5fe8c479c094e16d67e7ef23e5be4f5","8755c4ef5a1f688857495e694d886f39ed0b9e793b4aa3c4b1e495e86d90cc45","0eb44323a055be6b30e8835f5b8044b3f554b9c32e44347fba4465b5fa7dcf32","43293bf2a84946dc176360a670670cfdefa8c8922350cafa5ff69d7f7a1191c0","7331ce7d53e7c3c8bf41638e6691daaf5f2f205f94c45d39bbbf1b0d77477328","8eda13259dc163ceaa4b31030c5ca90340e2c0ad0fa535e8e992d2b2a76ad1f7","dbe9422de38ad661d056b771b632652d6ee42b2ea7f10bda6477093824672bea","24e143a8dc116d935b8f1ec8164b42e3006937df062e8827cf43dd55e38db2d3","990242e21688e72c8daba7001e35213f65b1c1c1c9457029d41b8b23dc8531e2"},
	{"90770a84646512101cadb195ee5e676388a004ea394cc7913051075dab2c27e6","3a34ac6e1042213afc32c759a773933ab4f3c915089d0d73577d10c8d19be586","ad8657af7f58c7711a0bc65ffa28585fb5fe8c479c094e16d67e7ef23e5be4f5","4cbfb023b480689fa5f8ced400b52fe36f4b595b653a05f21c4efa51c6b6d52b","6b760ede8113d73b697eb5a394710cafdfd3c00302dc461cfc111ded76756fe9","9a24a044e1b35f2c154f0f41f512345f96698c41027246c3346a4f72c0c9648a","ff23b670a37ab183a0a743ad31c381e1baa951b67a286531a17baf036bded1cf","f244f9d828095d4e2373ec6edabc8deffc83856b0825206f436a187b4e9db718","7408592ca41ce815653505fe774615412d4c226a37414e56a6b48c7c7f76ec5d","23d3023a5bdb8a43b9869504d1a9e00337192ae28f030b8499cb7a1ed6d11733","b92cf5010bfbe2d760a6b411b28eeb923931054613f60dc690641763c7b1263f"}}
	notifyVersions := [3]string{"20000000","20000000","20000000"}
	notifyNbitss := [3]string{"170901ba","170901ba","170901ba"}
	notifyNtimes := [3]string{"628004ab" ,"628004e7","62800523"}
	notifyCleans := [3]bool{false,false,false}

	// "prod.s9x8", //worker name
	// "d73b189a",  //job ID
	// "",          //extra nonce 2
	// "536dc802",  //time in bits
	// "222771801") //nonce

	// workerNames := [3]string{"prod.s9x8","prod.s9x8","prod.s9x8"}
	// jobIDs := [3]string{"d73b189a","616c4a28","616c4a28"}
	// extraNonce2s := [3]string{"","77f9020000000000","5035030000000000"} //5a7a010000000000
	// nTimes := [3]string{"536dc802","61e6f66c","61e6f66c"} // 61e6f66c
	// nOnces := [3]string{"222771801","d83d2cf9","602849db"} // e6b732f5

	workerNames := [3]string{"seanmcadam.worker0","seanmcadam.worker0","seanmcadam.worker0"}
	jobIDs := [3]string{"57ef4334","57ef4334","57ef4334"}
	extraNonce2s := [3]string{"6654010000000000","e85f020000000000","bd77020000000000"} //5a7a010000000000
	nTimes := [3]string{"62800523","62800523","62800523"} // 61e6f66c
	nOnces := [3]string{"8d921ec5","f57dc4c0","5256e200"} // e6b732f5

	time.Sleep(time.Second * 5)

	ps.SendValidateSetDiff(context.Background(), miner.ID, defaultDest.ID, 65535) //486604799 4294901789
	time.Sleep(time.Second * 5)

	// ps.SendValidateNotify(context.Background(), miner.ID, defaultDest.ID, notifyJobId, notifyPrevBlock, notifyGen1, notifyGen2, notifyMerkles, notifyVersion, notifyNbits, notifyNtime, notifyClean)
	// time.Sleep(time.Second * 10)
	for i:=0;i<3;i++ {
		ps.SendValidateNotify(context.Background(), miner.ID, defaultDest.ID, notifyJobIds[i], notifyPrevBlocks[i], notifyGen1s[i], notifyGen2s[i], notifyMerkless[i], notifyVersions[i], notifyNbitss[i], notifyNtimes[i], notifyCleans[i])
		time.Sleep(time.Second * 5)
	}

	for i:=0;i<3;i++ {
		ps.SendValidateSubmit(context.Background(), workerNames[i], miner.ID, defaultDest.ID, jobIDs[i], extraNonce2s[i], nTimes[i], nOnces[i])
		time.Sleep(time.Second * 3)
	}

	m1,_ := ps.MinerGetWait(miner1.ID)
	m2,_ := ps.MinerGetWait(miner2.ID)
	m3,_ := ps.MinerGetWait(miner3.ID)

	fmt.Println("Miner 1 Current Hashrate: ", m1.CurrentHashRate)
	fmt.Println("Miner 2 Current Hashrate: ", m2.CurrentHashRate)
	fmt.Println("Miner 3 Current Hashrate: ", m3.CurrentHashRate)
	time.Sleep(time.Second * 2)

	fmt.Print("\n\n/// 2 New available contracts found ///\n\n\n")

	contract1 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID01"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    150,
	}
	contract2 := msgbus.Contract{
		IsSeller: true,
		ID:       msgbus.ContractID("ContractID02"),
		State:    msgbus.ContAvailableState,
		Price:    10,
		Limit:    10,
		Speed:    150,
	}
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	ps.PubWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(time.Second * 2)
	
	fmt.Print("\n\n/// Contract 1 purchased and now running ///\n\n\n")

	targetDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:55555/",
	}
	event,_ := ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest.ID), targetDest)
	targetDestID := msgbus.DestID(event.ID)

	contract1.State = msgbus.ContRunningState
	contract1.Buyer = "buyer1"
	contract1.Dest = targetDest.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(reAdjustmentTime*time.Second)

	var slicedMiner *msgbus.Miner
	var fullMiner1 *msgbus.Miner
	var fullMiner2 *msgbus.Miner
	minerIDs,_ := ps.MinerGetAllWait()
	for _,m := range minerIDs {
		miner,_ := ps.MinerGetWait(m)
		if miner.TimeSlice {
			slicedMiner = miner
		} else if miner.Contracts[contract1.ID] && !miner.TimeSlice{
			fullMiner1 = miner
		} else {
			fullMiner2 = miner
		}
	}
	if slicedMiner.Dest != targetDestID {
		t.Errorf("Sliced miner dest field incorrect")
	}

	if fullMiner1.Dest != targetDestID {
		t.Errorf("Full miner dest field incorrect")
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	slicedMiner,_ = ps.MinerGetWait(slicedMiner.ID)
	fullMiner1,_ = ps.MinerGetWait(fullMiner1.ID)
	if slicedMiner.Dest != defaultDestID {
		t.Errorf("Sliced miner dest field incorrect")
	}

	if fullMiner1.Dest != targetDestID {
		t.Errorf("Full miner dest field incorrect")
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	fmt.Print("\n\n/// Contract 2 purchased and now running ///\n\n\n")

	targetDest2 := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.GetRandomIDString()),
		NetUrl: "stratum+tcp://127.0.0.1:66666/",
	}
	event,_ = ps.PubWait(msgbus.DestMsg, msgbus.IDString(targetDest2.ID), targetDest2)
	targetDest2ID := msgbus.DestID(event.ID)

	contract2.State = msgbus.ContRunningState
	contract2.Buyer = "buyer2"
	contract2.Dest = targetDest2.ID
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	slicedMiner,_ = ps.MinerGetWait(slicedMiner.ID)
	fullMiner1,_ = ps.MinerGetWait(fullMiner1.ID)
	fullMiner2,_ = ps.MinerGetWait(fullMiner2.ID)
	if slicedMiner.Dest != targetDestID {
		t.Errorf("Sliced miner dest field incorrect, Dest in Miner: %s", slicedMiner.Dest)
	}

	if fullMiner1.Dest != targetDestID {
		t.Errorf("Full miner 1 dest field incorrect, Dest in Miner: %s", fullMiner1.Dest)
	}

	if fullMiner2.Dest != targetDest2ID {
		t.Errorf("Full miner 2 dest field incorrect, Dest in Miner: %s", fullMiner2.Dest)
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	time.Sleep(reAdjustmentTime*time.Second)

	slicedMiner,_ = ps.MinerGetWait(slicedMiner.ID)
	fullMiner1,_ = ps.MinerGetWait(fullMiner1.ID)
	fullMiner2,_ = ps.MinerGetWait(fullMiner2.ID)

	if slicedMiner.Dest != targetDest2ID {
		t.Errorf("Sliced miner dest field incorrect, Dest in Miner: %s", slicedMiner.Dest)
	}

	if fullMiner1.Dest != targetDestID {
		t.Errorf("Full miner 1 dest field incorrect, Dest in Miner: %s", fullMiner1.Dest)
	}

	if fullMiner2.Dest != targetDest2ID {
		t.Errorf("Full miner 2 dest field incorrect, Dest in Miner: %s", fullMiner2.Dest)
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	fmt.Print("\n\n/// Contract 1 closes out ///\n\n\n")

	contract1.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract1.ID), contract1)
	time.Sleep(hashrateCalcLagTime*time.Second)
	time.Sleep(reAdjustmentTime*time.Second)

	slicedMiner,_ = ps.MinerGetWait(slicedMiner.ID)
	fullMiner1,_ = ps.MinerGetWait(fullMiner1.ID)
	fullMiner2,_ = ps.MinerGetWait(fullMiner2.ID)

	if slicedMiner.Dest != targetDest2ID {
		t.Errorf("Sliced miner dest field incorrect, Dest in Miner: %s", slicedMiner.Dest)
	}

	if fullMiner1.Dest != defaultDestID {
		t.Errorf("Full miner 1 dest field incorrect, Dest in Miner: %s", fullMiner1.Dest)
	}

	if fullMiner2.Dest != targetDest2ID {
		t.Errorf("Full miner 2 dest field incorrect, Dest in Miner: %s", fullMiner2.Dest)
	}

	time.Sleep((hashrateCalcLagTime/2 - reAdjustmentTime)*time.Second)

	fmt.Print("\n\n/// Contract 2 closes out ///\n\n\n")

	contract2.State = msgbus.ContAvailableState
	ps.SetWait(msgbus.ContractMsg, msgbus.IDString(contract2.ID), contract2)
	time.Sleep(reAdjustmentTime*time.Second)


	time.Sleep(40*time.Second)
}