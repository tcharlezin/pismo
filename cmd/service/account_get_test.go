package service

import (
	"pismo/app/setup"
	"pismo/cmd/service/account_get"
	"pismo/data/models"
	"pismo/test/repository"
	"testing"
	"time"
)

func TestAccountGet(t *testing.T) {

	t.Log("TestAccountGet")

	s := Service{
		Repository: &repository.MockRepository{
			AccountById: &models.Account{
				ID:             5,
				DocumentNumber: "123456",
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			},
		},
		Log: setup.SetupLog(),
	}

	input := account_get.AccountGetInput{AccountID: 5}

	response, err := s.AccountGet(input)

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

func TestAccountGetNotExist(t *testing.T) {

	t.Log("TestAccountGet")

	s := Service{
		Repository: &repository.MockRepository{
			AccountById: nil,
		},
		Log: setup.SetupLog(),
	}

	input := account_get.AccountGetInput{AccountID: 5}

	response, err := s.AccountGet(input)

	if err == nil {
		t.Log("Error! Should not be nil!")
		t.FailNow()
	}

	if err.Error() != "account not found" {
		t.Log("Error! Should be \"account not found\"!")
		t.FailNow()
	}

	if response != nil {
		t.Log("Error! Should be nil!")
		t.FailNow()
	}
}
