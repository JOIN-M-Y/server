package query

import (
	"errors"

	"github.com/JOIN-M-Y/server/study/model"
)

func (bus *Bus) handleReadStudyByIDQuery(
	query *ReadStudyByIDQuery,
) (*model.Study, error) {
	studyEntity, err := bus.repository.FindByID(
		query.StudyID,
	)
	if err != nil {
		return nil, err
	}
	if studyEntity.ID == "" {
		return nil, errors.New("study is not found")
	}
	return bus.entityToModel(studyEntity), nil
}
