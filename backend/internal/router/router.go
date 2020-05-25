package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/leogsouza/expenses-tracking/server/internal/transaction"
)

func New() http.Handler {

	// Logger
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		JSON: true,
		//Concise: true,
		// Tags: map[string]string{
		// 	"version": "v1.0-81aa4244d9fc8076a",
		// 	"env":     "dev",
		// },
	})

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Heartbeat("/ping"))

	repo, err := transaction.NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	serv := transaction.NewService(repo)
	trHandler := transaction.NewHandler(serv)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/", healthcheck)
	r.Route("/api", func(r chi.Router) {
		r.Mount("/transactions", trHandler.Routes())
	})
	return r

}

func healthcheck(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
