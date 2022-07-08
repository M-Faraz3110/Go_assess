package models

type TimeReq struct {
	DocId      int    `db:"doc_id"`
	PatId      int    `db:"pat_id"`
	Start_time string `db:"start_time"`
	End_time   string `db:"end_time"`
}
