package contractmanager

import (
	"testing"
)

func TestBoilerPlateFunc(t *testing.T) {
	msg, err := BoilerPlateFunc()
	if msg != "Contract Manager Package" && err != nil {
		t.Fatalf("Test Failed")
	}
}