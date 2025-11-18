package repository

import (
	"2Kang/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	CreateTukang(t *entity.Tukang) error
}

type userRepositoryImpl struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepositoryImpl) CreateTukang(t *entity.Tukang) error {
	return r.db.Create(t).Error
}

func (r *userRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Preload("TukangDetail").First(&user).Error
	return &user, err
}
