package integration

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

// QueryVectorDB performs semantic search for RAG (stub)
func QueryVectorDB(sfCase interface{}) string {
	// TODO: Integrate with Pinecone/Redis
	return "" // Return context string if found
}

// CallOpenAI sends prompt to OpenAI GPT API (stub)
type AIResponse struct {
	Resolved bool
	Reason   string
	Action   string
}

func CallOpenAI(prompt string) (AIResponse, float64) {
	log.Printf("[DEBUG] Sending prompt to OpenAI: %s", prompt)
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Printf("[ERROR] OPENAI_API_KEY not set")
		return AIResponse{Resolved: false, Reason: "OpenAI API key missing", Action: "Manual review required"}, 0.0
	}

	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": "You are a support case resolution and fraud detection agent."},
			{"role": "user", "content": prompt},
		},
		"max_tokens": 256,
	}
	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("[ERROR] Creating OpenAI request: %v", err)
		return AIResponse{Resolved: false, Reason: "OpenAI request error", Action: "Manual review required"}, 0.0
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[ERROR] Calling OpenAI: %v", err)
		return AIResponse{Resolved: false, Reason: "OpenAI call failed", Action: "Manual review required"}, 0.0
	}
	defer resp.Body.Close()
	var openaiResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&openaiResp); err != nil {
		log.Printf("[ERROR] Decoding OpenAI response: %v", err)
		return AIResponse{Resolved: false, Reason: "OpenAI decode error", Action: "Manual review required"}, 0.0
	}
	if len(openaiResp.Choices) == 0 {
		return AIResponse{Resolved: false, Reason: "No response from OpenAI", Action: "Manual review required"}, 0.0
	}
	content := openaiResp.Choices[0].Message.Content
	// Simple parsing: look for keywords in the response
	if containsFraudOrSpam(content) {
		return AIResponse{Resolved: true, Reason: "Detected as fraud/spam email.", Action: "Close as fraud"}, 0.99
	}
	if strings.Contains(content, "resolved") {
		return AIResponse{Resolved: true, Reason: content, Action: "Auto-close case."}, 0.95
	}
	return AIResponse{Resolved: false, Reason: content, Action: "Manual review required"}, 0.7
}

func containsFraudOrSpam(prompt string) bool {
	// Simple keyword check for demo; replace with real NLP/AI logic
	if strings.Contains(prompt, "fraud") || strings.Contains(prompt, "spam") {
		return true
	}
	return false
}

// PatchSalesforceCase auto-closes the case in Salesforce (stub)
func PatchSalesforceCase(caseId, reason string) {
	// TODO: Implement Salesforce PATCH API call
	log.Printf("[INFO] Auto-closing Salesforce case %s: %s", caseId, reason)
}
