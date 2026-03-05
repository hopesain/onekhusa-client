package authentication

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func GetAccessToken(input AccessTokenRequest) (AccessTokenResponse, error) {
	url := "https://api.onekhusa.com/sandbox/v1/account/getAccessToken"

	payload := input

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		slog.Error(
			"Failed to marshal payload",
			"Context", "Get Access Token from oneKhusa",
			"Error", err,
		)
		return AccessTokenResponse{}, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		slog.Error(
			"Failed to create a new request",
			"Context", "Get Access Token from oneKhusa",
			"Error", err,
		)
		return AccessTokenResponse{}, err
	}

	request.Header.Set("Accept-Language", "en")
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		slog.Error(
			"Failed to send an HTTP request",
			"Context", "Get access token from oneKhusa",
			"Error", err,
		)
		return AccessTokenResponse{}, err
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var output AccessTokenResponse

	err = json.Unmarshal(bodyBytes, &output)
	if err != nil {
		slog.Error(
			"Failed to unmarshal the output",
			"Context", "Get access token from oneKhusa",
			"Error", err,
		)
		return AccessTokenResponse{}, err
	}

	return output, nil
}
