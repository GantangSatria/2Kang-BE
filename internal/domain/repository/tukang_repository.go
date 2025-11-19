package repository

import (
	"2Kang/internal/domain/entity"
	"gorm.io/gorm"
)

type TukangRepository interface {
	Create(t *entity.Tukang) error
	FindByEmail(email string) (*entity.Tukang, error)
	GetTukangDetail(id uint) (*entity.Tukang, error)
	GetTukangList(category string) ([]entity.Tukang, error)
	UpdateCategory(id uint, category string) error
	UpdateBio(id uint, bio string) error
	UpdateServices(id uint, services string) error

}

type TukangRepositoryImpl struct {
	db *gorm.DB
}

func NewTukangRepository(db *gorm.DB) TukangRepository {
	return &TukangRepositoryImpl{db: db}
}

// CREATE
func (r *TukangRepositoryImpl) Create(t *entity.Tukang) error {
	return r.db.Create(t).Error
}

// FIND BY EMAIL
func (r *TukangRepositoryImpl) FindByEmail(email string) (*entity.Tukang, error) {
	var t entity.Tukang
	err := r.db.Where("email = ?", email).First(&t).Error
	return &t, err
}

// GET DETAIL
func (r *TukangRepositoryImpl) GetTukangDetail(id uint) (*entity.Tukang, error) {
	var t entity.Tukang
	err := r.db.First(&t, id).Error
	return &t, err
}

// GET LIST
func (r *TukangRepositoryImpl) GetTukangList(category string) ([]entity.Tukang, error) {
	var list []entity.Tukang
	query := r.db

	// Filter by category
	if category != "" {
		query = query.Where("category = ?", category)
	}

	err := query.Find(&list).Error
	return list, err
}

func (r *TukangRepositoryImpl) UpdateCategory(id uint, category string) error {
    return r.db.Model(&entity.Tukang{}).
        Where("id = ?", id).
        Update("category", category).
        Error
}

func (r *TukangRepositoryImpl) UpdateBio(id uint, bio string) error {
    return r.db.Model(&entity.Tukang{}).
        Where("id = ?", id).
        Update("bio", bio).
        Error
}

func (r *TukangRepositoryImpl) UpdateServices(id uint, services string) error {
    return r.db.Model(&entity.Tukang{}).
        Where("id = ?", id).
        Update("services", services).
        Error
}
