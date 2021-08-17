package configurationmanager

import (
	"testing"
)

func TestBoilerPlateFunc(t *testing.T) {
	msg, err := BoilerPlateFunc()
	if msg != "Configuration Manager Package" && err != nil {
		t.Fatalf("Test Failed")
	}
}