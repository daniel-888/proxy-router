package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
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
		var requestBody msgdata.SellerJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		
		seller.AddSeller(requestBody)
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
		c.Status(http.StatusOK)
	}
}