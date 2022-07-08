package services

import (
	"clinic/models"
	"clinic/repository"
	"errors"
	"fmt"
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
	res := []models.Appointment{}
	slots := []models.Appointment{}
	c.ar.Aslots(&slots, id)
	slot := models.Appointment{}
	// slot = models.Appointment{
	// 	DocId:      0,
	// 	PatId:      0,
	// 	Start_time: time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.Local),
	// 	End_time:   slots[0].Start_time,
	// }
	// res = append(res, slot)
	for k, _ := range slots {
		slot = models.Appointment{
			DocId:      slots[k].DocId,
			PatId:      slots[k].PatId,
			Start_time: slots[k].Start_time,
			End_time:   slots[k].End_time,
		}
		res = append(res, slot)
		// if k+1 < len(slots) {
		// 	if slots[k+1].Start_time != slots[k].End_time {
		// 		slot = models.Appointment{
		// 			DocId:      slots[k].DocId,
		// 			PatId:      slots[k].PatId,
		// 			Start_time: slots[k].End_time,
		// 			End_time:   slots[k+1].Start_time,
		// 		}
		// 	}
		// }

	}
	// slot = models.Appointment{
	// 	DocId:      0,
	// 	PatId:      0,
	// 	Start_time: slots[len(slots)-1].End_time,
	// 	End_time:   time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, time.Local),
	// }
	// res = append(res, slot)
	return res, nil
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
	// layout := "02 Jan 06 15:04"
	// tm1, err := time.Parse(layout, app.Start_time)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// tm2, err := time.Parse(layout, app.End_time)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	diff := app.End_time.Sub(app.Start_time)
	fmt.Println(diff.Minutes())
	if diff.Minutes() < 1 {
		return errors.New("time invalid")
	}
	return c.ar.Ains(&app, diff.Minutes())

}

func (c *appointmentServiceImpl) App(id int) (models.Appointment, error) {
	app := models.Appointment{}
	c.ar.Asel(&app, id)
	return app, nil

}
