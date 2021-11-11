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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"contractClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"contractFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"contractPurchased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractState\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"encryptedPoolData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_purchaser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validationFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lmn\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"port\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receivedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setContractCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFundContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_encryptedPoolData\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withValidator\",\"type\":\"bool\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"speed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611e3f806100206000396000f3fe6080604052600436106101145760003560e01c80638331ed49116100a0578063b6293b3411610064578063b6293b341461034e578063c20906ac14610365578063c5095d6814610390578063ce0c722a146103bb578063f8bbf27e146103e657610114565b80638331ed491461028657806385209ee0146102a25780638b7e4b13146102cd578063a035b1fe146102f8578063a4d66daf1461032357610114565b80631f7b6d32116100e75780631f7b6d32146101c557806326b3c68b146101f05780634897b9ac1461021b5780637150d8ae1461023257806374a348121461025d57610114565b806308551a5314610119578063089aa8a2146101445780630a61e2d91461016f57806316713b371461019a575b600080fd5b34801561012557600080fd5b5061012e610411565b60405161013b919061174d565b60405180910390f35b34801561015057600080fd5b50610159610437565b604051610166919061174d565b60405180910390f35b34801561017b57600080fd5b5061018461045d565b604051610191919061188e565b60405180910390f35b3480156101a657600080fd5b506101af610463565b6040516101bc919061188e565b60405180910390f35b3480156101d157600080fd5b506101da610469565b6040516101e7919061188e565b60405180910390f35b3480156101fc57600080fd5b5061020561046f565b604051610212919061188e565b60405180910390f35b34801561022757600080fd5b50610230610475565b005b34801561023e57600080fd5b506102476105db565b604051610254919061174d565b60405180910390f35b34801561026957600080fd5b50610284600480360381019061027f919061155f565b610601565b005b6102a0600480360381019061029b91906114af565b6107bb565b005b3480156102ae57600080fd5b506102b7610ad4565b6040516102c49190611791565b60405180910390f35b3480156102d957600080fd5b506102e2610ae7565b6040516102ef91906117ac565b60405180910390f35b34801561030457600080fd5b5061030d610b75565b60405161031a919061188e565b60405180910390f35b34801561032f57600080fd5b50610338610b7b565b604051610345919061188e565b60405180910390f35b34801561035a57600080fd5b50610363610b81565b005b34801561037157600080fd5b5061037a610ceb565b604051610387919061188e565b60405180910390f35b34801561039c57600080fd5b506103a5610cf1565b6040516103b2919061188e565b60405180910390f35b3480156103c757600080fd5b506103d0610cf7565b6040516103dd919061174d565b60405180910390f35b3480156103f257600080fd5b506103fb610d1d565b604051610408919061188e565b60405180910390f35b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035481565b60045481565b60095481565b600a5481565b60011515601260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514610508576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ff906117ce565b60405180910390fd5b6000600c544261051891906119a6565b9050600080600954831061052f576000915061055a565b6009548360095461054091906119a6565b60065461054d919061194c565b610557919061191b565b91505b8160065461056891906119a6565b90506105748183610d23565b6003600560146101000a81548160ff0219169083600381111561059a57610599611b4a565b5b02179055507faadd128c35976a01ffffa9dfb8d363b3358597ce6b30248bcf25e80bd3af4512336040516105ce919061174d565b60405180910390a1505050565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900460ff1680610627575060008054906101000a900460ff16155b610666576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065d9061184e565b60405180910390fd5b60008060019054906101000a900460ff1615905080156106b6576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b8860068190555087600781905550866008819055508560098190555083600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600b8190555082600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560146101000a81548160ff0219169083600381111561078157610780611b4a565b5b021790555061078f82610ecb565b80156107b05760008060016101000a81548160ff0219169083151502179055505b505050505050505050565b600060038111156107cf576107ce611b4a565b5b600560149054906101000a900460ff1660038111156107f1576107f0611b4a565b5b14610831576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108289061180e565b60405180910390fd5b8360119080519060200190610847929190611306565b5082600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061091b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600654610f72565b6001151581151514156109c5576001601260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600b5434146109c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109bb9061186e565b60405180910390fd5b5b6001601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160126000600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb33604051610ac6919061174d565b60405180910390a150505050565b600560149054906101000a900460ff1681565b60118054610af490611a89565b80601f0160208091040260200160405190810160405280929190818152602001828054610b2090611a89565b8015610b6d5780601f10610b4257610100808354040283529160200191610b6d565b820191906000526020600020905b815481529060010190602001808311610b5057829003601f168201915b505050505081565b60065481565b60075481565b6000610b8b611000565b14610bcb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc2906117ee565b60405180910390fd5b6001151560126000600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514610c80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c779061182e565b60405180910390fd5b42600c819055506002600560146101000a81548160ff02191690836003811115610cad57610cac611b4a565b5b02179055507f9a229752d1121ddc96ca24609b0837d8abded38d4749ac3048f1c71653bb9be033604051610ce1919061174d565b60405180910390a1565b60085481565b600c5481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600b5481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b8152600401610da2929190611768565b602060405180830381600087803b158015610dbc57600080fd5b505af1158015610dd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df49190611482565b50600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b8152600401610e74929190611768565b602060405180830381600087803b158015610e8e57600080fd5b505af1158015610ea2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ec69190611482565b505050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600381905550505050565b6000600354600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401611060919061174d565b60206040518083038186803b15801561107857600080fd5b505afa15801561108c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b09190611532565b111561124857600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600354600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401611175919061174d565b60206040518083038186803b15801561118d57600080fd5b505afa1580156111a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111c59190611532565b6111cf91906119a6565b6040518363ffffffff1660e01b81526004016111ec929190611768565b602060405180830381600087803b15801561120657600080fd5b505af115801561121a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123e9190611482565b5060009050611303565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016112a3919061174d565b60206040518083038186803b1580156112bb57600080fd5b505afa1580156112cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f39190611532565b60035461130091906119a6565b90505b90565b82805461131290611a89565b90600052602060002090601f016020900481019282611334576000855561137b565b82601f1061134d57805160ff191683800117855561137b565b8280016001018555821561137b579182015b8281111561137a57825182559160200191906001019061135f565b5b509050611388919061138c565b5090565b5b808211156113a557600081600090555060010161138d565b5090565b60006113bc6113b7846118ce565b6118a9565b9050828152602081018484840111156113d8576113d7611bdc565b5b6113e3848285611a47565b509392505050565b6000813590506113fa81611dc4565b92915050565b60008135905061140f81611ddb565b92915050565b60008151905061142481611ddb565b92915050565b600082601f83011261143f5761143e611bd7565b5b813561144f8482602086016113a9565b91505092915050565b60008135905061146781611df2565b92915050565b60008151905061147c81611df2565b92915050565b60006020828403121561149857611497611be6565b5b60006114a684828501611415565b91505092915050565b600080600080608085870312156114c9576114c8611be6565b5b600085013567ffffffffffffffff8111156114e7576114e6611be1565b5b6114f38782880161142a565b9450506020611504878288016113eb565b9350506040611515878288016113eb565b925050606061152687828801611400565b91505092959194509250565b60006020828403121561154857611547611be6565b5b60006115568482850161146d565b91505092915050565b600080600080600080600080610100898b0312156115805761157f611be6565b5b600061158e8b828c01611458565b985050602061159f8b828c01611458565b97505060406115b08b828c01611458565b96505060606115c18b828c01611458565b95505060806115d28b828c01611458565b94505060a06115e38b828c016113eb565b93505060c06115f48b828c016113eb565b92505060e06116058b828c016113eb565b9150509295985092959890939650565b61161e816119da565b82525050565b61162d81611a35565b82525050565b600061163e826118ff565b611648818561190a565b9350611658818560208601611a56565b61166181611beb565b840191505092915050565b600061167960318361190a565b915061168482611bfc565b604082019050919050565b600061169c602e8361190a565b91506116a782611c4b565b604082019050919050565b60006116bf60258361190a565b91506116ca82611c9a565b604082019050919050565b60006116e260238361190a565b91506116ed82611ce9565b604082019050919050565b6000611705602e8361190a565b915061171082611d38565b604082019050919050565b600061172860178361190a565b915061173382611d87565b602082019050919050565b61174781611a2b565b82525050565b60006020820190506117626000830184611615565b92915050565b600060408201905061177d6000830185611615565b61178a602083018461173e565b9392505050565b60006020820190506117a66000830184611624565b92915050565b600060208201905081810360008301526117c68184611633565b905092915050565b600060208201905081810360008301526117e78161166c565b9050919050565b600060208201905081810360008301526118078161168f565b9050919050565b60006020820190508181036000830152611827816116b2565b9050919050565b60006020820190508181036000830152611847816116d5565b9050919050565b60006020820190508181036000830152611867816116f8565b9050919050565b600060208201905081810360008301526118878161171b565b9050919050565b60006020820190506118a3600083018461173e565b92915050565b60006118b36118c4565b90506118bf8282611abb565b919050565b6000604051905090565b600067ffffffffffffffff8211156118e9576118e8611ba8565b5b6118f282611beb565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600061192682611a2b565b915061193183611a2b565b92508261194157611940611b1b565b5b828204905092915050565b600061195782611a2b565b915061196283611a2b565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561199b5761199a611aec565b5b828202905092915050565b60006119b182611a2b565b91506119bc83611a2b565b9250828210156119cf576119ce611aec565b5b828203905092915050565b60006119e582611a0b565b9050919050565b60008115159050919050565b6000819050611a0682611db0565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611a40826119f8565b9050919050565b82818337600083830152505050565b60005b83811015611a74578082015181840152602081019050611a59565b83811115611a83576000848401525b50505050565b60006002820490506001821680611aa157607f821691505b60208210811415611ab557611ab4611b79565b5b50919050565b611ac482611beb565b810181811067ffffffffffffffff82111715611ae357611ae2611ba8565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f746869732061646472657373206973206e6f7420616c6c6f77656420746f206360008201527f616c6c20746869732066756e6374696f6e000000000000000000000000000000602082015250565b7f6c756d6572696e20746f6b656e73206e65656420746f2062652073656e74207460008201527f6f2074686520636f6e7472616374000000000000000000000000000000000000602082015250565b7f636f6e7472616374206973206e6f7420696e20616e20617661696c61626c652060008201527f7374617465000000000000000000000000000000000000000000000000000000602082015250565b7f74686520636f6e747261637420686173206e6f74206265656e2070757263686160008201527f7365640000000000000000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f76616c69646174696f6e20666565206e6f742073656e74000000000000000000600082015250565b60048110611dc157611dc0611b4a565b5b50565b611dcd816119da565b8114611dd857600080fd5b50565b611de4816119ec565b8114611def57600080fd5b50565b611dfb81611a2b565b8114611e0657600080fd5b5056fea2646970667358221220346702b17e9de9d88c917d50d542115be435961c668dd66f685c97367c1c6fe164736f6c63430008070033",
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

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x8331ed49.
//
// Solidity: function setPurchaseContract(string _encryptedPoolData, address _buyer, address _validator, bool _withValidator) payable returns()
func (_Implementation *ImplementationTransactor) SetPurchaseContract(opts *bind.TransactOpts, _encryptedPoolData string, _buyer common.Address, _validator common.Address, _withValidator bool) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setPurchaseContract", _encryptedPoolData, _buyer, _validator, _withValidator)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x8331ed49.
//
// Solidity: function setPurchaseContract(string _encryptedPoolData, address _buyer, address _validator, bool _withValidator) payable returns()
func (_Implementation *ImplementationSession) SetPurchaseContract(_encryptedPoolData string, _buyer common.Address, _validator common.Address, _withValidator bool) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encryptedPoolData, _buyer, _validator, _withValidator)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x8331ed49.
//
// Solidity: function setPurchaseContract(string _encryptedPoolData, address _buyer, address _validator, bool _withValidator) payable returns()
func (_Implementation *ImplementationTransactorSession) SetPurchaseContract(_encryptedPoolData string, _buyer common.Address, _validator common.Address, _withValidator bool) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encryptedPoolData, _buyer, _validator, _withValidator)
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
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContractClosed is a free log retrieval operation binding the contract event 0xaadd128c35976a01ffffa9dfb8d363b3358597ce6b30248bcf25e80bd3af4512.
//
// Solidity: event contractClosed(address caller)
func (_Implementation *ImplementationFilterer) FilterContractClosed(opts *bind.FilterOpts) (*ImplementationContractClosedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractClosed")
	if err != nil {
		return nil, err
	}
	return &ImplementationContractClosedIterator{contract: _Implementation.contract, event: "contractClosed", logs: logs, sub: sub}, nil
}

// WatchContractClosed is a free log subscription operation binding the contract event 0xaadd128c35976a01ffffa9dfb8d363b3358597ce6b30248bcf25e80bd3af4512.
//
// Solidity: event contractClosed(address caller)
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

// ParseContractClosed is a log parse operation binding the contract event 0xaadd128c35976a01ffffa9dfb8d363b3358597ce6b30248bcf25e80bd3af4512.
//
// Solidity: event contractClosed(address caller)
func (_Implementation *ImplementationFilterer) ParseContractClosed(log types.Log) (*ImplementationContractClosed, error) {
	event := new(ImplementationContractClosed)
	if err := _Implementation.contract.UnpackLog(event, "contractClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationContractFundedIterator is returned from FilterContractFunded and is used to iterate over the raw logs and unpacked data for ContractFunded events raised by the Implementation contract.
type ImplementationContractFundedIterator struct {
	Event *ImplementationContractFunded // Event containing the contract specifics and raw log

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
func (it *ImplementationContractFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractFunded)
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
		it.Event = new(ImplementationContractFunded)
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
func (it *ImplementationContractFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractFunded represents a ContractFunded event raised by the Implementation contract.
type ImplementationContractFunded struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterContractFunded is a free log retrieval operation binding the contract event 0x9a229752d1121ddc96ca24609b0837d8abded38d4749ac3048f1c71653bb9be0.
//
// Solidity: event contractFunded(address caller)
func (_Implementation *ImplementationFilterer) FilterContractFunded(opts *bind.FilterOpts) (*ImplementationContractFundedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractFunded")
	if err != nil {
		return nil, err
	}
	return &ImplementationContractFundedIterator{contract: _Implementation.contract, event: "contractFunded", logs: logs, sub: sub}, nil
}

// WatchContractFunded is a free log subscription operation binding the contract event 0x9a229752d1121ddc96ca24609b0837d8abded38d4749ac3048f1c71653bb9be0.
//
// Solidity: event contractFunded(address caller)
func (_Implementation *ImplementationFilterer) WatchContractFunded(opts *bind.WatchOpts, sink chan<- *ImplementationContractFunded) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractFunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractFunded)
				if err := _Implementation.contract.UnpackLog(event, "contractFunded", log); err != nil {
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

// ParseContractFunded is a log parse operation binding the contract event 0x9a229752d1121ddc96ca24609b0837d8abded38d4749ac3048f1c71653bb9be0.
//
// Solidity: event contractFunded(address caller)
func (_Implementation *ImplementationFilterer) ParseContractFunded(log types.Log) (*ImplementationContractFunded, error) {
	event := new(ImplementationContractFunded)
	if err := _Implementation.contract.UnpackLog(event, "contractFunded", log); err != nil {
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
