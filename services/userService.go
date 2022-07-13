package services

import (
	"clinic/models"
	"clinic/repository"

	"go.uber.org/zap"
)

type UserService interface {
	Register(models.User) error
	Login(request models.User) (string, error)
}

type userServiceImpl struct {
	ur repository.UserRepository
	l  *zap.SugaredLogger
}

//=============================================	   Constructor 	========================================================
var _ UserService = (*userServiceImpl)(nil)

func UserServiceProvider(ur repository.UserRepository, l *zap.SugaredLogger) UserService {
	return &userServiceImpl{ur: ur, l: l}
}

//=============================================	 	SVC Functions		========================================================

func (c *userServiceImpl) Register(request models.User) error {
	c.l.Info("/register service SUCCESS...")
	return c.ur.UserIns(&request)
}

func (c *userServiceImpl) Login(request models.User) (string, error) {
	c.l.Info("/login service SUCCESS...")
	return c.ur.UserSel(&request)
}
