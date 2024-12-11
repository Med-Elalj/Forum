package structs

import "time"

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	UserName  string
	Content   string
	CreatedAt time.Time
}

type Post struct {
	ID         int
	UserID     int
	UserName   string
	Title      string
	Content    string
	CreatedAt  time.Time
	Categories []string
}

type View struct {
	UserID int
	PostID int
	IsPost bool
	IsLike bool
	Time   time.Time
}

type UserProfile struct {
	UID       int
	Username  string
	Posts     []Post
	CreatedAt time.Time
}
