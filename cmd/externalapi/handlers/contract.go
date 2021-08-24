package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
)


func ContractsGET(contract *msgdata.ContractRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := contract.GetAllContracts()
		c.JSON(http.StatusOK, results)
	}
}

func ContractGET(contract *msgdata.ContractRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result, err := contract.GetContract(id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

func ContractPOST(contract *msgdata.ContractRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody msgdata.ContractJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}
		
		contract.AddContract(requestBody)
		c.Status(http.StatusOK)
	}
}

func ContractPUT(contract *msgdata.ContractRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var requestBody msgdata.ContractJSON
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		if err := contract.UpdateContract(id, requestBody); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}

func ContractDELETE(contract *msgdata.ContractRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := contract.DeleteContract(id); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusOK)
	}
}