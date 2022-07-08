package models

import "time"

type Times struct {
	Start_time time.Time `db:"start_time"`
	End_time   time.Time `db:"end_time"`
}
