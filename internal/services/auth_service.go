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
	users   repository.UserRepository
	tukangs repository.TukangRepository
}

func NewAuthService(users repository.UserRepository, tukangs repository.TukangRepository) AuthService {
	return &authServiceImpl{
		users:   users,
		tukangs: tukangs,
	}
}

func (s *authServiceImpl) Register(req request.RegisterRequest) error {

	hashedPassword, _ := utils.HashPassword(req.Password)

	switch req.Role {

	// ===== REGISTER USER =====
	case "user":
		user := entity.User{
			Name:     req.FullName,
			Email:    req.Email,
			Password: hashedPassword,
		}
		return s.users.Create(&user)

	// ===== REGISTER TUKANG =====
	case "tukang":
		tukang := entity.Tukang{
			Name:     req.FullName,
			Email:    req.Email,
			Password: hashedPassword,
			Category: "",
			Bio:      "",
			Services: "",
			Rating:   0,
		}
		return s.tukangs.Create(&tukang)

	default:
		return errors.New("invalid role")
	}
}

func (s *authServiceImpl) Login(req request.LoginRequest) (string, error) {

	user, err := s.users.FindByEmail(req.Email)
	if err == nil {
		if utils.CheckPassword(req.Password, user.Password) {
			// kirim role = user
			token, _ := utils.GenerateJWT(user.ID, user.Email, "user")
			return token, nil
		}
	}

	tukang, err := s.tukangs.FindByEmail(req.Email)
	if err == nil {
		if utils.CheckPassword(req.Password, tukang.Password) {
			// kirim role = tukang
			token, _ := utils.GenerateJWT(tukang.ID, tukang.Email, "tukang")
			return token, nil
		}
	}

	return "", errors.New("invalid email or password")
}
