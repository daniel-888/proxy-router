package msgdata

import (
	"fmt"
	"testing"

	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func TestAddSeller(t *testing.T) {
	seller := SellerJSON{
		ID:                     "Test",
		DefaultDest:            "Test",
		TotalAvailableHashRate: 100,
		UnusedHashRate:         100,
	}
	seller.NewContracts = map[msgbus.ContractID]bool{
		"0x85A256C5688D012263D5A79EE37E84FC35EC4524": true,
        "0x89921E8D51D22252D64EA34340A4161696887271": false,
        "0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3": true,
	}
	seller.ReadyContracts = map[msgbus.ContractID]bool{
		"0x50937C047DB93CB5C87F65B6EFFEA47D03DF0F7D": true,
        "0xFB610E4C269DA110C97B92F5F34EAA50E5F3D500": false,
        "0x397729E80F77BA09D930FE24E8D1FC74372E86D3": true,
	}   
    seller.ActiveContracts = map[msgbus.ContractID]bool{
		"0x9F252E1EC723AF6D96A36B4EB2B75A262291497C": true,
        "0xBB2EAAAAA9B08EC320FC984D7D19E28835DD94DD": false,
        "0x407E8A225658FEE384859874952E2BBC11E98B5C": true,
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
		seller[i].TotalAvailableHashRate = 100
		seller[i].UnusedHashRate = 100
		seller[i].NewContracts = map[msgbus.ContractID]bool{
			"0x85A256C5688D012263D5A79EE37E84FC35EC4524": true,
			"0x89921E8D51D22252D64EA34340A4161696887271": false,
			"0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3": true,
		}
		seller[i].ReadyContracts = map[msgbus.ContractID]bool{
			"0x50937C047DB93CB5C87F65B6EFFEA47D03DF0F7D": true,
			"0xFB610E4C269DA110C97B92F5F34EAA50E5F3D500": false,
			"0x397729E80F77BA09D930FE24E8D1FC74372E86D3": true,
		}   
		seller[i].ActiveContracts = map[msgbus.ContractID]bool{
			"0x9F252E1EC723AF6D96A36B4EB2B75A262291497C": true,
			"0xBB2EAAAAA9B08EC320FC984D7D19E28835DD94DD": false,
			"0x407E8A225658FEE384859874952E2BBC11E98B5C": true,
		}
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
		seller[i].TotalAvailableHashRate = 100
		seller[i].UnusedHashRate = 100
		seller[i].NewContracts = map[msgbus.ContractID]bool{
			"0x85A256C5688D012263D5A79EE37E84FC35EC4524": true,
			"0x89921E8D51D22252D64EA34340A4161696887271": false,
			"0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3": true,
		}
		seller[i].ReadyContracts = map[msgbus.ContractID]bool{
			"0x50937C047DB93CB5C87F65B6EFFEA47D03DF0F7D": true,
			"0xFB610E4C269DA110C97B92F5F34EAA50E5F3D500": false,
			"0x397729E80F77BA09D930FE24E8D1FC74372E86D3": true,
		}   
		seller[i].ActiveContracts = map[msgbus.ContractID]bool{
			"0x9F252E1EC723AF6D96A36B4EB2B75A262291497C": true,
			"0xBB2EAAAAA9B08EC320FC984D7D19E28835DD94DD": false,
			"0x407E8A225658FEE384859874952E2BBC11E98B5C": true,
		}
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
		seller[i].TotalAvailableHashRate = 100
		seller[i].UnusedHashRate = 100
		seller[i].NewContracts = map[msgbus.ContractID]bool{
			"0x85A256C5688D012263D5A79EE37E84FC35EC4524": true,
			"0x89921E8D51D22252D64EA34340A4161696887271": false,
			"0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3": true,
		}
		seller[i].ReadyContracts = map[msgbus.ContractID]bool{
			"0x50937C047DB93CB5C87F65B6EFFEA47D03DF0F7D": true,
			"0xFB610E4C269DA110C97B92F5F34EAA50E5F3D500": false,
			"0x397729E80F77BA09D930FE24E8D1FC74372E86D3": true,
		}   
		seller[i].ActiveContracts = map[msgbus.ContractID]bool{
			"0x9F252E1EC723AF6D96A36B4EB2B75A262291497C": true,
			"0xBB2EAAAAA9B08EC320FC984D7D19E28835DD94DD": false,
			"0x407E8A225658FEE384859874952E2BBC11E98B5C": true,
		}
	}
	
	sellerRepo := NewSeller()
	for i := 0; i < 10; i++ {
		sellerRepo.AddSeller(seller[i])
	}

	sellerUpdates := SellerJSON{
		ID:                     "",
		DefaultDest:            "",
		TotalAvailableHashRate: 10001,
		UnusedHashRate:         0,
	}
	sellerUpdates.NewContracts = map[msgbus.ContractID]bool{}
	sellerUpdates.ReadyContracts = map[msgbus.ContractID]bool{}   
    sellerUpdates.ActiveContracts = map[msgbus.ContractID]bool{}
	
	var results [10]SellerJSON
	var errors [10]error
	for i := 0; i < 10; i++ {
		errors[i] = sellerRepo.UpdateSeller("Test" + fmt.Sprint(i), sellerUpdates)
		results[i],_ = sellerRepo.GetSeller("Test" + fmt.Sprint(i))
		if errors[i] != nil {
			t.Errorf("UpdateSeller function returned error for this ID: " + results[i].ID)
		}
		if results[i].TotalAvailableHashRate != 10001 {
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
		seller[i].TotalAvailableHashRate = 100
		seller[i].UnusedHashRate = 100
		seller[i].NewContracts = map[msgbus.ContractID]bool{
			"0x85A256C5688D012263D5A79EE37E84FC35EC4524": true,
			"0x89921E8D51D22252D64EA34340A4161696887271": false,
			"0xF68F06C4189F360D9D1AA7F3B5135E5F2765DAA3": true,
		}
		seller[i].ReadyContracts = map[msgbus.ContractID]bool{
			"0x50937C047DB93CB5C87F65B6EFFEA47D03DF0F7D": true,
			"0xFB610E4C269DA110C97B92F5F34EAA50E5F3D500": false,
			"0x397729E80F77BA09D930FE24E8D1FC74372E86D3": true,
		}   
		seller[i].ActiveContracts = map[msgbus.ContractID]bool{
			"0x9F252E1EC723AF6D96A36B4EB2B75A262291497C": true,
			"0xBB2EAAAAA9B08EC320FC984D7D19E28835DD94DD": false,
			"0x407E8A225658FEE384859874952E2BBC11E98B5C": true,
		}
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