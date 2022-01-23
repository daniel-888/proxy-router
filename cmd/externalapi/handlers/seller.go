package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

func SellersGET(seller *msgdata.SellerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := seller.GetAllSellers()
		c.JSON(http.StatusOK, results)
	}
}

func SellerGET(seller *msgdata.SellerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := seller.GetSeller(id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func SellerPOST(seller *msgdata.SellerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody []msgdata.SellerJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		for i := range(requestBody) {
			seller.AddSeller(requestBody[i])
			sellerMsg := msgdata.ConvertSellerJSONtoSellerMSG(requestBody[i])
			_,err := seller.Ps.PubWait(msgbus.SellerMsg, msgbus.IDString(sellerMsg.ID), sellerMsg)
			if err != nil {
				log.Printf("Seller POST request failed to update msgbus: %s", err)
			}
		}
		
		c.Status(http.StatusOK)
	}
}

func SellerPUT(seller *msgdata.SellerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var requestBody msgdata.SellerJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		if err := seller.UpdateSeller(id, requestBody); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		sellerMsg := msgdata.ConvertSellerJSONtoSellerMSG(requestBody)
		_,err := seller.Ps.SetWait(msgbus.SellerMsg, msgbus.IDString(sellerMsg.ID), sellerMsg)
		if err != nil {
			log.Printf("Seller PUT request failed to update msgbus: %s", err)
		}

		c.Status(http.StatusOK)
	}
}

func SellerDELETE(seller *msgdata.SellerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := seller.DeleteSeller(id); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		_,err := seller.Ps.UnpubWait(msgbus.SellerMsg, msgbus.IDString(id))
		if err != nil {
			log.Printf("Seller DELETE request failed to update msgbus: %s", err)
		}

		c.Status(http.StatusOK)
	}
}