package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"forum/database"
	"forum/structs"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
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
	Postid, err = strconv.Atoi(r.PathValue("id"))
	if err != nil || Postid < 0 {
		// TODO iso standard
		ErrorPage(w, 400, errors.New("invalid TawilPostHandler 0"))
	}
	post, err := database.GetPostByID(DB, Postid, uid)
	if err != nil {
		// TODO iso standard
		ErrorPage(w, 400, errors.New("invalid TawilPostHandler 1"))
	}

	comments, err := database.GetCommentsByPost(DB, post.ID)
	fmt.Println(comments)
	if err != nil {
		// TODO iso standard
		ErrorPage(w, 400, errors.New("invalid TawilPostHandler 2"))
	}

	template.ExecuteTemplate(w, "post.html", struct {
		Post     structs.Post
		Comments []structs.Comment
	}{Post: post, Comments: comments})
}

//localhost:8080/infinite-scroll?offset=0
func InfiniteScroll(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil && err.Error() != "http: named cookie not present" {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized"+err.Error()))
		fmt.Println(err)
		return
	}
	if err != nil {
		c = &http.Cookie{}
	}
	// adding funtcion to check the url type is it Cetegory type or home type of profile type
	// if it is category type then we will get the category name from the url and then we will get the posts by category
	// if it is profile type then we will get the username from the url and then we will get the posts by username
	// if it is home type then we will get the latest posts
	// if it is none of the above then we will return the error
	uid, err := database.GetUidFromToken(DB, c.Value)
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
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

	offset_str := r.URL.Query().Get("offset")
	offset, err := strconv.Atoi(offset_str)
	if err != nil {
		offset = 0
	}

	var posts []structs.Post
	if r.URL.Query().Get("type") == "category" {
		category := r.URL.Query().Get("category")
		posts, err = database.QuerryLatestPostsByCategory(DB, uid, category, offset)
		if err != nil {
			ErrorJs(w, http.StatusInternalServerError, err)
			return
		}
	} else if r.URL.Query().Get("type") == "profile" {
		posts, err = database.QuerryPostsbyUser(DB, profile.UserName, uid, structs.Limit)
		if err != nil {
			ErrorPage(w, http.StatusInternalServerError, errors.New("error fetching posts "+err.Error()))
			return
		}
	} else if r.URL.Query().Get("type") == "home" {
		posts, err = database.QuerryLatestPosts(DB, uid, structs.Limit, offset)
		if err != nil {
			ErrorJs(w, http.StatusInternalServerError, err)
			return
		}
	} else if r.URL.Query().Get("type") == "liked" {
		posts, err = database.QuerryLatestPostsByUserLikes(DB, profile.UserName, uid, structs.Limit)
		if err != nil {
			ErrorJs(w, http.StatusInternalServerError, err)
			return
		}
	} else if r.URL.Query().Get("type") == "trending" {
		posts, err = database.QuerryLatestPosts(DB, uid, structs.Limit, offset)
		if err != nil {
			ErrorJs(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		ErrorPage(w, http.StatusBadRequest, errors.New("invalid url"))
		return
	}
	fmt.Println("Offset:", offset)

	// Set the content type header to application/json
	w.Header().Add("Content-Type", "application/json")

	// Optionally set the status code to 200 OK
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(posts)
	fmt.Println(err)
}
