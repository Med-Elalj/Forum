package handlers

import (
	"encoding/json"
	"errors"
	"forum/database"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorJs(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorJs(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}

	UserId := CheckAuthentication(w, r)
	UserProfile, err := database.GetUserProfile(DB, UserId)
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized UserProfile"+err.Error()))
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
	if len(data.Title) > 60 {
		ErrorJs(w, http.StatusBadRequest, errors.New("title is too long"))
		return
	}
	if len(data.Content) > 1000 {
		ErrorJs(w, http.StatusBadRequest, errors.New("content is too long"))
		return
	}
	if len(data.Categories) == 0 {
		ErrorJs(w, http.StatusBadRequest, errors.New("no categories"))
		return
	}
	err, id := database.CreatePost(DB, UserId, data.Title, data.Content, data.Categories)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("error creating post"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        "ok",
		"ID":            id,
		"Title":         data.Title,           // TODO get username
		"UserName":      UserProfile.UserName, // TODO get username
		"CreatedAt":     "now",
		"Content":       data.Content,
		"Categories":    data.Categories,
		"LikeCount":     0,
		"DislikeCount":  0,
		"CommentsCount": 0,
	})
}
