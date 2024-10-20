package transaction_create

import "pismo/data/models"

type TransactionCreateResponse struct {
	TransactionID   uint    `json:"transaction_id" validate:"required"`
	AccountID       uint    `json:"account_id" validate:"required"`
	OperationTypeID uint    `json:"operation_type_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
}

func ResponseTo(transaction models.Transaction) TransactionCreateResponse {

	tmpAmount := transaction.Amount
	if tmpAmount < 0 {
		tmpAmount = tmpAmount * (-1)
	}

	return TransactionCreateResponse{
		TransactionID:   transaction.ID,
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          tmpAmount,
	}
}
