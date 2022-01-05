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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newCipherText\",\"type\":\"string\"}],\"name\":\"cipherTextUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"contractClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"contractPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"purchaseInfoUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractState\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"encryptedPoolData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_purchaser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cloneFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receivedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"closeOutType\",\"type\":\"uint256\"}],\"name\":\"setContractCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_encryptedPoolData\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newEncryptedPoolData\",\"type\":\"string\"}],\"name\":\"setUpdateMiningInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_closeoutType\",\"type\":\"uint256\"}],\"name\":\"setUpdatePurchaseInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"speed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612436806100206000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80638b7e4b13116100a2578063c20906ac11610071578063c20906ac1461029b578063c5095d68146102b9578063ce0c722a146102d7578063ddcb1bf2146102f5578063f1a6e3011461031157610116565b80638b7e4b13146102255780638e2e6d5d14610243578063a035b1fe1461025f578063a4d66daf1461027d57610116565b80631f7b6d32116100e95780631f7b6d32146101935780633a5d4213146101b15780637150d8ae146101cd578063719e6b5b146101eb57806385209ee01461020757610116565b806308551a531461011b578063089aa8a2146101395780630a61e2d91461015757806316713b3714610175575b600080fd5b61012361032d565b6040516101309190611a12565b60405180910390f35b610141610353565b60405161014e9190611a12565b60405180910390f35b61015f610379565b60405161016c9190611c13565b60405180910390f35b61017d61037f565b60405161018a9190611c13565b60405180910390f35b61019b610385565b6040516101a89190611c13565b60405180910390f35b6101cb60048036038101906101c691906116d7565b61038b565b005b6101d561057f565b6040516101e29190611a12565b60405180910390f35b610205600480360381019061020091906115d8565b6105a5565b005b61020f6106fb565b60405161021c9190611a56565b60405180910390f35b61022d61070e565b60405161023a9190611a71565b60405180910390f35b61025d6004803603810190610258919061167d565b61079c565b005b610267610b16565b6040516102749190611c13565b60405180910390f35b610285610b1c565b6040516102929190611c13565b60405180910390f35b6102a3610b22565b6040516102b09190611c13565b60405180910390f35b6102c1610b28565b6040516102ce9190611c13565b60405180910390f35b6102df610b2e565b6040516102ec9190611a12565b60405180910390f35b61030f600480360381019061030a9190611621565b610b54565b005b61032b6004803603810190610326919061178d565b610d7c565b005b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60045481565b60095481565b600060019054906101000a900460ff16806103b1575060008054906101000a900460ff16155b6103f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e790611b93565b60405180910390fd5b60008060019054906101000a900460ff161590508015610440576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b8860068190555087600781905550866008819055508560098190555084600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560146101000a81548160ff0219169083600181111561054557610544611ecf565b5b021790555061055384610f27565b80156105745760008060016101000a81548160ff0219169083151502179055505b505050505050505050565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610635576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062c90611a93565b60405180910390fd5b60018081111561064857610647611ecf565b5b600560149054906101000a900460ff16600181111561066a57610669611ecf565b5b146106aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a190611bd3565b60405180910390fd5b80600f90805190602001906106c0929190611444565b507f2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d816040516106f09190611a71565b60405180910390a150565b600560149054906101000a900460ff1681565b600f805461071b90611e0e565b80601f016020809104026020016040519081016040528092919081815260200182805461074790611e0e565b80156107945780601f1061076957610100808354040283529160200191610794565b820191906000526020600020905b81548152906001019060200180831161077757829003601f168201915b505050505081565b600081141561092357600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061084e5750600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b61088d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088490611b13565b60405180910390fd5b6000600a544261089d9190611d2b565b90506000600954826009546108b29190611d2b565b6006546108bf9190611cd1565b6108c99190611ca0565b90506000816006546108db9190611d2b565b90506108e78183610fce565b6108ef611176565b7ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a1505050610b13565b6001811415610a0a576000600a544261093c9190611d2b565b90506000600954826009546109519190611d2b565b60065461095e9190611cd1565b6109689190611ca0565b9050600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f190611bf3565b60405180910390fd5b610a038161122c565b5050610b12565b6002811480610a195750600381145b15610acd576000600a5442610a2e9190611d2b565b9050600954811015610a75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6c90611ab3565b60405180910390fd5b6003821415610a93576000806006549050610a908183610fce565b50505b610a9b611176565b7ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a150610b11565b60048110610b10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0790611b73565b60405180910390fd5b5b5b5b50565b60065481565b60075481565b60085481565b600a5481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006001811115610b6857610b67611ecf565b5b600560149054906101000a900460ff166001811115610b8a57610b89611ecf565b5b14610bca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc190611b33565b60405180910390fd5b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5190611b53565b60405180910390fd5b81600f9080519060200190610c70929190611444565b5080600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600a819055506001600560146101000a81548160ff02191690836001811115610cdf57610cde611ecf565b5b0217905550610d35600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166006546113b6565b3373ffffffffffffffffffffffffffffffffffffffff167f0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb60405160405180910390a25050565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e0c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e0390611bb3565b60405180910390fd5b600180811115610e1f57610e1e611ecf565b5b600560149054906101000a900460ff166001811115610e4157610e40611ecf565b5b14610e81576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7890611ad3565b60405180910390fd5b6002811480610e905750600381145b610ecf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec690611af3565b60405180910390fd5b84600681905550836007819055508260088190555081600981905550610ef48161079c565b7f03e052767f275c0c51cc93a76255d42498341feb7a5beef7cc11fd57c5b6681860405160405180910390a15050505050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b815260040161104d929190611a2d565b602060405180830381600087803b15801561106757600080fd5b505af115801561107b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109f91906115ab565b50600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161111f929190611a2d565b602060405180830381600087803b15801561113957600080fd5b505af115801561114d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061117191906115ab565b505050565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180602001604052806000815250600f90805190602001906111fe929190611444565b506000600560146101000a81548160ff0219169083600181111561122557611224611ecf565b5b0217905550565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016112e99190611a12565b60206040518083038186803b15801561130157600080fd5b505afa158015611315573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061133991906116aa565b6113439190611d2b565b6040518363ffffffff1660e01b8152600401611360929190611a2d565b602060405180830381600087803b15801561137a57600080fd5b505af115801561138e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113b291906115ab565b5050565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600381905550505050565b82805461145090611e0e565b90600052602060002090601f01602090048101928261147257600085556114b9565b82601f1061148b57805160ff19168380011785556114b9565b828001600101855582156114b9579182015b828111156114b857825182559160200191906001019061149d565b5b5090506114c691906114ca565b5090565b5b808211156114e35760008160009055506001016114cb565b5090565b60006114fa6114f584611c53565b611c2e565b90508281526020810184848401111561151657611515611f61565b5b611521848285611dcc565b509392505050565b600081359050611538816123bb565b92915050565b60008151905061154d816123d2565b92915050565b600082601f83011261156857611567611f5c565b5b81356115788482602086016114e7565b91505092915050565b600081359050611590816123e9565b92915050565b6000815190506115a5816123e9565b92915050565b6000602082840312156115c1576115c0611f6b565b5b60006115cf8482850161153e565b91505092915050565b6000602082840312156115ee576115ed611f6b565b5b600082013567ffffffffffffffff81111561160c5761160b611f66565b5b61161884828501611553565b91505092915050565b6000806040838503121561163857611637611f6b565b5b600083013567ffffffffffffffff81111561165657611655611f66565b5b61166285828601611553565b925050602061167385828601611529565b9150509250929050565b60006020828403121561169357611692611f6b565b5b60006116a184828501611581565b91505092915050565b6000602082840312156116c0576116bf611f6b565b5b60006116ce84828501611596565b91505092915050565b600080600080600080600080610100898b0312156116f8576116f7611f6b565b5b60006117068b828c01611581565b98505060206117178b828c01611581565b97505060406117288b828c01611581565b96505060606117398b828c01611581565b955050608061174a8b828c01611529565b94505060a061175b8b828c01611529565b93505060c061176c8b828c01611529565b92505060e061177d8b828c01611529565b9150509295985092959890939650565b600080600080600060a086880312156117a9576117a8611f6b565b5b60006117b788828901611581565b95505060206117c888828901611581565b94505060406117d988828901611581565b93505060606117ea88828901611581565b92505060806117fb88828901611581565b9150509295509295909350565b61181181611d5f565b82525050565b61182081611dba565b82525050565b600061183182611c84565b61183b8185611c8f565b935061184b818560208601611ddb565b61185481611f70565b840191505092915050565b600061186c604383611c8f565b915061187782611f81565b606082019050919050565b600061188f602a83611c8f565b915061189a82611ff6565b604082019050919050565b60006118b2602b83611c8f565b91506118bd82612045565b604082019050919050565b60006118d5602883611c8f565b91506118e082612094565b604082019050919050565b60006118f8603b83611c8f565b9150611903826120e3565b604082019050919050565b600061191b602583611c8f565b915061192682612132565b604082019050919050565b600061193e603a83611c8f565b915061194982612181565b604082019050919050565b6000611961602983611c8f565b915061196c826121d0565b604082019050919050565b6000611984602e83611c8f565b915061198f8261221f565b604082019050919050565b60006119a7604383611c8f565b91506119b28261226e565b606082019050919050565b60006119ca602883611c8f565b91506119d5826122e3565b604082019050919050565b60006119ed604183611c8f565b91506119f882612332565b606082019050919050565b611a0c81611db0565b82525050565b6000602082019050611a276000830184611808565b92915050565b6000604082019050611a426000830185611808565b611a4f6020830184611a03565b9392505050565b6000602082019050611a6b6000830184611817565b92915050565b60006020820190508181036000830152611a8b8184611826565b905092915050565b60006020820190508181036000830152611aac8161185f565b9050919050565b60006020820190508181036000830152611acc81611882565b9050919050565b60006020820190508181036000830152611aec816118a5565b9050919050565b60006020820190508181036000830152611b0c816118c8565b9050919050565b60006020820190508181036000830152611b2c816118eb565b9050919050565b60006020820190508181036000830152611b4c8161190e565b9050919050565b60006020820190508181036000830152611b6c81611931565b9050919050565b60006020820190508181036000830152611b8c81611954565b9050919050565b60006020820190508181036000830152611bac81611977565b9050919050565b60006020820190508181036000830152611bcc8161199a565b9050919050565b60006020820190508181036000830152611bec816119bd565b9050919050565b60006020820190508181036000830152611c0c816119e0565b9050919050565b6000602082019050611c286000830184611a03565b92915050565b6000611c38611c49565b9050611c448282611e40565b919050565b6000604051905090565b600067ffffffffffffffff821115611c6e57611c6d611f2d565b5b611c7782611f70565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000611cab82611db0565b9150611cb683611db0565b925082611cc657611cc5611ea0565b5b828204905092915050565b6000611cdc82611db0565b9150611ce783611db0565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d2057611d1f611e71565b5b828202905092915050565b6000611d3682611db0565b9150611d4183611db0565b925082821015611d5457611d53611e71565b5b828203905092915050565b6000611d6a82611d90565b9050919050565b60008115159050919050565b6000819050611d8b826123a7565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611dc582611d7d565b9050919050565b82818337600083830152505050565b60005b83811015611df9578082015181840152602081019050611dde565b83811115611e08576000848401525b50505050565b60006002820490506001821680611e2657607f821691505b60208210811415611e3a57611e39611efe565b5b50919050565b611e4982611f70565b810181811067ffffffffffffffff82111715611e6857611e67611f2d565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f74686973206163636f756e74206973206e6f7420617574686f72697a6564207460008201527f6f2075706461746520746865206369706865727465787420696e666f726d617460208201527f696f6e0000000000000000000000000000000000000000000000000000000000604082015250565b7f74686520636f6e7472616374206861732079657420746f20626520636172726960008201527f656420746f207465726d00000000000000000000000000000000000000000000602082015250565b7f74686973206973206163636f756e74206973206e6f7420696e2074686520727560008201527f6e6e696e67207374617465000000000000000000000000000000000000000000602082015250565b7f796f752063616e206f6e6c792075736520636c6f73656f7574206f7074696f6e60008201527f732032206f722033000000000000000000000000000000000000000000000000602082015250565b7f74686973206163636f756e74206973206e6f7420617574686f72697a6564207460008201527f6f207472696767657220616e206561726c7920636c6f73656f75740000000000602082015250565b7f636f6e7472616374206973206e6f7420696e20616e20617661696c61626c652060008201527f7374617465000000000000000000000000000000000000000000000000000000602082015250565b7f746869732061646472657373206973206e6f7420617070726f76656420746f2060008201527f63616c6c207468652070757263686173652066756e6374696f6e000000000000602082015250565b7f796f75206d757374206d616b6520612073656c656374696f6e2062657477656560008201527f6e203020616e6420330000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f74686973206973206163636f756e74206973206e6f7420617574686f72697a6560008201527f6420746f207570646174652074686520636f6e747261637420706172616d657460208201527f6572730000000000000000000000000000000000000000000000000000000000604082015250565b7f74686520636f6e7472616374206973206e6f7420696e207468652072756e6e6960008201527f6e67207374617465000000000000000000000000000000000000000000000000602082015250565b7f74686973206163636f756e74206973206e6f7420617574686f72697a6564207460008201527f6f20747269676765722061206d69642d636f6e747261637420636c6f73656f7560208201527f7400000000000000000000000000000000000000000000000000000000000000604082015250565b600281106123b8576123b7611ecf565b5b50565b6123c481611d5f565b81146123cf57600080fd5b50565b6123db81611d71565b81146123e657600080fd5b50565b6123f281611db0565b81146123fd57600080fd5b5056fea2646970667358221220a7bf8d7ebf9e51d1e72e17c4703b042e1d3fa2a4126b4fe741691ce3e5329ebb64736f6c63430008070033",
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

// EncryptedPoolData is a free data retrieval call binding the contract method 0x8b7e4b13.
//
// Solidity: function encryptedPoolData() view returns(string)
func (_Implementation *ImplementationCaller) EncryptedPoolData(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "encryptedPoolData")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EncryptedPoolData is a free data retrieval call binding the contract method 0x8b7e4b13.
//
// Solidity: function encryptedPoolData() view returns(string)
func (_Implementation *ImplementationSession) EncryptedPoolData() (string, error) {
	return _Implementation.Contract.EncryptedPoolData(&_Implementation.CallOpts)
}

// EncryptedPoolData is a free data retrieval call binding the contract method 0x8b7e4b13.
//
// Solidity: function encryptedPoolData() view returns(string)
func (_Implementation *ImplementationCallerSession) EncryptedPoolData() (string, error) {
	return _Implementation.Contract.EncryptedPoolData(&_Implementation.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x3a5d4213.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _seller, address _lmn, address _cloneFactory, address _validator) returns()
func (_Implementation *ImplementationTransactor) Initialize(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _seller common.Address, _lmn common.Address, _cloneFactory common.Address, _validator common.Address) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "initialize", _price, _limit, _speed, _length, _seller, _lmn, _cloneFactory, _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0x3a5d4213.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _seller, address _lmn, address _cloneFactory, address _validator) returns()
func (_Implementation *ImplementationSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _seller common.Address, _lmn common.Address, _cloneFactory common.Address, _validator common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _seller, _lmn, _cloneFactory, _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0x3a5d4213.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _seller, address _lmn, address _cloneFactory, address _validator) returns()
func (_Implementation *ImplementationTransactorSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _seller common.Address, _lmn common.Address, _cloneFactory common.Address, _validator common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _seller, _lmn, _cloneFactory, _validator)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x8e2e6d5d.
//
// Solidity: function setContractCloseOut(uint256 closeOutType) returns()
func (_Implementation *ImplementationTransactor) SetContractCloseOut(opts *bind.TransactOpts, closeOutType *big.Int) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setContractCloseOut", closeOutType)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x8e2e6d5d.
//
// Solidity: function setContractCloseOut(uint256 closeOutType) returns()
func (_Implementation *ImplementationSession) SetContractCloseOut(closeOutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetContractCloseOut(&_Implementation.TransactOpts, closeOutType)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x8e2e6d5d.
//
// Solidity: function setContractCloseOut(uint256 closeOutType) returns()
func (_Implementation *ImplementationTransactorSession) SetContractCloseOut(closeOutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetContractCloseOut(&_Implementation.TransactOpts, closeOutType)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0xddcb1bf2.
//
// Solidity: function setPurchaseContract(string _encryptedPoolData, address _buyer) returns()
func (_Implementation *ImplementationTransactor) SetPurchaseContract(opts *bind.TransactOpts, _encryptedPoolData string, _buyer common.Address) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setPurchaseContract", _encryptedPoolData, _buyer)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0xddcb1bf2.
//
// Solidity: function setPurchaseContract(string _encryptedPoolData, address _buyer) returns()
func (_Implementation *ImplementationSession) SetPurchaseContract(_encryptedPoolData string, _buyer common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encryptedPoolData, _buyer)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0xddcb1bf2.
//
// Solidity: function setPurchaseContract(string _encryptedPoolData, address _buyer) returns()
func (_Implementation *ImplementationTransactorSession) SetPurchaseContract(_encryptedPoolData string, _buyer common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encryptedPoolData, _buyer)
}

// SetUpdateMiningInformation is a paid mutator transaction binding the contract method 0x719e6b5b.
//
// Solidity: function setUpdateMiningInformation(string _newEncryptedPoolData) returns()
func (_Implementation *ImplementationTransactor) SetUpdateMiningInformation(opts *bind.TransactOpts, _newEncryptedPoolData string) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setUpdateMiningInformation", _newEncryptedPoolData)
}

// SetUpdateMiningInformation is a paid mutator transaction binding the contract method 0x719e6b5b.
//
// Solidity: function setUpdateMiningInformation(string _newEncryptedPoolData) returns()
func (_Implementation *ImplementationSession) SetUpdateMiningInformation(_newEncryptedPoolData string) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdateMiningInformation(&_Implementation.TransactOpts, _newEncryptedPoolData)
}

// SetUpdateMiningInformation is a paid mutator transaction binding the contract method 0x719e6b5b.
//
// Solidity: function setUpdateMiningInformation(string _newEncryptedPoolData) returns()
func (_Implementation *ImplementationTransactorSession) SetUpdateMiningInformation(_newEncryptedPoolData string) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdateMiningInformation(&_Implementation.TransactOpts, _newEncryptedPoolData)
}

// SetUpdatePurchaseInformation is a paid mutator transaction binding the contract method 0xf1a6e301.
//
// Solidity: function setUpdatePurchaseInformation(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _closeoutType) returns()
func (_Implementation *ImplementationTransactor) SetUpdatePurchaseInformation(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _closeoutType *big.Int) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setUpdatePurchaseInformation", _price, _limit, _speed, _length, _closeoutType)
}

// SetUpdatePurchaseInformation is a paid mutator transaction binding the contract method 0xf1a6e301.
//
// Solidity: function setUpdatePurchaseInformation(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _closeoutType) returns()
func (_Implementation *ImplementationSession) SetUpdatePurchaseInformation(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _closeoutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdatePurchaseInformation(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _closeoutType)
}

// SetUpdatePurchaseInformation is a paid mutator transaction binding the contract method 0xf1a6e301.
//
// Solidity: function setUpdatePurchaseInformation(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _closeoutType) returns()
func (_Implementation *ImplementationTransactorSession) SetUpdatePurchaseInformation(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _closeoutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdatePurchaseInformation(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _closeoutType)
}

// ImplementationCipherTextUpdatedIterator is returned from FilterCipherTextUpdated and is used to iterate over the raw logs and unpacked data for CipherTextUpdated events raised by the Implementation contract.
type ImplementationCipherTextUpdatedIterator struct {
	Event *ImplementationCipherTextUpdated // Event containing the contract specifics and raw log

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
func (it *ImplementationCipherTextUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationCipherTextUpdated)
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
		it.Event = new(ImplementationCipherTextUpdated)
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
func (it *ImplementationCipherTextUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationCipherTextUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationCipherTextUpdated represents a CipherTextUpdated event raised by the Implementation contract.
type ImplementationCipherTextUpdated struct {
	NewCipherText string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCipherTextUpdated is a free log retrieval operation binding the contract event 0x2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d.
//
// Solidity: event cipherTextUpdated(string newCipherText)
func (_Implementation *ImplementationFilterer) FilterCipherTextUpdated(opts *bind.FilterOpts) (*ImplementationCipherTextUpdatedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "cipherTextUpdated")
	if err != nil {
		return nil, err
	}
	return &ImplementationCipherTextUpdatedIterator{contract: _Implementation.contract, event: "cipherTextUpdated", logs: logs, sub: sub}, nil
}

// WatchCipherTextUpdated is a free log subscription operation binding the contract event 0x2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d.
//
// Solidity: event cipherTextUpdated(string newCipherText)
func (_Implementation *ImplementationFilterer) WatchCipherTextUpdated(opts *bind.WatchOpts, sink chan<- *ImplementationCipherTextUpdated) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "cipherTextUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationCipherTextUpdated)
				if err := _Implementation.contract.UnpackLog(event, "cipherTextUpdated", log); err != nil {
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

// ParseCipherTextUpdated is a log parse operation binding the contract event 0x2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d.
//
// Solidity: event cipherTextUpdated(string newCipherText)
func (_Implementation *ImplementationFilterer) ParseCipherTextUpdated(log types.Log) (*ImplementationCipherTextUpdated, error) {
	event := new(ImplementationCipherTextUpdated)
	if err := _Implementation.contract.UnpackLog(event, "cipherTextUpdated", log); err != nil {
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
// Solidity: event contractPurchased(address indexed _buyer)
func (_Implementation *ImplementationFilterer) FilterContractPurchased(opts *bind.FilterOpts, _buyer []common.Address) (*ImplementationContractPurchasedIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractPurchased", _buyerRule)
	if err != nil {
		return nil, err
	}
	return &ImplementationContractPurchasedIterator{contract: _Implementation.contract, event: "contractPurchased", logs: logs, sub: sub}, nil
}

// WatchContractPurchased is a free log subscription operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address indexed _buyer)
func (_Implementation *ImplementationFilterer) WatchContractPurchased(opts *bind.WatchOpts, sink chan<- *ImplementationContractPurchased, _buyer []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractPurchased", _buyerRule)
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
// Solidity: event contractPurchased(address indexed _buyer)
func (_Implementation *ImplementationFilterer) ParseContractPurchased(log types.Log) (*ImplementationContractPurchased, error) {
	event := new(ImplementationContractPurchased)
	if err := _Implementation.contract.UnpackLog(event, "contractPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationPurchaseInfoUpdatedIterator is returned from FilterPurchaseInfoUpdated and is used to iterate over the raw logs and unpacked data for PurchaseInfoUpdated events raised by the Implementation contract.
type ImplementationPurchaseInfoUpdatedIterator struct {
	Event *ImplementationPurchaseInfoUpdated // Event containing the contract specifics and raw log

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
func (it *ImplementationPurchaseInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationPurchaseInfoUpdated)
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
		it.Event = new(ImplementationPurchaseInfoUpdated)
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
func (it *ImplementationPurchaseInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationPurchaseInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationPurchaseInfoUpdated represents a PurchaseInfoUpdated event raised by the Implementation contract.
type ImplementationPurchaseInfoUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPurchaseInfoUpdated is a free log retrieval operation binding the contract event 0x03e052767f275c0c51cc93a76255d42498341feb7a5beef7cc11fd57c5b66818.
//
// Solidity: event purchaseInfoUpdated()
func (_Implementation *ImplementationFilterer) FilterPurchaseInfoUpdated(opts *bind.FilterOpts) (*ImplementationPurchaseInfoUpdatedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "purchaseInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &ImplementationPurchaseInfoUpdatedIterator{contract: _Implementation.contract, event: "purchaseInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPurchaseInfoUpdated is a free log subscription operation binding the contract event 0x03e052767f275c0c51cc93a76255d42498341feb7a5beef7cc11fd57c5b66818.
//
// Solidity: event purchaseInfoUpdated()
func (_Implementation *ImplementationFilterer) WatchPurchaseInfoUpdated(opts *bind.WatchOpts, sink chan<- *ImplementationPurchaseInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "purchaseInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationPurchaseInfoUpdated)
				if err := _Implementation.contract.UnpackLog(event, "purchaseInfoUpdated", log); err != nil {
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

// ParsePurchaseInfoUpdated is a log parse operation binding the contract event 0x03e052767f275c0c51cc93a76255d42498341feb7a5beef7cc11fd57c5b66818.
//
// Solidity: event purchaseInfoUpdated()
func (_Implementation *ImplementationFilterer) ParsePurchaseInfoUpdated(log types.Log) (*ImplementationPurchaseInfoUpdated, error) {
	event := new(ImplementationPurchaseInfoUpdated)
	if err := _Implementation.contract.UnpackLog(event, "purchaseInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
