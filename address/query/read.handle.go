package query

import "github.com/JOIN-M-Y/server/address/model"

func (bus *Bus) handleReadAddressQuery(
	query *ReadAddressQuery,
) (*model.Address, error) {
	address := &model.Address{}
	return address.GetFirstDepth(), nil
}

func (bus *Bus) handleReadAddressByFirstRegionNameQuery(
	query *ReadAddressByFirstRegionNameQuery,
) (*model.Address, error) {
	address := &model.Address{}
	return address.GetSecondDepth(query.FirstRegionName), nil
}
