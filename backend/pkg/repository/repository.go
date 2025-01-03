package repository

import (
	"context"
	"pentbook/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GetUser(username, password string) (models.User, error)
}

type Command interface {
	Create(command models.Command) (string, error)
	GetAll() ([]string, error)
	GetById(commandId string) (models.Command, error)
	Delete(commandId string) error
}

type Repository struct {
	Authorization
	Command
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
		Command:       NewCommandMongo(db),
	}
}

func deleteById(collection *mongo.Collection, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})

	if err != nil {
		return err
	}

	return nil
}
