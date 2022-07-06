package models

type Doctor struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
}
