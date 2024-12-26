package database

import (
	"database/sql"
	"fmt"

	"forum/database/querries"
	"forum/structs"
)

func CreateComment(db *sql.DB, UserId, PostId int, content string) (error, int) {
	stmt, err := db.Prepare("INSERT INTO comments(user_id, post_id, content) VALUES(?,?,?)")
	if err != nil {
		return err, 0
	}
	defer stmt.Close()

	res, err := stmt.Exec(UserId, PostId, content)
	if err != nil {
		return err, 0
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err, 0
	}
	return nil, int(id)
}

func GetCommentsByPost(db *sql.DB, postId int) ([]structs.Comment, error) {
	res := make([]structs.Comment, 0)
	rows, err := db.Query(querries.GetCommentsByPostL, postId, 10)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment structs.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.LikeCount, &comment.DislikeCount, &comment.CreatedAt, &comment.UserName)
		if err != nil {
			return res, err
		}
		res = append(res, comment)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetCommentById(db *sql.DB, commentId int) (structs.Comment, error) {
	var comment structs.Comment
	err := db.QueryRow(querries.GetCommentsByID, commentId).Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content,
		&comment.LikeCount, &comment.DislikeCount, &comment.CreatedAt, &comment.UserName)
	if err == sql.ErrNoRows {
		return comment, fmt.Errorf("comment with ID %d not found", commentId)
	}
	return comment, err
}
