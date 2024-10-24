package handle

import (
	"log/slog"
	"pismo/cmd/repository"
)

type Application struct {
	Repository repository.IRepository
	Log        *slog.Logger
}

func (app *Application) Boot() {
}

func (app *Application) Shutdown() {
}
