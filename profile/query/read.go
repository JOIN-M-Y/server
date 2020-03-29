package query

// ReadProfileByIDQuery read profile data by profile id
type ReadProfileByIDQuery struct {
	ProfileID string
}

// ReadProfileByIDListQuery read profile list by profileID list
type ReadProfileByIDListQuery struct {
	ProfileIDList []string
}

// ReadProfileByAccountIDQuery read profile data by accountID
type ReadProfileByAccountIDQuery struct {
	AccountID string
}
