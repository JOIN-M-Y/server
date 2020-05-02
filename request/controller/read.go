package controller

import (
	"net/http"

	"github.com/JOIN-M-Y/server/request/query"
	"github.com/gin-gonic/gin"
)

// @Descritpion read request list
// @Tags Requests
// @Produce json
// @Param study_id query string true "studyId"
// @Success 200 {object} model.Request
// @Router /requests [get]
func (controller *Controller) readList(context *gin.Context) {
	studyID := context.Query("study_id")
	if studyID == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	query := &query.ReadRequestByStudyID{
		StudyID: studyID,
	}
	request, err := controller.queryBus.Handle(query)
	if err != nil {
		httpError := controller.util.Error.HTTP.InternalServerError()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, request)
}
