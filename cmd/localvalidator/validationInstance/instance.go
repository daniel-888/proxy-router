package validationInstance

import (
	"example.com/blockHeader"
	"example.com/message"
	"fmt"
	//"math"
	"errors"
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
func (v *Validator) IncomingHash(credential string, nonce string, time string, hash string, difficulty string) (message.HashResult, error) {
	if credential != v.PoolCredentials {
		return false, errors.Sprintf("Hashrate Hijacking Detected. Check pool user %s", credential)
	}
	var result = message.HashResult{} //initialize result here to use in error response
	calcHash := v.BH.HashInput(nonce, time)
	var hashingResult bool //temp until revised logic put in place
	hashAsBigInt, hashingErr := blockHeader.BlockHashToBigInt(calcHash)
	if hashingErr != nil {
		return result, hashingErr
	}
	var bigDifficulty *big.Int = blockHeader.DifficultyToBigInt(uint32(v.DifficultyTarget))

	if hashAsBigInt.Cmp(bigDifficulty) < 1 {
		hashingResult = true
	} else {
		hashingResult = false
	}
	if hashingResult {
		v.HashesAnalyzed++
	} 
	if v.HashrateRemaining() == false {
		v.closeOutContract() 
	}
	result.IsCorrect = strconv.FormatBool(hashingResult)
	return result, nil
}
//function to update the validators block header
func (v *Validator) UpdateBlockHeader(bh blockHeader.BlockHeader) {
	v.BH = bh
}
