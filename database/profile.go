package database

import (
	"database/sql"

	"forum/structs"
)

func GetUserProfile(db *sql.DB, username string) (structs.UserProfile, error) {
	var userProfile structs.UserProfile
	var post structs.Post

	// Query the user and their posts
	rows, err := db.Query(`
	SELECT u.id, u.username, u.created_at, p.id AS post_id, p.title, p.content, p.created_at AS post_created_at
	FROM users u
	LEFT JOIN posts p ON u.id = p.user_id
	WHERE u.username = ?
	ORDER BY p.created_at`, username)
	if err != nil {
		return userProfile, err
	}
	defer rows.Close()

	// Loop over the rows and populate the UserProfile struct
	for rows.Next() {
		err := rows.Scan(&userProfile.UID, &userProfile.Username, &userProfile.CreatedAt,
			&post.ID, &post.Title, &post.Content, &post.CreatedAt)
		if err != nil {
			return userProfile, err
		}

		// Add post to the user's post list
		if post.ID != 0 { // Avoid adding empty posts if the user has no posts
			userProfile.Posts = append(userProfile.Posts, post)
		}
	}

	// Check for any error during iteration
	if err := rows.Err(); err != nil {
		return userProfile, err
	}

	return userProfile, nil
}
