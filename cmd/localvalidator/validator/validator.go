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
	"example.com/blockHeader"
	"example.com/channels"
	"example.com/message"
	"example.com/utils"
	"example.com/validationInstance"
	"fmt"
	"time"
)

//creates a channel object which can be used to access created validators
type Validator struct {
	channel channels.Channels
}

//creates a validator
func createValidator(bh blockHeader.BlockHeader, hashRate uint, limit uint, diff uint, messages chan message.Message) {
	go func() {
		myValidator := validationInstance.Validator{
			BH:               bh,
			StartTime:        time.Now(),
			HashesAnalyzed:   0,
			DifficultyTarget: diff,
			ContractHashRate: hashRate,
			ContractLimit:    limit,
		}
		for {
			//message is of type message, with messageType and content values
			m := <-messages
			if m.MessageType == "validate" {
				//potentially bubble up result of function call
				req := message.ReceiveHashingRequest(m.Message)
				result := myValidator.IncomingHash(req.Nonce, req.Time, req.Hash, req.Difficulty) //this function broadcasts a message
				newM := m
				newM.Message = message.ConvertMessageToString(result)
				messages <- newM //sends the message.HashResult struct into the channel
			} else if m.MessageType == "getHashCompleted" {
				fmt.Printf("%d", myValidator.HashesAnalyzed) //print number of hashes done
			} else if m.MessageType == "blockHeaderUpdate" {
				bh := blockHeader.ConvertToBlockHeader(m.Message)
				myValidator.UpdateBlockHeader(bh)
			}
		}
	}()
}

//entry point of all validators
//rite now it only returns whether or not a hash was successful. Future abilities should be able to return a response based on the input message
func (v *Validator) SendMessageToValidator(m message.Message) *message.Message {
	if m.MessageType == "createNew" {
		newChannel := v.channel.AddChannel(m.Address)
		//need to extract the block header out of m.Message
		creation := message.ReceiveNewValidatorRequest(m.Message)
		createValidator( //creation["BH"] is an embedded JSON object
			blockHeader.ConvertToBlockHeader(creation.BH),
			utils.ConvertStringToUint(creation.HashRate),
			utils.ConvertStringToUint(creation.Limit),
			utils.ConvertStringToUint(creation.Diff),
			newChannel,
		)
		return nil
	} else { //any other message will be sent to the validator, where the internal channel logic will handle the message
		channel, _ := v.channel.GetChannel(m.Address)
		channel <- m
		returnMessageMessage := <-channel
		//returnMessageMessage is a message of type message.HashResult
		var returnMessage = message.Message{}
		returnMessage.Address = m.Address
		returnMessage.MessageType = "response"
		returnMessage.Message = returnMessageMessage.Message
		return &returnMessage
	}
}

func MakeNewValidator() Validator {
	ch := channels.Channels{
		ValidationChannels: make(map[string]chan message.Message),
	}
	validator := Validator{channel: ch}
	return validator
}
