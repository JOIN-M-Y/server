package model

import "time"

// Request requested data to join study
type Request struct {
	ID        string    `json:"id" example:"requestId"`
	ProfileID string    `json:"profileId" example:"profileId"`
	StudyID   string    `json:"studyId" example:"studyId"`
	Status    string    `json:"status" example:"requested"`
	CreatedAt time.Time `json:"createdAt" example:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" example:"updatedAt"`
}
