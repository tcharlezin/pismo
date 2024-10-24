package handle

import (
	"errors"
	"net/http"
	"pismo/cmd/request"
	"pismo/cmd/service"
	"pismo/cmd/service/transaction_create"
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
func (app *Application) TransactionCreate(w http.ResponseWriter, r *http.Request) {

	var input transaction_create.TransactionCreateInput
	err := request.ReadJSON(w, r, &input)

	if err != nil {
		app.Log.Error("Invalid input: ", "err", err.Error())
		_ = request.ErrorJSON(w, errors.New("invalid input"))
		return
	}

	svc := service.Service{
		Repository: app.Repository,
		Log:        app.Log,
	}

	response, err := svc.TransactionCreate(input)

	if err != nil {
		_ = request.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = request.WriteJSON(w, http.StatusOK, &response)
}
