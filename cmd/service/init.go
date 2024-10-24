package service

import (
	"log/slog"
	"pismo/cmd/repository"
)

type Service struct {
	Repository repository.IRepository
	Log        *slog.Logger
}
