package body

import "time"

// CreateStudy http request body POST /studies
type CreateStudy struct {
	Title                  string    `json:"title" example:"title"`
	Description            string    `json:"description" example:"description"`
	InterestedField        string    `json:"interestedField" example:"develop"`
	Recruitment            int       `json:"recruitment" example:"10"`
	RecruitEndDate         time.Time `json:"recruitEndDate" example:"2020-03-31T10:50:32.666Z"`
	Public                 bool      `json:"public" example:"true"`
	AddressFirstDepthName  string    `json:"addressFirstDepthName" example:"서울"`
	AddressSecondDepthName string    `json:"addressSecondDepthName" example:"강남"`
}

// UpdateStudy http request body PUT /studies
type UpdateStudy struct {
	Title                  string    `json:"title" example:"title"`
	Description            string    `json:"description" example:"description"`
	Recruitment            int       `json:"recruitment" example:"10"`
	RecruitEndDate         time.Time `json:"recruitEndDate" example:"2020-03-31T10:50:32.666Z"`
	Public                 bool      `json:"public" example:"true"`
	AddressFirstDepthName  string    `json:"addressFirstDepthName" example:"서울"`
	AddressSecondDepthName string    `json:"addressSecondDepthName" example:"강남"`
}
