package router

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/httplog"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
	"github.com/leogsouza/expenses-tracking/backend/internal/account"
	"github.com/leogsouza/expenses-tracking/backend/internal/category"
	"github.com/leogsouza/expenses-tracking/backend/internal/transaction"
	"github.com/leogsouza/expenses-tracking/backend/internal/user"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

var logger = httplog.NewLogger("httplog-example", httplog.Options{
	JSON: true,
	//Concise: true,
	// Tags: map[string]string{
	// 	"version": "v1.0-81aa4244d9fc8076a",
	// 	"env":     "dev",
	// },
})

func New(db *gorm.DB) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Heartbeat("/ping"))

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
		r.Mount("/transactions", transactionRoutes(db))
		r.Mount("/users", userRoutes(db))
		r.Mount("/accounts", accountRoutes(db))
		r.Mount("/categories", categoryRoutes(db))
	})

	if *routes {
		// fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/go-chi/chi",
			Intro:       "Welcome to the chi/_examples/rest generated docs.",
		}))
		//return
	}
	return r

}

func healthcheck(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func transactionRoutes(db *gorm.DB) http.Handler {
	repo, err := transaction.NewRepository(db)
	if err != nil {
		logger.Fatal().Err(err)
	}

	serv := transaction.NewService(repo)

	return transaction.NewHandler(serv).Routes()
}

func userRoutes(db *gorm.DB) http.Handler {
	repo, err := user.NewRepository(db)
	if err != nil {
		logger.Fatal().Err(err)
	}

	serv := user.NewService(repo)

	return user.NewHandler(serv).Routes()
}

func accountRoutes(db *gorm.DB) http.Handler {
	repo, err := account.NewRepository(db)
	if err != nil {
		logger.Fatal().Err(err)
	}

	serv := account.NewService(repo)

	return account.NewHandler(serv).Routes()
}

func categoryRoutes(db *gorm.DB) http.Handler {
	repo, err := category.NewRepository(db)
	if err != nil {
		logger.Fatal().Err(err)
	}

	serv := category.NewService(repo)

	return category.NewHandler(serv).Routes()
}
