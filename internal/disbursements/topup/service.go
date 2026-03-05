package topup

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

// payload := TopupMerchantAccountRequest{
// 		MerchantAccountNumber: merchantAccountNumber,
// 		ConnectorID:           221500, //FDH Bank
// 		TopupAmount:           100000000,
// 		CreatedBy:             adminEmail,
// 	}

func TopupMerchantAccount(accessToken string, payloadInput TopupMerchantAccountRequest) (TopupMerchantAccountResponse, error) {
	payload := payloadInput

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		slog.Error(
			"Failed to marshal a payload",
			"Context", "Topup Merchant Account",
			"Error", err,
		)
		return TopupMerchantAccountResponse{}, err
	}

	url := "https://api.onekhusa.com/sandbox/v1/merchants/accounts/topup"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		slog.Error(
			"Failed to create a new request",
			"Context", "Topup Merchant Account",
			"Error", err,
		)
		return TopupMerchantAccountResponse{}, err
	}

	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 20 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		slog.Error(
			"Failed to send an HTTP request",
			"Context", "Topup Merchant Account",
			"Error", err,
		)
		return TopupMerchantAccountResponse{}, err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Error(
			"Failed to read response body",
			"Context", "Topup Merchant Account",
			"Error", err,
		)
		return TopupMerchantAccountResponse{}, err
	}

	var topupResponse TopupMerchantAccountResponse

	if err := json.Unmarshal(bodyBytes, &topupResponse); err != nil {
		slog.Error(
			"Failed to umarshal the response",
			"Context", "Topup Merchant Account",
			"Error", err,
		)
		return TopupMerchantAccountResponse{}, err
	}

	return topupResponse, nil
}
