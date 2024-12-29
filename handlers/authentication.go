package handlers

import (
	"errors"
	"fmt"
	"forum/database"
	"net/http"
)

func CheckAuthentication(w http.ResponseWriter, r *http.Request) (userID int) {
	c, err := r.Cookie("session")
	if err != nil && err.Error() != "http: named cookie not present" {
		ErrorPage(w,"error.html", http.StatusUnauthorized, errors.New("unauthorized"+err.Error()))
		fmt.Println(err)
		return
	}
	if err != nil {
		c = &http.Cookie{}
	}

	userID, err = database.GetUidFromToken(DB, c.Value)
	if err != nil {
		ErrorPage(w, "error.html", http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	return
}
