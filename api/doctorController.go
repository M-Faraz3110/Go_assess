package api

import (
	"clinic/middle"
	"clinic/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DoctorController interface {
	GetDoctor(c *gin.Context)
	GetDoctors(c *gin.Context)
	GetAvail(c *gin.Context)
	GetSix(c *gin.Context)

	/*
		This is where you would add your controller's routes. like this:
		CreateDoctor(c *gin.Context)
		GetDoctor(c *gin.Context)
		GetDoctors(c *gin.Context)
		GetDoctorAppointments(c *gin.Context)

		Note these are just examples. You can add your own route handler functions here.
	*/
	//router

	SetupRoutes(r *gin.RouterGroup)
}

type doctorControllerImpl struct {
	svc services.DoctorService // use this to call the appropriate service functions
}

//=============================================	   Constructor	========================================================

var _ DoctorController = (*doctorControllerImpl)(nil)

func DoctorControllerProvider(s services.DoctorService) DoctorController {
	return &doctorControllerImpl{svc: s}
}

//=============================================	  	 Router Functions		========================================================

func (c *doctorControllerImpl) SetupRoutes(r *gin.RouterGroup) { //(c *appointmentControllerImpl) means that we are making a function of the type appointmentControllerImpl to access the files in your struct like so c.svc.whatever()
	r.Use(middle.Auth())
	r.GET("/doctor/:id", c.GetDoctor)
	r.GET("/doctors", c.GetDoctors)
	r.GET("/availability", c.GetAvail)
	r.GET("/sixhours", c.GetSix)
}

//============================================= 		Controller Functions	========================================================

// func (c *doctorControllerImpl) GetDoctors(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, c.svc.Doctors())
// }

func (c *doctorControllerImpl) GetDoctor(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		fmt.Println(err)
		panic(err)
	}
	if res, err := c.svc.Doctor(id); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}

func (c *doctorControllerImpl) GetDoctors(ctx *gin.Context) {
	if res, err := c.svc.Doctors(); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}

func (c *doctorControllerImpl) GetAvail(ctx *gin.Context) {
	if res, err := c.svc.Avail(); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}

func (c *doctorControllerImpl) GetSix(ctx *gin.Context) {
	if res, err := c.svc.SixHours(); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}