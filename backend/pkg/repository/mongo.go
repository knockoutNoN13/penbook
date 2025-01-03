package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

const (
	usersCollection    = "users"
	commandsCollection = "commands"
)

func NewMongoDB(cfg Config) (*mongo.Database, error) {
	usr := ""
	if cfg.Username != "" {
		usr = fmt.Sprintf("%s:%s@", cfg.Username, cfg.Password)
	}
	applyURI := fmt.Sprintf("mongodb://%s%s:%s", usr, cfg.Host, cfg.Port)

	if applyURI == "" {
		return nil, errors.New("no connection specified")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(applyURI))
	if err != nil {
		return nil, err
	}
	db := client.Database(cfg.DBName)

	return db, nil
}
