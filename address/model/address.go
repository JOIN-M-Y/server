package model

// Address address moddel
type Address []string

// GetFirstDepth get address data
func (address *Address) GetFirstDepth() *Address {
	keys := make(Address, len(address.RegionData()))
	var index int
	for key := range address.RegionData() {
		keys[index] = key
		index++
	}
	return &keys
}

// GetSecondDepth get address data
func (address *Address) GetSecondDepth(firstDepth string) *Address {
	data := address.RegionData()[firstDepth]
	return &data
}
