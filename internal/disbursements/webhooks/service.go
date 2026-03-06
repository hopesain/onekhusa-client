package webhooks

import (
	"bytes"
	"encoding/json"
	"io"
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
			"Context", "Add Webhook",
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return AddWebhookResponse{}, err
	}

	slog.Info("Webhook API response", "status", response.StatusCode, "body", string(body))

	var output AddWebhookResponse

	if len(body) > 0 {
		err = json.Unmarshal(body, &output)
		if err != nil {
			slog.Error(
				"Failed to decode the response",
				"Context", "Add Webhook",
				"Error", err,
			)
			return AddWebhookResponse{}, err
		}
	}

	return output, nil
}
