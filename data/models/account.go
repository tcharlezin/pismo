package models

import (
	"encoding/json"
	"time"
)

type Account struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	DocumentNumber string    `gorm:"unique;not null" json:"document_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (model *Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		AccountID      uint   `json:"account_id"`
		DocumentNumber string `json:"document_number"`
	}{
		AccountID:      model.ID,
		DocumentNumber: model.DocumentNumber,
	})
}
