package database

import (
	"database/sql"
	"fmt"

	"forum/structs"
)

func QuerryLatestPosts(db *sql.DB, ammount int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query("SELECT p.*, u.username FROM posts p JOIN users u ON p.user_id = u.id ORDER BY p.created_at DESC LIMIT ?;", ammount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.UserName)
		if err != nil {
			return res, fmt.Errorf("failed to scan row: %w", err)
		}
		res = append(res, post)
		fmt.Printf("Post ID: %d, User ID: %d, Content: %s, Created At: %s, UserName: %s\n", post.ID, post.UserID, post.Content, post.CreatedAt, post.UserName)
	}

	err = rows.Err()
	if err != nil {
		return res, err
	}
	return res, nil
}

func QuerryPostsbyUser(db *sql.DB, username string, ammount int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query("SELECT * FROM posts WHERE user_id =(SELECT id FROM users WHERE username=? ) ORDER BY created_at DESC LIMIT?", username, ammount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return res, err
		}
		res = append(res, post)
		fmt.Printf("Post ID: %d, User ID: %d, Title: %s, Content: %s, Created At: %s\n", post.ID, post.UserID, post.Title, post.Content, post.CreatedAt)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreatePost(db *sql.DB, UserID int, title, content string) error {
	stmt, err := db.Prepare("INSERT INTO posts(user_id, title, content) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(UserID, title, content)
	if err != nil {
		return err
	}
	return nil
}

func AddPostCategories(db *sql.DB, UserID int, title, content string, categories []string) error {
	stmt, err := db.Prepare(`INSERT INTO posts(category_in, post_id) VALUES (
	(SELECT id FROM categories WHERE name = ? ),
	(SELECT id FROM posts WHERE user_id= ? AND title = ? AND content =? ORDER BY created_at DESC LIMIT 1)
	)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, category := range categories {
		_, err = stmt.Exec(category, UserID, title, content)
		if err != nil {
			return err
		}
	}
	return nil
}
