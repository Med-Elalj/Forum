package handlers

import (
	"forum/database"
	"forum/structs"
	"net/http"
	"strconv"
)

func TawilPostMobile(w http.ResponseWriter, r *http.Request) {

	template := getHtmlTemplate()
	// TODO make it post specific
	var Postid int
	// get the post id from /post/{id}
	// TODO check for edge cases

	Postid, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || Postid < 0 {
		ErrorPage(w, "error.html", map[string]interface{}{
			"StatuCode":    http.StatusBadRequest,
			"MessageError": "Bar Request error, incorrect post id" + err.Error(),
		})
		return
	}
	post, err := database.GetPostByID(DB, Postid, 0)
	if err != nil {
		ErrorPage(w, "error.html", map[string]interface{}{
			"StatuCode":    http.StatusBadRequest,
			"MessageError": "Bar Request error, incorrect post id" + err.Error(),
		})
		return
	}

	comments, err := database.GetCommentsByPost(DB, post.ID)
	if err != nil {
		ErrorPage(w, "error.html", map[string]interface{}{
			"StatuCode":    http.StatusInternalServerError,
			"MessageError": "internal server error, trying to get comments" + err.Error(),
		})
		return
	}

	template.ExecuteTemplate(w, "postmobile.html", struct {
		Post     structs.Post
		Comments []structs.Comment
	}{Post: post, Comments: comments})
}
