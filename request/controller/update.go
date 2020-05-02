package controller

import (
	"net/http"

	"github.com/JOIN-M-Y/server/request/command"
	"github.com/gin-gonic/gin"
)

// @Description update request
// @Tags Requests
// @Produce json
// @Param id path string true "requestId"
// @Success 200
// @Router /requests/{id} [put]
// @Security AccessToken
func (controller *Controller) update(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")
	if accessToken == "" {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	requestID := context.Param("requestId")
	if requestID == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	command := &command.UpdateRequestCommand{
		RequestID:   requestID,
		AccessToken: accessToken,
	}
	request, err := controller.commandBus.Handle(command)
	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, request)
}
