package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hopesain/onekhusa-client/models"
)

func GetAccessToken() (accessToken string, err error) {
	apiKey := GetOnekhusaApiKey()
	secretKey := GetOnekhusaSecretKey()
	organizationID := GetOrganizationID()
	merchantAccountNumber := GetMerchantAccountNumber()

	url := "https://api.onekhusa.com/sandbox/v1/account/getAccessToken"

	payload := models.AccessTokenRequest{
		APIKey:                apiKey,
		APISecret:             secretKey,
		OrganizationID:        organizationID,
		MerchantAccountNumber: merchantAccountNumber,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		slog.Error("failed to marshal payload", "error", err)
		return "", err
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		slog.Error("Failed to send a request", "error", err)
		return "", err
	}

	request.Header.Set("Accept-Language", "en")
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 30 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		slog.Error("HTTP request failed", "error", err)
		return "", err
	}

	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var tokenResponse models.AccessTokenResponse

	err = json.Unmarshal(bodyBytes, &tokenResponse)
	if err != nil {
		slog.Error("Failed decoding", "error", err)
		return "", err
	}

	return tokenResponse.AccessToken, nil
}

func GenerateIdempotencyKey() string {
	return uuid.NewString()
}
