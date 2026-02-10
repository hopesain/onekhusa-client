package models

type AccessTokenRequest struct {
	APIKey                string `json:"apiKey"`
	APISecret             string `json:"apiSecret"`
	OrganizationID        string `json:"organisationId"`
	MerchantAccountNumber int    `json:"merchantAccountNumber"`
}

type AccessTokenResponse struct {
	AccessToken     string `json:"accessToken"`
	ExpiresOn       string `json:"expiresOn"`
	ExpiryInMinutes int    `json:"expiryInMinutes"`
}
