package controller

import (
	"net/http"

	"github.com/ETjoel/task_managment_api/domain"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase domain.TaskUseCases
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetTasks(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")

	task, err := tc.TaskUsecase.GetTaskById(c, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) GetTasksById(c *gin.Context) {
	id := c.Param("id")

	tasks, err := tc.TaskUsecase.GetTaskById(c, id)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)

}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask domain.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := tc.TaskUsecase.UpdateTask(c, id, updatedTask); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := tc.TaskUsecase.DeleteTask(c, id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Item Deleted"})

}

func (tc *TaskController) AddTask(c *gin.Context) {
	var newTask domain.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.TaskUsecase.AddTask(c, &newTask)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newTask)
}
