package controller

import (
	"errors"

	"github.com/JOIN-M-Y/server/account/model"
	"github.com/JOIN-M-Y/server/address/api"
	"github.com/JOIN-M-Y/server/address/query"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
)

// Controller address controller strcut
type Controller struct {
	route    *gin.Engine
	queryBus *query.Bus
	util     *util.Util
	api      api.Interface
}

// New create address controller instance
func New(
	route *gin.Engine,
	queryBus *query.Bus,
	util *util.Util,
	api api.Interface,
) *Controller {
	controller := &Controller{
		route:    route,
		queryBus: queryBus,
		util:     util,
		api:      api,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup address route handler
func (controller *Controller) SetupRoutes() {
	controller.route.GET("address", func(context *gin.Context) {
		controller.read(context)
	})
	controller.route.GET("address/:first_depth_name", func(context *gin.Context) {
		controller.readSecondDepth(context)
	})
}

// GetAccountByAccessToken check http request auth
func (controller *Controller) GetAccountByAccessToken(
	accessToken string,
) (model.Account, error) {
	if accessToken == "" {
		return model.Account{}, errors.New("token is empty")
	}
	account, err := controller.api.GetAccountByAccessToken(
		accessToken,
	)
	if account == nil || err != nil {
		return model.Account{}, errors.New("token is invalid")
	}
	return *account, nil
}
