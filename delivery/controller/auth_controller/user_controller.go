package controller

import (
	"fmt"

	"github.com/ETjoel/task_managment_api/bootstrap"
	"github.com/ETjoel/task_managment_api/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecases domain.UserUsercases
	Env          bootstrap.Env
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		fmt.Print("bind: error % s", err.Error())
		return
	}

	if err := uc.UserUsecases.Register(c, &user); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User Registered Successfully"})

}

func (uc *UserController) Login(c *gin.Context) {

	jwtSecret := uc.Env.AccessTokenSecret
	expireHour := uc.Env.AccessTokenExpiryHour
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := uc.UserUsecases.Login(c, &user, jwtSecret, expireHour)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User Successfully login", "token": jwtToken})
}
