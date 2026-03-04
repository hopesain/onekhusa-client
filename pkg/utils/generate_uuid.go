package utils

import "github.com/google/uuid"

func GenerateIdempotencyKey() string {
	return uuid.NewString()
}
