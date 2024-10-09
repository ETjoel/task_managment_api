package domain

import (
	"context"
	"time"
)

const (
	TaskCollection = "tasks"
)

type Task struct {
	ID          string    `bson:"_id" json:"_id"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	DueDate     time.Time `bson:"due_date" json:"due_date"`
	Status      string    `bson:"status" json:"status"`
}

type TaskRepository interface {
	GetTasks(c context.Context) ([]Task, error)
	GetTaskById(c context.Context, id string) (Task, error)
	UpdateTask(c context.Context, id string, updatedTask Task) error
	DeleteTask(c context.Context, id string) error
	AddTask(c context.Context, task *Task) error
}

type TaskUseCases interface {
	GetTasks(c context.Context) ([]Task, error)
	GetTaskById(c context.Context, id string) (Task, error)
	UpdateTask(c context.Context, id string, updatedTask Task) error
	DeleteTask(c context.Context, id string) error
	AddTask(c context.Context, task *Task) error
}
