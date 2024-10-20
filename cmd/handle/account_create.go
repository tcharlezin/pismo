package handle

import (
	"errors"
	"net/http"
	"pismo/app"
	"pismo/cmd/handle/account_create"
	"pismo/cmd/request"
	"pismo/data/models"
)

// AccountCreate - Create an account
// @Summary Create one account for the document number.
// @Description Create one account for the document number.
// @Tags Account
// @Accept  json
// @Produce  json
// @Param account_create.AccountCreateInput body account_create.AccountCreateInput true "Payload"
// @Success 200 {object} account_create.AccountCreateResponse
// @Failure 400 "invalid input"
// @Failure 500 "not created account"
// @Router /accounts [post]
func AccountCreate(w http.ResponseWriter, r *http.Request) {

	var input account_create.AccountCreateInput
	err := request.ReadJSON(w, r, &input)

	if err != nil {
		app.Application.Log.Error("Invalid input: ", err.Error())
		_ = request.ErrorJSON(w, errors.New("invalid input"))
		return
	}

	if !account_create.Validate(input) {
		_ = request.ErrorJSON(w, errors.New("invalid input"))
		return
	}

	var account models.Account
	account.DocumentNumber = input.DocumentNumber

	result := app.Application.Repository.Create(&account)

	if result.RowsAffected == 0 {
		app.Application.Log.Error("Not created account: ", result.Error.Error())
		_ = request.ErrorJSON(w, errors.New("not created account"), http.StatusInternalServerError)
		return
	}

	_ = request.WriteJSON(w, http.StatusOK, account_create.ResponseTo(account))
}
