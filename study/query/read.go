package query

// ReadStudyByIDQuery read study by studyID
type ReadStudyByIDQuery struct {
	StudyID string
}

// ReadStudyByOwnerProfileID read study by study owner profileID
type ReadStudyByOwnerProfileID struct {
	OwnerProfileID string
}
