package webhooks

type AddWebhookRequest struct {
	MerchantAccountNumber int    `json:"merchantAccountNumber"`
	EventCode             string `json:"eventCode"`
	CallbackURL           string `json:"callbackURL"`
	CapturedBy            string `json:"capturedBy"`
}

type AddWebhookResponse struct {
	Message string `json:"message"`
}

type BatchFailedWebhookPayload struct {
	MerchantAccountNumber int      `json:"merchantAccountNumber"`
	IsBatchScheduled      bool     `json:"isBatchScheduled"`
	ScheduledDate         *string  `json:"scheduledDate"`
	UploadType            string   `json:"uploadType"`
	CapturerEmailAddress  string   `json:"capturerEmailAddress"`
	DateCaptured          string   `json:"dateCaptured"`
	ErrorMessage          string   `json:"errorMessage"`
	MetaData              MetaData `json:"metaData"`
}

type MetaData struct {
	BatchNumber          string `json:"batchNumber"`
	TotalAmount          string `json:"totalAmount"`
	NumberOfTransactions string `json:"numberOfTransactions"`
	DateOccurred         string `json:"dateOccurred"`
}
