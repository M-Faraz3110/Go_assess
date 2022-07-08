package models

type Available struct {
	Id    string `db:"doc_id"`
	Times []Times
}
