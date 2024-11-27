package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"forum/database"
	"forum/handlers"
)

func main() {
	fmt.Println(time.Now().Add(time.Hour * 24).Unix())
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
	// if err != nil {
	// 	panic(err.Error())
	// }
	handlers.DB = db
	fmt.Println("User created successfully!")
	fmt.Println(database.QuerryPostsbyUser(db, "username", 5))
	http.HandleFunc("/login", handlers.Login)
	http.Handle("/", http.FileServer(http.Dir("templates/")))
	// http.HandleFunc("/register", handlers.Register)
	fmt.Println("Server listening on :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
