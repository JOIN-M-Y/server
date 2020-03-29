package query

import (
	"errors"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/profile/entity"
	"github.com/JOIN-M-Y/server/profile/model"
	"github.com/JOIN-M-Y/server/profile/repository"
)

// Bus profile query bus
type Bus struct {
	config     config.Interface
	repository repository.Interface
}

// New create bus instance
func New(config config.Interface, repository repository.Interface) *Bus {
	return &Bus{config: config, repository: repository}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) ([]*model.Profile, error) {
	switch query := query.(type) {
	case *ReadProfileByIDQuery:
		return bus.handleReadProfileByIDQuery(query)
	case *ReadProfileByAccountIDQuery:
		return bus.handleReadProfileByAccountIDQuery(query)
	case *ReadProfileByIDListQuery:
		return bus.handleReadProfileByIDListQuery(query)
	default:
		return nil, errors.New("Invalid query")
	}
}

func (bus *Bus) entityToModel(entity entity.Profile) *model.Profile {
	var profileModel model.Profile
	profileModel.ID = entity.ID
	profileModel.AccountID = entity.AccountID
	profileModel.Gender = entity.Gender
	profileModel.InterestedField = entity.InterestedField
	profileModel.InterestedFieldDetail = entity.InterestedFieldDetail
	profileModel.CreatedAt = entity.CreatedAt
	profileModel.UpdatedAt = entity.UpdatedAt
	profileModel.ImageURL = bus.config.AWS().S3().Endpoint() + "/" +
		bus.config.AWS().S3().Bucket() + "/" + entity.FileID
	if entity.FileID == "" {
		profileModel.ImageURL = ""
	}
	return &profileModel
}
