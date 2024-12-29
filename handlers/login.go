package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
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

func Logout(w http.ResponseWriter, r *http.Request) {
	DeleteAllCookie(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// TODO fix db and link
	// uname := r.Form.Get("username")
	email := strings.ToLower(r.Form.Get("email"))
	upass := r.Form.Get("password")
	fmt.Println(email, upass)

	if email == "" {
		ErrorPage(w, "register.html", http.StatusBadRequest, errors.New("username or email is required"))
		return
	}

	var hpassword string
	var uid int

	var err error
	if email_RGX.MatchString(email) {
		hpassword, uid, err = database.GetUserByUemail(DB, email)
	} else if username_RGX.MatchString(email) {
		hpassword, uid, err = database.GetUserByUname(DB, email)
	} else {
		ErrorPage(w, "register.html", http.StatusBadRequest, errors.New("invalid email or username"))
		return
	}
	if err != nil || uid == 0 {
		log.Println(err)
		ErrorPage(w, "register.html", http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hpassword), []byte(upass))
	if err != nil {
		log.Println(errors.New("invalid email or password"))
		ErrorPage(w, "register.html", http.StatusInternalServerError, errors.New("invalid email or password"))
		return
	}

	token, err := tokening.GenerateSessionToken("email:" + email)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "register.html", http.StatusInternalServerError, err)
		return
	}
	err = database.AddSessionToken(DB, uid, token)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "register.html", http.StatusInternalServerError, err)
		return
	}
	SetCookie(w, token, "session", true)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	uemail := strings.ToLower(r.Form.Get("email"))
	uname := strings.ToLower(r.Form.Get("username"))
	upass := r.Form.Get("password")
	if !email_RGX.MatchString(uemail) || !username_RGX.MatchString(uname) || !validpassword(upass) {
		fmt.Println(!email_RGX.MatchString(uemail), uemail)
		fmt.Println(!email_RGX.MatchString(uname))
		fmt.Println(!validpassword(upass))
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid email or username or password"))
		return
	}

	uid, err := database.CreateUser(DB, uemail, uname, upass)
	if err != nil {
		log.Println(err)
		ErrorJs(w, http.StatusInternalServerError, err)
		return
	}
	token, err := tokening.GenerateSessionToken("username:" + uname)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "register.html", http.StatusInternalServerError, err)
		return
	}
	err = database.AddSessionToken(DB, uid, token)
	if err != nil {
		log.Println(err)
		ErrorPage(w, "register.html", http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	http.Redirect(w, r, "../", http.StatusFound)
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
		} else if !s && !unicode.IsDigit(char) && !unicode.IsLetter(char) {
			s = true
			continue
		}
		if a && A && d && s {
			return true
		}
	}
	return a && A && d && s
}
