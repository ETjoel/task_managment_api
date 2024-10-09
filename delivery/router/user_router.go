package router

import (
	"time"

	usecases "github.com/ETjoel/task_managment_api/Usecases"
	controller "github.com/ETjoel/task_managment_api/delivery/controller/auth_controller"
	"github.com/ETjoel/task_managment_api/domain"
	repository "github.com/ETjoel/task_managment_api/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRouter(db *mongo.Database, timeout time.Duration, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(*db, domain.UsersCollection)
	uu := usecases.NewUserUsecases(ur, timeout)
	uc := &controller.UserController{
		UserUsecases: uu,
	}

	group.POST("/login", uc.Login)
	group.POST("/register", uc.Register)
}
