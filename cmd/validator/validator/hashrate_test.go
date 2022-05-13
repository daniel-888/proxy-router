package validator

import(
	"testing"
	"context"
	"time"
	"fmt"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

func TestHashrate(t *testing.T) {
	ps := msgbus.New(10, nil)

	defaultpooladdr := "stratum+tcp://127.0.0.1:33334/"
	defaultDest := msgbus.Dest{
		ID:     msgbus.DestID(msgbus.DEFAULT_DEST_ID),
		NetUrl: msgbus.DestNetUrl(defaultpooladdr),
	}
	ps.PubWait(msgbus.DestMsg, msgbus.IDString(defaultDest.ID), defaultDest)

	mainContext := context.Background()
	cs := contextlib.NewContextStruct(nil, ps, nil, nil, nil)
	mainContext = context.WithValue(mainContext, contextlib.ContextKey, cs)

	v := MakeNewValidator(&mainContext)
	err := v.Start()
	if err != nil {
		panic(fmt.Sprintf("Validator failed to start: %v", err))
	}

	miner := msgbus.Miner{
		ID:                   msgbus.MinerID("MinerID01"),
		IP:                   "IpAddress1",
		CurrentHashRate:      0,
		State:                msgbus.OnlineState,
		Dest:                 defaultDest.ID,
		Contracts: 			  make(map[msgbus.ContractID]bool),	
	}
	ps.PubWait(msgbus.MinerMsg, msgbus.IDString(miner.ID), miner)

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

	notifyJobId := "616c4a28"
	notifyPrevBlock := "17c2c0507d5b4f32aa1ca39d82b83f16dfbf75d000093fd60000000000000000"
	notifyGen1 := "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2803cbf90a00046cf6e6610c"
	notifyGen2 := "0a746974616e2f6a74677261737369650affffffff024aa07e26000000001976a9143f1ad6ada38343e55cc4c332657e30a71c86b66188ac0000000000000000266a24aa21a9ed5617dd59c856a6ae1c00b6df9c4ed26727616d4d7b59f3eaacd16d810ff6dd3400000000"
	notifyMerkles := []interface{}{"e03d3ffb98db39658948c3e7e612b18d60a863b981b76d4fe17717d77818e4a3", "a3bc1f07702d7d17411fa3a7d686efc4ac6e45d7b3f3d5f6091194afcf5ce9ab", "5c60807c2e560c0ca85edc301136a4f3f442dc738fc2896cebd0088d7273d7ab", "62248c534cc4cf906f63ce71b3662bb986430c563225d7106ab1f14791ca6d90", "f957f25e5fac2ee6e1de76b3272ad5ddc751414ef2bc1d67fdcb50484eea9be7", "300c97cc0ce4179ef67dd0d974fd7f70bcb7f4d71a8b675f7f2955c780d94ecf", "243796dcef2ee48f1a5132c42abb9ab11ae47ddf87f77797390ef0cf0cdd3a3b", "5e26cbaa9f32657aac5be852e9c560e429f80019fda2da39fee69685086325f2", "df44bb117da9c9c79919d78a16255715fd4c62aa03e6b67c4667907ac897e15a", "94a8c79e827851ebb6f043e393db73111e754a7b0b55fb0c83c6cbc6235f1603", "1db100d99f37b95d429df28a11f0bef873dc63c8510787e2c9be199662236f06", "680aa7cd5a2007a9f2ae2a76ed2fb2e3231c8b6cd3ccaeea951a089a91912223"}
	notifyVersion := "20000000"
	notifyNbits := "170b8c8b"
	notifyNtime := "61e6f66c" 
	notifyClean := false

	// "prod.s9x8", //worker name
	// "d73b189a",  //job ID
	// "",          //extra nonce 2
	// "536dc802",  //time in bits
	// "222771801") //nonce
	workerNames := [3]string{"prod.s9x8","prod.s9x8","prod.s9x8"}
	jobIDs := [3]string{"d73b189a","616c4a28","616c4a28"}
	extraNonce2s := [3]string{"","77f9020000000000","5035030000000000"} //5a7a010000000000
	nTimes := [3]string{"536dc802","61e6f66c","61e6f66c"} // 61e6f66c
	nOnces := [3]string{"222771801","d83d2cf9","602849db"} // e6b732f5

	time.Sleep(time.Second * 10)

	ps.SendValidateSetDiff(context.Background(), miner.ID, defaultDest.ID, 289)
	time.Sleep(time.Second * 10)

	ps.SendValidateNotify(context.Background(), miner.ID, defaultDest.ID, notifyJobId, notifyPrevBlock, notifyGen1, notifyGen2, notifyMerkles, notifyVersion, notifyNbits, notifyNtime, notifyClean)
	time.Sleep(time.Second * 10)

	for i:=0;i<3;i++ {
		ps.SendValidateSubmit(context.Background(), workerNames[i], miner.ID, defaultDest.ID, jobIDs[i], extraNonce2s[i], nTimes[i], nOnces[i])
		time.Sleep(time.Second * 40)
	}
	time.Sleep(time.Second * 10)
}