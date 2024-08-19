package domain

type TransactionType string

const (
	Buy  TransactionType = "buy"
	Sell TransactionType = "sell"
)

type Transaction struct {
	ID       uint            `gorm:"primaryKey"`
	Type     TransactionType `json:"type"`
	StockID  string          `json:"stockId"`
	Price    float64         `json:"price"`
	Quantity int             `json:"quantity"`
}

type StockBalance struct {
	StockID  string
	Quantity int
}
