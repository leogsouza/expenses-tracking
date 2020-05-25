package main

import (
	"log"
	"net/http"

	"github.com/leogsouza/expenses-tracking/server/internal/router"
)

var port = "8080"

func main() {

	r := router.New()

	log.Printf("accepting connections on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}

}
