package controller

import (
	"fmt"
	"log"

	"github.com/ETjoel/task_managment_api/data"
	user_model "github.com/ETjoel/task_managment_api/models/user_model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user user_model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		fmt.Print("bind: error % s", err.Error())
		return
	}

	log.Printf("user passowrd: %s, user email: %s", user.Password, user.Email)
	if err := data.RegisterUser(user); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User Registered Successfully"})

}

func Login(c *gin.Context) {
	var user user_model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := data.LoginUser(user)

	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User Successfully login", "token": jwtToken})
}
