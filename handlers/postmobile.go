package handlers

import (
	"errors"
	"fmt"
	"forum/database"
	"forum/structs"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func TawilPostMobile(w http.ResponseWriter, r *http.Request) {

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

	Postid, err = strconv.Atoi(r.FormValue("id"))
	if err != nil || Postid < 0 {
		// TODO iso standard
		ErrorPage(w, "error.html", 400, errors.New("invalid TawilPostHandler 0"))
	}
	post, err := database.GetPostByID(DB, Postid, 0)
	if err != nil {
		// TODO iso standard
		ErrorPage(w, "error.html", 400, errors.New("invalid TawilPostHandler 1"))
	}

	comments, err := database.GetCommentsByPost(DB, post.ID)
	fmt.Println(comments)
	if err != nil {
		// TODO iso standard
		ErrorPage(w, "error.html", 400, errors.New("invalid TawilPostHandler 2"))
	}

	template.ExecuteTemplate(w, "postmobile.html", struct {
		Post     structs.Post
		Comments []structs.Comment
	}{Post: post, Comments: comments})
}
