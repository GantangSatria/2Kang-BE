package repository

import (
    "2Kang/internal/domain/entity"
    "gorm.io/gorm"
)

type TransactionRepository interface {
    Create(tx *entity.Transaction) error
    GetByID(id uint) (*entity.Transaction, error)
    GetByUser(userID uint) ([]entity.Transaction, error)
    GetByTukang(tukangID uint) ([]entity.Transaction, error)
    UpdateStatus(id uint, status entity.TransactionStatus) error
}

type transactionRepoImpl struct {
    db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
    return &transactionRepoImpl{db: db}
}

func (r *transactionRepoImpl) Create(tx *entity.Transaction) error {
    return r.db.Create(tx).Error
}

func (r *transactionRepoImpl) GetByID(id uint) (*entity.Transaction, error) {
    var t entity.Transaction
    if err := r.db.First(&t, id).Error; err != nil {
        return nil, err
    }
    return &t, nil
}

func (r *transactionRepoImpl) GetByUser(userID uint) ([]entity.Transaction, error) {
    var list []entity.Transaction
    if err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&list).Error; err != nil {
        return nil, err
    }
    return list, nil
}

func (r *transactionRepoImpl) GetByTukang(tukangID uint) ([]entity.Transaction, error) {
    var list []entity.Transaction
    if err := r.db.Where("tukang_id = ?", tukangID).Order("created_at desc").Find(&list).Error; err != nil {
        return nil, err
    }
    return list, nil
}

func (r *transactionRepoImpl) UpdateStatus(id uint, status entity.TransactionStatus) error {
    return r.db.Model(&entity.Transaction{}).Where("id = ?", id).Update("status", status).Error
}
