package controller

import (
	"github.com/JOIN-M-Y/server/address/query"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
)

// Controller address controller strcut
type Controller struct {
	route    *gin.Engine
	queryBus *query.Bus
	util     *util.Util
}

// New create address controller instance
func New(
	route *gin.Engine,
	queryBus *query.Bus,
	util *util.Util,
) *Controller {
	controller := &Controller{
		route:    route,
		queryBus: queryBus,
		util:     util,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup address route handler
func (controller *Controller) SetupRoutes() {
	controller.route.GET("address", controller.read)
	controller.route.GET("address/:first_depth_name", controller.readSecondDepth)
}
