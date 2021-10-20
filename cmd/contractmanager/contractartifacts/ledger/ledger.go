// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ledger

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

// LedgerMetaData contains all meta data concerning the Ledger contract.
var LedgerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getListOfContractsLedger\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rentalContract\",\"type\":\"address\"}],\"name\":\"setAddContractToStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516104273803806104278339818101604052810190610032919061008e565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610109565b600081519050610088816100f2565b92915050565b6000602082840312156100a4576100a36100ed565b5b60006100b284828501610079565b91505092915050565b60006100c6826100cd565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6100fb816100bb565b811461010657600080fd5b50565b61030f806101186000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063015e9d4b1461003b578063d70f8d6e14610057575b600080fd5b6100556004803603810190610050919061017e565b610075565b005b61005f6100db565b60405161006c9190610230565b60405180910390f35b6000819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6060600080548060200260200160405190810160405280929190818152602001828054801561015f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610115575b5050505050905090565b600081359050610178816102c2565b92915050565b600060208284031215610194576101936102bd565b5b60006101a284828501610169565b91505092915050565b60006101b783836101c3565b60208301905092915050565b6101cc8161028b565b82525050565b60006101dd82610262565b6101e7818561027a565b93506101f283610252565b8060005b8381101561022357815161020a88826101ab565b97506102158361026d565b9250506001810190506101f6565b5085935050505092915050565b6000602082019050818103600083015261024a81846101d2565b905092915050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b60006102968261029d565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6102cb8161028b565b81146102d657600080fd5b5056fea26469706673582212200316938975603c09689925c8083e30d7fcd1d745ec9f9b319fc3b47ba696dc0064736f6c63430008070033",
}

// LedgerABI is the input ABI used to generate the binding from.
// Deprecated: Use LedgerMetaData.ABI instead.
var LedgerABI = LedgerMetaData.ABI

// LedgerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LedgerMetaData.Bin instead.
var LedgerBin = LedgerMetaData.Bin

// DeployLedger deploys a new Ethereum contract, binding an instance of Ledger to it.
func DeployLedger(auth *bind.TransactOpts, backend bind.ContractBackend, _validator common.Address) (common.Address, *types.Transaction, *Ledger, error) {
	parsed, err := LedgerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LedgerBin), backend, _validator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ledger{LedgerCaller: LedgerCaller{contract: contract}, LedgerTransactor: LedgerTransactor{contract: contract}, LedgerFilterer: LedgerFilterer{contract: contract}}, nil
}

// Ledger is an auto generated Go binding around an Ethereum contract.
type Ledger struct {
	LedgerCaller     // Read-only binding to the contract
	LedgerTransactor // Write-only binding to the contract
	LedgerFilterer   // Log filterer for contract events
}

// LedgerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LedgerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LedgerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LedgerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LedgerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LedgerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LedgerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LedgerSession struct {
	Contract     *Ledger           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LedgerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LedgerCallerSession struct {
	Contract *LedgerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LedgerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LedgerTransactorSession struct {
	Contract     *LedgerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LedgerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LedgerRaw struct {
	Contract *Ledger // Generic contract binding to access the raw methods on
}

// LedgerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LedgerCallerRaw struct {
	Contract *LedgerCaller // Generic read-only contract binding to access the raw methods on
}

// LedgerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LedgerTransactorRaw struct {
	Contract *LedgerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLedger creates a new instance of Ledger, bound to a specific deployed contract.
func NewLedger(address common.Address, backend bind.ContractBackend) (*Ledger, error) {
	contract, err := bindLedger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ledger{LedgerCaller: LedgerCaller{contract: contract}, LedgerTransactor: LedgerTransactor{contract: contract}, LedgerFilterer: LedgerFilterer{contract: contract}}, nil
}

// NewLedgerCaller creates a new read-only instance of Ledger, bound to a specific deployed contract.
func NewLedgerCaller(address common.Address, caller bind.ContractCaller) (*LedgerCaller, error) {
	contract, err := bindLedger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LedgerCaller{contract: contract}, nil
}

// NewLedgerTransactor creates a new write-only instance of Ledger, bound to a specific deployed contract.
func NewLedgerTransactor(address common.Address, transactor bind.ContractTransactor) (*LedgerTransactor, error) {
	contract, err := bindLedger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LedgerTransactor{contract: contract}, nil
}

// NewLedgerFilterer creates a new log filterer instance of Ledger, bound to a specific deployed contract.
func NewLedgerFilterer(address common.Address, filterer bind.ContractFilterer) (*LedgerFilterer, error) {
	contract, err := bindLedger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LedgerFilterer{contract: contract}, nil
}

// bindLedger binds a generic wrapper to an already deployed contract.
func bindLedger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LedgerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ledger *LedgerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ledger.Contract.LedgerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ledger *LedgerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ledger.Contract.LedgerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ledger *LedgerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ledger.Contract.LedgerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ledger *LedgerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ledger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ledger *LedgerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ledger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ledger *LedgerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ledger.Contract.contract.Transact(opts, method, params...)
}

// GetListOfContractsLedger is a free data retrieval call binding the contract method 0xd70f8d6e.
//
// Solidity: function getListOfContractsLedger() view returns(address[])
func (_Ledger *LedgerCaller) GetListOfContractsLedger(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Ledger.contract.Call(opts, &out, "getListOfContractsLedger")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetListOfContractsLedger is a free data retrieval call binding the contract method 0xd70f8d6e.
//
// Solidity: function getListOfContractsLedger() view returns(address[])
func (_Ledger *LedgerSession) GetListOfContractsLedger() ([]common.Address, error) {
	return _Ledger.Contract.GetListOfContractsLedger(&_Ledger.CallOpts)
}

// GetListOfContractsLedger is a free data retrieval call binding the contract method 0xd70f8d6e.
//
// Solidity: function getListOfContractsLedger() view returns(address[])
func (_Ledger *LedgerCallerSession) GetListOfContractsLedger() ([]common.Address, error) {
	return _Ledger.Contract.GetListOfContractsLedger(&_Ledger.CallOpts)
}

// SetAddContractToStorage is a paid mutator transaction binding the contract method 0x015e9d4b.
//
// Solidity: function setAddContractToStorage(address _rentalContract) returns()
func (_Ledger *LedgerTransactor) SetAddContractToStorage(opts *bind.TransactOpts, _rentalContract common.Address) (*types.Transaction, error) {
	return _Ledger.contract.Transact(opts, "setAddContractToStorage", _rentalContract)
}

// SetAddContractToStorage is a paid mutator transaction binding the contract method 0x015e9d4b.
//
// Solidity: function setAddContractToStorage(address _rentalContract) returns()
func (_Ledger *LedgerSession) SetAddContractToStorage(_rentalContract common.Address) (*types.Transaction, error) {
	return _Ledger.Contract.SetAddContractToStorage(&_Ledger.TransactOpts, _rentalContract)
}

// SetAddContractToStorage is a paid mutator transaction binding the contract method 0x015e9d4b.
//
// Solidity: function setAddContractToStorage(address _rentalContract) returns()
func (_Ledger *LedgerTransactorSession) SetAddContractToStorage(_rentalContract common.Address) (*types.Transaction, error) {
	return _Ledger.Contract.SetAddContractToStorage(&_Ledger.TransactOpts, _rentalContract)
}
