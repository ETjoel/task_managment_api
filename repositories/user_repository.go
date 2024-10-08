package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/ETjoel/task_managment_api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userRepositoryImpl struct {
	db         mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepositoryImpl{db: db, collection: collection}
}

func (ur *userRepositoryImpl) Register(c context.Context, user domain.User) error {
	hashedPassword, err := bycrypt.GenerateFromPassword([]byte(user.Password))

	if err != nil {
		return errors.New("internal server error: " + err.Error())
	}
	newUser := domain.User{Email: strings.ToLower(user.Email), Password: string(hashedPassword)}

	collection := ur.db.Collection(ur.collection)
	_, err = collection.InsertOne(c, newUser)

	if err != nil && mongo.IsDuplicateKeyError(err) {
		return errors.New("email already in use")
	} else if err != nil {
		return errors.New("internal server error")
	} else {
		return nil
	}
}

func (ur *userRepositoryImpl) Login(c context.Context, user domain.User) (string, error) {
	var exitingUser domain.User
	collection := ur.db.Collection(ur.collection)

	user.Email = strings.ToLower(user.Email)

	if err := collection.FindOne(c, bson.M{"email": user.Email}).Decode(&exitingUser); err != nil {
		return "", errors.New("user not found or invalid credentials: " + err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(exitingUser.Password), []byte(user.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

}
