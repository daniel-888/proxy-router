package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
)


func DestsGET(dest *msgdata.DestRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := dest.GetAllDests()
		c.JSON(http.StatusOK, results)
	}
}

func DestGET(dest *msgdata.DestRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := dest.GetDest(id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func DestPOST(dest *msgdata.DestRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody msgdata.DestJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		
		dest.AddDest(requestBody)
		c.Status(http.StatusOK)
	}
}

func DestPUT(dest *msgdata.DestRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var requestBody msgdata.DestJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		if err := dest.UpdateDest(id, requestBody); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}

func DestDELETE(dest *msgdata.DestRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := dest.DeleteDest(id); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}