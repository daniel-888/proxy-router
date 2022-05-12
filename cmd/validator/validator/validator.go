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
	channel		Channels
	Ps			*msgbus.PubSub
	Ctx     	context.Context
	MinerDiffs	lumerinlib.ConcurrentMap //miners being with a validation channel
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

	// Monitor Validation Stratum Messages
	validateEventChan := msgbus.NewEventChan()
	_, err := v.Ps.Sub(msgbus.ValidateMsg, "", validateEventChan)
	if err != nil {
		contextlib.Logf(v.Ctx, log.LevelError, "Failed to subscribe to validate events, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
		return err
	}
	go v.validateHandler(validateEventChan)

	return nil
}

func (v *MainValidator) validateHandler(ch msgbus.EventChan) {
	for {
		select {
		case <-v.Ctx.Done():
			contextlib.Logf(v.Ctx, log.LevelInfo, "Cancelling current validator context: cancelling minerHandler go routine")
			return

		case event := <-ch:
			if event.EventType == msgbus.PublishEvent {
				id := msgbus.ValidateID(event.ID)
				validateMsg := event.Data.(msgbus.Validate)
				minerID := validateMsg.MinerID
				destID := validateMsg.DestID
				msgType := validateMsg.Data.(type)

				miner,err := v.Ps.MinerGetWait(minerID)
				if err != nil {
					contextlib.Logf(v.Ctx, log.LevelPanic, "Failed to get miner, Fileline::%s, Error::%v", lumerinlib.FileLine(), err)
				}
			
				switch msgType {
				case msgbus.SetDifficulty:
					setDifficultyMsg := validateMsg.Data.(msgbus.SetDifficulty)
					if !v.MinerDiffs.Exists(minerID) {
						v.MinerDiffs.Set(minerID, setDifficultyMsg.Diff)
					} else {

					}

				case msgbus.Notify:
					notifyMsg := validateMsg.Data.(msgbus.Notify)
					version := notifyMsg.Version
					previousBlockHash := notifyMsg.PrevBlockHash
					time := notifyMsg.Ntime
					difficulty := v.MinerDiffs.Get(minerID).(int)
					difficultyStr := strconv.Itoa(difficulty)
					merkleBranches := notifyMsg.MerkleBranches
					
					blockHeader := ConvertBlockHeaderToString(BlockHeader{
						Version:           version,
						PreviousBlockHash: previousBlockHash,
						MerkleRoot:        "915c887a2d9ec3f566a648bedcf4ed30d0988e22268cfe43ab5b0cf8638999d3", 
						Time:              time,
						Difficulty:        difficultyStr,
					})

					var createMessage = Message{}
					createMessage.Address = string(minerID)
					createMessage.MessageType = "createNew"
					v.SendMessageToValidator(createMessage)

				case msgbus.Submit:
					submitMsg := validateMsg.Data.(msgbus.Submit)
					jobID := submitMsg.JobID
					extraNonce := submitMsg.Extraonce
					nTime := submitMsg.NTime
					nonce := submitMsg.NOnce

					var tabulationMessage = Message{}
					mySubmit := MiningSubmit{}
					mySubmit.WorkerName = ""
					mySubmit.JobID = jobID
					mySubmit.ExtraNonce2 = extraNonce
					mySubmit.NTime = nTime
					mySubmit.NOnce = nonce
					tabulationMessage.Address = string(minerID)
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
				}
			}
		}
	}
}