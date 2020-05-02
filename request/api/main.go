package api

import (
	"bytes"
	"net/http"

	"encoding/json"

	"github.com/JOIN-M-Y/server/config"
	profile "github.com/JOIN-M-Y/server/profile/model"
	"github.com/JOIN-M-Y/server/study/body"
	study "github.com/JOIN-M-Y/server/study/model"
)

// Interface api
type Interface interface {
	GetProfileByAccessToken(accessToken string) (*profile.Profile, error)
	GetStudyByID(studyID string) (*study.Study, error)
	UpdateStudy(accessToken string, studyModel *study.Study) (*study.Study, error)
}

// API api struct
type API struct {
	profileAPIURL string
	studyAPIURL   string
}

// New create api instance
func New(config config.Interface) *API {
	return &API{
		profileAPIURL: config.Server().ProfileServiceEndPoint(),
		studyAPIURL:   config.Server().StudyServiceEndPoint(),
	}
}

// UpdateStudy http put request to update study
func (api *API) UpdateStudy(accessToken string, studyModel *study.Study) (*study.Study, error) {
	memberProfileIDList := []string{}
	for _, memberProfile := range studyModel.MembersProfile {
		memberProfileIDList = append(memberProfileIDList, memberProfile.ID)
	}
	responseBody := body.UpdateStudy{
		Title:                  studyModel.Title,
		Description:            studyModel.Description,
		Recruitment:            studyModel.Recruitment,
		RecruitEndDate:         studyModel.RecruitEndDate,
		Public:                 studyModel.Public,
		AddressFirstDepthName:  studyModel.AddressFirstDepthName,
		AddressSecondDepthName: studyModel.AddressSecondDepthName,
		InterestedField:        studyModel.InterestedField,
		InterestedFieldDetail:  studyModel.InterestedFieldDetail,
		MembersProfileID:       memberProfileIDList,
	}
	byteData, err := json.Marshal(responseBody)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("PUT", api.studyAPIURL+"/"+studyModel.ID, bytes.NewBuffer(byteData))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Authorization", accessToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var updatedStudy *study.Study
	err = decoder.Decode(&updatedStudy)
	if err != nil {
		panic(err)
	}
	return updatedStudy, err
}

// GetStudyByID http request to get study
func (api *API) GetStudyByID(studyID string) (*study.Study, error) {
	response, err := http.Get(api.studyAPIURL + "/" + studyID)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var study *study.Study
	err = decoder.Decode(&study)
	if err != nil {
		panic(err)
	}
	return study, err
}

// GetProfileByAccessToken request get profile to profile service
func (api *API) GetProfileByAccessToken(accessToken string) (*profile.Profile, error) {
	request, err := http.NewRequest("GET", api.profileAPIURL, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Authorization", accessToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var profileList []*profile.Profile
	err = decoder.Decode(&profileList)
	if err != nil {
		panic(err)
	}

	return profileList[0], err
}
