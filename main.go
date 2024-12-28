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
		fmt.Println("Usage: go run main.go <database_file>")
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

	handlers.DB = db
	fmt.Println("User created successfully!")
	fmt.Println(database.QuerryPostsbyUser(db, "test", 0, 5))

	http.HandleFunc("/", handlers.TawilHandelr)

	http.HandleFunc("POST /login", handlers.Login)
	http.HandleFunc("POST /register", handlers.Register)
	http.HandleFunc("/login", handlers.TawilHandelrRegister)
	http.HandleFunc("/register", handlers.TawilHandelrRegister)
	http.HandleFunc("/logout", handlers.Logout)

	http.HandleFunc("/post/{id}", handlers.TawilPostHandler)

	http.HandleFunc("/CreateComment", handlers.AddCommentHandler)
	http.HandleFunc("/createPost", handlers.CreatePost)
	http.HandleFunc("/PostReaction", handlers.PostReaction)

	http.HandleFunc("/infinite-scroll", handlers.InfiniteScroll)
	http.HandleFunc("/profile/{username}", handlers.TawilProfileHandler)
	// http.HandleFunc("/profile/{username}/likes", handlers.Likes) // TODO Implement profile like handler

	// http.HandleFunc("/create", handlers.Create) // TODO Implement create

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./frontend/assets"))))

	fmt.Println("Server listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
