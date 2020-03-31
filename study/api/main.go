package api

import (
	"encoding/json"
	"net/http"

	"github.com/JOIN-M-Y/server/config"
	profile "github.com/JOIN-M-Y/server/profile/model"
)

// Interface external api interface
type Interface interface {
	GetProfileByAccessToken(
		accessToken string,
	) (*profile.Profile, error)
	GetProfileByProfileIDList(
		profileIDList []string,
	) ([]*profile.Profile, error)
}

// API api struct
type API struct {
	profileAPIURL string
}

// New create api instance
func New(config config.Interface) *API {
	return &API{
		profileAPIURL: config.Server().ProfileServiceEndPoint(),
	}
}

// GetProfileByProfileIDList get profile list by profileId list
func (api *API) GetProfileByProfileIDList(
	profileIDList []string,
) ([]*profile.Profile, error) {
	profileServiceEndpoint := api.profileAPIURL
	request, createNewRequestError := http.NewRequest(
		"GET",
		profileServiceEndpoint,
		nil,
	)
	if createNewRequestError != nil {
		return nil, createNewRequestError
	}
	query := request.URL.Query()
	for _, profileID := range profileIDList {
		query.Add("id", profileID)
	}
	request.URL.RawQuery = query.Encode()
	client := &http.Client{}
	response, httpRequestError := client.Do(request)
	if httpRequestError != nil {
		return nil, httpRequestError
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var profileList []*profile.Profile
	responseBodyDecodeError := decoder.Decode(&profileList)
	if responseBodyDecodeError != nil {
		return nil, responseBodyDecodeError
	}
	return profileList, nil
}

// GetProfileByAccessToken get profile data from profile service by accesstoken
func (api *API) GetProfileByAccessToken(
	accessToken string,
) (*profile.Profile, error) {
	profileServiceEndpoint := api.profileAPIURL
	request, createNewRequestError := http.NewRequest(
		"GET",
		profileServiceEndpoint,
		nil,
	)
	if createNewRequestError != nil {
		return nil, createNewRequestError
	}

	request.Header.Add("Authorization", accessToken)

	client := &http.Client{}
	response, httpRequestError := client.Do(request)
	if httpRequestError != nil {
		return nil, httpRequestError
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var profileList []*profile.Profile
	responseBodyDecodeError := decoder.Decode(&profileList)
	if responseBodyDecodeError != nil {
		return nil, responseBodyDecodeError
	}
	return profileList[0], nil
}
