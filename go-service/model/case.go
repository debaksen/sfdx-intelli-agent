package model

type SalesforceCase struct {
	Id      string `json:"id"`
	Subject string `json:"subject"`
	Status  string `json:"status"`
	Details string `json:"details"`
}
