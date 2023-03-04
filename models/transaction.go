package models

import "time"

type Transaction struct {
	ID            int       `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int       `json:"user_id"`
	User          User      `json:"user"`
	TotalQuantity int       `json:"total_quantity"`
	TotalPrice    int       `json:"total_price"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

type TransactionUSerResponse struct {
	User          User `json:"user"`
	TotalQuantity int  `json:"total_quantity"`
	TotalPrice    int  `json:"total_price"`
}

func (TransactionUSerResponse) TableName() string {
	return "transactions"
}
