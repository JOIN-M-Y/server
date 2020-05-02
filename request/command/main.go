package command

import (
	"errors"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/request/api"
	"github.com/JOIN-M-Y/server/request/entity"
	"github.com/JOIN-M-Y/server/request/model"
	"github.com/JOIN-M-Y/server/request/repository"
)

// Bus request commnad
type Bus struct {
	repository repository.Interface
	api        api.Interface
	config     config.Interface
}

// New create Bus instance
func New(
	repository repository.Interface,
	api api.Interface,
	config config.Interface,
) *Bus {
	return &Bus{
		repository: repository,
		api:        api,
		config:     config,
	}
}

// Handle handle command
func (bus *Bus) Handle(command interface{}) (*model.Request, error) {
	switch command := command.(type) {
	case *CreateRequestCommand:
		return bus.handleCreateCommand(command)
	default:
		return nil, errors.New("invalid command type")
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
