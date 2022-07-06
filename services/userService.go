package services

import (
	"clinic/models"
	"clinic/repository"
)

type UserService interface {
	Register(models.Register) error
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

func (c *userServiceImpl) Register(request models.Register) error {

	return c.ur.UserIns(&request)
}
