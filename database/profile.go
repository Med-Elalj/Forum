package database

import (
	"database/sql"

	"forum/database/querries"
	"forum/structs"
)

func GetUserProfile(db *sql.DB, username string) (structs.UserProfile, error) {
	var userProfile structs.UserProfile
	var post structs.Post

	// Query the user and their posts
	rows, err := db.Query(querries.GetUserProfile, username)
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
