// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package implementation

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

// ImplementationMetaData contains all meta data concerning the Implementation contract.
var ImplementationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"contractCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"contractClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"contractPurchased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractState\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dueAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_purchaser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMiningPoolInformation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validationFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lmn\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"port\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receivedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setContractCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFundContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipaddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withValidator\",\"type\":\"bool\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"speed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611ee4806100206000396000f3fe60806040526004361061011f5760003560e01c806374a34812116100a0578063c204296511610064578063c204296514610372578063c20906ac1461039d578063c5095d68146103c8578063ce0c722a146103f3578063f8bbf27e1461041e5761011f565b806374a34812146102b157806385209ee0146102da578063a035b1fe14610305578063a4d66daf14610330578063b6293b341461035b5761011f565b80631f7b6d32116100e75780631f7b6d32146101ec57806326b3c68b146102175780634897b9ac1461024257806349823485146102595780637150d8ae146102865761011f565b806308551a5314610124578063089aa8a21461014f5780630a61e2d91461017a57806311afdfc9146101a557806316713b37146101c1575b600080fd5b34801561013057600080fd5b50610139610449565b60405161014691906117ee565b60405180910390f35b34801561015b57600080fd5b5061016461046f565b60405161017191906117ee565b60405180910390f35b34801561018657600080fd5b5061018f610495565b60405161019c9190611959565b60405180910390f35b6101bf60048036038101906101ba91906114f2565b61049b565b005b3480156101cd57600080fd5b506101d66107e8565b6040516101e39190611959565b60405180910390f35b3480156101f857600080fd5b506102016107ee565b60405161020e9190611959565b60405180910390f35b34801561022357600080fd5b5061022c6107f4565b6040516102399190611959565b60405180910390f35b34801561024e57600080fd5b506102576107fa565b005b34801561026557600080fd5b5061026e610955565b60405161027d9392919061184d565b60405180910390f35b34801561029257600080fd5b5061029b610ba2565b6040516102a891906117ee565b60405180910390f35b3480156102bd57600080fd5b506102d860048036038101906102d39190611600565b610bc8565b005b3480156102e657600080fd5b506102ef610da5565b6040516102fc9190611832565b60405180910390f35b34801561031157600080fd5b5061031a610db8565b6040516103279190611959565b60405180910390f35b34801561033c57600080fd5b50610345610dbe565b6040516103529190611959565b60405180910390f35b34801561036757600080fd5b50610370610dc4565b005b34801561037e57600080fd5b50610387610f74565b6040516103949190611959565b60405180910390f35b3480156103a957600080fd5b506103b2611033565b6040516103bf9190611959565b60405180910390f35b3480156103d457600080fd5b506103dd611039565b6040516103ea9190611959565b60405180910390f35b3480156103ff57600080fd5b5061040861103f565b60405161041591906117ee565b60405180910390f35b34801561042a57600080fd5b50610433611065565b6040516104409190611959565b60405180910390f35b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b600060038111156104af576104ae611c15565b5b600560149054906101000a900460ff1660038111156104d1576104d0611c15565b5b14610511576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610508906118d9565b60405180910390fd5b8560119080519060200190610527929190611349565b50846012908051906020019061053e929190611349565b508360139080519060200190610555929190611349565b5082600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610629600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660065461106b565b6001151581151514156106d6576001601460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550606434146106d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c890611919565b60405180910390fd5b6107a9565b6001601460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160146000600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b7f0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb336040516107d891906117ee565b60405180910390a1505050505050565b60045481565b60095481565b600a5481565b60011515601460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461088d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088490611899565b60405180910390fd5b6000600c544261089d9190611a71565b905060008060095483106108b457600091506108df565b600954836009546108c59190611a71565b6006546108d29190611a17565b6108dc91906119e6565b91505b816006546108ed9190611a71565b90506108f981836110f9565b6003600560146101000a81548160ff0219169083600381111561091f5761091e611c15565b5b02179055507ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a1505050565b606080606060011515601460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515146109ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e490611899565b60405180910390fd5b6011601260138280546109ff90611b54565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2b90611b54565b8015610a785780601f10610a4d57610100808354040283529160200191610a78565b820191906000526020600020905b815481529060010190602001808311610a5b57829003601f168201915b50505050509250818054610a8b90611b54565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab790611b54565b8015610b045780601f10610ad957610100808354040283529160200191610b04565b820191906000526020600020905b815481529060010190602001808311610ae757829003601f168201915b50505050509150808054610b1790611b54565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4390611b54565b8015610b905780601f10610b6557610100808354040283529160200191610b90565b820191906000526020600020905b815481529060010190602001808311610b7357829003601f168201915b50505050509050925092509250909192565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900460ff1680610bee575060008054906101000a900460ff16155b610c2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c24906118f9565b60405180910390fd5b60008060019054906101000a900460ff161590508015610c7d576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b8860068190555087600781905550866008819055508560098190555083600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600b8190555082600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560146101000a81548160ff02191690836003811115610d4857610d47611c15565b5b0217905550610d79600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836112a1565b8015610d9a5760008060016101000a81548160ff0219169083151502179055505b505050505050505050565b600560149054906101000a900460ff1681565b60065481565b60075481565b6000610dce610f74565b14610e0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e05906118b9565b60405180910390fd5b6001151560146000600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151480610f0157506001151560146000600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515145b610f40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3790611939565b60405180910390fd5b42600c819055506002600560146101000a81548160ff02191690836003811115610f6d57610f6c611c15565b5b0217905550565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610fd191906117ee565b60206040518083038186803b158015610fe957600080fd5b505afa158015610ffd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061102191906115d3565b60035461102e9190611a71565b905090565b60085481565b600c5481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b5481565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600381905550505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b8152600401611178929190611809565b602060405180830381600087803b15801561119257600080fd5b505af11580156111a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ca91906114c5565b50600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161124a929190611809565b602060405180830381600087803b15801561126457600080fd5b505af1158015611278573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129c91906114c5565b505050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b82805461135590611b54565b90600052602060002090601f01602090048101928261137757600085556113be565b82601f1061139057805160ff19168380011785556113be565b828001600101855582156113be579182015b828111156113bd5782518255916020019190600101906113a2565b5b5090506113cb91906113cf565b5090565b5b808211156113e85760008160009055506001016113d0565b5090565b60006113ff6113fa84611999565b611974565b90508281526020810184848401111561141b5761141a611ca7565b5b611426848285611b12565b509392505050565b60008135905061143d81611e69565b92915050565b60008135905061145281611e80565b92915050565b60008151905061146781611e80565b92915050565b600082601f83011261148257611481611ca2565b5b81356114928482602086016113ec565b91505092915050565b6000813590506114aa81611e97565b92915050565b6000815190506114bf81611e97565b92915050565b6000602082840312156114db576114da611cb1565b5b60006114e984828501611458565b91505092915050565b60008060008060008060c0878903121561150f5761150e611cb1565b5b600087013567ffffffffffffffff81111561152d5761152c611cac565b5b61153989828a0161146d565b965050602087013567ffffffffffffffff81111561155a57611559611cac565b5b61156689828a0161146d565b955050604087013567ffffffffffffffff81111561158757611586611cac565b5b61159389828a0161146d565b94505060606115a489828a0161142e565b93505060806115b589828a0161142e565b92505060a06115c689828a01611443565b9150509295509295509295565b6000602082840312156115e9576115e8611cb1565b5b60006115f7848285016114b0565b91505092915050565b600080600080600080600080610100898b03121561162157611620611cb1565b5b600061162f8b828c0161149b565b98505060206116408b828c0161149b565b97505060406116518b828c0161149b565b96505060606116628b828c0161149b565b95505060806116738b828c0161149b565b94505060a06116848b828c0161142e565b93505060c06116958b828c0161142e565b92505060e06116a68b828c0161142e565b9150509295985092959890939650565b6116bf81611aa5565b82525050565b6116ce81611b00565b82525050565b60006116df826119ca565b6116e981856119d5565b93506116f9818560208601611b21565b61170281611cb6565b840191505092915050565b600061171a6031836119d5565b915061172582611cc7565b604082019050919050565b600061173d602e836119d5565b915061174882611d16565b604082019050919050565b60006117606025836119d5565b915061176b82611d65565b604082019050919050565b6000611783602e836119d5565b915061178e82611db4565b604082019050919050565b60006117a66017836119d5565b91506117b182611e03565b602082019050919050565b60006117c9601a836119d5565b91506117d482611e2c565b602082019050919050565b6117e881611af6565b82525050565b600060208201905061180360008301846116b6565b92915050565b600060408201905061181e60008301856116b6565b61182b60208301846117df565b9392505050565b600060208201905061184760008301846116c5565b92915050565b6000606082019050818103600083015261186781866116d4565b9050818103602083015261187b81856116d4565b9050818103604083015261188f81846116d4565b9050949350505050565b600060208201905081810360008301526118b28161170d565b9050919050565b600060208201905081810360008301526118d281611730565b9050919050565b600060208201905081810360008301526118f281611753565b9050919050565b6000602082019050818103600083015261191281611776565b9050919050565b6000602082019050818103600083015261193281611799565b9050919050565b60006020820190508181036000830152611952816117bc565b9050919050565b600060208201905061196e60008301846117df565b92915050565b600061197e61198f565b905061198a8282611b86565b919050565b6000604051905090565b600067ffffffffffffffff8211156119b4576119b3611c73565b5b6119bd82611cb6565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60006119f182611af6565b91506119fc83611af6565b925082611a0c57611a0b611be6565b5b828204905092915050565b6000611a2282611af6565b9150611a2d83611af6565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611a6657611a65611bb7565b5b828202905092915050565b6000611a7c82611af6565b9150611a8783611af6565b925082821015611a9a57611a99611bb7565b5b828203905092915050565b6000611ab082611ad6565b9050919050565b60008115159050919050565b6000819050611ad182611e55565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611b0b82611ac3565b9050919050565b82818337600083830152505050565b60005b83811015611b3f578082015181840152602081019050611b24565b83811115611b4e576000848401525b50505050565b60006002820490506001821680611b6c57607f821691505b60208210811415611b8057611b7f611c44565b5b50919050565b611b8f82611cb6565b810181811067ffffffffffffffff82111715611bae57611bad611c73565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f746869732061646472657373206973206e6f7420616c6c6f77656420746f206360008201527f616c6c20746869732066756e6374696f6e000000000000000000000000000000602082015250565b7f6c756d6572696e20746f6b656e73206e65656420746f2062652073656e74207460008201527f6f2074686520636f6e7472616374000000000000000000000000000000000000602082015250565b7f636f6e7472616374206973206e6f7420696e20616e20617661696c61626c652060008201527f7374617465000000000000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f76616c69646174696f6e20666565206e6f742073656e74000000000000000000600082015250565b7f74686520627579657220686173206e6f74206265656e20736574000000000000600082015250565b60048110611e6657611e65611c15565b5b50565b611e7281611aa5565b8114611e7d57600080fd5b50565b611e8981611ab7565b8114611e9457600080fd5b50565b611ea081611af6565b8114611eab57600080fd5b5056fea2646970667358221220ecbf1beda227ec85143d1321c00692a94a3a315e0c63de24f7e15e90f8caa82264736f6c63430008070033",
}

// ImplementationABI is the input ABI used to generate the binding from.
// Deprecated: Use ImplementationMetaData.ABI instead.
var ImplementationABI = ImplementationMetaData.ABI

// ImplementationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ImplementationMetaData.Bin instead.
var ImplementationBin = ImplementationMetaData.Bin

// DeployImplementation deploys a new Ethereum contract, binding an instance of Implementation to it.
func DeployImplementation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Implementation, error) {
	parsed, err := ImplementationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ImplementationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Implementation{ImplementationCaller: ImplementationCaller{contract: contract}, ImplementationTransactor: ImplementationTransactor{contract: contract}, ImplementationFilterer: ImplementationFilterer{contract: contract}}, nil
}

// Implementation is an auto generated Go binding around an Ethereum contract.
type Implementation struct {
	ImplementationCaller     // Read-only binding to the contract
	ImplementationTransactor // Write-only binding to the contract
	ImplementationFilterer   // Log filterer for contract events
}

// ImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImplementationSession struct {
	Contract     *Implementation   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImplementationCallerSession struct {
	Contract *ImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImplementationTransactorSession struct {
	Contract     *ImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImplementationRaw struct {
	Contract *Implementation // Generic contract binding to access the raw methods on
}

// ImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImplementationCallerRaw struct {
	Contract *ImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// ImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImplementationTransactorRaw struct {
	Contract *ImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImplementation creates a new instance of Implementation, bound to a specific deployed contract.
func NewImplementation(address common.Address, backend bind.ContractBackend) (*Implementation, error) {
	contract, err := bindImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Implementation{ImplementationCaller: ImplementationCaller{contract: contract}, ImplementationTransactor: ImplementationTransactor{contract: contract}, ImplementationFilterer: ImplementationFilterer{contract: contract}}, nil
}

// NewImplementationCaller creates a new read-only instance of Implementation, bound to a specific deployed contract.
func NewImplementationCaller(address common.Address, caller bind.ContractCaller) (*ImplementationCaller, error) {
	contract, err := bindImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImplementationCaller{contract: contract}, nil
}

// NewImplementationTransactor creates a new write-only instance of Implementation, bound to a specific deployed contract.
func NewImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*ImplementationTransactor, error) {
	contract, err := bindImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImplementationTransactor{contract: contract}, nil
}

// NewImplementationFilterer creates a new log filterer instance of Implementation, bound to a specific deployed contract.
func NewImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*ImplementationFilterer, error) {
	contract, err := bindImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImplementationFilterer{contract: contract}, nil
}

// bindImplementation binds a generic wrapper to an already deployed contract.
func bindImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ImplementationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Implementation *ImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Implementation.Contract.ImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Implementation *ImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.Contract.ImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Implementation *ImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Implementation.Contract.ImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Implementation *ImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Implementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Implementation *ImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Implementation *ImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Implementation.Contract.contract.Transact(opts, method, params...)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Implementation *ImplementationCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "buyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Implementation *ImplementationSession) Buyer() (common.Address, error) {
	return _Implementation.Contract.Buyer(&_Implementation.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Implementation *ImplementationCallerSession) Buyer() (common.Address, error) {
	return _Implementation.Contract.Buyer(&_Implementation.CallOpts)
}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Implementation *ImplementationCaller) ContractState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "contractState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Implementation *ImplementationSession) ContractState() (uint8, error) {
	return _Implementation.Contract.ContractState(&_Implementation.CallOpts)
}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Implementation *ImplementationCallerSession) ContractState() (uint8, error) {
	return _Implementation.Contract.ContractState(&_Implementation.CallOpts)
}

// ContractTotal is a free data retrieval call binding the contract method 0x0a61e2d9.
//
// Solidity: function contractTotal() view returns(uint256)
func (_Implementation *ImplementationCaller) ContractTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "contractTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractTotal is a free data retrieval call binding the contract method 0x0a61e2d9.
//
// Solidity: function contractTotal() view returns(uint256)
func (_Implementation *ImplementationSession) ContractTotal() (*big.Int, error) {
	return _Implementation.Contract.ContractTotal(&_Implementation.CallOpts)
}

// ContractTotal is a free data retrieval call binding the contract method 0x0a61e2d9.
//
// Solidity: function contractTotal() view returns(uint256)
func (_Implementation *ImplementationCallerSession) ContractTotal() (*big.Int, error) {
	return _Implementation.Contract.ContractTotal(&_Implementation.CallOpts)
}

// DueAmount is a free data retrieval call binding the contract method 0xc2042965.
//
// Solidity: function dueAmount() view returns(uint256)
func (_Implementation *ImplementationCaller) DueAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "dueAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DueAmount is a free data retrieval call binding the contract method 0xc2042965.
//
// Solidity: function dueAmount() view returns(uint256)
func (_Implementation *ImplementationSession) DueAmount() (*big.Int, error) {
	return _Implementation.Contract.DueAmount(&_Implementation.CallOpts)
}

// DueAmount is a free data retrieval call binding the contract method 0xc2042965.
//
// Solidity: function dueAmount() view returns(uint256)
func (_Implementation *ImplementationCallerSession) DueAmount() (*big.Int, error) {
	return _Implementation.Contract.DueAmount(&_Implementation.CallOpts)
}

// EscrowPurchaser is a free data retrieval call binding the contract method 0x089aa8a2.
//
// Solidity: function escrow_purchaser() view returns(address)
func (_Implementation *ImplementationCaller) EscrowPurchaser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "escrow_purchaser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EscrowPurchaser is a free data retrieval call binding the contract method 0x089aa8a2.
//
// Solidity: function escrow_purchaser() view returns(address)
func (_Implementation *ImplementationSession) EscrowPurchaser() (common.Address, error) {
	return _Implementation.Contract.EscrowPurchaser(&_Implementation.CallOpts)
}

// EscrowPurchaser is a free data retrieval call binding the contract method 0x089aa8a2.
//
// Solidity: function escrow_purchaser() view returns(address)
func (_Implementation *ImplementationCallerSession) EscrowPurchaser() (common.Address, error) {
	return _Implementation.Contract.EscrowPurchaser(&_Implementation.CallOpts)
}

// EscrowSeller is a free data retrieval call binding the contract method 0xce0c722a.
//
// Solidity: function escrow_seller() view returns(address)
func (_Implementation *ImplementationCaller) EscrowSeller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "escrow_seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EscrowSeller is a free data retrieval call binding the contract method 0xce0c722a.
//
// Solidity: function escrow_seller() view returns(address)
func (_Implementation *ImplementationSession) EscrowSeller() (common.Address, error) {
	return _Implementation.Contract.EscrowSeller(&_Implementation.CallOpts)
}

// EscrowSeller is a free data retrieval call binding the contract method 0xce0c722a.
//
// Solidity: function escrow_seller() view returns(address)
func (_Implementation *ImplementationCallerSession) EscrowSeller() (common.Address, error) {
	return _Implementation.Contract.EscrowSeller(&_Implementation.CallOpts)
}

// GetMiningPoolInformation is a free data retrieval call binding the contract method 0x49823485.
//
// Solidity: function getMiningPoolInformation() view returns(string, string, string)
func (_Implementation *ImplementationCaller) GetMiningPoolInformation(opts *bind.CallOpts) (string, string, string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "getMiningPoolInformation")

	if err != nil {
		return *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetMiningPoolInformation is a free data retrieval call binding the contract method 0x49823485.
//
// Solidity: function getMiningPoolInformation() view returns(string, string, string)
func (_Implementation *ImplementationSession) GetMiningPoolInformation() (string, string, string, error) {
	return _Implementation.Contract.GetMiningPoolInformation(&_Implementation.CallOpts)
}

// GetMiningPoolInformation is a free data retrieval call binding the contract method 0x49823485.
//
// Solidity: function getMiningPoolInformation() view returns(string, string, string)
func (_Implementation *ImplementationCallerSession) GetMiningPoolInformation() (string, string, string, error) {
	return _Implementation.Contract.GetMiningPoolInformation(&_Implementation.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_Implementation *ImplementationCaller) Length(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_Implementation *ImplementationSession) Length() (*big.Int, error) {
	return _Implementation.Contract.Length(&_Implementation.CallOpts)
}

// Length is a free data retrieval call binding the contract method 0x1f7b6d32.
//
// Solidity: function length() view returns(uint256)
func (_Implementation *ImplementationCallerSession) Length() (*big.Int, error) {
	return _Implementation.Contract.Length(&_Implementation.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Implementation *ImplementationCaller) Limit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Implementation *ImplementationSession) Limit() (*big.Int, error) {
	return _Implementation.Contract.Limit(&_Implementation.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() view returns(uint256)
func (_Implementation *ImplementationCallerSession) Limit() (*big.Int, error) {
	return _Implementation.Contract.Limit(&_Implementation.CallOpts)
}

// Port is a free data retrieval call binding the contract method 0x26b3c68b.
//
// Solidity: function port() view returns(uint256)
func (_Implementation *ImplementationCaller) Port(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "port")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Port is a free data retrieval call binding the contract method 0x26b3c68b.
//
// Solidity: function port() view returns(uint256)
func (_Implementation *ImplementationSession) Port() (*big.Int, error) {
	return _Implementation.Contract.Port(&_Implementation.CallOpts)
}

// Port is a free data retrieval call binding the contract method 0x26b3c68b.
//
// Solidity: function port() view returns(uint256)
func (_Implementation *ImplementationCallerSession) Port() (*big.Int, error) {
	return _Implementation.Contract.Port(&_Implementation.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Implementation *ImplementationCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Implementation *ImplementationSession) Price() (*big.Int, error) {
	return _Implementation.Contract.Price(&_Implementation.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_Implementation *ImplementationCallerSession) Price() (*big.Int, error) {
	return _Implementation.Contract.Price(&_Implementation.CallOpts)
}

// ReceivedTotal is a free data retrieval call binding the contract method 0x16713b37.
//
// Solidity: function receivedTotal() view returns(uint256)
func (_Implementation *ImplementationCaller) ReceivedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "receivedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReceivedTotal is a free data retrieval call binding the contract method 0x16713b37.
//
// Solidity: function receivedTotal() view returns(uint256)
func (_Implementation *ImplementationSession) ReceivedTotal() (*big.Int, error) {
	return _Implementation.Contract.ReceivedTotal(&_Implementation.CallOpts)
}

// ReceivedTotal is a free data retrieval call binding the contract method 0x16713b37.
//
// Solidity: function receivedTotal() view returns(uint256)
func (_Implementation *ImplementationCallerSession) ReceivedTotal() (*big.Int, error) {
	return _Implementation.Contract.ReceivedTotal(&_Implementation.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Implementation *ImplementationCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Implementation *ImplementationSession) Seller() (common.Address, error) {
	return _Implementation.Contract.Seller(&_Implementation.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Implementation *ImplementationCallerSession) Seller() (common.Address, error) {
	return _Implementation.Contract.Seller(&_Implementation.CallOpts)
}

// Speed is a free data retrieval call binding the contract method 0xc20906ac.
//
// Solidity: function speed() view returns(uint256)
func (_Implementation *ImplementationCaller) Speed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "speed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Speed is a free data retrieval call binding the contract method 0xc20906ac.
//
// Solidity: function speed() view returns(uint256)
func (_Implementation *ImplementationSession) Speed() (*big.Int, error) {
	return _Implementation.Contract.Speed(&_Implementation.CallOpts)
}

// Speed is a free data retrieval call binding the contract method 0xc20906ac.
//
// Solidity: function speed() view returns(uint256)
func (_Implementation *ImplementationCallerSession) Speed() (*big.Int, error) {
	return _Implementation.Contract.Speed(&_Implementation.CallOpts)
}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_Implementation *ImplementationCaller) StartingBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "startingBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_Implementation *ImplementationSession) StartingBlockTimestamp() (*big.Int, error) {
	return _Implementation.Contract.StartingBlockTimestamp(&_Implementation.CallOpts)
}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_Implementation *ImplementationCallerSession) StartingBlockTimestamp() (*big.Int, error) {
	return _Implementation.Contract.StartingBlockTimestamp(&_Implementation.CallOpts)
}

// ValidationFee is a free data retrieval call binding the contract method 0xf8bbf27e.
//
// Solidity: function validationFee() view returns(uint256)
func (_Implementation *ImplementationCaller) ValidationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "validationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidationFee is a free data retrieval call binding the contract method 0xf8bbf27e.
//
// Solidity: function validationFee() view returns(uint256)
func (_Implementation *ImplementationSession) ValidationFee() (*big.Int, error) {
	return _Implementation.Contract.ValidationFee(&_Implementation.CallOpts)
}

// ValidationFee is a free data retrieval call binding the contract method 0xf8bbf27e.
//
// Solidity: function validationFee() view returns(uint256)
func (_Implementation *ImplementationCallerSession) ValidationFee() (*big.Int, error) {
	return _Implementation.Contract.ValidationFee(&_Implementation.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x74a34812.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _validationFee, address _seller, address _contractManager, address _lmn) returns()
func (_Implementation *ImplementationTransactor) Initialize(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validationFee *big.Int, _seller common.Address, _contractManager common.Address, _lmn common.Address) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "initialize", _price, _limit, _speed, _length, _validationFee, _seller, _contractManager, _lmn)
}

// Initialize is a paid mutator transaction binding the contract method 0x74a34812.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _validationFee, address _seller, address _contractManager, address _lmn) returns()
func (_Implementation *ImplementationSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validationFee *big.Int, _seller common.Address, _contractManager common.Address, _lmn common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _validationFee, _seller, _contractManager, _lmn)
}

// Initialize is a paid mutator transaction binding the contract method 0x74a34812.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _validationFee, address _seller, address _contractManager, address _lmn) returns()
func (_Implementation *ImplementationTransactorSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validationFee *big.Int, _seller common.Address, _contractManager common.Address, _lmn common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _validationFee, _seller, _contractManager, _lmn)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x4897b9ac.
//
// Solidity: function setContractCloseOut() returns()
func (_Implementation *ImplementationTransactor) SetContractCloseOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setContractCloseOut")
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x4897b9ac.
//
// Solidity: function setContractCloseOut() returns()
func (_Implementation *ImplementationSession) SetContractCloseOut() (*types.Transaction, error) {
	return _Implementation.Contract.SetContractCloseOut(&_Implementation.TransactOpts)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x4897b9ac.
//
// Solidity: function setContractCloseOut() returns()
func (_Implementation *ImplementationTransactorSession) SetContractCloseOut() (*types.Transaction, error) {
	return _Implementation.Contract.SetContractCloseOut(&_Implementation.TransactOpts)
}

// SetFundContract is a paid mutator transaction binding the contract method 0xb6293b34.
//
// Solidity: function setFundContract() returns()
func (_Implementation *ImplementationTransactor) SetFundContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setFundContract")
}

// SetFundContract is a paid mutator transaction binding the contract method 0xb6293b34.
//
// Solidity: function setFundContract() returns()
func (_Implementation *ImplementationSession) SetFundContract() (*types.Transaction, error) {
	return _Implementation.Contract.SetFundContract(&_Implementation.TransactOpts)
}

// SetFundContract is a paid mutator transaction binding the contract method 0xb6293b34.
//
// Solidity: function setFundContract() returns()
func (_Implementation *ImplementationTransactorSession) SetFundContract() (*types.Transaction, error) {
	return _Implementation.Contract.SetFundContract(&_Implementation.TransactOpts)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x11afdfc9.
//
// Solidity: function setPurchaseContract(string _ipaddress, string _username, string _password, address _buyer, address _validator, bool _withValidator) payable returns()
func (_Implementation *ImplementationTransactor) SetPurchaseContract(opts *bind.TransactOpts, _ipaddress string, _username string, _password string, _buyer common.Address, _validator common.Address, _withValidator bool) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setPurchaseContract", _ipaddress, _username, _password, _buyer, _validator, _withValidator)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x11afdfc9.
//
// Solidity: function setPurchaseContract(string _ipaddress, string _username, string _password, address _buyer, address _validator, bool _withValidator) payable returns()
func (_Implementation *ImplementationSession) SetPurchaseContract(_ipaddress string, _username string, _password string, _buyer common.Address, _validator common.Address, _withValidator bool) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _ipaddress, _username, _password, _buyer, _validator, _withValidator)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x11afdfc9.
//
// Solidity: function setPurchaseContract(string _ipaddress, string _username, string _password, address _buyer, address _validator, bool _withValidator) payable returns()
func (_Implementation *ImplementationTransactorSession) SetPurchaseContract(_ipaddress string, _username string, _password string, _buyer common.Address, _validator common.Address, _withValidator bool) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _ipaddress, _username, _password, _buyer, _validator, _withValidator)
}

// ImplementationContractCanceledIterator is returned from FilterContractCanceled and is used to iterate over the raw logs and unpacked data for ContractCanceled events raised by the Implementation contract.
type ImplementationContractCanceledIterator struct {
	Event *ImplementationContractCanceled // Event containing the contract specifics and raw log

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
func (it *ImplementationContractCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractCanceled)
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
		it.Event = new(ImplementationContractCanceled)
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
func (it *ImplementationContractCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractCanceled represents a ContractCanceled event raised by the Implementation contract.
type ImplementationContractCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractCanceled is a free log retrieval operation binding the contract event 0xcb3fbf71b2999f61d06483d265767337233ac476594db6b00a71ffa8306a1cb7.
//
// Solidity: event contractCanceled()
func (_Implementation *ImplementationFilterer) FilterContractCanceled(opts *bind.FilterOpts) (*ImplementationContractCanceledIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractCanceled")
	if err != nil {
		return nil, err
	}
	return &ImplementationContractCanceledIterator{contract: _Implementation.contract, event: "contractCanceled", logs: logs, sub: sub}, nil
}

// WatchContractCanceled is a free log subscription operation binding the contract event 0xcb3fbf71b2999f61d06483d265767337233ac476594db6b00a71ffa8306a1cb7.
//
// Solidity: event contractCanceled()
func (_Implementation *ImplementationFilterer) WatchContractCanceled(opts *bind.WatchOpts, sink chan<- *ImplementationContractCanceled) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractCanceled)
				if err := _Implementation.contract.UnpackLog(event, "contractCanceled", log); err != nil {
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

// ParseContractCanceled is a log parse operation binding the contract event 0xcb3fbf71b2999f61d06483d265767337233ac476594db6b00a71ffa8306a1cb7.
//
// Solidity: event contractCanceled()
func (_Implementation *ImplementationFilterer) ParseContractCanceled(log types.Log) (*ImplementationContractCanceled, error) {
	event := new(ImplementationContractCanceled)
	if err := _Implementation.contract.UnpackLog(event, "contractCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationContractClosedIterator is returned from FilterContractClosed and is used to iterate over the raw logs and unpacked data for ContractClosed events raised by the Implementation contract.
type ImplementationContractClosedIterator struct {
	Event *ImplementationContractClosed // Event containing the contract specifics and raw log

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
func (it *ImplementationContractClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractClosed)
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
		it.Event = new(ImplementationContractClosed)
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
func (it *ImplementationContractClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractClosed represents a ContractClosed event raised by the Implementation contract.
type ImplementationContractClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterContractClosed is a free log retrieval operation binding the contract event 0xf5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f4844.
//
// Solidity: event contractClosed()
func (_Implementation *ImplementationFilterer) FilterContractClosed(opts *bind.FilterOpts) (*ImplementationContractClosedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractClosed")
	if err != nil {
		return nil, err
	}
	return &ImplementationContractClosedIterator{contract: _Implementation.contract, event: "contractClosed", logs: logs, sub: sub}, nil
}

// WatchContractClosed is a free log subscription operation binding the contract event 0xf5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f4844.
//
// Solidity: event contractClosed()
func (_Implementation *ImplementationFilterer) WatchContractClosed(opts *bind.WatchOpts, sink chan<- *ImplementationContractClosed) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractClosed)
				if err := _Implementation.contract.UnpackLog(event, "contractClosed", log); err != nil {
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

// ParseContractClosed is a log parse operation binding the contract event 0xf5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f4844.
//
// Solidity: event contractClosed()
func (_Implementation *ImplementationFilterer) ParseContractClosed(log types.Log) (*ImplementationContractClosed, error) {
	event := new(ImplementationContractClosed)
	if err := _Implementation.contract.UnpackLog(event, "contractClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationContractPurchasedIterator is returned from FilterContractPurchased and is used to iterate over the raw logs and unpacked data for ContractPurchased events raised by the Implementation contract.
type ImplementationContractPurchasedIterator struct {
	Event *ImplementationContractPurchased // Event containing the contract specifics and raw log

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
func (it *ImplementationContractPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractPurchased)
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
		it.Event = new(ImplementationContractPurchased)
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
func (it *ImplementationContractPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractPurchased represents a ContractPurchased event raised by the Implementation contract.
type ImplementationContractPurchased struct {
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractPurchased is a free log retrieval operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address _buyer)
func (_Implementation *ImplementationFilterer) FilterContractPurchased(opts *bind.FilterOpts) (*ImplementationContractPurchasedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractPurchased")
	if err != nil {
		return nil, err
	}
	return &ImplementationContractPurchasedIterator{contract: _Implementation.contract, event: "contractPurchased", logs: logs, sub: sub}, nil
}

// WatchContractPurchased is a free log subscription operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address _buyer)
func (_Implementation *ImplementationFilterer) WatchContractPurchased(opts *bind.WatchOpts, sink chan<- *ImplementationContractPurchased) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractPurchased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractPurchased)
				if err := _Implementation.contract.UnpackLog(event, "contractPurchased", log); err != nil {
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

// ParseContractPurchased is a log parse operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address _buyer)
func (_Implementation *ImplementationFilterer) ParseContractPurchased(log types.Log) (*ImplementationContractPurchased, error) {
	event := new(ImplementationContractPurchased)
	if err := _Implementation.contract.UnpackLog(event, "contractPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
