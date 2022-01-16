package validationInstance

import (
	"example.com/blockHeader"
	"example.com/message"
	"fmt"
	//"math"
	"strconv"
	"time"
	"math/big"
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
func (v *Validator) IncomingHash(nonce string, time string, hash string, difficulty string) (message.HashResult, error) {
	/*
		remove section to reflect changes to blockHeader package
	*/
	calcHash := v.BH.HashInput(nonce, time)
	uintDifficulty, strConvErr := strconv.ParseUint(v.BH.Difficulty, 16, 32)
	if strConvErr != nil {
		return message.HashResult{}, strConvErr
	}
	var hashingResult bool //temp until revised logic put in place
	hashAsBigInt, hashingErr := blockHeader.BlockHashToBigInt(calcHash)
	if hashingErr != nil {
		return message.HashResult{}, hashingErr
	}
	var bigDifficulty *big.Int = blockHeader.DifficultyToBigInt(uint32(uintDifficulty))

	if hashAsBigInt.Cmp(bigDifficulty) < 1 {
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
	return result, nil
}

//function to update the validators block header
func (v *Validator) UpdateBlockHeader(bh blockHeader.BlockHeader) {
	v.BH = bh
}
