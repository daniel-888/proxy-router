// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package webfacing

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WebfacingMetaData contains all meta data concerning the Webfacing contract.
var WebfacingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ledgerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cloneFactoryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"contractPurchase\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getListOfContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setCreateRentalContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ip_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cfAddress\",\"type\":\"address\"}],\"name\":\"setUpdateCloneFactoryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ledgerAddress\",\"type\":\"address\"}],\"name\":\"setUpdateLedgerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620012ff380380620012ff8339818101604052810190620000379190620001c4565b620000576200004b620000e160201b60201c565b620000e960201b60201c565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200025e565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620001be8162000244565b92915050565b60008060408385031215620001de57620001dd6200023f565b5b6000620001ee85828601620001ad565b92505060206200020185828601620001ad565b9150509250929050565b600062000218826200021f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200024f816200020b565b81146200025b57600080fd5b50565b611091806200026e6000396000f3fe60806040526004361061007b5760003560e01c80637c88b5c71161004e5780637c88b5c7146101145780638da5cb5b14610130578063c938a5d41461015b578063f2fde38b146101985761007b565b80633ca58f961461008057806341a902f9146100a9578063515e104d146100d4578063715018a6146100fd575b600080fd5b34801561008c57600080fd5b506100a760048036038101906100a291906109b1565b6101c1565b005b3480156100b557600080fd5b506100be610281565b6040516100cb9190610cc7565b60405180910390f35b3480156100e057600080fd5b506100fb60048036038101906100f691906109b1565b61032d565b005b34801561010957600080fd5b506101126103ed565b005b61012e60048036038101906101299190610a0b565b610475565b005b34801561013c57600080fd5b50610145610524565b6040516101529190610cac565b60405180910390f35b34801561016757600080fd5b50610182600480360381019061017d9190610b23565b61054d565b60405161018f9190610cac565b60405180910390f35b3480156101a457600080fd5b506101bf60048036038101906101ba91906109b1565b6106a0565b005b6101c9610798565b73ffffffffffffffffffffffffffffffffffffffff166101e7610524565b73ffffffffffffffffffffffffffffffffffffffff161461023d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023490610d63565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d70f8d6e6040518163ffffffff1660e01b815260040160006040518083038186803b1580156102eb57600080fd5b505afa1580156102ff573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103289190610ada565b905090565b610335610798565b73ffffffffffffffffffffffffffffffffffffffff16610353610524565b73ffffffffffffffffffffffffffffffffffffffff16146103a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a090610d63565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6103f5610798565b73ffffffffffffffffffffffffffffffffffffffff16610413610524565b73ffffffffffffffffffffffffffffffffffffffff1614610469576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046090610d63565b60405180910390fd5b61047360006107a0565b565b8473ffffffffffffffffffffffffffffffffffffffff16632cf76de9848484886040518563ffffffff1660e01b81526004016104b49493929190610ce9565b600060405180830381600087803b1580156104ce57600080fd5b505af11580156104e2573d6000803e3d6000fd5b505050507f0900ee509329f0c587c70faa0224f4e63bc738c9756346744a0c016ee96f1704856040516105159190610cac565b60405180910390a15050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dc7a564487878787336040518663ffffffff1660e01b81526004016105b3959493929190610d83565b602060405180830381600087803b1580156105cd57600080fd5b505af11580156105e1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060591906109de565b9050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663015e9d4b826040518263ffffffff1660e01b81526004016106629190610cac565b600060405180830381600087803b15801561067c57600080fd5b505af1158015610690573d6000803e3d6000fd5b5050505080915050949350505050565b6106a8610798565b73ffffffffffffffffffffffffffffffffffffffff166106c6610524565b73ffffffffffffffffffffffffffffffffffffffff161461071c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071390610d63565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561078c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078390610d43565b60405180910390fd5b610795816107a0565b50565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600061087761087284610dfb565b610dd6565b9050808382526020820190508285602086028201111561089a57610899610f90565b5b60005b858110156108ca57816108b0888261092b565b84526020840193506020830192505060018101905061089d565b5050509392505050565b60006108e76108e284610e27565b610dd6565b90508281526020810184848401111561090357610902610f95565b5b61090e848285610ee9565b509392505050565b6000813590506109258161102d565b92915050565b60008151905061093a8161102d565b92915050565b600082601f83011261095557610954610f8b565b5b8151610965848260208601610864565b91505092915050565b600082601f83011261098357610982610f8b565b5b81356109938482602086016108d4565b91505092915050565b6000813590506109ab81611044565b92915050565b6000602082840312156109c7576109c6610f9f565b5b60006109d584828501610916565b91505092915050565b6000602082840312156109f4576109f3610f9f565b5b6000610a028482850161092b565b91505092915050565b600080600080600060a08688031215610a2757610a26610f9f565b5b6000610a3588828901610916565b9550506020610a4688828901610916565b945050604086013567ffffffffffffffff811115610a6757610a66610f9a565b5b610a738882890161096e565b935050606086013567ffffffffffffffff811115610a9457610a93610f9a565b5b610aa08882890161096e565b925050608086013567ffffffffffffffff811115610ac157610ac0610f9a565b5b610acd8882890161096e565b9150509295509295909350565b600060208284031215610af057610aef610f9f565b5b600082015167ffffffffffffffff811115610b0e57610b0d610f9a565b5b610b1a84828501610940565b91505092915050565b60008060008060808587031215610b3d57610b3c610f9f565b5b6000610b4b8782880161099c565b9450506020610b5c8782880161099c565b9350506040610b6d8782880161099c565b9250506060610b7e8782880161099c565b91505092959194509250565b6000610b968383610ba2565b60208301905092915050565b610bab81610ead565b82525050565b610bba81610ead565b82525050565b6000610bcb82610e68565b610bd58185610e8b565b9350610be083610e58565b8060005b83811015610c11578151610bf88882610b8a565b9750610c0383610e7e565b925050600181019050610be4565b5085935050505092915050565b6000610c2982610e73565b610c338185610e9c565b9350610c43818560208601610ef8565b610c4c81610fa4565b840191505092915050565b6000610c64602683610e9c565b9150610c6f82610fb5565b604082019050919050565b6000610c87602083610e9c565b9150610c9282611004565b602082019050919050565b610ca681610edf565b82525050565b6000602082019050610cc16000830184610bb1565b92915050565b60006020820190508181036000830152610ce18184610bc0565b905092915050565b60006080820190508181036000830152610d038187610c1e565b90508181036020830152610d178186610c1e565b90508181036040830152610d2b8185610c1e565b9050610d3a6060830184610bb1565b95945050505050565b60006020820190508181036000830152610d5c81610c57565b9050919050565b60006020820190508181036000830152610d7c81610c7a565b9050919050565b600060a082019050610d986000830188610c9d565b610da56020830187610c9d565b610db26040830186610c9d565b610dbf6060830185610c9d565b610dcc6080830184610bb1565b9695505050505050565b6000610de0610df1565b9050610dec8282610f2b565b919050565b6000604051905090565b600067ffffffffffffffff821115610e1657610e15610f5c565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610e4257610e41610f5c565b5b610e4b82610fa4565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610eb882610ebf565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610f16578082015181840152602081019050610efb565b83811115610f25576000848401525b50505050565b610f3482610fa4565b810181811067ffffffffffffffff82111715610f5357610f52610f5c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b61103681610ead565b811461104157600080fd5b50565b61104d81610edf565b811461105857600080fd5b5056fea264697066735822122069cc768362ee42cc89bcf3b4d8868980eaebb05211e0de48542af838e2c58c3c64736f6c63430008070033",
}

// WebfacingABI is the input ABI used to generate the binding from.
// Deprecated: Use WebfacingMetaData.ABI instead.
var WebfacingABI = WebfacingMetaData.ABI

// WebfacingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WebfacingMetaData.Bin instead.
var WebfacingBin = WebfacingMetaData.Bin

// DeployWebfacing deploys a new Ethereum contract, binding an instance of Webfacing to it.
func DeployWebfacing(auth *bind.TransactOpts, backend bind.ContractBackend, _ledgerAddress common.Address, _cloneFactoryAddress common.Address) (common.Address, *types.Transaction, *Webfacing, error) {
	parsed, err := WebfacingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WebfacingBin), backend, _ledgerAddress, _cloneFactoryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Webfacing{WebfacingCaller: WebfacingCaller{contract: contract}, WebfacingTransactor: WebfacingTransactor{contract: contract}, WebfacingFilterer: WebfacingFilterer{contract: contract}}, nil
}

// Webfacing is an auto generated Go binding around an Ethereum contract.
type Webfacing struct {
	WebfacingCaller     // Read-only binding to the contract
	WebfacingTransactor // Write-only binding to the contract
	WebfacingFilterer   // Log filterer for contract events
}

// WebfacingCaller is an auto generated read-only Go binding around an Ethereum contract.
type WebfacingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebfacingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WebfacingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebfacingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WebfacingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebfacingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WebfacingSession struct {
	Contract     *Webfacing        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WebfacingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WebfacingCallerSession struct {
	Contract *WebfacingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WebfacingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WebfacingTransactorSession struct {
	Contract     *WebfacingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WebfacingRaw is an auto generated low-level Go binding around an Ethereum contract.
type WebfacingRaw struct {
	Contract *Webfacing // Generic contract binding to access the raw methods on
}

// WebfacingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WebfacingCallerRaw struct {
	Contract *WebfacingCaller // Generic read-only contract binding to access the raw methods on
}

// WebfacingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WebfacingTransactorRaw struct {
	Contract *WebfacingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWebfacing creates a new instance of Webfacing, bound to a specific deployed contract.
func NewWebfacing(address common.Address, backend bind.ContractBackend) (*Webfacing, error) {
	contract, err := bindWebfacing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Webfacing{WebfacingCaller: WebfacingCaller{contract: contract}, WebfacingTransactor: WebfacingTransactor{contract: contract}, WebfacingFilterer: WebfacingFilterer{contract: contract}}, nil
}

// NewWebfacingCaller creates a new read-only instance of Webfacing, bound to a specific deployed contract.
func NewWebfacingCaller(address common.Address, caller bind.ContractCaller) (*WebfacingCaller, error) {
	contract, err := bindWebfacing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WebfacingCaller{contract: contract}, nil
}

// NewWebfacingTransactor creates a new write-only instance of Webfacing, bound to a specific deployed contract.
func NewWebfacingTransactor(address common.Address, transactor bind.ContractTransactor) (*WebfacingTransactor, error) {
	contract, err := bindWebfacing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WebfacingTransactor{contract: contract}, nil
}

// NewWebfacingFilterer creates a new log filterer instance of Webfacing, bound to a specific deployed contract.
func NewWebfacingFilterer(address common.Address, filterer bind.ContractFilterer) (*WebfacingFilterer, error) {
	contract, err := bindWebfacing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WebfacingFilterer{contract: contract}, nil
}

// bindWebfacing binds a generic wrapper to an already deployed contract.
func bindWebfacing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WebfacingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Webfacing *WebfacingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Webfacing.Contract.WebfacingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Webfacing *WebfacingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Webfacing.Contract.WebfacingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Webfacing *WebfacingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Webfacing.Contract.WebfacingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Webfacing *WebfacingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Webfacing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Webfacing *WebfacingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Webfacing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Webfacing *WebfacingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Webfacing.Contract.contract.Transact(opts, method, params...)
}

// GetListOfContracts is a free data retrieval call binding the contract method 0x41a902f9.
//
// Solidity: function getListOfContracts() view returns(address[])
func (_Webfacing *WebfacingCaller) GetListOfContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Webfacing.contract.Call(opts, &out, "getListOfContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetListOfContracts is a free data retrieval call binding the contract method 0x41a902f9.
//
// Solidity: function getListOfContracts() view returns(address[])
func (_Webfacing *WebfacingSession) GetListOfContracts() ([]common.Address, error) {
	return _Webfacing.Contract.GetListOfContracts(&_Webfacing.CallOpts)
}

// GetListOfContracts is a free data retrieval call binding the contract method 0x41a902f9.
//
// Solidity: function getListOfContracts() view returns(address[])
func (_Webfacing *WebfacingCallerSession) GetListOfContracts() ([]common.Address, error) {
	return _Webfacing.Contract.GetListOfContracts(&_Webfacing.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Webfacing *WebfacingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Webfacing.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Webfacing *WebfacingSession) Owner() (common.Address, error) {
	return _Webfacing.Contract.Owner(&_Webfacing.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Webfacing *WebfacingCallerSession) Owner() (common.Address, error) {
	return _Webfacing.Contract.Owner(&_Webfacing.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Webfacing *WebfacingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Webfacing.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Webfacing *WebfacingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Webfacing.Contract.RenounceOwnership(&_Webfacing.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Webfacing *WebfacingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Webfacing.Contract.RenounceOwnership(&_Webfacing.TransactOpts)
}

// SetCreateRentalContract is a paid mutator transaction binding the contract method 0xc938a5d4.
//
// Solidity: function setCreateRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length) returns(address)
func (_Webfacing *WebfacingTransactor) SetCreateRentalContract(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Webfacing.contract.Transact(opts, "setCreateRentalContract", _price, _limit, _speed, _length)
}

// SetCreateRentalContract is a paid mutator transaction binding the contract method 0xc938a5d4.
//
// Solidity: function setCreateRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length) returns(address)
func (_Webfacing *WebfacingSession) SetCreateRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Webfacing.Contract.SetCreateRentalContract(&_Webfacing.TransactOpts, _price, _limit, _speed, _length)
}

// SetCreateRentalContract is a paid mutator transaction binding the contract method 0xc938a5d4.
//
// Solidity: function setCreateRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length) returns(address)
func (_Webfacing *WebfacingTransactorSession) SetCreateRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Webfacing.Contract.SetCreateRentalContract(&_Webfacing.TransactOpts, _price, _limit, _speed, _length)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x7c88b5c7.
//
// Solidity: function setPurchaseContract(address _contract, address _buyer, string _ip_address, string _username, string _password) payable returns()
func (_Webfacing *WebfacingTransactor) SetPurchaseContract(opts *bind.TransactOpts, _contract common.Address, _buyer common.Address, _ip_address string, _username string, _password string) (*types.Transaction, error) {
	return _Webfacing.contract.Transact(opts, "setPurchaseContract", _contract, _buyer, _ip_address, _username, _password)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x7c88b5c7.
//
// Solidity: function setPurchaseContract(address _contract, address _buyer, string _ip_address, string _username, string _password) payable returns()
func (_Webfacing *WebfacingSession) SetPurchaseContract(_contract common.Address, _buyer common.Address, _ip_address string, _username string, _password string) (*types.Transaction, error) {
	return _Webfacing.Contract.SetPurchaseContract(&_Webfacing.TransactOpts, _contract, _buyer, _ip_address, _username, _password)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x7c88b5c7.
//
// Solidity: function setPurchaseContract(address _contract, address _buyer, string _ip_address, string _username, string _password) payable returns()
func (_Webfacing *WebfacingTransactorSession) SetPurchaseContract(_contract common.Address, _buyer common.Address, _ip_address string, _username string, _password string) (*types.Transaction, error) {
	return _Webfacing.Contract.SetPurchaseContract(&_Webfacing.TransactOpts, _contract, _buyer, _ip_address, _username, _password)
}

// SetUpdateCloneFactoryAddress is a paid mutator transaction binding the contract method 0x3ca58f96.
//
// Solidity: function setUpdateCloneFactoryAddress(address _cfAddress) returns()
func (_Webfacing *WebfacingTransactor) SetUpdateCloneFactoryAddress(opts *bind.TransactOpts, _cfAddress common.Address) (*types.Transaction, error) {
	return _Webfacing.contract.Transact(opts, "setUpdateCloneFactoryAddress", _cfAddress)
}

// SetUpdateCloneFactoryAddress is a paid mutator transaction binding the contract method 0x3ca58f96.
//
// Solidity: function setUpdateCloneFactoryAddress(address _cfAddress) returns()
func (_Webfacing *WebfacingSession) SetUpdateCloneFactoryAddress(_cfAddress common.Address) (*types.Transaction, error) {
	return _Webfacing.Contract.SetUpdateCloneFactoryAddress(&_Webfacing.TransactOpts, _cfAddress)
}

// SetUpdateCloneFactoryAddress is a paid mutator transaction binding the contract method 0x3ca58f96.
//
// Solidity: function setUpdateCloneFactoryAddress(address _cfAddress) returns()
func (_Webfacing *WebfacingTransactorSession) SetUpdateCloneFactoryAddress(_cfAddress common.Address) (*types.Transaction, error) {
	return _Webfacing.Contract.SetUpdateCloneFactoryAddress(&_Webfacing.TransactOpts, _cfAddress)
}

// SetUpdateLedgerAddress is a paid mutator transaction binding the contract method 0x515e104d.
//
// Solidity: function setUpdateLedgerAddress(address _ledgerAddress) returns()
func (_Webfacing *WebfacingTransactor) SetUpdateLedgerAddress(opts *bind.TransactOpts, _ledgerAddress common.Address) (*types.Transaction, error) {
	return _Webfacing.contract.Transact(opts, "setUpdateLedgerAddress", _ledgerAddress)
}

// SetUpdateLedgerAddress is a paid mutator transaction binding the contract method 0x515e104d.
//
// Solidity: function setUpdateLedgerAddress(address _ledgerAddress) returns()
func (_Webfacing *WebfacingSession) SetUpdateLedgerAddress(_ledgerAddress common.Address) (*types.Transaction, error) {
	return _Webfacing.Contract.SetUpdateLedgerAddress(&_Webfacing.TransactOpts, _ledgerAddress)
}

// SetUpdateLedgerAddress is a paid mutator transaction binding the contract method 0x515e104d.
//
// Solidity: function setUpdateLedgerAddress(address _ledgerAddress) returns()
func (_Webfacing *WebfacingTransactorSession) SetUpdateLedgerAddress(_ledgerAddress common.Address) (*types.Transaction, error) {
	return _Webfacing.Contract.SetUpdateLedgerAddress(&_Webfacing.TransactOpts, _ledgerAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Webfacing *WebfacingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Webfacing.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Webfacing *WebfacingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Webfacing.Contract.TransferOwnership(&_Webfacing.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Webfacing *WebfacingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Webfacing.Contract.TransferOwnership(&_Webfacing.TransactOpts, newOwner)
}

// WebfacingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Webfacing contract.
type WebfacingOwnershipTransferredIterator struct {
	Event *WebfacingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WebfacingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WebfacingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WebfacingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WebfacingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WebfacingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WebfacingOwnershipTransferred represents a OwnershipTransferred event raised by the Webfacing contract.
type WebfacingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Webfacing *WebfacingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WebfacingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Webfacing.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WebfacingOwnershipTransferredIterator{contract: _Webfacing.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Webfacing *WebfacingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WebfacingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Webfacing.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WebfacingOwnershipTransferred)
				if err := _Webfacing.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Webfacing *WebfacingFilterer) ParseOwnershipTransferred(log types.Log) (*WebfacingOwnershipTransferred, error) {
	event := new(WebfacingOwnershipTransferred)
	if err := _Webfacing.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WebfacingContractPurchaseIterator is returned from FilterContractPurchase and is used to iterate over the raw logs and unpacked data for ContractPurchase events raised by the Webfacing contract.
type WebfacingContractPurchaseIterator struct {
	Event *WebfacingContractPurchase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WebfacingContractPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WebfacingContractPurchase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WebfacingContractPurchase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WebfacingContractPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WebfacingContractPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WebfacingContractPurchase represents a ContractPurchase event raised by the Webfacing contract.
type WebfacingContractPurchase struct {
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractPurchase is a free log retrieval operation binding the contract event 0x0900ee509329f0c587c70faa0224f4e63bc738c9756346744a0c016ee96f1704.
//
// Solidity: event contractPurchase(address _contract)
func (_Webfacing *WebfacingFilterer) FilterContractPurchase(opts *bind.FilterOpts) (*WebfacingContractPurchaseIterator, error) {

	logs, sub, err := _Webfacing.contract.FilterLogs(opts, "contractPurchase")
	if err != nil {
		return nil, err
	}
	return &WebfacingContractPurchaseIterator{contract: _Webfacing.contract, event: "contractPurchase", logs: logs, sub: sub}, nil
}

// WatchContractPurchase is a free log subscription operation binding the contract event 0x0900ee509329f0c587c70faa0224f4e63bc738c9756346744a0c016ee96f1704.
//
// Solidity: event contractPurchase(address _contract)
func (_Webfacing *WebfacingFilterer) WatchContractPurchase(opts *bind.WatchOpts, sink chan<- *WebfacingContractPurchase) (event.Subscription, error) {

	logs, sub, err := _Webfacing.contract.WatchLogs(opts, "contractPurchase")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WebfacingContractPurchase)
				if err := _Webfacing.contract.UnpackLog(event, "contractPurchase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContractPurchase is a log parse operation binding the contract event 0x0900ee509329f0c587c70faa0224f4e63bc738c9756346744a0c016ee96f1704.
//
// Solidity: event contractPurchase(address _contract)
func (_Webfacing *WebfacingFilterer) ParseContractPurchase(log types.Log) (*WebfacingContractPurchase, error) {
	event := new(WebfacingContractPurchase)
	if err := _Webfacing.contract.UnpackLog(event, "contractPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
