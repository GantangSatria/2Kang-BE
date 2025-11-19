package repository

import (
	"2Kang/internal/domain/entity"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetReviewsByTukangID(tukangID uint) ([]entity.Review, error)
}

type reviewRepoImpl struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepoImpl{db: db}
}

func (r *reviewRepoImpl) GetReviewsByTukangID(tukangID uint) ([]entity.Review, error) {
	var list []entity.Review
	err := r.db.Where("tukang_id = ?", tukangID).
		Order("created_at desc").
		Find(&list).Error

	return list, err
}
