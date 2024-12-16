package handlers

import (
	"html/template"
	"net/http"

	"forum/database"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	posts, err := database.QuerryLatestPosts(DB, 0, 11)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, err)
		return
	}
	tmpl, err := template.ParseFiles("frontend/templates/index.html")
	if err != nil {
		panic(err)
	}
	// Render the template with post data
	err = tmpl.Execute(w, map[string]interface{}{
		"Posts": posts,
	})
	if err != nil {
		panic(err)
	}
}
