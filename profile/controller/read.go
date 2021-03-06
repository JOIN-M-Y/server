package controller

import (
	"net/http"

	"github.com/JOIN-M-Y/server/profile/model"
	"github.com/JOIN-M-Y/server/profile/query"
	"github.com/gin-gonic/gin"
)

// @Tags Profiles
// @Accept json
// @Produce json
// @Param id path string true "profile id"
// @Success 200 {object} model.Profile
// @Router /profiles/{id} [get]
func (controller *Controller) readByID(context *gin.Context) {
	id := context.Param("id")
	query := &query.ReadProfileByIDQuery{ProfileID: id}
	profile, _ := controller.queryBus.Handle(query)

	if profile == nil {
		httpError := controller.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, profile)
}

// @Tags Profiles
// @Accept json
// @Produce json
// @Success 200 {object} model.Profile
// @Router /profiles [get]
// @Param id query []string false "profileId list"
// @Security AccessToken
func (controller *Controller) read(context *gin.Context) {
	if _, existed := context.GetQueryArray("id"); existed == true {
		controller.readByProfileIDList(context)
		return
	}

	accessToken := context.GetHeader("Authorization")
	account, err := controller.GetAccountByAccessToken(accessToken)
	if account.ID == "" || err != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	query := &query.ReadProfileByAccountIDQuery{
		AccountID: account.ID,
	}
	profile, err := controller.queryBus.Handle(query)
	if err != nil {
		httpError := controller.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, profile)
}

func (controller *Controller) readByProfileIDList(
	context *gin.Context,
) {
	profileIDList, _ := context.GetQueryArray("id")
	query := &query.ReadProfileByIDListQuery{
		ProfileIDList: profileIDList,
	}
	profileList, err := controller.queryBus.Handle(query)
	if err != nil {
		profileList = []*model.Profile{}
		context.JSON(http.StatusOK, profileList)
		return
	}
	context.JSON(http.StatusOK, profileList)
}
