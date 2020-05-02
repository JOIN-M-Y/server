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

	request, err := bus.repository.Update(command.RequestID)
	return bus.entityToModel(request), err
}
