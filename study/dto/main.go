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
	RecruitEndDate         time.Time
	Public                 bool
	AddressFirstDepthName  string
	AddressSecondDepthName string
	InterestedField        string
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
	if time.Now().After(dto.RecruitEndDate) {
		return errors.New("recruitEndDate is invalid")
	}
	if dto.InterestedField == "" {
		return errors.New("interestedField is empty")
	}
	if dto.OwnerProfileID == "" {
		return errors.New("owner profileId is empty")
	}
	return nil
}
