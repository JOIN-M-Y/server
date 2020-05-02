package controller

import (
	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/request/command"
	"github.com/JOIN-M-Y/server/request/query"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
)

// Controller request controller struct
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	queryBus   *query.Bus
	config     config.Interface
	util       *util.Util
}

// New create request controller instance
func New(route *gin.Engine, commandBus *command.Bus, queryBus *query.Bus, util *util.Util) *Controller {
	controller := &Controller{route: route, commandBus: commandBus, queryBus: queryBus, util: util}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup request route handler
func (controller *Controller) SetupRoutes() {
	controller.route.POST("requests", controller.create)
	controller.route.GET("requests", controller.readList)
	controller.route.PUT("requests/:requestId", controller.update)
}
