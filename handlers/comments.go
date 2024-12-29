package handlers

import (
	"encoding/json"
	"errors"
	"forum/database"
	"net/http"
	"strconv"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorPage(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorPage(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}
	session, err := r.Cookie("session")
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	UserId, err := database.GetUidFromToken(DB, session.Value)
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	UserProfile, err := database.GetUserProfile(DB, UserId)

	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	data := struct {
		PostID  string
		Comment string
	}{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		ErrorPage(w, http.StatusBadRequest, errors.New("invalid json"))
		return
	}
	IdInt, err := strconv.Atoi(data.PostID)
	if err != nil {
		ErrorPage(w, http.StatusBadRequest, errors.New("invalid post id"))
		return
	}
	err, id := database.CreateComment(DB, UserId, IdInt, data.Comment)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, errors.New("error creating comment"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"UserName":  UserProfile.UserName, // TODO get username
		"CreatedAt": "now",
		"CommentID": id,
		"Content":   data.Comment,
	})

}
