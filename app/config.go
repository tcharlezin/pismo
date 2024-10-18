package app

import (
	"gorm.io/gorm"
	"log/slog"
	"pismo/app/setup"
	"sync"
)

type Config struct {
	Repository *gorm.DB
	Log        *slog.Logger
	Database   *sync.Mutex
}

var Application = Config{
	Log:        setup.SetupLog(),
	Repository: setup.ConnectDatabase(),
}

func (config *Config) Boot() {
}

func (config *Config) Shutdown() {
}
