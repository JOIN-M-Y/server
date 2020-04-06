package entity

import "time"

// Group study group entity
type Group struct {
	ID    string `json:"_id" bson:"_id"`
	Owner struct {
		ProfileID string `json:"profileId" bson:"profileId"`
	}
	Title                 string     `json:"title" bson:"title"`
	Description           string     `json:"description" bson:"description"`
	Recruitment           int        `json:"recruitment" bson:"recruitment"`
	RecruitEndDate        time.Time  `json:"recruitEndDate" bson:"recruitEndDate"`
	PublicAbility         bool       `json:"public" bson:"public"`
	InterestedField       string     `json:"interestedField" bson:"interestedField"`
	Recommendation        int        `json:"recommendation" bson:"recommendation"`
	FirstRegionDepthName  string     `json:"firstRegionDepthName" bson:"firstRegionDepthName"`
	SecondRegionDepthName string     `json:"secondRegionDepthName" bson:"secondRegionDepthName"`
	CreatedAt             time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time  `json:"updatedAt" bson:"updatedAt"`
	DeletedAt             *time.Time `json:"deletedAt" bson:"deletedAt"`
}
