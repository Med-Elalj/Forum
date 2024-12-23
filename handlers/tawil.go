package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

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
	template.ExecuteTemplate(w, "register.html", nil)
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
	data, err := database.QuerryLatestPosts(DB, 0, 5)
	if err != nil {
		log.Fatal(err)
	}
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
	posts, err := database.QuerryLatestPosts(DB, 0, 1)
	if err != nil {
		log.Fatal(err)
	}
	comments, err := database.GetCommentsByPost(DB, posts[1].ID, 0)
	if err != nil {
		log.Fatal(err)
	}

	template.ExecuteTemplate(w, "post.html", struct {
		Post     structs.Post
		Comments []structs.Comment
	}{Post: posts[1], Comments: comments})
}
