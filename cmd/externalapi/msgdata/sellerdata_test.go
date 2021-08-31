package msgdata

import (
	"fmt"
	"testing"
)

func TestAddSeller(t *testing.T) {
	seller := SellerJSON{
		ID:                     "Test",
		DefaultDest:            "Test",
		TotalAvailableHashRate: "Test",
		UnusedHashRate:         "Test",
	}
	
	sellerRepo := NewSeller()
	sellerRepo.AddSeller(seller)

	if len(sellerRepo.SellerJSONs) != 1 {
		t.Errorf("Seller struct not added")
	} 
}

func TestGetAllSellers(t *testing.T) {
	var seller [10]SellerJSON
	for i := 0; i < 10; i++ {
		seller[i].ID = "Test" + fmt.Sprint(i)
		seller[i].DefaultDest = "Test"
		seller[i].TotalAvailableHashRate = "Test"
		seller[i].UnusedHashRate = "Test"
	}
	
	sellerRepo := NewSeller()
	for i := 0; i < 10; i++ {
		sellerRepo.AddSeller(seller[i])
	}
	results := sellerRepo.GetAllSellers()

	if len(results) != 10 {
		t.Errorf("Could not get all seller structs")
	} 
} 

func TestGetSeller(t *testing.T) {
	var seller [10]SellerJSON
	for i := 0; i < 10; i++ {
		seller[i].ID = "Test" + fmt.Sprint(i)
		seller[i].DefaultDest = "Test"
		seller[i].TotalAvailableHashRate = "Test"
		seller[i].UnusedHashRate = "Test"
	}
	
	sellerRepo := NewSeller()
	for i := 0; i < 10; i++ {
		sellerRepo.AddSeller(seller[i])
	}

	var results [10]SellerJSON
	var errors [10]error
	for i := 0; i < 10; i++ {
		results[i], errors[i] = sellerRepo.GetSeller("Test" + fmt.Sprint(i))
		if errors[i] != nil {
			t.Errorf("GetSeller function returned error for this ID: " + results[i].ID)
		}
	}
}

func TestUpdateSeller(t *testing.T) {
	var seller [10]SellerJSON
	for i := 0; i < 10; i++ {
		seller[i].ID = "Test" + fmt.Sprint(i)
		seller[i].DefaultDest = "Test"
		seller[i].TotalAvailableHashRate = "Test"
		seller[i].UnusedHashRate = "Test"
	}
	
	sellerRepo := NewSeller()
	for i := 0; i < 10; i++ {
		sellerRepo.AddSeller(seller[i])
	}

	sellerUpdates := SellerJSON{
		ID:                     "",
		DefaultDest:            "",
		TotalAvailableHashRate: "Updated",
		UnusedHashRate:         "",
	}
	
	var results [10]SellerJSON
	var errors [10]error
	for i := 0; i < 10; i++ {
		errors[i] = sellerRepo.UpdateSeller("Test" + fmt.Sprint(i), sellerUpdates)
		results[i],_ = sellerRepo.GetSeller("Test" + fmt.Sprint(i))
		if errors[i] != nil {
			t.Errorf("UpdateSeller function returned error for this ID: " + results[i].ID)
		}
		if results[i].TotalAvailableHashRate != "Updated" {
			t.Errorf("UpdateSeller function did not update Seller Struct")
		}
		if results[i].ID != seller[i].ID {
			t.Errorf("UpdateSeller function updated all Seller fields instead of just filled in field")
		}
	}
}

func TestDeleteSeller(t *testing.T) {
	var seller [10]SellerJSON
	for i := 0; i < 10; i++ {
		seller[i].ID = "Test" + fmt.Sprint(i)
		seller[i].DefaultDest = "Test"
		seller[i].TotalAvailableHashRate = "Test"
		seller[i].UnusedHashRate = "Test"
	}
	
	sellerRepo := NewSeller()
	for i := 0; i < 10; i++ {
		sellerRepo.AddSeller(seller[i])
	}
	
	error := sellerRepo.DeleteSeller("Test7")
	if error != nil {
		t.Errorf("DeleteSeller function returned error")
	}
	if len(sellerRepo.SellerJSONs) != 9 {
		t.Errorf("Seller was not deleted")
	}
}