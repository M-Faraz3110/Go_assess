package models

type Appointment struct {
	DocId    int `db:"doc_id"`
	PatId    int `db:"pat_id"`
	Duration int `db:"duration"`
}
