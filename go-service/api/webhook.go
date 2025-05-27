package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"sfdx-intelli-agent/model"
	"sfdx-intelli-agent/service"
)

func SalesforceWebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] reading body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var sfCase model.SalesforceCase
	if err := json.Unmarshal(body, &sfCase); err != nil {
		log.Printf("[ERROR] unmarshalling: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := service.ProcessCase(sfCase)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
