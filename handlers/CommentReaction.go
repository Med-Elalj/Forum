package handlers

/* we Will Create Function To Like And Dislike Comments >>
We have 2 tables in database :
	"comments": `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY NOT NULL,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		like_count INTEGER DEFAULT 0,
		dislike_count INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		CONSTRAINT no_duplicates UNIQUE (post_id, user_id, content)
		);`,


			"comment_likes": `CREATE TABLE IF NOT EXISTS comment_likes (
		user_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL,
		is_like BOOLEAN NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, comment_id)
		);`,

/////////// This section Bellong to Reaction in the Comments Like And Dislike /////////////
// Crate HasUserLikedComment in database/likes.go
func HasUserLikedComment(db *sql.DB, userId, commentId int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 1`
	err := db.QueryRow(query, userId, commentId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Crate LikeComment in database/likes.go
func LikeComment(db *sql.DB, userId, commentId int) error {
	query := `INSERT INTO comment_likes (user_id, comment_id, is_like) VALUES (?, ?, 1)`
	_, err := db.Exec(query, userId, commentId)
	return err
}

// Crate UnlikeComment in database/likes.go
func UnlikeComment(db *sql.DB, userId, commentId int) error {
	query := `DELETE FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 1`
	_, err := db.Exec(query, userId, commentId)
	return err
}

// Crate GetCommentLikeCount in database/likes.go
func GetCommentLikeCount(db *sql.DB, commentId int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 1`
	err := db.QueryRow(query, commentId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Crate UpdateCommentLikeCount in database/likes.go
func UpdateCommentLikeCount(db *sql.DB, commentId int) error {
	query := `UPDATE comments SET like_count = (SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 1) WHERE id = ?`
	_, err := db.Exec(query, commentId, commentId)
	return err
}

// Crate HasUserDislikedComment in database/likes.go
func HasUserDislikedComment(db *sql.DB, userId, commentId int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 0`
	err := db.QueryRow(query, userId, commentId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Crate DislikeComment in database/likes.go
func DislikeComment(db *sql.DB, userId, commentId int) error {
	query := `INSERT INTO comment_likes (user_id, comment_id, is_like) VALUES (?, ?, 0)`
	_, err := db.Exec(query, userId, commentId)
	return err
}

// Crate UndislikeComment in database/likes.go
func UndislikeComment(db *sql.DB, userId, commentId int) error {
	query := `DELETE FROM comment_likes WHERE user_id = ? AND comment_id = ? AND is_like = 0`
	_, err := db.Exec(query, userId, commentId)
	return err
}

// Crate GetCommentDislikeCount in database/likes.go
func GetCommentDislikeCount(db *sql.DB, commentId int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 0`
	err := db.QueryRow(query, commentId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Crate UpdateCommentDislikeCount in database/likes.go
func UpdateCommentDislikeCount(db *sql.DB, commentId int) error {
	query := `UPDATE comments SET dislike_count = (SELECT COUNT(*) FROM comment_likes WHERE comment_id = ? AND is_like = 0) WHERE id = ?`
	_, err := db.Exec(query, commentId, commentId)
	return err
}
*/

// we will create CommnetReaction Handler to Handel all Comment Like and dislike from frontend

// func CommentReactionHandler(w http.ResponseWriter, r *http.Request) {
// 	// Parse the request body to get userId, commentId, and reaction type (like or dislike)
// 	var requestData struct {
// 		UserId    int
// 		CommentId int  `json:"comment_id"`
// 		IsLike    bool `json:"is_like"`
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	// Get the database connection from the context
// 	db, ok := r.Context().Value("db").(*sql.DB)
// 	if !ok {
// 		http.Error(w, "Database connection not found", http.StatusInternalServerError)
// 		return
// 	}
// 	session, err := r.Cookie("session")
// 	if err != nil {
// 		http.Error(w, "No Session Found, You are not authorized", http.StatusUnauthorized)
// 		return
// 	}
// 	requestData.UserId, err = database.GetUidFromToken(db, session.Value)
// 	if err != nil {
// 		http.Error(w, "You are not authorized", http.StatusUnauthorized)
// 		return
// 	}

// 	if requestData.IsLike {
// 		// Handle like action
// 		hasLiked, err := database.HasUserLikedComment(db, requestData.UserId, requestData.CommentId)
// 		if err != nil {
// 			http.Error(w, "Error checking like status", http.StatusInternalServerError)
// 			return
// 		}
// 		if hasLiked {
// 			err = database.UnlikeComment(db, requestData.UserId, requestData.CommentId)
// 		} else {
// 			err = database.LikeComment(db, requestData.UserId, requestData.CommentId)
// 		}
// 		if err != nil {
// 			http.Error(w, "Error updating like status", http.StatusInternalServerError)
// 			return
// 		}
// 		err = database.UpdateCommentLikeCount(db, requestData.CommentId)
// 		if err != nil {
// 			http.Error(w, "Error updating like count", http.StatusInternalServerError)
// 			return
// 		}
// 	} else {
// 		// Handle dislike action
// 		hasDisliked, err := database.HasUserDislikedComment(db, requestData.UserId, requestData.CommentId)
// 		if err != nil {
// 			http.Error(w, "Error checking dislike status", http.StatusInternalServerError)
// 			return
// 		}
// 		if hasDisliked {
// 			err = database.UndislikeComment(db, requestData.UserId, requestData.CommentId)
// 		} else {
// 			err = database.DislikeComment(db, requestData.UserId, requestData.CommentId)
// 		}
// 		if err != nil {
// 			http.Error(w, "Error updating dislike status", http.StatusInternalServerError)
// 			return
// 		}
// 		err = database.UpdateCommentDislikeCount(db, requestData.CommentId)
// 		if err != nil {
// 			http.Error(w, "Error updating dislike count", http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
