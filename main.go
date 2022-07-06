package main

import (
	"clinic/api"
	"clinic/db"
	"clinic/repository"
	"clinic/router"
	"clinic/services"
)

func main() {
	dbctx := db.GetDBCtx("user=postgres dbname=postgres sslmode=disable password=Salmon123")
	ur := repository.UserRepositoryProvider(dbctx)
	dr := repository.DoctorRepositoryProvider(dbctx)
	ar := repository.AppointmentRepositoryProvider(dbctx)
	us := services.UserServiceProvider(ur)
	ds := services.DoctorServiceProvider(dr)
	as := services.AppointmentServiceProvider(ar)
	uc := api.UserControllerProvider(us)
	dc := api.DoctorControllerProvider(ds)
	ac := api.AppointmentControllerProvider(as)
	router.Router(uc, dc, ac)
}
