package handle

import (
	"errors"
	"net/http"
	"pismo/app"
	"pismo/cmd/input/transaction_create"
	"pismo/cmd/request"
	"pismo/data/models"
	"pismo/src/Transaction"
)

func TransactionCreate(w http.ResponseWriter, r *http.Request) {

	var input transaction_create.TransactionCreateInput
	err := request.ReadJSON(w, r, &input)

	if err != nil {
		app.Application.Log.Error("Invalid input: ", err.Error())
		_ = request.ErrorJSON(w, errors.New("invalid input"))
		return
	}

	if !transaction_create.Validate(input) {
		_ = request.ErrorJSON(w, errors.New("invalid input"))
		return
	}

	var transaction models.Transaction
	transaction.AccountID = input.AccountID
	transaction.OperationTypeID = input.OperationTypeID
	transaction.Amount = Transaction.CalculateAmount(input)

	result := app.Application.Repository.Create(&transaction)

	if result.RowsAffected == 0 {
		app.Application.Log.Error("Not created transaction: ", result.Error.Error())
		_ = request.ErrorJSON(w, errors.New("not created transaction"), http.StatusInternalServerError)
		return
	}

	_ = request.WriteJSON(w, http.StatusOK, struct {
		TransactionID   uint    `json:"transaction_id"`
		AccountID       uint    `json:"account_id"`
		OperationTypeID uint    `json:"operation_type_id"`
		Amount          float64 `json:"amount"`
	}{
		TransactionID:   transaction.ID,
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          input.Amount,
	})
}
