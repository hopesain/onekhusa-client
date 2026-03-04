package utils

import (
	"log/slog"
	"os"
	"strconv"
)

func GetOnekhusaApiKey() string {
	apiKey := os.Getenv("ONEKHUSA_API_KEY")
	return apiKey
}

func GetOnekhusaSecretKey() string {
	secretKey := os.Getenv("ONEKHUSA_SECRET_KEY")
	return secretKey
}

func GetOrganizationID() string {
	organizationID := os.Getenv("ORGANIZATION_ID")
	return organizationID
}

func GetMerchantAccountNumber() (int, error) {
	stringAccountNumber := os.Getenv("MERCHANT_ACCOUNT_NUMBER")

	merchantAccountNumber, err := strconv.Atoi(stringAccountNumber)
	if err != nil {
		slog.Error("Failed to convert merchant account number to string", "error", err)
		return 0, err
	}
	return merchantAccountNumber, nil
}

func GetAdminEmail() string {
	adminEmail := os.Getenv("ADMIN_EMAIL")
	return adminEmail
}
