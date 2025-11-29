package service

import (
	"auth-service/model"
	"auth-service/repository"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(username string) (model.User, error) {
	return s.repo.Login(username)
}

func (s *AuthService) Register(user *model.User) (bool, error) {
	return s.repo.Register(user)
}
