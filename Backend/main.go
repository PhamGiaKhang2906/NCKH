package main

import (
	"Backend/api"
	"Backend/config"
	"Backend/initialize"
	"Backend/migrate"

	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnv()
	initialize.ConnectDB()
}

func main() {
	migrate.MigrateData()

	router := gin.Default()

	config.SetupCORS(router)

	api.SetupRouter(router)

	router.Run()
}
