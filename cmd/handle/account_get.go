package handle

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"pismo/app"
	"pismo/cmd/handle/account_get"
	"pismo/cmd/request"
	"pismo/data/models"
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
func AccountGet(w http.ResponseWriter, r *http.Request) {

	accountId := chi.URLParam(r, "accountID")

	var account models.Account

	result := app.Application.Repository.First(&account, "id = ?", accountId)

	if result.RowsAffected == 0 {
		_ = request.ErrorJSON(w, errors.New("account not found"))
		return
	}

	_ = request.WriteJSON(w, http.StatusOK, account_get.ResponseTo(account))
}
