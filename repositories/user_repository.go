package repository

import (
	"github.com/ETjoel/task_managment_api/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryImpl struct {
	db mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) *domain.UserRepository {
	return &userRepositoryImpl(db: db, collection: collection)
}

