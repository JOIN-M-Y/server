package api

import (
	"net/http"

	"encoding/json"

	"github.com/JOIN-M-Y/server/config"
	"github.com/JOIN-M-Y/server/profile/model"
)

// Interface api
type Interface interface {
	GetProfileByAccessToken(accessToken string) (*model.Profile, error)
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

// GetProfileByAccessToken request get profile to profile service
func (api *API) GetProfileByAccessToken(accessToken string) (*model.Profile, error) {
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
	var profileList []*model.Profile
	err = decoder.Decode(&profileList)
	if err != nil {
		panic(err)
	}

	return profileList[0], err
}
