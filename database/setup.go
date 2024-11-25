package database

import (
	"database/sql"
	"fmt"
	"log"

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
}

// var main sql.DB
var tables = map[string]string{
	"users": `CREATE TABLE users (
	  id INTEGER PRIMARY KEY,
	  email TEXT UNIQUE,
	  username TEXT UNIQUE,
	  password TEXT,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	);`,

	"posts": `CREATE TABLE posts (
	  id INTEGER PRIMARY KEY,
	  user_id INTEGER,
	  content TEXT,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`,

	"comments": `CREATE TABLE comments (
	  id INTEGER PRIMARY KEY,
	  post_id INTEGER,
	  user_id INTEGER,
	  content TEXT,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`,

	"categories": `CREATE TABLE categories (
	  id INTEGER PRIMARY KEY,
	  name TEXT UNIQUE
	);`,

	"post_categories": `CREATE TABLE post_categories (
	  post_id INTEGER,
	  category_id INTEGER,
	  PRIMARY KEY (post_id, category_id),
	  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	  FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
	);`,

	"likes": `CREATE TABLE likes (
	id INTEGER PRIMARY KEY,
	user_id INTEGER,
	post_id INTEGER,
	comment_id INTEGER,
	is_like BOOLEAN,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
	CHECK (
	(post_id IS NOT NULL AND comment_id IS NULL) OR 
	(post_id IS NULL AND comment_id IS NOT NULL)
  )
	);`,

	"sessions": `CREATE TABLE sessions (
	  id INTEGER PRIMARY KEY,
	  user_id INTEGER,
	  session_token TEXT,
	  expiration TIMESTAMP,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`,
}
