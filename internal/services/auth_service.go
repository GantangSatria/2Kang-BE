package services

import (
	"errors"

	"2Kang/internal/domain/entity"
	"2Kang/internal/domain/repository"
	"2Kang/pkg/dto/request"
	"2Kang/pkg/utils"
)

type AuthService interface {
	Register(req request.RegisterRequest) error
	Login(req request.LoginRequest) (string, error)
}

type authServiceImpl struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authServiceImpl{repo}
}

func (s *authServiceImpl) Register(req request.RegisterRequest) error {
	hashedPassword, _ := utils.HashPassword(req.Password)

	user := entity.User{
		Name:     req.FullName,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     entity.Role(req.Role),
	}

	// Create user
	if err := s.repo.Create(&user); err != nil {
		return err
	}

	// If role is tukang → create empty tukang record
	if req.Role == "tukang" {
		tukang := entity.Tukang{
			UserID:   user.ID,
			Category: "",
			Bio:      "",
			Services: "",
			Rating:   0,
		}
		if err := s.repo.CreateTukang(&tukang); err != nil {
			return err
		}
	}

	return nil
}


func (s *authServiceImpl) Login(req request.LoginRequest) (string, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("email not found")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, _ := utils.GenerateJWT(user)
	return token, nil
}
