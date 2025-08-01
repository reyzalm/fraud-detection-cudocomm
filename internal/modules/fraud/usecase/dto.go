package usecase

type FraudResultDTO struct {
	TransactionID   int64          `json:"transaction_id"`
	UserID          int64          `json:"user_id"`
	FinalFraudScore float64        `json:"final_fraud_score"`
	RiskLevel       string         `json:"risk_level"`
	Breakdown       ScoreBreakdown `json:"breakdown"`
}

type ScoreBreakdown struct {
	FrequencyScore float64 `json:"frequency_score"`
	AmountScore    float64 `json:"amount_score"`
	PatternScore   float64 `json:"pattern_score"`
}
