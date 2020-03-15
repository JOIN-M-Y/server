package address

import (
	"github.com/JOIN-M-Y/server/address/api"
	"github.com/JOIN-M-Y/server/address/controller"
	"github.com/JOIN-M-Y/server/address/query"
	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
)

// Initialize init address module
func Initialize(
	engine *gin.Engine, config config.Interface, util *util.Util,
) {
	api := api.New(config)
	queryBus := query.New(config)
	controller.New(engine, queryBus, util, api)
}
