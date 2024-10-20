package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"pismo/cmd/handle"
	_ "pismo/docs"
)

func Routes() http.Handler {

	router := chi.NewRouter()

	// specify who is allowed to connect
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(middleware.Heartbeat("/ping"))

	router.Post("/accounts", handle.AccountCreate)
	router.Get("/accounts/{accountID}", handle.AccountGet)
	router.Post("/transactions", handle.TransactionCreate)

	router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/docs/doc.json"), //The url pointing to API definition
	))

	return router
}
