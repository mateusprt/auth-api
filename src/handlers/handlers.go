package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mateusprt/auth-api/src/config"
	"github.com/mateusprt/auth-api/src/middlewares"
)

var db = config.CreateConnection()

func LoadAll() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/ping"))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         300,
	}))

	router.Post("/auth/register", RegisterHandler)
	router.Post("/auth/confirmation", ConfirmationHandler)
	router.Post("/auth/login", LoginHandler)

	router.Post("/auth/reset", ResetHandler)
	router.Post("/auth/reset/confirmation", ResetConfirmationHandler)
	//router.Post("/auth/logout", func(w http.ResponseWriter, r *http.Request) {})
	router.Get("/home", middlewares.Autentication(HomeHandler))

	return router
}
