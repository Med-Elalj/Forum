package database

import (
	"database/sql"
	"fmt"
	"strings"

	"forum/database/querries"
	"forum/structs"
)

func QuerryLatestPosts(db *sql.DB, user_id, ammount int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query(querries.GetLatestPostsL, user_id, ammount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		var categories sql.NullString
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.LikeCount, &post.DislikeCount, &post.CreatedAt, &post.UserName, &categories, &post.Liked)
		if categories.Valid {
			post.Categories = strings.Split(categories.String, "|")
		}
		if err != nil {
			return res, fmt.Errorf("failed to scan row: %w", err)
		}
		res = append(res, post)
		fmt.Printf("PID: %d, UID: %d, CONTENT: %12s, like:%d:%d , TIME:%15s, UName: %5s, categories %v %v\n", post.ID, post.UserID, post.Content, post.LikeCount, post.LikeCount, post.CreatedAt, post.UserName, post.Categories, post.Liked)
	}

	err = rows.Err()
	if err != nil {
		return res, err
	}
	return res, nil
}

func QuerryPostsbyUser(db *sql.DB, username string, user_id, ammount int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query(querries.GetPostsbyUserL, user_id, username, ammount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		var categories sql.NullString
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.LikeCount, &post.DislikeCount, &post.CreatedAt, &post.UserName, &categories, &post.Liked)
		if categories.Valid {
			post.Categories = strings.Split(categories.String, "|")
		}
		if err != nil {
			fmt.Println("azer", rows)
			return res, err
		}
		res = append(res, post)
		fmt.Printf("Post ID: %d, User ID: %d, Title: %15s, Content: %15s, Created At: %s\n", post.ID, post.UserID, post.Title, post.Content, post.CreatedAt)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreatePost(db *sql.DB, UserID int, title, content string, categories []string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO posts(user_id, title, content) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(UserID, title, content)
	if err != nil {
		return err
	}

	postID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	stmt_1, err := tx.Prepare(`INSERT INTO post_categories(category_id, post_id) VALUES((SELECT id FROM categories WHERE name = ?), ?)`)
	if err != nil {
		return err
	}
	defer stmt_1.Close()

	for _, category := range categories {
		_, err = stmt_1.Exec(category, postID)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
