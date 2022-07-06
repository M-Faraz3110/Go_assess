package models

type Available struct {
	Username  string `db:"username"`
	Time_Left string `db:"time_left"`
}
