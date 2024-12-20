package database

import (
	"database/sql"

	"forum/database/querries"
	"forum/structs"
)

func CreateComment(db *sql.DB, UserId, PostId int, content string) error {
	stmt, err := db.Prepare("INSERT INTO comments(user_id, post_id, content) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(UserId, PostId, content)
	if err != nil {
		return err
	}
	return nil
}

func GetCommentsByPost(db *sql.DB, postId, ammount int) ([]structs.Comment, error) {
	res := make([]structs.Comment, 0, ammount)
	rows, err := db.Query(querries.GetCommentsByPostL, postId, ammount)
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
