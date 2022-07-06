package middle

import (
	"clinic/auth"
	"clinic/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Utype    string `json:"type"`
}

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       string `json:"id"`
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		fmt.Print(tokenString[7:])
		err := auth.ValidateToken(tokenString[7:])
		fmt.Println(err)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
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
	fmt.Println(request.Type)
	cmd := fmt.Sprintf("SELECT username, password, id FROM users WHERE username = '%s' and password = '%s'", request.Username, request.Password)
	err := db.Get(&user, cmd)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	tokenString, err := auth.GenerateJWT(user.Username, user.Id, request.Type)
	if err != nil {
		// context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// context.Abort()
		return "", err
	}
	return tokenString, err
}
