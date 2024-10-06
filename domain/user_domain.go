package domain

import (
	"context"
)

type User struct {
	ID       string `bson:"_id,omitempty" json:"-"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}

type UserRepository interface {
	Register(c context.Context, user *User) error
	Login(c context.Context, user *User) (string, error)
}

type UserUsercases interface {
	Register(c context.Context, user *User) error
	Login(c context.Context, user *User) (string, error)
}
