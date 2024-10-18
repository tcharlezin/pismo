package models

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	ID              uint          `gorm:"primarykey" json:"id"`
	AccountID       uint          `json:"account_id"`
	Account         Account       `json:"account"`
	OperationTypeID uint          `json:"operation_type_id"`
	OperationType   OperationType `json:"operation-type"`
	Amount          float64       `json:"amount"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

func (model *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TransactionID   uint    `json:"transaction_id"`
		AccountID       uint    `json:"account_id"`
		OperationTypeID uint    `json:"operation_type_id"`
		Amount          float64 `json:"amount"`
	}{
		TransactionID:   model.ID,
		AccountID:       model.AccountID,
		OperationTypeID: model.OperationTypeID,
		Amount:          model.Amount,
	})
}
