package structs

type Post_omar struct {
	PostID           string         `json:"post_id"`
	UserProfileImage string         `json:"user_profile_image"`
	PostDetails      PostDetails    `json:"post_details"`
	PostContent      PostContent    `json:"post_content"`
	PostCategories   []string       `json:"post_categories"`
	UserEngagement   UserEngagement `json:"user_engagement"`
}

type PostDetails_omar struct {
	PostTitle        string `json:"post_title"`
	AuthorUsername   string `json:"author_username"`
	PostCreationTime string `json:"post_creation_time"`
}

type PostContent_omar struct {
	BodyText string `json:"body_text"`
}

type UserEngagement_omar struct {
	Reactions       Reactions       `json:"reactions"`
	CommentsSection CommentsSection `json:"comments_section"`
}

type Reactions_omar struct {
	Likes    ReactionDetail `json:"likes"`
	Dislikes ReactionDetail `json:"dislikes"`
}

type ReactionDetail_omar struct {
	PostID string `json:"post_id"`
	Count  int    `json:"like_count"` // or dislike_count depending on the reaction
}

type CommentsSection_omar struct {
	CommentCount int `json:"comment_count"`
}
