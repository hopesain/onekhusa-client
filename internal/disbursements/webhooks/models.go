package webhooks

type AddWebhookRequest struct {
	MerchantAccountNumber int
	EventCode             string
	CallbackURL           string
	CapturedBy            string
}

type AddWebhookResponse struct {
	message string
}
