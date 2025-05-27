package service

import (
	"log"
	"sfdx-intelli-agent/integration"
	"sfdx-intelli-agent/model"
)

type CaseResult struct {
	Resolved bool   `json:"resolved"`
	Reason   string `json:"reason"`
	Action   string `json:"action"`
}

func ProcessCase(sfCase model.SalesforceCase) CaseResult {
	prompt := BuildPrompt(sfCase)
	// Optionally augment with RAG
	contextDocs := integration.QueryVectorDB(sfCase)
	if contextDocs != "" {
		prompt += "\nContext:\n" + contextDocs
	}
	aiResp, confidence := integration.CallOpenAI(prompt)
	if confidence < 0.8 {
		log.Printf("[WARN] Low confidence: %.2f", confidence)
		return CaseResult{Resolved: false, Reason: "Low AI confidence", Action: "Manual review required"}
	}
	// Fraud detection: If AI flags as fraud, close the case
	if aiResp.Action == "Close as fraud" {
		integration.PatchSalesforceCase(sfCase.Id, aiResp.Reason)
		return CaseResult{Resolved: true, Reason: aiResp.Reason, Action: aiResp.Action}
	}
	if aiResp.Resolved {
		integration.PatchSalesforceCase(sfCase.Id, aiResp.Reason)
	}
	// Map AIResponse to CaseResult
	return CaseResult{
		Resolved: aiResp.Resolved,
		Reason:   aiResp.Reason,
		Action:   aiResp.Action,
	}
}

func BuildPrompt(sfCase model.SalesforceCase) string {
	prompt := "Case Subject: " + sfCase.Subject + "\nDetails: " + sfCase.Details + "\nStatus: " + sfCase.Status + "\n"
	prompt += "If this is a fraud or spam email, or the content does not make sense, respond with Action: 'Close as fraud' and provide a reason. Otherwise, determine if this case is resolved and provide a reason."
	return prompt
}
