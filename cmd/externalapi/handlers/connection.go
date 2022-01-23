package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
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
			connMsg := msgdata.ConvertConnectionJSONtoConnectionMSG(requestBody[i])
			_,err := conn.Ps.PubWait(msgbus.ConnectionMsg, msgbus.IDString(connMsg.ID), connMsg)
			if err != nil {
				log.Printf("Connection POST request failed to update msgbus: %s", err)
			}
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
		
		connMsg := msgdata.ConvertConnectionJSONtoConnectionMSG(requestBody)
		_,err := conn.Ps.SetWait(msgbus.ConnectionMsg, msgbus.IDString(connMsg.ID), connMsg)
		if err != nil {
			log.Printf("Connection PUT request failed to update msgbus: %s", err)
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

		_,err := conn.Ps.UnpubWait(msgbus.ConnectionMsg, msgbus.IDString(id))
		if err != nil {
			log.Printf("Connection DELETE request failed to update msgbus: %s", err)
		}

		c.Status(http.StatusOK)
	}
}