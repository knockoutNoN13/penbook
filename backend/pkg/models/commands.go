package models

type Command struct {
	Id         string   `bson:"_id,omitempty"`
	Name       string   `bson:"name"`
	Descripton string   `bson:"description"`
	Template   string   `bson:"template"`
	Args       []string `bson:"args"`
}
