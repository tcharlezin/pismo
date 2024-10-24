package handle

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"pismo/cmd/request"
	"pismo/cmd/service"
	"pismo/cmd/service/account_get"
	"strconv"
)

// AccountGet - Get an account
// @Summary Get one account based in the query string parameter.
// @Description Get one account based in the query string parameter.
// @Tags Account
// @Accept  json
// @Produce  json
// @Param accountID path string true "Account ID"
// @Success 200 {object} account_get.AccountGetResponse
// @Failure 400 "account not found"
// @Router /accounts/{accountID} [get]
func (app *Application) AccountGet(w http.ResponseWriter, r *http.Request) {

	accountId := chi.URLParam(r, "accountID")
	id, _ := strconv.Atoi(accountId)

	input := account_get.AccountGetInput{AccountID: uint(id)}

	svc := service.Service{
		Repository: app.Repository,
		Log:        app.Log,
	}

	response, err := svc.AccountGet(input)

	if err != nil {
		_ = request.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	_ = request.WriteJSON(w, http.StatusOK, &response)
}
