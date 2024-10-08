package main

import (
	"time"

	"github.com/ETjoel/task_managment_api/bootstrap"
	r "github.com/ETjoel/task_managment_api/delivery/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()
	app.CreateEmailUniqueIndex()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	db := app.Client.Database(app.Env.DBName)
	defer app.CloseDatabase()

	router := gin.Default()
	router.Use(cors.Default())

	r.SetupRouter(db, timeout, router)
	router.Run(app.Env.ServerAddress)
}
