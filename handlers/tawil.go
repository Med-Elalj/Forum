package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"forum/database"
	"forum/structs"
)

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
	fmt.Println(c.Name, c.Value, err)
	uid, err := database.GetUidFromToken(DB, c.Value)
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	// TODO LIMIT
	data, err := database.QuerryLatestPosts(DB, uid, structs.Limit)
	if err != nil {
		log.Fatal(err)
	}
	// TODO get user profile
	cat := map[string]int{"test": 1, "azer": 32}
	profile := structs.Profile{
		UserName:     "test",
		PFP:          "Vivian",
		ArticleCount: 2,
		CommentCount: 20,
		Categories:   cat,
	}
	fmt.Println(profile)
	template.ExecuteTemplate(w, "index.html", struct {
		Posts   []structs.Post
		Profile structs.Profile
	}{data, profile})
}

func TawilProfileHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseGlob("./frontend/templates/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
	template, err = template.ParseGlob("./frontend/templates/components/*.html")
	if err != nil {
		log.Fatal(err, "Error Parsing Data from Template hTl")
	}
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
	i, err := fmt.Sscanf(r.URL.Path, "/post1/%d", &Postid)
	if i != 1 || err != nil || Postid <= 0 {
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
