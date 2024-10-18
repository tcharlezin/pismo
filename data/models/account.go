package models

import (
	"time"
)

type Account struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	DocumentNumber string    `json:"document_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
