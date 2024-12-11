package querries

const (
	GetLatestPostsL = `SELECT p.id, p.user_id, p.title, p.content, p.created_at, u.username, GROUP_CONCAT(c.name , "|" ORDER BY c.name ) AS categories
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		GROUP BY p.id
		ORDER BY p.created_at ASC
		LIMIT ?;`
	GetPostsbyUserL    = `SELECT * FROM posts WHERE user_id =(SELECT id FROM users WHERE username=? ) ORDER BY created_at DESC LIMIT?`
	GetCommentsByPostL = "SELECT c.*, u.username FROM comments c JOIN users u ON c.user_id = u.id WHERE post_id=? ORDER BY c.created_at DESC LIMIT ?"
	GetUserProfile     = `SELECT u.id, u.username, u.created_at, p.id AS post_id, p.title, p.content, p.created_at AS post_created_at
	FROM users u
	LEFT JOIN posts p ON u.id = p.user_id
	WHERE u.username = ?
	ORDER BY p.created_at`
)
