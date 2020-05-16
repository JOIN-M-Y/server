package controller

import (
	"net/http"
	"strconv"

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

// @Tags Studies
// @Accept json
// @Produce json
// @Param limit query int true "list count limit"
// @Param cursor query string true "pagenation cursor"
// @Param interested query string true "kind of list"
// @Success 200 {object} model.StudyList
// @Router /studies [get]
func (controller *Controller) read(context *gin.Context) {
	limitString := context.Query("limit")
	cursor := context.Query("cursor")
	interested := context.Query("interested")

	if interested == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), "interested is empty")
		return
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		panic(err)
	}
	query := &query.ReadStudyQuery{
		Limit:      limit,
		Cursor:     cursor,
		Interested: interested,
	}

	result, err := controller.queryBus.Handle(query)
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, result)
}
