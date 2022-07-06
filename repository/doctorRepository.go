package repository

import (
	"clinic/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DoctorRepository interface {
	Dsel(doc *models.Doctor, id int) error
	Dselall(doc *[]models.Doctor) error
	// Dins(doc *models.User) error
	Davail(docs *[]models.Available) error
	Dsixhours(docs *[]models.Available) error
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

func (c *doctorrepositoryImpl) Davail(doctors *[]models.Available) error {
	cmd := "SELECT doc_id, COUNT(doc_id) as appointments, SUM(durationmins) as appointment_time FROM apps GROUP BY doc_id HAVING COUNT(doc_id) < 12 AND SUM(durationmins) < 480"
	err := c.db.Select(doctors, cmd)
	fmt.Println(err)
	return err
}

func (c *doctorrepositoryImpl) Dsixhours(doctors *[]models.Available) error {
	cmd := "SELECT doc_id, COUNT(doc_id) as appointments, SUM(durationmins) as appointment_time FROM apps GROUP BY doc_id HAVING SUM(durationmins) > 360"
	return c.db.Select(doctors, cmd)
}
