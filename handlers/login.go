package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"forum/database"
	tokening "forum/handlers/token"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	uname := r.Form.Get("username")
	upass := r.Form.Get("password")
	fmt.Println(uname, upass)
	hpassword, err := database.GetUserByUname(DB, uname)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hpassword), []byte(upass))
	if err != nil {
		log.Println(errors.New("invalid email or password"))
		ErrorPage(w, http.StatusInternalServerError, errors.New("invalid email or password"))
		return
	}

	token, err := tokening.GenerateSessionToken("username:" + uname)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	json.NewEncoder(w).Encode(struct{ Token string }{Token: token})
}

var DB *sql.DB

func ErrorPage(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
