package stratumv1

import (
	"fmt"
	"math/rand"
	"testing"
)

//
// Provide standard JSON testing
//
//

//
//
//
func TestUnmarshalSubscribe(t *testing.T) {

	id := rand.Intn(1000000) + 1

	subscribeString := fmt.Sprintf("{\"id\":%d, \"method\":\"mining.subscribe\", \"params\":[ \"cpuminer/2.5.1\" ] }", id)

	subscribe := []byte(subscribeString)

	ret, err := unmarshalMsg(subscribe)
	if err != nil {
		t.Errorf("unmarshalMsg failed: %s", err)
	}

	switch ret.(type) {
	case *stratumRequest:
	default:
		t.Errorf("unmarshalMsg wrong type returned: %T", ret)
	}

	idResult, err := ret.(*stratumRequest).getID()
	if err != nil {
		t.Errorf("unmarshalMsg getID() failed: %s", err)
	}
	if id != idResult {
		t.Errorf("unmarshalMsg ID result mis-match: id %d, result:%d", id, idResult)

	}
}
