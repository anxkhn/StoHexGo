package repositories

import (
	"stock-portfolio/internal/core/domain"

	"gorm.io/gorm"
)

type sqliteRepository struct {
	db *gorm.DB
}

func NewSQLiteRepository(db *gorm.DB) *sqliteRepository {
	return &sqliteRepository{db: db}
}

func (r *sqliteRepository) SaveTransaction(transaction domain.Transaction) error {
	return r.db.Create(&transaction).Error
}

func (r *sqliteRepository) GetAllTransactions() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
