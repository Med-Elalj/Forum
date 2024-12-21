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
		like_count INTEGER DEFAULT 0,
		dislike_count INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,

	"comments": `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY,
		post_id INTEGER,
		user_id INTEGER,
		content TEXT NOT NULL,
		like_count INTEGER DEFAULT 0,
		dislike_count INTEGER DEFAULT 0,
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
		id INTEGER PRIMARY KEY UNIQUE,
		user_id INTEGER NOT NULL,
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
		    IF EXISTS (
		            SELECT 1 FROM posts 
				  WHERE user_id = NEW.user_id 
				  AND title = NEW.title 
				  AND content = NEW.content 
				  AND created_at >= datetime('now', '-1 year')
		        ) THEN
		            RAISE (ABORT, 'A post with the same title and content already exists within the last year.')
		    END IF;
		END;
`,

	"post_react_count_insert": `CREATE TRIGGER post_react_count_insert
		BEFORE INSERT ON post_likes
		FOR EACH ROW
		BEGIN
		    -- If it's a like, increment like_count
		    IF NEW.is_like THEN
		        UPDATE posts
		        SET like_count = like_count + 1
		        WHERE id = NEW.post_id;
		    -- If it's a dislike, increment dislike_count
		    ELSEIF NOT NEW.is_like THEN
		        UPDATE posts
		        SET dislike_count = dislike_count + 1
		        WHERE id = NEW.post_id;
			ELSE 
			    RAISE (ABORT, "unable to insert.")
		    END IF;
		END;`,

	"post_react_count_update": `CREATE TRIGGER post_react_count_update
		BEFORE UPDATE ON post_likes
		FOR EACH ROW
		BEGIN
		    -- Update the created_at timestamp to now
	        SET NEW.created_at = CURRENT_TIMESTAMP;

		    -- If the like/dislike was changed from like to dislike
		    IF OLD.is_like AND NOT NEW.is_like THEN
		        UPDATE posts
		        SET like_count = like_count - 1,
		            dislike_count = dislike_count + 1
		        WHERE id = OLD.post_id;
		    -- If the like/dislike was changed from dislike to like
		    ELSEIF NOT OLD.is_like AND NEW.is_like THEN
		        UPDATE posts
		        SET like_count = like_count + 1,
		            dislike_count = dislike_count - 1
		        WHERE id = OLD.post_id;
			ELSE 
			    RAISE (ABORT, "unable to update.")
		    END IF;
		END;`,

	"post_react_count_delete": `CREATE TRIGGER post_react_count_delete
		BEFORE DELETE ON post_likes
		FOR EACH ROW
		BEGIN
		    -- If it was a like, decrement like_count
		    IF OLD.is_like THEN
		        UPDATE posts
		        SET like_count = like_count - 1
		        WHERE id = OLD.post_id;
		    -- If it was a dislike, decrement dislike_count
		    ELSEIF NOT OLD.is_like
		        UPDATE posts
		        SET dislike_count = dislike_count - 1
		        WHERE id = OLD.post_id;
			ELSE 
			    RAISE (ABORT, "unable to delete.")
		    END IF;
		END;`,

	"comment_react_count_insert": `CREATE TRIGGER comment_react_count_insert
		BEFORE INSERT ON comment_likes
		FOR EACH ROW
		BEGIN
		    -- If it's a like, increment like_count
		    IF NEW.is_like THEN
		        UPDATE comments
		        SET like_count = like_count + 1
		        WHERE id = NEW.comment_id;
		    -- If it's a dislike, increment dislike_count
		    ELSEIF NOT NEW.is_like THEN
		        UPDATE comments
		        SET dislike_count = dislike_count + 1
		        WHERE id = NEW.comment_id;
			ELSE 
			    RAISE (ABORT, "unable to insert.")
		    END IF;
		END;`,

	"comment_react_count_update": `CREATE TRIGGER comment_react_count_update
		BEFORE UPDATE ON comment_likes
		FOR EACH ROW
		BEGIN
		    -- Update the created_at timestamp to now
	        SET NEW.created_at = CURRENT_TIMESTAMP;

		    -- If the like/dislike was changed from like to dislike
		    IF OLD.is_like AND NOT NEW.is_like THEN
		        UPDATE comments
		        SET like_count = like_count - 1,
		            dislike_count = dislike_count + 1
		        WHERE id = OLD.comment_id;
		    -- If the like/dislike was changed from dislike to like
		    ELSEIF NOT OLD.is_like AND NEW.is_like THEN
		        UPDATE comments
		        SET like_count = like_count + 1,
		            dislike_count = dislike_count - 1
		        WHERE id = OLD.comment_id;
			ELSE 
			    RAISE (ABORT, "unable to update.")
		    END IF;
		END;`,

	"comment_react_count_delete": `CREATE TRIGGER comment_react_count_delete
		BEFORE DELETE ON comment_likes
		FOR EACH ROW
		BEGIN
		    -- If it was a like, decrement like_count
		    IF OLD.is_like THEN
		        UPDATE comments
		        SET like_count = like_count - 1
		        WHERE id = OLD.comment_id;
		    -- If it was a dislike, decrement dislike_count
		    ELSEIF NOT OLD.is_like
		        UPDATE comments
		        SET dislike_count = dislike_count - 1
		        WHERE id = OLD.comment_id;
			ELSE 
			    RAISE (ABORT, "unable to delete.")
		    END IF;
		END;`,
}
