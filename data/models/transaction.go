package models

import (
	"time"
)

type Transaction struct {
	ID              uint          `gorm:"primarykey" json:"id"`
	AccountID       int           `json:"account_id"`
	Account         Account       `json:"account"`
	OperationTypeID int           `json:"operation_type_id"`
	OperationType   OperationType `json:"operation-type"`
	Amount          float64       `json:"amount"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
