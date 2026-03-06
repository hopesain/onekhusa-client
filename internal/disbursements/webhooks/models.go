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
