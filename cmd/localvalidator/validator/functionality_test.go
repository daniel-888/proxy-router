package validator

import "testing"
import "example.com/message"
import "example.com/blockHeader"
import "strings"
import "fmt"

//function to generate a new validator message
func createTestValidator() message.Message {
	var returnMessage = message.Message{}
	returnMessage.Address = "123"
	returnMessage.MessageType = "createNew"

	//need function to convert message to string and string to message
	newValidatorMessage := message.NewValidator{
		BH:         blockHeader.ConvertBlockHeaderToString(createTestBlockHeader()),
		HashRate:   "10",
		Limit:      "10",
		Diff:       "100000000000000000000000000000000000000000000000000000000000",
		WorkerName: "prod.s9x8",
	}
	newValidatorString := message.ConvertMessageToString(newValidatorMessage)
	returnMessage.Message = newValidatorString
	return returnMessage
}

//function to provide a blockheader
//function to create a new blockheader for use in a new validator.
//using UpdateBlockHeader type since it contains all necessary information
//this information is also updated via the mining.Notify stratum v1 message
func createTestBlockHeader() blockHeader.BlockHeader {
	return blockHeader.BlockHeader{
		Version:           "00000001",
		PreviousBlockHash: "000000000002d01c1fccc21636b607dfd930d31d01c3a62104612a1719011250",
		MerkleRoot:        "f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766",
		Time:              "4d1b2237",
		Difficulty:        "1b04864c",
	}
}

//this returns an updated block header due to the block being attempted
//having been mined. output should be the same as createTestBloclHeader
func createBlockHeaderUpdate() blockHeader.BlockHeader{
	//this information needs to be changed. currently it is just a copy of the information in CreateTestBlockHeader
	return blockHeader.BlockHeader{
		Version:           "00000001",
		PreviousBlockHash: "000000000002d01c1fccc21636b607dfd930d31d01c3a62104612a1719011250",
		MerkleRoot:        "f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766",
		Time:              "4d1b2237",
		Difficulty:        "1b04864c",
	}
}

//creates a submit message which is of the same type as stratum mining.submit
func createSubmitMessage(wn string, jid string, en2 string, nt string, no string) message.Message {
	returnMessage := message.Message{}
	mySubmit := message.MiningSubmit{}
	mySubmit.WorkerName = wn
	mySubmit.JobID = jid
	mySubmit.ExtraNonce2 = en2
	mySubmit.NTime = nt
	mySubmit.NOnce = no
	returnMessage.Address = "123"
	returnMessage.MessageType = "validate"
	returnMessage.Message = message.ConvertMessageToString(mySubmit)
	fmt.Printf("%+v", returnMessage)
	return returnMessage
}

//creates a notify message which is of the same type as stratum mining.notify
func createNotifyMessage(JobID string, PreviousBlockHash string, GTP1 string, GTP2 string, MerkleList string, Version string, NBits string, NTime string, CleanJobs bool) message.Message {
	returnMessage := message.Message{}
	myNotify := message.MiningNotify{}
	myNotify.PreviousBlockHash = PreviousBlockHash
	myNotify.GTP1 = GTP1
	myNotify.GTP2 = GTP2
	myNotify.MerkleList = MerkleList
	myNotify.Version = Version
	myNotify.NBits = NBits
	myNotify.NTime = NTime
	myNotify.CleanJobs = CleanJobs
	returnMessage.Address = "123"
	returnMessage.MessageType = "validate"
	returnMessage.Message = message.ConvertMessageToString(myNotify)
	return returnMessage
}

func createHashSubmissionMessage() message.Message {
	returnMessage := createSubmitMessage("prod.s9x8", "d73b189a", "d9e9020000000000", "61e6f630", "11745e4a")
	return returnMessage
}

func createMinerNotifyMessage() message.Message {
	returnMessage := createNotifyMessage("616c4a28", "17c2c0507d5b4f32aa1ca39d82b83f16dfbf75d000093fd60000000000000000", "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff2803cbf90a00046cf6e6610c", "0a746974616e2f6a74677261737369650affffffff024aa07e26000000001976a9143f1ad6ada38343e55cc4c332657e30a71c86b66188ac0000000000000000266a24aa21a9ed5617dd59c856a6ae1c00b6df9c4ed26727616d4d7b59f3eaacd16d810ff6dd3400000000", `[ "e03d3ffb98db39658948c3e7e612b18d60a863b981b76d4fe17717d77818e4a3", "a3bc1f07702d7d17411fa3a7d686efc4ac6e45d7b3f3d5f6091194afcf5ce9ab", "5c60807c2e560c0ca85edc301136a4f3f442dc738fc2896cebd0088d7273d7ab", "62248c534cc4cf906f63ce71b3662bb986430c563225d7106ab1f14791ca6d90", "f957f25e5fac2ee6e1de76b3272ad5ddc751414ef2bc1d67fdcb50484eea9be7", "300c97cc0ce4179ef67dd0d974fd7f70bcb7f4d71a8b675f7f2955c780d94ecf", "243796dcef2ee48f1a5132c42abb9ab11ae47ddf87f77797390ef0cf0cdd3a3b", "5e26cbaa9f32657aac5be852e9c560e429f80019fda2da39fee69685086325f2", "df44bb117da9c9c79919d78a16255715fd4c62aa03e6b67c4667907ac897e15a", "94a8c79e827851ebb6f043e393db73111e754a7b0b55fb0c83c6cbc6235f1603", "1db100d99f37b95d429df28a11f0bef873dc63c8510787e2c9be199662236f06", "680aa7cd5a2007a9f2ae2a76ed2fb2e3231c8b6cd3ccaeea951a089a91912223" ]`, "20000000", "170b8c8b", "61e6f66c", false)
	return returnMessage
}

//create a message which asks the validator for the current hash count
func createHashCounterRequestMessage() message.Message {
	var returnMessage = message.Message{}
	returnMessage.Address = "123"
	returnMessage.MessageType = "getHashCompleted"
	returnMessage.Message = ""
	return returnMessage
}

//creates a generic close method to terminate the specific goroutine
func createCloseMethod() message.Message {
	var returnMessage = message.Message{}
	returnMessage.Address = "123"
	returnMessage.MessageType = "closeValidator"
	returnMessage.Message = ""
	return returnMessage
}

//create a validator, send a hash, close validator
func TestCreateValidator(t *testing.T) {
	//creating a validator
	creationMessage := createTestValidator()
	validator := MakeNewValidator()
	validator.SendMessageToValidator(creationMessage)

	//sending a hash message to the validator
	hashingMessage := createHashSubmissionMessage()
	hashingResult := validator.SendMessageToValidator(hashingMessage)
	if strings.Contains(hashingResult.Message, "ERROR") { //need to update this functionality
		t.Errorf("incorrect hash: %v", hashingResult.Message)
	}

	//closing the validator
	closingResult := validator.SendMessageToValidator(createCloseMethod())
	if strings.Contains(closingResult.Message, "ERROR") { //need to update this functionality
		t.Errorf("error closing validator: %v", closingResult.Message)
	}
}

/*
//create a validator, send a hash, confirm hash results in true, close validator
func TestCreateValidatorValidateHashCloseValidator(t *testing.T) {
	//creating a validator
	creationMessage := createTestValidator()
	validator := MakeNewValidator()
	validator.SendMessageToValidator(creationMessage)

	//sending a hash message to the validator
	hashingMessage := createHashSubmissionMessage()
	hashResult := validator.SendMessageToValidator(hashingMessage)
	//should be a message where the Message is of type HashResult
	if message.ReceiveHashResult(hashResult.Message).IsCorrect != "true" {
		t.Errorf("incorrect hash: %v", message.ReceiveHashResult(hashResult.Message))
	}

	//closing the validator
	validator.SendMessageToValidator(createCloseMethod())
}

//create a validator, send a hash, confirm hash results in true, close validator
func TestSubmit2HashesVerifyCount(t *testing.T) {
	//creating a validator
	creationMessage := createTestValidator()
	validator := MakeNewValidator()
	validator.SendMessageToValidator(creationMessage)

	//sending 2 mining.submit messages to validator
	hashingMessage := createHashSubmissionMessage()
	validator.SendMessageToValidator(hashingMessage)
	validator.SendMessageToValidator(hashingMessage)

	//creating hash request message
	hashRequestMessage := createHashCounterRequestMessage()

	//obtaining the hash verify response
	//resultingHashes := validator.SendMessageToValidator(hashRequestMessage)
	hashCount := validator.SendMessageToValidator(hashRequestMessage)
	resultingHashCount := message.ReceiveHashCount(hashCount.Message).HashCount
	validator.SendMessageToValidator(createCloseMethod())
	if resultingHashCount != "2" {
		t.Errorf("incorrect hashcount: %v", resultingHashCount)
	}
}
*/
/*
//create 2 validators, send each a hash, confirm hash results are true, close validators
func TestCreateTwoValidatorValidateHashCloseValidator(t *testing.T) {
}

//create a validator, send an invalid hash, confirm failure, close validator
func TestInvalidHashing(t *testing.T) {
}

//create 2 validators, send each a hash, close validators
func TestCreateTwoValidators(t *testing.T) {
}

//create a validator, send hashes to simulate 10 TH/s, confirm hashrate, closeout validator
func TestCaldulatedHashRate(t *testing.T) {
}

//create a validator, allow to run for 30 seconds, confirm that validator displays duration as 30 seconds
func TestTimeDurationTracking(t *testing.T) {
}

//create a validator, submit a hash, update blockheader, submit a hash, confirm that new hash conforms to new blockheader info
func TestBlockHeaderUpdate(t *testing.T) {
}
*/
