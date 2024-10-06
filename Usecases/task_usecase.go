package usecases

import (
	"context"
	"time"

	"github.com/ETjoel/task_managment_api/domain"
)

type taskUsecasesImpl struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecases(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUseCases {
	return &taskUsecasesImpl{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecasesImpl) GetTasks(c context.Context) ([]domain.Task, error) {
	cxt, cancel := context.WithTimeout(c, tu.contextTimeout)

	defer cancel()

	return tu.taskRepository.GetTasks(cxt)
}

func (tu *taskUsecasesImpl) GetTaskById(c context.Context, id string) (domain.Task, error) {
	cxt, cancel := context.WithTimeout(c, tu.contextTimeout)

	defer cancel()

	return tu.taskRepository.GetTaskById(cxt, id)
}

func (tu *taskUsecasesImpl) UpdateTask(c context.Context, id string, updatedTask domain.Task) error {
	cxt, cancel := context.WithTimeout(c, tu.contextTimeout)

	defer cancel()

	return tu.taskRepository.UpdateTask(cxt, id, updatedTask)
}

func (tu *taskUsecasesImpl) DeleteTask(c context.Context, id string) error {
	cxt, cancel := context.WithTimeout(c, tu.contextTimeout)

	defer cancel()

	return tu.taskRepository.DeleteTask(cxt, id)
}

func (tu *taskUsecasesImpl) AddTask(c context.Context, task *domain.Task) error {
	cxt, cancel := context.WithTimeout(c, tu.contextTimeout)

	defer cancel()

	return tu.taskRepository.AddTask(cxt, task)
}
