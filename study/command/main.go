package command

import (
	"errors"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/file/repository"
	profile "github.com/JOIN-M-Y/server/profile/model"
	"github.com/JOIN-M-Y/server/study/api"
	"github.com/JOIN-M-Y/server/study/entity"
	"github.com/JOIN-M-Y/server/study/model"
)

// Bus study command
type Bus struct {
	repository repository.Interface
	config     config.Interface
	api        api.Interface
}

// New create Bus instance
func New(
	repository repository.Interface,
	config config.Interface,
	api api.Interface,
) *Bus {
	return &Bus{
		repository: repository,
		config:     config,
		api:        api,
	}
}

// Handle handle command
func (bus *Bus) Handle(
	command interface{},
) (*model.Study, error) {
	switch command := command.(type) {
	case *CreateStudyCommand:
		return bus.handleCreateCommand(command)
	default:
		return nil, errors.New("invalid command type")
	}
}

func (bus *Bus) entityToModel(
	entity entity.Study,
	ownerProfile profile.Profile,
	membersProfile []profile.Profile,
) *model.Study {
	var studyModel model.Study
	studyModel.ID = entity.ID
	studyModel.Title = entity.Title
	studyModel.Description = entity.Description
	studyModel.Recruitment = entity.Recruitment
	studyModel.Public = entity.Public
	studyModel.AddressFirstDepthName = entity.AddressFirstDepthName
	studyModel.AddressSecondDepthName = entity.AddressSecondDepthName
	studyModel.InterestedField = entity.InterestedField
	// ownderProfile = bus.api.GetProfileByAccessToken(
	// 	entity.OwnerProfileID
	// )
	return &studyModel
}
