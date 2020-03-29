package controller

import (
	"net/http"

	"github.com/JOIN-M-Y/server/profile/command"
	"github.com/JOIN-M-Y/server/profile/dto"
	"github.com/JOIN-M-Y/server/profile/query"
	"github.com/gin-gonic/gin"
)

// @Description update profile
// @Tags Profiles
// @Accept json
// @Produce json
// @Param UpdateProfile body body.UpdateProfile true "update profile data"
// @Success 200 {object} model.Profile
// @Router /profiles [put]
// @Security AccessToken
func (controller *Controller) update(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")

	account, err := controller.GetAccountByAccessToken(accessToken)
	if account.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	var data dto.UpdateProfile
	data.AccountID = account.ID

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	query := &query.ReadProfileByAccountIDQuery{
		AccountID: data.AccountID,
	}
	profileList, err := controller.queryBus.Handle(query)
	profile := profileList[0]
	if profile.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	command := &command.UpdateProfileCommand{
		ID:                    profile.ID,
		FileID:                data.FileID,
		InterestedField:       data.InterestedField,
		InterestedFieldDetail: data.InterestedFieldDetail,
	}
	updatedProfile, err := controller.commandBus.Handle(command)
	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, updatedProfile)
}
