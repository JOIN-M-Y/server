package command

import (
	"errors"

	"github.com/JOIN-M-Y/server/study/model"
)

func (bus *Bus) handleUpdateCommand(
	command *UpdateStudyCommand,
) (*model.Study, error) {
	oldData, err := bus.repository.FindByID(
		command.StudyID,
	)
	if oldData.ID == "" || err != nil {
		return nil, errors.New("update target study data is not found")
	}
	updatedStudyEntity, err := bus.repository.Update(
		command.StudyID,
		command.Title,
		command.Description,
		command.Recruitment,
		command.RecruitEndDate,
		command.Public,
		command.AddressFirstDepthName,
		command.AddressSecondDepthName,
		command.InterestedField,
		command.InterestedFieldDetail,
		command.MembersProfileID,
	)
	if err != nil {
		return nil, err
	}
	studyModel := bus.entityToModel(updatedStudyEntity)
	return studyModel, nil
}
