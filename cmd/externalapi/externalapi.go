package externalapi

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/handlers"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
)

func InitializeJSONRepos() (*msgdata.ConfigInfoRepo, 
							*msgdata.ConnectionRepo, 
							*msgdata.ContractRepo, 
							*msgdata.DestRepo, 
							*msgdata.MinerRepo, 
							*msgdata.SellerRepo) {

	config := msgdata.NewConfigInfo()
	connection := msgdata.NewConnection()
	contract := msgdata.NewContract()
	dest := msgdata.NewDest()
	miner := msgdata.NewMiner()
	seller := msgdata.NewSeller()

	return config, connection, contract, dest, miner, seller
}

func RunAPI(config *msgdata.ConfigInfoRepo,
			connection *msgdata.ConnectionRepo,
			contract *msgdata.ContractRepo,
			dest *msgdata.DestRepo,
			miner *msgdata.MinerRepo,
			seller *msgdata.SellerRepo) {
	
	r := gin.Default()

	configRoutes := r.Group("/config")
	{
		configRoutes.GET("/", handlers.ConfigsGET(config))
		configRoutes.GET("/:id", handlers.ConfigGET(config))
		configRoutes.POST("/", handlers.ConfigPOST(config))
		configRoutes.PUT("/:id", handlers.ConfigPUT(config))
		configRoutes.DELETE("/:id", handlers.ConfigDELETE(config))
	}

	connectionRoutes := r.Group("/connection")
	{
		connectionRoutes.GET("/", handlers.ConnectionsGET(connection))
		connectionRoutes.GET("/:id", handlers.ConnectionGET(connection))
		connectionRoutes.POST("/", handlers.ConnectionPOST(connection))
		connectionRoutes.PUT("/:id", handlers.ConnectionPUT(connection))
		connectionRoutes.DELETE("/:id", handlers.ConnectionDELETE(connection))
	}

	contractRoutes := r.Group("/contract")
	{
		contractRoutes.GET("/", handlers.ContractsGET(contract))
		contractRoutes.GET("/:id", handlers.ContractGET(contract))
		contractRoutes.POST("/", handlers.ContractPOST(contract))
		contractRoutes.PUT("/:id", handlers.ContractPUT(contract))
		contractRoutes.DELETE("/:id", handlers.ContractDELETE(contract))
	}

	destRoutes := r.Group("/dest")
	{
		destRoutes.GET("/", handlers.DestsGET(dest))
		destRoutes.GET("/:id", handlers.DestGET(dest))
		destRoutes.POST("/", handlers.DestPOST(dest))
		destRoutes.PUT("/:id", handlers.DestPUT(dest))
		destRoutes.DELETE("/:id", handlers.DestDELETE(dest))
	}

	minerRoutes := r.Group("/miner")
	{
		minerRoutes.GET("/", handlers.MinersGET(miner))
		minerRoutes.GET("/:id", handlers.MinerGET(miner))
		minerRoutes.POST("/", handlers.MinerPOST(miner))
		minerRoutes.PUT("/:id", handlers.MinerPUT(miner))
		minerRoutes.DELETE("/:id", handlers.MinerDELETE(miner))
	}

	sellerRoutes := r.Group("/seller")
	{
		sellerRoutes.GET("/", handlers.SellersGET(seller))
		sellerRoutes.GET("/:id", handlers.SellerGET(seller))
		sellerRoutes.POST("/", handlers.SellerPOST(seller))
		sellerRoutes.PUT("/:id", handlers.SellerPUT(seller))
		sellerRoutes.DELETE("/:id", handlers.SellerDELETE(seller))
	}

	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
	
}
