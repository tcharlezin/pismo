package service

import (
	"errors"
	"pismo/cmd/service/account_get"
)

func (s *Service) AccountGet(input account_get.AccountGetInput) (*account_get.AccountGetResponse, error) {

	account := s.Repository.GetAccountById(input.AccountID)

	if account == nil {
		return nil, errors.New("account not found")
	}

	return account_get.ResponseTo(account), nil
}
