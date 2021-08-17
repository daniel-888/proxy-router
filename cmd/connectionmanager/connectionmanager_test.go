package connectionmanager

import (
	"testing"
)

func TestBoilerPlateFunc(t *testing.T) {
	msg, err := BoilerPlateFunc()
	if msg != "Connection Manager Package" && err != nil {
		t.Fatalf("Test Failed")
	}
}