package database

import "database/sql"

func GetCategoriesWithPostCount(db *sql.DB) (map[string]int, error) {
	categories := make(map[string]int)
	rows, err := db.Query(`
		SELECT c.name, COUNT(pc.post_id) as post_count
		FROM categories c
		LEFT JOIN post_categories pc ON c.id = pc.category_id
		GROUP BY c.name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var postCount int
		if err := rows.Scan(&name, &postCount); err != nil {
			return nil, err
		}
		categories[name] = postCount
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
