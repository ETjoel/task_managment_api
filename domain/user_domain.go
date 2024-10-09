package domain

import (
	"context"
)

const (
	UsersCollection = "users"
)

type User struct {
	ID       string `bson:"_id,omitempty" json:"-"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type UserRepository interface {
	Register(c context.Context, user *User) error
	Login(c context.Context, user *User, jwtSecret string, expiryHour int) (string, error)
}

type UserUsercases interface {
	Register(c context.Context, user *User) error
	Login(c context.Context, user *User, jwtSecret string, expiryHour int) (string, error)
}
