package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/siddhartharajbongshi/spendsense-backend/models"
)

type ParserService struct{}

func NewParserService() *ParserService {
	return &ParserService{}
}

func (p *ParserService) ParseCSV(filePath string) ([]models.Expense, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	// Validate columns
	required := map[string]bool{"date": false, "amount": false, "description": false}
	colMap := make(map[string]int)

	for i, col := range header {
		lowerCol := strings.ToLower(strings.TrimSpace(col))
		if _, exists := required[lowerCol]; exists {
			required[lowerCol] = true
			colMap[lowerCol] = i
		}
	}

	for col, found := range required {
		if !found {
			return nil, fmt.Errorf("missing required column: %s", col)
		}
	}

	var expenses []models.Expense

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		amountStr := record[colMap["amount"]]
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			continue // Skip invalid amounts
		}
		if amount <= 0 {
			continue // Skip non-expenses or zero values
		}

		dateStr := record[colMap["date"]]
		// Try parsing date to standardize it (optional, but good for validation)
		parsedDate, err := parseDate(dateStr)
		if err == nil {
			dateStr = parsedDate
		} else {
			// If duplicate formats or tricky ones, maybe try more parsers.
			// For now, keep original or error? Let's keep original if parse fails but log/print?
			// Actually, let's enforce YYYY-MM-DD for consistency if possible, or just pass through.
			// The user sample is YYYY-MM-DD.
		}

		expenses = append(expenses, models.Expense{
			Date:        dateStr,
			Description: record[colMap["description"]],
			Amount:      amount,
		})
	}

	// Sort by date
	sort.Slice(expenses, func(i, j int) bool {
		return expenses[i].Date < expenses[j].Date
	})

	return expenses, nil
}

func parseDate(dateStr string) (string, error) {
	formats := []string{"2006-01-02", "02-01-2006", "1/2/2006", "2006/01/02"}
	for _, format := range formats {
		t, err := time.Parse(format, dateStr)
		if err == nil {
			return t.Format("2006-01-02"), nil
		}
	}
	return "", fmt.Errorf("unknown date format")
}

func (p *ParserService) GenerateSampleData() []models.Expense {
	return []models.Expense{
		{Date: "2026-01-01", Description: "Zomato Order", Amount: 450},
		{Date: "2026-01-02", Description: "Netflix Subscription", Amount: 199},
		{Date: "2026-01-03", Description: "Autorickshaw", Amount: 50},
		{Date: "2026-01-04", Description: "Grocery Shopping", Amount: 1200},
		{Date: "2026-01-05", Description: "Amazon Prime", Amount: 179},
		{Date: "2026-01-06", Description: "Swiggy Delivery", Amount: 380},
		{Date: "2026-01-07", Description: "Spotify Premium", Amount: 119},
		{Date: "2026-01-08", Description: "Burger King", Amount: 320},
		{Date: "2026-01-09", Description: "Uber", Amount: 280},
		{Date: "2026-01-10", Description: "Mobile Recharge", Amount: 399},
		{Date: "2026-01-11", Description: "DMart Shopping", Amount: 850},
		{Date: "2026-01-12", Description: "Netflix Payment", Amount: 199},
		{Date: "2026-01-13", Description: "Coffee - Starbucks", Amount: 250},
		{Date: "2026-01-14", Description: "Swiggy Order", Amount: 520},
		{Date: "2026-01-15", Description: "Book Purchase", Amount: 450},
	}
}
