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
	if err := r.ParseForm(); err != nil {
		ErrorPage(w, http.StatusBadRequest, errors.New("failed to parse Loging form " + err.Error()))
		return
	}
	// TODO fix db and link
	uname := r.Form.Get("username")
	email := r.Form.Get("email")
	upass := r.Form.Get("password")
	fmt.Println(uname, email, upass)
	if email == "" && uname == "" {
		ErrorJs(w, http.StatusBadRequest, errors.New("username or email is required"))
		return
	}

	var hpassword string
	var uid int

	var err error
	if email != "" && email_RGX.MatchString(email) {
		hpassword, uid, err = database.GetUserByUemail(DB, email)
	} else if uname != "" && username_RGX.MatchString(uname) {
		hpassword, uid, err = database.GetUserByUname(DB, uname)
	} else {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid email or username"))
		return
	}
	if err != nil {
		log.Println(err)
		ErrorJs(w, http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hpassword), []byte(upass))
	if err != nil {
		ErrorJs(w, http.StatusUnauthorized, errors.New("invalid email or password " + err.Error()))
		return
	}

	token, err := tokening.GenerateSessionToken("username:" + uname)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, errors.New("Error in Generating Session Token " + err.Error()))
		return
	}

	err = database.AddSessionToken(DB, uid, token)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, errors.New("Error in Adding Session Token to DB " + err.Error()))
		return
	}
	SetCookie(w, token, "session", true)
	http.Redirect(w, r, "/index", http.StatusFound)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseForm(); err != nil {
		ErrorJs(w, http.StatusBadRequest, errors.New("failed to parse Register form"))
		return
	}

	uemail := r.Form.Get("email")
	uname := r.Form.Get("username")
	upass := r.Form.Get("password")
	if !email_RGX.MatchString(uemail) || !username_RGX.MatchString(uname) || !validpassword(upass) {
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid email or username or password"))
		return
	}

	uid, err := database.CreateUser(DB, uemail, uname, upass)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError, errors.New("Error in Creating User " + err.Error()))
		return
	}

	token, err := tokening.GenerateSessionToken("username:" + uname)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError,  errors.New("Error in Generating Session Token " + err.Error()))
		return
	}

	err = database.AddSessionToken(DB, uid, token)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError, errors.New("Error in Adding Session Token to DB " + err.Error()))
		return
	}
	json.NewEncoder(w).Encode(struct{ Token string }{token})

	w.WriteHeader(http.StatusCreated)
	
}

func validpassword(password string) bool {
	// Lowercase UPPERCASE digit {symbol}
	var a, A, d, s bool
	if len(password) < 8 || len(password) > 64 {
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
		} else if !s && unicode.IsSymbol(char) {
			s = true
			continue
		}
		if a && A && d && s {
			return true
		}
	}
	return a && A && d && s
}
