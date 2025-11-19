package services

import (
	"fmt"
	"2Kang/internal/domain/entity"
	"2Kang/internal/domain/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{UserRepo: r}
}

func (s *UserService) GetProfile(id uint) (*entity.User, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid user id")
	}

	user, err := s.UserRepo.GetProfile(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch profile: %w", err)
	}

	return user, nil
}

func (s *UserService) UpdateName(id uint, name string) error {
    if name == "" {
        return fmt.Errorf("name cannot be empty")
    }
    return s.UserRepo.UpdateName(id, name)
}
