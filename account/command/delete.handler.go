package command

import "github.com/JOIN-M-Y/server/account/model"

func (bus *Bus) handleDeleteCommand(command *DeleteCommand) (*model.Account, error) {
	err := bus.repository.Delete(command.AccountID)
	if err != nil {
		return nil, err
	}

	entity, err := bus.repository.FindByID(command.AccountID, true)
	if err != nil {
		return nil, err
	}

	return bus.entityToModel(*entity), nil
}
