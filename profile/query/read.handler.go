package query

import (
	"errors"

	"github.com/JOIN-M-Y/server/profile/model"
)

func (bus *Bus) handleReadProfileByIDQuery(
	query *ReadProfileByIDQuery,
) ([]*model.Profile, error) {
	profileEntity, err := bus.repository.FindByID(query.ProfileID)
	if err != nil {
		return nil, err
	}

	if profileEntity.ID == "" {
		return nil, errors.New("Profile is not found")
	}
	profileList := []*model.Profile{}
	profile := bus.entityToModel(profileEntity)
	profileList = append(profileList, profile)
	return profileList, nil
}

func (bus *Bus) handleReadProfileByAccountIDQuery(
	query *ReadProfileByAccountIDQuery,
) ([]*model.Profile, error) {
	profileEntity, err := bus.repository.FindByAccountID(
		query.AccountID,
	)
	if err != nil {
		return nil, err
	}
	profileList := []*model.Profile{}
	profile := bus.entityToModel(profileEntity)
	profileList = append(profileList, profile)
	return profileList, nil
}

func (bus *Bus) handleReadProfileByIDListQuery(
	query *ReadProfileByIDListQuery,
) ([]*model.Profile, error) {
	profileEntityList, err := bus.repository.FindByIDList(
		query.ProfileIDList,
	)
	if err != nil {
		return nil, err
	}
	profileList := []*model.Profile{}
	for _, profileEntity := range profileEntityList {
		profile := bus.entityToModel(profileEntity)
		profileList = append(profileList, profile)
	}
	return profileList, nil
}
