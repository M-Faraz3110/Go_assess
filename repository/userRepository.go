package repository

import (
	"clinic/middle"
	"clinic/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	UserIns(user *models.User) error
	UserSel(user *models.User) (string, error)
}

type userrepositoryImpl struct {
	db *sqlx.DB
}

//=============================================	   Constructor and DI		========================================================
var _ UserRepository = (*userrepositoryImpl)(nil)

func UserRepositoryProvider(db *sqlx.DB) UserRepository {
	return &userrepositoryImpl{db: db}
}

//=============================================	 	SVC Functions		========================================================

func (c *userrepositoryImpl) UserIns(user *models.User) error {
	hashpass, err := middle.HashPassword(user.Password)
	if err != nil {
		// handle error
		fmt.Println(err)
		return err
	}
	fmt.Println(user.Username, hashpass, user.Type)
	cmd := fmt.Sprintf("INSERT INTO users (username, password, user_type) values ('%s', '%s', '%s')", user.Username, hashpass, user.Type)
	_, err = c.db.Exec(cmd)
	return err
}

func (c *userrepositoryImpl) UserSel(user *models.User) (string, error) {
	token, err := middle.GenerateToken(user, c.db)
	return token, err
}
