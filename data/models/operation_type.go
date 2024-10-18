package models

import (
	"time"
)

type OperationType struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
