package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/siddhartharajbongshi/spendsense-backend/models"
	"github.com/siddhartharajbongshi/spendsense-backend/services"
)

var (
	parser      = services.NewParserService()
	categorizer = services.NewCategorizerService()
	insightGen  = services.NewInsightService()
	// Default to "llama3", user can change or we can make it an env var
	tutor = services.NewLLMService("tinyllama")

	// In-memory storage for demo
	userExpenses = make(map[string][]models.Expense)
)

func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // SvelteKit default port
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", enableCors(handleHealth))
	mux.HandleFunc("/upload", enableCors(handleUpload))
	mux.HandleFunc("/sample-data", enableCors(handleSampleData))
	mux.HandleFunc("/dashboard", enableCors(handleDashboard))
	mux.HandleFunc("/explain-insight", enableCors(handleExplainInsight))
	mux.HandleFunc("/generate-persona", enableCors(handleGeneratePersona))

	port := "8000"
	fmt.Printf("Backend running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// file, handler, err := r.FormFile("file") // Does not work with ParseMultipartForm not called?
	// ParseMultipartForm is called by FormFile automatically usually, but let's be safe
	r.ParseMultipartForm(10 << 20) // 10 MB

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save to temp file to reuse our parser logic
	tempFile, err := os.CreateTemp("", "upload-*.csv")
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	tempFile.Write(fileBytes)
	tempFile.Close()

	// Parse
	expenses, err := parser.ParseCSV(tempFile.Name())
	if err != nil {
		http.Error(w, fmt.Sprintf("Parsing failed: %v", err), http.StatusBadRequest)
		return
	}

	// Categorize
	expenses = categorizer.CategorizeExpenses(expenses)

	// Store
	userID := "default"
	userExpenses[userID] = expenses

	// Generate Response
	dashboard := insightGen.GenerateDashboardData(expenses)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboard)
}

func handleSampleData(w http.ResponseWriter, r *http.Request) {
	expenses := parser.GenerateSampleData()
	expenses = categorizer.CategorizeExpenses(expenses)

	userID := "default"
	userExpenses[userID] = expenses

	dashboard := insightGen.GenerateDashboardData(expenses)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboard)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	userID := "default"
	expenses, exists := userExpenses[userID]
	if !exists {
		http.Error(w, "No data uploaded", http.StatusNotFound)
		return
	}

	dashboard := insightGen.GenerateDashboardData(expenses)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboard)
}

type ExplainRequest struct {
	Insight  models.Insight `json:"insight"`
	FollowUp string         `json:"follow_up"`
	Style    string         `json:"style"`  // "polite" or "savage"
	Action   string         `json:"action"` // "explain" or "draft_cancel"
}

func handleExplainInsight(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ExplainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	explanation, err := tutor.GetExplanation(req.Insight, req.FollowUp, req.Style, req.Action)
	if err != nil {
		http.Error(w, fmt.Sprintf("AI Tutor error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"explanation": explanation})
}

func handleGeneratePersona(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := "default"
	expenses, exists := userExpenses[userID]
	if !exists {
		http.Error(w, "No data uploaded", http.StatusNotFound)
		return
	}

	persona, err := tutor.GeneratePersona(expenses, "savage")
	if err != nil {
		http.Error(w, fmt.Sprintf("Persona error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persona)
}
