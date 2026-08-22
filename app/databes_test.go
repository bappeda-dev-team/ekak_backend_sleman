package app

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func TestDatabase(t *testing.T) {

	godotenv.Load()

	// !! diganti dengan DB_URL agar bisa konek di docker-compose
	// contoh: "root@tcp(localhost:3306)/db_spbe?parseTime=true"
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")

	connStr := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?parseTime=true", dbUser, dbPass, dbPort, dbName)
	// connStr := os.Getenv("DB_URL")

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	defer db.Close()
}
