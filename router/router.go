package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	authcontroller "github.com/ETjoel/task_managment_api/controller/auth_controller"
	controller "github.com/ETjoel/task_managment_api/controller/task_controller"
	"github.com/ETjoel/task_managment_api/middleware"
)

// SetupRouter sets up the router for the application

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/tasks", controller.GetTasks)
	router.GET("/tasks/:id", controller.GetTasksById)
	router.PUT("/tasks/:id", controller.UpdateTask)
	router.DELETE("/tasks/:id", controller.DeleteTask)
	router.POST("/tasks", controller.AddTask)

	router.POST("/register", authcontroller.Register)
	router.POST("/login", authcontroller.Login)

	router.GET("/v2/tasks", middleware.AuthMiddleWare(), controller.GetTasks)
	router.POST("/v2/tasks", middleware.AuthMiddleWare(), controller.AddTask)
	// router.GET("/me", authcontroller.GetUser)
	return router
}
