package services

import (
	"2Kang/internal/domain/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{UserRepo: r}
}

func (s *UserService) GetProfile(userID uint) (interface{}, error) {
	return s.UserRepo.GetProfile(userID)
}
