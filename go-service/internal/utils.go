package internal

import (
	"log"
	"os"
)

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func LogCompliance(event, details string) {
	log.Printf("[COMPLIANCE] %s: %s", event, details)
}
