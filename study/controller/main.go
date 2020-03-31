package controller

import (
	"errors"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/profile/model"
	"github.com/JOIN-M-Y/server/study/api"
	"github.com/JOIN-M-Y/server/study/command"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
)

// Controller study controller strcut
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	util       *util.Util
	config     config.Interface
	api        api.Interface
}

// New create study controller instance
func New(
	route *gin.Engine,
	commandBus *command.Bus,
	util *util.Util,
	config config.Interface,
	api api.Interface,
) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: commandBus,
		util:       util,
		config:     config,
		api:        api,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup study route handler
func (controller *Controller) SetupRoutes() {
	controller.route.POST("studies", controller.create)
}

// GetProfileByAccessToken get profile data
func (controller *Controller) GetProfileByAccessToken(
	accessToken string,
) (model.Profile, error) {
	if accessToken == "" {
		return model.Profile{}, errors.New("token is empty")
	}
	profile, err := controller.api.GetProfileByAccessToken(
		accessToken,
	)
	if profile == nil || err != nil {
		return model.Profile{}, errors.New("token is invalid")
	}
	return *profile, nil
}
