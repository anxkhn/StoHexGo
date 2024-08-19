package ports

import "stock-portfolio/internal/core/domain"

type TransactionRepository interface {
	SaveTransaction(transaction domain.Transaction) error
	GetAllTransactions() ([]domain.Transaction, error)
}
