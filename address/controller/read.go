package controller

import (
	"net/http"

	"github.com/JOIN-M-Y/server/address/query"
	"github.com/gin-gonic/gin"
)

// @Tags Address
// @Accept json
// @Produce json
// @Success 200 {object} model.Address
// @Router /address [get]
func (controller *Controller) read(context *gin.Context) {
	query := &query.ReadAddressQuery{}
	address, err := controller.queryBus.Handle(query)
	if address == nil || err != nil {
		httpError := controller.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, address)
}

// @Tags Address
// @Accept json
// @Produce json
// @Success 200 {object} model.Address
// @Router /address/{first_depth_name} [get]
// @Param first_depth_name path string true "region first depth name"
func (controller *Controller) readSecondDepth(context *gin.Context) {
	firstRegionName := context.Param("first_depth_name")
	query := &query.ReadAddressByFirstRegionNameQuery{
		FirstRegionName: firstRegionName,
	}
	address, err := controller.queryBus.Handle(query)
	if address == nil || err != nil {
		httpError := controller.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, address)
}
