package clients

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/hopesain/onekhusa-client/models"
	"github.com/hopesain/onekhusa-client/services"
)

func UploadBatchJSONClient(
	requestData models.BatchJSONUploadRequest,
) (responseData *models.BatchJSONUploadResponse, err error) {
	payload, err := json.Marshal(requestData)
	if err != nil {
		slog.Error("Error marshalling request data", "error", err)
		return nil, err
	}

	slog.Info("Batch payload being sent", "payload", string(payload))

	url := "https://api.onekhusa.com/sandbox/v1/disbursements/batch/addJson"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))

	request.Header.Set("Content-Type", "application/json")

	token, err := services.GetAccessToken()
	if err != nil {
		slog.Error("Failed to retrieve access token", "error", err)
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+token)

	idempotencyKey := services.GenerateIdempotencyKey()
	request.Header.Set("X-Idempotency-Key", idempotencyKey)

	client := http.Client{Timeout: 20 * time.Second}

	response, err := client.Do(request)
	if err != nil {
		slog.Error("Request failed", "error", err)
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK || response.StatusCode != http.StatusAccepted {
		slog.Warn("response code", "statusCode", response.StatusCode)
	}

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		slog.Error("Failed to decode the response", "error", err)
		return nil, err
	}

	return responseData, nil
}
