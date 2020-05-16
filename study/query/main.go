package query

import (
	"errors"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/study/api"
	"github.com/JOIN-M-Y/server/study/entity"
	"github.com/JOIN-M-Y/server/study/model"
	"github.com/JOIN-M-Y/server/study/repository"
)

// Bus study query bus
type Bus struct {
	config     config.Interface
	repository repository.Interface
	api        api.Interface
}

// New create Bus instance
func New(
	config config.Interface,
	repository repository.Interface,
	api api.Interface,
) *Bus {
	return &Bus{
		config:     config,
		repository: repository,
		api:        api,
	}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) ([]model.Study, error) {
	switch query := query.(type) {
	case *ReadStudyByIDQuery:
		return bus.handleReadStudyByIDQuery(query)
	case *ReadStudyByOwnerProfileID:
		return bus.handleReadStudyByOwnerProfileID(query)
	case *ReadStudyQuery:
		return bus.handleRead(query)
	default:
		return nil, errors.New("query can not handled")
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
