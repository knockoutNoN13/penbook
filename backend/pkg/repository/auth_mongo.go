package repository

import (
	"context"
	"errors"
	"fmt"
	"pentbook/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthMongo struct {
	db *mongo.Database
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user models.User) (string, error) {
	collection := r.db.Collection(usersCollection)

	checkRes := collection.FindOne(context.TODO(), user.Username)
	if checkRes.Err() != mongo.ErrNoDocuments {
		return "", fmt.Errorf(fmt.Sprintf("user %v already exists", user.Username))
	}

	res, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).String()
	return id, nil
}

func (r *AuthMongo) GetUser(username, password string) (models.User, error) {
	collection := r.db.Collection(usersCollection)

	filter := bson.D{
		primitive.E{Key: "username", Value: username},
	}

	res := collection.FindOne(context.TODO(), filter)

	user := models.User{}
	err := res.Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		return models.User{}, errors.New("invalid password")
	} else if err != nil {
		return models.User{}, err
	}

	return user, nil

}
