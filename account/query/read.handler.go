package query

import (
	"errors"

	"github.com/JOIN-M-Y/server/account/model"
)

func (bus *Bus) handleReadAccountByIDQuery(
	query *ReadAccountByIDQuery,
) (*model.Account, error) {
	entity, err := bus.repository.FindByID(query.AccountID, false)
	if err != nil {
		return nil, err
	}
	if entity.ID == "" {
		return nil, errors.New("Account is not found")
	}

	model := bus.entityToModel(*entity)
	model.CreateAccessToken(
		bus.config.Auth().AccessTokenSecret(),
		bus.config.Auth().AccessTokenExpiration(),
	)
	return model, err
}

func (bus *Bus) handleReadAccountQuery(
	query *ReadAccountQuery,
) (*model.Account, error) {
	entityList, err := bus.repository.FindByEmailAndProvider(
		query.Email, query.Provider, query.Unscoped,
	)
	if err != nil {
		return &model.Account{}, err
	}

	entity := &entityList[0]

	if entity.ID == "" {
		return &model.Account{}, nil
	}

	if err := compareHashAndPassword(
		entity.Password,
		query.Password,
	); err != nil {
		return &model.Account{}, err
	}

	if err := compareHashAndPassword(
		entity.SocialID,
		query.SocialID,
	); err != nil {
		return &model.Account{}, err
	}

	model := bus.entityToModel(entity)
	model.CreateAccessToken(
		bus.config.Auth().AccessTokenSecret(),
		bus.config.Auth().AccessTokenExpiration())
	return model, nil
}

func (bus *Bus) handleReadAccountByEmailquery(
	query *ReadAccountByEmailQuery,
) (*model.Account, error) {
	entity := bus.repository.FindByEmail(query.Email)
	if entity.ID == "" {
		return &model.Account{}, nil
	}
	model := bus.entityToModel(entity)
	model.CreateAccessToken(
		bus.config.Auth().AccessTokenSecret(),
		bus.config.Auth().AccessTokenExpiration(),
	)
	return model, nil
}
