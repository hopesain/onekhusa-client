package uploads

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/hopesain/onekhusa-client/pkg/utils"
)

func UploadBatchJSONService(
	accessToken string,
	requestData BatchJSONUploadRequest,
) (responseData BatchJSONUploadResponse, err error) {

	payload, err := json.Marshal(requestData)
	if err != nil {
		slog.Error(
			"Failed to marshal the request data",
			"Context", "Upload Batch JSON Service",
			"Error", err,
		)
		return BatchJSONUploadResponse{}, err
	}

	url := "https://api.onekhusa.com/sandbox/v1/disbursements/batch/addJson"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)

	idempotencyKey := utils.GenerateIdempotencyKey()
	request.Header.Set("X-Idempotency-Key", idempotencyKey)

	client := http.Client{Timeout: 20 * time.Second}

	response, err := client.Do(request)
	if err != nil {
		slog.Error(
			"Failed to send an HTTP request",
			"Context", "Upload Batch JSON",
			"Error", err,
		)
		return BatchJSONUploadResponse{}, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		slog.Error(
			"Failed to decode the response",
			"Context", "Upload Batch JSON Service",
			"Error", err,
		)
		return BatchJSONUploadResponse{}, err
	}

	return responseData, nil
}
