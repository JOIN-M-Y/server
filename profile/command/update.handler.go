package command

import (
	"errors"

	"github.com/JOIN-M-Y/server/profile/model"
)

func (bus *Bus) handleUpdateCommand(
	command *UpdateProfileCommand,
) (*model.Profile, error) {
	oldData, err := bus.repository.FindByID(command.ID)
	if oldData.ID == "" || err != nil {
		return nil, errors.New("update target profile data is not found")
	}

	updatedProfileEntity, err := bus.repository.Update(
		command.ID,
		command.InterestedField,
		command.InterestedFieldDetail,
		command.FileID,
	)
	if err != nil {
		return nil, err
	}
	profileModel := bus.entityToModel(updatedProfileEntity)
	return profileModel, nil
}
