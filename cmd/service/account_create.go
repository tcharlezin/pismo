package service

import (
	"errors"
	"pismo/cmd/service/account_create"
	"pismo/data/models"
)

func (s *Service) AccountCreate(input account_create.AccountCreateInput) (*account_create.AccountCreateResponse, error) {

	if !input.Validate() {
		return nil, errors.New("invalid input")
	}

	var account *models.Account
	account = s.Repository.GetAccountByDocumentNumber(input.DocumentNumber)
	if account != nil {
		return nil, errors.New("invalid input")
	}

	account = &models.Account{
		DocumentNumber: input.DocumentNumber,
	}

	account, err := s.Repository.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	return account_create.ResponseTo(account), nil
}
