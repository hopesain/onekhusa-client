package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/hopesain/onekhusa-client/clients"
	"github.com/hopesain/onekhusa-client/internal/authentication"
	"github.com/hopesain/onekhusa-client/models"
	"github.com/hopesain/onekhusa-client/pkg/utils"
	"github.com/hopesain/onekhusa-client/services"
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
		APIKey: apiKey,
		APISecret: secretKey,
		OrganizationID: organizationID,
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
	fmt.Println(string(prettyAccessTokenData))


	accessToken := tokenOutput.AccessToken
	fmt.Println(accessToken)


	//Get Access Token
	// token, _ := services.GetAccessToken()

	// // Uncomment this to test account topup.
	// // TopUp Merchant Account
	// response, err := services.TopupMerchantAccount(token)
	// if err != nil {
	// 	slog.Error("Unable to topup merchant account", "error", err)
	// 	return
	// }

	// prettyRensonse, _ := json.MarshalIndent(response, "", " ")
	// fmt.Println(string(prettyRensonse))

	accountNumber := services.GetMerchantAccountNumber()
	userEmail := services.GetAdminEmail()

	var batchRequest = models.BatchJSONUploadRequest{
		Header: models.HeaderSection{
			MerchantAccountNumber: accountNumber,
			IsBatchScheduled:      false,
			ScheduledDate:         nil,
			CapturedBy:            userEmail,
		},
		Transactions: []models.TransactionsSection{
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

	batchResponse, err := clients.UploadBatchJSONClient(batchRequest)
	if err != nil {
		slog.Error("failed to upload json", "error", err)
		return
	}
	pretty, _ := json.MarshalIndent(batchResponse, "", "  ")
	fmt.Println(string(pretty))

}
