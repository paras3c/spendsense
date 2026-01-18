package services

import (
	"fmt"
	"math"

	"github.com/siddhartharajbongshi/spendsense-backend/models"
)

type InsightService struct{}

func NewInsightService() *InsightService {
	return &InsightService{}
}

func (s *InsightService) GenerateInsights(expenses []models.Expense) []models.Insight {
	categoryTotals := make(map[string]float64)
	var totalSpent float64

	for _, exp := range expenses {
		categoryTotals[exp.Category] += exp.Amount
		totalSpent += exp.Amount
	}

	var insights []models.Insight

	// Insight 1: Top Spending
	if totalSpent > 0 {
		var topCategory string
		var topAmount float64
		for cat, amt := range categoryTotals {
			if amt > topAmount {
				topAmount = amt
				topCategory = cat
			}
		}
		topPercentage := (topAmount / totalSpent) * 100
		insights = append(insights, models.Insight{
			Type:        "top_spending",
			MonthlyCost: math.Round(topAmount*100) / 100,
			Percentage:  math.Round(topPercentage*10) / 10,
			Message:     fmt.Sprintf("Your biggest expense is %s (₹%.0f, %.1f%%)", topCategory, topAmount, topPercentage),
			FlagLevel:   "info",
		})
	}

	// Insight 2: Subscription Waste
	subscriptions := categoryTotals["Subscriptions"]
	if totalSpent > 0 && subscriptions > 0 {
		subPercentage := (subscriptions / totalSpent) * 100
		flag := "info"
		if subPercentage > 15 {
			flag = "alert"
		} else if subPercentage > 10 {
			flag = "warning"
		}

		insights = append(insights, models.Insight{
			Type:        "subscription_waste",
			MonthlyCost: math.Round(subscriptions*100) / 100,
			Percentage:  math.Round(subPercentage*10) / 10,
			Message:     fmt.Sprintf("Subscriptions: ₹%.0f/month (%.1f%% of spend)", subscriptions, subPercentage),
			FlagLevel:   flag,
		})
	}

	// Insight 3: High Food Spending
	food := categoryTotals["Food"]
	if totalSpent > 0 && food > 0 {
		foodPercentage := (food / totalSpent) * 100
		flag := "info"
		if foodPercentage > 35 {
			flag = "alert"
		} else if foodPercentage > 25 {
			flag = "warning"
		}

		insights = append(insights, models.Insight{
			Type:        "high_food",
			MonthlyCost: math.Round(food*100) / 100,
			Percentage:  math.Round(foodPercentage*10) / 10,
			Message:     fmt.Sprintf("Food & Dining: ₹%.0f/month (%.1f%% of spend)", food, foodPercentage),
			FlagLevel:   flag,
		})
	}

	// Insight 4: Daily Average
	if len(expenses) > 0 {
		dailyAvg := totalSpent / float64(len(expenses))
		insights = append(insights, models.Insight{
			Type:        "daily_average",
			MonthlyCost: math.Round(dailyAvg*100) / 100,
			Message:     fmt.Sprintf("Daily spending average per transaction: ₹%.0f", dailyAvg),
			FlagLevel:   "info",
		})
	}

	// Insight 5: Fixed vs Variable
	fixedCategories := map[string]bool{"Rent": true, "Subscriptions": true, "Utilities": true}
	var fixedTotal, variableTotal float64

	for cat, amt := range categoryTotals {
		if fixedCategories[cat] {
			fixedTotal += amt
		} else {
			variableTotal += amt
		}
	}

	if totalSpent > 0 {
		fixedPct := (fixedTotal / totalSpent) * 100
		variablePct := (variableTotal / totalSpent) * 100

		msg := fmt.Sprintf("Fixed: %.0f%% vs Variable: %.0f%%. Ideally, keep fixed < 50%%.", fixedPct, variablePct)
		if fixedPct > 60 {
			msg = fmt.Sprintf("Alert: Fixed expenses are %.0f%% of your budget (Target < 50%%).", fixedPct)
		}

		insights = append(insights, models.Insight{
			Type:        "fixed_vs_variable",
			MonthlyCost: math.Round(fixedTotal*100) / 100, // Showing fixed cost as the primary metric
			Percentage:  math.Round(fixedPct*10) / 10,
			Message:     msg,
			FlagLevel:   "info",
		})
	}

	// Insight 5: Category Breakdown
	breakdown := make(map[string]float64)
	for k, v := range categoryTotals {
		breakdown[k] = math.Round(v*100) / 100
	}
	insights = append(insights, models.Insight{
		Type:           "category_breakdown",
		MonthlyCost:    math.Round(totalSpent*100) / 100,
		Message:        "Here's where your money goes",
		Breakdown:      breakdown,
		FlagLevel:      "info",
		ImpactContext:  fmt.Sprintf("Total Yearly Spend: ₹%.0f", totalSpent*12),
		ActionableStep: "Check if this aligns with your goals.",
	})

	// Add Impact Context and Actionable Steps to all insights
	for i := range insights {
		if insights[i].ImpactContext == "" {
			switch insights[i].Type {
			case "subscription_waste":
				yearlyLoss := insights[i].MonthlyCost * 12
				insights[i].ImpactContext = fmt.Sprintf("You lose ₹%.0f/year — that's a weekend trip!", yearlyLoss)
				insights[i].ActionableStep = "Cancel at least 1 unused sub today."
			case "high_food":
				saved := insights[i].MonthlyCost * 0.20 // Assume 20% saving target
				insights[i].ImpactContext = fmt.Sprintf("Cooking more could save you ₹%.0f/month.", saved)
				insights[i].ActionableStep = "Limit ordering out to weekends only."
			case "fixed_vs_variable":
				insights[i].ImpactContext = "High fixed costs limit your freedom to invest."
				insights[i].ActionableStep = "Negotiate rent or downsize plans."
			case "top_spending":
				insights[i].ImpactContext = "This is your biggest wealth leak."
				insights[i].ActionableStep = "Set a strict limit for this category."
			case "daily_average":
				insights[i].ImpactContext = "Small daily habits add up."
				insights[i].ActionableStep = "Try a 'No Spend Day' once a week."
			}
		}
	}

	if len(insights) > 6 {
		return insights[:6]
	}
	return insights
}

func (s *InsightService) CalculateConfidenceScore(insights []models.Insight) int {
	score := 100
	for _, insight := range insights {
		if insight.FlagLevel == "alert" {
			score -= 15
		} else if insight.FlagLevel == "warning" {
			score -= 5
		}
	}
	if score < 0 {
		return 0
	}
	return score
}

func (s *InsightService) GenerateDashboardData(expenses []models.Expense) models.DashboardData {
	var total float64
	for _, exp := range expenses {
		total += exp.Amount
	}

	avgDaily := 0.0
	if len(expenses) > 0 {
		avgDaily = total / float64(len(expenses))
	}

	insights := s.GenerateInsights(expenses)
	confidenceScore := s.CalculateConfidenceScore(insights)

	return models.DashboardData{
		TotalExpenses:    math.Round(total*100) / 100,
		ExpenseCount:     len(expenses),
		AverageDaily:     math.Round(avgDaily*100) / 100,
		Expenses:         expenses,
		Insights:         insights,
		MonthlyBreakdown: s.GetMonthlyBreakdown(expenses),
		ConfidenceScore:  confidenceScore,
	}
}

func (s *InsightService) GetMonthlyBreakdown(expenses []models.Expense) map[string]float64 {
	breakdown := make(map[string]float64)
	for _, exp := range expenses {
		if len(exp.Date) >= 7 {
			monthKey := exp.Date[:7] // YYYY-MM
			breakdown[monthKey] += exp.Amount
		}
	}
	// Round values
	for k, v := range breakdown {
		breakdown[k] = math.Round(v*100) / 100
	}
	return breakdown
}
