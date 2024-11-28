package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase(file string) (*sql.DB, error) {
	fmt.Println(file)
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the database is actually accessible
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database opened successfully!")
	return db, nil
}

func CreateTables(db *sql.DB) {
	for t, c := range tables {
		_, err := db.Exec(c)
		if err != nil {
			log.Fatalf("Error creating table %s: %v", t, err)
		}
		fmt.Printf("Created table: %s\n", t)
	}
	DeleteExpiredSessions(db)
}

func CreateTriggers(db *sql.DB) {
	for t, c := range trigers {
		rc, err := db.Exec(c)
		if err != nil && err.Error() == "Error creating table post_year_enforce: trigger "+t+" already exists" {
			log.Fatalf("Error creating table %s: %v, '%v'", t, err, rc)
		}
		fmt.Printf("Created trigger: %s\n", t)
	}
}

func DES_Ticker(ticker *time.Ticker, db *sql.DB) {
	for range ticker.C {
		err := DeleteExpiredSessions(db)
		if err != nil {
			log.Printf("Error deleting expired sessions: %v", err)
		} else {
			fmt.Println("Expired sessions deleted successfully.")
		}
	}
}

func DeleteExpiredSessions(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM sessions WHERE expiration < CURRENT_TIMESTAMP")
	return err
}
