// Connect to database
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func InitDB() {
	// Capture connection properties.
	var connString string = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	)

	var err error
	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal("Failed to connect to the database: ", pingErr)
	}
	log.Println("Connected to database!")
}
