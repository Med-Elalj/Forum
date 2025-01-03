package database

var tables = map[string]string{
	"Pragma": `PRAGMA foreign_keys = ON;`,
	"users": `
		CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,

	"posts": `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL, 
		content TEXT NOT NULL,
		like_count INTEGER DEFAULT 0,
		dislike_count INTEGER DEFAULT 0,
		comment_count INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,

	"comments": `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY NOT NULL,
		post_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		like_count INTEGER DEFAULT 0,
		dislike_count INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		CONSTRAINT no_duplicates UNIQUE (post_id, user_id, content)
		);`,

	"categories": `CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY  NOT NULL,
		name TEXT UNIQUE NOT NULL
		);`,

	"post_categories": `CREATE TABLE IF NOT EXISTS post_categories (
		post_id INTEGER NOT NULL,
		category_id INTEGER NOT NULL,
		PRIMARY KEY (post_id, category_id),
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
		FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
		);`,

	"post_likes": `CREATE TABLE IF NOT EXISTS post_likes (
		user_id INTEGER NOT NULL,
		post_id INTEGER NOT NULL,
		is_like BOOLEAN NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, post_id)
		);`,

	"comment_likes": `CREATE TABLE IF NOT EXISTS comment_likes (
		user_id INTEGER NOT NULL,
		comment_id INTEGER NOT NULL,
		is_like BOOLEAN NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
		PRIMARY KEY (user_id, comment_id)			
		);`,

	"sessions": `CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY UNIQUE NOT NULL,
		user_id INTEGER NOT NULL,
		session_token TEXT NOT NULL,
		expiration TIMESTAMP NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
}

type triger struct {
	name, statment string
	tables         []string
}

var trigers = []triger{
	{
		"check_post_time_difference",
		`CREATE TRIGGER check_post_time_difference
    BEFORE INSERT ON posts
    FOR EACH ROW
BEGIN
    -- Check if a similar post already exists in the last year
    SELECT
        CASE
            WHEN EXISTS (
                SELECT 1 FROM posts
                WHERE user_id = NEW.user_id
                AND title = NEW.title
                AND content = NEW.content
                AND created_at >= datetime('now', '-1 year')
            ) THEN
                RAISE (ABORT, 'A post with the same title and content already exists within the last year.')
        END;
END;`,
		[]string{},
	},
	{"one_session_per_user",
		`CREATE TRIGGER one_session_per_user
		BEFORE INSERT ON sessions
		FOR EACH ROW
		BEGIN
		DELETE FROM sessions WHERE user_id = NEW.user_id;
		END;`,
		[]string{},
	},
	{
		"comment_count_insert",
		`CREATE TRIGGER comment_count_insert
		BEFORE INSERT ON comments
		FOR EACH ROW
		BEGIN
		UPDATE posts
		SET comment_count = comment_count + 1
		WHERE id = NEW.post_id;
		END;`,
		[]string{},
	},
	{
		"comment_count_delete",
		`CREATE TRIGGER comment_count_delete
        BEFORE DELETE ON comments
        FOR EACH ROW
        BEGIN
        UPDATE posts
        SET comment_count = comment_count - 1
        WHERE id = OLD.post_id;
        END;`,
		[]string{},
	},
	{
		"_react_count_insert",
		`CREATE TRIGGER 1here2_react_count_insert
BEFORE INSERT ON 1here2_likes
FOR EACH ROW
BEGIN
	UPDATE 1here2s
	SET 
	like_count = like_count + (NEW.is_like = 1) ,
	dislike_count = dislike_count + (NEW.is_like = 0)
	  WHERE id = NEW.1here2_id;
END;`,
		[]string{"post", "comment"},
	},
	{
		"_react_count_insert",
		`CREATE TRIGGER 1here2_react_count_insert
BEFORE UPDATE ON 1here2_likes
FOR EACH ROW
BEGIN
UPDATE 1here2s
SET 
SET 
	like_count = like_count + 
	  (CASE WHEN NEW.is_like = 1 AND OLD.is_like = 0 THEN 1 ELSE 0 END) -
	  (CASE WHEN NEW.is_like = 0 AND OLD.is_like = 1 THEN 1 ELSE 0 END),
	dislike_count = dislike_count + 
	  (CASE WHEN NEW.is_like = 0 AND OLD.is_like = 1 THEN 1 ELSE 0 END) -
	  (CASE WHEN NEW.is_like = 1 AND OLD.is_like = 0 THEN 1 ELSE 0 END)
  WHERE id = NEW.1here2_id;
END;`,
		[]string{"post", "comment"},
	},
	{
		"_react_count_insert",
		`CREATE TRIGGER 1here2_react_count_insert
BEFORE DELETE ON 1here2_likes
FOR EACH ROW
BEGIN
UPDATE 1here2s
SET 
like_count = like_count - (OLD.is_like = 1) ,
dislike_count = dislike_count - (OLD.is_like = 0)
WHERE id = OLD.1here2_id;
END;`,
		[]string{"post", "comment"},
	},
	{
		"CreateCategories", `INSERT OR IGNORE INTO Categories (Name) VALUES
		('General'),
		('Entertainment'),
		('Health'),
		('Business'),
		('Sports'),
		('Technology');`,
		[]string{},
	},
}
