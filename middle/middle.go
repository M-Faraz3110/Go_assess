package middle

import (
	"clinic/auth"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Utype    string `json:"utype"`
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
		err := auth.ValidateToken(tokenString)
		fmt.Println(err)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=Salmon123")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	user := user{}
	fmt.Println(request.Utype)
	cmd := fmt.Sprintf("SELECT username, password, id FROM %s WHERE username = '%s' and password = '%s'", request.Utype, request.Username, request.Password)
	err = db.Get(&user, cmd)
	if err != nil {
		fmt.Println(err)
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username, user.Id, request.Utype)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
