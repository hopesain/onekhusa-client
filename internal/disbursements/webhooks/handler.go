package webhooks

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func BatchFailedWebhook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			slog.Error("Method not allowed", "method", r.Method)
			return
		}

		var payload BatchFailedWebhookPayload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			slog.Error("Failed to decode webhook payload", "error", err)
			return
		}

		defer r.Body.Close()

		slog.Info("Batch Failed Webhook received",
			"merchantAccountNumber", payload.MerchantAccountNumber,
			"batchNumber", payload.MetaData.BatchNumber,
			"totalAmount", payload.MetaData.TotalAmount,
			"numberOfTransactions", payload.MetaData.NumberOfTransactions,
			"errorMessage", payload.ErrorMessage,
		)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Webhook received successfully",
		})
	}
}
