package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"forum/database"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	uname := r.PathValue("username")
	if uname == "" {
		// TODO: handle empty username
		ErrorJs(w, http.StatusBadRequest, errors.New("username is required"))
		return
	}
	profile, err := database.GetUserProfile(DB, uname)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
