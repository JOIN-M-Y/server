package controller

import (
	"net/http"

	"github.com/JOIN-M-Y/server/study/command"
	"github.com/JOIN-M-Y/server/study/dto"
	"github.com/JOIN-M-Y/server/study/query"
	"github.com/gin-gonic/gin"
)

// @Description update study
// @Tags Studies
// @Accept json
// @Produce json
// @Param id path string true "studyId"
// @Param UpdateStudy body body.UpdateStudy true "update study"
// @Success 200
// @Router /studies{id} [put]
// @Security AccessToken
func (controller *Controller) update(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	profile, err := controller.GetProfileByAccessToken(accessToken)
	if profile.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	var data dto.UpdateStudy
	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	query := &query.ReadStudyByIDQuery{
		StudyID: context.Param("id"),
	}

	studyList, err := controller.queryBus.Handle(
		query,
	)

	if len(studyList) == 0 || err != nil {
		httpError := controller.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	command := &command.UpdateStudyCommand{
		StudyID:                context.Param("id"),
		Title:                  data.Title,
		Description:            data.Description,
		Recruitment:            data.Recruitment,
		RecruitEndDate:         data.RecruitEndDate,
		Public:                 data.Public,
		AddressFirstDepthName:  data.AddressFirstDepthName,
		AddressSecondDepthName: data.AddressSecondDepthName,
		InterestedField:        data.InterestedField,
		InterestedFieldDetail:  data.InterestedFieldDetail,
		MembersProfileID:       data.MembersProfileID,
	}

	_, err = controller.commandBus.Handle(
		command,
	)
	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, nil)
}
