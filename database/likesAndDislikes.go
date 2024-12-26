package database

import (
	"database/sql"
	"fmt"
)

// Crate HasUserLikedPost in database/likes.go Based on the Satment Query Above
func HasUserLikedPost(db *sql.DB, userId, postId int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM post_likes WHERE user_id = ? AND post_id = ? AND is_like = 1`
	err := db.QueryRow(query, userId, postId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Crate LikePost in database/likes.go
func LikePost(db *sql.DB, userId, postId int) error {
	query := `INSERT INTO post_likes (user_id, post_id, is_like) VALUES (?, ?, 1)`
	fmt.Println(query)
	_, err := db.Exec(query, userId, postId)
	fmt.Println(err)
	return err
}

// Crate UnlikePost in database/likes.go
func UnlikePost(db *sql.DB, userId, postId int) error {
	query := `DELETE FROM post_likes WHERE user_id = ? AND post_id = ? AND is_like = 1`
	_, err := db.Exec(query, userId, postId)
	return err
}

// Crate GetPostLikeCount in database/likes.go
func GetPostLikeCount(db *sql.DB, postId int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM post_likes WHERE post_id = ? AND is_like = 1`
	err := db.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Crate HasUserDislikedPost in database/likes.go
func HasUserDislikedPost(db *sql.DB, userId, postId int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM post_likes WHERE user_id = ? AND post_id = ? AND is_like = 0`
	err := db.QueryRow(query, userId, postId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Crate DislikePost
func DislikePost(db *sql.DB, userId, postId int) error {
	query := `INSERT INTO post_likes (user_id, post_id, is_like) VALUES (?, ?, 0)`
	_, err := db.Exec(query, userId, postId)
	return err
}

// Crate UndislikePost in
func UndislikePost(db *sql.DB, userId, postId int) error {
	query := `DELETE FROM post_likes WHERE user_id = ? AND post_id = ? AND is_like = 0`
	_, err := db.Exec(query, userId, postId)
	return err
}

// Crate GetPostDislikeCount in
func GetPostDislikeCount(db *sql.DB, postId int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM post_likes WHERE post_id = ? AND is_like = 0`
	err := db.QueryRow(query, postId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// / in the query above, we need after each like to update the like_count in the post table
// Crate UpdatePostLikeCount in database/likes.go
func UpdatePostLikeCount(db *sql.DB, postId int) error {
	query := `UPDATE posts SET like_count = (SELECT COUNT(*) FROM post_likes WHERE post_id = ? AND is_like = 1) WHERE id = ?`
	_, err := db.Exec(query, postId, postId)
	return err
}

// the same for dislike
// Crate UpdatePostDislikeCount in database/likes.go
func UpdatePostDislikeCount(db *sql.DB, postId int) error {
	query := `UPDATE posts SET dislike_count = (SELECT COUNT(*) FROM post_likes WHERE post_id = ? AND is_like = 0) WHERE id = ?`
	_, err := db.Exec(query, postId, postId)
	return err
}
