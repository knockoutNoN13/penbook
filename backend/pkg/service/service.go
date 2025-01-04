package service

import (
	"pentbook/pkg/models"
	"pentbook/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user models.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
}

type Command interface {
	Create(models.Command) (string, error)
	GetAll() ([]models.GetAllResponse, error)
	GetById(commandId string) (models.Command, error)
	Delete(commandId string) error
}

type Service struct {
	Authorization
	Command
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Command:       NewCommandService(repos.Command),
	}
}
