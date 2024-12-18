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
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run setup.go <database_file>")
		os.Exit(1)
	}
	fmt.Printf("Creating database at %v...\n", os.Args[1:])
	db := database.OpenDatabase(os.Args[1])
	database.CreateTables(db)
	database.CreateTriggers(db)
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
	fmt.Println(database.QuerryPostsbyUser(db, "test", 0, 5))
	http.HandleFunc("POST /login", handlers.Login)
	http.HandleFunc("POST /register", handlers.Register)
	http.HandleFunc("/post1", handlers.TawilPostHandler)
	http.HandleFunc("/profile", handlers.TawilProfileHandler)
	http.HandleFunc("/post", handlers.Post)
	http.HandleFunc("/home", handlers.HomePage)
	http.HandleFunc("/err", handlers.TawilHandelr)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./frontend/assets"))))
	fmt.Println("Server listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
