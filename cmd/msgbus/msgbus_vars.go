package msgbus

type ContractState string
type MinerState string

const (
	ContAvailableState ContractState = "AvailableState"
	ContActiveState    ContractState = "ActiveState"
	ContRunningState   ContractState = "RunningState"
	ContCompleteState  ContractState = "CompleteState"
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
	AvailableContracts     map[ContractID]bool
	ActiveContracts        map[ContractID]bool
	RunningContracts       map[ContractID]bool
	CompleteContracts      map[ContractID]bool
}

type Buyer struct {
	ID                BuyerID
	DefaultDest       DestID
	ActiveContracts   map[ContractID]bool
	RunningContracts  map[ContractID]bool
	CompleteContracts map[ContractID]bool
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
	ValidationFee          int
	StartingBlockTimestamp int
	Dest                   DestID
}
