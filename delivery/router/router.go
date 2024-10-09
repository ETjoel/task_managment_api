package router

import (
	"time"

	"github.com/ETjoel/task_managment_api/bootstrap"
	"github.com/ETjoel/task_managment_api/delivery/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRouter sets up the router for the application

func SetupRouter(db *mongo.Database, timeout time.Duration, router *gin.Engine, env *bootstrap.Env) {

	unpGroup := router.Group("/api/v1")
	pGroup := router.Group(("api/v2"))
	pGroup.Use(middleware.AuthMiddleWare(*env))

	NewTaskRouter(db, timeout, unpGroup)
	NewUserRouter(db, timeout, unpGroup)

	NewTaskRouter(db, timeout, pGroup)
}
