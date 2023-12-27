package line

import (
	"encoding/json"
	"net/http"
	"net/url"
	"user_management_service/internal/models"
)

type LineClientInterface interface {
	GetClientID() string
	GetClientSecret() string
	GetRedirectURI() string
	GetState() string
	GetAccessToken(code string) (result models.GetTokenResponse, err error)
	GetProfile(accessToken string) (result map[string]interface{}, err error)
}

type LineClient struct {
	clientID     string
	clientSecret string
	redirectURI  string
	state        string
	tokenURL     string
	profileURL   string
}

func New(clientID, clientSecret, redirectURI, state string) *LineClient {
	return &LineClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		state:        state,
		tokenURL:     "https://api.line.me/oauth2/v2.1/token",
		profileURL:   "https://api.line.me/v2/profile",
	}
}

func (l *LineClient) GetClientID() string {
	return l.clientID
}

func (l *LineClient) GetClientSecret() string {
	return l.clientID
}

func (l *LineClient) GetRedirectURI() string {
	return l.redirectURI
}

func (l *LineClient) GetState() string {
	return l.state
}

func (l *LineClient) GetAccessToken(code string) (result models.GetTokenResponse, err error) {
	data := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {l.redirectURI},
		"client_id":     {l.clientID},
		"client_secret": {l.clientSecret},
	}

	resp, err := http.PostForm(l.tokenURL, data)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}

func (l *LineClient) GetProfile(accessToken string) (result map[string]interface{}, err error) {
	req, err := http.NewRequest("GET", l.profileURL, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}
