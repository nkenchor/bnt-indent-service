package main

import (
	mongoRepository "bnt/bnt-indent-service/internal/adapter/repository/mongodb"
	"bnt/bnt-indent-service/internal/adapter/routes"

	message "bnt/bnt-indent-service/internal/core/helper/message-helper"
	logger "bnt/bnt-indent-service/internal/core/helper/log-helper"
	configuration "bnt/bnt-indent-service/internal/core/helper/configuration-helper"
	"fmt"
	"log"
)

func main() {
	
	//Initialize request Log
	logger.InitializeLog()
	//Start DB Connection
	mongoRepo := startMongo()
	logger.LogEvent("INFO", "MongoDB Connected and Initialized!")

	logger.LogEvent("INFO", "Redis Connected and Initialized!")
	//Set up routes
	router := routes.SetupRouter(mongoRepo.Indent)
	//Print custom message for server start

	fmt.Println(message.ServerStarted)
	//Log server start event
	logger.LogEvent("INFO", message.ServerStarted)
	//start server
	_ = router.Run(":" + configuration.GlobalConfig.ServicePort)
	//api.SetConfiguration
}

func startMongo() mongoRepository.MongoRepositories {
	logger.LogEvent("INFO", "Initializing Mongo!")
	mongoRepo, err := mongoRepository.ConnectToMongo(configuration.GlobalConfig.DbType, configuration.GlobalConfig.MongoDbUserName,
		configuration.GlobalConfig.MongoDbPassword, configuration.GlobalConfig.MongoDbHost,
		configuration.GlobalConfig.MongoDbPort, configuration.GlobalConfig.MongoDbAuthDb,
		configuration.GlobalConfig.MongoDbName)
	if err != nil {
		fmt.Println(err)
		logger.LogEvent("ERROR", "MongoDB database Connection Error: "+err.Error())
		log.Fatal()
	}
	return mongoRepo
}
