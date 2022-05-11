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
	"context"
	"fmt"
	"strconv"
	"time"

	"gitlab.com/TitanInd/lumerin/cmd/log"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
	"gitlab.com/TitanInd/lumerin/lumerinlib"
	contextlib "gitlab.com/TitanInd/lumerin/lumerinlib/context"
)

//creates a channel object which can be used to access created validators
type MainValidator struct {
	channel	Channels
	Ps		*msgbus.PubSub
	Ctx     context.Context
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
func MakeNewValidator(Ctx *context.Context) MainValidator {
	ch := Channels{
		ValidationChannels: make(map[string]chan Message),
	}
	ctxStruct := contextlib.GetContextStruct(*Ctx)
	validator := MainValidator{
		channel: ch,
		Ps: ctxStruct.MsgBus,
		Ctx: *Ctx,
	}
	return validator
}

func (v *MainValidator) Start() error {
	contextlib.Logf(v.Ctx, log.LevelInfo, "Validator Starting")

	// Monitor Miners
	minerEventChan := msgbus.NewEventChan()
	_, err := v.Ps.Sub(msgbus.MinerMsg, "", minerEventChan)
	if err != nil {
		contextlib.Logf(v.Ctx, log.LevelError, "Failed to subscribe to miner events, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		return err
	}
	go v.minerHandler(minerEventChan)

	// Monitor Miner Submits
	submitEventChan := msgbus.NewEventChan()
	_, err = v.Ps.Sub(msgbus.SubmitMsg, "", submitEventChan)
	if err != nil {
		contextlib.Logf(v.Ctx, log.LevelError, "Failed to subscribe to miner submit events, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		return err
	}
	go v.submitHandler(submitEventChan)

	return nil
}

func (v *MainValidator) minerHandler(ch msgbus.EventChan) {
	blockHeader := ConvertBlockHeaderToString(BlockHeader{
		Version:           "00000002",                                                         //bitcoin difficulty big endian
		PreviousBlockHash: "000000000000000067ecc744b5ae34eebbde14d21ca4db51652e4d67e155f07e", //big-endian expected
		MerkleRoot:        "915c887a2d9ec3f566a648bedcf4ed30d0988e22268cfe43ab5b0cf8638999d3", //big-endian expected
		Time:              "1399703554",                                                       //timestamp, not necessay and overwritten with a submission attempt
		Difficulty:        "1900896c",                                                         //big-endian the difficulty target that a block needs to meet
	})
	difficulty := "1d00ffff"

	for {
		select {
		case <-v.Ctx.Done():
			contextlib.Logf(v.Ctx, log.LevelInfo, "Cancelling current validator context: cancelling minerHandler go routine")
			return

		case event := <-ch:
			id := msgbus.MinerID(event.ID)

			switch event.EventType {

			//
			// Publish Event
			//
			case msgbus.PublishEvent:
				miner := event.Data.(msgbus.Miner)
				minerDest,err := v.Ps.DestGetWait(miner.Dest)
				if err != nil {
					contextlib.Logf(v.Ctx, log.LevelPanic, "Failed to get miner dest, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
				}
				workername := minerDest.Username() + ":" + minerDest.Password()

				if miner.State == msgbus.OnlineState {
					// create new validator for miner
					var createMessage = Message{}
					createMessage.Address = string(id)
					createMessage.MessageType = "createNew"
					createMessage.Message = ConvertMessageToString(NewValidator{
						BH:         blockHeader,
						HashRate:   "",        // not needed for now
						Limit:      "",        // not needed for now
						Diff:       difficulty,  //highest difficulty allowed using difficulty encoding
						WorkerName: workername, //worker name assigned to an individual mining rig. used to ensure that attempts are being allocated correctly
					})
					v.SendMessageToValidator(createMessage)
				}

				//
				// Unpublish Event
				//
			case msgbus.UnpublishEvent:
				var closeMessage = Message{}
				closeMessage.Address = string(id)
				closeMessage.MessageType = "closeValidator"
				closeMessage.Message = ""
				v.SendMessageToValidator(closeMessage)

				//
				// Update Event
				//
			case msgbus.UpdateEvent:
				miner := event.Data.(msgbus.Miner)
				if miner.State == msgbus.OfflineState {
					var closeMessage = Message{}
					closeMessage.Address = string(id)
					closeMessage.MessageType = "closeValidator"
					closeMessage.Message = ""
					v.SendMessageToValidator(closeMessage)
				}
				
			default:

			}
		}
	}
}

func (v *MainValidator) submitHandler(ch msgbus.EventChan) {
	for {
		select {
		case <-v.Ctx.Done():
			contextlib.Logf(v.Ctx, log.LevelInfo, "Cancelling current validator context: cancelling submitHandler go routine")
			return

		case event := <-ch:
			id := msgbus.SubmitID(event.ID)
		

			switch event.EventType {

			//
			// Publish Event
			//
			case msgbus.PublishEvent:
				submit := event.Data.(msgbus.Submit)
				miner,err := v.Ps.MinerGetWait(submit.Miner)
				if err != nil {
					contextlib.Logf(v.Ctx, log.LevelPanic, "Failed to get miner, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
				}
				minerDest,err := v.Ps.DestGetWait(miner.Dest)
				if err != nil {
					contextlib.Logf(v.Ctx, log.LevelPanic, "Failed to get miner dest, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
				}
				workername := minerDest.Username() + ":" + minerDest.Password()

				var tabulationMessage = Message{}
				mySubmit := MiningSubmit{}
				mySubmit.WorkerName = workername
				mySubmit.JobID = string(id)
				mySubmit.ExtraNonce2 = submit.Extraonce
				mySubmit.NTime = submit.NTime
				mySubmit.NOnce = submit.NOnce
				tabulationMessage.Address = string(submit.ID)
				tabulationMessage.MessageType = "tabulate"
				tabulationMessage.Message = ConvertMessageToString(mySubmit)

				m := v.SendMessageToValidator(tabulationMessage)
				hashrate,err := strconv.Atoi(m.Message)
				if err != nil {
					contextlib.Logf(v.Ctx, log.LevelPanic, "Failed to convert hashrate string to int, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
				}

				// set hashrate in miner
				miner.CurrentHashRate = hashrate
				v.Ps.MinerSetWait(*miner)

				//
				// Unpublish Event
				//
			case msgbus.UnpublishEvent:

				//
				// Update Event
				//
			case msgbus.UpdateEvent:
				
			default:

			}
		}
	}
}