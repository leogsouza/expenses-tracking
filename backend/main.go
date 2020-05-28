package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/leogsouza/expenses-tracking/server/internal/entity"
	"github.com/leogsouza/expenses-tracking/server/internal/router"
)

var port = "8080"

func main() {

	// Logger
	db, err := gorm.Open("sqlite3", "expenses.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&entity.Transaction{})

	r := router.New(db)

	log.Printf("accepting connections on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}

}
