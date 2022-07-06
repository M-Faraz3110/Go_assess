package services

import (
	"clinic/models"
	"clinic/repository"
)

type AppointmentService interface {
	Slots(id int) ([]models.Appointment, error)
	Cancel(id int) error
	History(id int) ([]models.Appointment, error)
	MostApps() ([]models.Mostapps, error)
	Book(app models.Appointment) error
	App(id int) (models.Appointment, error)
}

type appointmentServiceImpl struct {
	ar repository.AppointmentRepository
}

//=============================================	   Constructor		========================================================
var _ AppointmentService = (*appointmentServiceImpl)(nil)

func AppointmentServiceProvider(ar repository.AppointmentRepository) AppointmentService {
	return &appointmentServiceImpl{ar: ar}
}

//=============================================	 	SVC Functions		========================================================

//CALL REPO FUNCTIONS

func (c *appointmentServiceImpl) Slots(id int) ([]models.Appointment, error) {
	slots := []models.Appointment{}
	c.ar.Aslots(&slots, id)
	return slots, nil

}

func (c *appointmentServiceImpl) Cancel(id int) error {

	return c.ar.Adel(id)

}

func (c *appointmentServiceImpl) History(id int) ([]models.Appointment, error) {

	apps := []models.Appointment{}
	c.ar.Aselall(&apps, id)
	return apps, nil

}

func (c *appointmentServiceImpl) MostApps() ([]models.Mostapps, error) {

	apps := []models.Mostapps{}
	c.ar.AMostApps(&apps)
	return apps, nil

}

func (c *appointmentServiceImpl) Book(app models.Appointment) error {

	return c.ar.Ains(&app)

}

func (c *appointmentServiceImpl) App(id int) (models.Appointment, error) {
	app := models.Appointment{}
	c.ar.Asel(&app, id)
	return app, nil

}
