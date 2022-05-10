/*
this is the main package where a goroutine is spun off to be the validator
incoming messages are a JSON object with the following key-value pairs:
	messageType: string
	contractAddress: string
	message: string

	messageType is the type of message, one of the following: "create", "validate", "getHashRate", "updateBlockHeader" [more]
	contractAddress will always be a single ethereum address
*/

package validator

import (
	"fmt"
	"strconv"
	"time"
)

//creates a channel object which can be used to access created validators
type MainValidator struct {
	channel Channels
}

//creates a validator
//func createValidator(bh blockHeader.BlockHeader, hashRate uint, limit uint, diff uint, messages chan message.Message) error{
func createValidator(bh BlockHeader, hashRate uint, limit uint, diff uint, pc string, messages chan Message) {
	go func() {
		myValidator := Validator{
			BH:               bh,
			StartTime:        time.Now(),
			HashesAnalyzed:   0,
			DifficultyTarget: diff,
			ContractHashRate: hashRate,
			ContractLimit:    limit,
			PoolCredentials:  pc, // pool login credentials
		}
		for {
			//message is of type message, with messageType and content values
			m := <-messages
			if m.MessageType == "validate" {
				//potentially bubble up result of function call
				req, hashingRequestError := ReceiveHashingRequest(m.Message)
				if hashingRequestError != nil {
					//error handling for hashing request error
				}
				result, hashingErr := myValidator.IncomingHash(req.WorkerName, req.NOnce, req.NTime) //this function broadcasts a message
				newM := m
				if hashingErr != "" { //make this error the message contents precedded by ERROR
					newM.Message = fmt.Sprintf("ERROR: error encountered validating a mining.submit message: %s\n", hashingErr)
				} else {
					newM.Message = ConvertMessageToString(result)
				}
				messages <- newM //sends the message.HashResult struct into the channel
			} else if m.MessageType == "getHashCompleted" {
				//print number of hashes done
				result := HashCount{}
				result.HashCount = strconv.FormatUint(uint64(myValidator.HashesAnalyzed), 10)
				newM := m
				newM.Message = ConvertMessageToString(result)
				messages <- newM
				//create a response object where the result is the hashes analyzed

			} else if m.MessageType == "blockHeaderUpdate" {
				bh := ConvertToBlockHeader(m.Message)
				myValidator.UpdateBlockHeader(bh)
			} else if m.MessageType == "closeValidator" {
				close(messages)
				return
			} else if m.MessageType == "tabulate" {
				/*
					this is similar to the validation message, but instead of returning a boolean value, it returns the current hashrate after the message is sent to it
				*/
				result := TabulationCount{}
				req, hashingRequestError := ReceiveHashingRequest(m.Message)
				if hashingRequestError != nil {
					//error handling for hashing request error
				}
				myValidator.IncomingHash(req.WorkerName, req.NOnce, req.NTime) //this function broadcasts a message
				hashrate := myValidator.UpdateHashrate()
				result.HashCount = hashrate
				newM := m
				newM.Message = ConvertMessageToString(result)
				messages <- newM

			}
		}
	}()
}

//entry point of all validators
//rite now it only returns whether or not a hash was successful. Future abilities should be able to return a response based on the input message
func (v *MainValidator) SendMessageToValidator(m Message) *Message {
	if m.MessageType == "createNew" {
		newChannel := v.channel.AddChannel(m.Address)
		//need to extract the block header out of m.Message
		creation, creationErr  := ReceiveNewValidatorRequest(m.Message)
		if creationErr != nil {
			//error handling for validator creation
		}
		useDiff, _ := strconv.ParseUint(creation.Diff, 16, 64)
		createValidator( //creation["BH"] is an embedded JSON object
			ConvertToBlockHeader(creation.BH),
			ConvertStringToUint(creation.HashRate),
			ConvertStringToUint(creation.Limit),
			uint(useDiff),
			creation.WorkerName,
			newChannel,
		)
		return nil
	} else { //any other message will be sent to the validator, where the internal channel logic will handle the message
		channel, _ := v.channel.GetChannel(m.Address)
		channel <- m
		returnMessageMessage := <-channel
		//returnMessageMessage is a message of type message.HashResult
		var returnMessage = Message{}
		returnMessage.Address = m.Address
		returnMessage.MessageType = "response"
		returnMessage.Message = returnMessageMessage.Message
		return &returnMessage
	}
}

func (v *MainValidator) ReceiveJSONMessage(b []byte, id string) {

	//blindly try to convert the message to a submit message. If it returns true
	//process the message
	msg := Message{}
	msg.Address = id
	submit, err := convertJSONToSubmit(b)
	//we don't care about the error message
	if err == nil {
		msg.MessageType = "validate"
		msg.Message = ConvertMessageToString(submit)
	}

	//blindly try to convert the message to a notify message.
	notify, err := convertJSONToNotify(b)
	if err == nil {
		msg.MessageType = "blockHeaderUpdate"
		msg.Message = ConvertMessageToString(notify)
	}
	//send message to validator. 
	v.SendMessageToValidator(msg)


}

//creates a new validator which can spawn multiple validation instances
func MakeNewValidator() MainValidator {
	ch := Channels{
		ValidationChannels: make(map[string]chan Message),
	}
	validator := MainValidator{channel: ch}
	return validator
}
