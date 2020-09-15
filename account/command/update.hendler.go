package command

import (
	"errors"

	"github.com/JOIN-M-Y/server/account/model"
)

func (bus *Bus) handleUpdateCommand(
	command *UpdateCommand,
) (*model.Account, error) {
	oldData, err := bus.repository.FindByID(command.AccountID, false)
	if err != nil {
		return nil, err
	}
	if oldData.ID == "" {
		return nil, errors.New("Update target Account data is not found")
	}

	hashedPassword, _ :=
		getHashedPasswordAndSocialID(command.Password, "")

	if command.Password == "" {
		hashedPassword = oldData.Password
	}

	oldData.Password = hashedPassword
	oldData.FCMToken = command.FCMToken

	err = bus.repository.Update(oldData)
	if err != nil {
		return nil, err
	}

	bus.email.Send([]string{oldData.Email}, "Account is updated.")
	accountModel := bus.entityToModel(*oldData)
	accountModel.CreateAccessToken(
		bus.config.Auth().AccessTokenSecret(),
		bus.config.Auth().AccessTokenExpiration(),
	)
	return accountModel, nil
}
