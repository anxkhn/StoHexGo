package services

import (
	"fmt"
	"stock-portfolio/internal/core/domain"
	"stock-portfolio/internal/core/ports"
)

type portfolioService struct {
	repo ports.TransactionRepository
}

func NewPortfolioService(repo ports.TransactionRepository) ports.PortfolioService {
	return &portfolioService{repo: repo}
}

func (s *portfolioService) ExecuteTransactions(transactions []domain.Transaction) (float64, error) {
	totalProfit := 0.0
	stockBalances := make(map[string][]domain.Transaction)

	for _, t := range transactions {
		if t.Type == domain.Buy {
			stockBalances[t.StockID] = append(stockBalances[t.StockID], t)
		} else {
			sellQuantity := t.Quantity
			for sellQuantity > 0 {
				if len(stockBalances[t.StockID]) == 0 {
					return 0, fmt.Errorf("not enough shares to sell for stock %s", t.StockID)
				}
				oldestBuy := &stockBalances[t.StockID][0]
				if oldestBuy.Quantity <= sellQuantity {
					profit := float64(oldestBuy.Quantity) * (t.Price - oldestBuy.Price)
					totalProfit += profit
					sellQuantity -= oldestBuy.Quantity
					stockBalances[t.StockID] = stockBalances[t.StockID][1:]
				} else {
					profit := float64(sellQuantity) * (t.Price - oldestBuy.Price)
					totalProfit += profit
					oldestBuy.Quantity -= sellQuantity
					sellQuantity = 0
				}
			}
		}
		if err := s.repo.SaveTransaction(t); err != nil {
			return 0, err
		}
	}
	return totalProfit, nil
}

func (s *portfolioService) GetCurrentBalance() (map[string]int, error) {
	transactions, err := s.repo.GetAllTransactions()
	if err != nil {
		return nil, err
	}

	balances := make(map[string]int)
	for _, t := range transactions {
		if t.Type == domain.Buy {
			balances[t.StockID] += t.Quantity
		} else {
			balances[t.StockID] -= t.Quantity
		}
	}

	return balances, nil
}
