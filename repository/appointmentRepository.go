package repository

import (
	"clinic/models"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type AppointmentRepository interface {
	Asel(app *models.Appointment, id int) error
	Aselall(apps *[]models.Appointment, id int) error
	Adel(id int) error
	Ains(app *models.Appointment, duration float64) error
	Aslots(apps *[]models.Appointment, id int) error
	// AMostApps(apps *[]models.Mostapps) error
}

type appointmentrepositoryImpl struct {
	db *sqlx.DB
	l  *zap.SugaredLogger
}

//=============================================	   Constructor and DI		========================================================
var _ AppointmentRepository = (*appointmentrepositoryImpl)(nil)

func AppointmentRepositoryProvider(db *sqlx.DB, l *zap.SugaredLogger) AppointmentRepository {
	return &appointmentrepositoryImpl{db: db, l: l}
}

//=============================================	 	SVC Functions		========================================================

func (c *appointmentrepositoryImpl) Asel(app *models.Appointment, id int) error {
	cmd := fmt.Sprintf("SELECT doc_id, start_time, end_time, pat_id FROM apps WHERE id = %v", id)
	err := c.db.Get(app, cmd)
	fmt.Println(err)
	c.l.Info("select app repo SUCCESS...")
	return err
}

func (c *appointmentrepositoryImpl) Aselall(app *[]models.Appointment, id int) error {
	cmd := fmt.Sprintf("SELECT doc_id, start_time, end_time, pat_id FROM apps WHERE pat_id = %v", id)
	c.l.Info("select all apps repo SUCCESS...")
	return c.db.Select(app, cmd)
}

func (c *appointmentrepositoryImpl) Aslots(app *[]models.Appointment, id int) error {
	fmt.Println(time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	cmd := fmt.Sprintf("SELECT doc_id, pat_id, start_time, end_time FROM apps where DATE(start_time) = '%v-%v-%v' AND doc_id = %v", time.Now().Year(), int(time.Now().Month()), time.Now().Day(), id)
	c.l.Info("select slots repo SUCCESS...")
	return c.db.Select(app, cmd)
}

func (c *appointmentrepositoryImpl) Adel(id int) error {
	cmd := fmt.Sprintf("DELETE from apps WHERE id = %v", id)
	_, err := c.db.Exec(cmd)
	c.l.Info("delete app repo SUCCESS...")
	return err
}

func (c *appointmentrepositoryImpl) Ains(app *models.Appointment, duration float64) error {

	// doctorc1 := models.User{}
	// cmd := fmt.Sprintf("SELECT username, password, user_type FROM users WHERE id = %v", app.DocId)
	// err := c.db.Get(&doctorc1, cmd)
	// if err != nil {
	// 	// handle error
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// if doctorc1.Type != "doctor" {
	// 	return errors.New("invalid doctor")
	// }
	doctorc1 := models.Available{}
	cmd := fmt.Sprintf("SELECT doc_id from apps WHERE doc_id = %v and DATE(start_time) = '%v-%v-%v'", app.DocId, time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	err := c.db.Get(&doctorc1, cmd)
	if err != nil {
		cmd = fmt.Sprintf("INSERT INTO apps(start_time, end_time, doc_id, pat_id, duration) values ('%v', '%v', %v, %v, %v)", app.Start_time.Format("02 Jan 06 15:04"), app.End_time.Format("02 Jan 06 15:04"), app.DocId, app.PatId, duration)
		_, err = c.db.Exec(cmd) //COMMAND HERE
		if err != nil {
			// handle error
			c.l.Panic("invalid booking details,..")
			fmt.Println(err)
			return err
		}
		return err
	}
	cmd = fmt.Sprintf("SELECT doc_id FROM apps WHERE doc_id = %v and DATE(start_time) = '%v-%v-%v' GROUP BY doc_id HAVING COUNT(doc_id) < 12 AND SUM(duration) < 480", app.DocId, time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	err = c.db.Get(&doctorc1, cmd)
	if err != nil {
		// handle error
		c.l.Panicf("doctor %v unavailable...", app.DocId)
		fmt.Println(err)
		return err
	}

	cmd = fmt.Sprintf("INSERT INTO apps(start_time, end_time, doc_id, pat_id, duration) values ('%v', '%v', %v, %v, %v)", app.Start_time.Format("02 Jan 06 15:04"), app.End_time.Format("02 Jan 06 15:04"), app.DocId, app.PatId, duration)
	_, err = c.db.Exec(cmd)
	if err != nil {
		// handle error
		c.l.Panic("invalid booking details..")
		fmt.Println(err)
		return err
	}
	c.l.Info("insert app repo SUCCESS...")
	return err

}

// func (c *appointmentrepositoryImpl) AMostApps(app *[]models.Mostapps) error {
// 	cmd := "SELECT doc_id, COUNT(doc_id) FROM apps GROUP BY doc_id ORDER BY COUNT(doc_id) DESC"
// 	return c.db.Select(app, cmd)
// }
