package controller

import (
	"log"
	"net/http"

	"github.com/JOIN-M-Y/server/request/command"
	"github.com/JOIN-M-Y/server/request/dto"
	"github.com/gin-gonic/gin"
)

// @Description create request
// @Tags Requests
// @Accept json
// Produce json
// @Param CreateRequest body body.CreateRequest true "Create request data"
// @Success 201 {object} model.Request
// @Router /requests [post]
// @Security AccessToken
func (controller *Controller) create(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")
	log.Println("test", accessToken)
	if accessToken == "" {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	var data dto.CreateRequest
	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	command := &command.CreateRequestCommand{
		AccessToken: accessToken,
		StudyID:     data.StudyID,
	}
	request, err := controller.commandBus.Handle(command)
	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusCreated, request)
}
