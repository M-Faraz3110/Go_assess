package middle

import (
	"clinic/auth"
	"clinic/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Utype    string `json:"type"`
}

type user struct {
	Username string `json:"username"`
	Type     string `json:"type"`
	Password string `json:"password"`
	Id       int    `json:"id"`
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := auth.ValidateToken(tokenString[7:])
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateToken(request *models.User, db *sqlx.DB) (string, error) {
	// var request TokenRequest
	// if err := context.ShouldBindJSON(&request); err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	context.Abort()
	// 	return "", err
	// }
	// check if email exists and password is correct
	user := user{}
	cmd := fmt.Sprintf("SELECT username, password, id, user_type as type FROM users WHERE username = '%s' and user_type = '%s'", request.Username, request.Type)
	err := db.Get(&user, cmd)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("INVALID USER")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	tokenString, err := auth.GenerateJWT(user.Username, user.Id, user.Type)
	if err != nil {
		// context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// context.Abort()
		return "", err
	}
	return tokenString, err
}
