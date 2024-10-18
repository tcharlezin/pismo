package handle

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"pismo/app"
	"pismo/cmd/request"
	"pismo/data/models"
)

func AccountGet(w http.ResponseWriter, r *http.Request) {

	accountId := chi.URLParam(r, "accountId")

	var account models.Account

	result := app.Application.Repository.First(&account, "id = ?", accountId)

	if result.RowsAffected == 0 {
		_ = request.ErrorJSON(w, errors.New("account not found"))
		return
	}

	_ = request.WriteJSON(w, http.StatusOK, &account)
}
