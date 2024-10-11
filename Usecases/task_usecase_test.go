package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/ETjoel/task_managment_api/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetTasks(c context.Context) ([]domain.Task, error) {
	arg := m.Called(c)
	return arg.Get(0).([]domain.Task), arg.Error(1)
}
func (m *MockTaskRepository) GetTaskById(c context.Context, id string) (domain.Task, error) {
	arg := m.Called(c, id)
	return arg.Get(0).(domain.Task), arg.Error(1)
}

func (m *MockTaskRepository) UpdateTask(c context.Context, id string, updatedTask domain.Task) error {
	arg := m.Called(c, id, updatedTask)
	return arg.Error(0)
}

func (m *MockTaskRepository) DeleteTask(c context.Context, id string) error {
	arg := m.Called(c, id)
	return arg.Error(0)
}

func (m *MockTaskRepository) AddTask(c context.Context, task *domain.Task) error {
	arg := m.Called(c, task)
	return arg.Error(0)
}

func TestTaskUsecases(t *testing.T) {
	task := &domain.Task{
		Title:       "Task 1",
		Description: "Task 1 description",
		DueDate:     time.Now(),
		Status:      "In Progress",
	}

	tests := []struct {
		name          string
		method        string
		inputUser     *domain.Task
		expectedError error
		expectedTask  domain.Task
		expectedTasks []domain.Task
		mockReturn    struct {
			tasks []domain.Task
			task  domain.Task
			err   error
		}
	}{
		{
			name:          "successful get tasks",
			method:        "GetTasks",
			inputUser:     nil,
			expectedError: nil,
			expectedTask:  domain.Task{},
			expectedTasks: []domain.Task{*task},
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: []domain.Task{*task},
				task:  domain.Task{},
				err:   nil,
			},
		},
		{
			name:          "Failure get tasks",
			method:        "GetTasks",
			inputUser:     nil,
			expectedError: assert.AnError,
			expectedTask:  domain.Task{},
			expectedTasks: nil,
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: nil,
				task:  domain.Task{},
				err:   assert.AnError,
			},
		},
		{
			name:          "successful GetTaskById",
			method:        "GetTaskById",
			inputUser:     task,
			expectedError: nil,
			expectedTask:  *task,
			expectedTasks: nil,
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: nil,
				task:  *task,
				err:   nil,
			},
		},
		{
			name:          "failure GetTaskById",
			method:        "GetTaskById",
			inputUser:     task,
			expectedError: assert.AnError,
			expectedTask:  domain.Task{},
			expectedTasks: nil,
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: nil,
				task:  domain.Task{},
				err:   assert.AnError,
			},
		},
		{
			name:          "successful UpdateTask",
			method:        "UpdateTask",
			inputUser:     task,
			expectedError: nil,
			expectedTask:  domain.Task{},
			expectedTasks: nil,
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: nil,
				task:  domain.Task{},
				err:   nil,
			},
		},
		{
			name:          "failure UpdateTask",
			method:        "UpdateTask",
			inputUser:     task,
			expectedError: assert.AnError,
			expectedTask:  *task,
			expectedTasks: nil,
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: nil,
				task:  domain.Task{},
				err:   assert.AnError,
			},
		},
		{
			name:          "successful DeleteTask",
			method:        "DeleteTask",
			inputUser:     task,
			expectedError: nil,
			expectedTask:  *task,
			expectedTasks: nil,
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: nil,
				task:  domain.Task{},
				err:   nil,
			},
		},
		{
			name:          "failure DeleteTask",
			method:        "DeleteTask",
			inputUser:     task,
			expectedError: assert.AnError,
			expectedTask:  domain.Task{},
			expectedTasks: nil,
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: nil,
				task:  domain.Task{},
				err:   assert.AnError,
			},
		},
		{
			name:          "successful AddTask",
			method:        "AddTask",
			inputUser:     task,
			expectedError: nil,
			expectedTask:  *task,
			expectedTasks: nil,
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: nil,
				task:  domain.Task{},
				err:   nil,
			},
		},
		{
			name:          "failure AddTask",
			method:        "AddTask",
			inputUser:     task,
			expectedError: assert.AnError,
			expectedTask:  domain.Task{},
			expectedTasks: nil,
			mockReturn: struct {
				tasks []domain.Task
				task  domain.Task
				err   error
			}{
				tasks: nil,
				task:  domain.Task{},
				err:   assert.AnError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockTaskRepository := new(MockTaskRepository)
			taskUsercases := NewTaskUsecases(mockTaskRepository, time.Duration(2)*time.Second)

			if tt.method == "GetTasks" {
				mockTaskRepository.On("GetTasks", mock.Anything).Return(tt.mockReturn.tasks, tt.mockReturn.err)
			} else if tt.method == "GetTaskById" {
				mockTaskRepository.On("GetTaskById", mock.Anything, mock.Anything).Return(tt.mockReturn.task, tt.mockReturn.err)
			} else if tt.method == "UpdateTask" {
				mockTaskRepository.On("UpdateTask", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockReturn.err)
			} else if tt.method == "DeleteTask" {
				mockTaskRepository.On("DeleteTask", mock.Anything, mock.Anything).Return(tt.mockReturn.err)
			} else if tt.method == "AddTask" {
				mockTaskRepository.On("AddTask", mock.Anything, mock.Anything).Return(tt.mockReturn.err)
			}

			if tt.method == "GetTasks" {
				result, err := taskUsercases.GetTasks(context.Background())

				assert.Equal(t, tt.expectedError, err)
				assert.Equal(t, tt.expectedTasks, result)
			} else if tt.method == "GetTaskById" {
				result, err := taskUsercases.GetTaskById(context.Background(), "1")

				assert.Equal(t, tt.expectedError, err)
				assert.Equal(t, tt.expectedTask, result)
			} else if tt.method == "UpdateTask" {
				err := taskUsercases.UpdateTask(context.Background(), "1", *task)

				assert.Equal(t, tt.expectedError, err)
			} else if tt.method == "DeleteTask" {
				err := taskUsercases.DeleteTask(context.Background(), "1")

				assert.Equal(t, tt.expectedError, err)
			} else if tt.method == "AddTask" {
				err := taskUsercases.AddTask(context.Background(), task)

				assert.Equal(t, tt.expectedError, err)
			}

			mockTaskRepository.AssertExpectations(t)

		})
	}
}
