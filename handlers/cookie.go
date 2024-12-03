package handlers

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, sessioID string, name string, only bool) {
	cookie := http.Cookie{
		Name:     name,
		Value:    sessioID,
		Expires:  time.Now().Add(time.Hour * 1),
		HttpOnly: only,
	}
	http.SetCookie(w, &cookie)
}

func DeleteAllCookie(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		http.SetCookie(w, &http.Cookie{
			Name:     cookie.Name,
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
		})
	}
}
