package services

import (
	"clinic/models"
	"clinic/repository"
	"fmt"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type DoctorService interface {
	Doctor(id int) (models.Doctor, error)
	Doctors() ([]models.Doctor, error)
	MostApps() ([]models.Mostapps, error)
	Avail() ([]models.Available, error)
	SixHours() ([]models.Doctor, error)
}

type doctorServiceImpl struct {
	dr repository.DoctorRepository
	l  *zap.SugaredLogger
}

//=============================================	   Constructor 	========================================================
var _ DoctorService = (*doctorServiceImpl)(nil)

func DoctorServiceProvider(dr repository.DoctorRepository, l *zap.SugaredLogger) DoctorService {
	return &doctorServiceImpl{dr: dr, l: l}
}

//=============================================	 	SVC Functions		========================================================

//CALL REPO FUNCTIONS

func (c *doctorServiceImpl) Doctor(id int) (models.Doctor, error) {

	doc := models.Doctor{}
	c.dr.Dsel(&doc, id)
	c.l.Info("/doctor/:id service SUCCESS...")
	return doc, nil

}

func (c *doctorServiceImpl) Doctors() ([]models.Doctor, error) {
	docs := []models.Doctor{}
	c.dr.Dselall(&docs)
	c.l.Info("/doctors service SUCCESS...")
	return docs, nil
}

func (c *doctorServiceImpl) Avail() ([]models.Available, error) {
	docs := []models.Available{}
	ids := []int{}
	c.dr.Distinct(&ids)
	times := []models.Times{}
	fmt.Println(ids)
	for k := range ids {
		avail := []models.Times{}
		avail = append(avail, models.Times{
			Start_time: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 9, 0, 0, 0, time.Local),
			End_time:   time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 17, 0, 0, 0, time.Local),
		})
		c.dr.Davail(&times, ids[k])
		for k := range times {

			avail[len(avail)-1].End_time = times[k].Start_time
			if avail[len(avail)-1].Start_time == avail[len(avail)-1].End_time {
				avail = avail[:len(avail)-1]
			}
			avail = append(avail, models.Times{
				Start_time: times[k].End_time,
				End_time:   time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 17, 0, 0, 0, time.Local),
			})

		}
		docs = append(docs, models.Available{
			Id:    strconv.Itoa(ids[k]),
			Times: avail,
		})
	}
	c.l.Info("/avail service SUCCESS...")
	return docs, nil
}

// func (c *doctorServiceImpl) Avail() ([]models.Available, error) {
// 	res := []models.Available{}
// 	docs := []models.Available{}
// 	// now := time.Now()
// 	c.dr.Davail(&docs)
// 	fmt.Println(docs)

// 	return res, nil
// }

func (c *doctorServiceImpl) SixHours() ([]models.Doctor, error) {
	docs := []models.Doctor{}
	c.dr.Dsixhours(&docs)
	c.l.Info("/sixhours service SUCCESS...")
	return docs, nil
}

func (c *doctorServiceImpl) MostApps() ([]models.Mostapps, error) {
	docs := []models.Mostapps{}
	c.dr.DMostApps(&docs)
	c.l.Info("/mostapps service SUCCESS...")
	return docs, nil
}
