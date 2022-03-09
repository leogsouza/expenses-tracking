package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/leogsouza/expenses-tracking/backend/internal/account"
	"github.com/leogsouza/expenses-tracking/backend/internal/category"
	"github.com/leogsouza/expenses-tracking/backend/internal/container"
	"github.com/leogsouza/expenses-tracking/backend/internal/router"
	"github.com/leogsouza/expenses-tracking/backend/internal/transaction"
	"github.com/leogsouza/expenses-tracking/backend/internal/user"
	"github.com/tinrab/kit/retry"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found to be loaded")
	}
}

type Config struct {
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresDB       string `envconfig:"POSTGRES_DB"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	PostgresPort     string `envconfig:"POSTGRES_PORT"`
}

func main() {

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var (
		databaseAddr = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresDB, cfg.PostgresPassword)
		port = env("PORT", "8080")
		db   *gorm.DB
	)

	// Logger
	log.Println("DatabaseAddr", databaseAddr)

	retry.ForeverSleep(2*time.Second, func(attempt int) error {
		db, err = gorm.Open("postgres", databaseAddr)

		if err != nil {
			log.Printf("could not open db connection: %v\n", err)
			return err
		}
		return nil
	})

	defer db.Close()

	// Migrate the schema
	err = dbMigration(db.DB())
	if err != nil {
		log.Printf("could not process the db migration: %v\n", err)
		//return
	}

	if err = db.DB().Ping(); err != nil {
		log.Printf("could not ping to db: %v\n", err)
		//return
	}

	accRepo, _ := account.NewRepository(db)
	catRepo, _ := category.NewRepository(db)
	txRepo, _ := transaction.NewRepository(db)
	userRepo, _ := user.NewRepository(db)

	ctr := &container.Services{}
	ctr.Account = account.NewService(accRepo)
	ctr.Category = category.NewService(catRepo)
	ctr.Transaction = transaction.NewService(txRepo)
	ctr.User = user.NewService(userRepo)

	r := router.New(ctr)

	log.Printf("accepting connections on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}

}

func env(key, fallbackValue string) string {
	s := os.Getenv(key)
	if s == "" {
		return fallbackValue
	}

	return s
}

func dbMigration(db *sql.DB) error {

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not get db driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/db/migrations",
		"expenses",
		driver)
	log.Println("Starting migration")
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Executing migration")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println(err)
		return err
	}

	return nil
}
