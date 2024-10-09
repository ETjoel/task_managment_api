package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRouter sets up the router for the application

func SetupRouter(db *mongo.Database, timeout time.Duration, router *gin.Engine) {

	taskGroup := router.Group("/api/v1")

	NewTaskRouter(db, timeout, taskGroup)

}
