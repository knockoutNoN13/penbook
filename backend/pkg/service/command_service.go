package service

import (
	"pentbook/pkg/models"
	"pentbook/pkg/repository"
)

type CommandService struct {
	repo repository.Command
}

func NewCommandService(repo repository.Command) *CommandService {
	return &CommandService{repo: repo}
}

func (s *CommandService) Create(command models.Command) (string, error) {
	return s.repo.Create(command)
}

func (s *CommandService) GetAll() ([]models.GetAllResponse, error) {
	return s.repo.GetAll()
}

func (s *CommandService) GetById(commandId string) (models.Command, error) {
	return s.repo.GetById(commandId)
}

func (s *CommandService) Delete(commandId string) error {
	return s.repo.Delete(commandId)
}
