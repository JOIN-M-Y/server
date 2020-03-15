package controller

import (
	"errors"

	"github.com/JOIN-M-Y/server/account/api"
	"github.com/JOIN-M-Y/server/account/command"
	"github.com/JOIN-M-Y/server/account/model"
	"github.com/JOIN-M-Y/server/account/query"
	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
)

// Controller account controller struct
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	queryBus   *query.Bus
	util       *util.Util
	config     config.Interface
	api        api.Interface
}

// New create account controller instance
func New(
	route *gin.Engine,
	commandBus *command.Bus,
	queryBus *query.Bus,
	util *util.Util,
	config config.Interface,
	api api.Interface,
) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: commandBus,
		queryBus:   queryBus,
		util:       util,
		config:     config,
		api:        api,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup accounts route handler
func (controller *Controller) SetupRoutes() {
	controller.route.POST("accounts", func(context *gin.Context) {
		controller.create(context)
	})

	controller.route.GET("accounts", func(context *gin.Context) {
		controller.readAccount(context)
	})

	controller.route.PUT("accounts", func(context *gin.Context) {
		controller.update(context)
	})

	controller.route.DELETE("accounts", func(context *gin.Context) {
		controller.delete(context)
	})
}

// GetAccountByAccessToken get account data by accesstoken
func (controller *Controller) GetAccountByAccessToken(
	accessToken string,
) (model.Account, error) {
	if accessToken == "" {
		return model.Account{}, errors.New("token is empty")
	}

	account := &model.Account{AccessToken: accessToken}

	accountID, err := account.GetTokenIssuer(
		controller.config.Auth().AccessTokenSecret(),
	)
	if accountID == "" || err != nil {
		return model.Account{}, errors.New("token is invalid")
	}

	query := &query.ReadAccountByIDQuery{AccountID: accountID}
	account, queryError := controller.queryBus.Handle(query)
	if queryError != nil {
		return model.Account{}, errors.New("account query error is occurred")
	}
	return *account, nil
}
