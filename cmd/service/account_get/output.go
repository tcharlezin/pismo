package account_get

import "pismo/data/models"

type AccountGetResponse struct {
	AccountID      uint   `json:"account_id" validate:"required"`
	DocumentNumber string `json:"document_number" validate:"required"`
}

func ResponseTo(account *models.Account) *AccountGetResponse {
	return &AccountGetResponse{
		AccountID:      account.ID,
		DocumentNumber: account.DocumentNumber,
	}
}
