package models

import "time"

type Appointment struct {
	DocId      int       `db:"doc_id"`
	PatId      int       `db:"pat_id"`
	Start_time time.Time `db:"start_time"`
	End_time   time.Time `db:"end_time"`
}
