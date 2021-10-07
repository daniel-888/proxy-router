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
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"contractCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"contractClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"contractPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"val\",\"type\":\"string\"}],\"name\":\"dataEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractState\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentState\",\"outputs\":[{\"internalType\":\"enumEscrow.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dueAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_purchaser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow_seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lmn\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipaddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"length\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"password\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"port\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receivedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setContractCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setEarlyCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFundContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hashesCompleted\",\"type\":\"uint256\"}],\"name\":\"setPenaltyCloseOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ipaddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_username\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"speed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"username\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008060026101000a81548160ff0219169083600281111561003657610035610040565b5b021790555061006f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b61213b8061007e6000396000f3fe60806040526004361061014b5760003560e01c806361fac54d116100b6578063c20429651161006f578063c20429651461040b578063c20906ac14610436578063cc2de3ce14610461578063ce0c722a1461048a578063dcd161fb146104b5578063f8bbf27e146104de5761014b565b806361fac54d1461031d5780637150d8ae1461034857806385209ee014610373578063a035b1fe1461039e578063a4d66daf146103c9578063b6293b34146103f45761014b565b8063224b610b11610108578063224b610b1461025257806326b3c68b1461027d57806329dda763146102a85780632cf76de9146102d357806339c1765a146102ef5780634897b9ac146103065761014b565b806308551a5314610150578063089aa8a21461017b5780630a61e2d9146101a65780630c3f6acf146101d157806316713b37146101fc5780631f7b6d3214610227575b600080fd5b34801561015c57600080fd5b50610165610509565b604051610172919061193a565b60405180910390f35b34801561018757600080fd5b5061019061052f565b60405161019d919061193a565b60405180910390f35b3480156101b257600080fd5b506101bb610555565b6040516101c89190611a76565b60405180910390f35b3480156101dd57600080fd5b506101e661055b565b6040516101f39190611999565b60405180910390f35b34801561020857600080fd5b5061021161056e565b60405161021e9190611a76565b60405180910390f35b34801561023357600080fd5b5061023c610574565b6040516102499190611a76565b60405180910390f35b34801561025e57600080fd5b5061026761057a565b60405161027491906119b4565b60405180910390f35b34801561028957600080fd5b50610292610608565b60405161029f9190611a76565b60405180910390f35b3480156102b457600080fd5b506102bd61060e565b6040516102ca91906119b4565b60405180910390f35b6102ed60048036038101906102e89190611619565b61069c565b005b3480156102fb57600080fd5b5061030461087e565b005b34801561031257600080fd5b5061031b610980565b005b34801561032957600080fd5b50610332610a76565b60405161033f91906119b4565b60405180910390f35b34801561035457600080fd5b5061035d610b04565b60405161036a919061193a565b60405180910390f35b34801561037f57600080fd5b50610388610b2a565b604051610395919061197e565b60405180910390f35b3480156103aa57600080fd5b506103b3610b3d565b6040516103c09190611a76565b60405180910390f35b3480156103d557600080fd5b506103de610b43565b6040516103eb9190611a76565b60405180910390f35b34801561040057600080fd5b50610409610b49565b005b34801561041757600080fd5b50610420610bc0565b60405161042d9190611a76565b60405180910390f35b34801561044257600080fd5b5061044b610c7f565b6040516104589190611a76565b60405180910390f35b34801561046d57600080fd5b506104886004803603810190610483919061172e565b610c85565b005b34801561049657600080fd5b5061049f610ed4565b6040516104ac919061193a565b60405180910390f35b3480156104c157600080fd5b506104dc60048036038101906104d791906116d4565b610efa565b005b3480156104ea57600080fd5b506104f36110c7565b6040516105009190611a76565b60405180910390f35b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060039054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b600060029054906101000a900460ff1681565b60055481565b600a5481565b6013805461058790611d48565b80601f01602080910402602001604051908101604052809291908181526020018280546105b390611d48565b80156106005780601f106105d557610100808354040283529160200191610600565b820191906000526020600020905b8154815290600101906020018083116105e357829003601f168201915b505050505081565b600b5481565b6011805461061b90611d48565b80601f016020809104026020016040519081016040528092919081815260200182805461064790611d48565b80156106945780601f1061066957610100808354040283529160200191610694565b820191906000526020600020905b81548152906001019060200180831161067757829003601f168201915b505050505081565b600060038111156106b0576106af611e09565b5b600660149054906101000a900460ff1660038111156106d2576106d1611e09565b5b14610712576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070990611a36565b60405180910390fd5b8360119080519060200190610728929190611485565b50826012908051906020019061073f929190611485565b508160139080519060200190610756929190611485565b5080600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506107e9600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007546110cd565b6001601460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb33604051610870919061193a565b60405180910390a150505050565b60011515601460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514610911576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610908906119f6565b60405180910390fd5b33601060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fcb3fbf71b2999f61d06483d265767337233ac476594db6b00a71ffa8306a1cb760405160405180910390a1565b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a07906119d6565b60405180910390fd5b610a1d6007546000611192565b6003600660146101000a81548160ff02191690836003811115610a4357610a42611e09565b5b02179055507ff5e1a452bb76d7335225182a97ad694be2c7b4b5d75dcffb67ddf15db95f484460405160405180910390a1565b60128054610a8390611d48565b80601f0160208091040260200160405190810160405280929190818152602001828054610aaf90611d48565b8015610afc5780601f10610ad157610100808354040283529160200191610afc565b820191906000526020600020905b815481529060010190602001808311610adf57829003601f168201915b505050505081565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660149054906101000a900460ff1681565b60075481565b60085481565b6000610b53610bc0565b14610b93576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8a90611a16565b60405180910390fd5b6002600660146101000a81548160ff02191690836003811115610bb957610bb8611e09565b5b0217905550565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610c1d919061193a565b60206040518083038186803b158015610c3557600080fd5b505afa158015610c49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6d9190611701565b600454610c7a9190611c40565b905090565b60095481565b600060019054906101000a900460ff1680610cab575060008054906101000a900460ff16155b610cea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce190611a56565b60405180910390fd5b60008060019054906101000a900460ff161590508015610d3a576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b87600781905550866008819055508560098190555084600a8190555083600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600660146101000a81548160ff02191690836003811115610dfe57610dfd611e09565b5b0217905550600160146000600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610ea9600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168361139c565b8015610eca5760008060016101000a81548160ff0219169083151502179055505b5050505050505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610f8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f81906119d6565b60405180910390fd5b600060085482600854610f9d9190611c40565b600754610faa9190611be6565b610fb49190611bb5565b9050600060085483600754610fc99190611be6565b610fd39190611bb5565b905060008183610fe39190611b5f565b600754610ff09190611c40565b9050600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16601060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561107d5780826110769190611b5f565b915061108c565b80836110899190611b5f565b92505b6110968383611192565b6003600660146101000a81548160ff021916908360038111156110bc576110bb611e09565b5b021790555050505050565b600c5481565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600060036101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806004819055507f88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb426040516111859190611a91565b60405180910390a1505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b8152600401611211929190611955565b602060405180830381600087803b15801561122b57600080fd5b505af115801561123f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061126391906115ec565b50600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060039054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016112e3929190611955565b602060405180830381600087803b1580156112fd57600080fd5b505af1158015611311573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061133591906115ec565b506001600060026101000a81548160ff0219169083600281111561135c5761135b611e09565b5b02179055507f88265e9be093ab2ee66f829ad3ca909591f25cd6685323b555215283e78148eb426040516113909190611abf565b60405180910390a15050565b81600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b82805461149190611d48565b90600052602060002090601f0160209004810192826114b357600085556114fa565b82601f106114cc57805160ff19168380011785556114fa565b828001600101855582156114fa579182015b828111156114f95782518255916020019190600101906114de565b5b509050611507919061150b565b5090565b5b8082111561152457600081600090555060010161150c565b5090565b600061153b61153684611b12565b611aed565b90508281526020810184848401111561155757611556611e9b565b5b611562848285611d06565b509392505050565b600081359050611579816120c0565b92915050565b60008151905061158e816120d7565b92915050565b600082601f8301126115a9576115a8611e96565b5b81356115b9848260208601611528565b91505092915050565b6000813590506115d1816120ee565b92915050565b6000815190506115e6816120ee565b92915050565b60006020828403121561160257611601611ea5565b5b60006116108482850161157f565b91505092915050565b6000806000806080858703121561163357611632611ea5565b5b600085013567ffffffffffffffff81111561165157611650611ea0565b5b61165d87828801611594565b945050602085013567ffffffffffffffff81111561167e5761167d611ea0565b5b61168a87828801611594565b935050604085013567ffffffffffffffff8111156116ab576116aa611ea0565b5b6116b787828801611594565b92505060606116c88782880161156a565b91505092959194509250565b6000602082840312156116ea576116e9611ea5565b5b60006116f8848285016115c2565b91505092915050565b60006020828403121561171757611716611ea5565b5b6000611725848285016115d7565b91505092915050565b600080600080600080600060e0888a03121561174d5761174c611ea5565b5b600061175b8a828b016115c2565b975050602061176c8a828b016115c2565b965050604061177d8a828b016115c2565b955050606061178e8a828b016115c2565b945050608061179f8a828b0161156a565b93505060a06117b08a828b0161156a565b92505060c06117c18a828b0161156a565b91505092959891949750929550565b6117d981611c74565b82525050565b6117e881611ce2565b82525050565b6117f781611cf4565b82525050565b600061180882611b43565b6118128185611b4e565b9350611822818560208601611d15565b61182b81611eaa565b840191505092915050565b6000611843600e83611b4e565b915061184e82611ebb565b602082019050919050565b6000611866603483611b4e565b915061187182611ee4565b604082019050919050565b6000611889603183611b4e565b915061189482611f33565b604082019050919050565b60006118ac602e83611b4e565b91506118b782611f82565b604082019050919050565b60006118cf602583611b4e565b91506118da82611fd1565b604082019050919050565b60006118f2602e83611b4e565b91506118fd82612020565b604082019050919050565b6000611915601283611b4e565b91506119208261206f565b602082019050919050565b61193481611cd8565b82525050565b600060208201905061194f60008301846117d0565b92915050565b600060408201905061196a60008301856117d0565b611977602083018461192b565b9392505050565b600060208201905061199360008301846117df565b92915050565b60006020820190506119ae60008301846117ee565b92915050565b600060208201905081810360008301526119ce81846117fd565b905092915050565b600060208201905081810360008301526119ef81611859565b9050919050565b60006020820190508181036000830152611a0f8161187c565b9050919050565b60006020820190508181036000830152611a2f8161189f565b9050919050565b60006020820190508181036000830152611a4f816118c2565b9050919050565b60006020820190508181036000830152611a6f816118e5565b9050919050565b6000602082019050611a8b600083018461192b565b92915050565b6000604082019050611aa6600083018461192b565b8181036020830152611ab781611836565b905092915050565b6000604082019050611ad4600083018461192b565b8181036020830152611ae581611908565b905092915050565b6000611af7611b08565b9050611b038282611d7a565b919050565b6000604051905090565b600067ffffffffffffffff821115611b2d57611b2c611e67565b5b611b3682611eaa565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000611b6a82611cd8565b9150611b7583611cd8565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611baa57611ba9611dab565b5b828201905092915050565b6000611bc082611cd8565b9150611bcb83611cd8565b925082611bdb57611bda611dda565b5b828204905092915050565b6000611bf182611cd8565b9150611bfc83611cd8565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611c3557611c34611dab565b5b828202905092915050565b6000611c4b82611cd8565b9150611c5683611cd8565b925082821015611c6957611c68611dab565b5b828203905092915050565b6000611c7f82611cb8565b9050919050565b60008115159050919050565b6000819050611ca082612098565b919050565b6000819050611cb3826120ac565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611ced82611c92565b9050919050565b6000611cff82611ca5565b9050919050565b82818337600083830152505050565b60005b83811015611d33578082015181840152602081019050611d18565b83811115611d42576000848401525b50505050565b60006002820490506001821680611d6057607f821691505b60208210811415611d7457611d73611e38565b5b50919050565b611d8382611eaa565b810181811067ffffffffffffffff82111715611da257611da1611e67565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f457363726f772043726561746564000000000000000000000000000000000000600082015250565b7f746869732066756e6374696f6e206d7573742062652063616c6c65642062792060008201527f74686520636f6e7472616374206d616e61676572000000000000000000000000602082015250565b7f746869732061646472657373206973206e6f7420616c6c6f77656420746f206360008201527f616c6c20746869732066756e6374696f6e000000000000000000000000000000602082015250565b7f6c756d6572696e20746f6b656e73206e65656420746f2062652073656e74207460008201527f6f2074686520636f6e7472616374000000000000000000000000000000000000602082015250565b7f636f6e7472616374206973206e6f7420696e20616e20617661696c61626c652060008201527f7374617465000000000000000000000000000000000000000000000000000000602082015250565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b7f436f6e747261637420436f6d706c657465640000000000000000000000000000600082015250565b600481106120a9576120a8611e09565b5b50565b600381106120bd576120bc611e09565b5b50565b6120c981611c74565b81146120d457600080fd5b50565b6120e081611c86565b81146120eb57600080fd5b50565b6120f781611cd8565b811461210257600080fd5b5056fea264697066735822122034efbcdd7ac15c16d20f80e6578bd1a2ce9caa5b4992364e2323667860c18c6764736f6c63430008070033",
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

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x2cf76de9.
//
// Solidity: function setPurchaseContract(string _ipaddress, string _username, string _password, address _buyer) payable returns()
func (_Implementation *ImplementationTransactor) SetPurchaseContract(opts *bind.TransactOpts, _ipaddress string, _username string, _password string, _buyer common.Address) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setPurchaseContract", _ipaddress, _username, _password, _buyer)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x2cf76de9.
//
// Solidity: function setPurchaseContract(string _ipaddress, string _username, string _password, address _buyer) payable returns()
func (_Implementation *ImplementationSession) SetPurchaseContract(_ipaddress string, _username string, _password string, _buyer common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _ipaddress, _username, _password, _buyer)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x2cf76de9.
//
// Solidity: function setPurchaseContract(string _ipaddress, string _username, string _password, address _buyer) payable returns()
func (_Implementation *ImplementationTransactorSession) SetPurchaseContract(_ipaddress string, _username string, _password string, _buyer common.Address) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _ipaddress, _username, _password, _buyer)
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
