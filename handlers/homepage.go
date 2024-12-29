package handlers

import (
	"errors"
	"forum/structs"
	"log"
	"net/http"

	"forum/database"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w,"error.html", http.StatusNotFound, errors.New("page not found"))
		return
	}

	template := getHtmlTemplate()

	userId := CheckAuthentication(w, r)

	var profile structs.Profile
	var err error
	if userId != 0 {
		profile, err = database.GetUserProfile(DB, userId)
		if err != nil {
			log.Fatal(err)
		}
	}

	categories, err := database.GetCategoriesWithPostCount(DB)
	if err != nil {
		ErrorPage(w, "error.html", http.StatusInternalServerError, errors.New("error getting categories from database"))
	}
	if r.FormValue("type") != "" { //?type=home
		profile.CurrentPage = r.FormValue("type")
		profile.Category = r.FormValue("category")
	} else {
		profile.CurrentPage = "home"
	}
	categor := structs.Categories(categories)

	err = template.ExecuteTemplate(w, "index.html", struct {
		Posts      []structs.Post
		Profile    structs.Profile
		Categories structs.Categories
	}{nil, profile, categor})
	if err != nil {
		// TODO remove fatal so the server doesn't stop
		log.Fatal(err)
	}
}
