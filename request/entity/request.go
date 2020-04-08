package entity

import "time"

// Request requested to join study data
type Request struct {
	ID        string    `json:"_id" bson:"_id"`
	ProfileID string    `json:"profileId" bson:"profileId"`
	StudyID   string    `json:"studyId" bson:"studyId"`
	Status    string    `json:"status" bson:"status"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
