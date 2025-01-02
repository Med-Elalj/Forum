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
	Liked        string
}

type Post struct {
	ID                  int       `json:"post_id"`
	UserID              int       `json:"-"`
	UserName            string    `json:"author_username"`
	Title               string    `json:"post_title"`
	Content             string    `json:"post_content"`
	LikeCount           int       `json:"like_count"`
	DislikeCount        int       `json:"dislike_count"`
	CommentCount        int       `json:"comment_count"`
	CreatedAt           time.Time `json:"post_creation_time"`
	Categories          []string  `json:"post_categories"`
	Liked               string    `json:"view"`
	UserPostsCounts     int       `json:"user_posts_count"`
	UsersCommentsCounts int       `json:"user_comments_count"`
}

type View struct {
	UserID int
	ID     int
	IsLike bool
}

type Profile struct {
	UserID       int
	UserName     string
	PFP          string
	ArticleCount int
	CommentCount int
	CreatedAt    time.Time
	CurrentPage  string
	Category     string
}

type Categories map[string]int

const (
	NotaUser = 0
	Limit    = 10
	NoOffSet = 0
)
