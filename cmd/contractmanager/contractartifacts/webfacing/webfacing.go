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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ledgerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cloneFactoryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"contractPurchase\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getListOfContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validationFee\",\"type\":\"uint256\"}],\"name\":\"setCreateRentalContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withValidator\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_ip_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620011bc380380620011bc8339818101604052810190620000379190620001c4565b620000576200004b620000e160201b60201c565b620000e960201b60201c565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200025e565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620001be8162000244565b92915050565b60008060408385031215620001de57620001dd6200023f565b5b6000620001ee85828601620001ad565b92505060206200020185828601620001ad565b9150509250929050565b600062000218826200021f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200024f816200020b565b81146200025b57600080fd5b50565b610f4e806200026e6000396000f3fe6080604052600436106100555760003560e01c8063081349951461005a57806341a902f9146100975780634b40a874146100c2578063715018a6146100de5780638da5cb5b146100f5578063f2fde38b14610120575b600080fd5b34801561006657600080fd5b50610081600480360381019061007c9190610970565b610149565b60405161008e9190610b1c565b60405180910390f35b3480156100a357600080fd5b506100ac61029f565b6040516100b99190610b37565b60405180910390f35b6100dc60048036038101906100d79190610831565b61034b565b005b3480156100ea57600080fd5b506100f3610400565b005b34801561010157600080fd5b5061010a610488565b6040516101179190610b1c565b60405180910390f35b34801561012c57600080fd5b50610147600480360381019061014291906107d7565b6104b1565b005b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633e45cccd8888888888336040518763ffffffff1660e01b81526004016101b196959493929190610c0f565b602060405180830381600087803b1580156101cb57600080fd5b505af11580156101df573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102039190610804565b9050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663015e9d4b826040518263ffffffff1660e01b81526004016102609190610b1c565b600060405180830381600087803b15801561027a57600080fd5b505af115801561028e573d6000803e3d6000fd5b505050508091505095945050505050565b6060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d70f8d6e6040518163ffffffff1660e01b815260040160006040518083038186803b15801561030957600080fd5b505afa15801561031d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103469190610927565b905090565b8673ffffffffffffffffffffffffffffffffffffffff166311afdfc98484848a8a8a6040518763ffffffff1660e01b815260040161038e96959493929190610b59565b600060405180830381600087803b1580156103a857600080fd5b505af11580156103bc573d6000803e3d6000fd5b505050507f0900ee509329f0c587c70faa0224f4e63bc738c9756346744a0c016ee96f1704876040516103ef9190610b1c565b60405180910390a150505050505050565b6104086105a9565b73ffffffffffffffffffffffffffffffffffffffff16610426610488565b73ffffffffffffffffffffffffffffffffffffffff161461047c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047390610bef565b60405180910390fd5b61048660006105b1565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6104b96105a9565b73ffffffffffffffffffffffffffffffffffffffff166104d7610488565b73ffffffffffffffffffffffffffffffffffffffff161461052d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052490610bef565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561059d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059490610bcf565b60405180910390fd5b6105a6816105b1565b50565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600061068861068384610c95565b610c70565b905080838252602082019050828560208602820111156106ab576106aa610e36565b5b60005b858110156106db57816106c1888261073c565b8452602084019350602083019250506001810190506106ae565b5050509392505050565b60006106f86106f384610cc1565b610c70565b90508281526020810184848401111561071457610713610e3b565b5b61071f848285610d8f565b509392505050565b60008135905061073681610ed3565b92915050565b60008151905061074b81610ed3565b92915050565b600082601f83011261076657610765610e31565b5b8151610776848260208601610675565b91505092915050565b60008135905061078e81610eea565b92915050565b600082601f8301126107a9576107a8610e31565b5b81356107b98482602086016106e5565b91505092915050565b6000813590506107d181610f01565b92915050565b6000602082840312156107ed576107ec610e45565b5b60006107fb84828501610727565b91505092915050565b60006020828403121561081a57610819610e45565b5b60006108288482850161073c565b91505092915050565b600080600080600080600060e0888a0312156108505761084f610e45565b5b600061085e8a828b01610727565b975050602061086f8a828b01610727565b96505060406108808a828b01610727565b95505060606108918a828b0161077f565b945050608088013567ffffffffffffffff8111156108b2576108b1610e40565b5b6108be8a828b01610794565b93505060a088013567ffffffffffffffff8111156108df576108de610e40565b5b6108eb8a828b01610794565b92505060c088013567ffffffffffffffff81111561090c5761090b610e40565b5b6109188a828b01610794565b91505092959891949750929550565b60006020828403121561093d5761093c610e45565b5b600082015167ffffffffffffffff81111561095b5761095a610e40565b5b61096784828501610751565b91505092915050565b600080600080600060a0868803121561098c5761098b610e45565b5b600061099a888289016107c2565b95505060206109ab888289016107c2565b94505060406109bc888289016107c2565b93505060606109cd888289016107c2565b92505060806109de888289016107c2565b9150509295509295909350565b60006109f78383610a03565b60208301905092915050565b610a0c81610d47565b82525050565b610a1b81610d47565b82525050565b6000610a2c82610d02565b610a368185610d25565b9350610a4183610cf2565b8060005b83811015610a72578151610a5988826109eb565b9750610a6483610d18565b925050600181019050610a45565b5085935050505092915050565b610a8881610d59565b82525050565b6000610a9982610d0d565b610aa38185610d36565b9350610ab3818560208601610d9e565b610abc81610e4a565b840191505092915050565b6000610ad4602683610d36565b9150610adf82610e5b565b604082019050919050565b6000610af7602083610d36565b9150610b0282610eaa565b602082019050919050565b610b1681610d85565b82525050565b6000602082019050610b316000830184610a12565b92915050565b60006020820190508181036000830152610b518184610a21565b905092915050565b600060c0820190508181036000830152610b738189610a8e565b90508181036020830152610b878188610a8e565b90508181036040830152610b9b8187610a8e565b9050610baa6060830186610a12565b610bb76080830185610a12565b610bc460a0830184610a7f565b979650505050505050565b60006020820190508181036000830152610be881610ac7565b9050919050565b60006020820190508181036000830152610c0881610aea565b9050919050565b600060c082019050610c246000830189610b0d565b610c316020830188610b0d565b610c3e6040830187610b0d565b610c4b6060830186610b0d565b610c586080830185610b0d565b610c6560a0830184610a12565b979650505050505050565b6000610c7a610c8b565b9050610c868282610dd1565b919050565b6000604051905090565b600067ffffffffffffffff821115610cb057610caf610e02565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610cdc57610cdb610e02565b5b610ce582610e4a565b9050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610d5282610d65565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610dbc578082015181840152602081019050610da1565b83811115610dcb576000848401525b50505050565b610dda82610e4a565b810181811067ffffffffffffffff82111715610df957610df8610e02565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b610edc81610d47565b8114610ee757600080fd5b50565b610ef381610d59565b8114610efe57600080fd5b50565b610f0a81610d85565b8114610f1557600080fd5b5056fea264697066735822122057bfa11df02d3c552590e53555332abd4052b5ebcd38b9b3451e049b7f6009d664736f6c63430008070033",
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

// SetCreateRentalContract is a paid mutator transaction binding the contract method 0x08134995.
//
// Solidity: function setCreateRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _validationFee) returns(address)
func (_Webfacing *WebfacingTransactor) SetCreateRentalContract(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validationFee *big.Int) (*types.Transaction, error) {
	return _Webfacing.contract.Transact(opts, "setCreateRentalContract", _price, _limit, _speed, _length, _validationFee)
}

// SetCreateRentalContract is a paid mutator transaction binding the contract method 0x08134995.
//
// Solidity: function setCreateRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _validationFee) returns(address)
func (_Webfacing *WebfacingSession) SetCreateRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validationFee *big.Int) (*types.Transaction, error) {
	return _Webfacing.Contract.SetCreateRentalContract(&_Webfacing.TransactOpts, _price, _limit, _speed, _length, _validationFee)
}

// SetCreateRentalContract is a paid mutator transaction binding the contract method 0x08134995.
//
// Solidity: function setCreateRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _validationFee) returns(address)
func (_Webfacing *WebfacingTransactorSession) SetCreateRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validationFee *big.Int) (*types.Transaction, error) {
	return _Webfacing.Contract.SetCreateRentalContract(&_Webfacing.TransactOpts, _price, _limit, _speed, _length, _validationFee)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x4b40a874.
//
// Solidity: function setPurchaseContract(address _contract, address _buyer, address _validator, bool _withValidator, string _ip_address, string _username, string _password) payable returns()
func (_Webfacing *WebfacingTransactor) SetPurchaseContract(opts *bind.TransactOpts, _contract common.Address, _buyer common.Address, _validator common.Address, _withValidator bool, _ip_address string, _username string, _password string) (*types.Transaction, error) {
	return _Webfacing.contract.Transact(opts, "setPurchaseContract", _contract, _buyer, _validator, _withValidator, _ip_address, _username, _password)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x4b40a874.
//
// Solidity: function setPurchaseContract(address _contract, address _buyer, address _validator, bool _withValidator, string _ip_address, string _username, string _password) payable returns()
func (_Webfacing *WebfacingSession) SetPurchaseContract(_contract common.Address, _buyer common.Address, _validator common.Address, _withValidator bool, _ip_address string, _username string, _password string) (*types.Transaction, error) {
	return _Webfacing.Contract.SetPurchaseContract(&_Webfacing.TransactOpts, _contract, _buyer, _validator, _withValidator, _ip_address, _username, _password)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x4b40a874.
//
// Solidity: function setPurchaseContract(address _contract, address _buyer, address _validator, bool _withValidator, string _ip_address, string _username, string _password) payable returns()
func (_Webfacing *WebfacingTransactorSession) SetPurchaseContract(_contract common.Address, _buyer common.Address, _validator common.Address, _withValidator bool, _ip_address string, _username string, _password string) (*types.Transaction, error) {
	return _Webfacing.Contract.SetPurchaseContract(&_Webfacing.TransactOpts, _contract, _buyer, _validator, _withValidator, _ip_address, _username, _password)
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
