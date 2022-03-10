package simple

import (
	"context"
	_ "fmt"
	"net"
	_ "reflect"
	"testing"
)

/*
testing system for the SIMPLE layer
The SIMPLE layer will need to be tested to
ensure that is can route packets depending on outside instructions
As additional functionality is added to the SIMPLE layer
(multithreading, compressing, rerouting, error handling, etc)
additional feature and functional tests will need to be written.

The Stratum layer, lower level layers, and MSG are not being tested in this
testing suite, however their messages may either be used or simulated for testing purposes
*/

/*
this is just a cookie cutter placehodler so new tests
can be quickly implemented
func TestTemplate(t *testing.T) {
}
*/
func generateTestContext() context.Context {
	returnContext := context.TODO()
	return returnContext
}

func generateTestAddr() net.Addr {
	conn, _ := net.Dial("tcp", "golang.org:http")
	return conn.RemoteAddr()
}

//generate a SimpleStruct for testing purposes
func generateSimpleStruct() SimpleStruct {
	myContext := generateTestContext()
	mySimpleStruct := SimpleStruct{
		ctx: myContext,
		cancel: dummyFunc,
		eventHandler: 0,
		eventChan: make(chan SimpleEvent),
		protocolChan: make(chan SimpleEvent),
		commChan: make(chan []byte),
	}
	return mySimpleStruct
}

//generate a SimpleListenStruct for testing purposes
func generateSimpleListenStruct() SimpleListenStruct {
	myContext := generateTestContext()
	myAddr := generateTestAddr()
	myStruct, _ := New(myContext, myAddr)
	return myStruct
}

/*
test steps
1. protocol layer calls the New function in the SIMPLE package
	a context and an Addr are passed into New
2. checks the error message, test fails if the error message is anything but nil

*/
func TestInitializeSimpleListenStruct(t *testing.T) {
	_, err := New(generateTestContext(), generateTestAddr())
	if err != nil {
		t.Error("failed to initialize SimpleListenStruct")
	}

}

/*
test steps
1. create a dummy communication layer to listen for information from a context
2. create a new simple struct routine
3. call the run function and pass in a standard context
4. check the dummy communication layer to see if the context inforation has been passed fown
*/
func TestSimpleStructRun(t *testing.T) {
	simpleStruct := generateSimpleStruct()
	context := generateTestContext() //this needs to be replaced to accept a context from the protocol
	simpleStruct.Run(context) //run is working but needs to do "something" with the context
}

/*
test steps
1. create a new SimpleListenStruct
2. call the close function on the SimpleListenStruct
3. ensure all associated routines are closed
*/
func TestSimpleListenStructClose(t *testing.T) {
	listenStruct, _ := New(generateTestContext(), generateTestAddr())
	listenStruct.Close()
}

/*
test steps
1. create a new SimpleListenStruct
2. call the close function on the SimpleListenStruct
3. ensure all associated routines are closed
*/
func TestSimpleStructClose(t *testing.T) {
	simpleStruct := generateSimpleStruct()
	context := generateTestContext() //this needs to be replaced to accept a context from the protocol
	simpleStruct.Run(context) //run is working but needs to do "something" with the context
	simpleStruct.Close()
}

func TestSetMessageSizeDefault(t *testing.T) {
	simpleStruct := generateSimpleStruct()
	if simpleStruct.maxMessageSize != 0 {
		t.Errorf("message expected to be 0, actually is: %d", simpleStruct.maxMessageSize)
	}
	simpleStruct.SetMessageSizeDefault(100)
	if simpleStruct.maxMessageSize != 100 {
		t.Errorf("message expected to be 100, actually is: %d", simpleStruct.maxMessageSize)
	}

}

/*
testing that a SimpleStruct will dial a connection and accuratley store the resulting
connection in the mapping, and retrieve the mapping
*/
func TestDialFunctionality(t *testing.T) {
	simpleStruct := generateSimpleStruct()
	testAddr := generateTestAddr()

	if simpleStruct.connectionIndex !=0 {
		t.Error("testing index is not 0")
	}

	uID, e := simpleStruct.Dial(testAddr)
	if uID != 0 {
		t.Error("conn index is not 0")
	}

	if e != nil {
		t.Errorf("%s",e)
	}

}





