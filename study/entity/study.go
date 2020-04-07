package entity

import "time"

// Study study entity
type Study struct {
	ID                     string     `json:"_id" bson:"_id"`
	Title                  string     `json:"title" bson:"title"`
	Description            string     `json:"description" bson:"description"`
	Recruitment            int        `json:"recruitment" bson:"recruitment"`
	RecruitEndDate         time.Time  `json:"recruitEndDate" bson:"recruitEndDate"`
	Public                 bool       `json:"public" bson:"public"`
	AddressFirstDepthName  string     `json:"addressFirstDepthName" bson:"addressFirstDepthName"`
	AddressSecondDepthName string     `json:"addressSecondDepthName" bson:"addressSecondDepthName"`
	InterestedField        string     `json:"interestedField" bson:"interestedField"`
	InterestedFieldDetail  []string   `json:"interestedFieldDetail" bson:"interestedFieldDetail"`
	OwnerProfileID         string     `json:"ownerProfileId" bson:"ownerProfileId"`
	MembersProfileID       []string   `json:"membersProfileId" bson:"membersProfileId"`
	CreatedAt              time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt              time.Time  `json:"updatedAt" bson:"updatedAt"`
	DeletedAt              *time.Time `json:"deletedAt" bson:"deletedAt"`
}
