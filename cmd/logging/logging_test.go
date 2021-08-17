package logging

import (
	"testing"
)

func TestBoilerPlateFunc(t *testing.T) {
	msg, err := BoilerPlateFunc()
	if msg != "Logging Package" && err != nil {
		t.Fatalf("Test Failed")
	}
}