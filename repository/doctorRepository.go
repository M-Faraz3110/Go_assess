package repository

import (
	"clinic/models"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type DoctorRepository interface {
	Dsel(doc *models.Doctor, id int) error
	Dselall(doc *[]models.Doctor) error
	Distinct(docs *[]int) error
	// Dins(doc *models.User) error
	Davail(times *[]models.Times, id int) error
	Dsixhours(docs *[]models.Doctor) error
	DMostApps(docs *[]models.Mostapps) error
	// SelectDoctors()
	// InsertDoctors()
	// UpdateDoctors()
}

type doctorrepositoryImpl struct {
	db *sqlx.DB
}

//=============================================	   Constructor and DI		========================================================
var _ DoctorRepository = (*doctorrepositoryImpl)(nil)

func DoctorRepositoryProvider(db *sqlx.DB) DoctorRepository {
	return &doctorrepositoryImpl{db: db}
}

//=============================================	 	SVC Functions		========================================================

func (c *doctorrepositoryImpl) Dsel(doctor *models.Doctor, id int) error {
	cmd := fmt.Sprintf("SELECT id, username FROM users WHERE id = %v AND user_type = 'doctor'", id)
	fmt.Println(cmd)
	return c.db.Get(doctor, cmd)
}

func (c *doctorrepositoryImpl) Dselall(doctors *[]models.Doctor) error {
	cmd := "SELECT id, username FROM users WHERE user_type = 'doctor'"
	return c.db.Select(doctors, cmd)
}

// func (c *doctorrepositoryImpl) Dins(doctor *models.User) error {
// 	cmd := fmt.Sprintf("INSERT INTO doctors (username, password) values ('%s', '%s')", doctor.Username, doctor.Password)
// 	_, err := c.db.Exec(cmd)
// 	return err
// }

func (c *doctorrepositoryImpl) Davail(times *[]models.Times, id int) error {
	cmd := fmt.Sprintf("SELECT start_time, end_time FROM apps where DATE(start_time) = '%v-%v-%v' and doc_id = %v GROUP BY doc_id,start_time,end_time HAVING COUNT(doc_id) < 12 AND SUM(duration) < 480", time.Now().Year(), int(time.Now().Month()), time.Now().Day(), id)
	return c.db.Select(times, cmd)
}

func (c *doctorrepositoryImpl) Distinct(docs *[]int) error {
	// res := []models.Available{}
	cmd := "SELECT distinct doc_id from apps"
	return c.db.Select(docs, cmd)

}

// func (c *doctorrepositoryImpl) Davail(doctors *[]models.Available) error {
// 	cmd := fmt.Sprintf("SELECT doc_id, start_time, end_time FROM apps where DATE(start_time) = '%v-%v-%v' GROUP BY doc_id,start_time,end_time HAVING COUNT(doc_id) < 12 AND SUM(duration) < 480", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
// 	err := c.db.Select(doctors, cmd)
// 	fmt.Println(err)
// 	return err
// }

func (c *doctorrepositoryImpl) Dsixhours(doctors *[]models.Doctor) error {
	cmd := "SELECT users.id, users.username FROM users, apps where users.id = apps.doc_id GROUP BY users.id, users.username HAVING SUM(apps.duration) > 360"
	return c.db.Select(doctors, cmd)
}

func (c *doctorrepositoryImpl) DMostApps(doctors *[]models.Mostapps) error {
	cmd := "SELECT doc_id, COUNT(doc_id) FROM apps GROUP BY doc_id ORDER BY COUNT(doc_id) DESC"
	return c.db.Select(doctors, cmd)
}
