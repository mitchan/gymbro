package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mitchan/gymbro/util"
)

func NewDatabase() *sql.DB {
	connStr := util.GetEnv("DATABASE_URL", "")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	log.Printf("========= DATABASE CONNECTION =========")
	log.Printf("Full Connection String: %s", connStr)
	log.Printf("=======================================")

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	return db
}
