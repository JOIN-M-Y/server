package query

import (
	"errors"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/request/entity"
	"github.com/JOIN-M-Y/server/request/model"
	"github.com/JOIN-M-Y/server/request/repository"
)

// Bus request query bus
type Bus struct {
	config     config.Interface
	repository repository.Interface
}

// New create Bus instance
func New(config config.Interface, repository repository.Interface) *Bus {
	return &Bus{config: config, repository: repository}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) ([]*model.Request, error) {
	switch query := query.(type) {
	case *ReadRequestByStudyID:
		return bus.handleReadRequestByStudyID(query)
	default:
		return nil, errors.New("invalid query type")
	}
}

func (bus *Bus) entityToModel(entity entity.Request) *model.Request {
	var requestModel model.Request
	requestModel.ID = entity.ID
	requestModel.ProfileID = entity.ProfileID
	requestModel.StudyID = entity.StudyID
	requestModel.Status = entity.Status
	requestModel.CreatedAt = entity.CreatedAt
	requestModel.UpdatedAt = entity.UpdatedAt
	return &requestModel
}
