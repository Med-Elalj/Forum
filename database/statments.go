package database

var tables = map[string]string{
	"users": `CREATE TABLE IF NOT EXISTS users (
	  id INTEGER PRIMARY KEY,
	  email TEXT UNIQUE NOT NULL,
	  username TEXT UNIQUE NOT NULL,
	  password TEXT NOT NULL,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`,

	"posts": `CREATE TABLE IF NOT EXISTS posts (
	  id INTEGER PRIMARY KEY,
	  user_id INTEGER,
	  title TEXT NOT NULL, 
	  content TEXT NOT NULL,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`,

	"comments": `CREATE TABLE IF NOT EXISTS comments (
	  id INTEGER PRIMARY KEY,
	  post_id INTEGER,
	  user_id INTEGER,
	  content TEXT NOT NULL,
	  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	  CONSTRAINT no_duplicates UNIQUE (post_id, user_id, content)
	);`,

	"categories": `CREATE TABLE IF NOT EXISTS categories (
	  id INTEGER PRIMARY KEY,
	  name TEXT UNIQUE NOT NULL
	);`,

	"post_categories": `CREATE TABLE IF NOT EXISTS post_categories (
	  post_id INTEGER,
	  category_id INTEGER,
	  PRIMARY KEY (post_id, category_id),
	  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	  FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
	);`,

	"likes": `CREATE TABLE IF NOT EXISTS likes (
	id INTEGER PRIMARY KEY,
	user_id INTEGER,
	post_id INTEGER,
	comment_id INTEGER,
	is_like BOOLEAN NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
	FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
	CHECK (
	(post_id IS NOT NULL AND comment_id IS NULL) OR 
	(post_id IS NULL AND comment_id IS NOT NULL)
  )
	);`,

	"sessions": `CREATE TABLE IF NOT EXISTS sessions (
	  id INTEGER PRIMARY KEY,
	  user_id INTEGER,
	  session_token TEXT NOT NULL,
	  expiration TIMESTAMP NOT NULL,
	  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`,
}

var trigers = map[string]string{
	"check_post_time_difference": `CREATE TRIGGER check_post_time_difference
	  BEFORE INSERT ON posts
	  FOR EACH ROW
	  BEGIN
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
	  END;
`,
}
