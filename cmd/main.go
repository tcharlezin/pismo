package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pismo/app"
	"pismo/cmd/api"
)

func main() {

	app.Application.Boot()
	defer app.Application.Shutdown()

	webPort := os.Getenv("WEB_PORT")

	// Define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api.Routes(),
	}

	// Start server
	app.Application.Log.Info(fmt.Sprintf("Starting webserver at port %s", webPort))

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic("[Server] ", err)
	}
}
