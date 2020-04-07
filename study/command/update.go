package command

import "time"

// UpdateStudyCommand update study command
type UpdateStudyCommand struct {
	StudyID                string
	Title                  string
	Description            string
	Recruitment            int
	RecruitEndDate         time.Time
	Public                 bool
	AddressFirstDepthName  string
	AddressSecondDepthName string
	InterestedField        string
	InterestedFieldDetail  []string
	MembersProfileID       []string
}
