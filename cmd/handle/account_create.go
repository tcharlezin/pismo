package handle

import (
	"errors"
	"net/http"
	"pismo/app"
	"pismo/cmd/input/account_create"
	"pismo/cmd/request"
	"pismo/data/models"
)

func AccountCreate(w http.ResponseWriter, r *http.Request) {

	var input account_create.AccountCreateInput
	err := request.ReadJSON(w, r, &input)

	if err != nil {
		_ = request.ErrorJSON(w, err)
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

	_ = request.WriteJSON(w, http.StatusOK, &account)
}
