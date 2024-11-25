package structs

import "time"

type Post struct {
	ID        int
	UserID    int
	Content   string
	CreatedAt time.Time
}
