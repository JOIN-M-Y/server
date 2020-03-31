package controller

import (
	"net/http"
	"time"

	"github.com/JOIN-M-Y/server/study/command"
	"github.com/JOIN-M-Y/server/study/dto"
	"github.com/gin-gonic/gin"
)

// @Description create study
// @Tags Studies
// @Accept json
// @Produce json
// @Param CreateStudy body body.CreateStudy true "Create Study data"
// @Success 201 {object} model.Study
// @Router /studies [post]
// @Security AccessToken
func (controller *Controller) create(context *gin.Context) {
	profile, err := controller.GetProfileByAccessToken(
		context.GetHeader("Authorization"),
	)
	if profile.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	var data dto.CreateStudy

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	data.OwnerProfileID = profile.ID

	if validationError := data.ValidationData(); validationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	parsedTime, err := time.Parse(time.RFC3339, data.RecruitEndDate)
	if err != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	command := &command.CreateStudyCommand{
		Title:                  data.Title,
		Description:            data.Description,
		Recruitment:            data.Recruitment,
		RecruitEndDate:         parsedTime,
		Public:                 data.Public,
		AddressFirstDepthName:  data.AddressFirstDepthName,
		AddressSecondDepthName: data.AddressSecondDepthName,
		InterestedField:        data.InterestedField,
		OwnerProfileID:         data.OwnerProfileID,
	}

	createdStudy, handlingError := controller.commandBus.Handle(command)
	if handlingError != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	context.JSON(http.StatusOK, createdStudy)
}
