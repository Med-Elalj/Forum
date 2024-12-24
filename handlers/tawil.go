package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"forum/database"
	"forum/structs"
)

// func TawilHandelrCreate(w http.ResponseWriter, r *http.Request) {
// 	// Middler Ware to check if user is logged in
// 	if r.Method == "POST" {

// 	// if not redirect to login
// 	// if logged in continue
// 	// if post request
// 	// get data from post
// 	// create post
// 	// redirect to post
// 	// if get request
// 	// render create post page
// 	template, err := template.ParseGlob("./frontend/templates/*.html")
// 	if err != nil {
// 		log.Fatal(err, "Error Parsing Data from Template hTl")
// 	}
// 	template, err = template.ParseGlob("./frontend/templates/components/*.html")
// 	if err != nil {
// 		log.Fatal(err, "Error Parsing Data from Template hTl")
// 	}
// 	template.ExecuteTemplate(w, "create.html", nil)

// }
func TawilHandelrRegister(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseGlob("./frontend/templates/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
	template, err = template.ParseGlob("./frontend/templates/components/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
	template.ExecuteTemplate(w, "register.html", struct{ Register bool }{Register: r.URL.Path == "/register"})
}

func TawilHandelr(w http.ResponseWriter, r *http.Request) {
	template, err := template.New("index").Funcs(template.FuncMap{
		"timeAgo": structs.TimeAgo,
	}).ParseGlob("./frontend/templates/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
	template, err = template.ParseGlob("./frontend/templates/components/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}

	c, err := r.Cookie("session")
	if err != nil && err.Error() != "http: named cookie not present" {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized"+err.Error()))
		fmt.Println(err)
		return
	}
	if err != nil {
		c = &http.Cookie{}
	}

	uid, err := database.GetUidFromToken(DB, c.Value)
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	// TODO LIMIT
	data, err := database.QuerryLatestPosts(DB, uid, structs.Limit)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, errors.New("error fetching posts "+err.Error()))
		return
	}

	var profile structs.Profile
	if uid != 0 {
		profile, err = database.GetUserProfile(DB, uid)
		if err != nil {
			log.Fatal(err)
		}
	}
	// TODO PROFILE PICTURES
	profile.PFP = "Vivian"
	// TODO dynamic categories
	cat := map[string]int{"test": 1, "azer": 32}

	categor := structs.Categories(cat)
	fmt.Println(profile)

	err = template.ExecuteTemplate(w, "index.html", struct {
		Posts      []structs.Post
		Profile    structs.Profile
		Categories structs.Categories
	}{data, profile, categor})
	if err != nil {
		// TODO remove fatal so the server doesn't stop
		log.Fatal(err)
	}
}

func TawilProfileHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.New("index").Funcs(template.FuncMap{
		"timeAgo": structs.TimeAgo,
	}).ParseGlob("./frontend/templates/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
	template, err = template.ParseGlob("./frontend/templates/components/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}

	c, err := r.Cookie("session")
	if err != nil && err.Error() != "http: named cookie not present" {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized"+err.Error()))
		fmt.Println(err)
		return
	}
	if err != nil {
		c = &http.Cookie{}
	}

	uid, err := database.GetUidFromToken(DB, c.Value)
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	// TODO LIMIT
	uname := r.PathValue("username")
	if !username_RGX.MatchString(uname) {
		ErrorPage(w, http.StatusBadRequest, errors.New("invalid username"))
		fmt.Println("username atempt", uname)
		return
	}
	data, err := database.QuerryPostsbyUser(DB, uname, uid, structs.Limit)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, errors.New("error fetching posts "+err.Error()))
		return
	}
	var profile structs.Profile
	if uid != 0 {
		fmt.Println("3mer profile")
		profile, err = database.GetUserProfile(DB, uid)
		if err != nil {
			log.Fatal(err)
		}
	}
	// TODO PROFILE PICTURES
	profile.PFP = "Vivian"
	// TODO dynamic categories
	cat := map[string]int{"test": 1, "azer": 32}

	categor := structs.Categories(cat)
	fmt.Println(profile)

	err = template.ExecuteTemplate(w, "index.html", struct {
		Posts      []structs.Post
		Profile    structs.Profile
		Categories structs.Categories
	}{data, profile, categor})
	if err != nil {
		// TODO remove fatal so the server doesn't stop
		log.Fatal(err)
	}
	// TODO
	template.ExecuteTemplate(w, "profile.html", nil)
}

func TawilPostHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseGlob("./frontend/templates/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
	template, err = template.ParseGlob("./frontend/templates/components/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}

	// TODO make it post specific
	var Postid int
	// get the post id from /post/{id}
	// TODO check for edge cases

	Postid, err = strconv.Atoi(r.PathValue("id"))
	if err != nil || Postid <= 0 {
		// TODO iso standard
		ErrorPage(w, 400, errors.New("invalid TawilPostHandler 0"))
	}
	post, err := database.GetPostByID(DB, Postid, 0)
	if err != nil {
		// TODO iso standard
		ErrorPage(w, 400, errors.New("invalid TawilPostHandler 1"))
	}

	comments, err := database.GetCommentsByPost(DB, post.ID, 5)
	if err != nil {
		// TODO iso standard
		ErrorPage(w, 400, errors.New("invalid TawilPostHandler 2"))
	}

	template.ExecuteTemplate(w, "post.html", struct {
		Post     structs.Post
		Comments []structs.Comment
	}{Post: post, Comments: comments})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	DeleteAllCookie(w, r)
	http.Redirect(w, r, "/index", http.StatusFound)
}

// TODO
func Likes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome", r.URL.Path, r.PathValue("username"))
}

// TODO
func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome", r.URL.Path)
	// TODO validate and sanitize inputs
	r.ParseForm()
	content := r.FormValue("content")
	title := r.FormValue("title")
	cat := r.Form["category"]
	uid, _ := strconv.Atoi(r.FormValue("uid"))

	fmt.Println("test", content, title, uid, cat)

	http.Redirect(w, r, "/index", http.StatusFound)
}
