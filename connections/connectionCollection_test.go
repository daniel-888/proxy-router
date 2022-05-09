package connections

import (
	"testing"
)

var testMinerAddress = "testconnectionaddress"
var testPoolAddress = "testpooladdress"

func TestCreateConnection(t *testing.T) {
	col := CreateConnectionCollection()

	connection := col.CreateNewConnection(testMinerAddress, testPoolAddress, "status").(*Connection)

	if connection.address != testMinerAddress {
		t.Errorf("Expected connection.address to equal %v; equals %v", testMinerAddress, connection.address)
	}

	if connection.destination != testPoolAddress {
		t.Errorf("Expected connection.destination to equal %v; equals %v", testPoolAddress, connection.destination)
	}
}

func TestAddAndGetConnection(t *testing.T) {
	col := CreateConnectionCollection()

	err := col.AddConnection(testMinerAddress, testPoolAddress, "status")

	if err != nil {
		t.Error(err)
	}

	gotConnection := col.GetConnection(testMinerAddress).(*Connection)

	if gotConnection.address != testMinerAddress {
		t.Errorf("Expected connection.address to equal %v; equals %v", testMinerAddress, gotConnection.address)
	}

	if gotConnection.destination != testPoolAddress {
		t.Errorf("Expected connection.destination to equal %v; equals %v", testPoolAddress, gotConnection.destination)
	}
}

func TestAddAndGetConnections(t *testing.T) {
	col := CreateConnectionCollection()

	testConnectionAddress2 := testMinerAddress + "2"
	testConnectionAddress3 := testMinerAddress + "3"

	col.AddConnection(testMinerAddress, testPoolAddress, "status")
	col.AddConnection(testConnectionAddress2, testPoolAddress, "status2")
	col.AddConnection(testConnectionAddress3, testPoolAddress, "status3")

	connections, err := col.GetConnections()

	if err != nil {
		t.Errorf("failed to get connections with GetConnection function; %v", err)
	}

	gotConnection := connections[0].(*Connection)
	gotConnection2 := connections[1].(*Connection)
	gotConnection3 := connections[2].(*Connection)

	if gotConnection.address != testMinerAddress {
		t.Errorf("Expected gotConnection.address to equal '%v'; equals '%v'", testMinerAddress, gotConnection.address)
	}

	if gotConnection.destination != testPoolAddress {
		t.Errorf("Expected gotConnection.destination to equal '%v'; equals '%v'", testPoolAddress, gotConnection.destination)
	}

	if gotConnection.status != "status" {
		t.Errorf("Expected gotConnection.status to equal 'status'; equals '%v'", gotConnection.status)
	}

	if gotConnection2.address != testConnectionAddress2 {
		t.Errorf("Expected gotConnection2.address to equal '%v'; equals '%v'", testConnectionAddress2, gotConnection2.address)
	}

	if gotConnection2.destination != testPoolAddress {
		t.Errorf("Expected gotConnection.destination to equal '%v'; equals '%v'", testPoolAddress, gotConnection2.destination)
	}

	if gotConnection2.status != "status2" {
		t.Errorf("Expected gotConnection.status to equal 'status2'; equals '%v'", gotConnection2.status)
	}

	if gotConnection3.address != testConnectionAddress3 {
		t.Errorf("Expected gotConnection3.address to equal '%v'; equals '%v'", testConnectionAddress3, gotConnection3.address)
	}

	if gotConnection3.destination != testPoolAddress {
		t.Errorf("Expected gotConnection3.destination to equal '%v'; equals '%v'", testPoolAddress, gotConnection3.destination)
	}

	if gotConnection3.status != "status3" {
		t.Errorf("Expected gotConnection.status to equal 'status3'; equals '%v'", gotConnection3.status)
	}
}
