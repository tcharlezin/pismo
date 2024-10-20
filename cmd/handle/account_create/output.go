package account_create

import "pismo/data/models"

type AccountCreateResponse struct {
	AccountID      uint   `json:"account_id" validate:"required"`
	DocumentNumber string `json:"document_number" validate:"required"`
}

func ResponseTo(account models.Account) AccountCreateResponse {
	return AccountCreateResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	}
}
