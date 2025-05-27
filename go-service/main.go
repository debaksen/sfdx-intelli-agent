package main

import (
	"log"
	"net/http"
	"sfdx-intelli-agent/api"
)

func main() {
	http.HandleFunc("/webhook", api.SalesforceWebhookHandler)
	log.Println("[INFO] Service started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("[FATAL] %v", err)
	}
}
