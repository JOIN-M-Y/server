package main

import (
	"log"

	"github.com/JOIN-M-Y/server/account"
	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/file"
	"github.com/JOIN-M-Y/server/profile"
	"github.com/JOIN-M-Y/server/util"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey AccessToken
// @in header
// @name Authorization

func main() {
	config := config.Initialize()
	util := util.Initialize()
	gin.SetMode(config.Server().Mode())
	route := gin.Default()
	account.Initialize(route, config, util)
	file.Initialize(route, config, util)
	profile.Initialize(route, config, util)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(route.Run(":" + config.Server().Port()))
}
