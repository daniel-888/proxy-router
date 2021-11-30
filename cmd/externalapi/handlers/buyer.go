package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
)


func BuyersGET(buyer *msgdata.BuyerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := buyer.GetAllBuyers()
		c.JSON(http.StatusOK, results)
	}
}

func BuyerGET(buyer *msgdata.BuyerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := buyer.GetBuyer(id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func BuyerPOST(buyer *msgdata.BuyerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody []msgdata.BuyerJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		for i := range(requestBody) {
			buyer.AddBuyer(requestBody[i])
		}
		
		c.Status(http.StatusOK)
	}
}

func BuyerPUT(buyer *msgdata.BuyerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var requestBody msgdata.BuyerJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		if err := buyer.UpdateBuyer(id, requestBody); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}

func BuyerDELETE(buyer *msgdata.BuyerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := buyer.DeleteBuyer(id); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}