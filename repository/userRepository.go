package repository

import (
	"clinic/middle"
	"clinic/models"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserRepository interface {
	UserIns(user *models.User) error
	UserSel(user *models.User) (string, error)
}

type userrepositoryImpl struct {
	db *sqlx.DB
	l  *zap.SugaredLogger
}

//=============================================	   Constructor and DI		========================================================
var _ UserRepository = (*userrepositoryImpl)(nil)

func UserRepositoryProvider(db *sqlx.DB, l *zap.SugaredLogger) UserRepository {
	return &userrepositoryImpl{db: db, l: l}
}

//=============================================	 	SVC Functions		========================================================

func (c *userrepositoryImpl) UserIns(user *models.User) error {
	hashpass, err := middle.HashPassword(user.Password)
	if err != nil {
		// handle error
		c.l.Info("invalid password...")
		fmt.Println(err)
		return err
	}
	cmd := fmt.Sprintf("INSERT INTO users (username, password, user_type) values ('%s', '%s', '%s')", user.Username, hashpass, user.Type)
	_, err = c.db.Exec(cmd)
	c.l.Info("register repo SUCCESS...")
	return err
}

func (c *userrepositoryImpl) UserSel(user *models.User) (string, error) {
	token, err := middle.GenerateToken(user, c.db)
	c.l.Info("login repo SUCCESS...")
	return token, err
}
