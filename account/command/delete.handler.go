package command

import "github.com/JOIN-M-Y/server/account/model"

func (bus *Bus) handleDeleteCommand(command *DeleteCommand) (*model.Account, error) {
	accountEntity := bus.repository.Delete(command.AccountID)
	return bus.entityToModel(accountEntity), nil
}
