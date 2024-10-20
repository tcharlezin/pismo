package account_create

import (
	"pismo/app"
	"pismo/data/models"
)

type AccountCreateInput struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

func Validate(input AccountCreateInput) bool {

	var account models.Account
	account.DocumentNumber = input.DocumentNumber

	result := app.Application.Repository.First(&account, "document_number = ?", account.DocumentNumber)

	if result.RowsAffected != 0 {
		return false
	}

	return true
}
