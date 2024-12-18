package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"unicode"

	"forum/database"
	tokening "forum/handlers/token"

	"golang.org/x/crypto/bcrypt"
)

var (
	DB                      *sql.DB
	email_RGX, username_RGX *regexp.Regexp
)

func init() {
	email_RGX = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	username_RGX = regexp.MustCompile(`^\w{3,19}$`)
}

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
	err = database.AddSessionToken(DB, uname, token)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct{ Token string }{token})
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	uemail := r.Form.Get("email")
	uname := r.Form.Get("username")
	upass := r.Form.Get("password")
	if !email_RGX.MatchString(uemail) || !username_RGX.MatchString(uname) || !validpassword(upass) {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid email or username or password"))
		return
	}

	err := database.CreateUser(DB, uemail, uname, upass)
	if err != nil {
		log.Println(err)
		ErrorJs(w, http.StatusInternalServerError, err)
		return
	}
	token, err := tokening.GenerateSessionToken("username:" + uname)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	err = database.AddSessionToken(DB, uname, token)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct{ Token string }{token})
}

func validpassword(password string) bool {
	// Lowercase UPPERCASE digit
	var a, A, d bool
	if len(password) < 8 {
		return false
	}
	for _, char := range password {
		if !a && unicode.IsLower(char) {
			a = true
			continue
		} else if !A && unicode.IsUpper(char) {
			A = true
			continue
		} else if !d && unicode.IsDigit(char) {
			d = true
			continue
		}
		if a && A && d {
			return true
		}
	}
	return a && A && d
}
