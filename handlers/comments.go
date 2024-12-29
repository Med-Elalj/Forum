package handlers

import (
	"encoding/json"
	"errors"
	"forum/database"
	"net/http"
	"strconv"
	"strings"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorJs(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
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
	UserProfile, err := database.GetUserProfile(DB, UserId)

	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	data := struct {
		PostID  string
		Comment string
	}{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if len(data.Comment) == 0 || strings.TrimSpace(data.Comment) == "" {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid comment"))
		return
	}
	if err != nil {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid json"))
		return
	}
	IdInt, err := strconv.Atoi(data.PostID)
	if err != nil {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid post id"))
		return
	}
	err, id := database.CreateComment(DB, UserId, IdInt, data.Comment)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error creating comment"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"UserName":  UserProfile.UserName,
		"CreatedAt": "just now",
		"CommentID": id,
		"Content":   data.Comment,
	})

}
