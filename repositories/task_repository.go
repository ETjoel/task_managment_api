package repository

import (
	"context"

	"github.com/ETjoel/task_managment_api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepositoryimpl struct {
	db         mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepositoryimpl{db: db, collection: collection}
}

func (tr *taskRepositoryimpl) GetTasks(c context.Context) ([]domain.Task, error) {
	collection := tr.db.Collection(tr.collection)
	cur, err := collection.Find(c, bson.M{})

	if err != nil {
		return nil, err
	}

	var tasks []domain.Task
	for cur.Next(c) {
		var task domain.Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr *taskRepositoryimpl) GetTaskById(c context.Context, id string) (domain.Task, error) {
	collection := tr.db.Collection(tr.collection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}

	var task domain.Task
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, nil // No document found
		}
		return domain.Task{}, err // Other errors
	}
	return task, nil
}

func (tr *taskRepositoryimpl) UpdateTask(c context.Context, id string, updatedTask domain.Task) error {
	collection := tr.db.Collection(tr.collection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.UpdateByID(c, objectId, bson.M{"$set": updatedTask})
	if err != nil {
		return err
	}

	return nil

}

func (tr *taskRepositoryimpl) DeleteTask(c context.Context, id string) error {

	collection := tr.db.Collection(tr.collection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil
}

func (tr *taskRepositoryimpl) AddTask(c context.Context, task *domain.Task) error {
	collection := tr.db.Collection(tr.collection)
	_, err := collection.InsertOne(c, task)
	if err != nil {
		return err
	}
	return nil
}
