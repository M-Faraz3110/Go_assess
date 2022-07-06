package router

import (
	"clinic/api"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Router(uc api.UserController, dc api.DoctorController, ac api.AppointmentController) {
	router := gin.Default()
	baseUrl := router.Group("/api/v1")
	//============ Calling the Controllers' SetupRoutes function ======================
	uc.SetupRoutes(baseUrl)
	dc.SetupRoutes(baseUrl)
	ac.SetupRoutes(baseUrl)
	router.Run(":8080")
}
