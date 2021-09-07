package connectionmanager

import (
	"encoding/json"
	"fmt"

	"gitlab.com/TitanInd/lumerin/lumerinlib"
)

type stratumMethods string
type stratumStates string
type stratumErrors string

const (
	SErrNull           stratumErrors = "null"
	SErrUnknown        stratumErrors = "-1"
	SErrSrvNotFound    stratumErrors = "-2"
	SErrMethodNotFound stratumErrors = "-3"
	SErrFeeReq         stratumErrors = "-10"
	SErrSigReq         stratumErrors = "-20"
	SErrSigUnavail     stratumErrors = "-21"
	SErrUnkSigTyp      stratumErrors = "-22"
	SErrBadSig         stratumErrors = "-23"
)

const (
	StratumNew        stratumStates = "StratumNew"
	StratumSubscribed stratumStates = "StratumSubscribed"
	StratumAuthorized stratumStates = "StratumAuthorized"
	StratumMsgError   stratumStates = "StratumMsgError"
)

const (
	MINING_AUTHORIZE      stratumMethods = "mining.authorize"
	MINING_CONFIGURE      stratumMethods = "mining.configure"
	MINING_NOTIFY         stratumMethods = "mining.notify"
	MINING_SET_DIFFICULTY stratumMethods = "mining.set_difficulty"
	MINING_SET_TARGET     stratumMethods = "mining.set_target"
	MINING_SUBMIT         stratumMethods = "mining.submit"
	MINING_SUBSCRIBE      stratumMethods = "mining.subscribe"
)

// type jsonarray []interface{}
type jsonarray interface{}

type stratumMsg struct {
	ID     interface{} `json:"id,omitempty"`
	Method interface{} `json:"method,omitempty"`
	Error  interface{} `json:"error,omitempty"`
	Params interface{} `json:"params,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Reject interface{} `json:"reject-reason,omitempty"`
	//	Params interface{} `json:"params,omitempty"`
	//	Result interface{} `json:"result,omitempty"`
}

type request struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

// notice ID is always null
type notice struct {
	ID     *string     `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type responce struct {
	ID     int         `json:"id"`
	Error  *string     `json:"error"`
	Result interface{} `json:"result"`
	Reject interface{} `json:"reject-reason,omitempty"`
}

//------------------------------------------------------
//
//------------------------------------------------------
func unmarshalMsg(b []byte) (ret interface{}, err error) {

	msg := stratumMsg{}
	j := map[string]interface{}{}

	err = json.Unmarshal(b, &j)

	if err == nil {
		for key, value := range j {
			switch key {
			case "id":
				switch vtype := value.(type) {
				case float32:
					msg.ID = int(value.(float32))
				case float64:
					msg.ID = int(value.(float64))
				case int:
					msg.ID = value.(int)
				case string:
					msg.ID = -1
				case nil:
					msg.ID = nil
				default:
					panic(fmt.Sprintf("Value Type: %t", vtype))
				}

			case "method":
				switch vtype := value.(type) {
				case string:
					msg.Method = value.(string)
				case nil:
					msg.Method = ""
				default:
					panic(fmt.Sprintf("Value Type: %t", vtype))
				}

			case "error":
				switch vtype := value.(type) {
				case string:
					msg.Error = value.(string)
				case nil:
					msg.Error = nil
				default:
					panic(fmt.Sprintf("Value Type: %t", vtype))
				}

			case "reject-reason":
				msg.Reject = value
			//	switch vtype := value.(type) {
			//	case string:
			//		msg.Reject = value.(string)
			//	case nil:
			//		msg.Reject = nil
			//	default:
			//		panic(fmt.Sprintf("Value Type: %t", vtype))
			//	}

			case "params":
				msg.Params = value

			case "result":
				msg.Result = value

			default:
				panic(fmt.Sprintf("Key Value: %s", key))
			}
		}

		// Is this a Responce Msg?
		if msg.Result != nil {
			r := responce{}
			r.ID = msg.ID.(int)
			r.Result = msg.Result
			if msg.Error == nil {
				r.Error = nil
			} else {
				r.Error = msg.Error.(*string)
			}
			if msg.Reject == nil {
				r.Reject = nil
			} else {
				r.Reject = msg.Reject
			}
			ret = &r

			// Is this a Notice?
		} else if msg.ID == nil {
			ret = &notice{
				ID:     nil,
				Method: msg.Method.(string),
				Params: msg.Params,
			}

		} else {
			// Must be a Request
			ret = &request{
				ID:     msg.ID.(int),
				Method: msg.Method.(string),
				Params: msg.Params,
			}

		}

	} else {
		fmt.Printf(lumerinlib.FileLine()+"Error unmarshaling msg:%s\n", err)
	}

	return ret, err
}

//------------------------------------------------------
//
//------------------------------------------------------
func getStratumMsg(msg []byte) (ret interface{}, err error) {

	fmt.Printf(lumerinlib.FileLine()+"Stratum Msg: %s\n", msg)

	ret, err = unmarshalMsg(msg)

	if err != nil {
		panic(fmt.Sprintf(lumerinlib.FileLine() + "Error unmarshaling Notice msg\n"))
	}

	fmt.Printf(lumerinlib.FileLine()+"unmarshaled Stratum %T, Msg: %v\n", ret, ret)

	return ret, err
}

//------------------------------------------------------
//
//------------------------------------------------------
func createResponceMsg(r *responce) (msg []byte, err error) {

	err = nil
	fmt.Printf("Create Stratum Responce: %v\n", r)

	msg, err = json.Marshal(r)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error Marshaling Responce Err:%s\n", err)
		return nil, err
	}

	return msg, err
}

//------------------------------------------------------
//
//------------------------------------------------------
func createRequestMsg(r *request) (msg []byte, err error) {

	err = nil
	fmt.Printf("Create Stratum Request: %v\n", r)

	msg, err = json.Marshal(r)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error Marshaling Request Err:%s\n", err)
		return nil, err
	}

	return msg, err
}

//------------------------------------------------------
//
//------------------------------------------------------
func createNoticeMsg(n *notice) (msg []byte, err error) {

	err = nil
	fmt.Printf("Create Stratum Request: %v\n", n)

	msg, err = json.Marshal(n)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error Marshaling Request Err:%s\n", err)
		return nil, err
	}

	return msg, err
}
