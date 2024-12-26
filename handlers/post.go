package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"forum/database"
	"forum/structs"
)

func Post(w http.ResponseWriter, r *http.Request) {
	// TODO Single add user from token
	posts, err := database.QuerryLatestPosts(DB, structs.NotaUser, structs.Limit, structs.NoOffSet)
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

func InfiniteScroll(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil && err.Error() != "http: named cookie not present" {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized"+err.Error()))
		fmt.Println(err)
		return
	}
	if err != nil {
		c = &http.Cookie{}
	}

	uid, err := database.GetUidFromToken(DB, c.Value)
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	offset_str := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offset_str)
	if err != nil {
		offset = 0
	}
	fmt.Println("Offset:", offset)
	posts, err := database.QuerryLatestPosts(DB, uid, structs.Limit, offset)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, err)
		return
	}
	// Set the content type header to application/json
	w.Header().Add("Content-Type", "application/json")

	// Optionally set the status code to 200 OK
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(posts)
	fmt.Println(err)
}
