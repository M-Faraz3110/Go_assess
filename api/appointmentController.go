package api

import (
	"clinic/middle"
	"clinic/models"
	"clinic/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentController interface {
	GetSlots(c *gin.Context)
	CreateAppointment(c *gin.Context)
	CancelAppointment(c *gin.Context)
	GetAppointment(c *gin.Context)
	/*
		This is where you would add your controller's routes. like this:
		CreateAppointment(c *gin.Context)
		CancelAppointment(c *gin.Context)
		GetAppointment(c *gin.Context)
		GetAppointments(c *gin.Context)

		Note these are just examples. You can add your own route handler functions here.
	*/
	//router
	SetupRoutes(r *gin.RouterGroup)
}

type appointmentControllerImpl struct {
	svc services.AppointmentService // use this to call the appropriate service functions
}

//=============================================	   Constructor 		========================================================

var _ AppointmentController = (*appointmentControllerImpl)(nil)

func AppointmentControllerProvider(s services.AppointmentService) AppointmentController {
	return &appointmentControllerImpl{svc: s}
}

//=============================================	  	 Router Functions		========================================================

func (c *appointmentControllerImpl) SetupRoutes(r *gin.RouterGroup) { //(c *appointmentControllerImpl) means that we are making a function of the type appointmentControllerImpl to access the files in your struct like so c.svc.whatever()
	r.Use(middle.Auth())
	r.GET("/doctor/:id/slots", c.GetSlots)
	r.GET("/app/:id", c.GetSlots)
	r.POST("/book", c.CreateAppointment)
	r.DELETE("/cancel/:id", c.CancelAppointment)
	// r.GET("/history/:id", c.History)
	//c.JSON()
}

//============================================= 		Controller Functions	========================================================

// INTERFACE FUCNTIONS

func (c *appointmentControllerImpl) GetSlots(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		fmt.Println(err)
		panic(err)
	}
	if res, err := c.svc.Slots(id); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}

func (c *appointmentControllerImpl) CreateAppointment(ctx *gin.Context) {
	var request models.Appointment
	if err := ctx.BindJSON(&request); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Println(err)
		panic(err)
	}
	if err := c.svc.Book(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, "SUCCESS")
	}
	// RETURN OTHER Status
}

func (c *appointmentControllerImpl) CancelAppointment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		fmt.Println(err)
		panic(err)
	}
	if err := c.svc.Cancel(id); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, "SUCCESS")
	}
	// RETURN OTHER Status
}

func (c *appointmentControllerImpl) GetAppointment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		fmt.Println(err)
		panic(err)
	}
	if res, err := c.svc.App(id); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}

func (c *appointmentControllerImpl) History(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		fmt.Println(err)
		panic(err)
	}
	if res, err := c.svc.History(id); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}
