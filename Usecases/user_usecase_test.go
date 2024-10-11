package usecases

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ETjoel/task_managment_api/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(c context.Context, user *domain.User) error {
	arg := m.Called(c, user)
	return arg.Error(0)
}

func (m *MockUserRepository) Login(c context.Context, user *domain.User, jwtSecret string, expiryHour int) (string, error) {
	arg := m.Called(c, user, jwtSecret, expiryHour)
	return arg.String(0), arg.Error(1)
}

func TestUserUsecases(t *testing.T) {
	user := &domain.User{Email: "papajones@gmail.com", Password: "password"}

	tests := []struct {
		name          string
		method        string
		inputUser     *domain.User
		expectedError error
		expectedToken string
		mockReturn    struct {
			registerError error
			loginToken    string
			loginError    error
		}
	}{
		{
			name:      "successful registration and login",
			method:    "both",
			inputUser: user,
			mockReturn: struct {
				registerError error
				loginToken    string
				loginError    error
			}{registerError: nil, loginToken: "some token", loginError: nil},
			expectedError: nil,
			expectedToken: "some token",
		},
		{
			name:      "failed registration",
			method:    "register",
			inputUser: user,
			mockReturn: struct {
				registerError error
				loginToken    string
				loginError    error
			}{registerError: errors.New("register error"), loginToken: "", loginError: nil},
			expectedError: errors.New("register error"),
			expectedToken: "",
		},
		{
			name:      "failed login",
			method:    "login",
			inputUser: user,
			mockReturn: struct {
				registerError error
				loginToken    string
				loginError    error
			}{registerError: nil, loginToken: "", loginError: errors.New("login error")},
			expectedError: errors.New("login error"),
			expectedToken: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := new(MockUserRepository)
			userUsercases := NewUserUsecases(mockUserRepository, time.Duration(2)*time.Second)

			if tt.method == "both" || tt.method == "register" {
				mockUserRepository.On("Register", mock.Anything, tt.inputUser).Return(tt.mockReturn.registerError)
			}

			if tt.method == "both" || tt.method == "login" {
				mockUserRepository.On("Login", mock.Anything, tt.inputUser, mock.Anything, mock.Anything).Return(tt.mockReturn.loginToken, tt.mockReturn.loginError)
			}

			if tt.method == "both" || tt.method == "register" {
				err := userUsercases.Register(context.Background(), tt.inputUser)
				if tt.expectedError != nil {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}

			if tt.method == "both" || tt.method == "login" {
				token, err := userUsercases.Login(context.Background(), tt.inputUser, "some_token", 10)
				assert.Equal(t, tt.expectedToken, token)
				if tt.expectedError != nil {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}

			mockUserRepository.AssertExpectations(t)
		})
	}
}
