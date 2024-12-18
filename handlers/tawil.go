package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"forum/database"
	"forum/structs"
)

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
	template.ExecuteTemplate(w, "post.html", nil)
}
