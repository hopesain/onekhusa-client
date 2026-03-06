package webhooks

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

func AddWebhook(accessToken string, requestData AddWebhookRequest) (AddWebhookResponse, error) {
	payload, err := json.Marshal(requestData)
	if err != nil {
		slog.Error(
			"Failed to marshal the requested input",
			"Context", "Add Webhook",
			"Error", err,
		)
		return AddWebhookResponse{}, err
	}
	url := "https://api.onekhusa.com/sandbox/v1/merchants/webhooks/add"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		slog.Error(
			"Failed to create a new request",
			"Context", "Add Webook",
			"Error", err,
		)
		return AddWebhookResponse{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)

	client := http.Client{Timeout: 20 * time.Second}

	response, err := client.Do(request)
	if err != nil {
		slog.Error(
			"Failed to send an HTTP request",
			"Context", "Add Webhook",
			"Error", err,
		)
		return AddWebhookResponse{}, err
	}

	defer response.Body.Close()

	var message string

	err = json.NewDecoder(response.Body).Decode(&message)
	if err != nil {
		slog.Error(
			"Failed to decode the response", 
			"Context", "Add Webook",
			"Error", err,
		)
		return AddWebhookResponse{}, err
	}

	output := AddWebhookResponse{
		message: message,
	}
	
	return output, nil
}
