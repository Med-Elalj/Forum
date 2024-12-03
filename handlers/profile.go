package handlers

import (
	"encoding/json"
	"net/http"

	"forum/database"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	uname := r.PathValue("username")
	profile, err := database.GetUserProfile(DB, uname)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
