package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
)

func ConfigsGET(config *msgdata.ConfigInfoRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := config.GetAllConfigInfos()
		c.JSON(http.StatusOK, results)
	}
}

func ConfigGET(config *msgdata.ConfigInfoRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := config.GetConfigInfo(id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func ConfigPOST(config *msgdata.ConfigInfoRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody msgdata.ConfigInfoJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		
		config.AddConfigInfo(requestBody)
		c.Status(http.StatusOK)
	}
}

func ConfigPUT(config *msgdata.ConfigInfoRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var requestBody msgdata.ConfigInfoJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		if err := config.UpdateConfigInfo(id, requestBody); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}

func ConfigDELETE(config *msgdata.ConfigInfoRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := config.DeleteConfigInfo(id); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}