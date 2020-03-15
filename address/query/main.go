package query

import (
	"errors"

	"github.com/JOIN-M-Y/server/address/model"
	"github.com/JOIN-M-Y/server/config"
)

// Bus address query bus
type Bus struct {
	config config.Interface
}

// New create Bus instance
func New(config config.Interface) *Bus {
	return &Bus{config: config}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) (*model.Address, error) {
	switch query := query.(type) {
	case *ReadAddressQuery:
		return bus.handleReadAddressQuery(query)
	case *ReadAddressByFirstRegionNameQuery:
		return bus.handleReadAddressByFirstRegionNameQuery(query)
	default:
		return nil, errors.New("Query can not handled")
	}
}
