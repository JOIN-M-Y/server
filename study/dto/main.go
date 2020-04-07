package dto

import (
	"errors"
	"time"
)

// CreateStudy dto for create study
type CreateStudy struct {
	Title                  string
	Description            string
	Recruitment            int
	RecruitEndDate         string
	Public                 bool
	AddressFirstDepthName  string
	AddressSecondDepthName string
	InterestedField        string
	InterestedFieldDetail  []string
	OwnerProfileID         string
}

// ValidationData validate dto attribute
func (dto *CreateStudy) ValidationData() error {
	if dto.Title == "" {
		return errors.New("title is empty")
	}
	if dto.Description == "" {
		return errors.New("description is empty")
	}
	if dto.Recruitment < 1 {
		return errors.New("recruitment is invalid")
	}
	if dto.InterestedField == "" {
		return errors.New("interestedField is empty")
	}
	if dto.OwnerProfileID == "" {
		return errors.New("owner profileId is empty")
	}
	return nil
}

// UpdateStudy dto for update study
type UpdateStudy struct {
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
