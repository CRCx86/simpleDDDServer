package service

import (
	"simpleDDDServer/config"
	"simpleDDDServer/model"
	"simpleDDDServer/repository"
)

type UserService struct {
	config *config.Config
	userRepository *repository.UserRepository
}

func (s *UserService) FindAll() []*model.User {
	if s.config.Enabled {
		return s.userRepository.FindAll()
	}
	return []*model.User{}
}

func NewUserService(config *config.Config, repository *repository.UserRepository) *UserService {
	return &UserService{
		config:         config,
		userRepository: repository,
	}
}