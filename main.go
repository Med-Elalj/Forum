package main

import (
	"fmt"
	"os"
	"time"

	"forum/database"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run setup.go <database_file>")
		os.Exit(1)
	}
	fmt.Printf("Creating database at %v...\n", os.Args[1:])
	db, err := database.OpenDatabase(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	database.CreateTables(db)
	fmt.Println("Database setup complete!")
	ticker := time.NewTicker(time.Hour)
	go database.DES_Ticker(ticker, db)
	defer ticker.Stop()
	defer db.Close()
	// err = database.CreateUser(db, "exazerample@website.net", "userqsdfname", "password")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("User created successfully!")
	fmt.Println(database.QuerryPostsbyUser(db, "username", 5))
}
