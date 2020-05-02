package query

import "github.com/JOIN-M-Y/server/request/model"

func (bus *Bus) handleReadRequestByStudyID(query *ReadRequestByStudyID) ([]*model.Request, error) {
	entityList, err := bus.repository.FindByStudyID(query.StudyID)
	if err != nil {
		panic(err)
	}
	modelList := []*model.Request{}
	for _, entity := range entityList {
		model := bus.entityToModel(entity)
		modelList = append(modelList, model)
	}
	return modelList, err
}
