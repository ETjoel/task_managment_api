package main

import (
	"github.com/ETjoel/task_managment_api/database"
	"github.com/ETjoel/task_managment_api/router"
)

func main() {
	database.ConnectMongoDB()
	database.CreateEmailUniqueIndex()
	router := router.SetupRouter()
	router.Run(":8080")
}
