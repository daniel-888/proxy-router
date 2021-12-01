package connectionmanager

//
// ToDo
// Move GetRandomIDString to be getRandomIDString containing it in this module
//
//
//
//

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
	CLIENT_MINING_AUTHORIZE stratumMethods = "mining.authorize"
	//	CLIENT_MINING_CAPABILITIES       stratumMethods = "mining.capabilities"
	CLIENT_MINING_EXTRANONCE         stratumMethods = "mining.extranonce.subscribe"
	CLIENT_MINING_SUBMIT             stratumMethods = "mining.submit"
	CLIENT_MINING_SUBSCRIBE          stratumMethods = "mining.subscribe"
	CLIENT_MINING_SUGGEST_DIFFICULTY stratumMethods = "mining.suggest_difficulty"
	CLIENT_MINING_SUGGEST_TARGET     stratumMethods = "mining.suggest_target"
	MINING_CONFIGURE                 stratumMethods = "mining.configure"
	MINING_SET_TARGET                stratumMethods = "mining.set_target"
	SERVER_GET_VERSION               stratumMethods = "client.get_version"
	SERVER_RECONNECT                 stratumMethods = "client.reconnect"
	SERVER_SHOW_MESSAGE              stratumMethods = "client.show_message"
	SERVER_MINING_NOTIFY             stratumMethods = "mining.notify"
	// Not sure about this method, got it once
	SERVER_MINING_PING           stratumMethods = "mining.ping"
	SERVER_MINING_SET_DIFFICULTY stratumMethods = "mining.set_difficulty"
	SERVER_MINING_SET_EXTRANONCE stratumMethods = "mining.set_extranonce"

//	SERVER_MINING_SET_GOAL           stratumMethods = "mining.set_goal"
)

//
// Used for recieving incoming stratum JSON  messages
//
type stratumMsg struct {
	ID      interface{} `json:"id,omitempty"`
	Jsonrpc interface{} `jsonrpc:"jsonrpc,omitempty"`
	Method  interface{} `json:"method,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Reject  interface{} `json:"reject-reason,omitempty"`
}

//
// Used to build outgoing JSON message
//
type request struct {
	ID     int      `json:"id"`
	Method string   `json:"method"`
	Params []string `json:"params"`
}

// notice ID is always null
type notice struct {
	ID     *string     `json:"id"`
	Params interface{} `json:"params"`
	Method string      `json:"method"`
}

type noticeMiningSetDifficulty struct {
	ID     *string `json:"id"`
	Method string  `json:"method"`
	Params []int   `json:"params"`
}

type noticeMiningNotify struct {
	ID      *string       `json:"id"`
	Jsonrpc *string       `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
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

			case "params":
				msg.Params = value
				//msg.Params = make([]interface{}, 0)
				//for _, v := range value.([]interface{}) {
				//	msg.Params = append(msg.Params, v)
				//}

			case "result":
				msg.Result = value

			case "jsonrpc":
				msg.Jsonrpc = value

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
			r := &request{
				ID:     msg.ID.(int),
				Method: msg.Method.(string),
				Params: make([]string, 0),
			}
			for _, v := range msg.Params.([]interface{}) {
				r.Params = append(r.Params, v.(string))
			}

			ret = r

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
// {"id":2,"method":"mining.authorize","params":["testrig.worker1",""]}
//------------------------------------------------------
func (r *request) getAuthName() (name string, err error) {

	if r.Method != string(CLIENT_MINING_AUTHORIZE) {
		return "", fmt.Errorf("wrong method, expetected mining.authorize")
	}

	fmt.Printf(" type:%T", r.Params)

	name = r.Params[0]
	// name = r.Params.([]string)[0]

	return name, err
}

//------------------------------------------------------
// extracts the mining.submit information being sent from the miner to the pool
//------------------------------------------------------
func (r *request) getSubmit() (userName string, jobId string, ExtraNonce2 string, nTime string, nonce string, err error) {

	if r.Method != string(CLIENT_MINING_SUBMIT) {
		return "", "","","","", fmt.Errorf("wrong method, expetected mining.submit")
	}

	fmt.Printf(" type:%T", r.Params)

	userName = r.Params[0]
	jobId = r.Params[1]
	ExtraNonce2 = r.Params[2]
	nTime = r.Params[3]
	nonce = r.Params[4]

	return userName, jobId, ExtraNonce2, nTime, nonce, err
}

//------------------------------------------------------
// extracts the mining.notify information from the pool to the miner
// returns a []byte to be used in parent function to construct a block header
//------------------------------------------------------
func (n *notice) getNotify() (msg []byte, err error) {

	if n.Method != string(CLIENT_MINING_SUBMIT) {
		return []byte(""), fmt.Errorf("wrong method, expetected mining.submit")
	}

	fmt.Printf(" type:%T", n.Params)

	/*
	code to obtain msg info and convert into a string
	*/

	return []byte(""), err
}

//------------------------------------------------------
// extracts the mining.set_difficulty information from the pool to the miner
// returns a string which can be used to update the validators target difficulty
//------------------------------------------------------
func (n *notice) getDifficulty() (msg string, err error) {

	if n.Method != string(SERVER_MINING_SET_DIFFICULTY) {
		return "", fmt.Errorf("wrong method, expetected mining.submit")
	}

	fmt.Printf(" type:%T", n.Params)

	/*
	difficulty = n.Params[0]
	code to obtain the difficulty and return as a string
	*/

	return "", err
}

//------------------------------------------------------
//
//------------------------------------------------------

func (r *request) createRequestMsg() (msg []byte, err error) {

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
func (r *responce) getAuthResult() (ret bool, err error) {

	_, ok := r.Result.(bool)
	if !ok {
		err = fmt.Errorf(lumerinlib.FileLine()+" result is wrong type:%T", r.Result)
	} else {
		ret = r.Result.(bool)
	}

	return ret, err
}

//------------------------------------------------------
//
//------------------------------------------------------
func (r *responce) createResponceMsg() (msg []byte, err error) {

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
func (n *notice) createNoticeMsg() (msg []byte, err error) {

	err = nil
	fmt.Printf("Create Stratum Notice: %v\n", n)

	switch n.Method {
	case string(SERVER_MINING_SET_DIFFICULTY):
		msg, err = n.createNoticeSetDifficultyMsg()
	case string(SERVER_MINING_NOTIFY):
		msg, err = n.createNoticeMiningNotify()
	default:
		msg, err = json.Marshal(n)
	}

	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error Marshaling Request Err:%s\n", err)
		return nil, err
	}

	return msg, err
}

//------------------------------------------------------
//
//------------------------------------------------------
func (n *notice) createNoticeSetDifficultyMsg() (msg []byte, err error) {

	fmt.Printf(lumerinlib.Funcname()+": %v\n", n)

	err = nil

	var nsd noticeMiningSetDifficulty
	nsd.ID = n.ID
	nsd.Method = n.Method
	nsd.Params = make([]int, 0)

	for _, v := range n.Params.([]interface{}) {
		i := int(v.(float64))
		if err != nil {
			panic("")
		}
		nsd.Params = append(nsd.Params, i)
	}

	msg, err = json.Marshal(nsd)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error Marshaling Request Err:%s\n", err)
		return nil, err
	}

	return msg, err
}

//------------------------------------------------------
//
// {
//   "params": [
// #0(string)   "613a0f04000001bc",
// #1(string)   "36847fbbe629819b9c0e23ddb4a80e68339e1b420002630c0000000000000000",
// #2(string)   "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff360394ad0a0004f14e3a61046ff4f4050c",
// #3(string)   "0a636b706f6f6c122f6d696e6564206279204c756d6572696e2fffffffff031cda8225000000001976a91422ddd9233f44ac2e9f183ec755adf134c12cdbf188ac0000000000000000266a24aa21a9ed852376e1fca95e42b3b1c080c3dc14b0db71c1b683511b0037b091c6f28acab96a413000000000001976a91422ddd9233f44ac2e9f183ec755adf134c12cdbf188ac00000000",
// #4([]string) [
// (string)      "773418c442067fdd5c3caf10653537041db14d13249cab724d9d892d8427a66a",
// (string)      "4126854f7bd3dc91bf666f53c35930685ee245239242ced1254f43e7b51b97e2",
// (string)      "d89213f7501f4f6123c5d24403801b7d978957e9ecbee82869fefb295025caff",
// (string)      "b4817f2f1e86914186c5acf715db97f753b84b9cc2cbd3a977e021df09ccf46d",
// (string)      "51c91bbfb65e328063dbfe020913a5e92c2973796f7cd84c74806e33eaf48116",
// (string)      "6e006d18ed55017612adf0e334b94d52e16b06f11adb14058a91caee161a304f",
// (string)      "633c5a641b57c0fc0fc9ed669d04686634f17ff34b6d509cc9a50c58e7cd9771",
// (string)      "e90773f4f44dc4a6a13e60956cad1612549e5c23a8f4ba42e760eb8661177464",
// (string)      "de5fc02be1faa3dbbb59e9799ea1fae886ab25e6b154413d2e2d35204fedbaf2",
// (string)      "79b109bdf26dd068446afa66c62f7d5ba30b179fcf032bb299f5a2591e0e3fce",
// (string)      "fe16f0630558f6564ec212ed700b1d5469b0a9d1cd39f4b7ce344d3d01d650b7",
// (string)      "03802c6be8643a09f8f74254ebf6f3704cfc622ab55f94687299fc32ca4a31da"
//              ],
// #5(string)   "20000000",
// #6(string)   "170f48e4",
// #7(string)   "613a4ee8",
// #8(bool)     true
//   ],
//   "id": null,
//   "method": "mining.notify"
// }
//
//
//------------------------------------------------------
func (n *notice) createNoticeMiningNotify() (msg []byte, err error) {

	fmt.Printf(lumerinlib.Funcname()+": %v\n", n)

	err = nil

	var nsd noticeMiningNotify
	nsd.ID = n.ID
	nsd.Method = n.Method

	if len(n.Params.([]interface{})) != 9 {
		panic("")
	}

	nsd.Params = make([]interface{}, 9)

	// nsd.Params = append(nsd.Params, a)
	for i, v := range n.Params.([]interface{}) {
		switch i {
		case 0:
			fallthrough
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			nsd.Params[i] = v.(string)

		case 4:
			arr := make([]string, 0)
			if len(v.([]interface{})) > 0 {
				for _, w := range v.([]interface{}) {
					arr = append(arr, w.(string))
				}
			}
			nsd.Params[i] = arr

		case 5:
			fallthrough
		case 6:
			fallthrough
		case 7:
			nsd.Params[i] = v.(string)

		case 8:
			if v == "true" {
				nsd.Params[i] = true
			} else if v == "false" {
				nsd.Params[i] = false
			}
		}
	}

	msg, err = json.Marshal(nsd)
	if err != nil {
		fmt.Printf(lumerinlib.FileLine()+"Error Marshaling Request Err:%s\n", err)
		return nil, err
	}

	return msg, err
}
