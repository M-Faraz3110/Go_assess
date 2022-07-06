package services

import (
	"clinic/models"
	"clinic/repository"
)

type DoctorService interface {
	Doctor(id int) (models.Doctor, error)
	Doctors() ([]models.Doctor, error)
	Avail() ([]models.Available, error)
	SixHours() ([]models.Available, error)
}

type doctorServiceImpl struct {
	dr repository.DoctorRepository
}

//=============================================	   Constructor 	========================================================
var _ DoctorService = (*doctorServiceImpl)(nil)

func DoctorServiceProvider(dr repository.DoctorRepository) DoctorService {
	return &doctorServiceImpl{dr: dr}
}

//=============================================	 	SVC Functions		========================================================

//CALL REPO FUNCTIONS

func (c *doctorServiceImpl) Doctor(id int) (models.Doctor, error) {

	doc := models.Doctor{}
	c.dr.Dsel(&doc, id)
	return doc, nil

}

func (c *doctorServiceImpl) Doctors() ([]models.Doctor, error) {
	docs := []models.Doctor{}
	c.dr.Dselall(&docs)
	return docs, nil
}

func (c *doctorServiceImpl) Avail() ([]models.Available, error) {
	docs := []models.Available{}
	c.dr.Davail(&docs)
	return docs, nil
}

func (c *doctorServiceImpl) SixHours() ([]models.Available, error) {
	docs := []models.Available{}
	c.dr.Dsixhours(&docs)
	return docs, nil
}
