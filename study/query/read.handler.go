package query

import (
	"errors"

	"github.com/JOIN-M-Y/server/study/model"
)

func (bus *Bus) handleReadStudyByIDQuery(
	query *ReadStudyByIDQuery,
) ([]model.Study, error) {
	studyEntity, err := bus.repository.FindByID(
		query.StudyID,
	)
	if err != nil {
		return nil, err
	}
	if studyEntity.ID == "" {
		return nil, errors.New("study is not found")
	}
	studyList := []model.Study{}
	study := bus.entityToModel(studyEntity)
	studyList = append(studyList, *study)
	return studyList, nil
}

func (bus *Bus) handleReadStudyByOwnerProfileID(
	query *ReadStudyByOwnerProfileID,
) ([]model.Study, error) {
	studyEntityList, err := bus.repository.FindByOwnerProfileID(
		query.OwnerProfileID,
	)
	if err != nil {
		return nil, err
	}
	studyList := []model.Study{}
	for _, studyEntity := range studyEntityList {
		study := bus.entityToModel(studyEntity)
		studyList = append(studyList, *study)
	}
	return studyList, nil
}

func (bus *Bus) handleRead(query *ReadStudyQuery) ([]model.Study, error) {
	studyEntityList, err := bus.repository.Find(query.Limit, query.Cursor, query.Interested)
	if err != nil {
		panic(err)
	}

	studyList := []model.Study{}
	for _, studyEntity := range studyEntityList {
		study := bus.entityToModel(studyEntity)
		studyList = append(studyList, *study)
	}
	return studyList, nil
}
