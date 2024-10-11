package repository

import "github.com/stretchr/testify/mock"

type Bcrypte interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(exPassword []byte, password []byte) error
}

type MockBcrypt struct {
	mock.Mock
}

func (m *MockBcrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	arg := m.Called(password, cost)
	return arg.Get(0).([]byte), arg.Error(1)
}

func (m *MockBcrypt) CompareHashAndPassword(exPassword []byte, password []byte) error {
	arg := m.Called(exPassword, password)
	return arg.Error(0)
}

type MockMongoDatabase struct {
	mock.Mock
}
