package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"forum/database"
)

/// TODO No need i have already implemented it in Tawil.go file
/// under The same NAME Remove X anothe duplicate in posts.go files under X

func CreatePostXX(w http.ResponseWriter, r *http.Request) {
	// Handling adding posts based on createPostInputEventListeners function
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorJs(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}
	token, err := r.Cookie("session")
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	UserId, err := database.GetUidFromToken(DB, token.Value)
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized ID"+err.Error()))
		return
	}

	data := struct {
		Title      string
		Content    string
		Categories []string
	}{}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid json"))
		return
	}

	Pid, err := database.CreatePost(DB, UserId, data.Title, data.Content, data.Categories)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error creating post"))
		return
	}
	post, err := database.GetPostByID(DB, Pid, UserId)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error getting post"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	// Handling adding comments based on addCommentInputEventListeners function
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorJs(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}
	session, err := r.Cookie("session")
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	UserId, err := database.GetUidFromToken(DB, session.Value)
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}

	data := struct {
		PostID  string
		Comment string
	}{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("errorX", err)
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid json"))
		return
	}

	Pid, err := strconv.Atoi(data.PostID)
	if err != nil {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid post id"))
		return
	}
	fmt.Println("User id ", UserId)
	err, id := database.CreateComment(DB, UserId, Pid, data.Comment)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error creating comment"))
		return
	}

	comment, err := database.GetCommentById(DB, id)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error getting comment"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comment)
}
