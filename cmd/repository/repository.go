package repository

import (
	"errors"
	"gorm.io/gorm"
	"pismo/app/setup"
	"pismo/data/models"
)

type Repository struct {
	db *gorm.DB
}

func InitDB() Repository {
	repo := Repository{
		db: setup.ConnectDatabase(),
	}

	return repo
}

func (repo *Repository) GetAccountById(id uint) *models.Account {
	account := models.Account{}
	result := repo.db.First(&account, "id = ?", id)

	if result.RowsAffected != 0 {
		return &account
	}

	return nil
}

func (repo *Repository) GetAccountByDocumentNumber(documentNumber string) *models.Account {
	account := models.Account{}
	result := repo.db.First(&account, "document_number = ?", documentNumber)

	if result.RowsAffected != 0 {
		return &account
	}

	return nil
}

func (repo *Repository) CreateAccount(account *models.Account) (*models.Account, error) {
	result := repo.db.Create(&account)
	if result.RowsAffected == 0 {
		return nil, errors.New("not created account")
	}
	return account, nil
}

func (repo *Repository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	result := repo.db.Create(&transaction)
	if result.RowsAffected == 0 {
		return nil, errors.New("not created transaction")
	}
	return transaction, nil
}
