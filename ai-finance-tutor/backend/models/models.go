package models

type Expense struct {
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
}

type Insight struct {
	Type           string             `json:"type"` // "subscription_waste", "high_food", etc.
	MonthlyCost    float64            `json:"monthly_cost"`
	Percentage     float64            `json:"percentage,omitempty"`
	Message        string             `json:"message"`
	FlagLevel      string             `json:"flag_level"` // "info", "warning", "alert"
	Breakdown      map[string]float64 `json:"breakdown,omitempty"`
	ImpactContext  string             `json:"impact_context"`  // e.g. "At this rate, you lose ₹21,600/year."
	ActionableStep string             `json:"actionable_step"` // e.g. "Cancel 2 unused subscriptions."
}

type DashboardData struct {
	TotalExpenses    float64            `json:"total_expenses"`
	ExpenseCount     int                `json:"expense_count"`
	AverageDaily     float64            `json:"average_daily"`
	Expenses         []Expense          `json:"expenses"`
	Insights         []Insight          `json:"insights"`
	MonthlyBreakdown map[string]float64 `json:"monthly_breakdown"`
	ConfidenceScore  int                `json:"confidence_score"` // 0-100 Financial Health Score
}

type TutorRequest struct {
	Insight  Insight `json:"insight"`
	FollowUp string  `json:"follow_up,omitempty"`
}

type TutorResponse struct {
	Explanation string `json:"explanation"`
	Tip         string `json:"tip"`
	NextStep    string `json:"next_step"`
}

type PersonaResponse struct {
	Archetype   string `json:"archetype"`    // e.g. "The Coffee Shop Philanthropist"
	Emoji       string `json:"emoji"`        // e.g. ☕️
	Description string `json:"description"`  // Fun description
	SavageQuote string `json:"savage_quote"` // A roast
}
