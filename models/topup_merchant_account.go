package models

type TopupMerchantAccountRequest struct {
	MerchantAccountNumber int    `json:"merchantAccountNumber"`
	ConnectorID           int    `json:"connectorId"`
	TopupAmount           int    `json:"topupAmount"`
	CreatedBy             string `json:"createdBy"`
}

type TopupMerchantAccountResponse struct {
	TransactionReferenceNumber string  `json:"transactionReferenceNumber"`
	BeneficiaryAccountNumber   int     `json:"beneficiaryAccountNumber"`
	BeneficiaryName            string  `json:"beneficiaryName"`
	BeneficiaryCurrencyCode    string  `json:"beneficiaryCurrencyCode"`
	SourceAccountNumber        string  `json:"sourceAccountNumber"`
	SourceCustomerName         string  `json:"sourceCustomerName"`
	SourceReferenceNumber      string  `json:"sourceReferenceNumber"`
	SourceCurrencyCode         string  `json:"sourceCurrencyCode"`
	ConnectorName              string  `json:"connectorName"`
	AmountSent                 float64 `json:"amountSent"`
	AmountReceived             float64 `json:"amountReceived"`
	TransactionFee             float64 `json:"transactionFee"`
	TransactionDescription     string  `json:"transactionDescription"`
	TransactionDate            string  `json:"transactionDate"`
	ValueDate                  string  `json:"valueDate"`
	DateCreated                string  `json:"dateCreated"`
	TransactionCode            string  `json:"transactionCode"`
	TransactionTypeName        string  `json:"transactionTypeName"`
	TransactionStatusCode      string  `json:"transactionStatusCode"`
	TransactionStatusName      string  `json:"transactionStatusName"`
	BridgeReferenceNumber      string  `json:"bridgeReferenceNumber"`
	ResponseCode               string  `json:"responseCode"`
	ResponseMessage            string  `json:"responseMessage"`
}
