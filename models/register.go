package models

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}
