package repository

import (
	"pismo/data/models"
)

type MockRepository struct {
	AccountByDocumentNumber *models.Account
	AccountById             *models.Account
	CreateAccountModel      *models.Account
	CreateAccountError      error
	CreateTransactionModel  *models.Transaction
	CreateTransactionError  error
}

func (repo *MockRepository) GetAccountByDocumentNumber(documentNumber string) *models.Account {
	return repo.AccountByDocumentNumber
}

func (repo *MockRepository) GetAccountById(id uint) *models.Account {
	return repo.AccountById
}

func (repo *MockRepository) CreateAccount(account *models.Account) (*models.Account, error) {
	return repo.CreateAccountModel, repo.CreateAccountError
}

func (repo *MockRepository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	return repo.CreateTransactionModel, repo.CreateTransactionError
}
