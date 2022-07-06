package repository

import (
	"clinic/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DoctorRepository interface {
	Dsel(doc *models.Doctor, id int) error
	Dselall(doc *[]models.Doctor) error
	Dins(doc *models.Register) error
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
	cmd := fmt.Sprintf("SELECT id, username FROM doctors WHERE id = %v", id)
	fmt.Println(cmd)
	return c.db.Select(doctor, cmd)
}

func (c *doctorrepositoryImpl) Dselall(doctors *[]models.Doctor) error {
	cmd := "SELECT id, username FROM doctors"
	return c.db.Select(doctors, cmd)
}

func (c *doctorrepositoryImpl) Dins(doctor *models.Register) error {
	cmd := fmt.Sprintf("INSERT INTO doctors (username, password) values ('%s', '%s')", doctor.Username, doctor.Password)
	_, err := c.db.Exec(cmd)
	return err
}

func (c *doctorrepositoryImpl) Davail(doctors *[]models.Available) error {
	cmd := "SELECT username, time_left FROM doctors ORDER BY id ASC"
	return c.db.Select(doctors, cmd)
}

func (c *doctorrepositoryImpl) Dsixhours(doctors *[]models.Available) error {
	cmd := "SELECT username, time_left FROM doctors WHERE time_left <= 2 ORDER BY time_left ASC"
	return c.db.Select(doctors, cmd)
}
