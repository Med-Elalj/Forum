package database

import (
	"database/sql"

	"forum/structs"
)

func CreateView(db *sql.DB, view structs.View) error {
	var stmt *sql.Stmt
	var err error
	if view.IsPost {
		stmt, err = db.Prepare("INSERT INTO likes(user_id, post_id, is_like) VALUES(?,?,?)")
	} else {
		stmt, err = db.Prepare("INSERT INTO likes(user_id, comment_id, is_like) VALUES(?,?,?)")
	}
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(view.UserID, view.PostID, view.IsPost, view.IsLike)
	if err != nil {
		return err
	}
	return nil
}

func GetLikes(db *sql.DB, postID int, isPost bool) (int, int, error) {
	var totalLikes, totalDislikes int
	if isPost {
		err := db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id =? AND is_like =?", postID, true).Scan(&totalLikes)
		if err != nil && err != sql.ErrNoRows {
			return 0, 0, err
		}
		err = db.QueryRow("SELECT COUNT(*) FROM likes WHERE post_id =? AND is_like =?", postID, false).Scan(&totalDislikes)
		if err != nil && err != sql.ErrNoRows {
			return 0, 0, err
		}
	} else {
		err := db.QueryRow("SELECT COUNT(*) FROM likes WHERE comment_id =? AND is_like =?", postID, true).Scan(&totalLikes)
		if err != nil && err != sql.ErrNoRows {
			return 0, 0, err
		}
		err = db.QueryRow("SELECT COUNT(*) FROM likes WHERE comment_id =? AND is_like =?", postID, false).Scan(&totalDislikes)
		if err != nil && err != sql.ErrNoRows {
			return 0, 0, err
		}
	}
	return totalLikes, totalDislikes, nil
}
