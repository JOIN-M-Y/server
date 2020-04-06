package group

import "time"

// Group study group model
type Group struct {
	ID                    string    `json:"id" example:"groupId"`
	Title                 string    `json:"title" example:"title"`
	Description           string    `json:"description" example:"description"`
	Recruitment           int       `json:"recruitment" example:"10"`
	RecruitEndDate        time.Time `json:"recruitEndDate" example:"2019-12-23 12:27:37"`
	PublicAbility         bool      `json:"publicAbility" example:"true"`
	InterestedField       string    `json:"interestedField" example:"develop"`
	Recommendation        int       `json:"recommendation" example:"5"`
	FirstRegionDepthName  string    `json:"firstRegionDepthName" example:"서울"`
	SecondRegionDepthName string    `json:"secondRegionDepthName" example:"강남구"`
}
