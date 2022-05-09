package interfaces

type IConnectionController interface {
	CreateNewConnection(minerAddress string, destinationAddress string, status string) IConnectionModel
	AddConnection(minerAddress string, destinationAddress string, status string, id string) (IConnectionModel, error)
	RemoveConnection(id string) error
	GetConnection(id string) IConnectionModel
	GetOrAddConnection(connectionAddress string, destinationAddress string, status string, id string) (IConnectionModel, error)
	GetConnections() ([]IConnectionModel, error)
}
