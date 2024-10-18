package transaction_create

import (
	"pismo/app"
	"pismo/data/models"
)

type TransactionCreateInput struct {
	AccountID       uint    `json:"account_id"`
	OperationTypeID uint    `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

func Validate(input TransactionCreateInput) bool {
	if input.OperationTypeID < 1 || input.OperationTypeID > 4 {
		return false
	}

	if input.Amount < 0 {
		return false
	}

	var account models.Account
	result := app.Application.Repository.First(&account, "id = ?", input.AccountID)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
