package database

import (
	"database/sql"

	"forum/database/querries"
	"forum/structs"
)

func CreateComment(db *sql.DB, UserId, PostId int, content string) (error, int64) {
	stmt, err := db.Prepare("INSERT INTO comments(user_id, post_id, content) VALUES(?,?,?)")
	if err != nil {
		return err, -1
	}
	defer stmt.Close()

	returnResult, err := stmt.Exec(UserId, PostId, content)
	if err != nil {
		return err, -1
	}
	id, err := returnResult.LastInsertId()
	if err != nil {
		return err, -1
	}
	return nil, id
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
