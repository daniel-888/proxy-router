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
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"contractCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"contractClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"contractPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"dataEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkForDepositedTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractState\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentState\",\"outputs\":[{\"internalType\":\"enumEscrow.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dueAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_purchaser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lmn\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipaddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"password\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"port\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receivedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setContractCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setEarlyCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFundContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hashesCompleted\",\"type\":\"uint256\"}],\"name\":\"setPenaltyCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipaddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"speed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"username\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008060026101000a81548160ff0219169083600281111561003657610035610040565b5b021790555061006f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6124568061007e6000396000f3fe6080604052600436106101665760003560e01c80637150d8ae116100d1578063c20906ac1161008a578063d7cba96011610064578063d7cba960146104b4578063dcd161fb146104d0578063f3fc13f2146104f9578063f8bbf27e1461050357610166565b8063c20906ac14610435578063cc2de3ce14610460578063ce0c722a1461048957610166565b80637150d8ae1461034757806385209ee014610372578063a035b1fe1461039d578063a4d66daf146103c8578063b6293b34146103f3578063c20429651461040a57610166565b8063224b610b11610123578063224b610b1461026d57806326b3c68b1461029857806329dda763146102c357806339c1765a146102ee5780634897b9ac1461030557806361fac54d1461031c57610166565b806308551a531461016b578063089aa8a2146101965780630a61e2d9146101c15780630c3f6acf146101ec57806316713b37146102175780631f7b6d3214610242575b600080fd5b34801561017757600080fd5b5061018061052e565b60405161018d9190611b76565b60405180910390f35b3480156101a257600080fd5b506101ab610554565b6040516101b89190611b76565b60405180910390f35b3480156101cd57600080fd5b506101d661057a565b6040516101e39190611cb2565b60405180910390f35b3480156101f857600080fd5b50610201610580565b60405161020e9190611bd5565b60405180910390f35b34801561022357600080fd5b5061022c610593565b6040516102399190611cb2565b60405180910390f35b34801561024e57600080fd5b50610257610599565b6040516102649190611cb2565b60405180910390f35b34801561027957600080fd5b5061028261059f565b60405161028f9190611bf0565b60405180910390f35b3480156102a457600080fd5b506102ad61062d565b6040516102ba9190611cb2565b60405180910390f35b3480156102cf57600080fd5b506102d8610633565b6040516102e59190611bf0565b60405180910390f35b3480156102fa57600080fd5b506103036106c1565b005b34801561031157600080fd5b5061031a6107c3565b005b34801561032857600080fd5b506103316108b9565b60405161033e9190611bf0565b60405180910390f35b34801561035357600080fd5b5061035c610947565b6040516103699190611b76565b60405180910390f35b34801561037e57600080fd5b5061038761096d565b6040516103949190611bba565b60405180910390f35b3480156103a957600080fd5b506103b2610980565b6040516103bf9190611cb2565b60405180910390f35b3480156103d457600080fd5b506103dd610986565b6040516103ea9190611cb2565b60405180910390f35b3480156103ff57600080fd5b5061040861098c565b005b34801561041657600080fd5b5061041f61098e565b60405161042c9190611cb2565b60405180910390f35b34801561044157600080fd5b5061044a610a64565b6040516104579190611cb2565b60405180910390f35b34801561046c57600080fd5b5061048760048036038101906104829190611901565b610a6a565b005b34801561049557600080fd5b5061049e610d33565b6040516104ab9190611b76565b60405180910390f35b6104ce60048036038101906104c99190611800565b610d59565b005b3480156104dc57600080fd5b506104f760048036038101906104f291906118a7565b610f57565b005b610501611129565b005b34801561050f57600080fd5b50610518611204565b6040516105259190611cb2565b60405180910390f35b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060039054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b600060029054906101000a900460ff1681565b60055481565b600a5481565b601380546105ac9061200e565b80601f01602080910402602001604051908101604052809291908181526020018280546105d89061200e565b80156106255780601f106105fa57610100808354040283529160200191610625565b820191906000526020600020905b81548152906001019060200180831161060857829003601f168201915b505050505081565b600b5481565b601180546106409061200e565b80601f016020809104026020016040519081016040528092919081815260200182805461066c9061200e565b80156106b95780601f1061068e576101008083540402835291602001916106b9565b820191906000526020600020905b81548152906001019060200180831161069c57829003601f168201915b505050505081565b60011515601460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514610754576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074b90611c32565b60405180910390fd5b33601060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fcb3fbf71b2999f61d06483d265767337233ac476594db6b00a71ffa8306a1cb760405160405180910390a1565b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610853576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084a90611c12565b60405180910390fd5b610860600754600061120a565b6003600660146101000a81548160ff02191690836003811115610886576108856120cf565b5b02179055507ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a1565b601280546108c69061200e565b80601f01602080910402602001604051908101604052809291908181526020018280546108f29061200e565b801561093f5780601f106109145761010080835404028352916020019161093f565b820191906000526020600020905b81548152906001019060200180831161092257829003601f168201915b505050505081565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660149054906101000a900460ff1681565b60075481565b60085481565b565b600080600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016109ec9190611b76565b60206040518083038186803b158015610a0457600080fd5b505afa158015610a18573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3c91906118d4565b905060008114610a5c5760045481610a549190611f06565b915050610a61565b809150505b90565b60095481565b600060019054906101000a900460ff1680610a90575060008054906101000a900460ff16155b610acf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac690611c72565b60405180910390fd5b60008060019054906101000a900460ff161590508015610b1f576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b87600781905550866008819055508560098190555084600a8190555083600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600660146101000a81548160ff02191690836003811115610be357610be26120cf565b5b0217905550600160146000600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600160146000600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610d08600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168361148b565b8015610d295760008060016101000a81548160ff0219169083151502179055505b5050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600c543414610d9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9490611c92565b60405180910390fd5b60006003811115610db157610db06120cf565b5b600660149054906101000a900460ff166003811115610dd357610dd26120cf565b5b14610e13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e0a90611c52565b60405180910390fd5b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610e7b573d6000803e3d6000fd5b508260119080519060200190610e9292919061166c565b508160129080519060200190610ea992919061166c565b508060139080519060200190610ec092919061166c565b506002600660146101000a81548160ff02191690836003811115610ee757610ee66120cf565b5b0217905550610f1b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1633600754611574565b7f0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb33604051610f4a9190611b76565b60405180910390a1505050565b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fe7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fde90611c12565b60405180910390fd5b600060085482600854610ffa9190611f06565b6007546110079190611eac565b6110119190611e7b565b90506000600854836007546110269190611eac565b6110309190611e7b565b9050600061103e8383611639565b6110488484611652565b6110529190611f06565b9050600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156110df5780826110d89190611e25565b91506110ee565b80836110eb9190611e25565b92505b6110f8838361120a565b6003600660146101000a81548160ff0219169083600381111561111e5761111d6120cf565b5b021790555050505050565b600061113361098e565b14156111a0576002600060026101000a81548160ff0219169083600281111561115f5761115e6120cf565b5b02179055507f88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb426040516111939190611cfb565b60405180910390a1611202565b60008060026101000a81548160ff021916908360028111156111c5576111c46120cf565b5b02179055507f88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb426040516111f99190611d29565b60405180910390a15b565b600c5481565b60028081111561121d5761121c6120cf565b5b600060029054906101000a900460ff16600281111561123f5761123e6120cf565b5b14611280577f88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb426040516112739190611d57565b60405180910390a1611487565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b81526004016112ff929190611b91565b602060405180830381600087803b15801561131957600080fd5b505af115801561132d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061135191906117d3565b50600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060039054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016113d1929190611b91565b602060405180830381600087803b1580156113eb57600080fd5b505af11580156113ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061142391906117d3565b506001600060026101000a81548160ff0219169083600281111561144a576114496120cf565b5b02179055507f88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb4260405161147e9190611d85565b60405180910390a15b5050565b81600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600060036101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806004819055507f88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb4260405161162c9190611ccd565b60405180910390a1505050565b6000818310611648578161164a565b825b905092915050565b6000818310156116625781611664565b825b905092915050565b8280546116789061200e565b90600052602060002090601f01602090048101928261169a57600085556116e1565b82601f106116b357805160ff19168380011785556116e1565b828001600101855582156116e1579182015b828111156116e05782518255916020019190600101906116c5565b5b5090506116ee91906116f2565b5090565b5b8082111561170b5760008160009055506001016116f3565b5090565b600061172261171d84611dd8565b611db3565b90508281526020810184848401111561173e5761173d612161565b5b611749848285611fcc565b509392505050565b600081359050611760816123db565b92915050565b600081519050611775816123f2565b92915050565b600082601f8301126117905761178f61215c565b5b81356117a084826020860161170f565b91505092915050565b6000813590506117b881612409565b92915050565b6000815190506117cd81612409565b92915050565b6000602082840312156117e9576117e861216b565b5b60006117f784828501611766565b91505092915050565b6000806000606084860312156118195761181861216b565b5b600084013567ffffffffffffffff81111561183757611836612166565b5b6118438682870161177b565b935050602084013567ffffffffffffffff81111561186457611863612166565b5b6118708682870161177b565b925050604084013567ffffffffffffffff81111561189157611890612166565b5b61189d8682870161177b565b9150509250925092565b6000602082840312156118bd576118bc61216b565b5b60006118cb848285016117a9565b91505092915050565b6000602082840312156118ea576118e961216b565b5b60006118f8848285016117be565b91505092915050565b600080600080600080600060e0888a0312156119205761191f61216b565b5b600061192e8a828b016117a9565b975050602061193f8a828b016117a9565b96505060406119508a828b016117a9565b95505060606119618a828b016117a9565b94505060806119728a828b01611751565b93505060a06119838a828b01611751565b92505060c06119948a828b01611751565b91505092959891949750929550565b6119ac81611f3a565b82525050565b6119bb81611fa8565b82525050565b6119ca81611fba565b82525050565b60006119db82611e09565b6119e58185611e14565b93506119f5818560208601611fdb565b6119fe81612170565b840191505092915050565b6000611a16600e83611e14565b9150611a2182612181565b602082019050919050565b6000611a39603483611e14565b9150611a44826121aa565b604082019050919050565b6000611a5c603183611e14565b9150611a67826121f9565b604082019050919050565b6000611a7f601683611e14565b9150611a8a82612248565b602082019050919050565b6000611aa2602583611e14565b9150611aad82612271565b604082019050919050565b6000611ac5602e83611e14565b9150611ad0826122c0565b604082019050919050565b6000611ae8601d83611e14565b9150611af38261230f565b602082019050919050565b6000611b0b601b83611e14565b9150611b1682612338565b602082019050919050565b6000611b2e601883611e14565b9150611b3982612361565b602082019050919050565b6000611b51601283611e14565b9150611b5c8261238a565b602082019050919050565b611b7081611f9e565b82525050565b6000602082019050611b8b60008301846119a3565b92915050565b6000604082019050611ba660008301856119a3565b611bb36020830184611b67565b9392505050565b6000602082019050611bcf60008301846119b2565b92915050565b6000602082019050611bea60008301846119c1565b92915050565b60006020820190508181036000830152611c0a81846119d0565b905092915050565b60006020820190508181036000830152611c2b81611a2c565b9050919050565b60006020820190508181036000830152611c4b81611a4f565b9050919050565b60006020820190508181036000830152611c6b81611a95565b9050919050565b60006020820190508181036000830152611c8b81611ab8565b9050919050565b60006020820190508181036000830152611cab81611afe565b9050919050565b6000602082019050611cc76000830184611b67565b92915050565b6000604082019050611ce26000830184611b67565b8181036020830152611cf381611a09565b905092915050565b6000604082019050611d106000830184611b67565b8181036020830152611d2181611a72565b905092915050565b6000604082019050611d3e6000830184611b67565b8181036020830152611d4f81611adb565b905092915050565b6000604082019050611d6c6000830184611b67565b8181036020830152611d7d81611b21565b905092915050565b6000604082019050611d9a6000830184611b67565b8181036020830152611dab81611b44565b905092915050565b6000611dbd611dce565b9050611dc98282612040565b919050565b6000604051905090565b600067ffffffffffffffff821115611df357611df261212d565b5b611dfc82612170565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000611e3082611f9e565b9150611e3b83611f9e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611e7057611e6f612071565b5b828201905092915050565b6000611e8682611f9e565b9150611e9183611f9e565b925082611ea157611ea06120a0565b5b828204905092915050565b6000611eb782611f9e565b9150611ec283611f9e565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611efb57611efa612071565b5b828202905092915050565b6000611f1182611f9e565b9150611f1c83611f9e565b925082821015611f2f57611f2e612071565b5b828203905092915050565b6000611f4582611f7e565b9050919050565b60008115159050919050565b6000819050611f66826123b3565b919050565b6000819050611f79826123c7565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611fb382611f58565b9050919050565b6000611fc582611f6b565b9050919050565b82818337600083830152505050565b60005b83811015611ff9578082015181840152602081019050611fde565b83811115612008576000848401525b50505050565b6000600282049050600182168061202657607f821691505b6020821081141561203a576120396120fe565b5b50919050565b61204982612170565b810181811067ffffffffffffffff821117156120685761206761212d565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f457363726f772043726561746564000000000000000000000000000000000000600082015250565b7f746869732066756e6374696f6e206d7573742062652063616c6c65642062792060008201527f74686520636f6e7472616374206d616e61676572000000000000000000000000602082015250565b7f746869732061646472657373206973206e6f7420616c6c6f77656420746f206360008201527f616c6c20746869732066756e6374696f6e000000000000000000000000000000602082015250565b7f436f6e74726163742066756c6c792066756e6465642100000000000000000000600082015250565b7f636f6e7472616374206973206e6f7420696e20616e20617661696c61626c652060008201527f7374617465000000000000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f436f6e7472616374206973206e6f742066756c6c792066756e64656421000000600082015250565b7f76616c69646174696f6e2066656520697320696e636f72726563740000000000600082015250565b7f4572726f722c206e6f742066756c6c792066756e646564210000000000000000600082015250565b7f436f6e747261637420436f6d706c657465640000000000000000000000000000600082015250565b600481106123c4576123c36120cf565b5b50565b600381106123d8576123d76120cf565b5b50565b6123e481611f3a565b81146123ef57600080fd5b50565b6123fb81611f4c565b811461240657600080fd5b50565b61241281611f9e565b811461241d57600080fd5b5056fea264697066735822122090b96576a5b9db0a85f3afc1982c3e81ef383be4ac149d5712b9a36bf597dfa064736f6c63430008070033",
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

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() view returns(uint8)
func (_Implementation *ImplementationCaller) CurrentState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "currentState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() view returns(uint8)
func (_Implementation *ImplementationSession) CurrentState() (uint8, error) {
	return _Implementation.Contract.CurrentState(&_Implementation.CallOpts)
}

// CurrentState is a free data retrieval call binding the contract method 0x0c3f6acf.
//
// Solidity: function currentState() view returns(uint8)
func (_Implementation *ImplementationCallerSession) CurrentState() (uint8, error) {
	return _Implementation.Contract.CurrentState(&_Implementation.CallOpts)
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

// Ipaddress is a free data retrieval call binding the contract method 0x29dda763.
//
// Solidity: function ipaddress() view returns(string)
func (_Implementation *ImplementationCaller) Ipaddress(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "ipaddress")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Ipaddress is a free data retrieval call binding the contract method 0x29dda763.
//
// Solidity: function ipaddress() view returns(string)
func (_Implementation *ImplementationSession) Ipaddress() (string, error) {
	return _Implementation.Contract.Ipaddress(&_Implementation.CallOpts)
}

// Ipaddress is a free data retrieval call binding the contract method 0x29dda763.
//
// Solidity: function ipaddress() view returns(string)
func (_Implementation *ImplementationCallerSession) Ipaddress() (string, error) {
	return _Implementation.Contract.Ipaddress(&_Implementation.CallOpts)
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

// Password is a free data retrieval call binding the contract method 0x224b610b.
//
// Solidity: function password() view returns(string)
func (_Implementation *ImplementationCaller) Password(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "password")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Password is a free data retrieval call binding the contract method 0x224b610b.
//
// Solidity: function password() view returns(string)
func (_Implementation *ImplementationSession) Password() (string, error) {
	return _Implementation.Contract.Password(&_Implementation.CallOpts)
}

// Password is a free data retrieval call binding the contract method 0x224b610b.
//
// Solidity: function password() view returns(string)
func (_Implementation *ImplementationCallerSession) Password() (string, error) {
	return _Implementation.Contract.Password(&_Implementation.CallOpts)
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

// Username is a free data retrieval call binding the contract method 0x61fac54d.
//
// Solidity: function username() view returns(string)
func (_Implementation *ImplementationCaller) Username(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "username")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Username is a free data retrieval call binding the contract method 0x61fac54d.
//
// Solidity: function username() view returns(string)
func (_Implementation *ImplementationSession) Username() (string, error) {
	return _Implementation.Contract.Username(&_Implementation.CallOpts)
}

// Username is a free data retrieval call binding the contract method 0x61fac54d.
//
// Solidity: function username() view returns(string)
func (_Implementation *ImplementationCallerSession) Username() (string, error) {
	return _Implementation.Contract.Username(&_Implementation.CallOpts)
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

// CheckForDepositedTokens is a paid mutator transaction binding the contract method 0xf3fc13f2.
//
// Solidity: function checkForDepositedTokens() payable returns()
func (_Implementation *ImplementationTransactor) CheckForDepositedTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "checkForDepositedTokens")
}

// CheckForDepositedTokens is a paid mutator transaction binding the contract method 0xf3fc13f2.
//
// Solidity: function checkForDepositedTokens() payable returns()
func (_Implementation *ImplementationSession) CheckForDepositedTokens() (*types.Transaction, error) {
	return _Implementation.Contract.CheckForDepositedTokens(&_Implementation.TransactOpts)
}

// CheckForDepositedTokens is a paid mutator transaction binding the contract method 0xf3fc13f2.
//
// Solidity: function checkForDepositedTokens() payable returns()
func (_Implementation *ImplementationTransactorSession) CheckForDepositedTokens() (*types.Transaction, error) {
	return _Implementation.Contract.CheckForDepositedTokens(&_Implementation.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2de3ce.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _seller, address _contractManager, address _lmn) returns()
func (_Implementation *ImplementationTransactor) Initialize(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _seller common.Address, _contractManager common.Address, _lmn common.Address) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "initialize", _price, _limit, _speed, _length, _seller, _contractManager, _lmn)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2de3ce.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _seller, address _contractManager, address _lmn) returns()
func (_Implementation *ImplementationSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _seller common.Address, _contractManager common.Address, _lmn common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _seller, _contractManager, _lmn)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2de3ce.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _seller, address _contractManager, address _lmn) returns()
func (_Implementation *ImplementationTransactorSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _seller common.Address, _contractManager common.Address, _lmn common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _seller, _contractManager, _lmn)
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

// SetEarlyCloseOut is a paid mutator transaction binding the contract method 0x39c1765a.
//
// Solidity: function setEarlyCloseOut() returns()
func (_Implementation *ImplementationTransactor) SetEarlyCloseOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setEarlyCloseOut")
}

// SetEarlyCloseOut is a paid mutator transaction binding the contract method 0x39c1765a.
//
// Solidity: function setEarlyCloseOut() returns()
func (_Implementation *ImplementationSession) SetEarlyCloseOut() (*types.Transaction, error) {
	return _Implementation.Contract.SetEarlyCloseOut(&_Implementation.TransactOpts)
}

// SetEarlyCloseOut is a paid mutator transaction binding the contract method 0x39c1765a.
//
// Solidity: function setEarlyCloseOut() returns()
func (_Implementation *ImplementationTransactorSession) SetEarlyCloseOut() (*types.Transaction, error) {
	return _Implementation.Contract.SetEarlyCloseOut(&_Implementation.TransactOpts)
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

// SetPenaltyCloseOut is a paid mutator transaction binding the contract method 0xdcd161fb.
//
// Solidity: function setPenaltyCloseOut(uint256 _hashesCompleted) returns()
func (_Implementation *ImplementationTransactor) SetPenaltyCloseOut(opts *bind.TransactOpts, _hashesCompleted *big.Int) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setPenaltyCloseOut", _hashesCompleted)
}

// SetPenaltyCloseOut is a paid mutator transaction binding the contract method 0xdcd161fb.
//
// Solidity: function setPenaltyCloseOut(uint256 _hashesCompleted) returns()
func (_Implementation *ImplementationSession) SetPenaltyCloseOut(_hashesCompleted *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetPenaltyCloseOut(&_Implementation.TransactOpts, _hashesCompleted)
}

// SetPenaltyCloseOut is a paid mutator transaction binding the contract method 0xdcd161fb.
//
// Solidity: function setPenaltyCloseOut(uint256 _hashesCompleted) returns()
func (_Implementation *ImplementationTransactorSession) SetPenaltyCloseOut(_hashesCompleted *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetPenaltyCloseOut(&_Implementation.TransactOpts, _hashesCompleted)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0xd7cba960.
//
// Solidity: function setPurchaseContract(string _ipaddress, string _username, string _password) payable returns()
func (_Implementation *ImplementationTransactor) SetPurchaseContract(opts *bind.TransactOpts, _ipaddress string, _username string, _password string) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setPurchaseContract", _ipaddress, _username, _password)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0xd7cba960.
//
// Solidity: function setPurchaseContract(string _ipaddress, string _username, string _password) payable returns()
func (_Implementation *ImplementationSession) SetPurchaseContract(_ipaddress string, _username string, _password string) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _ipaddress, _username, _password)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0xd7cba960.
//
// Solidity: function setPurchaseContract(string _ipaddress, string _username, string _password) payable returns()
func (_Implementation *ImplementationTransactorSession) SetPurchaseContract(_ipaddress string, _username string, _password string) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _ipaddress, _username, _password)
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

// ImplementationDataEventIterator is returned from FilterDataEvent and is used to iterate over the raw logs and unpacked data for DataEvent events raised by the Implementation contract.
type ImplementationDataEventIterator struct {
	Event *ImplementationDataEvent // Event containing the contract specifics and raw log

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
func (it *ImplementationDataEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationDataEvent)
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
		it.Event = new(ImplementationDataEvent)
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
func (it *ImplementationDataEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationDataEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationDataEvent represents a DataEvent event raised by the Implementation contract.
type ImplementationDataEvent struct {
	Date *big.Int
	Val  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDataEvent is a free log retrieval operation binding the contract event 0x88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb.
//
// Solidity: event dataEvent(uint256 date, string val)
func (_Implementation *ImplementationFilterer) FilterDataEvent(opts *bind.FilterOpts) (*ImplementationDataEventIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "dataEvent")
	if err != nil {
		return nil, err
	}
	return &ImplementationDataEventIterator{contract: _Implementation.contract, event: "dataEvent", logs: logs, sub: sub}, nil
}

// WatchDataEvent is a free log subscription operation binding the contract event 0x88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb.
//
// Solidity: event dataEvent(uint256 date, string val)
func (_Implementation *ImplementationFilterer) WatchDataEvent(opts *bind.WatchOpts, sink chan<- *ImplementationDataEvent) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "dataEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationDataEvent)
				if err := _Implementation.contract.UnpackLog(event, "dataEvent", log); err != nil {
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

// ParseDataEvent is a log parse operation binding the contract event 0x88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb.
//
// Solidity: event dataEvent(uint256 date, string val)
func (_Implementation *ImplementationFilterer) ParseDataEvent(log types.Log) (*ImplementationDataEvent, error) {
	event := new(ImplementationDataEvent)
	if err := _Implementation.contract.UnpackLog(event, "dataEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
