package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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

	UserId, err := CheckAuthentication(w, r)
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
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
		fmt.Println("invalid json - error 1")
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid json - error 1"))
		return
	}
	if len(data.Title) > 60 {
		fmt.Println("title is too long - error 2")
		ErrorJs(w, http.StatusBadRequest, errors.New("title is too long - error 2"))
		return
	}
	if len(data.Content) > 1000 {
		fmt.Println("content is too long - error 3")
		ErrorJs(w, http.StatusBadRequest, errors.New("content is too long - error 3"))
		return
	}
	if len(data.Categories) == 0 {
		fmt.Println("no categories - error 4")
		ErrorJs(w, http.StatusBadRequest, errors.New("no categories - error 4"))
		return
	}
	err, id := database.CreatePost(DB, UserId, data.Title, data.Content, data.Categories)
	if err != nil {
		fmt.Println("error creating post - error 5")
		ErrorJs(w, http.StatusInternalServerError, errors.New("error creating post - error 5"))
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
