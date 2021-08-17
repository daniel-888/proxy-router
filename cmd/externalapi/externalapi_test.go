package externalapi

import (
	"testing"
)

func TestBoilerPlateFunc(t *testing.T) {
	msg, err := BoilerPlateFunc()
	if msg != "External API Package" && err != nil {
		t.Fatalf("Test Failed")
	}
}