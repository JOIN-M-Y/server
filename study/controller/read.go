package controller

import (
	"net/http"

	"github.com/JOIN-M-Y/server/study/query"
	"github.com/gin-gonic/gin"
)

// @Tags Studies
// @Accept json
// @Produce json
// @Param id path string true "studyId"
// @Success 200 {object} model.Study
// @Router /studies/{id} [get]
func (controller *Controller) readByID(context *gin.Context) {
	id := context.Param("id")
	query := &query.ReadStudyByIDQuery{
		StudyID: id,
	}
	studyList, _ := controller.queryBus.Handle(
		query,
	)
	if studyList == nil || studyList[0].ID == "" {
		httpError := controller.util.Error.HTTP.NotFound()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	context.JSON(http.StatusOK, studyList[0])
}
