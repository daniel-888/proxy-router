package connectionscheduler

import (
	"testing"
)

func TestBoilerPlateFunc(t *testing.T) {
	msg, err := BoilerPlateFunc()
	if msg != "Connection Scheduler Package" && err != nil {
		t.Fatalf("Test Failed")
	}
}