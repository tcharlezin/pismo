package repository

import "pismo/data/models"

type IRepository interface {
	GetAccountById(id uint) *models.Account
	GetAccountByDocumentNumber(documentNumber string) *models.Account
	CreateAccount(account *models.Account) (*models.Account, error)
	CreateTransaction(transaction *models.Transaction) (*models.Transaction, error)
}
