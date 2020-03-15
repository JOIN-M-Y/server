package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JOIN-M-Y/server/account/api"
	"github.com/JOIN-M-Y/server/account/command"
	"github.com/JOIN-M-Y/server/account/controller"
	"github.com/JOIN-M-Y/server/account/query"
	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/util"
	"github.com/gin-gonic/gin"
)

// TestNew test controller's New method
func TestNew(t *testing.T) {
	_, engine := gin.CreateTestContext(httptest.NewRecorder())
	util := &util.Util{}
	commandBus := &command.Bus{}
	queryBus := &query.Bus{}
	config := &config.Config{}
	api := &api.API{}
	controllerInstance := controller.New(
		engine, commandBus, queryBus, util, config, api,
	)
	if controllerInstance == nil {
		t.Error("Can not create controller instance")
	}
}
