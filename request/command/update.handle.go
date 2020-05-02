package command

import (
	"errors"

	"github.com/JOIN-M-Y/server/request/model"
)

func (bus *Bus) handleUpdateCommand(command *UpdateRequestCommand) (*model.Request, error) {
	oldData := bus.repository.FindByID(command.RequestID)
	if oldData.ID == "" {
		return nil, errors.New("update target request is not found")
	}

	profile, err := bus.api.GetProfileByAccessToken(command.AccessToken)
	if err != nil {
		panic(err)
	}

	study, err := bus.api.GetStudyByID(oldData.StudyID)
	if err != nil {
		panic(err)
	}

	study.MembersProfile = append(study.MembersProfile, *profile)

	bus.api.UpdateStudy(command.AccessToken, study)

	request, err := bus.repository.Update(command.RequestID)
	return bus.entityToModel(request), err
}
