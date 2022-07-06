package repository

import (
	"clinic/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AppointmentRepository interface {
	Asel(app *models.Appointment, id int) error
	Aselall(apps *[]models.Appointment, id int) error
	Adel(id int) error
	Ains(app *models.Appointment) error
	Aslots(apps *[]models.Appointment, id int) error
	AMostApps(apps *[]models.Mostapps) error
}

type appointmentrepositoryImpl struct {
	db *sqlx.DB
}

//=============================================	   Constructor and DI		========================================================
var _ AppointmentRepository = (*appointmentrepositoryImpl)(nil)

func AppointmentRepositoryProvider(db *sqlx.DB) AppointmentRepository {
	return &appointmentrepositoryImpl{db: db}
}

//=============================================	 	SVC Functions		========================================================

func (c *appointmentrepositoryImpl) Asel(app *models.Appointment, id int) error {
	cmd := fmt.Sprintf("SELECT doc_id, duration, pat_id FROM apps WHERE id = %v", id)
	return c.db.Select(app, cmd)
}

func (c *appointmentrepositoryImpl) Aselall(app *[]models.Appointment, id int) error {
	cmd := fmt.Sprintf("SELECT doc_id, duration, pat_id FROM apps WHERE pat_id = %v", id)
	return c.db.Select(app, cmd)
}

func (c *appointmentrepositoryImpl) Aslots(app *[]models.Appointment, id int) error {
	cmd := fmt.Sprintf("SELECT doc_id, duration, pat_id FROM apps WHERE doc_id = %v", id)
	return c.db.Select(app, cmd)
}

func (c *appointmentrepositoryImpl) Adel(id int) error {
	cmd := fmt.Sprintf("DELETE from apps WHERE id = %v", id)
	_, err := c.db.Exec(cmd)
	return err
}

func (c *appointmentrepositoryImpl) Ains(app *models.Appointment) error {
	cmd := fmt.Sprintf("INSERT INTO apps(duration, doc_id, pat_id) values (%v, %v, %v)", app.Duration, app.DocId, app.PatId)
	_, err := c.db.Exec(cmd)
	if err != nil {
		// handle error
		fmt.Println(err)
		panic(err)
	}
	cmd = fmt.Sprintf("UPDATE doctors SET time_left = time_left - %v WHERE id = %v", float64(app.Duration)/60, app.DocId)
	_, err = c.db.Exec(cmd)
	return err

}

func (c *appointmentrepositoryImpl) AMostApps(app *[]models.Mostapps) error {
	cmd := "SELECT doc_id, COUNT(doc_id) FROM apps GROUP BY doc_id ORDER BY COUNT(doc_id) DESC"
	return c.db.Select(app, cmd)
}
