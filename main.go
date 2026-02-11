package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/hopesain/onekhusa-client/clients"
	"github.com/hopesain/onekhusa-client/models"
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
	token, _ := services.GetAccessToken()
	fmt.Println(token)

	resp, err := services.TopupMerchantAccount(token)
	if err != nil {
		log.Fatalf("Topup failed: %v", err)
	}

	fmt.Printf("Topup Response:\n%+v\n", resp)

	accountNumber := services.GetMerchantAccountNumber()
	userEmail := services.GetAdminEmail()

	var batchRequest = models.BatchJSONUploadRequest{
		Header: models.HeaderSection{
			MerchantAccountNumber: accountNumber,
			IsBatchScheduled:      false,
			ScheduledDate:         "2025-09-19",
			CapturedBy:            userEmail,
		},
		Transactions: []models.TransactionsSection{
			{
				BeneficiaryName:          "John Phiri",
				ConnectorID:              221500,
				BeneficiaryAccountNumber: "12345678",
				TransactionDescription:   "Salary Payment",
				TransactionAmount:        45000,
				SourceReferenceNumber:    "QKAYXD208367",
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
