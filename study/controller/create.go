package controller

import (
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

	data.OwnerProfileID = profile.ID

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if validationError := data.ValidationData(); validationError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

}
