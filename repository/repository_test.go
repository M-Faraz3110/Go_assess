package repository

import (
	"clinic/db"
	"clinic/logger"
	"clinic/models"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAsel(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := AppointmentRepositoryProvider(dbctx, logger)
	err1 := dr.Asel(&models.Appointment{}, 9)
	err2 := dr.Asel(&models.Appointment{}, 2)
	assert.Equal(t, nil, err1)
	assert.Equal(t, errors.New("sql: no rows in result set"), err2)

}

func TestAselall(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := AppointmentRepositoryProvider(dbctx, logger)
	err1 := dr.Aselall(&[]models.Appointment{}, 10)
	err2 := dr.Aselall(&[]models.Appointment{}, 2)
	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)
}

func TestAslots(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := AppointmentRepositoryProvider(dbctx, logger)
	err1 := dr.Aslots(&[]models.Appointment{}, 8)
	err2 := dr.Aslots(&[]models.Appointment{}, 2)
	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)
}

func TestAdel(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := AppointmentRepositoryProvider(dbctx, logger)
	err1 := dr.Adel(8)
	err2 := dr.Adel(2)
	assert.Equal(t, nil, err1)
	assert.Equal(t, nil, err2)
}

func TestAins(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := AppointmentRepositoryProvider(dbctx, logger)
	test := models.Appointment{
		DocId:      8,
		PatId:      9,
		Start_time: time.Now(),
		End_time:   time.Now().Add(time.Minute * 20),
	}
	err1 := dr.Ains(&test, test.End_time.Sub(test.Start_time).Minutes())
	assert.Equal(t, nil, err1)

}

func TestDsel(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := DoctorRepositoryProvider(dbctx, logger)
	err1 := dr.Dsel(&models.Doctor{}, 8)
	assert.Equal(t, nil, err1)

}

func TestDselall(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := DoctorRepositoryProvider(dbctx, logger)
	err1 := dr.Dselall(&[]models.Doctor{})
	assert.Equal(t, nil, err1)

}

func TestDavail(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := DoctorRepositoryProvider(dbctx, logger)
	err1 := dr.Davail(&[]models.Times{}, 8)
	assert.Equal(t, nil, err1)
}

func TestDsixhours(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := DoctorRepositoryProvider(dbctx, logger)
	err1 := dr.Dsixhours(&[]models.Doctor{})
	assert.Equal(t, nil, err1)

}

func TestDMostapps(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := DoctorRepositoryProvider(dbctx, logger)
	err1 := dr.DMostApps(&[]models.Mostapps{})
	assert.Equal(t, nil, err1)
}

func TestUserIns(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := UserRepositoryProvider(dbctx, logger)
	test := models.User{
		Username: "TEST2",
		Password: "TEST2",
		Type:     "patient",
	}
	err1 := dr.UserIns(&test)
	err2 := dr.UserIns(&models.User{})
	assert.Equal(t, nil, err1)
	assert.Equal(t, errors.New("INVALID USER TYPE"), err2)

}

func TestUsersel(t *testing.T) {
	dbctx := db.GetDBCtx()
	logger := logger.ProvideLogger()
	dr := UserRepositoryProvider(dbctx, logger)
	_, err1 := dr.UserSel(&models.User{})
	_, err2 := dr.UserSel(&models.User{
		Username: "doctor1",
		Password: "doc1",
		Type:     "doctor",
	})
	assert.Equal(t, errors.New("INVALID USER"), err1)
	assert.Equal(t, nil, err2)

}
