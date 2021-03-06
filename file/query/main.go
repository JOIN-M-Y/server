package query

import (
	"errors"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/file/entity"
	"github.com/JOIN-M-Y/server/file/model"
	"github.com/JOIN-M-Y/server/file/repository"
)

// Bus file query bus
type Bus struct {
	config     config.Interface
	repository repository.Interface
}

// New create Bus instance
func New(config config.Interface, repository repository.Interface) *Bus {
	return &Bus{config: config, repository: repository}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) (*model.File, error) {
	switch query := query.(type) {
	case *ReadFileByIDQuery:
		return bus.handleReadFileByIDQuery(query)
	default:
		return nil, errors.New("Query can not handled")
	}
}

func (bus *Bus) entityToModel(entity entity.File) *model.File {
	var fileModel model.File
	fileModel.ID = entity.ID
	fileModel.AccountID = entity.AccountID
	fileModel.Usage = entity.Usage
	fileModel.CreatedAt = entity.CreatedAt
	fileModel.ImageURL = bus.config.AWS().S3().Endpoint() + "/" +
		bus.config.AWS().S3().Bucket() + "/" + entity.ID
	return &fileModel
}
