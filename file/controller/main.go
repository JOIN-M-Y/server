package controller

import (
	"errors"

	"github.com/JOIN-M-Y/server/account/model"
	"github.com/JOIN-M-Y/server/file/api"
	"github.com/JOIN-M-Y/server/file/command"
	"github.com/JOIN-M-Y/server/file/query"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
)

// Controller file controller struct
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	queryBus   *query.Bus
	util       *util.Util
	api        api.Interface
}

// New create file controller instance
func New(
	route *gin.Engine,
	commandBus *command.Bus,
	queryBus *query.Bus,
	util *util.Util,
	api api.Interface,
) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: commandBus,
		queryBus:   queryBus,
		util:       util,
		api:        api,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup files route handler
func (controller *Controller) SetupRoutes() {
	controller.route.POST("/files", func(context *gin.Context) {
		controller.create(context)
	})

	controller.route.GET("/files/:id", func(context *gin.Context) {
		controller.readFileByID(context)
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
