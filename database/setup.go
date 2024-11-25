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

// var main sql.DB
var tables = map[string]string{
	"users": `CREATE TABLE IF NOT EXISTS users (
	  id INTEGER PRIMARY KEY,
	  email TEXT UNIQUE NOT NULL,
	  username TEXT UNIQUE NOT NULL,
	  password TEXT NOT NULL,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`,

	"posts": `CREATE TABLE IF NOT EXISTS posts (
	  id INTEGER PRIMARY KEY,
	  user_id INTEGER,
	  content TEXT NOT NULL,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`,

	"comments": `CREATE TABLE IF NOT EXISTS comments (
	  id INTEGER PRIMARY KEY,
	  post_id INTEGER,
	  user_id INTEGER,
	  content TEXT NOT NULL,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`,

	"categories": `CREATE TABLE IF NOT EXISTS categories (
	  id INTEGER PRIMARY KEY,
	  name TEXT UNIQUE NOT NULL
	);`,

	"post_categories": `CREATE TABLE IF NOT EXISTS post_categories (
	  post_id INTEGER,
	  category_id INTEGER,
	  PRIMARY KEY (post_id, category_id),
	  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	  FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
	);`,

	"likes": `CREATE TABLE IF NOT EXISTS likes (
	id INTEGER PRIMARY KEY,
	user_id INTEGER,
	post_id INTEGER,
	comment_id INTEGER,
	is_like BOOLEAN NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
	CHECK (
	(post_id IS NOT NULL AND comment_id IS NULL) OR 
	(post_id IS NULL AND comment_id IS NOT NULL)
  )
	);`,

	"sessions": `CREATE TABLE IF NOT EXISTS sessions (
	  id INTEGER PRIMARY KEY,
	  user_id INTEGER,
	  session_token TEXT NOT NULL,
	  expiration TIMESTAMP NOT NULL,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`,
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
