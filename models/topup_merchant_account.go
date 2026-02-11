package models

type TopupMerchantAccountRequest struct {
	MerchantAccountNumber int    `json:"merchantAccountNumber"`
	ConnectorID           int    `json:"connectorId"`
	TopupAmount           int    `json:"topupAmount"`
	CreatedBy             string `json:"createdBy"`
}

type TopupMerchantAccountResponse struct {
	Message string `json:"message"`
}