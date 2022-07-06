package api

import (
	"clinic/models"
	"clinic/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
}

//=============================================	   Constructor 	========================================================

var _ UserController = (*userControllerImpl)(nil)

func UserControllerProvider(s services.UserService) UserController {
	return &userControllerImpl{svc: s}
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
		fmt.Println(err)
		panic(err)
	}
	if err := c.svc.Register(request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, "SUCCESS")
	}

}

func (c *userControllerImpl) Login(ctx *gin.Context) {
	var request models.User
	if err := ctx.BindJSON(&request); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Println(err)
		panic(err)
	}
	if token, err := c.svc.Login(request); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"token:": token})
	}

}
