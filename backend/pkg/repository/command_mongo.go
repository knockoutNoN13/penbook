package repository

import (
	"context"
	"pentbook/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommandMongo struct {
	db *mongo.Database
}

func NewCommandMongo(db *mongo.Database) *CommandMongo {
	return &CommandMongo{db: db}
}

// Вместо return стоят заглушки для запуска проекта

func (r *CommandMongo) Create(command models.Command) (string, error) {

	collection := r.db.Collection(commandsCollection)

	res, err := collection.InsertOne(context.TODO(), command)
	if err != nil {
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (r *CommandMongo) GetAll() ([]models.GetAllResponse, error) {
	collection := r.db.Collection(commandsCollection)

	var commands []models.GetAllResponse
	commands, err := GetAll(collection)
	if err != nil {
		return nil, err
	}
	return commands, nil
}

func (r *CommandMongo) GetById(commandId string) (models.Command, error) {

	collection := r.db.Collection(commandsCollection)
	var command models.Command

	objectId, err := primitive.ObjectIDFromHex(commandId)
	if err != nil {
		return models.Command{}, err
	}

	res := collection.FindOne(context.TODO(), bson.M{"_id": objectId})
	if res.Err() != nil {
		return models.Command{}, res.Err()
	}

	err = res.Decode(&command)

	if err != nil {
		return models.Command{}, err
	}

	return command, nil
}

func (r *CommandMongo) Delete(commandId string) error {
	collection := r.db.Collection(commandsCollection)

	err := deleteById(collection, commandId)

	if err != nil {
		return err
	}

	return nil
}

func GetAll(collection *mongo.Collection) ([]models.GetAllResponse, error) {

	var commandsList []models.GetAllResponse
	cur, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var item models.Command
		err := cur.Decode(&item)
		if err != nil {
			return nil, err
		}
		var temp models.GetAllResponse
		temp.Id = item.Id
		temp.Name = item.Name
		commandsList = append(commandsList, temp)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(context.TODO())
	return commandsList, nil
}
