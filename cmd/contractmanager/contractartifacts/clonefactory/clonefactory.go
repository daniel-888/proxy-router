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
	Bin: "0x608060405234801561001057600080fd5b506040516125663803806125668339818101604052810190610032919061018d565b60006040516100409061016b565b604051809103906000f08015801561005c573d6000803e3d6000fd5b509050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505061022e565b611e288061073e83390190565b60008151905061018781610217565b92915050565b6000806000606084860312156101a6576101a5610212565b5b60006101b486828701610178565b93505060206101c586828701610178565b92505060406101d686828701610178565b9150509250925092565b60006101eb826101f2565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b610220816101e0565b811461022b57600080fd5b50565b6105018061023d6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80633e45cccd14610030575b600080fd5b61004a6004803603810190610045919061029b565b610060565b6040516100579190610369565b60405180910390f35b60008061008c60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661019c565b90508073ffffffffffffffffffffffffffffffffffffffff166374a34812898989898989600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518963ffffffff1660e01b81526004016101199897969594939291906103a4565b600060405180830381600087803b15801561013357600080fd5b505af1158015610147573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff167ffcf9a0c9dedbfcd1a047374855fc36baaf605bd4f4837802a0cc938ba1b5f30260405160405180910390a2809150509695505050505050565b60006040517f3d602d80600a3d3981f3363d3d373d3d3d363d7300000000000000000000000081528260601b60148201527f5af43d82803e903d91602b57fd5bf3000000000000000000000000000000000060288201526037816000f0915050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561026c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026390610384565b60405180910390fd5b919050565b6000813590506102808161049d565b92915050565b600081359050610295816104b4565b92915050565b60008060008060008060c087890312156102b8576102b761046f565b5b60006102c689828a01610286565b96505060206102d789828a01610286565b95505060406102e889828a01610286565b94505060606102f989828a01610286565b935050608061030a89828a01610286565b92505060a061031b89828a01610271565b9150509295509295509295565b61033181610433565b82525050565b6000610344601683610422565b915061034f82610474565b602082019050919050565b61036381610465565b82525050565b600060208201905061037e6000830184610328565b92915050565b6000602082019050818103600083015261039d81610337565b9050919050565b6000610100820190506103ba600083018b61035a565b6103c7602083018a61035a565b6103d4604083018961035a565b6103e1606083018861035a565b6103ee608083018761035a565b6103fb60a0830186610328565b61040860c0830185610328565b61041560e0830184610328565b9998505050505050505050565b600082825260208201905092915050565b600061043e82610445565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600080fd5b7f455243313136373a20637265617465206661696c656400000000000000000000600082015250565b6104a681610433565b81146104b157600080fd5b50565b6104bd81610465565b81146104c857600080fd5b5056fea2646970667358221220a3dcef72ab1edf0e2b8df594b30207cc127967876d0d84f699f6139d6f1150fd64736f6c63430008070033608060405234801561001057600080fd5b50611e08806100206000396000f3fe6080604052600436106101145760003560e01c80638331ed49116100a0578063b6293b3411610064578063b6293b341461034e578063c20906ac14610365578063c5095d6814610390578063ce0c722a146103bb578063f8bbf27e146103e657610114565b80638331ed491461028657806385209ee0146102a25780638b7e4b13146102cd578063a035b1fe146102f8578063a4d66daf1461032357610114565b80631f7b6d32116100e75780631f7b6d32146101c557806326b3c68b146101f05780634897b9ac1461021b5780637150d8ae1461023257806374a348121461025d57610114565b806308551a5314610119578063089aa8a2146101445780630a61e2d91461016f57806316713b371461019a575b600080fd5b34801561012557600080fd5b5061012e610411565b60405161013b9190611716565b60405180910390f35b34801561015057600080fd5b50610159610437565b6040516101669190611716565b60405180910390f35b34801561017b57600080fd5b5061018461045d565b6040516101919190611857565b60405180910390f35b3480156101a657600080fd5b506101af610463565b6040516101bc9190611857565b60405180910390f35b3480156101d157600080fd5b506101da610469565b6040516101e79190611857565b60405180910390f35b3480156101fc57600080fd5b5061020561046f565b6040516102129190611857565b60405180910390f35b34801561022757600080fd5b50610230610475565b005b34801561023e57600080fd5b506102476105db565b6040516102549190611716565b60405180910390f35b34801561026957600080fd5b50610284600480360381019061027f9190611528565b610601565b005b6102a0600480360381019061029b9190611478565b6107bb565b005b3480156102ae57600080fd5b506102b7610ad4565b6040516102c4919061175a565b60405180910390f35b3480156102d957600080fd5b506102e2610ae7565b6040516102ef9190611775565b60405180910390f35b34801561030457600080fd5b5061030d610b75565b60405161031a9190611857565b60405180910390f35b34801561032f57600080fd5b50610338610b7b565b6040516103459190611857565b60405180910390f35b34801561035a57600080fd5b50610363610b81565b005b34801561037157600080fd5b5061037a610cb4565b6040516103879190611857565b60405180910390f35b34801561039c57600080fd5b506103a5610cba565b6040516103b29190611857565b60405180910390f35b3480156103c757600080fd5b506103d0610cc0565b6040516103dd9190611716565b60405180910390f35b3480156103f257600080fd5b506103fb610ce6565b6040516104089190611857565b60405180910390f35b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60045481565b60095481565b600a5481565b60011515601260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514610508576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ff90611797565b60405180910390fd5b6000600c5442610518919061196f565b9050600080600954831061052f576000915061055a565b60095483600954610540919061196f565b60065461054d9190611915565b61055791906118e4565b91505b81600654610568919061196f565b90506105748183610cec565b6003600560146101000a81548160ff0219169083600381111561059a57610599611b13565b5b02179055507faadd128c35976a01ffffa9dfb8d363b3358597ce6b30248bcf25e80bd3af4512336040516105ce9190611716565b60405180910390a1505050565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900460ff1680610627575060008054906101000a900460ff16155b610666576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065d90611817565b60405180910390fd5b60008060019054906101000a900460ff1615905080156106b6576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b8860068190555087600781905550866008819055508560098190555083600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600b8190555082600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560146101000a81548160ff0219169083600381111561078157610780611b13565b5b021790555061078f82610e94565b80156107b05760008060016101000a81548160ff0219169083151502179055505b505050505050505050565b600060038111156107cf576107ce611b13565b5b600560149054906101000a900460ff1660038111156107f1576107f0611b13565b5b14610831576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610828906117d7565b60405180910390fd5b83601190805190602001906108479291906112cf565b5082600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061091b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600654610f3b565b6001151581151514156109c5576001601260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600b5434146109c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109bb90611837565b60405180910390fd5b5b6001601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160126000600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb33604051610ac69190611716565b60405180910390a150505050565b600560149054906101000a900460ff1681565b60118054610af490611a52565b80601f0160208091040260200160405190810160405280929190818152602001828054610b2090611a52565b8015610b6d5780601f10610b4257610100808354040283529160200191610b6d565b820191906000526020600020905b815481529060010190602001808311610b5057829003601f168201915b505050505081565b60065481565b60075481565b6000610b8b610fc9565b14610bcb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc2906117b7565b60405180910390fd5b6001151560126000600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514610c80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c77906117f7565b60405180910390fd5b42600c819055506002600560146101000a81548160ff02191690836003811115610cad57610cac611b13565b5b0217905550565b60085481565b600c5481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b5481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b8152600401610d6b929190611731565b602060405180830381600087803b158015610d8557600080fd5b505af1158015610d99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dbd919061144b565b50600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610e3d929190611731565b602060405180830381600087803b158015610e5757600080fd5b505af1158015610e6b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8f919061144b565b505050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600381905550505050565b6000600354600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016110299190611716565b60206040518083038186803b15801561104157600080fd5b505afa158015611055573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061107991906114fb565b111561121157600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600354600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161113e9190611716565b60206040518083038186803b15801561115657600080fd5b505afa15801561116a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061118e91906114fb565b611198919061196f565b6040518363ffffffff1660e01b81526004016111b5929190611731565b602060405180830381600087803b1580156111cf57600080fd5b505af11580156111e3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611207919061144b565b50600090506112cc565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161126c9190611716565b60206040518083038186803b15801561128457600080fd5b505afa158015611298573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112bc91906114fb565b6003546112c9919061196f565b90505b90565b8280546112db90611a52565b90600052602060002090601f0160209004810192826112fd5760008555611344565b82601f1061131657805160ff1916838001178555611344565b82800160010185558215611344579182015b82811115611343578251825591602001919060010190611328565b5b5090506113519190611355565b5090565b5b8082111561136e576000816000905550600101611356565b5090565b600061138561138084611897565b611872565b9050828152602081018484840111156113a1576113a0611ba5565b5b6113ac848285611a10565b509392505050565b6000813590506113c381611d8d565b92915050565b6000813590506113d881611da4565b92915050565b6000815190506113ed81611da4565b92915050565b600082601f83011261140857611407611ba0565b5b8135611418848260208601611372565b91505092915050565b60008135905061143081611dbb565b92915050565b60008151905061144581611dbb565b92915050565b60006020828403121561146157611460611baf565b5b600061146f848285016113de565b91505092915050565b6000806000806080858703121561149257611491611baf565b5b600085013567ffffffffffffffff8111156114b0576114af611baa565b5b6114bc878288016113f3565b94505060206114cd878288016113b4565b93505060406114de878288016113b4565b92505060606114ef878288016113c9565b91505092959194509250565b60006020828403121561151157611510611baf565b5b600061151f84828501611436565b91505092915050565b600080600080600080600080610100898b03121561154957611548611baf565b5b60006115578b828c01611421565b98505060206115688b828c01611421565b97505060406115798b828c01611421565b965050606061158a8b828c01611421565b955050608061159b8b828c01611421565b94505060a06115ac8b828c016113b4565b93505060c06115bd8b828c016113b4565b92505060e06115ce8b828c016113b4565b9150509295985092959890939650565b6115e7816119a3565b82525050565b6115f6816119fe565b82525050565b6000611607826118c8565b61161181856118d3565b9350611621818560208601611a1f565b61162a81611bb4565b840191505092915050565b60006116426031836118d3565b915061164d82611bc5565b604082019050919050565b6000611665602e836118d3565b915061167082611c14565b604082019050919050565b60006116886025836118d3565b915061169382611c63565b604082019050919050565b60006116ab6023836118d3565b91506116b682611cb2565b604082019050919050565b60006116ce602e836118d3565b91506116d982611d01565b604082019050919050565b60006116f16017836118d3565b91506116fc82611d50565b602082019050919050565b611710816119f4565b82525050565b600060208201905061172b60008301846115de565b92915050565b600060408201905061174660008301856115de565b6117536020830184611707565b9392505050565b600060208201905061176f60008301846115ed565b92915050565b6000602082019050818103600083015261178f81846115fc565b905092915050565b600060208201905081810360008301526117b081611635565b9050919050565b600060208201905081810360008301526117d081611658565b9050919050565b600060208201905081810360008301526117f08161167b565b9050919050565b600060208201905081810360008301526118108161169e565b9050919050565b60006020820190508181036000830152611830816116c1565b9050919050565b60006020820190508181036000830152611850816116e4565b9050919050565b600060208201905061186c6000830184611707565b92915050565b600061187c61188d565b90506118888282611a84565b919050565b6000604051905090565b600067ffffffffffffffff8211156118b2576118b1611b71565b5b6118bb82611bb4565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60006118ef826119f4565b91506118fa836119f4565b92508261190a57611909611ae4565b5b828204905092915050565b6000611920826119f4565b915061192b836119f4565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561196457611963611ab5565b5b828202905092915050565b600061197a826119f4565b9150611985836119f4565b92508282101561199857611997611ab5565b5b828203905092915050565b60006119ae826119d4565b9050919050565b60008115159050919050565b60008190506119cf82611d79565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611a09826119c1565b9050919050565b82818337600083830152505050565b60005b83811015611a3d578082015181840152602081019050611a22565b83811115611a4c576000848401525b50505050565b60006002820490506001821680611a6a57607f821691505b60208210811415611a7e57611a7d611b42565b5b50919050565b611a8d82611bb4565b810181811067ffffffffffffffff82111715611aac57611aab611b71565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f746869732061646472657373206973206e6f7420616c6c6f77656420746f206360008201527f616c6c20746869732066756e6374696f6e000000000000000000000000000000602082015250565b7f6c756d6572696e20746f6b656e73206e65656420746f2062652073656e74207460008201527f6f2074686520636f6e7472616374000000000000000000000000000000000000602082015250565b7f636f6e7472616374206973206e6f7420696e20616e20617661696c61626c652060008201527f7374617465000000000000000000000000000000000000000000000000000000602082015250565b7f74686520636f6e747261637420686173206e6f74206265656e2070757263686160008201527f7365640000000000000000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f76616c69646174696f6e20666565206e6f742073656e74000000000000000000600082015250565b60048110611d8a57611d89611b13565b5b50565b611d96816119a3565b8114611da157600080fd5b50565b611dad816119b5565b8114611db857600080fd5b50565b611dc4816119f4565b8114611dcf57600080fd5b5056fea26469706673582212201169b9f3c2708f0876b5ed85c540ba658aa4bc068583e6d40ecb4eb1fe576f1f64736f6c63430008070033",
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
