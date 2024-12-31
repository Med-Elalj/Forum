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
		Title          string
		Content        string
		Categories     []string
		CategoriesList []string
	}{}
	err = json.NewDecoder(r.Body).Decode(&data)
	fmt.Println("data===>", data)
	if (!title_RGX.MatchString(data.Title)) || (!content_RGX.MatchString(data.Content)) || (len(data.Categories) == 0) || (err != nil) {
		fmt.Println("some required input not provided")
		ErrorJs(w, http.StatusBadRequest, errors.New("required input not provided"))
		return
	}

	id, err := database.CreatePost(DB, UserId, data.Title, data.Content, data.Categories)
	if err != nil {
		ErrorJs(w, http.StatusInternalServerError, errors.New("Somethign went wrong creating post"+err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println(data.CategoriesList)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        "ok",
		"ID":            id,
		"Title":         data.Title,           // TODO get username
		"UserName":      UserProfile.UserName, // TODO get username
		"CreatedAt":     "now",
		"Content":       data.Content,
		"Categories":    data.CategoriesList,
		"LikeCount":     0,
		"DislikeCount":  0,
		"CommentsCount": 0,
	})
}
