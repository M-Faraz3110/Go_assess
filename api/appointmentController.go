package api

import (
	"clinic/auth"
	"clinic/middle"
	"clinic/models"
	"clinic/services"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	l   *zap.SugaredLogger
}

//=============================================	   Constructor 		========================================================

var _ AppointmentController = (*appointmentControllerImpl)(nil)

func AppointmentControllerProvider(s services.AppointmentService, l *zap.SugaredLogger) AppointmentController {
	return &appointmentControllerImpl{svc: s, l: l}
}

//=============================================	  	 Router Functions		========================================================

func (c *appointmentControllerImpl) SetupRoutes(r *gin.RouterGroup) { //(c *appointmentControllerImpl) means that we are making a function of the type appointmentControllerImpl to access the files in your struct like so c.svc.whatever()
	r.Use(middle.Auth())
	r.GET("/doctor/:id/slots", c.GetSlots)
	r.GET("/app/:id", c.GetAppointment)
	r.POST("/book", c.CreateAppointment)
	r.DELETE("/cancel/:id", c.CancelAppointment)
	r.GET("/history/:id", c.History)
	//c.JSON()
}

//============================================= 		Controller Functions	========================================================

// INTERFACE FUCNTIONS

func (c *appointmentControllerImpl) GetSlots(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		c.l.Panic(err)
		fmt.Println(err)
		panic(err)
	}
	if res, err := c.svc.Slots(id); err != nil {
		c.l.Panic(err)
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		c.l.Info("/slots controller SUCCESS...")
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}

func (c *appointmentControllerImpl) CreateAppointment(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	claims, err := auth.GetClaims(tokenString)
	if err != nil {
		// handle error
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(claims)
	if claims.Utype == "patient" {
		var request models.TimeReq
		if err := ctx.BindJSON(&request); err != nil {
			// DO SOMETHING WITH THE ERROR
			fmt.Println(err)
			panic(err)
		}
		layout := "02 Jan 06 15:04"
		stime, err := time.Parse(layout, request.Start_time)
		if err != nil {
			// handle error
			fmt.Println(err)
			panic(err)
		}
		etime, err := time.Parse(layout, request.End_time)
		if err != nil {
			// handle error
			fmt.Println(err)
			panic(err)
		}
		app := models.Appointment{
			DocId:      request.DocId,
			PatId:      request.PatId,
			Start_time: stime,
			End_time:   etime,
		}
		if request.PatId != claims.ID {
			c.l.Panic(err)
			ctx.JSON(http.StatusOK, "Incorrect Pat ID")
		} else if err := c.svc.Book(app); err != nil {
			c.l.Panic(err)
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			c.l.Info("/book controller SUCCESS...")
			ctx.JSON(http.StatusOK, "SUCCESS")
		}

	} else {
		ctx.JSON(http.StatusOK, "INSUFFICIENT PERMISSIONS")
	}

	// RETURN OTHER Status
}

func (c *appointmentControllerImpl) CancelAppointment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		c.l.Panic(err)
		fmt.Println(err)
		panic(err)
	}
	if err := c.svc.Cancel(id); err != nil {
		c.l.Panic(err)
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		c.l.Info("/cancel controller SUCCESS...")
		ctx.JSON(http.StatusOK, "SUCCESS")
	}
	// RETURN OTHER Status
}

func (c *appointmentControllerImpl) GetAppointment(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		c.l.Panic(err)
		fmt.Println(err)
		panic(err)
	}
	if res, err := c.svc.App(id); err != nil {
		c.l.Panic(err)
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		c.l.Info("/appointment controller SUCCESS...")
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}

func (c *appointmentControllerImpl) History(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		c.l.Panic(err)
		fmt.Println(err)
		panic(err)
	}
	if res, err := c.svc.History(id); err != nil {
		c.l.Panic(err)
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		c.l.Info("/history controller SUCCESS...")
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}
