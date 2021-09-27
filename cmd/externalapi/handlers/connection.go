package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
)


func ConnectionsGET(conn *msgdata.ConnectionRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := conn.GetAllConnections()
		c.JSON(http.StatusOK, results)
	}
}

func ConnectionGET(conn *msgdata.ConnectionRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := conn.GetConnection(id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func ConnectionPOST(conn *msgdata.ConnectionRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody []msgdata.ConnectionJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		for i := range(requestBody) {
			conn.AddConnection(requestBody[i])
		}

		c.Status(http.StatusOK)
	}
}

func ConnectionPUT(conn *msgdata.ConnectionRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var requestBody msgdata.ConnectionJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		if err := conn.UpdateConnection(id, requestBody); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}

func ConnectionDELETE(conn *msgdata.ConnectionRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := conn.DeleteConnection(id); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}