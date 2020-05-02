package body

// CreateRequest request body for POST /requests
type CreateRequest struct {
	StudyID string `json:"studyId" example:"studyId" binding:"required"`
}
