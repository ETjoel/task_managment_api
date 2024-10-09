package repository

import (
	"context"
	"errors"

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
		return nil, errors.New("tasks not found!: " + err.Error())
	}

	var tasks []domain.Task
	for cur.Next(c) {
		var task domain.Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, errors.New("internal server error: " + err.Error())
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (tr *taskRepositoryimpl) GetTaskById(c context.Context, id string) (domain.Task, error) {
	collection := tr.db.Collection(tr.collection)

	var task domain.Task
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, errors.New("Item not found: " + err.Error())
		}
		return domain.Task{}, err
	}
	return task, nil
}

func (tr *taskRepositoryimpl) UpdateTask(c context.Context, id string, updatedTask domain.Task) error {
	collection := tr.db.Collection(tr.collection)
	updatedResult, err := collection.UpdateByID(c, id, bson.M{"$set": updatedTask})
	if err != nil {
		return errors.New("internal server error: " + err.Error())
	}

	if updatedResult.MatchedCount == 0 {
		return errors.New("item not found")
	}

	return nil

}

func (tr *taskRepositoryimpl) DeleteTask(c context.Context, id string) error {

	collection := tr.db.Collection(tr.collection)

	deletedResult, err := collection.DeleteOne(c, bson.M{"_id": id})
	if err != nil {
		return errors.New("internal server error: " + err.Error())
	}
	if deletedResult.DeletedCount == 0 {
		return errors.New("item not found")
	}

	return nil
}

func (tr *taskRepositoryimpl) AddTask(c context.Context, task *domain.Task) error {
	collection := tr.db.Collection(tr.collection)
	objectId := primitive.NewObjectID()

	task.ID = objectId.Hex()
	_, err := collection.InsertOne(c, task)
	if err != nil {
		return errors.New("internal server error: " + err.Error())
	}
	return nil
}
