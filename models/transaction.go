package models

import "time"

type Transaction struct {
	ID            int                       `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int                       `json:"user_id"`
	User          UserTransactionResponse   `json:"user"`
	Cart          []CartTransactionResponse `json:"carts"`
	TotalQuantity int                       `json:"total_quantity"`
	TotalPrice    int                       `json:"total_price"`
	CreatedAt     time.Time                 `json:"-"`
	UpdatedAt     time.Time                 `json:"-"`
}

type TransactionUSerResponse struct {
	ID            int `json:"transaction_id"`
	UserID        int `json:"-"`
	TotalQuantity int `json:"total_quantity"`
	TotalPrice    int `json:"total_price"`
}

func (TransactionUSerResponse) TableName() string {
	return "transactions"
}
