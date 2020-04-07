package command

import (
	"errors"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/study/api"
	"github.com/JOIN-M-Y/server/study/entity"
	"github.com/JOIN-M-Y/server/study/model"
	"github.com/JOIN-M-Y/server/study/repository"
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
	case *UpdateStudyCommand:
		return bus.handleUpdateCommand(command)
	default:
		return nil, errors.New("invalid command type")
	}
}

func (bus *Bus) entityToModel(
	entity entity.Study,
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
	studyModel.InterestedFieldDetail = entity.InterestedFieldDetail

	profileIDList := []string{}
	profileIDList = append(profileIDList, entity.OwnerProfileID)
	profileIDList = append(profileIDList, entity.MembersProfileID...)

	profileList, err := bus.api.GetProfileByProfileIDList(
		profileIDList,
	)
	if len(profileList) == 0 || err != nil {
		panic(err)
	}
	ownerProfile := profileList[0]
	studyModel.OwnerProfile = ownerProfile
	if 1 < len(profileList) {
		membersProfile := profileList[1:]
		for _, memberProfile := range membersProfile {
			studyModel.MembersProfile = append(studyModel.MembersProfile, memberProfile)
		}
	}
	return &studyModel
}
