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
	data, err := database.QuerryLatestPosts(DB, uid, structs.Limit, structs.NoOffSet)
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

	categories, err := database.GetCategoriesWithPostCount(DB)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, errors.New("error getting categories from database"))
	}
	// TODO PROFILE PICTURES
	profile.PFP = "Vivian"

	categor := structs.Categories(categories)

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
	categories, err := database.GetCategoriesWithPostCount(DB)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, errors.New("error getting categories from database"))
	}
	categor := structs.Categories(categories)
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorPage(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorPage(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}
	token, err := r.Cookie("session")
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	UserId, err := database.GetUidFromToken(DB, token.Value)
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized ID"+err.Error()))
		return
	}
	UserProfile, err := database.GetUserProfile(DB, UserId)

	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized UserProfile"+err.Error()))
		return
	}
	data := struct {
		Title      string
		Content    string
		Categories []string
	}{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		ErrorPage(w, http.StatusBadRequest, errors.New("invalid json"))
		return
	}
	if len(data.Title) > 60 {
		ErrorPage(w, http.StatusBadRequest, errors.New("title is too long"))
		return
	}
	if len(data.Content) > 1000 {
		ErrorPage(w, http.StatusBadRequest, errors.New("content is too long"))
		return
	}
	if len(data.Categories) == 0 {
		ErrorPage(w, http.StatusBadRequest, errors.New("no categories"))
		return
	}
	err, id := database.CreatePost(DB, UserId, data.Title, data.Content, data.Categories)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, errors.New("error creating post"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        "ok",
		"ID":            id,
		"Title":         data.Title,           // TODO get username
		"UserName":      UserProfile.UserName, // TODO get username
		"CreatedAt":     "now",
		"Content":       data.Content,
		"Categories":    data.Categories,
		"LikeCount":     0,
		"DislikeCount":  0,
		"CommentsCount": 0,
	})
}

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorPage(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorPage(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}
	session, err := r.Cookie("session")
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	UserId, err := database.GetUidFromToken(DB, session.Value)
	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	UserProfile, err := database.GetUserProfile(DB, UserId)

	if err != nil {
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	data := struct {
		PostID  string
		Comment string
	}{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		ErrorPage(w, http.StatusBadRequest, errors.New("invalid json"))
		return
	}
	IdInt, err := strconv.Atoi(data.PostID)
	if err != nil {
		ErrorPage(w, http.StatusBadRequest, errors.New("invalid post id"))
		return
	}
	err, id := database.CreateComment(DB, UserId, IdInt, data.Comment)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, errors.New("error creating comment"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "ok",
		"UserName":  UserProfile.UserName, // TODO get username
		"CreatedAt": "now",
		"CommentID": id,
		"Content":   data.Comment,
	})

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

func Logout(w http.ResponseWriter, r *http.Request) {
	DeleteAllCookie(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}

// handling likes and dislikes in the backend Golang
func PostReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorPage(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorPage(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}
	session, err := r.Cookie("session")
	if err != nil {
		fmt.Println("unauthorized")
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}
	UserId, err := database.GetUidFromToken(DB, session.Value)
	if err != nil {
		fmt.Println("unauthorized")
		ErrorPage(w, http.StatusUnauthorized, errors.New("unauthorized "+err.Error()))
		return
	}

	var requestData struct {
		PostID string `json:"postId"`
		Type   string `json:"type"`
		Post   bool   `json:"post"`
	}

	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		fmt.Println("invalid jsonXV")
		ErrorPage(w, http.StatusBadRequest, errors.New("invalid jsonX"))
		return
	}
	if requestData.Type != "like" && requestData.Type != "dislike" {
		fmt.Println("invalid type", requestData.Type)
		ErrorPage(w, http.StatusBadRequest, errors.New("invalid type"))
		return
	}
	PostIdInt, err := strconv.Atoi(requestData.PostID)
	if err != nil {
		fmt.Println("invalid like post id")
		ErrorPage(w, 400, errors.New("invalid like post id"))
		return
	}
	if PostIdInt < 0 {
		fmt.Println("invalid like post id")
		ErrorPage(w, 400, errors.New("invalid like post id"))
	}
	liked, err := database.HasUserLikedPost(DB, UserId, PostIdInt, requestData.Post)
	if err != nil {
		fmt.Println("error checking if user has liked post")
		ErrorPage(w, http.StatusInternalServerError, errors.New("error checking if user has liked post"))
		return
	}

	dislike, err := database.HasUserDislikedPost(DB, UserId, PostIdInt, requestData.Post)
	if err != nil {
		fmt.Println("error checking if user has liked post")
		ErrorPage(w, http.StatusInternalServerError, errors.New("error checking if user has liked post"))
		return
	}
	// /// // / / / / / / /
	addeddStatus := false
	if requestData.Type == "like" {
		if dislike {
			err = database.UndislikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error unliking post")
				ErrorPage(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		}
		if liked {
			// remove the like from the post in database
			err = database.UnlikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error unliking post2")
				ErrorPage(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		} else {
			err = database.LikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error liking post3")
				ErrorPage(w, http.StatusInternalServerError, errors.New("error liking post"))
				return
			}
			addeddStatus = true
		}
	} else {
		if liked {
			err = database.UnlikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error unliking post3")
				ErrorPage(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		}
		if dislike {
			// remove the like from the post in database
			err = database.UndislikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error unliking post4")
				ErrorPage(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		} else {
			err = database.DislikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error liking post2")
				ErrorPage(w, http.StatusInternalServerError, errors.New("error liking post"))
				return
			}
			addeddStatus = true
		}
	}

	// get the new like count
	likeCount, err := database.GetPostLikeCount(DB, PostIdInt, requestData.Post)
	if err != nil {
		fmt.Println("error getting like count")
		ErrorPage(w, http.StatusInternalServerError, errors.New("error getting like count"))
		return
	}
	dislikeCount, err := database.GetPostDislikeCount(DB, PostIdInt, requestData.Post)
	if err != nil {
		fmt.Println("error getting dislike count")

		ErrorPage(w, http.StatusInternalServerError, errors.New("error getting like count"))
		return
	}
	database.UpdatePostLikeCount(DB, PostIdInt, requestData.Post)
	database.UpdatePostDislikeCount(DB, PostIdInt, requestData.Post)
	// return the new like count
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(map[string]interface{}{
		"status":   "ok",
		"likes":    likeCount,
		"dislikes": dislikeCount,
	})
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "ok",
		"likes":    likeCount,
		"dislikes": dislikeCount,
		"added":    addeddStatus,
	})

}

// Create Handler of Categories
