package models

type BatchJSONUploadRequest struct {
	Header       HeaderSection         `json:"header"`
	Transactions []TransactionsSection `json:"transactions"`
}

type HeaderSection struct {
	MerchantAccountNumber int     `json:"merchantAccountNumber"`
	IsBatchScheduled      bool    `json:"isBatchScheduled"`
	ScheduledDate         string `json:"scheduledDate"`
	CapturedBy            string  `json:"capturedBy"`
}

type TransactionsSection struct {
	BeneficiaryName          string `json:"beneficiaryName"`
	ConnectorID              int    `json:"connectorId"`
	BeneficiaryAccountNumber string `json:"beneficiaryAccountNumber"`
	TransactionDescription   string `json:"transactionDescription"`
	TransactionAmount        int    `json:"transactionAmount"`
	SourceReferenceNumber    string `json:"sourceReferenceNumber"`
}

type BatchJSONUploadResponse struct {
	MerchantAccountNumber int64 `json:"merchantAccountNumber"`
	BatchNumber           int64 `json:"batchNumber"`
}
