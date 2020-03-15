package command

import (
	"errors"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/file/api"
	"github.com/JOIN-M-Y/server/file/aws"
	"github.com/JOIN-M-Y/server/file/entity"
	"github.com/JOIN-M-Y/server/file/model"
	"github.com/JOIN-M-Y/server/file/repository"
)

// Bus file command bus
type Bus struct {
	repository repository.Interface
	api        api.Interface
	aws        aws.Interface
	config     config.Interface
}

// New create bus instance
func New(
	repository repository.Interface,
	api api.Interface,
	aws aws.Interface,
	config config.Interface,
) *Bus {
	return &Bus{
		repository: repository,
		api:        api,
		aws:        aws,
		config:     config,
	}
}

// Handle handle command
func (bus *Bus) Handle(command interface{}) (*model.File, error) {
	switch command := command.(type) {
	case *CreateCommand:
		return bus.handleCreateCommand(command)
	default:
		return nil, errors.New("Command is not handled")
	}
}

func (bus *Bus) entityToModel(entity entity.File) *model.File {
	var fileModel model.File
	fileModel.ID = entity.ID
	fileModel.AccountID = entity.AccountID
	fileModel.Usage = entity.Usage
	fileModel.CreatedAt = entity.CreatedAt
	imageURL := bus.config.AWS().S3().Endpoint() +
		"/" + bus.config.AWS().S3().Bucket() + "/" + entity.ID
	fileModel.ImageURL = imageURL

	return &fileModel
}
