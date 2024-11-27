// Connect to database
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sqlx.DB

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
	DB, err = sqlx.Open("postgres", connString)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal("Failed to connect to the database: ", pingErr)
	}
	log.Println("Connected to database!")
}
