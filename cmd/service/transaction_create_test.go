package service

import (
	"pismo/app/setup"
	"pismo/cmd/service/transaction_create"
	"pismo/data/models"
	"pismo/test/repository"
	"testing"
	"time"
)

func TestTransactionCreate(t *testing.T) {

	t.Log("TestTransactionCreate")

	s := Service{
		Repository: &repository.MockRepository{
			AccountById: &models.Account{
				ID:             1,
				DocumentNumber: "123456",
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			},
			CreateTransactionModel: &models.Transaction{
				ID:              1,
				AccountID:       1,
				OperationTypeID: 1,
				Amount:          123.50,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
			},
			CreateTransactionError: nil,
		},
		Log: setup.SetupLog(),
	}

	input := transaction_create.TransactionCreateInput{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          123.50,
	}

	response, err := s.TransactionCreate(input)

	if err != nil {
		t.Log("Error! Should be nil!")
		t.FailNow()
	}

	if response.AccountID != 1 {
		t.Log("Error! Should be 1!")
		t.FailNow()
	}

	if response.TransactionID != 1 {
		t.Log("Error! Should be 1!")
		t.FailNow()
	}

	if response.OperationTypeID != 1 {
		t.Log("Error! Should be 1!")
		t.FailNow()
	}

	if response.Amount != 123.50 {
		t.Log("Error! Should be 123.50!")
		t.FailNow()
	}
}
