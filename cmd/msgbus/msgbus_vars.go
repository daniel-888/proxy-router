package msgbus

type ContractState string
type MinerState string

const (
	ContAvailableState ContractState = "AvailableState"
	ContRunningState   ContractState = "RunningState"
)

// Need to figure out the IDString for this, for now it is just a string
type IDString string
type ConfigID IDString
type SellerID IDString
type BuyerID IDString
type ContractID IDString

// Do we still need this with the config package in place?
type ConfigInfo struct {
	ID          ConfigID
	DefaultDest DestID
	Seller      SellerID
}

type Seller struct {
	ID                     SellerID
	DefaultDest            DestID
	TotalAvailableHashRate int
	UnusedHashRate         int
	Contracts       	   map[ContractID]ContractState
}

type Buyer struct {
	ID                BuyerID
	DefaultDest       DestID
	Contracts         map[ContractID]ContractState
}

type Contract struct {
	IsSeller               bool
	ID                     ContractID
	State                  ContractState
	Buyer                  BuyerID
	Price                  int
	Limit                  int
	Speed                  int
	Length                 int
	StartingBlockTimestamp int
	Dest                   DestID
}
