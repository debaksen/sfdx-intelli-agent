# sfdx-intelli-agent Go Microservice

This project is a Go-based middleware that integrates Salesforce Service Cloud with OpenAI's GPT APIs, optionally enhanced with Retrieval-Augmented Generation (RAG) via a vector database.

## Features
- Receives Salesforce Case data via webhooks or REST polling
- Formats and sends case data to OpenAI GPT for resolution analysis
- Auto-closes resolved cases in Salesforce via REST PATCH
- Optional RAG: semantic search over internal knowledge base (Pinecone/Redis)
- Structured prompts and response validation
- Logging and compliance safeguards

## Getting Started
1. `go mod tidy` to install dependencies
2. Implement your Salesforce, OpenAI, and vector DB credentials in config
3. Run the service: `go run main.go`

## Project Structure
- `main.go` — Entry point, HTTP server setup
- `api/` — REST endpoints for Salesforce integration
- `service/` — Business logic, AI orchestration
- `integration/` — Salesforce, OpenAI, and vector DB clients
- `internal/` — Utilities, logging, config, validation

## Extending
- Add new RAG providers in `integration/vector/`
- Update prompt templates in `service/prompt.go`

---
This is a starter template. Extend as needed for your use case.
