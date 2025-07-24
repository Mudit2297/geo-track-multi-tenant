package client

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	user := os.Getenv("POSTGRES_USER")
	pw := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	if user == "" || pw == "" || dbName == "" || host == "" || port == "" {
		log.Fatalf("DB connection details not provided. Connection failed")
	}

	conn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, pw, host, port, dbName)

	// conn := "postgres://root:password@postgres:5432/postgres?sslmode=disable"

	var err error

	// Retry logic for container startup order
	for i := 0; i < 5; i++ {
		DB, err = sql.Open("postgres", conn)
		if err == nil && DB.Ping() == nil {
			break
		}
		log.Printf("Retrying DB connection attempt %d: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL")
}
