package usecases

import (
	"context"
	"time"

	"github.com/ETjoel/task_managment_api/domain"
)

type userUsercases struct {
	ur             domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecases(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsercases {
	return &userUsercases{
		ur:             userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsercases) Register(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.ur.Register(ctx, user)
}

func (uu *userUsercases) Login(c context.Context, user *domain.User, jwtSecret string, expiryHour int) (string, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()

	return uu.ur.Login(ctx, user, jwtSecret, expiryHour)
}
