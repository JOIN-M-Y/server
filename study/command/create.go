package command

import "time"

// CreateStudyCommand create study command
type CreateStudyCommand struct {
	Title                  string
	Description            string
	Recruitment            int
	RecruitEndDate         time.Time
	Public                 bool
	AddressFirstDepthName  string
	AddressSecondDepthName string
	InterestedField        string
	OwnerProfileID         string
}
