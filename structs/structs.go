package structs

import "time"

type Comment struct {
	ID           int
	UserID       int
	PostID       int
	UserName     string
	LikeCount    int
	DislikeCount int
	Content      string
	CreatedAt    time.Time
}

type Post struct {
	ID           int
	UserID       int
	UserName     string
	Title        string
	Content      string
	LikeCount    int
	DislikeCount int
	CreatedAt    time.Time
	Categories   []string
}

type View struct {
	UserID int
	ID     int
	IsLike bool
}

type UserProfile struct {
	UID       int
	Username  string
	Posts     []Post
	CreatedAt time.Time
}
