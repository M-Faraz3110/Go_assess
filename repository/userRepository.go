package repository

import (
	"clinic/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	UserIns(user *models.Register) error
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
func (c *userrepositoryImpl) UserIns(user *models.Register) error {
	cmd := fmt.Sprintf("INSERT INTO %s (username, password) values ('%s', '%s')", user.Type, user.Username, user.Password)
	_, err := c.db.Exec(cmd)
	return err
}
