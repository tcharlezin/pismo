package handle

import (
	"errors"
	"net/http"
	"pismo/app"
	"pismo/cmd/handle/transaction_create"
	"pismo/cmd/request"
	"pismo/data/models"
	"pismo/src/Transaction"
)

// TransactionCreate - Create a transaction
// @Summary Create a transaction based in the request
// @Description Create a transaction based in the request
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param transaction_create.TransactionCreateInput body transaction_create.TransactionCreateInput true "Payload"
// @Success 200 {object} transaction_create.TransactionCreateResponse
// @Failure 400 "invalid input"
// @Failure 500 "not created transaction"
// @Router /transactions [post]
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

	_ = request.WriteJSON(w, http.StatusOK, transaction_create.ResponseTo(transaction))
}
