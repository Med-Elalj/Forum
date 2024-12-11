package database

import (
	"database/sql"
	"fmt"
	"strings"

	"forum/database/querries"
	"forum/structs"
)

func QuerryLatestPosts(db *sql.DB, ammount int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query(querries.GetLatestPostsL, ammount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		var categories sql.NullString
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UserName, &categories)
		if categories.Valid {
			post.Categories = strings.Split(categories.String, "|")
		}
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
	rows, err := db.Query(querries.GetPostsbyUserL, username, ammount)
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

func CreatePost(db *sql.DB, UserID int, title, content string) (int, error) {
	stmt, err := db.Prepare("INSERT INTO posts(user_id, title, content) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(UserID, title, content)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

func AddPostCategories(db *sql.DB, post_id, categories []string) error {
	stmt, err := db.Prepare(`INSERT INTO posts(category_in, post_id) VALUES (
	(SELECT id FROM categories WHERE name = ? ),?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, category := range categories {
		_, err = stmt.Exec(category, post_id)
		if err != nil {
			return err
		}
	}
	return nil
}
