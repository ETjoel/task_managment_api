package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	usecases "github.com/ETjoel/task_managment_api/Usecases"
	controller "github.com/ETjoel/task_managment_api/delivery/controller/task_controller"
	repository "github.com/ETjoel/task_managment_api/repositories"
)

// SetupRouter sets up the router for the application

func SetupRouter(db *mongo.Database, timeout time.Duration, router *gin.Engine) {
	tr := repository.NewTaskRepository(*db, "tasks")
	tc := &controller.TaskController{
		TaskUsecase: usecases.NewTaskUsecases(tr, timeout),
	}

	group := router.Group("/api/v1")

	group.GET("/tasks", tc.GetTasks)
	group.GET("/tasks/:id", tc.GetTasksById)
	group.PUT("/tasks/:id", tc.UpdateTask)
	group.DELETE("/tasks/:id", tc.DeleteTask)
	group.POST("/tasks", tc.AddTask)

}
