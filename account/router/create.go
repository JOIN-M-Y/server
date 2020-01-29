package router

import (
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/kyhsa93/gin-rest-example/account/dto"

	"github.com/gin-gonic/gin"
)

// @Description create account group
// @Tags Accounts
// @Accept  json
// @Produce  json
// @Param account body dto.Account true "Add account"
// @Success 200
// @Router /accounts [post]
func (router *Router) create(context *gin.Context) {
	var data dto.Account

	if bindError := context.ShouldBindJSON(&data); bindError != nil {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	if !emailAndProviderValidation(data.Email, data.Provider) {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	emaiFormatlValidationError := checkmail.ValidateFormat(data.Email)
	if emaiFormatlValidationError != nil {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	emaiHostlValidationError := checkmail.ValidateHost(data.Email)
	if emaiHostlValidationError != nil {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	_, existedProvider := dto.Provider()[data.Provider]

	if existedProvider == false {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	duplicated, _ := router.service.ReadAccount(data.Email, "", "", "", true)

	if duplicated != nil {
		httpError := router.util.Error.HTTP.Conflict()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	dto.FilterAccountAttributeByProvider(&data)

	validate := dto.ValidateAccountAttributeByProvider(&data)
	if validate == false {
		httpError := router.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	router.service.Create(data.Email, data.Provider, data.SocialID, data.Password)

	context.JSON(http.StatusCreated, "Account created")
}
