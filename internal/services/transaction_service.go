package services

import (
    "errors"
    "time"
    "fmt"
    "2Kang/internal/domain/entity"
    "2Kang/internal/domain/repository"
)

type TransactionService struct {
    trxRepo repository.TransactionRepository
    // optionally you may want access to TukangRepository to validate tukang exists
    tukRepo repository.TukangRepository
}

func NewTransactionService(trx repository.TransactionRepository, tuk repository.TukangRepository) *TransactionService {
    return &TransactionService{trxRepo: trx, tukRepo: tuk}
}

// Create an order
func (s *TransactionService) CreateOrder(userID uint, reqCreate interface{}) (*entity.Transaction, error) {
    req := reqCreate.(map[string]interface{}) // we'll adapt in handler; simpler is to pass typed struct

    // We'll parse fields safer in handler; here assume values provided:
    tukangID := uint(req["tukang_id"].(float64))
    scheduledAt := req["scheduled_at"].(time.Time)
    address := req["address"].(string)
    paymentMethod := req["payment_method"].(string)
    price := req["price"].(float64)
    notes := ""
    if v, ok := req["notes"].(string); ok {
        notes = v
    }

    // Validate tukang exists
    if _, err := s.tukRepo.FindByID(tukangID); err != nil {
        return nil, errors.New("tukang not found")
    }

    trx := &entity.Transaction{
        UserID:        userID,
        TukangID:      tukangID,
        ScheduledAt:   scheduledAt,
        Address:       address,
        PaymentMethod: paymentMethod,
        Price:         price,
        Status:        entity.StatusPending,
        Notes:         notes,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
    }

    if err := s.trxRepo.Create(trx); err != nil {
        return nil, fmt.Errorf("create order failed: %w", err)
    }

    return trx, nil
}

func (s *TransactionService) GetOrderByID(id uint) (*entity.Transaction, error) {
    return s.trxRepo.GetByID(id)
}

func (s *TransactionService) GetOrdersByUser(userID uint) ([]entity.Transaction, error) {
    return s.trxRepo.GetByUser(userID)
}

func (s *TransactionService) GetOrdersByTukang(tukangID uint) ([]entity.Transaction, error) {
    return s.trxRepo.GetByTukang(tukangID)
}

func (s *TransactionService) UpdateStatus(id uint, status entity.TransactionStatus) error {
    // optional: validate status transition
    return s.trxRepo.UpdateStatus(id, status)
}
