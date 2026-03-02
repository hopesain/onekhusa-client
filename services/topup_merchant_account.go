package services

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/hopesain/onekhusa-client/models"
)

func TopupMerchantAccount(accessToken string) (models.TopupMerchantAccountResponse, error) {
	merchantAccountNumber := GetMerchantAccountNumber()
	adminEmail := GetAdminEmail()

	payload := models.TopupMerchantAccountRequest{
		MerchantAccountNumber: merchantAccountNumber,
		ConnectorID:           221500, //FDH Bank
		TopupAmount:           10000000,
		CreatedBy:             adminEmail,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		slog.Error("Failed to marshal payload", "error", err)
		return models.TopupMerchantAccountResponse{}, err
	}

	url := "https://api.onekhusa.com/sandbox/v1/merchants/accounts/topup"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		slog.Error("Failed to create request", "error", err)
		return models.TopupMerchantAccountResponse{}, err
	}

	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 20 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		slog.Error("Request failed", "error", err)
		return models.TopupMerchantAccountResponse{}, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error("Failed to read response body", "error", err)
		return models.TopupMerchantAccountResponse{}, err
	}

	var topupResponse models.TopupMerchantAccountResponse

	if err := json.Unmarshal(bodyBytes, &topupResponse); err != nil {
		return models.TopupMerchantAccountResponse{}, err
	}

	return topupResponse, nil
}
