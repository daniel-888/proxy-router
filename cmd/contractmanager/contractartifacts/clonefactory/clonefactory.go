// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package clonefactory

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

// ClonefactoryMetaData contains all meta data concerning the Clonefactory contract.
var ClonefactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"contractCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validationFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"setCreateNewRentalContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516126423803806126428339818101604052810190610032919061018d565b60006040516100409061016b565b604051809103906000f08015801561005c573d6000803e3d6000fd5b509050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505061022e565b611f048061073e83390190565b60008151905061018781610217565b92915050565b6000806000606084860312156101a6576101a5610212565b5b60006101b486828701610178565b93505060206101c586828701610178565b92505060406101d686828701610178565b9150509250925092565b60006101eb826101f2565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b610220816101e0565b811461022b57600080fd5b50565b6105018061023d6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633e45cccd14610030575b600080fd5b61004a6004803603810190610045919061029b565b610060565b6040516100579190610369565b60405180910390f35b60008061008c60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661019c565b90508073ffffffffffffffffffffffffffffffffffffffff166374a34812898989898989600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518963ffffffff1660e01b81526004016101199897969594939291906103a4565b600060405180830381600087803b15801561013357600080fd5b505af1158015610147573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff167ffcf9a0c9dedbfcd1a047374855fc36baaf605bd4f4837802a0cc938ba1b5f30260405160405180910390a2809150509695505050505050565b60006040517f3d602d80600a3d3981f3363d3d373d3d3d363d7300000000000000000000000081528260601b60148201527f5af43d82803e903d91602b57fd5bf3000000000000000000000000000000000060288201526037816000f0915050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561026c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026390610384565b60405180910390fd5b919050565b6000813590506102808161049d565b92915050565b600081359050610295816104b4565b92915050565b60008060008060008060c087890312156102b8576102b761046f565b5b60006102c689828a01610286565b96505060206102d789828a01610286565b95505060406102e889828a01610286565b94505060606102f989828a01610286565b935050608061030a89828a01610286565b92505060a061031b89828a01610271565b9150509295509295509295565b61033181610433565b82525050565b6000610344601683610422565b915061034f82610474565b602082019050919050565b61036381610465565b82525050565b600060208201905061037e6000830184610328565b92915050565b6000602082019050818103600083015261039d81610337565b9050919050565b6000610100820190506103ba600083018b61035a565b6103c7602083018a61035a565b6103d4604083018961035a565b6103e1606083018861035a565b6103ee608083018761035a565b6103fb60a0830186610328565b61040860c0830185610328565b61041560e0830184610328565b9998505050505050505050565b600082825260208201905092915050565b600061043e82610445565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600080fd5b7f455243313136373a20637265617465206661696c656400000000000000000000600082015250565b6104a681610433565b81146104b157600080fd5b50565b6104bd81610465565b81146104c857600080fd5b5056fea2646970667358221220108dfb6ceda536403477baa4b86e8782793bb6e8f54c3826d5f15bb284fe0a4264736f6c63430008070033608060405234801561001057600080fd5b50611ee4806100206000396000f3fe60806040526004361061011f5760003560e01c806374a34812116100a0578063c204296511610064578063c204296514610372578063c20906ac1461039d578063c5095d68146103c8578063ce0c722a146103f3578063f8bbf27e1461041e5761011f565b806374a34812146102b157806385209ee0146102da578063a035b1fe14610305578063a4d66daf14610330578063b6293b341461035b5761011f565b80631f7b6d32116100e75780631f7b6d32146101ec57806326b3c68b146102175780634897b9ac1461024257806349823485146102595780637150d8ae146102865761011f565b806308551a5314610124578063089aa8a21461014f5780630a61e2d91461017a57806311afdfc9146101a557806316713b37146101c1575b600080fd5b34801561013057600080fd5b50610139610449565b60405161014691906117ee565b60405180910390f35b34801561015b57600080fd5b5061016461046f565b60405161017191906117ee565b60405180910390f35b34801561018657600080fd5b5061018f610495565b60405161019c9190611959565b60405180910390f35b6101bf60048036038101906101ba91906114f2565b61049b565b005b3480156101cd57600080fd5b506101d66107e8565b6040516101e39190611959565b60405180910390f35b3480156101f857600080fd5b506102016107ee565b60405161020e9190611959565b60405180910390f35b34801561022357600080fd5b5061022c6107f4565b6040516102399190611959565b60405180910390f35b34801561024e57600080fd5b506102576107fa565b005b34801561026557600080fd5b5061026e610955565b60405161027d9392919061184d565b60405180910390f35b34801561029257600080fd5b5061029b610ba2565b6040516102a891906117ee565b60405180910390f35b3480156102bd57600080fd5b506102d860048036038101906102d39190611600565b610bc8565b005b3480156102e657600080fd5b506102ef610da5565b6040516102fc9190611832565b60405180910390f35b34801561031157600080fd5b5061031a610db8565b6040516103279190611959565b60405180910390f35b34801561033c57600080fd5b50610345610dbe565b6040516103529190611959565b60405180910390f35b34801561036757600080fd5b50610370610dc4565b005b34801561037e57600080fd5b50610387610f74565b6040516103949190611959565b60405180910390f35b3480156103a957600080fd5b506103b2611033565b6040516103bf9190611959565b60405180910390f35b3480156103d457600080fd5b506103dd611039565b6040516103ea9190611959565b60405180910390f35b3480156103ff57600080fd5b5061040861103f565b60405161041591906117ee565b60405180910390f35b34801561042a57600080fd5b50610433611065565b6040516104409190611959565b60405180910390f35b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b600060038111156104af576104ae611c15565b5b600560149054906101000a900460ff1660038111156104d1576104d0611c15565b5b14610511576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610508906118d9565b60405180910390fd5b8560119080519060200190610527929190611349565b50846012908051906020019061053e929190611349565b508360139080519060200190610555929190611349565b5082600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610629600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660065461106b565b6001151581151514156106d6576001601460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550606434146106d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c890611919565b60405180910390fd5b6107a9565b6001601460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160146000600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b7f0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb336040516107d891906117ee565b60405180910390a1505050505050565b60045481565b60095481565b600a5481565b60011515601460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461088d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088490611899565b60405180910390fd5b6000600c544261089d9190611a71565b905060008060095483106108b457600091506108df565b600954836009546108c59190611a71565b6006546108d29190611a17565b6108dc91906119e6565b91505b816006546108ed9190611a71565b90506108f981836110f9565b6003600560146101000a81548160ff0219169083600381111561091f5761091e611c15565b5b02179055507ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a1505050565b606080606060011515601460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515146109ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e490611899565b60405180910390fd5b6011601260138280546109ff90611b54565b80601f0160208091040260200160405190810160405280929190818152602001828054610a2b90611b54565b8015610a785780601f10610a4d57610100808354040283529160200191610a78565b820191906000526020600020905b815481529060010190602001808311610a5b57829003601f168201915b50505050509250818054610a8b90611b54565b80601f0160208091040260200160405190810160405280929190818152602001828054610ab790611b54565b8015610b045780601f10610ad957610100808354040283529160200191610b04565b820191906000526020600020905b815481529060010190602001808311610ae757829003601f168201915b50505050509150808054610b1790611b54565b80601f0160208091040260200160405190810160405280929190818152602001828054610b4390611b54565b8015610b905780601f10610b6557610100808354040283529160200191610b90565b820191906000526020600020905b815481529060010190602001808311610b7357829003601f168201915b50505050509050925092509250909192565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900460ff1680610bee575060008054906101000a900460ff16155b610c2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c24906118f9565b60405180910390fd5b60008060019054906101000a900460ff161590508015610c7d576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b8860068190555087600781905550866008819055508560098190555083600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600b8190555082600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560146101000a81548160ff02191690836003811115610d4857610d47611c15565b5b0217905550610d79600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836112a1565b8015610d9a5760008060016101000a81548160ff0219169083151502179055505b505050505050505050565b600560149054906101000a900460ff1681565b60065481565b60075481565b6000610dce610f74565b14610e0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e05906118b9565b60405180910390fd5b6001151560146000600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151480610f0157506001151560146000600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515145b610f40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f3790611939565b60405180910390fd5b42600c819055506002600560146101000a81548160ff02191690836003811115610f6d57610f6c611c15565b5b0217905550565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610fd191906117ee565b60206040518083038186803b158015610fe957600080fd5b505afa158015610ffd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061102191906115d3565b60035461102e9190611a71565b905090565b60085481565b600c5481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b5481565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600381905550505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b8152600401611178929190611809565b602060405180830381600087803b15801561119257600080fd5b505af11580156111a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ca91906114c5565b50600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161124a929190611809565b602060405180830381600087803b15801561126457600080fd5b505af1158015611278573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129c91906114c5565b505050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b82805461135590611b54565b90600052602060002090601f01602090048101928261137757600085556113be565b82601f1061139057805160ff19168380011785556113be565b828001600101855582156113be579182015b828111156113bd5782518255916020019190600101906113a2565b5b5090506113cb91906113cf565b5090565b5b808211156113e85760008160009055506001016113d0565b5090565b60006113ff6113fa84611999565b611974565b90508281526020810184848401111561141b5761141a611ca7565b5b611426848285611b12565b509392505050565b60008135905061143d81611e69565b92915050565b60008135905061145281611e80565b92915050565b60008151905061146781611e80565b92915050565b600082601f83011261148257611481611ca2565b5b81356114928482602086016113ec565b91505092915050565b6000813590506114aa81611e97565b92915050565b6000815190506114bf81611e97565b92915050565b6000602082840312156114db576114da611cb1565b5b60006114e984828501611458565b91505092915050565b60008060008060008060c0878903121561150f5761150e611cb1565b5b600087013567ffffffffffffffff81111561152d5761152c611cac565b5b61153989828a0161146d565b965050602087013567ffffffffffffffff81111561155a57611559611cac565b5b61156689828a0161146d565b955050604087013567ffffffffffffffff81111561158757611586611cac565b5b61159389828a0161146d565b94505060606115a489828a0161142e565b93505060806115b589828a0161142e565b92505060a06115c689828a01611443565b9150509295509295509295565b6000602082840312156115e9576115e8611cb1565b5b60006115f7848285016114b0565b91505092915050565b600080600080600080600080610100898b03121561162157611620611cb1565b5b600061162f8b828c0161149b565b98505060206116408b828c0161149b565b97505060406116518b828c0161149b565b96505060606116628b828c0161149b565b95505060806116738b828c0161149b565b94505060a06116848b828c0161142e565b93505060c06116958b828c0161142e565b92505060e06116a68b828c0161142e565b9150509295985092959890939650565b6116bf81611aa5565b82525050565b6116ce81611b00565b82525050565b60006116df826119ca565b6116e981856119d5565b93506116f9818560208601611b21565b61170281611cb6565b840191505092915050565b600061171a6031836119d5565b915061172582611cc7565b604082019050919050565b600061173d602e836119d5565b915061174882611d16565b604082019050919050565b60006117606025836119d5565b915061176b82611d65565b604082019050919050565b6000611783602e836119d5565b915061178e82611db4565b604082019050919050565b60006117a66017836119d5565b91506117b182611e03565b602082019050919050565b60006117c9601a836119d5565b91506117d482611e2c565b602082019050919050565b6117e881611af6565b82525050565b600060208201905061180360008301846116b6565b92915050565b600060408201905061181e60008301856116b6565b61182b60208301846117df565b9392505050565b600060208201905061184760008301846116c5565b92915050565b6000606082019050818103600083015261186781866116d4565b9050818103602083015261187b81856116d4565b9050818103604083015261188f81846116d4565b9050949350505050565b600060208201905081810360008301526118b28161170d565b9050919050565b600060208201905081810360008301526118d281611730565b9050919050565b600060208201905081810360008301526118f281611753565b9050919050565b6000602082019050818103600083015261191281611776565b9050919050565b6000602082019050818103600083015261193281611799565b9050919050565b60006020820190508181036000830152611952816117bc565b9050919050565b600060208201905061196e60008301846117df565b92915050565b600061197e61198f565b905061198a8282611b86565b919050565b6000604051905090565b600067ffffffffffffffff8211156119b4576119b3611c73565b5b6119bd82611cb6565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60006119f182611af6565b91506119fc83611af6565b925082611a0c57611a0b611be6565b5b828204905092915050565b6000611a2282611af6565b9150611a2d83611af6565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611a6657611a65611bb7565b5b828202905092915050565b6000611a7c82611af6565b9150611a8783611af6565b925082821015611a9a57611a99611bb7565b5b828203905092915050565b6000611ab082611ad6565b9050919050565b60008115159050919050565b6000819050611ad182611e55565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611b0b82611ac3565b9050919050565b82818337600083830152505050565b60005b83811015611b3f578082015181840152602081019050611b24565b83811115611b4e576000848401525b50505050565b60006002820490506001821680611b6c57607f821691505b60208210811415611b8057611b7f611c44565b5b50919050565b611b8f82611cb6565b810181811067ffffffffffffffff82111715611bae57611bad611c73565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f746869732061646472657373206973206e6f7420616c6c6f77656420746f206360008201527f616c6c20746869732066756e6374696f6e000000000000000000000000000000602082015250565b7f6c756d6572696e20746f6b656e73206e65656420746f2062652073656e74207460008201527f6f2074686520636f6e7472616374000000000000000000000000000000000000602082015250565b7f636f6e7472616374206973206e6f7420696e20616e20617661696c61626c652060008201527f7374617465000000000000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f76616c69646174696f6e20666565206e6f742073656e74000000000000000000600082015250565b7f74686520627579657220686173206e6f74206265656e20736574000000000000600082015250565b60048110611e6657611e65611c15565b5b50565b611e7281611aa5565b8114611e7d57600080fd5b50565b611e8981611ab7565b8114611e9457600080fd5b50565b611ea081611af6565b8114611eab57600080fd5b5056fea2646970667358221220ecbf1beda227ec85143d1321c00692a94a3a315e0c63de24f7e15e90f8caa82264736f6c63430008070033",
}

// ClonefactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ClonefactoryMetaData.ABI instead.
var ClonefactoryABI = ClonefactoryMetaData.ABI

// ClonefactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClonefactoryMetaData.Bin instead.
var ClonefactoryBin = ClonefactoryMetaData.Bin

// DeployClonefactory deploys a new Ethereum contract, binding an instance of Clonefactory to it.
func DeployClonefactory(auth *bind.TransactOpts, backend bind.ContractBackend, _lmn common.Address, _validator common.Address, _proxy common.Address) (common.Address, *types.Transaction, *Clonefactory, error) {
	parsed, err := ClonefactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClonefactoryBin), backend, _lmn, _validator, _proxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Clonefactory{ClonefactoryCaller: ClonefactoryCaller{contract: contract}, ClonefactoryTransactor: ClonefactoryTransactor{contract: contract}, ClonefactoryFilterer: ClonefactoryFilterer{contract: contract}}, nil
}

// Clonefactory is an auto generated Go binding around an Ethereum contract.
type Clonefactory struct {
	ClonefactoryCaller     // Read-only binding to the contract
	ClonefactoryTransactor // Write-only binding to the contract
	ClonefactoryFilterer   // Log filterer for contract events
}

// ClonefactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClonefactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonefactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClonefactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonefactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClonefactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonefactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClonefactorySession struct {
	Contract     *Clonefactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonefactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClonefactoryCallerSession struct {
	Contract *ClonefactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ClonefactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClonefactoryTransactorSession struct {
	Contract     *ClonefactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ClonefactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClonefactoryRaw struct {
	Contract *Clonefactory // Generic contract binding to access the raw methods on
}

// ClonefactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClonefactoryCallerRaw struct {
	Contract *ClonefactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ClonefactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClonefactoryTransactorRaw struct {
	Contract *ClonefactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClonefactory creates a new instance of Clonefactory, bound to a specific deployed contract.
func NewClonefactory(address common.Address, backend bind.ContractBackend) (*Clonefactory, error) {
	contract, err := bindClonefactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clonefactory{ClonefactoryCaller: ClonefactoryCaller{contract: contract}, ClonefactoryTransactor: ClonefactoryTransactor{contract: contract}, ClonefactoryFilterer: ClonefactoryFilterer{contract: contract}}, nil
}

// NewClonefactoryCaller creates a new read-only instance of Clonefactory, bound to a specific deployed contract.
func NewClonefactoryCaller(address common.Address, caller bind.ContractCaller) (*ClonefactoryCaller, error) {
	contract, err := bindClonefactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryCaller{contract: contract}, nil
}

// NewClonefactoryTransactor creates a new write-only instance of Clonefactory, bound to a specific deployed contract.
func NewClonefactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ClonefactoryTransactor, error) {
	contract, err := bindClonefactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryTransactor{contract: contract}, nil
}

// NewClonefactoryFilterer creates a new log filterer instance of Clonefactory, bound to a specific deployed contract.
func NewClonefactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ClonefactoryFilterer, error) {
	contract, err := bindClonefactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryFilterer{contract: contract}, nil
}

// bindClonefactory binds a generic wrapper to an already deployed contract.
func bindClonefactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonefactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clonefactory *ClonefactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clonefactory.Contract.ClonefactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clonefactory *ClonefactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonefactory.Contract.ClonefactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clonefactory *ClonefactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clonefactory.Contract.ClonefactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clonefactory *ClonefactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clonefactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clonefactory *ClonefactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonefactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clonefactory *ClonefactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clonefactory.Contract.contract.Transact(opts, method, params...)
}

// SetCreateNewRentalContract is a paid mutator transaction binding the contract method 0x3e45cccd.
//
// Solidity: function setCreateNewRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _validationFee, address _seller) returns(address)
func (_Clonefactory *ClonefactoryTransactor) SetCreateNewRentalContract(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validationFee *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setCreateNewRentalContract", _price, _limit, _speed, _length, _validationFee, _seller)
}

// SetCreateNewRentalContract is a paid mutator transaction binding the contract method 0x3e45cccd.
//
// Solidity: function setCreateNewRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _validationFee, address _seller) returns(address)
func (_Clonefactory *ClonefactorySession) SetCreateNewRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validationFee *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContract(&_Clonefactory.TransactOpts, _price, _limit, _speed, _length, _validationFee, _seller)
}

// SetCreateNewRentalContract is a paid mutator transaction binding the contract method 0x3e45cccd.
//
// Solidity: function setCreateNewRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _validationFee, address _seller) returns(address)
func (_Clonefactory *ClonefactoryTransactorSession) SetCreateNewRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validationFee *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContract(&_Clonefactory.TransactOpts, _price, _limit, _speed, _length, _validationFee, _seller)
}

// ClonefactoryContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the Clonefactory contract.
type ClonefactoryContractCreatedIterator struct {
	Event *ClonefactoryContractCreated // Event containing the contract specifics and raw log

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
func (it *ClonefactoryContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryContractCreated)
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
		it.Event = new(ClonefactoryContractCreated)
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
func (it *ClonefactoryContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryContractCreated represents a ContractCreated event raised by the Clonefactory contract.
type ClonefactoryContractCreated struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0xfcf9a0c9dedbfcd1a047374855fc36baaf605bd4f4837802a0cc938ba1b5f302.
//
// Solidity: event contractCreated(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) FilterContractCreated(opts *bind.FilterOpts, _address []common.Address) (*ClonefactoryContractCreatedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "contractCreated", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryContractCreatedIterator{contract: _Clonefactory.contract, event: "contractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0xfcf9a0c9dedbfcd1a047374855fc36baaf605bd4f4837802a0cc938ba1b5f302.
//
// Solidity: event contractCreated(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *ClonefactoryContractCreated, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "contractCreated", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryContractCreated)
				if err := _Clonefactory.contract.UnpackLog(event, "contractCreated", log); err != nil {
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
// Solidity: event contractCreated(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) ParseContractCreated(log types.Log) (*ClonefactoryContractCreated, error) {
	event := new(ClonefactoryContractCreated)
	if err := _Clonefactory.contract.UnpackLog(event, "contractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
