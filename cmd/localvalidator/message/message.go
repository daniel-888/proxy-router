package message

import "encoding/json"
import "reflect"
import "fmt"
import "strings"

//JSON can desearealize to have these values
//base message that will be decoded to first to determine where message should go
type Message struct {
	//address is an ethereum address
	//MessageType is a string which describes which message is being sent
	//Message is a stringified JSON message
	Address, MessageType, Message string
}

//struct for new validator message
type NewValidator struct {
	BH, HashRate, Limit, Diff string
}

//struct for hashing message
type HashingInstance struct {
	Nonce, Time, Hash, Difficulty string
}

//struct for requesting information from validator
type GetValidationInfo struct {
	Hashes, Duration string
}

//struct to update the block header information within the validator
type UpdateBlockHeader struct {
	Version, PreviousBlockHash, MerkleRoot, Time, Difficulty string
}

type HashResult struct { //string will be true or false
	IsCorrect string
}

//function to take any given message struct and convert it into a string
func ConvertMessageToString(i interface{}) string {
	v := reflect.ValueOf(i)
	myString := "{"
	for j := 0; j < v.NumField(); j++ {
		var tempString []string
		newString := fmt.Sprintf(`"%s":"%s"`, v.Type().Field(j).Name, v.Field(j).Interface())
		tempString = []string{myString, newString}
		if myString == "{" {
			myString = strings.Join(tempString, "")
		} else {
			myString = strings.Join(tempString, ",")
		}
	}
	myString += "}"
	return myString
}

//request to compare the given hash with the calculated hash given the nonce and timestamp compared
//to the current block
func ReceiveHashingRequest(m string) HashingInstance {
	res := HashingInstance{}
	json.Unmarshal([]byte(m), &res)
	return res
}

//request to compare the given hash with the calculated hash given the nonce and timestamp compared
//to the current block
func ReceiveHashResult(m string) HashResult {
	res := HashResult{}
	json.Unmarshal([]byte(m), &res)
	return res
}

//request to make a new validation object
func ReceiveNewValidatorRequest(m string) NewValidator {
	res := NewValidator{}
	json.Unmarshal([]byte(m), &res)
	return res
}

//message requesting info from the validator. Validator returns everything
//and its up to the recipient to figure out what it is looking for
func ReceiveValidatorInfoRequest(m string) GetValidationInfo {
	res := GetValidationInfo{}
	json.Unmarshal([]byte(m), &res)
	return res
}

//message for when a new blockheader is updated
func ReceiveHeaderUpdateRequest(m string) UpdateBlockHeader {
	res := UpdateBlockHeader{}
	json.Unmarshal([]byte(m), &res)
	return res
}
