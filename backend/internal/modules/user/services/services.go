package services

import (
	"backend/internal/modules/user/domain"
	"backend/internal/modules/user/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.Repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.Repo.FindAll()
}
