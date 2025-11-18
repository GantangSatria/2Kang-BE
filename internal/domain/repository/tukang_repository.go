package repository

import (
	"2Kang/internal/domain/entity"
	"gorm.io/gorm"
)

type TukangRepository struct {
	DB *gorm.DB
}

func NewTukangRepository(db *gorm.DB) TukangRepository {
	return TukangRepository{DB: db}
}

func (r TukangRepository) GetTukangList(kategori string) ([]entity.Tukang, error) {
	var tukangs []entity.Tukang
	q := r.DB

	if kategori != "" {
		q = q.Where("kategori = ?", kategori)
	}

	if err := q.Find(&tukangs).Error; err != nil {
		return nil, err
	}

	return tukangs, nil
}

func (r TukangRepository) GetTukangDetail(userID uint) (*entity.Tukang, error) {
	var data entity.Tukang

	if err := r.DB.Where("user_id = ?", userID).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
