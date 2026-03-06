package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/hopesain/onekhusa-client/internal/authentication"
	"github.com/hopesain/onekhusa-client/internal/disbursements/topup"
	"github.com/hopesain/onekhusa-client/internal/disbursements/uploads"
	"github.com/hopesain/onekhusa-client/internal/disbursements/webhooks"
	"github.com/hopesain/onekhusa-client/pkg/utils"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Unable to load the env variables", "error", err)
		os.Exit(1)
	}
}

func main() {
	// Get Access Token
	apiKey := utils.GetOnekhusaApiKey()
	secretKey := utils.GetOnekhusaSecretKey()
	organizationID := utils.GetOrganizationID()
	merchantAccountNumber, err := utils.GetMerchantAccountNumber()
	if err != nil {
		slog.Error(
			"Failed to retrieve merchant account number",
			"Error", err,
		)
		return
	}

	getAccessTokenInput := authentication.AccessTokenRequest{
		APIKey:                apiKey,
		APISecret:             secretKey,
		OrganizationID:        organizationID,
		MerchantAccountNumber: merchantAccountNumber,
	}

	tokenOutput, err := authentication.GetAccessToken(getAccessTokenInput)
	if err != nil {
		slog.Error(
			"Failed to retrieve the access token",
			"Error", err,
		)
		return
	}

	prettyAccessTokenData, err := json.MarshalIndent(tokenOutput, "", " ")
	if err != nil {
		slog.Error(
			"Failed to marshal Indent",
			"Error", err,
		)
		return
	}
	fmt.Println("Access Token Details")
	fmt.Println(string(prettyAccessTokenData))

	accessToken := tokenOutput.AccessToken

	// Topup Merchant Account
	adminEmail := utils.GetAdminEmail()
	topupAccountInput := topup.TopupMerchantAccountRequest{
		MerchantAccountNumber: merchantAccountNumber,
		ConnectorID:           221500, //FDH Bank
		TopupAmount:           10000000,
		CreatedBy:             adminEmail,
	}

	accountTopupResponse, err := topup.TopupMerchantAccount(accessToken, topupAccountInput)

	prettyAccountTopupResponse, err := json.MarshalIndent(accountTopupResponse, "", " ")
	if err != nil {
		slog.Error(
			"Failed to marshal Indent accountTopupResponse",
			"Error", err,
		)
		return
	}

	fmt.Println("Topup Merchant Account Details")
	fmt.Println(string(prettyAccountTopupResponse))

	//Upload Batch JSON
	accountNumber := merchantAccountNumber
	userEmail := adminEmail

	var batchRequestInputData = uploads.BatchJSONUploadRequest{
		Header: uploads.HeaderSection{
			MerchantAccountNumber: accountNumber,
			IsBatchScheduled:      false,
			ScheduledDate:         nil,
			CapturedBy:            userEmail,
		},
		Transactions: []uploads.TransactionsSection{
			{
				BeneficiaryName:          "CASEY CONROY",
				ConnectorID:              221500,
				BeneficiaryAccountNumber: "3333888800",
				TransactionDescription:   "Salary Payment",
				TransactionAmount:        450000,
				SourceReferenceNumber:    "C8BQLWY1UUU5",
			},
		},
	}

	uploadBatchJSONResponse, err := uploads.UploadBatchJSONService(accessToken, batchRequestInputData)
	if err != nil {
		slog.Error(
			"Failed to upload batch JSON",
			"Error", err,
		)
		return
	}

	pretty, err := json.MarshalIndent(uploadBatchJSONResponse, "", "  ")
	if err != nil {
		slog.Error(
			"Failed to marshal Indent uploadBatchJSONResponse",
			"Error", err,
		)
		return
	}
	fmt.Println("Upload Batch JSON Service")
	fmt.Println(string(pretty))

	addWebhookInput := webhooks.AddWebhookRequest{
		MerchantAccountNumber: accountNumber,
		EventCode:             "batch.payout.success",
		CallbackURL:           "http://localhost:8080/webhooks/batch-proceed",
		CapturedBy:            userEmail,
	}

	addWebhookOutput, err := webhooks.AddWebhook(accessToken, addWebhookInput)
	prettyAddWebhook, err := json.MarshalIndent(addWebhookOutput, "", "  ")
	if err != nil {
		slog.Error(
			"Failed to marshal Indent uploadBatchJSONResponse",
			"Error", err,
		)
		return
	}
	fmt.Println(string(prettyAddWebhook))

	//Batch Failed Webhook
	http.HandleFunc("/webhooks/batch-failed", webhooks.BatchFailedWebhook())

	slog.Info("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}
