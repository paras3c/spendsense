package services

import (
	"strings"

	"github.com/siddhartharajbongshi/spendsense-backend/models"
)

type CategorizerService struct {
	keywordMap map[string][]string
}

func NewCategorizerService() *CategorizerService {
	return &CategorizerService{
		keywordMap: map[string][]string{
			"Food": {
				"zomato", "swiggy", "dunzo", "burger king", "domino",
				"pizza hut", "mcd", "starbucks", "coffee", "restaurant",
				"food", "meal", "eat", "dine",
			},
			"Transport": {
				"uber", "ola", "auto", "autorickshaw", "taxi", "cab",
				"petrol", "fuel", "parking", "metro", "bus", "train",
				"travel", "ride",
			},
			"Subscriptions": {
				"netflix", "amazon prime", "spotify", "youtube", "hulu",
				"disney", "hotstar", "subscription", "premium", "plan",
				"membership",
			},
			"Shopping": {
				"amazon", "flipkart", "dmart", "supermarket", "mall",
				"clothing", "dress", "shoe", "shopping", "purchase",
				"retail", "store",
			},
			"Rent": {
				"rent", "housing", "landlord", "deposit", "lease",
			},
			"Utilities": {
				"electricity", "water", "internet", "wifi", "broadband",
				"mobile recharge", "phone", "bill",
			},
		},
	}
}

func (c *CategorizerService) Categorize(description string) string {
	descLower := strings.ToLower(description)

	for category, keywords := range c.keywordMap {
		for _, keyword := range keywords {
			if strings.Contains(descLower, keyword) {
				return category
			}
		}
	}
	return "Misc"
}

func (c *CategorizerService) CategorizeExpenses(expenses []models.Expense) []models.Expense {
	for i := range expenses {
		expenses[i].Category = c.Categorize(expenses[i].Description)
	}
	return expenses
}
