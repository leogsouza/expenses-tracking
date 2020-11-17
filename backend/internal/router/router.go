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
	"github.com/leogsouza/expenses-tracking/backend/internal/account"
	"github.com/leogsouza/expenses-tracking/backend/internal/category"
	"github.com/leogsouza/expenses-tracking/backend/internal/container"
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

func New(container *container.Services) http.Handler {

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
		r.Mount("/transactions", transactionRoutes(container.Transaction))
		r.Mount("/users", userRoutes(container.User))
		r.Mount("/accounts", accountRoutes(container.Account))
		r.Mount("/categories", categoryRoutes(container.Category))
	})

	flag.Parse()

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

func transactionRoutes(serv transaction.Service) http.Handler {

	return transaction.NewHandler(serv).Routes()
}

func userRoutes(serv user.Service) http.Handler {

	return user.NewHandler(serv).Routes()
}

func accountRoutes(serv account.Service) http.Handler {

	return account.NewHandler(serv).Routes()
}

func categoryRoutes(serv category.Service) http.Handler {
	return category.NewHandler(serv).Routes()
}
