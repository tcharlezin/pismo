package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pismo/app/setup"
	"pismo/cmd/api"
	"pismo/cmd/handle"
	"pismo/cmd/repository"
)

// @title         Pismo Test API
// @version       1.0
// @description   Implementation of Pismo Test API.
// @contact.name  Tcharles Pereira da Silva
// @contact.url   https://tcharlez.in
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	app := handle.Application{
		Log: setup.SetupLog(),
	}
	repo := repository.InitDB()
	app.Repository = &repo

	app.Boot()
	defer app.Shutdown()

	webPort := os.Getenv("WEB_PORT")

	// Define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api.Routes(app),
	}

	// Start server
	app.Log.Info(fmt.Sprintf("Starting webserver at port %s", webPort))

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic("[Server] ", err)
	}
}
