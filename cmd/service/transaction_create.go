package service

import (
	"errors"
	"pismo/cmd/service/transaction_create"
	"pismo/data/models"
	"pismo/src/Transaction"
)

func (s *Service) TransactionCreate(input transaction_create.TransactionCreateInput) (*transaction_create.TransactionCreateResponse, error) {

	if !input.Validate() {
		return nil, errors.New("invalid input")
	}

	account := s.Repository.GetAccountById(input.AccountID)
	if account == nil {
		return nil, errors.New("invalid input")
	}

	transaction := &models.Transaction{
		AccountID:       input.AccountID,
		OperationTypeID: input.OperationTypeID,
		Amount:          Transaction.CalculateAmount(input),
	}

	transaction, err := s.Repository.CreateTransaction(transaction)

	if err != nil {
		return nil, errors.New("not created transaction")
	}

	return transaction_create.ResponseTo(transaction), nil

}
