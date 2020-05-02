package command

import (
	"github.com/JOIN-M-Y/server/request/model"
	"github.com/google/uuid"
)

func (bus *Bus) handleCreateCommand(
	command *CreateRequestCommand,
) (*model.Request, error) {
	uuid, _ := uuid.NewRandom()
	profile, err := bus.api.GetProfileByAccessToken(command.AccessToken)
	if err != nil {
		panic(err)
	}

	request, err := bus.repository.Create(
		uuid.String(),
		profile.ID,
		command.StudyID,
	)
	if err != nil {
		panic(err)
	}
	return bus.entityToModel(request), err
}
