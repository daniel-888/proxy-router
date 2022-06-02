package connections

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/beevik/guid"
	"gitlab.com/TitanInd/lumerin/interfaces"
)

type ConnectionCollection struct {
	interfaces.IConnectionController

	connections sync.Map
}

func (m *ConnectionCollection) getCollection() ([]interfaces.IConnectionModel, error) {

	collection, ok := m.connections.Load("collection")

	if !ok {
		err := errors.New("Failed to get connection collection")
		log.Println(err)
		return nil, err
	}

	return collection.([]interfaces.IConnectionModel), nil
}

func (m *ConnectionCollection) getCollectionItem(query func(interfaces.IConnectionModel) bool) (interfaces.IConnectionModel, error) {

	collection, err := m.getCollection()

	if err != nil {
		return nil, err
	}

	for _, model := range collection {
		if query(model) {
			return model, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Failed to retrieve connection from internal collection - ConnectionCollection.getCollectionItem"))
}

func (m *ConnectionCollection) CreateNewConnection(connectionAddress string, destinationAddress string, status string) interfaces.IConnectionModel {
	return &Connection{
		Id:          guid.NewString(),
		Address:     connectionAddress,
		Destination: destinationAddress,
		isAvailable: false,
	}
}

func (m *ConnectionCollection) AddConnection(connectionAddress string, destinationAddress string, status string, id string) (interfaces.IConnectionModel, error) {
	model := m.CreateNewConnection(connectionAddress, destinationAddress, status)

	collection, err := m.getCollection()

	if err != nil {
		return nil, err
	}

	newCollection := append(collection, model)
	m.connections.Store(model.GetId(), model)
	m.connections.Store("collection", newCollection)

	return model, nil
}

func (m *ConnectionCollection) RemoveConnection(id string) error {
	collection, err := m.getCollection()

	if err != nil {
		return err
	}

	// newCollection := make([]interfaces.IConnectionModel, len(collection))

	for index, model := range collection {
		if model.GetId() == id {
			collection = append(collection[:index], collection[index+1:]...)
		}
	}

	m.connections.Delete(id)
	m.connections.Store("collection", collection)

	return nil
}

func (m *ConnectionCollection) GetOrAddConnection(connectionAddress string, destinationAddress string, status string, id string) (interfaces.IConnectionModel, error) {

	connection, ok := m.connections.Load(id)

	if !ok {
		var err error
		connection, err = m.AddConnection(connectionAddress, destinationAddress, status, id)

		if err != nil {
			log.Printf("Failed to get or add connection %v", id)

			return connection.(interfaces.IConnectionModel), err
		}
	}

	return connection.(interfaces.IConnectionModel), nil
}

func (m *ConnectionCollection) GetConnection(id string) interfaces.IConnectionModel {
	connection, ok := m.connections.Load(id)

	if !ok {
		log.Printf("Failed to get connection %v", id)
	}

	return connection.(interfaces.IConnectionModel)
}

func (m *ConnectionCollection) GetConnectionByAddresses(minerAddress string, destinationAddress string) (interfaces.IConnectionModel, error) {
	return m.getCollectionItem(func(model interfaces.IConnectionModel) bool {
		return model.AddressesMatch(minerAddress, destinationAddress)
	})
}

func (m *ConnectionCollection) GetConnections() ([]interfaces.IConnectionModel, error) {
	return m.getCollection()
}

func CreateConnectionCollection() interfaces.IConnectionController {
	connectionCollection := &ConnectionCollection{}

	connectionCollection.connections.Store("collection", []interfaces.IConnectionModel{})

	return connectionCollection
}
