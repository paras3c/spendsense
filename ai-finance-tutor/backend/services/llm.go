package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/siddhartharajbongshi/spendsense-backend/models"
)

type LLMService struct {
	BaseURL string
	Model   string
	Client  *http.Client
}

func NewLLMService(model string) *LLMService {
	// Default to local Ollama port
	return &LLMService{
		BaseURL: "http://localhost:11434/api/chat",
		Model:   model, // e.g. "llama3" or "mistral"
		Client:  &http.Client{Timeout: 30 * time.Second},
	}
}

type OllamaRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Format   string    `json:"format,omitempty"` // "json" or empty
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaResponse struct {
	Model     string  `json:"model"`
	CreatedAt string  `json:"created_at"`
	Message   Message `json:"message"`
	Done      bool    `json:"done"`
}

const SystemPrompt = `You are a financial literacy tutor for students and young professionals in India.

RULES:
1. Explain insights clearly using simple language
2. NEVER invent numbers - only explain given data
3. NEVER predict future spending or income
4. NEVER give risky financial advice (no crypto, no stocks tips, no borrowing advice)
5. Always explain WHY the insight matters
6. Give ONE safe, practical improvement tip
7. Use Indian context (₹, UPI, local services)
8. Be friendly and encouraging
9. End with one actionable next step (e.g. "Try setting a limit for next week")

SAFE TIPS ONLY:
- Track spending
- Set category limits
- Cancel unused subscriptions
- Cook at home more
- Use public transport`

func (s *LLMService) GetExplanation(insight models.Insight, followUp string, style string, action string) (string, error) {
	var systemPrompt string
	var userPrompt string

	if action == "draft_cancel" {
		systemPrompt = `You are a professional assistant.
TASK: Draft a polite but firm cancellation email for the subscription service mentioned below. Use placeholders like [Name] where necessary.
Keep it brief and formal.`
		userPrompt = fmt.Sprintf(`CONTEXT:
Service: %s
Cost: %.2f`, insight.Message, insight.MonthlyCost)
	} else {
		// Normal or Savage Explanation
		systemPrompt = "You are a helpful and encouraging financial tutor. Start by acknowledging the user's situation, then give one specific, actionable tip to save money. Be concise (max 2 sentences)."

		if style == "savage" {
			systemPrompt = "You are a ruthless, stand-up comedian financial roaster. Your goal is to roast the user for their spending habits using Gen-Z slang (no cap, fr, based). Be mean, funny, and unforgettable. Do NOT be polite. Make them regret spending money."
		}

		userPrompt = fmt.Sprintf(`DATA:
Insight: %s
Amount: %.2f
Breakdown: %v

USER QUESTION: %s`, insight.Message, insight.MonthlyCost, insight.Breakdown, followUp)
	}

	reqBody := OllamaRequest{
		Model: s.Model,
		Messages: []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := s.Client.Post(s.BaseURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to call Ollama: %v. Is it running?", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ollama API error: %s", resp.Status)
	}

	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", fmt.Errorf("failed to parse Ollama response: %v", err)
	}

	return ollamaResp.Message.Content, nil
}

func (s *LLMService) GeneratePersona(expenses []models.Expense, style string) (models.PersonaResponse, error) {
	// Calculate total and top category for context
	total := 0.0
	catMap := make(map[string]float64)
	for _, e := range expenses {
		total += e.Amount
		catMap[e.Category] += e.Amount
	}
	// Simplified context for prompt
	context := fmt.Sprintf("Total Spent: %.2f. Breakown: %v", total, catMap)

	prompt := fmt.Sprintf(`SYSTEM: You are a creative writer for a 'Spotify Wrapped' style finance app.
TASK: Analyze the spending summary below and assign the user a hilarious 'Financial Archetype'.
valid archetypes examples: "The Subscription Hoarder", "The Caffeine Investor", "The Midnight Snacker", "The Impulse Buyer".

OUTPUT JSON ONLY:
{
	"archetype": "Title",
	"emoji": "emoji",
	"description": "Short, funny description of why they got this.",
	"savage_quote": "A short, savage roast."
}

DATA: %s`, context)

	reqBody := OllamaRequest{
		Model: s.Model,
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
		Format: "json", // Request JSON format from Ollama
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return models.PersonaResponse{}, err
	}

	resp, err := s.Client.Post(s.BaseURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return models.PersonaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.PersonaResponse{}, fmt.Errorf("ollama error: %s", resp.Status)
	}

	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return models.PersonaResponse{}, err
	}

	// Parse the JSON string inside the message content
	var persona models.PersonaResponse
	if err := json.Unmarshal([]byte(ollamaResp.Message.Content), &persona); err != nil {
		// Fallback if LLM didn't return perfect JSON (common with small models)
		// Return a generic error persona or try to clean up
		fmt.Printf("Failed to parse JSON content: %v\nRaw: %s\n", err, ollamaResp.Message.Content)
		return models.PersonaResponse{
			Archetype:   "The Mystery Spender",
			Emoji:       "👻",
			Description: "Your spending was too chaotic even for AI.",
			SavageQuote: "I have no words.",
		}, nil
	}

	return persona, nil
}
