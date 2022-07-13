package api

import (
	"clinic/models"
	"clinic/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController interface {
	Register(c *gin.Context)
	Login(ctx *gin.Context)
	/*
		This is where you would add your controller's routes. like this:
		CreateUser(c *gin.Context)
		GetUser(c *gin.Context)
		GetUsers(c *gin.Context)
		GetUserAppointments(c *gin.Context)
		UpdateUser(c *gin.Context)

		Note these are just examples. You can add your own route handler functions here.
	*/
	//router
	SetupRoutes(r *gin.RouterGroup)
}

type userControllerImpl struct {
	svc services.UserService // use this to call the appropriate service functions
	l   *zap.SugaredLogger
}

//=============================================	   Constructor 	========================================================

var _ UserController = (*userControllerImpl)(nil)

func UserControllerProvider(s services.UserService, l *zap.SugaredLogger) UserController {
	return &userControllerImpl{svc: s, l: l}
}

//=============================================	  	 Router Functions		========================================================

func (c *userControllerImpl) SetupRoutes(r *gin.RouterGroup) { //(c *appointmentControllerImpl) means that we are making a function of the type appointmentControllerImpl to access the files in your struct like so c.svc.whatever()
	r.POST("/login", c.Login)
	r.POST("/register", c.Register)
}

//============================================= 		Controller Functions	========================================================

func (c *userControllerImpl) Register(ctx *gin.Context) {
	var request models.User
	if err := ctx.BindJSON(&request); err != nil {
		// DO SOMETHING WITH THE ERROR
		c.l.Panic(err)
		fmt.Println(err)
		panic(err)
	}
	if err := c.svc.Register(request); err != nil {
		c.l.Panic(err)
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		c.l.Info("/register controller SUCCESS...")
		ctx.JSON(http.StatusOK, "SUCCESS")
	}

}

func (c *userControllerImpl) Login(ctx *gin.Context) {
	var request models.User
	if err := ctx.BindJSON(&request); err != nil {
		// DO SOMETHING WITH THE ERROR
		c.l.Panic(err)
		fmt.Println(err)
		panic(err)
	}
	if token, err := c.svc.Login(request); err != nil {
		c.l.Info("/login controller SUCCESS...")
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.l.Info("/login controller SUCCESS...")
		ctx.JSON(http.StatusOK, gin.H{"token:": token})
	}

}
