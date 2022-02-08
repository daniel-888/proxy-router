package externalapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/handlers"
	"gitlab.com/TitanInd/lumerin/cmd/externalapi/msgdata"
	"gitlab.com/TitanInd/lumerin/cmd/msgbus"
)

type APIRepos struct {
	Config					*msgdata.ConfigInfoRepo 
	ContractManagerConfig	*msgdata.ContractManagerConfigRepo
	Connection 				*msgdata.ConnectionRepo 
	Contract				*msgdata.ContractRepo 
	Dest					*msgdata.DestRepo 
	Miner					*msgdata.MinerRepo 
	NodeOperator			*msgdata.NodeOperatorRepo
}

func (api *APIRepos) InitializeJSONRepos(ps *msgbus.PubSub) {
	api.Config = msgdata.NewConfigInfo(ps)
	go api.Config.SubscribeToConfigInfoMsgBus()

	api.ContractManagerConfig = msgdata.NewContractManagerConfig(ps)
	go api.ContractManagerConfig.SubscribeToContractManagerConfigMsgBus()

	api.Connection = msgdata.NewConnection(ps)
	go api.Connection.SubscribeToConnectionMsgBus()

	api.Contract = msgdata.NewContract(ps)
	go api.Contract.SubscribeToContractMsgBus()

	api.Dest = msgdata.NewDest(ps)
	go api.Dest.SubscribeToDestMsgBus()

	api.Miner = msgdata.NewMiner(ps)
	go api.Miner.SubscribeToMinerMsgBus()

	api.NodeOperator = msgdata.NewNodeOperator(ps)
	go api.NodeOperator.SubscribeToNodeOperatorMsgBus()
}

func (api *APIRepos) RunAPI() {
	r := gin.Default()

	configRoutes := r.Group("/config")
	{
		configRoutes.GET("/", handlers.ConfigsGET(api.Config))
		configRoutes.GET("/:id", handlers.ConfigGET(api.Config))
		configRoutes.POST("/", handlers.ConfigPOST(api.Config))
		configRoutes.PUT("/:id", handlers.ConfigPUT(api.Config))
		configRoutes.DELETE("/:id", handlers.ConfigDELETE(api.Config))
	}

	contractManagerConfigRoutes := r.Group("/contractmanagerconfig")
	{
		contractManagerConfigRoutes.GET("/", handlers.ContractManagerConfigsGET(api.ContractManagerConfig))
		contractManagerConfigRoutes.GET("/:id", handlers.ContractManagerConfigGET(api.ContractManagerConfig))
		contractManagerConfigRoutes.POST("/", handlers.ContractManagerConfigPOST(api.ContractManagerConfig))
		contractManagerConfigRoutes.PUT("/:id", handlers.ContractManagerConfigPUT(api.ContractManagerConfig))
		contractManagerConfigRoutes.DELETE("/:id", handlers.ContractManagerConfigDELETE(api.ContractManagerConfig))
	}

	connectionRoutes := r.Group("/connection")
	{
		connectionRoutes.GET("/", handlers.ConnectionsGET(api.Connection))
		connectionRoutes.GET("/:id", handlers.ConnectionGET(api.Connection))
		connectionRoutes.POST("/", handlers.ConnectionPOST(api.Connection))
		connectionRoutes.PUT("/:id", handlers.ConnectionPUT(api.Connection))
		connectionRoutes.DELETE("/:id", handlers.ConnectionDELETE(api.Connection))
	}

	contractRoutes := r.Group("/contract")
	{
		contractRoutes.GET("/", handlers.ContractsGET(api.Contract))
		contractRoutes.GET("/:id", handlers.ContractGET(api.Contract))
		contractRoutes.POST("/", handlers.ContractPOST(api.Contract))
		contractRoutes.PUT("/:id", handlers.ContractPUT(api.Contract))
		contractRoutes.DELETE("/:id", handlers.ContractDELETE(api.Contract))
	}

	destRoutes := r.Group("/dest")
	{
		destRoutes.GET("/", handlers.DestsGET(api.Dest))
		destRoutes.GET("/:id", handlers.DestGET(api.Dest))
		destRoutes.POST("/", handlers.DestPOST(api.Dest))
		destRoutes.PUT("/:id", handlers.DestPUT(api.Dest))
		destRoutes.DELETE("/:id", handlers.DestDELETE(api.Dest))
	}

	minerRoutes := r.Group("/miner")
	{
		minerRoutes.GET("/", handlers.MinersGET(api.Miner))
		minerRoutes.GET("/:id", handlers.MinerGET(api.Miner))
		minerRoutes.POST("/", handlers.MinerPOST(api.Miner))
		minerRoutes.PUT("/:id", handlers.MinerPUT(api.Miner))
		minerRoutes.DELETE("/:id", handlers.MinerDELETE(api.Miner))
	}

	nodeOperatorRoutes := r.Group("/nodeoperator")
	{
		nodeOperatorRoutes.GET("/", handlers.NodeOperatorsGET(api.NodeOperator))
		nodeOperatorRoutes.GET("/:id", handlers.NodeOperatorGET(api.NodeOperator))
		nodeOperatorRoutes.POST("/", handlers.NodeOperatorPOST(api.NodeOperator))
		nodeOperatorRoutes.PUT("/:id", handlers.NodeOperatorPUT(api.NodeOperator))
		nodeOperatorRoutes.DELETE("/:id", handlers.NodeOperatorDELETE(api.NodeOperator))
	}

	if err := r.Run(); err != nil {
		panic(fmt.Sprintf("external api failed to run:%s", err))
	}
}
