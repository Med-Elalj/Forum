package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"forum/database"
	"forum/structs"
)

func Post(w http.ResponseWriter, r *http.Request) {
	// TODO Single add user from token
	posts, err := database.QuerryLatestPosts(DB, structs.NotaUser, 10)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, err)
		return
	}

	// Set the content type header to application/json
	w.Header().Add("Content-Type", "application/json")

	// Optionally set the status code to 200 OK
	w.WriteHeader(http.StatusOK)

	err1 := json.NewEncoder(w).Encode(struct{ Posts []structs.Post }{posts})
	fmt.Println(err1)
}
