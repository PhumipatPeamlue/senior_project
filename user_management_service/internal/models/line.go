package models

type GetTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type GetTokenBadRequest struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
