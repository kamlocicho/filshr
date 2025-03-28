package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// TODO: move those variables to dotenv file when going production-ready
const (
	dbUser     = "myuser"
	dbPassword = "mypassword"
	dbName     = "mydb"
	dbHost     = "localhost"
	dbPort     = 5432
)

func ConnectDatabase() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	return db
}
