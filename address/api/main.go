package api

import (
	"encoding/json"
	"net/http"

	account "github.com/JOIN-M-Y/server/account/model"
	"github.com/JOIN-M-Y/server/config"
)

// Interface extermal api interface
type Interface interface {
	GetAccountByAccessToken(
		accessToken string,
	) (*account.Account, error)
}

// API api struct
type API struct {
	accountAPIURL string
}

// New create api instance
func New(config config.Interface) *API {
	return &API{
		accountAPIURL: config.Server().AccountServiceEndPoint(),
	}
}

// GetAccountByAccessToken get account data from account service by accessToken
func (api *API) GetAccountByAccessToken(
	accessToken string,
) (*account.Account, error) {
	accountServiceEndpoint := api.accountAPIURL
	request, createNewRequestError := http.NewRequest(
		"GET",
		accountServiceEndpoint,
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
	var account *account.Account
	responseBodyDecodeError := decoder.Decode(&account)
	if responseBodyDecodeError != nil {
		return nil, responseBodyDecodeError
	}
	return account, nil
}
