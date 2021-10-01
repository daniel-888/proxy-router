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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ledgerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cloneFactoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"contractCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getListOfContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setCreateRentalContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cfAddress\",\"type\":\"address\"}],\"name\":\"setUpdateCloneFactoryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ledgerAddress\",\"type\":\"address\"}],\"name\":\"setUpdateLedgerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001157380380620011578339818101604052810190620000379190620002e8565b620000576200004b6200020560201b60201c565b6200020d60201b60201c565b83600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d8811ede836040518263ffffffff1660e01b81526004016200013691906200036b565b600060405180830381600087803b1580156200015157600080fd5b505af115801562000166573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639ac4e37f826040518263ffffffff1660e01b8152600401620001c791906200036b565b600060405180830381600087803b158015620001e257600080fd5b505af1158015620001f7573d6000803e3d6000fd5b5050505050505050620003db565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620002e281620003c1565b92915050565b60008060008060808587031215620003055762000304620003bc565b5b60006200031587828801620002d1565b94505060206200032887828801620002d1565b93505060406200033b87828801620002d1565b92505060606200034e87828801620002d1565b91505092959194509250565b620003658162000388565b82525050565b60006020820190506200038260008301846200035a565b92915050565b600062000395826200039c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620003cc8162000388565b8114620003d857600080fd5b50565b610d6c80620003eb6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063715018a61161005b578063715018a6146100d85780638da5cb5b146100e2578063c938a5d414610100578063f2fde38b1461011c5761007d565b80633ca58f961461008257806341a902f91461009e578063515e104d146100bc575b600080fd5b61009c60048036038101906100979190610848565b610138565b005b6100a66101f8565b6040516100b39190610a7f565b60405180910390f35b6100d660048036038101906100d19190610848565b6102a4565b005b6100e0610364565b005b6100ea6103ec565b6040516100f79190610a3b565b60405180910390f35b61011a600480360381019061011591906108eb565b610415565b005b61013660048036038101906101319190610848565b6105a7565b005b61014061069f565b73ffffffffffffffffffffffffffffffffffffffff1661015e6103ec565b73ffffffffffffffffffffffffffffffffffffffff16146101b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ab90610ac1565b60405180910390fd5b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d70f8d6e6040518163ffffffff1660e01b815260040160006040518083038186803b15801561026257600080fd5b505afa158015610276573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061029f91906108a2565b905090565b6102ac61069f565b73ffffffffffffffffffffffffffffffffffffffff166102ca6103ec565b73ffffffffffffffffffffffffffffffffffffffff1614610320576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031790610ac1565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61036c61069f565b73ffffffffffffffffffffffffffffffffffffffff1661038a6103ec565b73ffffffffffffffffffffffffffffffffffffffff16146103e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103d790610ac1565b60405180910390fd5b6103ea60006106a7565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dc7a564486868686336040518663ffffffff1660e01b815260040161047a959493929190610ae1565b602060405180830381600087803b15801561049457600080fd5b505af11580156104a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104cc9190610875565b9050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638d8c48f782336040518363ffffffff1660e01b815260040161052b929190610a56565b600060405180830381600087803b15801561054557600080fd5b505af1158015610559573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff167ffcf9a0c9dedbfcd1a047374855fc36baaf605bd4f4837802a0cc938ba1b5f30260405160405180910390a25050505050565b6105af61069f565b73ffffffffffffffffffffffffffffffffffffffff166105cd6103ec565b73ffffffffffffffffffffffffffffffffffffffff1614610623576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061a90610ac1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610693576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068a90610aa1565b60405180910390fd5b61069c816106a7565b50565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600061077e61077984610b59565b610b34565b905080838252602082019050828560208602820111156107a1576107a0610c70565b5b60005b858110156107d157816107b788826107f0565b8452602084019350602083019250506001810190506107a4565b5050509392505050565b6000813590506107ea81610d08565b92915050565b6000815190506107ff81610d08565b92915050565b600082601f83011261081a57610819610c6b565b5b815161082a84826020860161076b565b91505092915050565b60008135905061084281610d1f565b92915050565b60006020828403121561085e5761085d610c7a565b5b600061086c848285016107db565b91505092915050565b60006020828403121561088b5761088a610c7a565b5b6000610899848285016107f0565b91505092915050565b6000602082840312156108b8576108b7610c7a565b5b600082015167ffffffffffffffff8111156108d6576108d5610c75565b5b6108e284828501610805565b91505092915050565b6000806000806080858703121561090557610904610c7a565b5b600061091387828801610833565b945050602061092487828801610833565b935050604061093587828801610833565b925050606061094687828801610833565b91505092959194509250565b600061095e838361096a565b60208301905092915050565b61097381610bcf565b82525050565b61098281610bcf565b82525050565b600061099382610b95565b61099d8185610bad565b93506109a883610b85565b8060005b838110156109d95781516109c08882610952565b97506109cb83610ba0565b9250506001810190506109ac565b5085935050505092915050565b60006109f3602683610bbe565b91506109fe82610c90565b604082019050919050565b6000610a16602083610bbe565b9150610a2182610cdf565b602082019050919050565b610a3581610c01565b82525050565b6000602082019050610a506000830184610979565b92915050565b6000604082019050610a6b6000830185610979565b610a786020830184610979565b9392505050565b60006020820190508181036000830152610a998184610988565b905092915050565b60006020820190508181036000830152610aba816109e6565b9050919050565b60006020820190508181036000830152610ada81610a09565b9050919050565b600060a082019050610af66000830188610a2c565b610b036020830187610a2c565b610b106040830186610a2c565b610b1d6060830185610a2c565b610b2a6080830184610979565b9695505050505050565b6000610b3e610b4f565b9050610b4a8282610c0b565b919050565b6000604051905090565b600067ffffffffffffffff821115610b7457610b73610c3c565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610bda82610be1565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b610c1482610c7f565b810181811067ffffffffffffffff82111715610c3357610c32610c3c565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b610d1181610bcf565b8114610d1c57600080fd5b50565b610d2881610c01565b8114610d3357600080fd5b5056fea264697066735822122020707f37da917bba98b6bb5c790f17cf72c0505634ef927a7919e01a4addff8e64736f6c63430008070033",
}

// WebfacingABI is the input ABI used to generate the binding from.
// Deprecated: Use WebfacingMetaData.ABI instead.
var WebfacingABI = WebfacingMetaData.ABI

// WebfacingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WebfacingMetaData.Bin instead.
var WebfacingBin = WebfacingMetaData.Bin

// DeployWebfacing deploys a new Ethereum contract, binding an instance of Webfacing to it.
func DeployWebfacing(auth *bind.TransactOpts, backend bind.ContractBackend, _ledgerAddress common.Address, _cloneFactoryAddress common.Address, _contractManager common.Address, _proxy common.Address) (common.Address, *types.Transaction, *Webfacing, error) {
	parsed, err := WebfacingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WebfacingBin), backend, _ledgerAddress, _cloneFactoryAddress, _contractManager, _proxy)
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
// Solidity: function setCreateRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length) returns()
func (_Webfacing *WebfacingTransactor) SetCreateRentalContract(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Webfacing.contract.Transact(opts, "setCreateRentalContract", _price, _limit, _speed, _length)
}

// SetCreateRentalContract is a paid mutator transaction binding the contract method 0xc938a5d4.
//
// Solidity: function setCreateRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length) returns()
func (_Webfacing *WebfacingSession) SetCreateRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Webfacing.Contract.SetCreateRentalContract(&_Webfacing.TransactOpts, _price, _limit, _speed, _length)
}

// SetCreateRentalContract is a paid mutator transaction binding the contract method 0xc938a5d4.
//
// Solidity: function setCreateRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length) returns()
func (_Webfacing *WebfacingTransactorSession) SetCreateRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Webfacing.Contract.SetCreateRentalContract(&_Webfacing.TransactOpts, _price, _limit, _speed, _length)
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

// WebfacingContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the Webfacing contract.
type WebfacingContractCreatedIterator struct {
	Event *WebfacingContractCreated // Event containing the contract specifics and raw log

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
func (it *WebfacingContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WebfacingContractCreated)
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
		it.Event = new(WebfacingContractCreated)
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
func (it *WebfacingContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WebfacingContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WebfacingContractCreated represents a ContractCreated event raised by the Webfacing contract.
type WebfacingContractCreated struct {
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0xfcf9a0c9dedbfcd1a047374855fc36baaf605bd4f4837802a0cc938ba1b5f302.
//
// Solidity: event contractCreated(address indexed _contract)
func (_Webfacing *WebfacingFilterer) FilterContractCreated(opts *bind.FilterOpts, _contract []common.Address) (*WebfacingContractCreatedIterator, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Webfacing.contract.FilterLogs(opts, "contractCreated", _contractRule)
	if err != nil {
		return nil, err
	}
	return &WebfacingContractCreatedIterator{contract: _Webfacing.contract, event: "contractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0xfcf9a0c9dedbfcd1a047374855fc36baaf605bd4f4837802a0cc938ba1b5f302.
//
// Solidity: event contractCreated(address indexed _contract)
func (_Webfacing *WebfacingFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *WebfacingContractCreated, _contract []common.Address) (event.Subscription, error) {

	var _contractRule []interface{}
	for _, _contractItem := range _contract {
		_contractRule = append(_contractRule, _contractItem)
	}

	logs, sub, err := _Webfacing.contract.WatchLogs(opts, "contractCreated", _contractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WebfacingContractCreated)
				if err := _Webfacing.contract.UnpackLog(event, "contractCreated", log); err != nil {
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

// ParseContractCreated is a log parse operation binding the contract event 0xfcf9a0c9dedbfcd1a047374855fc36baaf605bd4f4837802a0cc938ba1b5f302.
//
// Solidity: event contractCreated(address indexed _contract)
func (_Webfacing *WebfacingFilterer) ParseContractCreated(log types.Log) (*WebfacingContractCreated, error) {
	event := new(WebfacingContractCreated)
	if err := _Webfacing.contract.UnpackLog(event, "contractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
