package api

import (
	"clinic/auth"
	"clinic/middle"
	"clinic/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DoctorController interface {
	GetDoctor(c *gin.Context)
	GetDoctors(c *gin.Context)
	GetAvail(c *gin.Context)
	GetSix(c *gin.Context)
	GetMostApps(c *gin.Context)

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
	l   *zap.SugaredLogger
}

//=============================================	   Constructor	========================================================

var _ DoctorController = (*doctorControllerImpl)(nil)

func DoctorControllerProvider(s services.DoctorService, l *zap.SugaredLogger) DoctorController {
	return &doctorControllerImpl{svc: s, l: l}
}

//=============================================	  	 Router Functions		========================================================

func (c *doctorControllerImpl) SetupRoutes(r *gin.RouterGroup) { //(c *appointmentControllerImpl) means that we are making a function of the type appointmentControllerImpl to access the files in your struct like so c.svc.whatever()
	r.Use(middle.Auth())
	r.GET("/doctor/:id", c.GetDoctor)
	r.GET("/doctors", c.GetDoctors)
	r.GET("/availability", c.GetAvail)
	r.GET("/sixhours", c.GetSix)
	r.GET("/mostapps", c.GetMostApps)
}

//============================================= 		Controller Functions	========================================================

// func (c *doctorControllerImpl) GetDoctors(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, c.svc.Doctors())
// }

func (c *doctorControllerImpl) GetDoctor(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		// handle error
		c.l.Panic(err)
		fmt.Println(err)
		panic(err)
	}
	if res, err := c.svc.Doctor(id); err != nil {
		c.l.Panic(err)
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		c.l.Info("/doctor controller SUCCESS...")
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}

func (c *doctorControllerImpl) GetDoctors(ctx *gin.Context) {
	if res, err := c.svc.Doctors(); err != nil {
		c.l.Panic(err)
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		c.l.Info("/doctor/:id controller SUCCESS...")
		ctx.JSON(http.StatusOK, res)
	}
	// RETURN OTHER Status
}

func (c *doctorControllerImpl) GetAvail(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	claims, err := auth.GetClaims(tokenString)
	if err != nil {
		// handle error
		fmt.Println(err)
		c.l.Panic(err)
		panic(err)
	}
	fmt.Println(claims)
	if claims.Utype == "admin" {
		if res, err := c.svc.Avail(); err != nil {
			c.l.Panic(err)
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			c.l.Info("/availability controller SUCCESS...")
			ctx.JSON(http.StatusOK, res)
		}
	} else {
		ctx.JSON(http.StatusOK, "INSUFFICIENT PERMISSIONS")
	}
	// RETURN OTHER Status
}

func (c *doctorControllerImpl) GetSix(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	claims, err := auth.GetClaims(tokenString)
	if err != nil {
		// handle error
		c.l.Panic(err)
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(claims)
	if claims.Utype == "admin" {
		if res, err := c.svc.SixHours(); err != nil {
			c.l.Panic(err)
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			c.l.Info("/getsix controller SUCCESS...")
			ctx.JSON(http.StatusOK, res)
		}
	}
	// RETURN OTHER Status
}

func (c *doctorControllerImpl) GetMostApps(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	claims, err := auth.GetClaims(tokenString)
	if err != nil {
		// handle error
		c.l.Panic(err)
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(claims)
	if claims.Utype == "admin" {
		if res, err := c.svc.MostApps(); err != nil {
			c.l.Panic(err)
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			c.l.Info("/mostapps controller SUCCESS...")
			ctx.JSON(http.StatusOK, res)
		}
	}
	// RETURN OTHER Status
}
