package noboiler

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var (
	// DB : shared db
	DB *sqlx.DB
)

// User :
type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

// Language :
type Language struct {
	ID       int64  `json:"id" db:"id"`
	Language string `json:"language" db:"language"`
}

// UserLanguage :
type UserLanguage struct {
	UserID     int64 `json:"user_id" db:"user_id"`
	LanguageID int64 `json:"language_id" db:"language_id"`
}

// InitDB :
func InitDB() *sqlx.DB {
	var err error
	DB, err = sqlx.Connect("postgres", "user=postgres password=postgres dbname=sqlboiler_exp sslmode=disable")
	if err != nil {
		log.Fatal("InitDB", err)
	}
	return DB
}

// InsertUser :
func InsertUser(username string) (int64, error) {
	var userID int64
	err := DB.QueryRow(`
		INSERT INTO users(username) VALUES ($1) RETURNING id
	`, username).Scan(&userID)
	if err != nil {
		return userID, err
	}
	return userID, nil
}

// GetUser :
func GetUser(userID int64) (*User, error) {
	var u User
	err := DB.QueryRowx(`
		SELECT id, username FROM users WHERE id = $1
	`, userID).StructScan(&u)
	if err != nil {
		return &u, err
	}
	return &u, nil
}
