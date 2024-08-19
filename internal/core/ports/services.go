package ports

import "stock-portfolio/internal/core/domain"

type PortfolioService interface {
	ExecuteTransactions(transactions []domain.Transaction) (float64, error)
	GetCurrentBalance() (map[string]int, error)
}
