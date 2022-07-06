package models

type User struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Type     string `json:"type" db:"user_type"`
}
