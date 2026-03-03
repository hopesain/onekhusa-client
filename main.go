package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

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
	//Get Access Token
	token, _ := services.GetAccessToken()

	// Uncomment this to test account topup.
	// TopUp Merchant Account
	response, err := services.TopupMerchantAccount(token)
	if err != nil {
		slog.Error("Unable to topup merchant account", "error", err)
		return
	}

	prettyRensonse, _ := json.MarshalIndent(response, "", " ")
	fmt.Println(string(prettyRensonse))

	// accountNumber := services.GetMerchantAccountNumber()
	// userEmail := services.GetAdminEmail()

	// var batchRequest = models.BatchJSONUploadRequest{
	// 	Header: models.HeaderSection{
	// 		MerchantAccountNumber: accountNumber,
	// 		IsBatchScheduled:      false,
	// 		ScheduledDate:         nil,
	// 		CapturedBy:            userEmail,
	// 	},
	// 	Transactions: []models.TransactionsSection{
	// 		{
	// 			BeneficiaryName:          "CASEY CONROY",
	// 			ConnectorID:              221500,
	// 			BeneficiaryAccountNumber: "3333888800",
	// 			TransactionDescription:   "Salary Payment",
	// 			TransactionAmount:        450000,
	// 			SourceReferenceNumber:    "C8BQLWY1UUU5",
	// 		},
	// 	},
	// }

	// batchResponse, err := clients.UploadBatchJSONClient(batchRequest)
	// if err != nil {
	// 	slog.Error("failed to upload json", "error", err)
	// 	return
	// }
	// pretty, _ := json.MarshalIndent(batchResponse, "", "  ")
	// fmt.Println(string(pretty))

}
