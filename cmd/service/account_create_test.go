package service

import (
	"pismo/app/setup"
	"pismo/cmd/service/account_create"
	"pismo/data/models"
	"pismo/test/repository"
	"testing"
	"time"
)

func TestAccountCreate(t *testing.T) {

	t.Log("TestAccountCreate")

	s := Service{
		Repository: &repository.MockRepository{
			AccountByDocumentNumber: nil,
			CreateAccountModel: &models.Account{
				ID:             5,
				DocumentNumber: "123456",
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			},
			CreateAccountError: nil,
		},
		Log: setup.SetupLog(),
	}

	input := account_create.AccountCreateInput{DocumentNumber: "123456"}

	response, err := s.AccountCreate(input)

	if err != nil {
		t.Log("Error! Should be nil!")
		t.FailNow()
	}

	if response.AccountID != 5 {
		t.Log("Error! Should be 5!")
		t.FailNow()
	}

	if response.DocumentNumber != "123456" {
		t.Log("Error! Should be 123456!")
		t.FailNow()
	}
}

func TestAccountAlreadyExistent(t *testing.T) {

	t.Log("TestAccountAlreadyExistent")

	s := Service{
		Repository: &repository.MockRepository{
			AccountByDocumentNumber: &models.Account{
				ID:             5,
				DocumentNumber: "123456",
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			},
		},
		Log: setup.SetupLog(),
	}

	input := account_create.AccountCreateInput{DocumentNumber: "123456"}

	response, err := s.AccountCreate(input)

	if err == nil {
		t.Log("Error! Should have error!")
		t.FailNow()
	}

	if err.Error() != "invalid input" {
		t.Log("Error! Should be 'invalid input'")
		t.FailNow()
	}

	if response != nil {
		t.Log("Error! Should be nil")
		t.FailNow()
	}

}
