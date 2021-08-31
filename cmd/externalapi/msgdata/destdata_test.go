package msgdata

import (
	"fmt"
	"testing"
)

func TestAddDest(t *testing.T) {
	dest := DestJSON{
		ID:   "DestID01",
		IP:   "127.0.0.1",
		Port: "80",
	}
	
	destRepo := NewDest()
	destRepo.AddDest(dest)

	if len(destRepo.DestJSONs) != 1 {
		t.Errorf("Dest struct not added")
	} 
}

func TestGetAllDests(t *testing.T) {
	var dest [10]DestJSON
	for i := 0; i < 10; i++ {
		dest[i].ID = "Test" + fmt.Sprint(i)
		dest[i].IP = "Test"
		dest[i].Port = "Test"
	}
	
	destRepo := NewDest()
	for i := 0; i < 10; i++ {
		destRepo.AddDest(dest[i])
	}
	results := destRepo.GetAllDests()

	if len(results) != 10 {
		t.Errorf("Could not get all dest structs")
	} 
} 

func TestGetDest(t *testing.T) {
	var dest [10]DestJSON
	for i := 0; i < 10; i++ {
		dest[i].ID = "Test" + fmt.Sprint(i)
		dest[i].IP = "Test"
		dest[i].Port = "Test"
	}
	
	destRepo := NewDest()
	for i := 0; i < 10; i++ {
		destRepo.AddDest(dest[i])
	}

	var results [10]DestJSON
	var errors [10]error
	for i := 0; i < 10; i++ {
		results[i], errors[i] = destRepo.GetDest("Test" + fmt.Sprint(i))
		if errors[i] != nil {
			t.Errorf("GetDest function returned error for this ID: " + results[i].ID)
		}
	}
}

func TestUpdateDest(t *testing.T) {
	var dest [10]DestJSON
	for i := 0; i < 10; i++ {
		dest[i].ID = "Test" + fmt.Sprint(i)
		dest[i].IP = "Test"
		dest[i].Port = "Test"
	}
	
	destRepo := NewDest()
	for i := 0; i < 10; i++ {
		destRepo.AddDest(dest[i])
	}

	destUpdates := DestJSON{
		ID:   "",
		IP:   "Updated",
		Port: "",
	}
	
	var results [10]DestJSON
	var errors [10]error
	for i := 0; i < 10; i++ {
		errors[i] = destRepo.UpdateDest("Test" + fmt.Sprint(i), destUpdates)
		results[i],_ = destRepo.GetDest("Test" + fmt.Sprint(i))
		if errors[i] != nil {
			t.Errorf("UpdateDest function returned error for this ID: " + results[i].ID)
		}
		if results[i].IP != "Updated" {
			t.Errorf("UpdateDest function did not update Dest Struct")
		}
		if results[i].ID != dest[i].ID {
			t.Errorf("UpdateDest function updated all Dest fields instead of just filled in field")
		}
	}
}

func TestDeleteDest(t *testing.T) {
	var dest [10]DestJSON
	for i := 0; i < 10; i++ {
		dest[i].ID = "Test" + fmt.Sprint(i)
		dest[i].IP = "Test"
		dest[i].Port = "Test"
	}
	
	destRepo := NewDest()
	for i := 0; i < 10; i++ {
		destRepo.AddDest(dest[i])
	}
	
	error := destRepo.DeleteDest("Test7")
	if error != nil {
		t.Errorf("DeleteDest function returned error")
	}
	if len(destRepo.DestJSONs) != 9 {
		t.Errorf("Dest was not deleted")
	}
}