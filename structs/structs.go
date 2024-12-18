package structs

import (
	"time"
)

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
	ID           int       `json:"post_id"`
	UserID       int       `json:"-"`
	UserName     string    `json:"author_username"`
	Title        string    `json:"post_title"`
	Content      string    `json:"post_content"`
	LikeCount    int       `json:"like_count"`
	DislikeCount int       `json:"dislike_count"`
	CommentCount int       `json:"comment_count"`
	CreatedAt    time.Time `json:"post_creation_time"`
	Categories   []string  `json:"post_categories"`
	Liked        string    `json:"view"`
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

type Profile struct {
	UserName     string
	PFP          string
	ArticleCount int
	CommentCount int
	Categories   map[string]int
}
