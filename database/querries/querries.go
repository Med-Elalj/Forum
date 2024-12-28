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
		ORDER BY p.created_at DESC
		LIMIT ?
		OFFSET ?;`
	GetPostsbyUserL = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
		COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		WHERE u.username = ?
		GROUP BY p.id
		ORDER BY p.created_at DESC 
		LIMIT ?;`
	GetPostsbyCategoryL = `SELECT p.*, c.name AS category_name, GROUP_CONCAT(c.name, "|") AS categories,
        COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		WHERE c.name = ?
		GROUP BY p.id
		ORDER BY p.created_at DESC 
		LIMIT ?;`
	GetPostsbyUserLikeL = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
		COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ? AND is_like = 1
		WHERE u.username = ?
		GROUP BY p.id
		ORDER BY p.created_at DESC 
		LIMIT ?;`
	GetPostByID = `SELECT p.*, u.username, GROUP_CONCAT(c.name , "|") AS categories ,
		COALESCE(pl.is_like, "null") AS is_like
		FROM posts p
		JOIN users u ON p.user_id = u.id
		LEFT JOIN post_categories pc ON p.id = pc.post_id
		LEFT JOIN categories c ON pc.category_id = c.id
		LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = ?
		WHERE p.id = ?;`
	GetCommentsByPostL    = "SELECT c.*, u.username FROM comments c JOIN users u ON c.user_id = u.id WHERE post_id=? ORDER BY c.created_at DESC LIMIT ?"
	GetCommentsByID       = "SELECT c.*, u.username FROM comments c JOIN users u ON c.user_id = u.id WHERE id=? ORDER BY c.created_at DESC LIMIT ?"
	GetUserProfileByUname = `SELECT u.id, u.username, u.created_at, COUNT(p.id) AS post_count, COUNT(c.id) AS comment_count
	FROM users u
	LEFT JOIN posts p ON u.id = p.user_id
	LEFT JOIN comments c ON u.id = c.user_id
	WHERE u.username = ?
	ORDER BY p.created_at`
	GetUserProfileByID = `SELECT u.id, u.username, u.created_at, COUNT(p.id) AS post_count, COUNT(c.id) AS comment_count
	FROM users u
	LEFT JOIN posts p ON u.id = p.user_id
	LEFT JOIN comments c ON u.id = c.user_id
	WHERE u.id = ?
	ORDER BY p.created_at`
	GetCategoriesWithPostCount = `SELECT c.name, COUNT(pc.post_id) as post_count
			FROM categories c
			JOIN post_categories pc ON c.id = pc.category_id
			GROUP BY c.name 
			ORDER BY post_count DESC
			LIMIT 6`
)
