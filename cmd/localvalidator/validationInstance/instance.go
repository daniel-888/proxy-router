package validationInstance

import (
	"example.com/blockHeader"
	"example.com/message"
	"fmt"
	//"math"
	"strconv"
	"time"
)

//an individual validator which will operate as a thread
type Validator struct {
	BH               blockHeader.BlockHeader
	StartTime        time.Time
	HashesAnalyzed   uint
	DifficultyTarget uint
	ContractHashRate uint
	ContractLimit    uint
}

//emits a message notifying whether or not the block was hashed correctly
func (v *Validator) blockAnalysisMessage(validHash bool) string {
	//revie this to return a JSON message
	if validHash {
		return "block was valid"
	} else {
		return "block was invalid"
	}
}

//function to determine if there's any remaining hashrate on the contract
func (v *Validator) HashrateRemaining() bool {
	return true
	//return v.ContractLimit > v.HashesAnalyzed
}

//function to send message to end contract
//intended to be called by the contract manager. The contract manager will call the smart contract
func (v *Validator) closeOutContract() {
	fmt.Println("web3 call to smart contract to initiate closeout procedure")
}

//receives a nonce and a hash, compares the two, and updates instance parameters
//need to modify to check to see if the resulting hash is below the given difficulty level
func (v *Validator) IncomingHash(ExtraNonce2 string, NOnce string, NTime string) message.HashResult {
	/*
		remove section to reflect changes to blockHeader package
	*/
	var hashingResult bool //temp until revised logic put in place
	//function from blockheader to get current difficulty
	currentDifficulty := 10
	calcHash := v.BH.HashInput(NOnce, NTime) 
	fmt.Printf("%v", ExtraNonce2) //temporarily here to get rid of compiler issue
	/*CHRIMBUS returns the hash of the mining.submit message in big endian format. 
	this will be changed to a boolean output which states if the resulting 
	hash is lower than the pool difficulty.
	*/
	newHash := fmt.Sprintf("%x", calcHash)
	if newHash <= currentDifficulty {
		hashingResult = true
	} else {
		hashingResult = false
	}
	if hashingResult {
		//increase the hashes calculated by 1
		v.HashesAnalyzed++
	} else {
		//message to indicate hashing failed
	}
	if v.HashrateRemaining() == false {
		v.closeOutContract()
	} //send out the message stating that the contract is closed with 0 hashes left
	var result = message.HashResult{}
	result.IsCorrect = strconv.FormatBool(hashingResult)
	return result
}

//function to update the validators block header
func (v *Validator) UpdateBlockHeader(bh blockHeader.BlockHeader) {
	v.BH = bh
}
