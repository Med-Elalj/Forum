package database

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, email, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO users(email ,username, password) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, username, hashedPassword)
	if err != nil {
		return err
	}
	fmt.Println("User created successfully!")
	return nil
}

func GetUserByUname(db *sql.DB, username string) (string, error) {
	var hpassword string
	fmt.Println(username, hpassword)
	err := db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&hpassword)
	fmt.Println(username, hpassword)

	if err == sql.ErrNoRows {
		fmt.Println("User not found")
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		fmt.Println(err)
		return "", err
	}
	return hpassword, nil
}

func GetUserByUemail(db *sql.DB, email string) (string, error) {
	var hpassword string
	fmt.Println(email, hpassword)
	err := db.QueryRow("SELECT password FROM users WHERE email=?", email).Scan(&hpassword)
	fmt.Println(email, hpassword)

	if err == sql.ErrNoRows {
		fmt.Println("User not found")
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		fmt.Println(err)
		return "", err
	}
	return hpassword, nil
}

func AddSessionToken(db *sql.DB, username, token string) error {
	fmt.Println("session", username)
	stmt, err := db.Prepare("INSERT INTO sessions(user_id, session_token, expiration) VALUES((SELECT id FROM users WHERE username=? ),?,DATETIME(CURRENT_TIMESTAMP, '+1 hour'))")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		return err
	}
	fmt.Println("Session token created successfully!")
	return nil
}

// Returns 0 if uid not found
func GetUidFromToken(db *sql.DB, token string) (int, error) {
	if token == "" {
		// case of new user
		return 0, nil
	}
	var uid int
	err := db.QueryRow("SELECT user_id FROM sessions WHERE token=?", token).Scan(&uid)
	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("session not found")
	} else if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return uid, nil
}
