package database

import (
	"database/sql"
	"fmt"

	"forum/structs"

	"golang.org/x/crypto/bcrypt"
)

func QuerryLatestPosts(db *sql.DB, ammount int) ([]structs.Post, error) {
	res := make([]structs.Post, 0, ammount)
	rows, err := db.Query("SELECT * FROM posts ORDER BY created_at DESC LIMIT ?", ammount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post structs.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt)
		if err != nil {
			return res, fmt.Errorf("failed to scan row: %w", err)
		}
		res = append(res, post)
		fmt.Printf("Post ID: %d, User ID: %d, Content: %s, Created At: %s\n", post.ID, post.UserID, post.Content, post.CreatedAt)
	}

	err = rows.Err()
	if err != nil {
		return res, err
	}
	return res, nil
}

func CreateUser(db *sql.DB, email, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO users(email ,username, password) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, username, hashedPassword)
	if err != nil {
		return err
	}
	fmt.Println("User created successfully!")
	return nil
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
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, post)
		fmt.Printf("Post ID: %d, User ID: %d, Content: %s, Created At: %s\n", post.ID, post.UserID, post.Content, post.CreatedAt)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return res, nil
}
