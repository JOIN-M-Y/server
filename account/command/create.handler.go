package command

import (
	"github.com/JOIN-M-Y/server/account/entity"
	"github.com/JOIN-M-Y/server/account/model"
	"github.com/google/uuid"
)

func (bus *Bus) handleCreateCommand(
	command *CreateCommand,
) (*model.Account, error) {
	uuid, _ := uuid.NewRandom()
	hashedPassword, hashedSocialID :=
		getHashedPasswordAndSocialID(command.Password, command.SocialID)

	entity := entity.Account{
		ID:       uuid.String(),
		Email:    command.Email,
		Provider: command.Provider,
		SocialID: hashedSocialID,
		Password: hashedPassword,
		FCMToken: command.FCMToken,
		Gender:   command.Gender,
	}

	err := bus.repository.Create(&entity)
	if err != nil {
		return nil, err
	}
	bus.email.Send([]string{command.Email}, "Account is created.")
	accountModel := bus.entityToModel(entity)
	accountModel.CreateAccessToken(
		bus.config.Auth().AccessTokenSecret(),
		bus.config.Auth().AccessTokenExpiration(),
	)
	profile, err := bus.api.CreateProfile(
		accountModel.AccessToken,
		accountModel.ID,
		command.Email,
		command.Gender,
		command.InterestedField,
		command.InterestedFieldDetail,
	)

	if err != nil || profile == nil {
		panic(err)
	}
	return accountModel, nil
}
