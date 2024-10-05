package data

import (
	"context"

	task_model "github.com/ETjoel/task_managment_api/models/task_model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ETjoel/task_managment_api/database"
)

func GetTasks() ([]task_model.Task, error) {
	collection := database.GetCollection(database.TaskCollection)
	cur, err := collection.Find(context.TODO(), bson.M{})
	var tasks []task_model.Task
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var task task_model.Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func GetTaskById(id string) (task_model.Task, error) {
	collection := database.GetCollection(database.TaskCollection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task_model.Task{}, err
	}

	var task task_model.Task
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return task_model.Task{}, nil // No document found
		}
		return task_model.Task{}, err // Other errors
	}

	return task, nil
}

func UpdateTask(id string, updatedTask task_model.Task) (bool, error) {
	collection := database.GetCollection(database.TaskCollection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	updatedResult, err := collection.UpdateByID(context.TODO(), objectId, bson.M{"$set": updatedTask})
	if err != nil {
		return false, err
	}

	return updatedResult.ModifiedCount != 0, nil
}

func DeleteTask(id string) (bool, error) {
	collection := database.GetCollection(database.TaskCollection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return false, err
	}

	return deleteResult.DeletedCount != 0, nil
}

func AddTask(newTask task_model.Task) error {
	collection := database.GetCollection(database.TaskCollection)
	_, err := collection.InsertOne(context.TODO(), newTask)

	return err
}
