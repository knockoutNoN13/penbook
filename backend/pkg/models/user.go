package models

type User struct {
	Id       string `bson:"_id"`
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}
