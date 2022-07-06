package services

import (
	"clinic/models"
	"clinic/repository"
)

type UserService interface {
	Register(models.User) error
	Login(request models.User) (string, error)
}

type userServiceImpl struct {
	ur repository.UserRepository
}

//=============================================	   Constructor 	========================================================
var _ UserService = (*userServiceImpl)(nil)

func UserServiceProvider(ur repository.UserRepository) UserService {
	return &userServiceImpl{ur: ur}
}

//=============================================	 	SVC Functions		========================================================

func (c *userServiceImpl) Register(request models.User) error {

	return c.ur.UserIns(&request)
}

func (c *userServiceImpl) Login(request models.User) (string, error) {

	return c.ur.UserSel(&request)
}
