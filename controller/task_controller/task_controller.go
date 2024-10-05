package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	task_service "github.com/ETjoel/task_managment_api/data"
	task_model "github.com/ETjoel/task_managment_api/models/task_model"
)

func GetTasks(c *gin.Context) {
	tasks, err := task_service.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTasksById(c *gin.Context) {
	id := c.Param("id")

	task, err := task_service.GetTaskById(id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	} else if task.ID != "" {
		c.IndentedJSON(http.StatusOK, task)
	} else {

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
	}
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask task_model.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if found, err := task_service.UpdateTask(id, updatedTask); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item Not Fount"})
	} else {
		c.IndentedJSON(http.StatusOK, updatedTask)
	}
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	found, err := task_service.DeleteTask(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item Not Fount"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Item Deleted"})
	}
}

func AddTask(c *gin.Context) {
	var newTask task_model.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := task_service.AddTask(newTask)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusCreated, newTask)
}
