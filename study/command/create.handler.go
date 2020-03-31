package command

import (
	"github.com/JOIN-M-Y/server/study/model"
	"github.com/google/uuid"
)

func (bus *Bus) handleCreateCommand(
	command *CreateStudyCommand,
) (*model.Study, error) {
	uuid, _ := uuid.NewRandom()
	createStudyEntity, createError := bus.repository.Create(
		uuid.String(),
		command.Title,
		command.Description,
		command.Recruitment,
		command.RecruitEndDate,
		command.Public,
		command.AddressFirstDepthName,
		command.AddressSecondDepthName,
		command.InterestedField,
		command.OwnerProfileID,
	)
	if createError != nil {
		return nil, createError
	}
	return bus.entityToModel(createStudyEntity), nil
}
