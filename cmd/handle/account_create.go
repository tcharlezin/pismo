package handle

import (
	"errors"
	"net/http"
	"pismo/cmd/request"
	"pismo/cmd/service"
	"pismo/cmd/service/account_create"
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
func (app *Application) AccountCreate(w http.ResponseWriter, r *http.Request) {

	var input account_create.AccountCreateInput
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

	response, err := svc.AccountCreate(input)

	if err != nil {
		_ = request.ErrorJSON(w, errors.New("not created account"), http.StatusInternalServerError)
		return
	}

	_ = request.WriteJSON(w, http.StatusOK, &response)
}
