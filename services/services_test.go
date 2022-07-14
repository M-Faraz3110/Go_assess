package services

import (
	"clinic/db"
	"clinic/logger"
	"clinic/models"
	"clinic/repository"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoctorsvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.DoctorRepositoryProvider(dbctx, logger)
	ds := DoctorServiceProvider(dr, logger)
	_, err1 := ds.Doctor(7)
	_, err2 := ds.Doctor(1)
	assert.Equal(t, nil, err1)
	assert.Equal(t, errors.New("INVALID ID"), err2)
}

func TestAvailsvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.DoctorRepositoryProvider(dbctx, logger)
	ds := DoctorServiceProvider(dr, logger)
	_, err1 := ds.Avail()
	assert.Equal(t, nil, err1)
}

func TestSixHourssvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.DoctorRepositoryProvider(dbctx, logger)
	ds := DoctorServiceProvider(dr, logger)
	_, err1 := ds.SixHours()
	assert.Equal(t, nil, err1)
}

func TestMostAppssvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.DoctorRepositoryProvider(dbctx, logger)
	ds := DoctorServiceProvider(dr, logger)
	_, err1 := ds.MostApps()
	assert.Equal(t, nil, err1)
}

func TestSlotssvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.AppointmentRepositoryProvider(dbctx, logger)
	ds := AppointmentServiceProvider(dr, logger)
	_, err1 := ds.Slots(8)
	assert.Equal(t, nil, err1)

}

func TestCancelsvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.AppointmentRepositoryProvider(dbctx, logger)
	ds := AppointmentServiceProvider(dr, logger)
	err1 := ds.Cancel(7)
	assert.Equal(t, nil, err1)
}

func TestHistorysvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.AppointmentRepositoryProvider(dbctx, logger)
	ds := AppointmentServiceProvider(dr, logger)
	_, err1 := ds.History(7)
	assert.Equal(t, nil, err1)
}

func TestBooksvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.AppointmentRepositoryProvider(dbctx, logger)
	ds := AppointmentServiceProvider(dr, logger)
	err1 := ds.Book(models.Appointment{DocId: 8, PatId: 9, Start_time: time.Now(), End_time: time.Now().Add(time.Minute * 20)})
	err2 := ds.Book(models.Appointment{DocId: 8, PatId: 9, Start_time: time.Now(), End_time: time.Now().Add(time.Hour * 6)})
	assert.Equal(t, nil, err1)
	assert.Equal(t, errors.New("time invalid"), err2)
}

func TestAppsvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.AppointmentRepositoryProvider(dbctx, logger)
	ds := AppointmentServiceProvider(dr, logger)
	_, err1 := ds.App(10)
	assert.Equal(t, nil, err1)
}

func TestRegistersvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.UserRepositoryProvider(dbctx, logger)
	ds := UserServiceProvider(dr, logger)
	err1 := ds.Register(models.User{
		Username: "TEST",
		Password: "TEST",
		Type:     "doctor",
	})
	err2 := ds.Register(models.User{
		Username: "TEST",
		Password: "TEST",
		Type:     "TEST",
	})
	assert.Equal(t, nil, err1)
	assert.Equal(t, errors.New("INVALID USER TYPE"), err2)
}

func TestLoginsvc(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := repository.UserRepositoryProvider(dbctx, logger)
	ds := UserServiceProvider(dr, logger)
	_, err1 := ds.Login(models.User{
		Username: "doctor1",
		Password: "doc1",
		Type:     "doctor",
	})
	_, err2 := ds.Login(models.User{
		Username: "TEST",
		Password: "TEST",
		Type:     "TEST",
	})
	assert.Equal(t, nil, err1)
	assert.Equal(t, errors.New("INVALID USER"), err2)
}
