package usecase

import (
	"math"
	"sync"
	"time"

	"github.com/CudoCommunication/cudocomm/internal/domain"
	"github.com/CudoCommunication/cudocomm/internal/domain/models"
	"github.com/CudoCommunication/cudocomm/internal/domain/repository"
)

type FraudUseCase interface {
	DetectFraud(userID, transactionID int64) (*FraudResultDTO, error)
}

type fraudUseCaseImpl struct {
	logger    domain.Logger
	transRepo repository.TransactionRepository
}

func NewFraudUseCase(logger domain.Logger, transRepo repository.TransactionRepository) FraudUseCase {
	return &fraudUseCaseImpl{
		logger:    logger,
		transRepo: transRepo,
	}
}

func (u *fraudUseCaseImpl) DetectFraud(userID, transactionID int64) (*FraudResultDTO, error) {
	currentTx, err := u.transRepo.GetTransactionByID(transactionID)
	if err != nil {
		u.logger.Error(&domain.LoggerPayload{Loc: "DetectFraud.GetTransactionByID", Msg: err.Error()})
		return nil, err
	}

	historicalTxs, err := u.transRepo.GetUserTransactions(userID, currentTx.TransactionDate)
	if err != nil {
		u.logger.Error(&domain.LoggerPayload{Loc: "DetectFraud.GetUserTransactions", Msg: err.Error()})
		return nil, err
	}

	var wg sync.WaitGroup
	results := make(chan float64, 3)

	wg.Add(3)

	go func() {
		defer wg.Done()
		score := u.frequencyCheck(currentTx, historicalTxs)
		results <- score
	}()

	go func() {
		defer wg.Done()
		score := u.amountCheck(currentTx, historicalTxs)
		results <- score
	}()

	go func() {
		defer wg.Done()
		score := u.patternCheck(currentTx, historicalTxs)
		results <- score
	}()

	wg.Wait()
	close(results)

	var scores []float64
	for score := range results {
		scores = append(scores, score)
	}

	breakdown := ScoreBreakdown{
		FrequencyScore: scores[0],
		AmountScore:    scores[1],
		PatternScore:   scores[2],
	}

	finalScore := (breakdown.FrequencyScore * 0.4) + (breakdown.AmountScore * 0.3) + (breakdown.PatternScore * 0.3)

	var riskLevel string
	if finalScore >= 80 {
		riskLevel = "HIGH"
	} else if finalScore >= 50 {
		riskLevel = "MEDIUM"
	} else {
		riskLevel = "LOW"
	}

	return &FraudResultDTO{
		TransactionID:   transactionID,
		UserID:          userID,
		FinalFraudScore: finalScore,
		RiskLevel:       riskLevel,
		Breakdown:       breakdown,
	}, nil
}

// Logika Pengecekan

func (u *fraudUseCaseImpl) frequencyCheck(currentTx *models.Transaction, historicalTxs []models.Transaction) float64 {
	oneHourAgo := currentTx.TransactionDate.Add(-1 * time.Hour)
	count := 1

	for _, tx := range historicalTxs {
		if tx.TransactionDate.After(oneHourAgo) {
			count++
		} else {
			break
		}
	}

	switch {
	case count > 8:
		return 95.0
	case count >= 7:
		return 85.0
	case count >= 6:
		return 75.0
	case count >= 5:
		return 60.0
	default:
		return 20.0
	}
}

func (u *fraudUseCaseImpl) amountCheck(currentTx *models.Transaction, historicalTxs []models.Transaction) float64 {
	if len(historicalTxs) < 2 {
		return 0.0
	}

	var sum, stdDevSum float64
	for _, tx := range historicalTxs {
		sum += tx.Amount
	}
	mean := sum / float64(len(historicalTxs))

	for _, tx := range historicalTxs {
		stdDevSum += math.Pow(tx.Amount-mean, 2)
	}
	stdDev := math.Sqrt(stdDevSum / float64(len(historicalTxs)))

	if stdDev == 0 {
		return 100.0
	}

	zScore := (currentTx.Amount - mean) / stdDev

	if zScore <= 2.0 {
		return 0.0
	}

	score := (zScore - 2.0) * 20.0
	return math.Min(100.0, score)
}

func (u *fraudUseCaseImpl) patternCheck(currentTx *models.Transaction, historicalTxs []models.Transaction) float64 {
	if len(historicalTxs) == 0 {
		return 0.0
	}

	var sum float64
	for _, tx := range historicalTxs {
		sum += tx.Amount
	}
	baseline := sum / float64(len(historicalTxs))

	if baseline == 0 {
		return 100.0
	}

	percentIncrease := ((currentTx.Amount - baseline) / baseline) * 100

	if percentIncrease <= 0 {
		return 0.0
	}

	score := percentIncrease / 4.0
	return math.Min(100.0, score)
}
