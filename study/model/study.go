package model

import (
	"time"
)

// Study study model
type Study struct {
	ID                     string    `json:"id" example:"studyId"`
	Title                  string    `json:"title" example:"title"`
	Description            string    `json:"description" example:"description"`
	Recruitment            int       `json:"recruitment" example:"10"`
	RecruitEndDate         time.Time `json:"recruitEndDate" example:"2020-03-31T10:50:32.666Z"`
	Public                 bool      `json:"public" example:"true"`
	AddressFirstDepthName  string    `json:"addressFirstDepthName" example:"서울"`
	AddressSecondDepthName string    `json:"addressSecondDepthName" example:"강남"`
	InterestedField        string    `json:"interestedField" example:"develop"`
	InterestedFieldDetail  []string  `json:"interestedFieldDetail" example:"web,server"`
	OwnerProfile           struct {
		ID                    string    `json:"id" example:"profileId"`
		AccountID             string    `json:"accountId" example:"accountId"`
		ImageURL              string    `json:"imageUrl" example:"profile.image_url.com"`
		Gender                string    `json:"gender" example:"male"`
		InterestedField       string    `json:"interestedField" example:"develop"`
		InterestedFieldDetail []string  `json:"interestedFieldDetail" example:"web,server"`
		CreatedAt             time.Time `json:"createdAt" example:"2019-12-23 12:27:37"`
		UpdatedAt             time.Time `json:"updatedAt" example:"2019-12-23 12:27:37"`
	}
	MembersProfile []struct {
		ID                    string    `json:"id" example:"profileId"`
		AccountID             string    `json:"accountId" example:"accountId"`
		ImageURL              string    `json:"imageUrl" example:"profile.image_url.com"`
		Gender                string    `json:"gender" example:"male"`
		InterestedField       string    `json:"interestedField" example:"develop"`
		InterestedFieldDetail []string  `json:"interestedFieldDetail" example:"web,server"`
		CreatedAt             time.Time `json:"createdAt" example:"2019-12-23 12:27:37"`
		UpdatedAt             time.Time `json:"updatedAt" example:"2019-12-23 12:27:37"`
	}
}
