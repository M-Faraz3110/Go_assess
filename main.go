package main

import (
	"clinic/api"
	"clinic/db"
	"clinic/logger"
	"clinic/repository"
	"clinic/router"
	"clinic/services"

	"go.uber.org/fx"
)

// func main() {
// 	dbctx := db.GetDBCtx("user=postgres dbname=postgres sslmode=disable password=Salmon123")
// 	logger := logger.ProvideLogger()
// 	ur := repository.UserRepositoryProvider(dbctx, logger)
// 	dr := repository.DoctorRepositoryProvider(dbctx, logger)
// 	ar := repository.AppointmentRepositoryProvider(dbctx, logger)
// 	us := services.UserServiceProvider(ur, logger)
// 	ds := services.DoctorServiceProvider(dr, logger)
// 	as := services.AppointmentServiceProvider(ar, logger)
// 	uc := api.UserControllerProvider(us, logger)
// 	dc := api.DoctorControllerProvider(ds, logger)
// 	ac := api.AppointmentControllerProvider(as, logger)
// 	router.Router(uc, dc, ac)
// }

func main() {
	app := fx.New(
		fx.Provide(db.GetDBCtx),
		fx.Provide(logger.ProvideLogger),
		fx.Provide(repository.UserRepositoryProvider),
		fx.Provide(repository.DoctorRepositoryProvider),
		fx.Provide(repository.AppointmentRepositoryProvider),
		fx.Provide(services.UserServiceProvider),
		fx.Provide(services.DoctorServiceProvider),
		fx.Provide(services.AppointmentServiceProvider),
		fx.Provide(api.UserControllerProvider),
		fx.Provide(api.DoctorControllerProvider),
		fx.Provide(api.AppointmentControllerProvider),
		fx.Invoke(router.Router),
	)
	app.Run()

	// dbctx := db.GetDBCtx("")
	// logger := logger.ProvideLogger()
	// ur := repository.UserRepositoryProvider(dbctx, logger)
	// dr := repository.DoctorRepositoryProvider(dbctx, logger)
	// ar := repository.AppointmentRepositoryProvider(dbctx, logger)
	// us := services.UserServiceProvider(ur, logger)
	// ds := services.DoctorServiceProvider(dr, logger)
	// as := services.AppointmentServiceProvider(ar, logger)
	// uc := api.UserControllerProvider(us, logger)
	// dc := api.DoctorControllerProvider(ds, logger)
	// ac := api.AppointmentControllerProvider(as, logger)
	// router.Router(uc, dc, ac)
}
