package querries

const (
	GetLatestPostsL = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
	    COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		GROUP BY p.id
		ORDER BY p.created_at ASC
		LIMIT ?;`
	GetPostsbyUserL = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
		COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		WHERE u.username = ?
		GROUP BY p.id
		ORDER BY created_at ASC 
		LIMIT ?;`
	GetPostByID = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
		COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		WHERE p.id = ?;`
	GetCommentsByPostL = "SELECT c.*, u.username FROM comments c JOIN users u ON c.user_id = u.id WHERE post_id=? ORDER BY c.created_at DESC LIMIT ?"
	GetUserProfile     = `SELECT u.id, u.username, u.created_at, p.id AS post_id, p.title, p.content, p.created_at AS post_created_at
	FROM users u
	LEFT JOIN posts p ON u.id = p.user_id
	WHERE u.username = ?
	ORDER BY p.created_at`
)
