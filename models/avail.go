package models

type Available struct {
	Id               string `db:"doc_id"`
	Appointments     string `db:"appointments"`
	Appointment_time string `db:"appointment_time"`
}
