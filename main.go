package main

import (
	"fmt"
	"log"
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
	fmt.Println("Hello world.")
	fmt.Println(services.GetMerchantAccountNumber())
	token, _ := services.GetAccessToken()
	fmt.Println(token)
	resp, err := services.TopupMerchantAccount(token)
	if err != nil {
		log.Fatalf("Topup failed: %v", err)
	}

	fmt.Printf("Topup Response:\n%+v\n", resp)
}
