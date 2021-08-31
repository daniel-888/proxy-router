package msgdata

import (
	"fmt"
	"testing"
)

func TestAddMiner(t *testing.T) {
	miner := MinerJSON{
		ID:						"Test",
		State: 					"Test",
		Seller:   				"Test",
		Dest:					"Test",	
		InitialMeasuredHashRate: "Test",
		CurrentHashRate:         "Test",
	}
	
	minerRepo := NewMiner()
	minerRepo.AddMiner(miner)

	if len(minerRepo.MinerJSONs) != 1 {
		t.Errorf("Miner struct not added")
	} 
}

func TestGetAllMiners(t *testing.T) {
	var miner [10]MinerJSON
	for i := 0; i < 10; i++ {
		miner[i].ID = "Test" + fmt.Sprint(i)
		miner[i].State = "Test"
		miner[i].Seller = "Test"
		miner[i].Dest = "Test"
		miner[i].InitialMeasuredHashRate = "Test"
		miner[i].CurrentHashRate = "Test"
	}
	
	minerRepo := NewMiner()
	for i := 0; i < 10; i++ {
		minerRepo.AddMiner(miner[i])
	}
	results := minerRepo.GetAllMiners()

	if len(results) != 10 {
		t.Errorf("Could not get all miner structs")
	} 
} 

func TestGetMiner(t *testing.T) {
	var miner [10]MinerJSON
	for i := 0; i < 10; i++ {
		miner[i].ID = "Test" + fmt.Sprint(i)
		miner[i].State = "Test"
		miner[i].Seller = "Test"
		miner[i].Dest = "Test"
		miner[i].InitialMeasuredHashRate = "Test"
		miner[i].CurrentHashRate = "Test"
	}
	
	minerRepo := NewMiner()
	for i := 0; i < 10; i++ {
		minerRepo.AddMiner(miner[i])
	}

	var results [10]MinerJSON
	var errors [10]error
	for i := 0; i < 10; i++ {
		results[i], errors[i] = minerRepo.GetMiner("Test" + fmt.Sprint(i))
		if errors[i] != nil {
			t.Errorf("GetMiner function returned error for this ID: " + results[i].ID)
		}
	}
}

func TestUpdateMiner(t *testing.T) {
	var miner [10]MinerJSON
	for i := 0; i < 10; i++ {
		miner[i].ID = "Test" + fmt.Sprint(i)
		miner[i].State = "Test"
		miner[i].Seller = "Test"
		miner[i].Dest = "Test"
		miner[i].InitialMeasuredHashRate = "Test"
		miner[i].CurrentHashRate = "Test"
	}
	
	minerRepo := NewMiner()
	for i := 0; i < 10; i++ {
		minerRepo.AddMiner(miner[i])
	}

	minerUpdates := MinerJSON{
		ID:						"",
		State: 					"Updated",
		Seller:   				"",
		Dest:					"",	
		InitialMeasuredHashRate: "",
		CurrentHashRate:         "",
	}
	
	var results [10]MinerJSON
	var errors [10]error
	for i := 0; i < 10; i++ {
		errors[i] = minerRepo.UpdateMiner("Test" + fmt.Sprint(i), minerUpdates)
		results[i],_ = minerRepo.GetMiner("Test" + fmt.Sprint(i))
		if errors[i] != nil {
			t.Errorf("UpdateMiner function returned error for this ID: " + results[i].ID)
		}
		if results[i].State != "Updated" {
			t.Errorf("UpdateMiner function did not update Miner Struct")
		}
		if results[i].ID != miner[i].ID {
			t.Errorf("UpdateMiner function updated all Miner fields instead of just filled in field")
		}
	}
}

func TestDeleteMiner(t *testing.T) {
	var miner [10]MinerJSON
	for i := 0; i < 10; i++ {
		miner[i].ID = "Test" + fmt.Sprint(i)
		miner[i].State = "Test"
		miner[i].Seller = "Test"
		miner[i].Dest = "Test"
		miner[i].InitialMeasuredHashRate = "Test"
		miner[i].CurrentHashRate = "Test"
	}
	
	minerRepo := NewMiner()
	for i := 0; i < 10; i++ {
		minerRepo.AddMiner(miner[i])
	}
	
	error := minerRepo.DeleteMiner("Test7")
	if error != nil {
		t.Errorf("DeleteMiner function returned error")
	}
	if len(minerRepo.MinerJSONs) != 9 {
		t.Errorf("Miner was not deleted")
	}
}