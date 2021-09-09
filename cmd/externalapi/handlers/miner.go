package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
)


func MinersGET(miner *msgdata.MinerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := miner.GetAllMiners()
		c.JSON(http.StatusOK, results)
	}
}

func MinerGET(miner *msgdata.MinerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := miner.GetMiner(id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func MinerPOST(miner *msgdata.MinerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody []msgdata.MinerJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		for i := range(requestBody) {
			miner.AddMiner(requestBody[i])
		}
		
		c.Status(http.StatusOK)
	}
}

func MinerPUT(miner *msgdata.MinerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var requestBody msgdata.MinerJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		if err := miner.UpdateMiner(id, requestBody); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}

func MinerDELETE(miner *msgdata.MinerRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := miner.DeleteMiner(id); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}