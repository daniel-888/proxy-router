package validator

import "testing"
import "example.com/message"
import "example.com/blockHeader"

//function to generate a new validator message
func createTestValidator() message.Message {
	var returnMessage = message.Message{}
	returnMessage.Address = "123"
	returnMessage.MessageType = "createNew"

	//need function to convert message to string and string to message
	newValidatorMessage := message.NewValidator{
		BH:       blockHeader.ConvertBlockHeaderToString(createTestBlockHeader()),
		HashRate: "10",
		Limit:    "10",
		Diff:     "100000000000000000000000000000000000000000000000000000000000",
	}
	newValidatorString := message.ConvertMessageToString(newValidatorMessage)
	returnMessage.Message = newValidatorString
	return returnMessage
}

//function to provide a blockheader
//function to create a new blockheader for use in a new validator.
//using UpdateBlockHeader type since it contains all necessary information
func createTestBlockHeader() blockHeader.BlockHeader {
	return blockHeader.BlockHeader{
		Version:           "00000001",
		PreviousBlockHash: "000000000002d01c1fccc21636b607dfd930d31d01c3a62104612a1719011250",
		MerkleRoot:        "f3e94742aca4b5ef85488dc37c06c3282295ffec960994b2c0d5ac2a25a95766",
		Time:              "4d1b2237",
		Difficulty:        "1b04864c",
	}
}

func createBlockHeaderUpdate() {
}

//creates a submit message which is of the same type as stratum mining.submit
func createSubmitMessage(wn string, jid string, en2 string, nt string, no string) {
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
}

//creates a notify message which is of the same type as stratum mining.notify
func createNotifyMessage(wn string, jid string, en2 string, nt string, no string) {
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
}


func createHashSubmissionMessage() message.Message {
	var returnMessage = message.Message{}
	hashingMessage := message.HashingInstance{
		Nonce:      "274148111",
		Time:       "4d1b2237",
		Hash:       "000000000003ba27aa200b1cecaad478d2b00432346c3f1f3986da1afd33e506",
		Difficulty: "1b04864c",
	}
	returnMessage.Address = "123"
	returnMessage.MessageType = "validate"
	returnMessage.Message = message.ConvertMessageToString(hashingMessage)
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

//create a validator, send a hash, close validator
func TestCreateValidator(t *testing.T) {
	//creating a validator
	creationMessage := createTestValidator()
	validator := MakeNewValidator()
	validator.SendMessageToValidator(creationMessage)

	//sending a hash message to the validator
	hashingMessage := createHashSubmissionMessage()
	validator.SendMessageToValidator(hashingMessage)

	//closing the validator
	//still need a message to shut the validator down
}

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
	//still need a message to shut the validator down
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
	if resultingHashCount != "2" {
		t.Errorf("incorrect hashcount: %v", resultingHashCount)
	}
}

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
