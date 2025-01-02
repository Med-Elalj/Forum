package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"forum/database"
)

func PostReaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorJs(w, http.StatusMethodNotAllowed, errors.New("invalid method"))
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		ErrorJs(w, http.StatusBadRequest, errors.New(r.Header.Get("Content-Type")))
		return
	}
	UserId, err := CheckAuthentication(w, r)
	fmt.Println("==UserId==", UserId == 0)
	if err != nil || UserId == 0 {
		fmt.Println("unauthorized")
		ErrorJs(w, http.StatusUnauthorized, errors.New("unauthorized "))
		return
	}

	var requestData struct {
		PostID string `json:"postId"`
		Type   string `json:"type"`
		Post   bool   `json:"post"`
	}
	err = json.NewDecoder(r.Body).Decode(&requestData)
	fmt.Println("====>", requestData)
	if err != nil {
		fmt.Println("invalid jsonXV")
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid jsonX"))
		return
	}
	if requestData.Type != "like" && requestData.Type != "dislike" {
		fmt.Println("invalid type", requestData.Type)
		ErrorJs(w, http.StatusBadRequest, errors.New("invalid type"))
		return
	}
	PostIdInt, err := strconv.Atoi(requestData.PostID)
	if err != nil {
		fmt.Println("invalid like post id")
		ErrorJs(w, 400, errors.New("invalid like post id"))
		return
	}
	if PostIdInt < 0 {
		fmt.Println("invalid like post id")
		ErrorJs(w, 400, errors.New("invalid like post id"))
	}
	liked, err := database.HasUserLikedPost(DB, UserId, PostIdInt, requestData.Post)
	if err != nil {
		fmt.Println("error checking if user has liked post")
		ErrorJs(w, http.StatusInternalServerError, errors.New("error checking if user has liked post"))
		return
	}

	dislike, err := database.HasUserDislikedPost(DB, UserId, PostIdInt, requestData.Post)
	if err != nil {
		fmt.Println("error checking if user has liked post")
		ErrorJs(w, http.StatusInternalServerError, errors.New("error checking if user has liked post"))
		return
	}
	// /// // / / / / / / /
	addeddStatus := false
	if requestData.Type == "like" {
		fmt.Println("******************", liked, dislike)
		if dislike {
			err = database.UndislikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error unliking post")
				ErrorJs(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		}
		if liked {
			// remove the like from the post in database
			err = database.UnlikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error unliking post2")
				ErrorJs(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		} else {
			err = database.LikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error liking post3", requestData.Post)
				ErrorJs(w, http.StatusInternalServerError, errors.New("error liking post"))
				return
			}
			addeddStatus = true
		}
	} else {
		if liked {
			err = database.UnlikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error unliking post3")
				ErrorJs(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		}
		if dislike {
			// remove the like from the post in database
			err = database.UndislikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error unliking post4")
				ErrorJs(w, http.StatusInternalServerError, errors.New("error unliking post"))
				return
			}
		} else {
			err = database.DislikePost(DB, UserId, PostIdInt, requestData.Post)
			if err != nil {
				fmt.Println("error liking post2")
				ErrorJs(w, http.StatusInternalServerError, errors.New("error liking post"))
				return
			}
			addeddStatus = true
		}
	}

	// get the new like count
	likeCount, err := database.GetPostLikeCount(DB, PostIdInt, requestData.Post)
	if err != nil {
		fmt.Println("error getting like count")
		ErrorJs(w, http.StatusInternalServerError, errors.New("error getting like count"))
		return
	}
	dislikeCount, err := database.GetPostDislikeCount(DB, PostIdInt, requestData.Post)
	if err != nil {
		fmt.Println("error getting dislike count")

		ErrorJs(w, http.StatusInternalServerError, errors.New("error getting like count"))
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
